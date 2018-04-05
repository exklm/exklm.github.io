---
title:  "your db design at a glance with schemaSpy"
date:   2016-05-06 00:00:00
description: how to use schemaSpy to chart out the db schema & how that helps with db work in general
---

When trying to grasp new projects and understand the logic behind it, looking at how data is stored
can be incredibly helpful. In SQL, you would look at the tables, columns, constraints & the
relationships between them. This task can be quite daunting, however.

If your projects are complex,
it might involve looking at over 20 to 30+ tables, each with at least 10+ columns. The challenge
are even harder on a poorly designed, non-normalized database.

For me, [schemaSpy](http://schemaspy.sourceforge.net/), an open-sourced tool has proved to be an indispensable  tool
by providing a **graphical summary** of tables, columns & relationships
(in the form of foreign key constraint).

So I went from looking at text:

<a href="#" data-featherlight="/assets/images/2016-05-06-schemaSpy_text_view.png"><img src="/assets/images/2016-05-06-schemaSpy_text_view.png"/></a>

To this:

<a href="#" data-featherlight="/assets/images/2016-05-06-schemaSpy_graphical_view.png"><img src="/assets/images/2016-05-06-schemaSpy_graphical_view.png"/></a>

You can also zoom in to specific tables and look at tables related to it directly (1 degree of separation) or indirectly (2 degree of separation)

<a href="#" data-featherlight="/assets/images/2016-05-06-schemaSpy_2_degree_table.png"><img src="/assets/images/2016-05-06-schemaSpy_2_degree_table.png"/></a>

Which is a great thing to do when you want to understand what are linked together.

> our database design right now is terrifically messy :(

There are some caveats for ORM users, however.

For example, as ActiveRecord's associations are managed by the app,
it's not going to show up as foreign key constraint
(and a link between tables) on schemaSpy.
Similarly, uniqueness and other type of constraints
won't be shown if you don't have it in the SQL schema.

## Install and usage

I work mostly with Postgres, so you can <br/> [get a cached version of schemaSpy & postgres' JDBC driver from here](https://github.com/exklamationmark/dotfiles/tree/master/tools/schemaspy)

The most-updated version should be there on:

- <http://schemaspy.sourceforge.net/>
- <https://jdbc.postgresql.org/download.html>
