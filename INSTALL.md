# Install

Red Pen installs as an agent skill. The skill format is an open standard (agentskills.io), and major coding agents read it. A skill is a folder with a `SKILL.md` file and supporting files. Install one copy per agent, or one shared copy that several agents read.

## What gets installed

The complete skill is two pieces:

- `SKILL.md`, the core rules
- `references/`, loaded on demand for commands, templates, structure, diagrams, and word swaps

The Go linter ships separately. See [Linter](#linter).

## Locations by agent

| Agent | Project scope | Personal scope |
|-------|---------------|----------------|
| Claude Code | `.claude/skills/red-pen/` | `~/.claude/skills/red-pen/` |
| OpenCode | `.opencode/skills/red-pen/` | `~/.config/opencode/skills/red-pen/` |
| Cursor 2.4 or newer | `.cursor/skills/red-pen/` | `~/.cursor/skills/red-pen/` |
| Copilot in VS Code | `.github/skills/red-pen/` | `~/.copilot/skills/red-pen/` |
| Copilot CLI and Copilot app | `.github/skills/red-pen/` | `~/.copilot/skills/red-pen/` |
| Codex CLI | `.codex/skills/red-pen/` | `~/.codex/skills/red-pen/` |
| Codex in VS Code, via integration | `.github/skills/red-pen/` | (not supported) |
| Gemini CLI, preview | `.gemini/skills/red-pen/` | (see the CLI docs) |

Project scope shares the skill with everyone who clones the repo. Commit the folder. Personal scope makes the skill available in every project for that user on that machine.

Agents read shared locations beyond the row in the table:

- OpenCode also reads `.claude/skills/`, `.agents/skills/`, `~/.claude/skills/`, and `~/.agents/skills/`.
- Copilot in VS Code also reads `.claude/skills/` and `.agents/skills/` in the project, and `~/.claude/skills/` and `~/.agents/skills/` in the home directory.
- Codex also reads `.agents/skills/` in the project.
- Claude Code reads the `.claude/skills/` locations only.

One install, two agents: a personal copy at `~/.claude/skills/red-pen/` serves Claude Code, OpenCode, and personal-scope Copilot in VS Code.

## Install a copy

Clone the repo:

```bash
git clone https://github.com/TheFynx/red-pen.git
cd red-pen
```

Copy the two pieces. Personal scope for Claude Code, for example:

```bash
mkdir -p ~/.claude/skills/red-pen
cp SKILL.md ~/.claude/skills/red-pen/
cp -r references ~/.claude/skills/red-pen/references
```

For another agent, run the same three commands against the target path in the table. Keep the folder name `red-pen`. Copilot in VS Code requires the directory name to match the `name` field in `SKILL.md`.

Project scope, shared with the team:

```bash
mkdir -p .claude/skills/red-pen
cp SKILL.md .claude/skills/red-pen/
cp -r references .claude/skills/red-pen/references
```

Commit the folder, so everyone who clones the repo gets it.

## Symlinks for development

A path in Claude Code can be a symlink to the clone:

```bash
ln -s <clone-dir> ~/.claude/skills/red-pen
```

Claude Code reads `SKILL.md` through the link. Cursor does not follow symlinks and cannot see the skill. With Cursor in the mix, use copies everywhere.

Personal-scope skills in Claude Code load in terminal and desktop sessions on that machine. Cowork and cloud sessions load the skills enabled for your claude.ai account. Cloud sessions outside a project thread also load the skill from `.claude/skills/` in the repo.

## Invoke the skill

| Agent | Auto-load | Explicit invocation |
|-------|-----------|---------------------|
| Claude Code | Yes, from the description | `/red-pen docs --fix`. The skill name is the slash command |
| OpenCode | Yes | `/red-pen docs` with the command file below, or `@red-pen` in a prompt |
| Copilot in VS Code | Yes | `red-pen` in the `/` menu, with arguments after the name |
| Cursor | Yes | Ask for it by name: "run the red-pen skill on this README" |
| Codex CLI | Yes, check the available skills list | Ask for it by name in the prompt |
| Gemini CLI | Yes, preview | Ask for it by name in the prompt |

OpenCode needs one command file for the slash form. Save it at `~/.config/opencode/commands/red-pen.md`:

```markdown
---
description: Clean docs, code comments, or platform pages in plain language
---
Load the red-pen skill with the skill tool (id red-pen), then apply it to this task: $ARGUMENTS.
```

Claude Code needs no command file. The skill registers `/red-pen` on its own.

## Agents without skill support

Any agent that reads repository files can follow the skill without a loader.

1. Vendor the folder: copy `SKILL.md` and `references/` into the repo, for example under `tools/red-pen/`.
2. Add one line to `AGENTS.md` or the instructions file: "For documentation work, read `tools/red-pen/SKILL.md` first, then follow it."
3. The agent loads the reference files when `SKILL.md` points at them.

For claude.ai on the web: paste the `SKILL.md` body into project knowledge, or enable it through the skills feature in account settings.

## Platform connectors

The `/red-pen docs` flow can read and push pages in Notion, Confluence, and Google Docs. That requires the platform connector of the host tool. Without the connector, the skill produces the cleaned output as a local file. See `references/commands.md`.

## Linter

The linter runs on any machine with Go. It applies the same rules as the skill and works without any agent.

```bash
go install github.com/TheFynx/red-pen/cmd/red-pen-lint@latest
```

```bash
red-pen-lint README.md docs/*.md
```

Exit 1 means violations to fix. Word-swap suggestions print as findings and do not fail the run. See [Automated lint](SKILL.md#automated-lint) in `SKILL.md` for the full behavior.

## Update and remove

The installed skill is a copy. After editing the source repo, run the copy commands again.

- OpenCode and Claude Code watch the skill folder. The running session picks up the change without a restart.
- Other agents pick up the change in the next session.

Uninstall by deleting the skill folder and the OpenCode command file, if installed.