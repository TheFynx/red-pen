# Examples

Before/after pairs showing red-pen rules applied to real documentation.

Each "before" file is the original. Each "after" file rewrites the same content with red-pen rules. The technical facts stay the same. The words change.

## Examples

| Type | Before | After |
|------|--------|-------|
| Profile README | [before-profile-readme.md](before-profile-readme.md) | [after-profile-readme.md](after-profile-readme.md) |
| Getting started guide | [before-getting-started.md](before-getting-started.md) | [after-getting-started.md](after-getting-started.md) |
| Bootstrap process docs | [before-bootstrap.md](before-bootstrap.md) | [after-bootstrap.md](after-bootstrap.md) |
| GitHub Action README | [before-github-action-readme.md](before-github-action-readme.md) | [after-github-action-readme.md](after-github-action-readme.md) |
| eBook tool README | [before-ebook2audiobook-readme.md](before-ebook2audiobook-readme.md) | [after-ebook2audiobook-readme.md](after-ebook2audiobook-readme.md) |

## Sources and licenses

The before files come from public docs or the maintainer's own writing. The after files are red-pen rewrites of the same content. The facts stay with the source. The words change.

| Files | Source | License |
|-------|--------|---------|
| before/after-profile-readme.md | Maintainer's GitHub profile README | Maintainer's own |
| before/after-getting-started.md | Maintainer's Reeve project docs (unpublished) | Maintainer's own |
| before/after-bootstrap.md | Maintainer's RWR project docs (unpublished) | Maintainer's own |
| before/after-github-action-readme.md | [actions-template-sync](https://github.com/AndreasAugustin/actions-template-sync) | MIT |
| before/after-ebook2audiobook-readme.md | [ebook2audiobook](https://github.com/DrewThomasson/ebook2audiobook) | Apache-2.0 |

## Linting the examples

The technical "after" files pass the linter under the default docs voice:

```bash
go run ./cmd/red-pen-lint after-getting-started.md after-bootstrap.md
```

Profile READMEs use the personal voice. Lint them with `--voice personal`:

```bash
go run ./cmd/red-pen-lint --voice personal after-profile-readme.md
```

Under the docs voice, that file reports its kept contractions, which is expected. See [personal voice](../references/personal-voice.md).

## What to look for

<!-- red-pen-ignore-start -->
- Sentences split at the 20/25-word limit
- Semicolons replaced with periods
- "Should", "may", "might", "could" replaced with "can", "will", "must" or deleted
- Em dashes replaced with commas, periods, or named conjunctions
- "Simply", "just", "easily" deleted
- Filler transitions ("It is worth noting", "It is important to note") deleted
- Summaries at the end of sections removed
- Cheerleading phrases ("Great job!", "Awesome work!") removed
<!-- red-pen-ignore-end -->
