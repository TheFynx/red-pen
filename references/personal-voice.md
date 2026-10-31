# Personal voice

Part of the red-pen skill. The core rules in `SKILL.md` apply here too.

Some writing belongs to a person, not a project. Examples are a GitHub profile README (the `username/username` repo), an "About me" page, a bio, or a personal site intro. The author's voice is part of the content. Red-pen still helps here, but its job is to get the author to the point, not to make them sound like a manual.

**Ask first.** When the target is personal writing, ask the user which voice they want before rewriting. Offer two options, through an interactive multiple-choice prompt when the assistant has one:

- **Keep my voice.** Cut the filler and the rambling. Keep the tone, the first person, and the personality.
- **Full red-pen.** Apply every rule, the same as technical docs.

Skip the question if the user already said which one they want. If working unattended, use "Keep my voice" and say so in the output.

**What "Keep my voice" relaxes:**

- Contractions are fine. Do not expand them.
- First person is fine. "I build", "my work", "I like" stay.
- Casual words and humor stay if the author wrote them. This includes words from the banned lists.
- Modals, minimizing words, semicolons, and em dashes are the author's call. Do not flag them.
- Emoji headers and emoji markers are fine. A profile is a personal page, not a reference doc.
- Sentence length is a guide, not a hard limit. Split a sentence only when it rambles.

**What "Keep my voice" still enforces:**

- Get to the point. Lead with who the author is and what they do. Cut throat-clearing openers and life stories that do not support the point.
- Do not wax poetic. Cut inflated symbolism, grand claims about passion or journeys, and metaphors that carry no fact.
- No AI slop vocabulary. Words from the slop list in `references/word-swaps.md` do not sound like a person.
- No empty self-promotion. Replace adjectives about the author with facts: what they built, who uses it, what it does.
- No hedging stacks, compulsive summaries, or formulaic transitions.
- Every claim stays specific. "Passionate about cloud" becomes what the author builds in the cloud.

Edit mode rules still apply in both voices. Do not delete sections, links, or projects without approval.

To lint personal writing in this voice, run the linter with `--voice personal`. See [Lint check](../SKILL.md#lint-check).

