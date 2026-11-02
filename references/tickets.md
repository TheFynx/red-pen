# Tickets

Part of the red-pen skill. The core rules in `SKILL.md` apply here too.

Use this file to write or rewrite a story, bug, task, spike, or epic in any tracker.

## Write for a person

A ticket is for the person who picks up the work. It is not a record for an agent to parse. That person must understand the work after one read.

Test every ticket before you post it:

- A reader new to the work can say what changes and why after one read.
- The title and the first two lines are enough to pick the ticket from a board.
- Every line helps someone do or review the work.

## Rules

1. Answer three questions: what changes, why, and how you know it is done.
2. State each fact once. Do not repeat a step in the criteria, the checklist, and the rollout plan.
3. Delete empty sections and template placeholders before you post. A heading with nothing under it is noise.
4. Every acceptance criterion is a check a reviewer can pass or fail.
5. State the impact in one line with a number. Do not argue for the ticket.
6. Put reference material in a link: metric catalogs, sample queries, cost math, and long code excerpts.
7. Put history in a comment: incident timelines, related PRs, moved scope, and investigation notes.
8. Add optional sections only when the work needs them. Rollout and backout belong on work that touches production or shared systems.
9. Use tables only for data the reader compares. Summarize the rest in a sentence.

## What to cut

| Pattern | Fix |
|---------|-----|
| Two opening sections that say the same thing | Keep one. Call it Problem. |
| A fact repeated across criteria, checklist, and rollout steps | Keep it in the section where someone acts on it. |
| Several spike sections that restate the question | Keep the question and the done-when list. |
| A data table with one row that matters | One sentence with the number. |
| Paragraphs that argue for the ticket | One line of impact. |
| Rationale inside every criterion | The criterion only. Put the reason in Problem once. |
| `TBD`, `[option A / option B]`, "fill in later" | The real value, or delete the line. |
| Hedging such as "it seems like" or "possibly related" | State what you know and what you do not know. |

## Acceptance criteria

Write a checklist by default. One check per line.

- Name the value, behavior, or output a reviewer checks.
- Do not use words a reviewer cannot test: "works", "properly", "correctly", "as expected", "fast", "better".
- Use Given/When/Then only for behavior with several branches. One scenario per branch.

Untestable:

```markdown
- [ ] Export works correctly.
```

Testable:

```markdown
- [ ] The Export button downloads the filtered rows as a CSV file.
- [ ] An export with zero rows downloads a CSV with only the header row.
```

Given/When/Then, for branching behavior:

```gherkin
Given a saved card that expired last month
When the customer pays with it
Then checkout shows "Card expired" and keeps the cart
```

## Story

New behavior a user or system needs.

```markdown
# [Verb] [thing]: 10 words max

## Problem
What is wrong or missing now, for whom. One line of impact with a number.

## Outcome
What is true when this ships.

## Acceptance criteria
- [ ] One testable check per line.

## Links
Optional. Design, docs, related tickets.

## Risks
Optional. What breaks if this goes wrong, and how to tell.

## Rollout and backout
Optional. Only for production or shared systems.
```

Some teams require the "As a / I want / so that" form. Put it under a `## User story` heading, in place of Problem and Outcome:

```markdown
## User story
As a support agent, I want to export filtered rows, so that I stop copying them by hand.
```

## Bug

Behavior that differs from what the product promises.

```markdown
# [Thing] [fails how] [when]

## Steps to reproduce
1. Numbered. Real values. Start from a known state.

## Expected
What the product does when it works.

## Actual
What it does instead. Error text in a code block. Rate or count if intermittent.

## Environment
Service, version or commit, region, browser. Only what narrows the cause.

## Logs
Optional. The shortest excerpt that shows the failure.

## Workaround
Optional. What users can do until the fix ships.
```

## Task

Work with no user-facing change: an upgrade, a refactor, a config change.

```markdown
# [Verb] [thing]

## What
The change, in one or two lines.

## Why
The reason, with a number when one exists.

## Done when
- [ ] One testable check per line.
```

## Spike

Time-boxed research that answers one question.

```markdown
# Decide [question]

## Question
One question. A yes/no or a choice between named options.

## Timebox
How long, such as 2 days.

## Deliverable
What the spike produces, such as a decision memo or a proof of concept.

## Done when
- [ ] The question has an answer with evidence.
- [ ] Follow-up stories exist, or the memo says none are needed.

## Constraints
Optional. Environments, approvals, budget.
```

The spike ticket holds the question. The findings go in the deliverable, not in empty sections of the ticket.

## Epic

Before writing an epic, ask the user: "Is this a true epic or a bucket epic for ad hoc work?"

- A true epic has a finish line. It closes when the end state is true.
- A bucket epic holds ad hoc work with no finish line, as in Scrumban. It stays open.

True epic:

```markdown
# [Outcome]

## Problem
What is wrong or missing now, for whom.

## Outcome
What is true when the epic closes.

## End state
The check that closes the epic, such as "all invoices send without manual steps."

## Stories
- Link or one line per story.
```

Bucket epic:

```markdown
# [Area]

## Purpose
What kind of work lives here, in one or two lines.

## In scope
- Kinds of work that belong.

## Out of scope
- Kinds of work that go somewhere else, and where.

## Entry rule
What a story needs before it goes in, such as a problem and an owner.

## Review cadence
When someone grooms the bucket and closes or moves stale stories.
```

## Trackers

Red-pen writes tickets for Jira, Linear, GitHub Issues, GitLab, Asana, Shortcut, Trello, and any other connected issue tracker. Write Markdown. Most trackers render it. Convert only when the tracker does not.

To find the team's tracker, check in this order:

1. A line in the project's `AGENTS.md` or `CLAUDE.md`, such as `red-pen tracker: linear`.
2. A `.red-pen.yml` file in the repo root with a `tracker` key:

   ```yaml
   tracker: linear
   ```

3. Ask the user.

To read or rewrite an existing ticket from a link, see `references/commands.md`.

## Lint a ticket

The linter runs ticket checks when the file covers three sections of one ticket type. To force ticket mode, pass `--ticket`. To set the type, pass `--ticket=story`, `bug`, `task`, `spike`, or `epic`:

```bash
red-pen-lint --ticket=bug ticket.md
```

Ticket checks:

| Rule | Severity | What it checks |
|------|----------|----------------|
| `ticket-section` | Violation | A required section for the type is missing. |
| `ticket-title` | Violation | The `#` title is over 10 words. |
| `ticket-placeholder` | Violation | Placeholder text such as `TBD` or `[option A / option B]` is left in. |
| `ticket-empty-section` | Violation | A heading has no content under it. |
| `ticket-type` | Violation | `--ticket` was passed, and no heading shows the type. |
| `ticket-vague` | Finding | An acceptance criterion uses an untestable word such as "works" or "properly". |

The linter accepts a bold-only line, such as `**Steps to Reproduce**`, as a heading. Tracker exports often use that form. An epic with a Purpose, Entry rule, or Review cadence heading is checked as a bucket epic.

To lint a ticket from a tracker, write it to a temporary `.md` file first.
