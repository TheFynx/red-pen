# Red Pen commands

Part of the red-pen skill. The core rules in `SKILL.md` apply here too.

When invoked with arguments, red-pen applies Edit mode to the target. The target can be repo files, an attached file, a link to a platform doc, or a connected service.

**The default is always preview.** Red-pen never writes, commits, or pushes back unless the user passes `--fix`. Preview generates a side-by-side comparison showing the original and the cleaned version.

## No arguments

When invoked as `/red-pen` with no arguments, present the user with choices:

- Clean documentation files in this project
- Clean code comments in this project
- Clean a specific file (if one is currently open or referenced)
- Clean a doc from a link or attached file

Ask the user to pick one. Use an interactive multiple-choice prompt if the assistant has one. Then start on the chosen task.

## Source detection

Red-pen accepts any of these as input:

| Input | Source type | How it reads |
|-------|------------|--------------|
| No source specified | Repo files | Glob the project for matching extensions |
| `--path docs/` | Repo files (scoped) | Glob the given path |
| `--source notion <page URL or title>` | Notion | Read via Notion connector |
| `--source confluence <page URL>` | Confluence | Read via Confluence connector |
| `--source gdocs <doc URL>` | Google Docs | Read via Google Drive connector |
| A Notion/Confluence/Google Docs link pasted directly | Auto-detect | Identify the platform from the URL, read via connector |
| An attached `.md`, `.txt`, `.adoc`, `.rst` file | Local file | Read the attachment |
| An attached `.docx` file | Word document | Read via docx skill |
| An attached `.pdf` file | PDF document | Read via pdf skill |

If a connector is needed but not connected, tell the user which connector to add. Do not guess or skip.

## `/red-pen docs`

Scan for documentation. Source depends on what is provided (see source detection above).

**For repo files**, match these extensions: `.md`, `.adoc`, `.asciidoc`, `.txt`, `.rst`, `.rdoc`.

Exclude: `vendor/`, `node_modules/`, `.git/`, `CHANGELOG.md`, `LICENSE`, generated files.

**Default behavior (no `--fix`).** Generate a side-by-side comparison. For repo scans with many files, group violations by file with line numbers, the violation, and the suggested fix. For single documents (attached files, platform links), produce a full before/after comparison doc.

**With `--fix`.** Apply the rewrites.

- Repo files: rewrite in place, stage with `git add`.
- Platform docs (Notion, Confluence, Google Docs): show the preview, wait for the user to approve it, then push the cleaned version back. See [Platform writes](#platform-writes).
- Attached files: generate and deliver the cleaned file.

**With `--fix --commit`.** Repo files only. Same as `--fix`, then commit with message `red-pen: clean docs` and a body listing files changed.

**With `--path <path>`.** Scope repo scans to the given directory or file.

## `/red-pen comments`

Scan code files for comments. Auto-detect languages from the repository: look at file extensions present in the project, match against known comment syntaxes.

**Language detection map:**

| Extensions | Comment syntax |
|------------|---------------|
| `.go` | `//`, `/* */` |
| `.rb`, `.sh`, `.bash`, `.zsh` | `#` |
| `.py` | `#`, docstrings (`"""` or `'''` as the first statement of a module, class, or function) |
| `.ts`, `.tsx`, `.js`, `.jsx` | `//`, `/* */` |
| `.rs` | `//`, `/* */`, `///`, `//!` |
| `.java`, `.kt`, `.scala` | `//`, `/* */` |
| `.c`, `.cpp`, `.h`, `.hpp` | `//`, `/* */` |
| `.cs` | `//`, `/* */` |
| `.swift` | `//`, `/* */` |
| `.lua` | `--`, `--[[ ]]` |
| `.ex`, `.exs` | `#` |

Skip extensions not present in the project.

Exclude: `vendor/`, `node_modules/`, `.git/`, generated files, build output directories.

**Never touch these comments.** Some comments are instructions to a compiler, interpreter, linter, or legal reader. Changing one changes how the code builds, runs, or is licensed. Leave them exactly as they are, even if they break other red-pen rules:

| Kind | Examples |
|------|----------|
| Shebangs and interpreter lines | `#!/usr/bin/env bash`, `# -*- coding: utf-8 -*-`, `# frozen_string_literal: true` |
| Build and compiler directives | `//go:build`, `// +build`, `//go:generate`, `//go:embed`, `//go:noinline`, `#pragma`, `/// <reference path=...>` |
| Linter and type-checker suppressions | `//nolint`, `// eslint-disable`, `// @ts-ignore`, `// @ts-expect-error`, `# type: ignore`, `# noqa`, `# pylint: disable`, `// prettier-ignore`, `# rubocop:disable`, `#[allow(...)]` |
| Coverage and tool markers | `/* istanbul ignore next */`, `# pragma: no cover`, `// NOSONAR`, `# shellcheck disable=` |
| Code generation markers | `// Code generated ... DO NOT EDIT.`, `@generated`, region markers such as `// #region` |
| License and copyright headers | `// Copyright ...`, `SPDX-License-Identifier:`, any license block at the top of a file |
| Doc comments that tooling parses for behavior | `@deprecated`, `@param`, `@returns`, `@throws` tags in JSDoc, TSDoc, and Javadoc, and Go `Deprecated:` paragraphs |
| Commented-out code | Report it as a finding. Do not delete it or rewrite it. |

Rules for these comments:

- Do not reword, reformat, re-case, or move a directive. Many tools match the exact text and position. `//go:build` breaks if you add a space after `//`.
- For doc comments with tags, you can rewrite the prose. Keep every tag, its name, and its order.
- If you are not sure whether a comment is a directive, leave it and list it in the report as "skipped: possible directive."

**Default behavior (no `--fix`).** Report comments that violate red-pen rules. Group by file. For each violation, show the line number, the comment, and the rule broken. Rules: restates the obvious, too long, narrates the code, missing ticket on TODO.

**With `--fix`.** Rewrite or delete comments in place. Stage changed files. Show a summary.

**With `--fix --commit`.** Same as `--fix`, then commit with message `red-pen: clean comments` and a body listing files changed.

**With `--path <path>`.** Scope the scan to the given directory or file.

## `/red-pen <link>`

When a Notion, Confluence, or Google Docs link is passed directly:

1. Detect the platform from the URL.
2. Read the document via the platform's connector.
3. Generate a side-by-side comparison (default) or rewrite in place (`--fix`).

## Platform writes

A Notion, Confluence, or Google Docs page is shared. Other people read it, and the platform does not stage changes the way git does. `--fix` on a platform doc does not write right away:

1. Show the side-by-side preview.
2. Name the page and say that the cleaned version replaces its current content.
3. Wait for an explicit yes in the conversation. `--fix` alone is not approval to push.
4. Push the cleaned version through the connector.
5. Report the page link and what changed.

If working unattended, do not push. Write the cleaned version as a local file and report it as pending approval.

## Output format

The cleaned output matches the source format unless the user asks for something different.

| Source | Default output | User can request |
|--------|---------------|-----------------|
| `.md` file | `.md` file | PDF, Word |
| `.adoc` / `.rst` | Same format | PDF, Word |
| Google Doc | Google Doc (via connector) | PDF, Word, Markdown |
| Notion page | Notion page (via connector) | PDF, Word, Markdown |
| Confluence page | Confluence page (via connector) | PDF, Word, Markdown |
| `.docx` file | `.docx` file | PDF, Markdown |
| `.pdf` file | Markdown (PDF is read-only) | Word |

To request a different output format: `/red-pen docs --fix --format pdf`

Use the `pdf` skill for PDF generation and the `docx` skill for Word output.

## Processing order

1. Detect source type and read the content.
2. For repo files: glob for matching files.
3. Run the linter on the text. See [Automated lint](../SKILL.md#automated-lint) for how to run it and what to do when Go is missing.
4. Generate side-by-side comparison (default) or apply rewrites (`--fix`).
5. For `--fix` on repo files: stage changed files with `git add <file>`.
6. For `--fix` on platform docs: get approval, then push changes back via connector. See [Platform writes](#platform-writes).
7. For `--fix --commit` on repo files: commit all staged changes.
8. For `--fix --format <format>`: export to the requested format.
9. Report results: source, violations found, files changed (if `--fix`).

## Scan limits

Process files in batches. For large projects (over 50 matching files), ask the user before proceeding. Show the file count and estimated scope.

Do not rewrite code. Only rewrite comments and documentation content. The code stays exactly as it is.

Platform writes require the correct connector. If the connector is missing, generate the cleaned output as a local file and tell the user which connector to add for direct platform writes.

