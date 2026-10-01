# Agent Instructions

This repository is a fork of `ayn2op/discordo`. Keep changes small, upstream-friendly, and compatible with GPL-3.0.

## Go Style

- Prefer idiomatic Go over clever abstractions. Small explicit functions are better than framework-shaped indirection.
- Read nearby code before editing. Match existing package boundaries, model/update patterns, error style, and tests.
- Keep package APIs narrow. Do not export names unless another package actually needs them.
- Return errors with useful context and wrap underlying errors with `%w`.
- Avoid package-level mutable state unless the surrounding package already uses it for configuration or test seams.
- Use goroutines, channels, and contexts only when ownership, cancellation, and shutdown behavior are clear.
- Prefer table-driven tests for parsing, configuration, and edge-case behavior.
- Do not run `go mod tidy` or update dependencies unless the task explicitly involves dependency maintenance.

## Verification

Before claiming a Go change is complete, run:

```sh
make check
```

`make check` requires a failing format check, `go test ./...`, a linker-metadata build, and the smoke target. Smoke verifies version/distribution metadata and config loading.
On Windows, use the existing `pwsh ./scripts/check.ps1` fallback: it throws on gofmt output, runs `go test ./...` and `go build -trimpath -ldflags=-s .`, then invokes `scripts/smoke.ps1` for metadata/config checks. Inspect each native command exit status as well as the wrapper status; the script does not explicitly promote every native nonzero exit into a PowerShell exception. Do not substitute the previous three-command list, which omits smoke and does not fail on formatting output.

For CI or docs-only changes, run the closest relevant check and explain anything that was not applicable.

## Pull Requests

- Use the PR template to mark whether the change is upstreamable or fork-specific.
- Open generally useful fixes upstream when practical.
- Do not merge without explicit user authorization.
