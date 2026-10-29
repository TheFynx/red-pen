# Document templates

Part of the red-pen skill. The core rules in `SKILL.md` apply here too.

## README

```
# Project Name
One sentence: what it does and who it is for.

## Requirements
What the reader needs before starting. List format.

## Install
Numbered steps. One action per step. Show the command, then confirm the result.

## Usage
Most common use case first. Show a working example with real values.

## Configuration
Table format: variable, type, default, description. Only if the project has configuration.

## Troubleshooting
Common errors. State the problem, then the cause, then the fix.
```

## How-To Guide

```
# How to [accomplish task]
One sentence: what the reader will have when done.

## Before you start
What they need installed, configured, or understood.

## Steps
Numbered. One action per step.
State the action. Show the command or procedure. Confirm the result before moving on.

## Verify
How to confirm the task succeeded.

## Common errors
Problem. Cause. Fix. No filler.
```

## Reference / API Doc

```
# [Feature or endpoint name]
One sentence: what it does.

## Syntax
The call signature or usage pattern.

## Parameters
Table: name, type, required, description.

## Example
A working request and response with real values.

## Response
The shape of what comes back. Field names, types, descriptions.

## Errors
Table: code, meaning, what to do.
```

## Checklist / Procedure

Use this template for operational procedures, setup guides, and multi-phase tasks. It produces one document that serves two reading speeds.

**Bold/unbold rule.** Bold lines form a standalone shortlist. An experienced reader must be able to read only the bold lines, in order, and have a complete procedure. Unbold lines expand on the bold line above them with detail, sub-steps, screen-by-screen instructions, or expected results. Every bold line must make sense without the unbold text below it.

```
# [System or Task Name]

## Overview
Numbered list of major phases. One line each.
This is the table of contents. The reader
sees the full scope before starting.

## [Phase Number]. [Phase Name]

NOTE: Conditions or context. Only if needed.
Keep notes before the steps they apply to.

**[Device/Location]: [Action]**
Detail for beginners: what screen loads, what
button to press, what value to enter.

**[Device/Location]: [Next action]**
Detail if needed. Not every bold line needs
expansion.

Verify:
What the operator should see before moving on.
List each indicator: light, value, message.

## [Phase Number]. [Phase Name with Branch]

Steps common to all paths go here.

**FOR [CONDITION A]: Continue with section [X-X].**
**FOR [CONDITION B]: Continue with section [X-X].**

## [X-X]. [Branch A Name]
(Continued from [parent section] step [N])

**[Device]: [Action]**
Detail.

## [X-X]. [Branch B Name]
(Continued from [parent section] step [N])

**[Device]: [Action]**
Detail.

## [Shutdown / Cleanup Phase]
Numbered sub-phases if multiple components
need separate shutdown sequences.

## Troubleshooting

### [Problem Category]

**[Symptom or error message]**
Cause. What to check. How to fix it.
Reference the main procedure step to
return to after the fix.
```

**Rules for this template:**

- Bold lines are the shortlist. They must read as a complete procedure on their own.
- Unbold lines are the detail. They expand on the bold line above them.
- Name the device or location before the action in every bold line.
- One action per bold line.
- Verify blocks go after a group of steps, not after every step.
- Notes go before the steps they apply to, not after.
- Branch points name the condition and the target section. Both branches live in the same document.
- Troubleshooting is a separate section at the end. Reference it inline only where failures are common.
- Numbering is hierarchical: 1-1, 1-2, 1-3 for phases. 1-4-1, 1-4-2 for sub-phases.

## Changelog Entry

```
## [Version] - [Date]

### Added
- What is new. One line per feature. Start with a verb.

### Changed
- What behaves differently. Name the old behavior and the new behavior.

### Fixed
- What broke and what the fix does.

### Removed
- What is gone. State what replaces it, or say nothing replaces it.
```

