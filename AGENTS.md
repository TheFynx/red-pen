# AGENTS.md

Context for agents working in this repository.

## What this is

The red-pen project: a writing skill plus a Go linter. Read `README.md` for scope, `INSTALL.md` for setup, and `SKILL.md` for the rules the linter checks.

## Build and test

The linter source lives in `cmd/red-pen-lint/`. After changing Go code:

```bash
go test ./...
go vet ./...
gofmt -l ./cmd
```

`gofmt -l` must print nothing.

## Setup

1. Run `mise install`. It installs the pinned Go toolchain, golangci-lint, hk, and the security gate tools, then registers the git hooks.
2. Run `hk run check` before finishing any change. It covers every local gate: format, Go lint, docs self-lint, tests, and security scans.
3. GitHub Actions runs the same gates on every push and pull request. A `v*` tag publishes a release through goreleaser.

Keep the local hooks and CI equal in strength. A gate that exists on only one side blocks contributors without warning.

## Documentation style

Every Markdown file in this repo passes the linter. The script `scripts/self-lint.sh` builds the linter and runs it over the whole documentation set. Run it before finishing a doc change:

```bash
./scripts/self-lint.sh
```

The `examples/before-*` files are intentional slop. The script skips them. The profile example is a personal-voice sample, and the script lints it with `--voice personal`.

Fix every violation. Findings are suggestions. Decide each one, then move on.