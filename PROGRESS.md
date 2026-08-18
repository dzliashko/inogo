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

`Functions`

## Current exercise

`Not assigned yet`

## Current project milestone

`M1 — Go fundamentals`

## Recommended next action

1. Learn multiple return values and why related results can be returned together.
2. Practice using both returned values at the call site.
3. Continue reinforcing clear function responsibilities and input validation.

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
- Variables, constants, basic types, zero values, and assignment forms.
- Conditions and control flow: comparisons, boolean expressions, branching, early returns, and `switch`.

---

# Topics in progress

- Functions.

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

## Exercise 003 — Values and zero values

**Phase:** Phase 1 — Go foundations

**Topic:** Variables, constants, basic types, and zero values

**Project relevance:** Feed metadata, article counters, URLs, and boolean state all require clearly typed values and predictable defaults.

**Result:** Completed

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Correctly used `var`, assignment, `:=`, `const`, `string`, `int`, and `bool`; explained zero values and the same-scope restriction on `:=`.

**Problems encountered:** One spelling error in a variable name and initially imprecise domain names; the spelling error was corrected independently.

**Hints required:** Small naming feedback

**Important mistake:** None affecting behavior.

**Needs repetition:** No immediate repetition; reinforce type choice and naming in later domain exercises.

**Next useful variation:** Use typed values in branching logic for feed and article state.

## Exercise 004 — Classify a feed by unread count

**Phase:** Phase 1 — Go foundations

**Topic:** Conditions and control flow

**Project relevance:** Backend code frequently classifies input and state using explicit, mutually exclusive rules and must handle boundary values correctly.

**Result:** Completed

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Correctly implemented an `if` / `else if` / `else` chain and handled the boundaries `-1`, `0`, `1`, `9`, and `10`.

**Problems encountered:** None affecting correctness; the next refinement is removing unnecessary `else` blocks after unconditional returns.

**Hints required:** None

**Important mistake:** None.

**Needs repetition:** No immediate repetition; reinforce boundary analysis in later validation exercises.

**Next useful variation:** Express the same behavior using early returns and explain why the branches remain mutually exclusive.

## Exercise 005 — Refactor branching with early returns

**Phase:** Phase 1 — Go foundations

**Topic:** Early returns and control flow

**Project relevance:** Early returns keep validation and backend decision logic flat and make invalid cases explicit.

**Result:** Completed

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Refactored a conditional chain without changing behavior and explained that `return` terminates the current function call, constraining values that reach later conditions.

**Problems encountered:** The initial submission omitted the requested control-flow explanation; it was supplied correctly after one prompt.

**Hints required:** Small

**Important mistake:** None affecting code behavior.

**Needs repetition:** No immediate repetition; reinforce early returns in future validation functions.

**Next useful variation:** Combine early returns with compound boolean expressions in feed validation.

## Exercise 006 — Boolean expressions for feed and article state

**Phase:** Phase 1 — Go foundations

**Topic:** Compound boolean expressions and short-circuit evaluation

**Project relevance:** Feed refresh eligibility and article visibility depend on combining several boolean states correctly.

**Result:** Completed with one correction

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Used `&&`, `||`, and `!`, verified boolean combinations, and explained left-to-right short-circuit behavior for an `&&` chain.

**Problems encountered:** Initially used `isRead || isStarred`, which treated a read article as satisfying the unread rule and rejected an unread unstarred article.

**Hints required:** Small conceptual hint

**Important mistake:** A positive boolean name was used directly where the requirement needed its negation.

**Needs repetition:** Yes, through boolean rules in a different domain context.

**Next useful variation:** Write truth tables before implementing compound authorization or filtering rules.

## Exercise 007 — Classify feed status with switch

**Phase:** Phase 1 — Go foundations

**Topic:** `switch` and control flow

**Project relevance:** Feed states and other finite backend states often need clear value-based branching with an explicit fallback.

**Result:** Completed with conceptual clarification

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Implemented an expression-based `switch` with `default` and distinguished exiting a `switch` from returning from a function.

**Problems encountered:** Initially attributed case termination to `return`, then reversed the roles of implicit switch termination and the function-level effect of `return`.

**Hints required:** Significant conceptual clarification

**Important mistake:** Confused leaving a selected `case`/`switch` with terminating the surrounding function.

**Needs repetition:** Yes, when `switch` branches perform work without returning.

**Next useful variation:** Use a `switch` whose cases update a value and then continue with code after the switch.

## Exercise 008 — Process article counts with for

**Phase:** Phase 1 — Go foundations

**Topic:** Three-clause `for`, `break`, and `continue`

**Project relevance:** Backend batch processing often limits the amount of work, skips invalid items, and tracks successfully processed items.

**Result:** Completed with corrections

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Used a three-clause `for`, bounded work with `break`, skipped invalid items with `continue`, and returned a counter representing processed items.

**Problems encountered:** Initially stopped before processing item 10 and returned `total - count`, which mixed examined and unexamined items. Initial predictions matched the faulty implementation rather than the requirements.

**Hints required:** Significant

**Important mistake:** Did not first define precisely what the counter represented and used the input total in a result after the loop intentionally stopped early.

**Needs repetition:** Yes, through another bounded-processing exercise with different boundary rules.

**Next useful variation:** Trace the counter and termination condition before running a condition-style retry loop.

## Exercise 009 — Retry loop with condition-style for

**Phase:** Phase 1 — Go foundations

**Topic:** Condition-style `for` and loop termination

**Project relevance:** Feed refresh jobs may retry failed work a bounded number of times and must always make progress toward termination.

**Result:** Completed

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Used a condition-style `for`, updated the attempt counter inside the loop, stopped on simulated success, and produced the expected results for invalid, successful, and exhausted-attempt cases.

**Problems encountered:** The termination explanation mentioned the counter increment but did not initially connect it explicitly to the finite upper bound in the loop condition.

**Hints required:** Small conceptual clarification

**Important mistake:** None affecting behavior.

**Needs repetition:** No immediate repetition; reinforce bounded termination when real retry behavior is introduced.

**Next useful variation:** Compare this loop with a three-clause `for`, then begin decomposing logic into small functions.

## Exercise 010 — Decompose feed refresh decision into functions

**Phase:** Phase 1 — Go foundations

**Topic:** Parameters, return values, scope, and function decomposition

**Project relevance:** Backend decision logic is easier to test and reuse when calculations and policy decisions have clear, separate responsibilities.

**Result:** Completed with a small conceptual correction

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Implemented two cooperating functions, passed values through parameters and return values, kept state local, and predicted all boundary-case results correctly.

**Problems encountered:** Described `canAttemptRefresh` only as checking whether attempts remain, omitting that it also checks whether the feed is enabled.

**Hints required:** Small terminology correction

**Important mistake:** None affecting behavior.

**Needs repetition:** No immediate repetition; continue naming complete function responsibilities precisely.

**Next useful variation:** Return a computed value together with a validity flag and handle both at the call site.

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
- Understands `var`, assignment, `:=`, `const`, and zero values for `string`, `int`, and `bool`.
- Can implement mutually exclusive conditional branches and verify boundary values.
- Understands that an unconditional `return` ends the current function call and makes a following `else` unnecessary.
- Can combine boolean conditions with `&&`, `||`, and `!` and explain basic short-circuit evaluation.
- Can implement a value-based `switch` with `default`.
- Can use a three-clause `for` with `break` and `continue` after correcting boundary and counter-semantics mistakes.
- Can implement a bounded condition-style `for` whose progress variable is updated in the loop body.
- Can split a calculation and a decision into cooperating functions and explain local function scope.

---

# Concepts needing practice

- Precise terminology for functions and ordered collections.
- Systematic debugging based on observable evidence.
- Git staging area versus commit history.
- SQL fundamentals (not studied yet).
- Distinguishing exit from a `switch` from exit from the surrounding function.

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
