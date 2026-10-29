#!/usr/bin/env bash
# Self-lint: build the linter, then run it over the documentation set.
# Docs voice covers every Markdown file. The profile example runs under the
# personal voice. The examples/before-* files are intentional slop and stay out.
set -euo pipefail

root="$(cd "$(dirname "$0")/.." && pwd)"
cd "$root"

dir="$(mktemp -d)"
trap 'rm -rf "$dir"' EXIT

go build -o "$dir/red-pen-lint" ./cmd/red-pen-lint

"$dir/red-pen-lint" \
  SKILL.md README.md INSTALL.md AGENTS.md \
  references/*.md examples/README.md \
  examples/after-getting-started.md examples/after-bootstrap.md \
  examples/after-github-action-readme.md examples/after-ebook2audiobook-readme.md

"$dir/red-pen-lint" --voice personal examples/after-profile-readme.md
