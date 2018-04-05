---
title:  "Building cross-platform subscription"
date:   2016-06-22 00:00:00
description: Working with 3rd party subscription engine (Android)
---

As part of offering better subscription experience, my company decided to
launch a new subscription engine in early 2016. It was supposed to unify subscription services
we had in Stripe, iTunes and Android Play. This series tell the story of how this happened,
and our choices/lesson.

In particular, this post is about building a fundamental block for cross-platform subscription
logic on top of Google Play's subscription engine.

> Jun 2016: this is work in progress, the approach seems good, so we are starting with implementation

# System overview

We can roughly break the unified subscription engine into logical parts like this:

<a href="#" data-featherlight="/assets/images/2016-05-22-building_cross_platform_subscription.png"><img src="/assets/images/2016-05-22-building_cross_platform_subscription.png"/></a>

We moved subscription logic from Stripe to in-house as a separate system.
After that, there is another logic layer that pulls data from both this web subscription engine,
as well as Android and iOS's. The data is then fed to a database, which is picked up by our public
Web API and analytics pipeline. Finally, user-facing features from web/mobile clients,
customer support portal,... make use of data from the Web API.

To completely offer cross-platform subscription capability, we need a better way to capture and
synchronize subscription status. This is basically capturing states and transitions in a user's
subscription, such as: **trial**, **paying**, **grace** (wait and retry after failed charge), **ended**.

# Challenges

The main challenges to create a cross-platform subscription system on top of Android's subscription
involves 2 things:

- Figuring out she undocumented (and unofficial) life cycle of the subscription from Google.
- Getting an accurate reading of subscription states from Google.

Let's go over each problem in more details:

## State of Google subscription

If you subscribe on Android, your subscription is managed by Google's billing engine.
However, this engine only allow you to pull the subscription's current state.
For example, after subscribing, you can use Android SDK or Google's publisher API
to know the start and end time of the subscription, from which you can tell whether the user
is currently a subscriber or not.

This works well enough for apps that only run on Android, but once you go cross-platform, there's
a need to know about state and state transitions of the Android subscription.
Sadly, there is no call-back mechanism from Google to tell you when things changes, so even
simple things like notifying failed payment, renewal and tracking historical changes becomes hard.

At the same time, not all user-action on the subscription goes through the mobile.
So we can't rely on our clients to send us state change events either.
For example, you can cancel the subscription in iTunes or Google Play store. 

## Getting accurate reading on Google's subscription state

Luckily, after poking around, we found a deterministic pattern of Android subscription's states.
With that, we decided to sample the subscription at various point in time,
checking with previous samples to then infer the current state of the subscription.
Though this model is undocumented, it (hopefully) won't change any time soon.

The next problem comes with the act of sampling. It's possible to miss out changes if we don't
sample at small enough intervals. Furthermore, sampling large number of subscription presents
the problem of rate-limiting and speed.

Finally, we found out that Google tends to run their subscriptions in batches, which create some
differences in subscription time. For example, a subscription starting at 1458092602407 (ms) UTC,
with 7 days of trial and 1 month in duration will not end at exactly 1461375802407, but might end
around 1461386605058, due to the batch job being ran later. The difference is usually within
1 day, which is fine from user perspective, but would make some calculations harder.

<a href="#" data-featherlight="/assets/images/2016-05-22-building_cross_platform_subscription_sample_google_sub_bad_date.png"><img src="/assets/images/2016-05-22-building_cross_platform_subscription_sample_google_sub_bad_date.png"/></a>

# Approach

## Sampling data

We call Android Publisher API to get something like this:

{% highlight bash %}
$> curl -sX GET 'https://www.googleapis.com/androidpublisher/v2/applications/com.viki.android/purchases/subscriptions/com.viki.vikipass.subscription.1month/tokens/userPurchaseToken?access_token=oauthToken'
{
 "kind": "androidpublisher#subscriptionPurchase",
 "startTimeMillis": "1458092602407",
 "expiryTimeMillis": "1469245392749",
 "autoRenewing": true,
 "priceCurrencyCode": "USD",
 "priceAmountMicros": "4380000",
 "countryCode": "US",
 "developerPayload": "",
 "paymentState": 1
}
{% endhighlight %}

When the Android subscription change states, the `expiryTimeMillis` value returned
from the sample call changes. By compare it with previous values, taking into account the
inaccuracy caused by the batch job, you can track what state it is in.

## Android subscription's life cycle

There are **5 different (1 dummy & 4 real) states** that an Android subscription goes through:

- **empty**     : dummy state, this is before the subscription is created
- **trial**     : when in-app subscription offer trial, and the user is getting it for the first time
- **subscribed**  : paying for benefits
- **grace**     : tried to pay earlier, but payment failed. However, it's still within a duration where we retry the payment
- **ended**     : ended by user action or because of failed payment

The state transitions that can happen are:

- **empty -> in_trial**         : this happen when user subscribe to an in-app product for the first time
- **empty -> subscribed**       : if user re-subscribed to an in-app product, they don't get the trial
- **trial -> ended**            : user can cancel their subscription during trial
- **trial -> grace**            : the first charge after trial fail, the user go into grace period with multiple retry on charge
- **trial -> subscribed**       : charge after trial is successful
- **subscribed -> ended**       : user stop renewing (cancel) during subscription period, it will stop at the end of the period
- **subscribed -> grace**       : the first renewal charge after the subscription period end fails, user go into grace
- **subscribed -> subscribed**  : this is a successful renewal
- **grace -> ended**            : user can also cancel while they are in grace period
- **grace -> grace**            : grace period consists of a number of days, each day, the `expiryTimeMillis` is extended by 1 day
- **grace -> subscribed**       : charing problem fixed, user move a subscription period starting from the first failed renewal

An example life cycle is sampled here:

<a href="#" data-featherlight="/assets/images/2016-05-22-building_cross_platform_subscription_sampling_google_sub.png"><img src="/assets/images/2016-05-22-building_cross_platform_subscription_sampling_google_sub.png"/></a>

## Sampling accuracy 

By now, keen readers would that there is a problem with this sampling: you only know when
the subscription change state **after** the fact. Plus, if your sampling frequency is too big, you might
end up missing out on changes altogether.

For now, we are trying to work with one sample every 12 hour for every active subscription we have.
Since the smallest interval between changes is 1 day, should be will be sufficient.

# After thoughts

Using the method of sampling, we built an abstraction on top of Google's subscription engine
(or a trick to peek into Google's system, depends on how you look). Other part of the unified
subscription engine can then build on top of this state data of user subscriptions.

However, due to the fact that sampling was used, we can only provide the data to a certain
accuracy level. My worry is mainly that when other people/systems call this abstraction, they
would start to mistake the states for **WHAT ACTUALLY HAPPENED**. This would be pretty frustrating
in debugging, where you based your guessing on inaccurate facts.

Perhaps more observation, as well as periodic evaluation of Google's subscription model is needed
to keep things going.
