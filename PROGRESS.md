# PROGRESS.md — Go Backend Learning Progress

## Student profile

Language:

`Go`

Primary goal:

`Become capable of independently developing backend services in Go and build an Inoreader-like RSS/news reader backend`

Target direction:

`Backend development`

Main learning project:

`Inoreader-like backend`

Starting level:

`Not assessed yet`

Study start:

`2026-08-15`

Study pace:

`Not fixed; progress by demonstrated understanding rather than calendar deadlines`

---

# Current status

## Current phase

`Phase 0 — Baseline assessment and Go environment`

## Current topic

`Initial assessment / Go toolchain and first program`

## Current exercise

`None assigned yet`

## Current project milestone

`M0 — Repository and learning baseline`

## Recommended next action

1. Assess current programming fundamentals.
2. Verify the Go development environment.
3. Confirm that a minimal Go program can be created, formatted, built, run, and tested.
4. Start Phase 1 at the appropriate depth based on the assessment.

---

# Target project scope

The learning project is a backend for an RSS/news reader inspired by Inoreader.

Planned MVP capabilities:

- users and authentication;
- feed subscriptions;
- RSS/Atom fetching and parsing;
- feed/article persistence in PostgreSQL;
- article deduplication;
- folders and tags;
- read/unread state;
- starred articles;
- article listing, filtering, sorting, and pagination;
- periodic feed refresh;
- background workers;
- retry/failure handling;
- automated tests.

Later capabilities may include:

- OPML import/export;
- search;
- feed discovery;
- automation/rules;
- caching;
- rate limiting;
- observability;
- deployment hardening.

---

# Current technical direction

Use these as the default direction unless the study plan is deliberately changed:

- Go;
- Go modules;
- Go standard library first;
- HTTP/JSON API;
- `net/http` before optional router/framework abstractions;
- PostgreSQL;
- explicit SQL before considering higher-level database abstractions;
- Go testing tools;
- Git;
- Docker later;
- modular monolith architecture initially.

Do not mark a technology as learned merely because it is listed here.

---

# Completed topics

None yet.

---

# Topics in progress

None yet.

---

# Exercises

No exercises completed yet.

For completed exercises, use this format:

## Exercise XXX — Exercise name

**Phase:**  
Phase name

**Topic:**  
Topic name

**Project relevance:**  
How this concept connects to the Inoreader-like backend.

**Result:**  
Completed / Completed with help / Needs revision

**Tests:**  
Passed / Partially passed / Not applicable

**What I understood:**  
Short note.

**Problems encountered:**  
Short note.

**Hints required:**  
None / Small / Significant / Full solution

**Important mistake:**  
Short note if relevant.

**Needs repetition:**  
Yes / No

**Next useful variation:**  
A different exercise that would test the same concept if repetition is needed.

---

# Project milestones

## M0 — Learning and repository baseline

**Status:** `Not started`

Goal:

- working Go environment;
- repository structure understood;
- baseline skill assessment completed;
- basic Go command workflow understood.

---

## M1 — Go fundamentals

**Status:** `Not started`

Goal:

- write small Go programs independently;
- use variables, control flow, functions, strings, slices, and maps;
- handle basic errors;
- write basic tests.

---

## M2 — Core Go design

**Status:** `Not started`

Goal:

- structs and methods;
- pointers;
- interfaces;
- packages/modules;
- error handling;
- basic project organization;
- testing and refactoring.

---

## M3 — First HTTP API

**Status:** `Not started`

Goal:

- build a small JSON API with `net/http`;
- understand handlers, methods, status codes, JSON, middleware concepts, validation, and tests.

---

## M4 — PostgreSQL persistence

**Status:** `Not started`

Goal:

- design a relational schema;
- connect Go to PostgreSQL;
- write explicit SQL;
- use migrations and transactions;
- test persistence.

---

## M5 — Feed subscription domain

**Status:** `Not started`

Goal:

- create users/subscriptions model;
- add/list/delete subscriptions;
- persist feed metadata.

---

## M6 — Feed fetching and parsing

**Status:** `Not started`

Goal:

- fetch feeds safely over HTTP;
- parse RSS/Atom;
- normalize entries;
- persist articles;
- deduplicate articles.

---

## M7 — Reading workflow

**Status:** `Not started`

Goal:

- article listing;
- pagination;
- filters;
- read/unread state;
- starred state;
- folders;
- tags.

---

## M8 — Background refresh and concurrency

**Status:** `Not started`

Goal:

- scheduled refresh;
- worker model;
- bounded concurrency;
- context cancellation;
- retries;
- graceful shutdown;
- race-safe behavior.

---

## M9 — Authentication and API hardening

**Status:** `Not started`

Goal:

- registration/login;
- password handling;
- authorization;
- validation;
- rate limiting basics;
- safer configuration and secrets handling.

---

## M10 — Search, import/export, and rules

**Status:** `Not started`

Goal:

- OPML import/export;
- article search;
- simple user-defined automation/rules.

---

## M11 — Production readiness

**Status:** `Not started`

Goal:

- structured logging;
- metrics;
- health endpoints;
- Dockerized environment;
- configuration;
- database backup/recovery awareness;
- performance investigation;
- deployment and operational documentation.

---

## M12 — Independent capstone

**Status:** `Not started`

Goal:

Independently design and implement a meaningful new backend capability without Codex generating most of the solution.

---

# Projects

## Main project — Inoreader-like backend

**Start date:** `2026-08-15`

**Status:** `Planned / not implemented yet`

**Architecture direction:**  
Start as a modular monolith.

**Primary persistence:**  
PostgreSQL.

**Primary API style:**  
HTTP/JSON.

**Current implemented features:**  
None yet.

**Current technical debt:**  
None yet.

**Major difficulties:**  
None recorded yet.

**Important lessons:**  
None recorded yet.

---

# Concepts understood well

None recorded yet.

---

# Concepts needing practice

None recorded yet.

---

# Go-specific concepts to watch

Do not mark these weak by default. Record them only when evidence appears.

Potential areas:

- value vs pointer semantics;
- slice behavior;
- map behavior;
- error handling;
- interfaces;
- package boundaries;
- `defer`;
- `context.Context`;
- goroutine lifetime;
- channels;
- synchronization;
- race conditions;
- HTTP handler lifecycle;
- SQL transactions.

---

# Recurring mistakes

None recorded yet.

Only add a recurring mistake after the same pattern appears more than once.

Possible examples:

- ignoring returned errors;
- returning an error without useful context;
- using pointers unnecessarily;
- accidental slice aliasing;
- nil map misuse;
- leaking HTTP response bodies;
- forgetting context cancellation;
- starting goroutines without a clear lifetime;
- unbounded concurrency;
- failing to check SQL errors;
- incorrect transaction handling;
- mixing HTTP, domain, and database logic in one large function;
- writing interfaces before there is a concrete boundary.

---

# Debugging skills

## Current level

`Not assessed`

Track progress in:

- reading compiler errors;
- reading test failures;
- reading panics and stack traces;
- following wrapped errors;
- inspecting HTTP requests/responses;
- inspecting logs;
- diagnosing SQL errors;
- isolating minimal reproductions;
- reasoning about concurrent failures;
- using the race detector later;
- testing hypotheses rather than guessing.

---

# Problem-solving skills

## Current level

`Not assessed`

Track:

- understanding requirements;
- defining acceptance criteria;
- breaking features into tasks;
- choosing appropriate data structures;
- writing pseudocode when useful;
- designing API behavior;
- designing database schemas;
- reasoning about failure paths;
- identifying edge cases;
- solving without AI assistance;
- reading documentation to unblock myself;
- evaluating tradeoffs.

---

# Backend design skills

## Current level

`Not assessed`

Track:

- separating responsibilities;
- choosing package boundaries;
- API design;
- data modeling;
- error boundaries;
- validation;
- transaction boundaries;
- dependency direction;
- concurrency design;
- operational concerns.

---

# Testing skills

## Current level

`Not assessed`

Track:

- writing a basic test;
- table-driven tests;
- boundary cases;
- testable design;
- HTTP handler tests;
- database integration tests;
- deterministic concurrent tests where feasible;
- understanding what should and should not be mocked.

---

# SQL/PostgreSQL skills

## Current level

`Not assessed`

Track later:

- schema design;
- constraints;
- joins;
- indexes;
- transactions;
- migrations;
- pagination;
- query performance basics;
- safe parameterized queries.

---

# Independence

Track how much assistance I require.

### Level 1 — Heavy assistance

Needs detailed instructions and substantial hints.

### Level 2 — Guided

Can solve problems with several hints.

### Level 3 — Mostly independent

Usually solves normal tasks alone but occasionally needs conceptual hints.

### Level 4 — Independent

Can design, implement, debug, and test normal junior backend tasks independently.

### Level 5 — Strong independence

Can approach unfamiliar backend problems, research documentation, evaluate tradeoffs, and defend design decisions independently.

Current independence level:

`Not assessed`

Target:

`Level 4–5 on junior Go backend tasks`

---

# Topics to repeat

None yet.

For each topic record:

- topic;
- reason for repetition;
- last practiced;
- observed mistake;
- suggested different exercise;
- review priority.

---

# Assessments

No assessments completed yet.

Use assessments at meaningful checkpoints instead of after every lesson.

For each assessment, record:

- date;
- phases/topics tested;
- strengths;
- weaknesses;
- tasks solved independently;
- tasks requiring hints;
- debugging performance;
- test-writing performance;
- recommended next steps.

---

# Learning observations

None yet.

Keep this section concise.

Record only observations that should change future teaching.

---

# Next milestone

Complete **M0 — Learning and repository baseline**, then begin **Phase 1 — Go foundations** at the level justified by the assessment.
