---
name: red-pen
argument-hint: "[docs | comments | file or link] [--fix]"
description: Write or rewrite technical documentation and code comments in plain, direct language. Use when writing or editing READMEs, how-to guides, reference docs, runbooks, changelogs, code comments, PR descriptions, commit messages, or personal profile READMEs and bios, or when the user runs /red-pen. Not for chat replies, marketing copy, blog posts, or fiction.
---

# Red Pen

Write text that a tired reader understands on one read. No filler. No qualifiers. No AI slop. State the fact, show the step, move on.

This skill applies to technical writing: documentation, code comments, PR descriptions, and commit messages. Documentation includes READMEs, how-tos, guides, reference docs, runbooks, and changelogs.

## Two modes

**Write mode (default).** Apply all rules below when writing new documentation or drafting from scratch.

**Edit mode.** When rewriting or editing existing docs, apply the same rules but preserve the author's technical claims, facts, citations, and argument structure. Change the words, not the meaning.

### Edit mode guardrails

These apply whenever you rewrite or restructure existing documentation. They are not optional.

**Never downgrade notices.** The original can use a callout, admonition, or alert block (`[!IMPORTANT]`, `[!WARNING]`, `[!CAUTION]`, `> **Warning:**`, or a platform equivalent). If it does, the rewrite must keep the same notice level or higher. Rewrite the text inside the notice for clarity. Do not convert it to plain text, move it out of the callout, or reduce its visual weight. Safety notices, legal notices, and deprecation warnings exist to stop the reader. Making them quieter is a defect.

**Never silently delete sections.** Red-pen rewrites words. It does not delete content the author chose to include. Some sections belong somewhere else: a roadmap in a README, a changelog in a getting-started guide, or a style guide in a troubleshooting page. To move one:

1. Identify the section and where it belongs (a separate file, a wiki page, a project board).
2. Draft the new location. Create the file or describe where it goes.
3. Replace the original section with a one-line reference linking to the new location.
4. Present the plan to the user before applying it.

Do not delete a section and say nothing. Do not replace content with a shorter summary unless the user approves. The reader who relied on that section will notice it is gone.

**Preserve visual indicators in tables.** Keep emoji markers (`:heavy_check_mark:`, `:x:`, ✅, ❌), colored HTML cells, and other visual encodings for yes/no or pass/fail. Convert GitHub shortcodes to their Unicode equivalents (✅, ❌) when moving from markdown shortcodes to plain markdown, but never flatten them to plain "Yes"/"No" text. The visual weight carries meaning the reader scans for. If the original table uses HTML with inline color styles, convert to a standard markdown table with the emoji markers instead.

**Structural changes are the user's choice.** Rewriting prose within a section is red-pen's job. Removing sections, merging sections, splitting a document into multiple files, or changing the document's outline requires the user's approval. Present the proposed structure, explain why, and wait for a yes. If working unattended, flag structural proposals in the output and mark them as pending approval.

## Reference files

This file holds the rules that apply to all writing. Read the matching reference file before starting these tasks:

| Task | Read first |
|------|-----------|
| The user runs `/red-pen`, with or without arguments | `references/commands.md` |
| Cleaning code comments in a project | `references/commands.md` (the "Never touch these comments" table) |
| Personal writing: a profile README, bio, or About page | `references/personal-voice.md` |
| Planning a doc set, or writing examples or method references | `references/doc-structure.md` |
| Writing a README, how-to, reference doc, checklist, or changelog | `references/templates.md` |
| Adding or editing a diagram | `references/diagrams.md` |
| Replacing a word or checking a term | `references/word-swaps.md` |

## Personal voice

Personal writing includes a GitHub profile README, a bio, an About page, or a personal site intro. For these, ask the user before rewriting: keep their voice, or apply full red-pen. "Keep my voice" keeps contractions, first person, and personality. It still cuts filler, slop, and empty self-promotion. The full rules are in `references/personal-voice.md`.

## Commands

`/red-pen` cleans docs or code comments in a repo, an attached file, or a Notion, Confluence, or Google Docs page. The default is a preview. Nothing is written until the user passes `--fix`. Platform pages also need an explicit yes before the push. The full command reference is in `references/commands.md`.

## Sentence rules

Borrowed from ASD-STE100 and FAA Order 1000.36A. These are hard limits, not suggestions.

- Procedural sentences (telling the reader what to do): 20 words max.
- Descriptive sentences (explaining how something works): 25 words max.
- One instruction per sentence. One topic per paragraph.
- Six sentences per paragraph max.
- If a sentence needs a second read to parse, it is too long. Split it.

## Voice and tone

Borrowed from DigitalOcean Product Docs voice standards.

- Calm. Neutral. Authoritative. No hype, no urgency, no cheerleading.
- Task-first. Lead with what the reader needs to do or know.
<!-- red-pen-ignore-start -->
- No motivational or encouraging language. No "Great job!" or "You are now ready to explore the possibilities."
- No minimizing language: "just", "simply", "easily", "merely", "straightforward."
- No promotional language. Do not describe features as "powerful", "robust", "elegant", or "cutting-edge."
<!-- red-pen-ignore-end -->
- Do not promise outcomes. State what the tool does. The reader decides if it works for them.
- Do not speculate about the reader's experience or intent.
- Use "you" for the reader. Use "we" only when speaking as the maintainer or team.

## Grammar

- Active voice. Name the actor. "The server starts on port 3000." Not "Port 3000 is started by the server."
- Use passive voice only when the actor is unknown or irrelevant: "The file is created during install."
- Simple tenses only. No present perfect ("has completed" becomes "completed"). No progressive forms used as asides (", making it easy" becomes a new sentence).
- No contractions. Write "do not", "cannot", "it is." This keeps text translation-friendly and unambiguous for non-native English readers.
- Keep articles and "that." Write complete grammar, not telegraph style.
<!-- red-pen-ignore-next-line -->
- Modals: use can, will, must. Do not use should, would, may, might, could. A required "should" becomes "must." An optional one gets deleted.
- Condition before action, with a comma: "If the build fails, check the log."

## Banned patterns

These are AI writing fingerprints. Do not use them.

<!-- red-pen-ignore-start -->
- No em dashes. Write two sentences, use a comma, or name the relation ("because", "but", "so").
- No semicolons. Write two sentences.
- No rule-of-three adjective stacks. "Fast, flexible, and scalable" is filler. Pick the one that matters and say why.
- No negative parallelisms. "It is not just a tool, it is a platform" is promotional language dressed as insight. Delete it.
- No false ranges. "From small teams to enterprise organizations" implies a spectrum with no data behind it. Be specific or delete it.
- No compulsive summaries. No "In summary," "Overall," "In conclusion," "To sum up." If the doc is clear, the reader does not need a recap.
- No formulaic transitions. Do not use "Moreover," "Furthermore," "Additionally," "It is worth noting," "It is important to note." Start the sentence with the fact.
- No hedging. Do not write "This can potentially help" when you mean "This helps."
- No vague attribution. "Studies show" or "experts say" with no named source is filler. Name the source or delete the claim.
- No bolded-term-colon-definition pattern. "**Idempotent:** A property where..." is a ChatGPT signature. Use a table, a bullet list, or work the definition into the sentence.
- No decorative bold. Bold is for UI labels, commands the reader must type, warnings, and shortlist lines in checklists (see the Checklist / Procedure template in `references/templates.md`). Not for emphasis.
- No emoji as structural elements. No emoji bullets, no emoji headers, no emoji section markers.
- No inflated symbolism. "Serves as a testament to" means nothing. State the fact.
<!-- red-pen-ignore-end -->

## Word choice

Use `references/word-swaps.md` for the full list of replacements. Core principle: if a simpler word carries the same meaning, use it. If the swap changes the meaning, keep the original word. The "Technical terms" table in that file lists the common cases, such as "deprecate" and "authenticate."

- One word, one meaning, for the whole document. Pick a term and stick with it. Do not alternate between "config", "configuration", "settings", and "options" for the same thing.
- Break noun chains over three words with a preposition. "The connection pool timeout value" becomes "the timeout value for the connection pool."
- Technical nouns and product names are fine. Do not simplify "Kubernetes" to "container system" or "PostgreSQL" to "database" when precision matters.
- Define a concept term at first use, in under ten words. Example: "idempotent (safe to run more than once)." Do not define product names, standard names (HTTP, S3, REST), or the tool the document is about.

## Code and commands

- State what a command does before showing it. Not after.
- Inline backticks for short references: Run `npm install` to add dependencies.
- Code blocks for anything the reader will copy or run.
- Show expected output when the result is not obvious or when the reader needs to confirm success.
- Mark destructive commands before the code block: "This deletes all rows in the table."
- Use consistent placeholder format throughout a document. Pick one: `your_api_key`, `<API_KEY>`, or `$API_KEY`. Document the convention at first use.
- Never put a command in a sentence without backticks.
- Do not explain what a command does line-by-line unless the reader needs to modify it. If they need to modify it, use inline comments.

## Code comments

Comments explain why, not what. The code already says what it does. A comment that restates the function name, the parameter types, or the obvious logic is noise.

Before editing existing comments, read the "Never touch these comments" table in `references/commands.md`. Directives, suppressions, and license headers look like comments, but tools depend on their exact text.

- Do not restate the function signature. If the name and types tell the reader what happens, skip the comment. `pluralItems(n int) string` does not need `// pluralItems returns "item" or "items" for the given count.`
- A comment must never be longer than the function it describes. The exception is a language doc standard (Go doc, TSDoc, RubyDoc, Javadoc) that tooling publishes as API reference.
- Comment the why, not the what. "Bypass the run-once marker because an explicit invocation implies intent" is useful. "Process packages from the blueprint file" above a call to `ProcessPackages` is not.
- Language doc standards get a pass. Go doc comments on exported functions, TSDoc on public interfaces, RubyDoc on module APIs are fine because tooling consumes them. Keep them tight. One sentence for the what, then the why or the gotcha.
- Inline comments earn their place when they explain a non-obvious decision: a workaround, a performance tradeoff, a constraint the next developer will hit. Delete inline comments that narrate the line they sit on.
- Section dividers (`// --- section name ---`) are fine when a file has distinct logical blocks that do not warrant separate files. They are navigational aids, not documentation.
- Do not comment closing braces, end-of-block markers, or obvious control flow.
- Do not write TODO, FIXME, or HACK comments without a name or ticket number. `// TODO(levi): migrate to v2 API` stays. `// TODO: fix this later` is trash.

## Formatting

- Headings carry structure, not commentary. "Install" not "Getting started with installation."
- A heading must cover at least two sentences or items. If it covers one sentence, fold it into the previous section.
- Vertical lists are for three or more parallel items or steps. Two items go in a sentence.
- Numbered lists for sequential steps. Bullet lists for unordered items.
- Tables for structured data with three or more columns. Two columns can be a bullet list ("term: description"). Do not use definition lists in Markdown. GitHub and most Markdown renderers do not support them. AsciiDoc and reStructuredText support them, so use them there.
- One blank line between sections. No decorative whitespace.
- No horizontal rules unless separating fundamentally different content on the same page.

## Lint check

The linter is a quick sweep, not the verdict. It catches the mechanical patterns fast, so you do not have to eyeball them. It does not judge the writing. Delivering clean prose is still your job. Use the first option that works:

1. If `red-pen-lint` is on the PATH, run `red-pen-lint <file>`.
2. If Go is installed, run `go run github.com/TheFynx/red-pen/cmd/red-pen-lint@latest <file>`. The first run needs network access to download the module.
3. From a clone of this repo, run `go run ./cmd/red-pen-lint <file>`. This works offline.

If none of these work, that is fine. Do not skip the self-check. Run the [self-check](#self-check-before-delivery) by hand. Tell the user the linter did not run, and why. Example: "Go is not installed." Give them the install command: `go install github.com/TheFynx/red-pen/cmd/red-pen-lint@latest`.

The linter reads files. To lint a platform doc or a draft, write it to a temporary `.md` file first.

It reports two severities:

**Violations** are mechanical pattern hits:

- Punctuation: em dashes, ` -- ` and unspaced `--` typed as a dash, en dashes used as a dash, and semicolons. En dashes in number ranges such as "1–5" pass.
- Words: banned modals, minimizing words, cheerleading, AI slop vocabulary, and promotional words.
- Phrases: formulaic transitions and compulsive summaries.
- Style: bolded-term-colon-definition patterns (colon inside the bold or right after it), contractions, and sentence length. The pattern check skips notice blockquotes and list labels in bold. A bold label that ends in a period is fine, as in "**Write mode (default).**".

**Findings** are word-swap candidates from `references/word-swaps.md`:

- A finding names the word and the line. It is a suggestion, not a defect. Do not treat it as a ban.
- Apply the swap when the simpler word carries the same meaning.
- Keep the word when it names the technical thing in context: compiler or SQL optimization, a product name, or a Technical terms entry.
- To remove a kept word from the report, add an ignore directive around the text (see [Ignore a line or region](#ignore-a-line-or-region)).

Sentence length uses 20 words for numbered list items, because those are steps. Everything else uses 25 words. Word checks and sentence length also run on table cells. Headings skip the word checks.

It skips YAML frontmatter, fenced code blocks (backticks or tildes), indented code blocks, inline code, HTML blocks, HTML tags, and HTML comments.

Exit code 0 means no mechanical violations. Exit code 1 means violations. Neither means the writing is good or bad. Fix the violations, decide each finding, then run the self-check. The self-check is the delivery gate.

**What the linter cannot see.** Do not rely on it alone. It cannot catch:

- Whether a sentence is true, specific, or in the right place.
- Rule-of-three stacks, negative parallelisms, false ranges, hedging, and vague attribution. These are rules here, but regexes cannot judge them.
- Whether a step names the real command, file, or value. It only counts words.
- Tone, structure, and reader fit: is this page for a beginner or an expert, and does it say the right thing first.
- Anything inside code blocks, inline code, HTML, frontmatter, or ignored regions. That content is skipped on purpose, not judged.

For personal writing in "Keep my voice" mode, add `--voice personal`. This skips the contraction, modal, minimizing-word, semicolon, em-dash, and sentence-length checks. All other checks still run.

```bash
red-pen-lint --voice personal README.md
```

### Ignore a line or region

Some text must keep a flagged word. Examples are a quoted error message, a UI label, a technical use, or a style guide that names the banned words. Mark that text so the linter skips it. Each directive goes on its own line.

```markdown
<!-- red-pen-ignore-next-line -->
The installer prints "You should restart now."

<!-- red-pen-ignore-start -->
Do not write "simply", "just", or "easily."
<!-- red-pen-ignore-end -->
```

Use these only when the text cannot change, or when the hit is a false positive. A false positive is a technical use of a word, or prose you did not write. Do not use them to hide your own prose from the linter.

## Self-check before delivery

This is the delivery gate. Run the lint check first if you can. Then run this on every piece of writing before delivering it. A clean lint run does not replace this. The linter only sees mechanical patterns. This list is where the actual quality decisions happen.

1. Find your three longest sentences. Are any over 20 words (procedural) or 25 words (descriptive)? Split them.
2. Search for em dashes. Remove all of them.
3. Search for words in `references/word-swaps.md`. Replace any hits, unless the word names the technical thing. Examples: a "Technical terms" entry, compiler or SQL optimization, a product name.
<!-- red-pen-ignore-start -->
4. Search for "should", "may", "might", "could", "would." Replace or delete.
5. Search for bold text. Is each instance a UI label, a command, a warning, or a checklist shortlist line? If not, remove the bold.
6. Search for "just", "simply", "easily", "straightforward." Delete them.
7. Search for "Overall", "In summary", "In conclusion", "It is worth noting", "It is important to note." Delete the sentence or rewrite it as a fact.
<!-- red-pen-ignore-end -->
8. Read each step in a procedure. Does it name the specific command, file, service, or value the reader needs? If not, add it.
9. Read the first sentence of each section. Does it state a fact or an action? If it restates the heading, delete it.
10. Read the last sentence of each section. Is it a summary of what the section already said? Delete it.
11. For checklists: read only the bold lines in sequence. Do they form a complete procedure without the unbold text? If not, rewrite the bold lines.
12. For examples: does every example have a label, working code, and expected output? If not, add them.
13. For method/function references: does every method have a signature, params table, example, and expected output? If not, add them.
14. For diagrams: did you test the diagram in the target renderer? Does a prose description exist near the diagram?

## Out of scope

This skill is for facts, instructions, and clear communication. It is not for marketing copy, blog posts, or brand writing because it deletes persuasion and personality. Say so if asked to apply it there. Personal profiles and bios are in scope through [Personal voice](#personal-voice).
