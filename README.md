# Red Pen

Plain language rules, word swaps, and document templates for writing technical docs that do not waste the reader's time.

## What this is

A reusable writing skill for AI assistants (Claude, Cursor, Copilot) and a reference for human writers. It enforces plain language rules drawn from:

- **ASD-STE100** Simplified Technical English (sentence length limits, one word = one meaning)
- **FAA Order 1000.36A** (cut excess words, active voice, reader-first ordering)
- **DigitalOcean Product Docs** voice (calm, neutral, task-focused, no qualifiers or filler)
- **Wikipedia AI Cleanup** patterns (banned AI writing fingerprints)

## What it covers

| Section | What it does |
|---------|-------------|
| Sentence rules | 20-word max procedural, 25-word max descriptive |
| Voice and tone | No hype, no minimizing, no promises |
| Grammar | No contractions, active voice, limited modals (can/will/must) |
| Banned patterns | 13 AI and slop fingerprints to avoid |
| Word choice | 160+ word and phrase swaps from verbose to plain |
| Code and commands | How to present code blocks, placeholders, destructive commands |
| Examples | Labeling, input/output pairing, legends, method/function reference pattern |
| Diagrams | Mermaid rules, platform support, declaration order, type guidance |
| Formatting | Headings, lists, tables, whitespace |
| Templates | README, How-To, Reference/API, Checklist/Procedure, Changelog |
| Lint check | Go linter sweep, plus the manual self-check as the delivery gate |
| Self-check | 14-item checklist to run before delivery |

## Use as a Claude skill

Red Pen installs as an agent skill in Claude Code, OpenCode, Cursor, VS Code with Copilot, Copilot CLI, Codex CLI, and Gemini CLI (preview). See [INSTALL.md](INSTALL.md) for the location table, copy commands, and invocation notes per agent.

The skill triggers on any documentation writing or editing task.

## Use as a human reference

Read `SKILL.md` as a style guide. The `references/` files cover commands, templates, document structure, diagrams, and personal voice. Use `references/word-swaps.md` as a lookup table when editing.

`SKILL.md` holds only the rules that apply to all writing, so it stays short when an assistant loads it. The assistant reads a reference file only when the task needs it.

## Linter

A Go linter that checks files against the red-pen rules. No external dependencies.

Install:

```bash
go install github.com/TheFynx/red-pen/cmd/red-pen-lint@latest
```

Run:

```bash
red-pen-lint README.md docs/*.md
```

Or run directly from a clone:

```bash
go run ./cmd/red-pen-lint myfile.md
```

The linter checks prose and table cells and reports two severities:

Violations fail the run:

- Punctuation: em dashes, dashes typed as ` -- ` or unspaced `--`, en dashes used as a sentence dash, and semicolons.
- Words: banned modals, minimizing words, cheerleading, AI slop vocabulary, and promotional words.
- Patterns: formulaic transitions, compulsive summaries, bolded-term-colon definitions, and contractions.
- Sentence length: 20 words for numbered steps, 25 words for everything else.

Findings do not fail the run:

- Verbose words and phrases from `references/word-swaps.md`. The linter names the word and the line. Apply the swap, or keep the word when it names the technical thing, such as compiler or SQL optimization.

It skips YAML frontmatter, fenced and indented code blocks, inline code, and HTML.

Exit 0 = no violations, findings allowed. Exit 1 = violations found.

For a personal profile README or bio, `--voice personal` keeps the author's voice. It skips the contraction, modal, minimizing-word, semicolon, em-dash, and sentence-length checks:

```bash
red-pen-lint --voice personal README.md
```

To skip text that keeps a flagged word for a good reason, put `<!-- red-pen-ignore-next-line -->` on the line above it. Reasons include a technical use, a quoted error message, or a style guide that names the banned words. For several lines, wrap them in `<!-- red-pen-ignore-start -->` and `<!-- red-pen-ignore-end -->`.

## Files

```
SKILL.md                      # Core rules, lint, and self-check (loads every time)
INSTALL.md                    # Install paths and steps for each agent
AGENTS.md                     # Build, test, and self-lint entry points for agents
references/commands.md        # /red-pen commands, sources, comment safety rules
references/personal-voice.md  # Keep-my-voice rules for profiles and bios
references/doc-structure.md   # Reader's path, page scope, examples, method references
references/templates.md       # README, how-to, reference, checklist, changelog templates
references/diagrams.md        # Mermaid rules and diagram guidance
references/word-swaps.md      # Word and phrase replacement tables
cmd/red-pen-lint/             # Go linter source and tests
scripts/self-lint.sh          # Builds the linter, lints the documentation set
mise.toml                     # Local toolchain pins and tasks
hk.pkl                        # pre-commit, pre-push, and check hook definitions
.golangci.yml                 # Lint rules shared by hooks and CI
.goreleaser.yaml              # Release build, checksum, signing, and SBOM
.github/workflows/            # CI and release pipelines
examples/                     # Before/after samples showing the rules applied
```

## Development

Local tooling uses mise and hk, the same as the maintainer's other Go projects.

1. Run `mise install`. It installs the pinned toolchain, downloads the modules, and registers the git hooks.
2. Edit in any shape. The pre-commit hook runs gofmt, go vet, golangci-lint, and the docs self-lint. The pre-push hook runs the race tests and the security scans.
3. Run `mise run ci` or `hk run check` to run every local gate before a push.

CI runs the same gates on every push and pull request. It adds a docs gate that self-lints every Markdown file, a snapshot build of every release target, and a full snapshot of the release pipeline. Pushing a `v*` tag cuts the release through goreleaser: archives per platform, a checksums file, keyless cosign signatures, and SPDX SBOMs.

Keep the local hooks and CI equal in strength. If one adds a gate, add it to both. One quirk: hk's pre-push hook cannot diff a tag push against the remote. Cut releases with `git push --no-verify origin <tag>`. The tag triggers the release workflow, which is the gate that matters for a release.

## Before and after

The `examples/` directory has documentation samples in typical verbose AI style ("before") alongside the same content rewritten with red-pen rules ("after"). See what changes and why.

## License

[MIT](LICENSE).
