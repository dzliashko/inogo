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

`Beginner — familiar with basic Go syntax; broader backend fundamentals are not yet established`

Study start:

`2026-08-15`

Study pace:

`Not fixed; progress by demonstrated understanding rather than calendar deadlines`

---

# Current status

## Current phase

`Phase 1 — Go foundations`

## Current topic

`Variables, constants, and basic types`

## Current exercise

`Exercise 003 — Values and zero values`

## Current project milestone

`M1 — Go fundamentals`

## Recommended next action

1. Learn declarations, assignment, short declarations, and zero values.
2. Practice basic `string`, `int`, and `bool` values in a small program.
3. Continue with conditions after demonstrating the concepts independently.

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

- Initial programming baseline assessment.
- Go environment and introductory toolchain workflow.
- Go program structure: module, package, imports, functions, and `main`.

---

# Topics in progress

- Variables, constants, and basic types.

---

# Exercises

## Exercise 001 — Toolchain workflow and reading compiler errors

**Phase:** Phase 0 — Baseline and development environment

**Topic:** Go toolchain and compiler diagnostics

**Project relevance:** Establishes the build, run, formatting, and diagnostic workflow used throughout the backend project.

**Result:** Completed with small conceptual correction

**Tests:** Passed (`go test ./...` and `go vet ./...`; no test files exist yet)

**What I understood:** Can format, run, and build a Go package and use compiler locations and messages to find a naming error.

**Problems encountered:** Initially described `go run .` as running one file and attributed the temporary binary to `go build`; corrected to package-level behavior and explicit build output.

**Hints required:** Small

**Important mistake:** Imprecise distinction between the temporary executable used by `go run` and the explicit output written by `go build -o`.

**Needs repetition:** No immediate repetition; reinforce during later exercises.

**Next useful variation:** Build and run a multi-file `main` package when program structure is introduced.

## Exercise 002 — Read and explain the current Go program

**Phase:** Phase 1 — Go foundations

**Topic:** Program structure and execution flow

**Project relevance:** Establishes the meaning of module, package, imports, functions, and the executable entry point before the backend grows into multiple files and packages.

**Result:** Completed with a small terminology correction

**Tests:** Not applicable (code-reading exercise)

**What I understood:** Identified the module path, package, imported package and function, program entry point, and correctly traced a loop with a boolean condition and counter.

**Problems encountered:** Described `main` generally as the starting point without distinguishing `package main` from `func main()`.

**Hints required:** None during the exercise

**Important mistake:** `package main` marks an executable package; `func main()` is its entry point.

**Needs repetition:** No immediate repetition; reinforce when the first additional package is created.

**Next useful variation:** Explain execution and dependency flow after the program is split across multiple files or packages.

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

**Status:** `Completed`

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

- Can write a small Go function using a slice, `range`, a condition, a counter, and a return value.
- Comfortable with basic terminal navigation commands.
- Understands the basic HTTP request/response distinction and common `GET`, `POST`, `404`, and `500` meanings.
- Can trace a loop and boolean condition accurately by hand.
- Understands the basic roles of module, package, import, and `func main()`.

---

# Concepts needing practice

- Precise terminology for functions and ordered collections.
- Systematic debugging based on observable evidence.
- Git staging area versus commit history.
- SQL fundamentals (not studied yet).

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

`Beginner — needs a more concrete evidence-first debugging process`

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

`Beginner — can solve a small specified task using basic Go syntax`

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

`Level 2 — Guided (initial estimate based on one small exercise)`

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

## Initial assessment — 2026-08-17

**Topics assessed:** programming fundamentals, small Go code, terminal, Git, HTTP, SQL, and debugging.

**Strengths:** independently implemented a correct formatted Go function using a slice, loop, condition, counter, and return value; comfortable with basic terminal commands; has a basic HTTP mental model.

**Needs work:** conceptual precision, manual code tracing, systematic debugging, Git staging/commit distinction, and SQL fundamentals.

**Practical result:** `countNonEmptyTitles` produced the expected result; `gofmt`, `go test ./...`, and `go vet ./...` passed (there are no test files yet).

**Recommended start:** complete Phase 0.2 together with Phase 1.1, then cover the remaining Phase 1 topics at beginner depth.

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

The student knows some Go syntax but should build a precise mental model from program structure and toolchain behavior rather than skip directly to backend features.

Keep this section concise.

Record only observations that should change future teaching.

---

# Next milestone

Complete **M1 — Go fundamentals** at beginner depth, starting with program structure.
