# AGENTS.md — Repository Context for AI Agents

## Repository Overview

**`gamekit-coffeemachine`** is a gamification toolkit for the [Coffee Machine kata](https://simcap.github.io/coffee-machine/). It provides language-agnostic tools—a **Progress Runner** and a **CLI utility**—that communicate with kata implementations (in Java, C++, Python, etc.) via an inter-process text protocol over Unix pipes. The actual kata implementations live in a separate companion repository (`kata-coffeemachine`).

- **Owner:** [Murex](https://github.com/murex)
- **Language:** Go (1.26+)
- **License:** MIT
- **Module path:** `github.com/murex/gamekit-coffeemachine`

## Project Structure

```
gamekit-coffeemachine/
├── main.go                  # Entry point — launches the CLI runner (cli.Run)
├── go.mod / go.sum          # Go module definition (Go 1.26, uses tool directives)
├── Makefile                 # Build, test, lint, release targets
├── .goreleaser.yml          # GoReleaser v2 config for cross-platform releases
├── .golangci.yml            # Linter config (golangci-lint v2 with revive, gosec, etc.)
├── progress-runner.sh       # Shell script to build & run progress tests against an implementation
│
├── cli/
│   └── runner.go            # CLI utility: interactive terminal to send commands to an implementation
│
├── process/
│   ├── coffee_machine_process.go       # Process lifecycle: spawn, run, send messages, read responses
│   ├── coffee_machine_process_test.go  # Tests for process management
│   ├── coffee_machine_message.go       # Message types (instructions) sent over the protocol
│   └── coffee_machine_message_test.go  # Tests for message formatting
│
├── progress/
│   ├── progress_test.go     # TestMain + Test_Progress: orchestrates all iteration tests
│   ├── iteration/
│   │   ├── test_runner.go   # Generic TestRunner: runs test cases sequentially per iteration
│   │   └── test_helper.go   # Assertion helpers (AssertDrinkIsServed, etc.)
│   ├── iteration1/          # Drink instructions, sugar, stick
│   ├── iteration2/          # Payment handling
│   ├── iteration3/          # Extra hot drinks, orange juice
│   ├── iteration4/          # Drink quantity reporting, total amount
│   └── iteration5/          # Empty tank handling, mail notifications
│
├── ref/
│   ├── drinks.go            # Drink definitions (Coffee, Tea, Chocolate, OrangeJuice) with prices
│   ├── liquids.go           # Liquid types (Water, Milk)
│   ├── tank_status.go       # Tank statuses (Full, Empty)
│   ├── amount.go            # AmountRegexp: locale-agnostic money formatting for assertions
│   └── amount_test.go
│
├── settings/
│   ├── build_info.go        # Build metadata (version, os, arch, commit, date, author) set via ldflags
│   └── build_info_test.go
│
├── dev-doc/
│   ├── README.md                       # Index of developer documentation
│   ├── inter-process-text-protocol.md  # Protocol specification (critical reference)
│   ├── build.md
│   ├── release.md
│   ├── debug.md
│   ├── quality-metrics.md
│   └── add-language.md
│
├── .github/
│   ├── workflows/
│   │   ├── go.yml              # CI: build & test on macOS, Ubuntu, Windows
│   │   ├── golangci_lint.yml   # CI: linting
│   │   └── go_releaser.yml     # CI: release pipeline
│   ├── CODEOWNERS
│   └── dependabot.yml
│
├── bin/                     # Build output directory (gitignored, populated by `make build`)
├── dist/                    # GoReleaser output directory (gitignored)
├── CONTRIBUTING.md
├── LICENSE.md
└── README.md
```

## Architecture & Key Concepts

### Inter-Process Text Protocol

Both tools communicate with kata implementations through a text-based protocol over stdin/stdout pipes. The toolkit spawns a child process (the language implementation's `run.sh`) and exchanges line-delimited messages.

**Protocol messages** (defined in `process/coffee_machine_message.go`):

| Instruction      | Direction | Purpose                                    | Response type   |
|------------------|-----------|--------------------------------------------|-----------------|
| `iteration`      | →         | Query which iteration is implemented       | Single line     |
| `restart`        | →         | Reset implementation state                 | Single line     |
| `shutdown`       | →         | Terminate the implementation process       | Single line     |
| `make-drink`     | →         | Request a drink instruction                | Single line     |
| `print-report`   | →         | Request a sales report                     | Multi-line (ends with `END-OF-REPORT`)  |
| `set-tank`       | →         | Set liquid tank status (empty/full)        | Single line     |
| `dump-mailbox`   | →         | Dump notification mailbox                  | Multi-line (ends with `END-OF-MAILBOX`) |

The `process.P` struct manages the child process lifecycle and message exchange with a 5-second response timeout.

### Progress Runner

The progress runner (`progress/`) is built as a **compiled Go test binary** (via `go test -c`). It:

1. Spawns the language implementation process.
2. Queries the implementation's current iteration number.
3. Runs black-box tests for iterations 1–5, skipping iterations beyond what the implementation reports.
4. Sends `restart` between test cases to reset state.
5. Sends `shutdown` on teardown.

Each iteration sub-package (`iteration1/` through `iteration5/`) defines test cases covering specific kata requirements:

- **Iteration 1:** Drink instructions, sugar, stick
- **Iteration 2:** Payment (enough/not enough money)
- **Iteration 3:** Extra hot drinks, orange juice
- **Iteration 4:** Drink quantity tracking, total revenue reporting
- **Iteration 5:** Empty liquid tanks, email notifications

### CLI Utility

The CLI (`cli/runner.go`) is an interactive terminal tool for manually sending protocol commands to an implementation. It reads user input from stdin, forwards it to the implementation process, and displays responses. Primarily used for debugging protocol integration when adding new language support.

### Reference Data (`ref/`)

Contains domain constants shared across the codebase:

- **Drinks:** Coffee (€0.60, code `C`), Tea (€0.40, code `T`), Chocolate (€0.50, code `H`), Orange Juice (€0.60, code `O`)
- **Liquids:** Water, Milk
- **Tank statuses:** Full, Empty
- **AmountRegexp:** Generates locale-agnostic regex patterns for monetary amounts (handles both `.` and `,` decimal separators)

## Build, Test, Lint & Release

All commands use the `Makefile`. The environment variable `CGO_ENABLED=0` is set globally.

| Command            | Description                                                       |
|--------------------|-------------------------------------------------------------------|
| `make build`       | Builds all binaries into `bin/`: `cli`, `progress-runner`, `gotestsum`, `test2json` |
| `make test`        | Runs all unit tests via `gotestsum`                               |
| `make lint`        | Runs `golangci-lint` (config in `.golangci.yml`)                  |
| `make modernize`   | Runs `gopls modernize` analysis                                   |
| `make vet`         | Runs `go vet`                                                     |
| `make tidy`        | Runs `go mod tidy`                                                |
| `make deps`        | Updates all dependencies                                          |
| `make prepare`     | Full pre-commit pipeline: deps → tidy → lint → modernize → build → test |
| `make clean`       | Removes `bin/`, `dist/`, `_test_results/`                         |
| `make release`     | Creates a release via GoReleaser                                  |
| `make snapshot`    | Creates a snapshot release (no publish)                           |

### Running Progress Tests Against an Implementation

```
./progress-runner.sh <path-to-language-implementation> [-v|-vv]
```

This script builds the toolkit, then runs the progress runner binary against the specified implementation directory. The last directory component of the path is used as the language name. Results are written to `_test_results/` as JUnit XML.

The `LANG_IMPL_PATH` environment variable is used internally to pass the implementation path to the progress runner process.

### Build Metadata

Build version, OS, architecture, commit, date, and author are injected at build time via `-ldflags` into the `settings` package variables. These are displayed when the tools start.

## Code Quality Standards

The project enforces strict linting via `golangci-lint` with the following notable rules:

- **Max function length:** 25 lines (revive `function-length`)
- **Max line length:** 120 characters (revive), 200 characters (lll)
- **Max arguments per function:** 4
- **Cognitive complexity limit:** 13
- **Cyclomatic complexity limit:** 16 (gocyclo)
- **Revive rules** are disabled for test files (`_test.go`)
- Security scanning via `gosec`

## CI/CD

Three GitHub Actions workflows run on push/PR:

1. **`go.yml`** — Build and test on macOS, Ubuntu, and Windows (Go 1.26)
2. **`golangci_lint.yml`** — Lint checks
3. **`go_releaser.yml`** — Cross-platform release (Linux/macOS/Windows, amd64/arm64)

## Testing Conventions

- Unit tests use `testify` (`assert` and `require` packages).
- The progress runner test binary uses `TestMain` for setup/teardown of the implementation process.
- Each iteration test runner sends a `restart` message before each test case.
- Test assertions use regex patterns to validate drink maker command format (e.g., `C:1:0` for coffee with 1 sugar and a stick).

## Key Files to Read First

1. `README.md` — High-level overview and architecture diagram
2. `dev-doc/inter-process-text-protocol.md` — Protocol specification (essential for understanding message flow)
3. `process/coffee_machine_message.go` — All message types and their wire format
4. `process/coffee_machine_process.go` — Process lifecycle and message exchange logic
5. `progress/progress_test.go` — How iterations are orchestrated
6. `ref/drinks.go` — Domain constants (drink names, prices, command codes)