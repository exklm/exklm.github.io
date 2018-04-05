---
layout: post
title: Getting buy-ins as an introvert
tags: [buy-ins]
---

If you are a software engineer, it's highly likely that you would be in situations where earlier design choices are questioned. Sometimes, as luck would have it, the conversation feels like an attack to what you decided on earlier. IMO, this indicates lack of buy-ins from other people (colleagues, management, etc).

Buy-ins, however, are not easy to get as an introvert.

### Identify the root cause

While it's easy to blame it on politics or people being difficult, lack of buy-ins might have deeper cause. Especially if you feel that your colleagues are reasonable human beings with good intentions.

I would attribute the root cause to 2 things:

- The technical requirements are not detailed and clear enough
- People forget things, and sometimes invent details on their own

Unclear requirements is probably the main issue. However, there's not much one can do about it, since most of the time we are designing solutions to new problems (to us, at least), the requirements will change as we understand things better.

This effect is amplified when multiple people are involved in designing the solution, since "problem cache" are usually cleared, or worse, recreated unfaithfully after a few days.

### Tackling it, the introvert way

I have seen people who are good at handling this. Usually they have a lot of energy and would be able to get people in a room, walk through decisions and end it with a positive note where everyone agrees. And they **retain** the session in their head, so if somebody come up and dispute, the objector's opinion/group concensus is quoted and things get resolved.

This doesn't work for me, though, due to the amount of energy required to organize all the people-related things (which eats up the coding time). Instead, I like this approach more:

- Spend more time working out the design, its goal, the trade-offs and why you do so. Discussion in person if you need to.
- Write down the ideas you have in mind, as concrete as possible. End-to-end examples are useful here. For example, I can write:

{% highlight bash %}
the service will receive request to endpoint X with payload Y
it should produce a response like Z
{% endhighlight %}

If this is a product design, you will need a convincing user story.
Which means you need to explain why your product decision benefits the user.
It's also better if you can contrast it with other possibilities/compare with existing products. For example:

{% highlight bash %}
We should present a line graph instead of bar graph, because the data
to be shown are supposed to be analog in nature.

Plus, if we observe data over a large interval, there might be
too many points and a bar graph will look more messy...
{% endhighlight %}

Then, you can point dispute and discussions to the written docs.

The caveat here are people who like to read short, **tl;dr-ish** description, who will skip over the docs. I haven't found a more effective approach than asking them to read the docs and give them a convenient URL, though.
