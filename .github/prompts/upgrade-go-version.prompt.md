---
agent: 'agent'
description: 'Upgrade Go version across the entire project (go.mod, CI workflows, linter config, documentation)'
---

# Upgrade Go Version

You are tasked with upgrading the Go version used in this project.

## Input

The user must provide the **target Go version** (e.g. `1.27`). If they haven't, ask them before proceeding.

## Steps

### 1. Identify the current Go version

Search for the current Go version in `go.mod` (the `go` directive) to confirm the starting point.

### 2. Find the latest compatible golangci-lint version

Search the web for the latest `golangci-lint` release that supports the target Go version.
The golangci-lint changelog is available at https://golangci-lint.run/docs/product/changelog/.
Look for a release entry that mentions support for the target Go version (e.g. "go1.27 support").
Use that version (without patch number, e.g. `v2.9`) as the golangci-lint version to set in CI.

### 3. Update all files

Update the Go version in **all** of the following files:

| File | What to update |
|------|----------------|
| `go.mod` | The `go` directive (set to the new major.minor version, e.g. `go 1.27`, dropping any previous patch version) |
| `.golangci.yml` | The `run.go` field |
| `.github/workflows/go.yml` | The `go-version` value in the `actions/setup-go` step |
| `.github/workflows/go_releaser.yml` | The `go-version` value in the `actions/setup-go` step |
| `.github/workflows/golangci_lint.yml` | The `go` version in the strategy matrix **AND** the `version` field of the `golangci/golangci-lint-action` step (set to the latest compatible golangci-lint version found in step 2) |
| `dev-doc/build.md` | The Go version mentioned in the "Install Go SDK" section |
| `AGENTS.md` | **All** references to the Go version (language version, go.mod comment, CI/CD section) |

### 4. Verify no old references remain

After making all edits, use `grep` to search for the **old** Go version pattern (e.g. `\b1\.26\b`) across `**/*.yml`, `**/*.md`, and `**/go.mod` files. Ignore matches found in `go.sum` or in third-party dependency paths — those are unrelated.

If any stale references are found, update them as well.

### 5. Remind the user about post-upgrade steps

After all files are updated, remind the user to run the following commands:

- `make tidy` — to update `go.sum` and resolve the `go.mod` toolchain directive
- `make prepare` — to validate that everything builds, passes lint, and tests correctly with the new versions