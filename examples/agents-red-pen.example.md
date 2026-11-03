# The agent-facing half of red-pen onboarding. Paste the section below into
# the repo's AGENTS.md or CLAUDE.md, or let /red-pen onboard write it.

## Red pen

- Tracker: linear
- Story style: Problem and Outcome, with acceptance criteria as a checklist
- Platforms: Notion, Confluence
- Ticket auto-detect: on
- Doc paths to skip: vendor/, docs/drafts/

# Line meanings:
# Tracker: the connected tracker /red-pen reads and rewrites tickets through.
# Story style: the form new stories take. Problem and Outcome (this skill's
#   form), the "As a / I want / so that" sentence, or custom. With custom,
#   also list the headings in .red-pen.yml as a story-sections key.
# Platforms: which doc platforms the team uses, so /red-pen warns about
#   missing connectors only for platforms in use.
# Ticket auto-detect: mirrors ticket-detect in .red-pen.yml. off is the
#   answer when repo owners want ticket checks only on an explicit --ticket.
# Doc paths to skip: folders /red-pen docs never rewrites.