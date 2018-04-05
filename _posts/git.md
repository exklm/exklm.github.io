---
title:  "A lightweight release model with git and docker"
date:   2016-04-07 00:00:00
description: keeping track of what's released without too much crappy process
---

> This is a personal blog. All opinions (and stupidity) are my own. <br/>
> They do not represent thoughts and strategies of my employer.

In this blog post, I would like to talk about a branching/release model at Viki for
our backend services. It allows for moderate tracking of what has been released, while trying to
stay the heck out of everyone's workflow.


## Context

This is what we are using at Viki for source control:

- **git** for code
- **github.com** for hosting code
- **hub.docker.com** for hosting, err... docker images

<br/>
We then deploy docker images with built-in source code to the servers.

## Guidelines

> **commit on master is safe** <br/>
> **once deployed, considered gone (no live changes)**

This made debugging and rolling back bad changes a lot easier.

There are still emergencies that require you to go cowboy mode & debug live on the server.
But we try to avoid these as much as we can, and commit changes as a hot fix.

<hr/>

## Source control

We have micro-services, so the repos never get really big. Git works fine here.

The branch model we have is a simplified version of
[this workflow](http://nvie.com/posts/a-successful-git-branching-model/), as we try not to
enforce too many process.

<a href="#" data-featherlight="/assets/images/2016-04-07-lightweight-release-model-with-git-and-docker_git_branching_model.png"><img src="/assets/images/2016-04-07-lightweight-release-model-with-git-and-docker_git_branching_model.png"/></a>

The only requirements are:

> **Plan for releases on a branch** <br/>
> **Squash feature branches before RC merge** <br/>
> This helps us to track what changed easily and encourage commiters to group
related changes together (less low-effort "update", "fixes" commit) <br/>
> **Master branch is protected, only project owner can merge** <br/>
> This encourage owners to review code and be more aware about their projects
(in case of emergency, our devops can still merge and release).

We use github's **protected branches** feature, which benefits us in 2 ways:

- Only commits that passed [drone](https://drone.io/) (our CI tool) tests can be deployed (assuming we put decent amount of tests in the drone build)
- Project owner are aware of what gets released

<br/>
<a href="#" data-featherlight="/assets/images/2016-04-07-lightweight-release-model-with-git-and-docker_github_protected_branch.png"><img src="/assets/images/2016-04-07-lightweight-release-model-with-git-and-docker_github_protected_branch.png" /></a>

<hr/>

## Docker image control

We set up **drone** so that it builds a new docker for every commits, which is named as ``org/service:git-hash``

We then deploy this image onto server, and (almost) never change the code for it.

<hr/>

## Useful things

> **git add -patch** <br/>
> very useful if you need to separate changes in a single file to smaller commits
