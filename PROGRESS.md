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

`Phase 2 — Core Go data handling`

## Current topic

`Maps`

## Current exercise

`Not assigned yet`

## Current project milestone

`M1 — Go fundamentals`

## Recommended next action

1. Build a mental model of maps as key/value storage.
2. Learn creation, insertion, lookup, and the comma-ok form.
3. Practice mapping feed identifiers to unread counts.

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
- Loops: three-clause and condition-style `for`, `break`, `continue`, counters, and bounded termination.
- Functions: parameters, single and multiple return values, scope, decomposition, and named versus unnamed results.
- Phase 1 checkpoint: independently combined variables, branching, a bounded loop, helper functions, multiple returns, and basic compiler-error reasoning.
- Arrays and slices: arrays conceptually; slice length, capacity, indexing, append, range, copying, backing arrays, nil/empty behavior, filtering, and independent results.

---

# Topics in progress

- Maps.

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

## Exercise 011 — Return remaining attempts with a validity flag

**Phase:** Phase 1 — Go foundations

**Topic:** Multiple return values and result interpretation

**Project relevance:** Backend calculations often need to return a value together with information about whether that value is valid; this prepares for Go's common value-and-error pattern.

**Result:** Completed with conceptual correction

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Returned and received an `(int, bool)` pair, used the validity flag to distinguish an exhausted retry budget from invalid input, and mapped both states to different caller-level results.

**Problems encountered:** Initially answered questions about the helper's `(int, bool)` results using strings returned by the calling function, then correctly traced the exact pairs and their interpretation.

**Hints required:** Small specific hint

**Important mistake:** Confused a helper function's return values with the higher-level result produced by its caller.

**Needs repetition:** Yes, reinforce tracing values across function-call boundaries in the cumulative functions exercise.

**Next useful variation:** Trace multiple return values through another caller using a different domain rule.

## Exercise 012 — Refactor named results into explicit returns

**Phase:** Phase 1 — Go foundations

**Topic:** Named and unnamed result parameters

**Project relevance:** Explicit return values make validation and configuration functions easier to trace as backend logic grows.

**Result:** Completed after one revision

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Identified the zero values of named result parameters, explained bare returns, preserved behavior while changing to unnamed results, and ultimately returned explicit value pairs without temporary result variables.

**Problems encountered:** The first refactor retained local `interval` and `valid` variables even though every branch already knew the final pair to return.

**Hints required:** Specific hint

**Important mistake:** Initially changed the signature and return syntax without fully applying the requested simplification.

**Needs repetition:** Yes, reinforce direct returns and removing unnecessary `else` branches after unconditional returns.

**Next useful variation:** Use explicit returns and early exits in a cumulative function-decomposition exercise.

## Exercise 013 — Build a refresh plan from helper results

**Phase:** Phase 1 — Go foundations

**Topic:** Cumulative functions exercise: decomposition, multiple returns, validation, scope, and early returns

**Project relevance:** Higher-level backend operations compose small validation and calculation helpers while preserving the meaning of each returned value.

**Result:** Completed after multiple revisions

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Composed two helpers, preserved validation precedence, used explicit early returns, and ultimately distinguished a valid exhausted state from invalid input using validity flags.

**Problems encountered:** Initially merged helper results with the caller's result, then treated a valid zero remaining-attempt count as invalid. Misleading boolean names and a redundant validation check remained through additional revisions.

**Hints required:** Significant

**Important mistake:** Used the computed value `attempts == 0` as an input-validity test even though the accompanying boolean already represented validity.

**Needs repetition:** Yes, reinforce valid zero values, precise boolean naming, and complete application of review feedback during the Phase 1 checkpoint.

**Next useful variation:** Solve a new value-plus-validity problem without hints and explain each function boundary before running it.

## Exercise 014 — Phase 1 foundations checkpoint

**Phase:** Phase 1 — Go foundations

**Topic:** Cumulative assessment of control flow, loops, functions, value validity, code tracing, and compiler diagnostics

**Project relevance:** Establishes that the language foundations needed for in-memory feed and article collections are sufficiently stable before moving to slices.

**Result:** Completed independently

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Correctly traced execution after a `switch`, explained an undefined-name compiler error, predicted all boundary cases, and independently implemented a helper-driven counting loop returning a valid zero distinctly from invalid input.

**Problems encountered:** No correctness problems. One answer about the post-`switch` instruction was initially broad but became precise after clarification.

**Hints required:** None for implementation; question clarification only

**Important mistake:** None.

**Needs repetition:** No immediate repetition; continue applying precise names and value-validity reasoning in later collection exercises.

**Next useful variation:** Process a slice of feed data using `range`, validation, and a returned count.

## Exercise 015 — Build and inspect a slice of feed titles

**Phase:** Phase 2 — Core Go data handling

**Topic:** Arrays versus slices, length, indexing, `append`, and `range`

**Project relevance:** Feed subscriptions and articles will be processed as variable-length collections, with safe indexing and iteration required throughout the backend.

**Result:** Completed with terminology clarification

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Created and appended to a string slice, used `len`, accessed first and last elements safely, iterated with `range`, and handled an empty slice with a validity flag.

**Problems encountered:** Initially described arrays and slices together and did not explain that indexing an empty slice would panic or that `append` specifically returns an updated slice.

**Hints required:** Small conceptual clarification

**Important mistake:** Imprecise distinction between a fixed-length array type and a slice type.

**Needs repetition:** Yes, reinforce the slice model while introducing capacity and backing arrays.

**Next useful variation:** Observe `len` and `cap` before and after `append`, then compare nil and empty slices.

## Exercise 016 — Observe slice length, capacity, and shared storage

**Phase:** Phase 2 — Core Go data handling

**Topic:** Slice length, capacity, half-open bounds, backing arrays, and `append` reallocation

**Project relevance:** Backend collection code must understand when slices share storage so that mutations and appends do not unexpectedly affect other views of feed or article data.

**Result:** Completed with significant conceptual clarification

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Created a slice with reserved capacity, observed length and capacity changes, demonstrated shared storage through a subslice, and showed how exceeding capacity separates the updated slice from the original backing array.

**Problems encountered:** Repeatedly calculated `len(base[:2])` as `3` before correctly distinguishing the two accessible elements from capacity `3`. Initial formatted output also omitted line breaks.

**Hints required:** Significant

**Important mistake:** Confused the length of a half-open slice range with the remaining capacity of its backing array.

**Needs repetition:** Yes, reinforce `[low:high]` as a half-open range and calculate `len` and `cap` independently.

**Next useful variation:** Compare nil and empty slices, then revisit subslice bounds in a filtering exercise.

## Exercise 017 — Compare nil and empty slices

**Phase:** Phase 2 — Core Go data handling

**Topic:** Nil slices, empty slices, comparison, iteration, and append behavior

**Project relevance:** Backend functions should accept zero-value slices safely and deliberately distinguish missing from explicitly empty collections only when the API semantics require it.

**Result:** Completed after one revision

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Distinguished nil and non-nil empty slices, observed their shared `len`, `cap`, `range`, and `append` behavior, compared each with `nil`, and used a function to classify their state.

**Problems encountered:** Initially said two slices had different types and described a slice as only a pointer. The first iteration counters stored the last index instead of counting iterations; the next revision retained an unnecessary index solely to print it.

**Hints required:** Significant conceptual and specific hints

**Important mistake:** Empty inputs made an incorrect iteration-counting implementation appear correct because both counters happened to remain zero.

**Needs repetition:** Yes, test counting logic with non-empty inputs and continue using `for range` when neither index nor value is needed.

**Next useful variation:** Compare shared storage after slice assignment with independent storage created using `copy`.

## Exercise 018 — Compare slice assignment with copy

**Phase:** Phase 2 — Core Go data handling

**Topic:** Slice assignment, shared backing storage, `make`, and the built-in `copy` function

**Project relevance:** Backend functions may need to return or retain independent collections so callers cannot accidentally mutate shared feed or article data.

**Result:** Completed after one correction

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Demonstrated shared storage after slice assignment, created independent storage with `make`, copied full and partial data, and observed that destination length rather than capacity limits `copy`.

**Problems encountered:** Printed `len(capacityOnly)` in place of `cap(capacityOnly)`, producing a plausible but incorrect observation that static tools could not detect. Initially attributed allocation of independent storage to `copy` rather than `make`.

**Hints required:** Small specific correction

**Important mistake:** Passed the wrong but type-compatible expression to formatted output, so compilation and static checks still succeeded.

**Needs repetition:** No immediate repetition; reinforce operation responsibilities when building a filtering function that allocates a result slice.

**Next useful variation:** Filter feed titles into a newly allocated result without mutating the input.

## Exercise 019 — Filter non-empty feed titles into a new slice

**Phase:** Phase 2 — Core Go data handling

**Topic:** Filtering, result allocation, `range`, `append`, and nil/empty result contracts

**Project relevance:** Article and feed endpoints frequently filter collections while preserving source data and returning predictable empty results.

**Result:** Completed after multiple revisions

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Filtered non-empty titles into an independently allocated slice, preserved order and input data, reserved capacity without creating zero-value elements, and returned non-nil empty results for nil and empty inputs.

**Problems encountered:** Initially used input capacity instead of length, then created a result with nonzero length and appended after its zero values. Repeatedly inspected or reassigned the nil input instead of storing and checking the function's direct return value; the empty-input call was omitted until specifically shown.

**Hints required:** Significant, including specific test-call structure

**Important mistake:** Confused reserving capacity with creating length, and tested an input/intermediate expression instead of the function result named in the acceptance criteria.

**Needs repetition:** Yes, independently construct a zero-length result with planned capacity and test direct return values in the cumulative slices exercise.

**Next useful variation:** Return an independent copy of the first N titles while treating zero as a valid empty request and invalid bounds separately.

## Exercise 020 — Return an independent prefix of feed titles

**Phase:** Phase 2 — Core Go data handling

**Topic:** Cumulative slices assessment: bounds, half-open slicing, valid empty results, `make`, `copy`, and storage independence

**Project relevance:** Backend pagination and limiting operations must validate bounds, distinguish empty from invalid results, and avoid exposing mutable internal storage.

**Result:** Completed after multiple revisions

**Tests:** Passed (`gofmt`, `go test ./...`, and `go vet ./...`; no test files exist yet)

**What I understood:** Correctly calculated half-open prefix length and capacity, validated bounds before slicing, created valid non-nil empty results, copied prefixes into independent storage, and verified that mutation did not affect the source.

**Problems encountered:** The first implementation copied directly from the full source, omitted nil-state and independence checks, used a separate zero-limit path, and renamed the required function. Initial explanations focused on internal pointers instead of the observable `(result, valid)` contract.

**Hints required:** Significant, including explicit missing acceptance criteria

**Important mistake:** Correct return values alone were treated as sufficient even though required slicing, observable-state checks, naming, and independence verification were missing.

**Needs repetition:** Yes, independently check every acceptance criterion before declaring later exercises complete.

**Next useful variation:** Use maps to associate feed identifiers with unread counts and distinguish a missing key from a stored zero value.

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

**Status:** `In progress`

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
- Can use multiple return values to distinguish a valid zero value from invalid input after a small correction.
- Understands named result zero values and bare returns, and can refactor them to explicit unnamed results after feedback.
- Can compose multiple helpers into a higher-level function while preserving their contracts after iterative review.
- Can independently combine a three-clause loop, helper function, validation, and an `(int, bool)` result while preserving a valid zero value.
- Can create, append to, index, and range over a slice while safely handling an empty slice.
- Can demonstrate shared backing-array behavior and reallocation after exceeding capacity following guided correction.
- Can distinguish nil and empty slices and use `len`, `cap`, `range`, comparison to `nil`, and `append` safely after correction.
- Can distinguish shared slice assignment from an independent destination created with `make` and populated using `copy`.
- Can filter a slice into independent storage with a deliberate non-nil empty-result contract after guided corrections.
- Can validate slice bounds, use half-open prefixes, create valid empty results, and return independent copies with `make` and `copy` after iterative review.

---

# Concepts needing practice

- Precise terminology for functions and ordered collections.
- Systematic debugging based on observable evidence.
- Git staging area versus commit history.
- SQL fundamentals (not studied yet).
- Distinguishing an iteration count from the last index when using `range`.
- Checking every acceptance criterion before declaring an exercise complete.

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
