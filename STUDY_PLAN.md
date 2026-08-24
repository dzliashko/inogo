# STUDY_PLAN.md — Go Backend Learning Roadmap

## Goal

This document defines my programming roadmap for becoming a **Go backend developer**.

The main learning project is an **Inoreader-like backend**: a service that subscribes to RSS/Atom feeds, collects articles, stores and organizes them, tracks user reading state, and refreshes content in the background.

Codex should adapt this roadmap to my actual progress.

Topics may be:

- added;
- removed;
- reordered;
- expanded;
- revisited.

Do not mark a topic as mastered merely because it was explained once.

A topic is considered sufficiently learned only when I can use it in exercises or project work with an appropriate level of independence.

---

# Current learning goal

Language:

`Go`

Main goal:

`Backend development`

Main project:

`Backend for an RSS/news reader inspired by Inoreader`

Current level:

`Not assessed yet`

Approximate study time:

`Not specified; advance by mastery rather than fixed calendar deadlines`

Architecture direction:

`Modular monolith first`

Core technical direction:

- Go standard library first;
- HTTP/JSON;
- PostgreSQL;
- explicit SQL;
- testing;
- Git;
- Docker later;
- concurrency only after the fundamentals are stable.

---

# Project outcome

By the end of the roadmap I should be able to independently build and explain a backend that supports a useful subset of an RSS reader.

## MVP target

The MVP should eventually include:

- user registration/login;
- RSS/Atom feed subscriptions;
- feed metadata;
- safe feed fetching over HTTP;
- feed parsing;
- normalized article storage;
- article deduplication;
- folders;
- tags;
- unread/read state;
- starred articles;
- article list endpoints;
- filtering, sorting, and pagination;
- scheduled feed refresh;
- background workers;
- retries and failure tracking;
- automated tests.

## Post-MVP target

After the core backend works:

- OPML import/export;
- full-text article search;
- feed discovery;
- simple automation/rules;
- rate limiting;
- caching when justified;
- structured logging;
- metrics;
- health checks;
- Dockerized local environment;
- deployment hardening.

## Not a priority initially

Do not expand the course early into:

- frontend development;
- mobile development;
- social network ingestion;
- newsletter email ingestion;
- recommendation systems;
- machine learning;
- microservices;
- Kubernetes.

---

# Phase 0 — Baseline and development environment

## 0.1 Initial assessment

Determine my current ability in:

- variables;
- conditions;
- loops;
- functions;
- collections;
- basic debugging;
- command line;
- Git;
- HTTP;
- SQL.

Do not assume I am a complete beginner and do not assume I already know programming.

Use a short combination of questions and small exercises.

Status:

`Completed`

---

## 0.2 Go environment and toolchain

Learn:

- what compiled code means;
- Go installation/toolchain concept;
- workspace vs module concept at a beginner level;
- `go version`;
- `go mod init`;
- `go run`;
- `go build`;
- `go test`;
- `gofmt`;
- basic terminal navigation.

Practice:

- create a small Go module;
- run a program;
- build it;
- deliberately cause and read a compiler error;
- format code.

Status:

`Completed`

---

# Phase 1 — Go foundations

## 1.1 Program structure

Learn:

- packages;
- `package main`;
- `import`;
- `func main`;
- statements;
- expressions;
- exported vs unexported names at a basic level.

Practice:

- very small command-line programs.

Status:

`Completed`

---

## 1.2 Variables, constants, and basic types

Learn:

- declarations;
- short variable declaration;
- assignment;
- zero values;
- constants;
- integers;
- floating point values;
- booleans;
- strings;
- basic conversions.

Practice:

- calculations;
- validation;
- small transformations.

Status:

`Completed`

---

## 1.3 Conditions and control flow

Learn:

- comparisons;
- boolean expressions;
- `if`;
- `else`;
- `switch`;
- early returns.

Practice:

- validation;
- branching logic;
- small decision-making programs.

Status:

`Completed`

---

## 1.4 Loops

Learn:

- Go's `for`;
- condition-style loops;
- range loops later when collections are introduced;
- `break`;
- `continue`.

Practice:

- counters;
- totals;
- searches;
- repeated processing.

Status:

`Completed`

---

## 1.5 Functions

Learn:

- parameters;
- return values;
- multiple return values;
- named vs unnamed results conceptually;
- scope;
- decomposition;
- small single-purpose functions.

Practice:

- split programs into functions;
- return value + error style later.

Status:

`Completed`

---

## Phase 1 checkpoint

Before continuing, I should be able to:

- read small Go programs;
- write small programs without copying a complete solution;
- use variables, branching, loops, and functions;
- explain compiler errors at a basic level.

Assessment:

`Completed`

---

# Phase 2 — Core Go data handling

## 2.1 Arrays and slices

Learn:

- arrays conceptually;
- slices as the normal dynamic sequence type;
- length and capacity;
- indexing;
- append;
- iteration;
- copying;
- backing array intuition;
- nil vs empty slice.

Practice:

- lists of article titles;
- filtering;
- transformations;
- deduplication exercises.

Status:

`Completed`

---

## 2.2 Maps

Learn:

- key/value storage;
- creation;
- zero-value behavior;
- lookup with existence check;
- insertion;
- deletion;
- iteration;
- maps as sets using an appropriate value type.

Practice:

- frequency tables;
- deduplication;
- grouping article data.

Status:

`In progress`

---

## 2.3 Strings, bytes, and Unicode basics

Learn:

- strings are bytes;
- UTF-8 concept;
- byte vs rune;
- indexing pitfalls;
- standard string operations;
- parsing and formatting.

Practice:

- normalize titles;
- detect empty/invalid fields;
- process URLs and text carefully.

Status:

`Not started`

---

## 2.4 Structs

Learn:

- modeling related data;
- struct fields;
- zero values;
- composite literals;
- nested structs;
- comparing data models.

Practice:

Create models such as:

- Feed;
- Article;
- User;
- Subscription.

Keep them simple and in-memory at this stage.

Status:

`Not started`

---

## 2.5 Methods and receivers

Learn:

- methods;
- value receivers;
- pointer receivers;
- behavior attached to types;
- when not to create methods.

Practice:

- small domain operations on feed/article models.

Status:

`Not started`

---

## 2.6 Pointers

Learn:

- addresses;
- dereferencing;
- passing values vs pointers;
- mutation;
- nil pointers;
- why pointers are used in Go;
- when values are simpler.

Do not overuse pointers.

Status:

`Not started`

---

## Phase 2 checkpoint

I should be able to model and manipulate in-memory feed/article data without AI writing most of the code.

Assessment:

`Not completed`

---

# Phase 3 — Errors, packages, interfaces, and tooling

## 3.1 Errors

Learn:

- `error`;
- returning errors;
- checking errors;
- creating errors;
- wrapping errors;
- preserving context;
- sentinel/type checking conceptually when needed;
- why errors are values.

Practice:

- parsing failures;
- validation failures;
- file/HTTP-like simulated failures.

Status:

`Not started`

---

## 3.2 `defer` and resource cleanup

Learn:

- execution timing;
- common cleanup pattern;
- files/connections/response bodies conceptually;
- common mistakes.

Status:

`Not started`

---

## 3.3 Packages and modules

Learn:

- package responsibility;
- imports;
- exported identifiers;
- module structure;
- dependency management;
- avoiding circular dependencies;
- keeping package APIs small.

Practice:

Split a small application into sensible packages.

Status:

`Not started`

---

## 3.4 Interfaces

Learn:

- implicit implementation;
- small interfaces;
- interface satisfaction;
- dependency boundaries;
- accepting behavior rather than concrete types;
- why premature interfaces are harmful.

Practice:

Introduce interfaces only where a real test/design boundary appears.

Status:

`Not started`

---

## 3.5 Go documentation and code navigation

Learn:

- reading package documentation;
- finding symbols;
- examples;
- reading function signatures;
- reading source when useful.

Goal:

Use documentation before asking AI to reconstruct an API from memory.

Status:

`Not started`

---

# Phase 4 — Testing, debugging, and code quality

## 4.1 Unit testing

Learn:

- `testing`;
- test functions;
- assertions using normal Go code;
- table-driven tests;
- subtests;
- deterministic tests.

Practice:

Test parsing/validation/domain functions.

Status:

`Not started`

---

## 4.2 Debugging

Learn:

- compiler errors;
- runtime panics;
- logical bugs;
- test failures;
- stack traces;
- minimal reproduction;
- hypothesis-driven debugging.

Status:

`Not started`

---

## 4.3 Refactoring

Learn:

- naming;
- duplication;
- extracting functions;
- reducing responsibilities;
- simplifying control flow;
- refactoring after tests.

Status:

`Not started`

---

## 4.4 Git fundamentals

Learn:

- repository;
- status;
- diff;
- add;
- commit;
- log;
- branches;
- merge concept.

Use Git throughout later project milestones.

Status:

`Not started`

---

## Phase 4 checkpoint

I should be able to create a small multi-package Go program, test it, debug it, and commit it.

Assessment:

`Not completed`

---

# Phase 5 — HTTP and web backend fundamentals

## 5.1 HTTP mental model

Learn:

- client/server;
- request/response;
- methods;
- paths;
- query parameters;
- headers;
- status codes;
- body;
- JSON;
- idempotency concept;
- statelessness concept.

Status:

`Not started`

---

## 5.2 Go `net/http`

Learn:

- server;
- handler;
- handler function;
- request;
- response writer;
- routing basics;
- server configuration;
- timeouts.

Build a small API using the standard library.

Status:

`Not started`

---

## 5.3 JSON APIs

Learn:

- encoding;
- decoding;
- request models;
- response models;
- validation;
- malformed input;
- consistent error responses.

Practice:

Create an in-memory feeds/articles API.

Status:

`Not started`

---

## 5.4 Middleware concept

Learn:

- cross-cutting behavior;
- logging;
- request IDs;
- authentication boundary later;
- recovery concept;
- why middleware order matters.

Do not build a large middleware stack yet.

Status:

`Not started`

---

## 5.5 HTTP testing

Learn:

- testing handlers;
- request construction;
- response recording;
- table-driven handler tests;
- API behavior as a contract.

Status:

`Not started`

---

## Project Milestone M3 — In-memory reader API

Build a small API that stores data only in memory.

Suggested capabilities:

- create a feed record;
- list feeds;
- create sample articles;
- list articles;
- mark an article read;
- star/unstar an article.

Purpose:

Practice HTTP and API design before adding a database.

Status:

`Not started`

---

# Phase 6 — SQL and PostgreSQL

## 6.1 Relational database foundations

Learn:

- table;
- row;
- column;
- data type;
- primary key;
- foreign key;
- unique constraint;
- nullability;
- normalization intuition.

Model:

- users;
- feeds;
- subscriptions;
- articles;
- user article state.

Status:

`Not started`

---

## 6.2 Core SQL

Learn:

- `SELECT`;
- `INSERT`;
- `UPDATE`;
- `DELETE`;
- `WHERE`;
- `ORDER BY`;
- `LIMIT`;
- joins;
- aggregates.

Practice with feed/article queries.

Status:

`Not started`

---

## 6.3 Indexes and query behavior

Learn:

- why indexes exist;
- tradeoff between reads and writes;
- uniqueness;
- common index use;
- basic query-plan intuition.

Do not optimize without evidence.

Status:

`Not started`

---

## 6.4 Transactions

Learn:

- atomicity;
- begin/commit/rollback;
- consistency of multi-step changes;
- transaction boundaries.

Use realistic article/subscription examples.

Status:

`Not started`

---

## 6.5 PostgreSQL from Go

Learn:

- connection;
- connection pool concept;
- context;
- parameterized SQL;
- scanning rows;
- errors;
- transaction API;
- resource cleanup.

Start with explicit SQL.

Status:

`Not started`

---

## 6.6 Migrations

Learn:

- why schema changes need versioning;
- forward migrations;
- rollback strategy conceptually;
- migration discipline.

Status:

`Not started`

---

## Project Milestone M4 — Persistent API

Replace in-memory storage with PostgreSQL.

Requirements:

- schema under version control;
- persistent feeds/articles;
- tests for important database behavior;
- clear error handling.

Status:

`Not started`

---

# Phase 7 — Domain modeling for the RSS reader

## 7.1 Users and subscriptions

Design:

- user;
- feed;
- subscription;
- uniqueness rules;
- subscription lifecycle.

Important question:

A feed is a global source; a subscription is the relationship between a user and that feed.

Understand why separating them can matter.

Status:

`Not started`

---

## 7.2 Article identity

Design:

- source identifier;
- canonical URL;
- publication time;
- fallback identity;
- uniqueness constraints;
- duplicate prevention.

Do not solve deduplication by title alone.

Status:

`Not started`

---

## 7.3 Per-user article state

Model:

- read/unread;
- starred;
- timestamps if useful.

Understand the difference between global article data and per-user state.

Status:

`Not started`

---

## 7.4 Folders and tags

Model:

- subscription folders;
- article tags;
- many-to-many relationships where needed.

Avoid over-generalizing the schema.

Status:

`Not started`

---

# Phase 8 — Feed fetching and parsing

## 8.1 HTTP client

Learn:

- outbound HTTP request;
- client reuse;
- context;
- timeout;
- status handling;
- headers;
- redirects conceptually;
- body closing;
- response-size limits;
- user agent.

Status:

`Not started`

---

## 8.2 XML fundamentals

Learn only what is needed:

- elements;
- attributes;
- namespaces conceptually;
- decoding XML in Go;
- malformed input.

Status:

`Not started`

---

## 8.3 RSS parsing

Start with a small known RSS sample.

Learn to extract:

- feed title;
- site URL;
- item title;
- item link;
- item GUID/id;
- publication date;
- description/content where available.

Status:

`Not started`

---

## 8.4 Atom and format normalization

Learn:

- why different formats need normalization;
- internal `Feed` and `Article` models;
- missing fields;
- date parsing variations;
- stable internal representation.

A third-party parser may be introduced only after the underlying problems are understood.

Status:

`Not started`

---

## 8.5 Fetch + parse + persist pipeline

Build the sequential version first:

1. load feed;
2. fetch URL;
3. parse response;
4. normalize entries;
5. deduplicate;
6. persist new articles;
7. update feed metadata;
8. record errors.

No goroutines yet unless prerequisites are complete.

Status:

`Not started`

---

## Project Milestone M6 — Real feeds

The backend can accept a real RSS/Atom URL, fetch it, parse it, store new articles, and avoid duplicate inserts.

Status:

`Not started`

---

# Phase 9 — API for the reading workflow

## 9.1 Article listing

Learn:

- filtering;
- sorting;
- stable ordering;
- query parameters;
- response shape.

Filters may include:

- subscription;
- folder;
- tag;
- unread;
- starred.

Status:

`Not started`

---

## 9.2 Pagination

Learn:

- offset pagination first if appropriate;
- limitations of large offsets;
- cursor/keyset pagination later;
- stable sort requirements.

Status:

`Not started`

---

## 9.3 Read/unread state

Implement:

- mark one article read/unread;
- bulk marking later;
- authorization checks.

Status:

`Not started`

---

## 9.4 Starred articles

Implement:

- star;
- unstar;
- list starred.

Status:

`Not started`

---

## 9.5 Folders and tags

Implement only after the schema and API behavior are understood.

Status:

`Not started`

---

## Project Milestone M7 — Usable reading API

A client can:

- subscribe to feeds;
- list feeds;
- browse articles;
- filter/paginate articles;
- mark read/unread;
- star articles;
- organize content.

Status:

`Not started`

---

# Phase 10 — Concurrency and background work

## 10.1 Concurrency foundations

Learn:

- goroutines;
- happens-before intuition;
- shared state;
- channels;
- mutexes;
- wait groups or current synchronization primitives when appropriate.

Start from sequential code.

Status:

`Not started`

---

## 10.2 Context and cancellation

Learn:

- cancellation;
- deadline;
- timeout;
- propagation;
- why context belongs at request/job boundaries;
- why context should not be stored casually in structs.

Status:

`Not started`

---

## 10.3 Worker model

Design a bounded feed-refresh worker system.

Learn:

- job queue concept;
- worker pool;
- bounded concurrency;
- backpressure;
- ownership of channels;
- shutdown.

Status:

`Not started`

---

## 10.4 Scheduling

Implement periodic feed refresh.

Think about:

- next fetch time;
- failed feed;
- disabled feed;
- different refresh intervals;
- avoiding duplicate simultaneous refresh.

Status:

`Not started`

---

## 10.5 Retries and backoff

Learn:

- transient vs permanent failure;
- retry budget;
- backoff;
- jitter conceptually;
- respecting remote servers;
- failure recording.

Status:

`Not started`

---

## 10.6 Race detection and concurrency tests

Learn to use tests and the race detector to validate concurrency assumptions.

Status:

`Not started`

---

## Project Milestone M8 — Background refresh

Feeds refresh automatically with bounded concurrency, cancellation, retries, and graceful shutdown.

Status:

`Not started`

---

# Phase 11 — Authentication, authorization, and security

## 11.1 Authentication fundamentals

Learn:

- identity;
- password storage;
- session/token concept;
- login lifecycle;
- logout/revocation considerations.

Do not implement homemade cryptography.

Status:

`Not started`

---

## 11.2 Authorization

Learn:

- authentication vs authorization;
- ownership checks;
- preventing access to another user's subscriptions/articles/state.

Status:

`Not started`

---

## 11.3 Input and HTTP hardening

Learn:

- validation;
- body limits;
- HTTP server timeouts;
- safe outbound requests;
- rate limiting concept;
- logging without secrets.

Status:

`Not started`

---

## 11.4 SQL and secret safety

Learn:

- parameterized queries;
- configuration;
- environment variables;
- credentials;
- secret leakage risks.

Status:

`Not started`

---

## Project Milestone M9 — Multi-user backend

Different users can securely manage their own feeds and reading state.

Status:

`Not started`

---

# Phase 12 — Import/export, search, and automation

## 12.1 OPML

Learn:

- purpose of OPML in feed readers;
- XML parsing/generation;
- import validation;
- duplicate subscriptions;
- export.

Status:

`Not started`

---

## 12.2 Search

Start simple.

Learn:

- search requirements;
- indexing tradeoffs;
- PostgreSQL full-text search concepts if appropriate;
- result ranking conceptually;
- filtering search results by user access.

Do not introduce a separate search engine without a real reason.

Status:

`Not started`

---

## 12.3 Automation/rules

Implement a deliberately limited rule model, for example:

- if title contains text -> add tag;
- if feed is X -> mark starred;
- if author matches -> add tag.

Learn:

- rule representation;
- validation;
- evaluation order;
- deterministic behavior;
- testing.

Status:

`Not started`

---

## Project Milestone M10 — Power-user features

The project supports OPML, useful search, and a small testable automation system.

Status:

`Not started`

---

# Phase 13 — Reliability and observability

## 13.1 Structured logging

Learn:

- log levels;
- structured fields;
- request/job correlation;
- useful errors;
- avoiding secrets.

Status:

`Not started`

---

## 13.2 Metrics

Learn useful backend metrics such as:

- HTTP request count/duration;
- feed fetch duration;
- successful/failed refreshes;
- queued jobs;
- new articles discovered;
- database latency conceptually.

Status:

`Not started`

---

## 13.3 Health and readiness

Learn the difference between:

- process is alive;
- service is ready to serve;
- dependency health.

Status:

`Not started`

---

## 13.4 Graceful shutdown

Coordinate:

- HTTP server shutdown;
- worker cancellation;
- in-flight jobs;
- database cleanup.

Status:

`Not started`

---

# Phase 14 — Performance and scalability fundamentals

## 14.1 Measure before optimizing

Learn:

- benchmark concept;
- profiling concept;
- database query measurement;
- latency vs throughput;
- allocation awareness.

Status:

`Not started`

---

## 14.2 Database performance

Learn:

- useful indexes;
- N+1 query pattern;
- pagination performance;
- connection pool basics;
- transaction duration.

Status:

`Not started`

---

## 14.3 Caching

Introduce only after a measurable or architectural reason exists.

Learn:

- cacheable data;
- invalidation;
- TTL;
- stale data tradeoffs;
- why caches add complexity.

Status:

`Not started`

---

## 14.4 Scaling thought exercise

Reason about how the system could evolve if it had:

- many users;
- many feeds;
- slow feeds;
- large article history;
- bursts of refresh work.

Do not automatically redesign it as microservices.

Status:

`Not started`

---

# Phase 15 — Delivery and production workflow

## 15.1 Configuration

Learn:

- development vs production config;
- environment;
- defaults;
- validation;
- secrets.

Status:

`Not started`

---

## 15.2 Docker

Learn:

- image;
- container;
- Dockerfile;
- build context;
- ports;
- environment;
- local PostgreSQL setup;
- multi-stage build concept when appropriate.

Status:

`Not started`

---

## 15.3 CI basics

Automate at least:

- formatting check;
- tests;
- vet/static checks as appropriate;
- build.

Status:

`Not started`

---

## 15.4 Deployment fundamentals

Understand:

- process lifecycle;
- migrations during deployment;
- logs;
- health checks;
- rollback concept;
- backups;
- basic TLS/reverse proxy concepts.

Status:

`Not started`

---

## Project Milestone M11 — Production-oriented release

Produce a documented version that another developer can run locally and that has a credible deployment path.

Status:

`Not started`

---

# Phase 16 — Independent capstone

Choose one meaningful extension.

Possible examples:

- robust feed discovery;
- advanced filtering;
- content retention policies;
- per-feed refresh policies;
- bulk article operations;
- improved search;
- API compatibility layer;
- notification subsystem;
- performance improvement backed by measurements.

Before implementation:

1. write requirements;
2. define acceptance criteria;
3. identify affected components;
4. propose schema/API changes;
5. identify risks;
6. write an implementation plan;
7. implement independently;
8. test;
9. document tradeoffs.

Codex should guide and review rather than build the feature for me.

Status:

`Not started`

---

# Review checkpoints

After every major phase or project milestone:

- give a short knowledge assessment;
- give several practical tasks without naming every concept being tested;
- include at least one debugging/code-reading task periodically;
- identify weak areas;
- update `PROGRESS.md`;
- revisit weak topics before advancing.

Do not use quiz performance alone.

Project implementation quality and independence matter more.

---

# Spaced repetition

Regularly revisit:

- errors;
- slices/maps;
- pointers;
- interfaces;
- package boundaries;
- HTTP status/error handling;
- SQL;
- transactions;
- context;
- concurrency;
- testing.

Prefer a new exercise in a different context instead of repeating an old solution.

---

# Architecture checkpoints

Before these milestones, require me to explain the design in my own words:

- first multi-package program;
- first HTTP API;
- first database schema;
- first feed ingestion pipeline;
- first worker pool;
- authentication;
- search/rules;
- production release.

Ask me to explain:

- responsibilities;
- data flow;
- failure paths;
- dependencies;
- tradeoffs.

---

# Graduation criteria

Do not consider the roadmap complete until I can independently:

- read unfamiliar Go code at junior-backend level;
- write idiomatic small and medium Go programs;
- model data using structs and appropriate types;
- use errors correctly;
- organize packages without unnecessary abstraction;
- write and test HTTP handlers;
- design a small REST-style HTTP/JSON API;
- write SQL and work with PostgreSQL;
- use transactions appropriately;
- fetch external HTTP resources safely;
- parse and normalize RSS/Atom data;
- reason about deduplication;
- implement background jobs;
- use goroutines and synchronization without obvious leaks/races;
- use contexts, deadlines, and cancellation appropriately;
- write unit and integration tests;
- debug compiler, runtime, HTTP, SQL, and concurrency problems;
- use Git in a normal development workflow;
- containerize and run the service locally;
- read official documentation to solve unfamiliar problems;
- explain the important parts of my own backend;
- add a non-trivial feature without AI generating most of the implementation.

The final goal is not memorizing Go syntax.

The final goal is being able to solve backend engineering problems independently.
