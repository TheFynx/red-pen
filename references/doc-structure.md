# Document structure and examples

Part of the red-pen skill. The core rules in `SKILL.md` apply here too.

Each page answers one question for one reader at one stage. A page that tries to answer "what is this," "how do I start," and "how do I configure it for AWS" fails all three readers.

## The reader's path

Documentation is a funnel. Each page assumes the reader arrived from the page before it. Do not repeat what the previous page covered.

1. **README.** What is this and what does it do. One sentence, then install, then the most common use case. The reader decides in 30 seconds whether to continue.
2. **Getting started.** The fastest path from zero to working. One platform, one workflow, no options. The reader proves the tool works before investing in configuration. Link to other platforms or advanced setup at the end, not inline.
3. **Config / reference landing page.** Index of configuration sections. Each entry names the config area and what it controls. The reader clicks the one they need.
4. **Individual config / reference pages.** One config area per page. Purpose in one sentence, a clean example, then a table of every option with type, default, and description. No option left undocumented.
5. **Troubleshooting.** Common problems, gotchas, and "if you see X, do Y." Point the reader to the issue tracker or support channel for problems not listed.

Not every project needs all five. Small tools can collapse getting started into the README. The principle holds: each page is one stage, and the reader moves forward, not sideways.

## One question per page

- If a page has conditional paragraphs ("For AWS, do X. For GCP, do Y. For Azure, do Z."), each condition is its own page. The landing page links to each one.
- If a getting-started guide includes edge cases, provider-specific wiring, or configuration tuning, those are separate pages. The getting-started page is the happy path only.
- If a tutorial step pauses to define a concept, the definition belongs in a glossary or reference page. Define it once, link it. Do not interrupt the reader's momentum.

## Two reading speeds

Docs serve two readers at the same time: the new user following the path, and the experienced user who knows exactly what they need.

- New users read linearly. Each page ends with a clear next step. They never hit a dead end or a wall of options before they understand the basics.
- Experienced users jump directly to reference pages. Config tables, option lists, and troubleshooting sections must stand alone without requiring the reader to have followed the getting-started path first.
- A page that forces the experienced user through beginner context is as broken as a page that drops a beginner into advanced configuration.

## Caveats and edge cases

- Do not weave caveats between action steps. The reader loses the thread.
- Group caveats in a "Common issues" or "Gotchas" table at the bottom of the page: problem, cause, fix.
- Link to the gotchas table from the step where the problem is most likely to surface. One line: "No output? See [common issues](#common-issues)."

## Examples

Examples are how readers confirm they understand. Every example must work.

- Label examples with a number and a short name: "Example 1: Create a user." This makes them referenceable and scannable.
- State what the example shows before the code block. One sentence.
- Use real values. No `foo`, `bar`, `lorem ipsum`, or `example.com` when a plausible real value works. Use `acme-api.example.com` over `example.com` when a hostname matters.
- Show expected output after commands and API calls. The reader needs to confirm they got the right result.
- Use input/output pairing for transformations. Show the input, then the output. Use a two-column table when the format supports it. Use stacked code blocks (labeled `Input:` and `Output:`) when it does not.
- Add a legend after complex examples. Map each highlighted element to the concept it demonstrates. Use a list, not a paragraph.
- One concept per example. If an example demonstrates two things, split it into two examples.

### Method / function reference

Use this pattern for every method, function, or endpoint in a reference doc. Every method looks the same.

```
## method_name

One sentence: what it does.

### Signature

\`\`\`
function_signature
\`\`\`

### Parameters

| Name | Type | Required | Description |
|------|------|----------|-------------|

### Returns

Type and what it contains.

### Example

\`\`\`
working_example_with_real_values
\`\`\`

### Expected output

\`\`\`
what_comes_back
\`\`\`
```

**Rules for this pattern:**

- Signature is required.
- Parameters table is required if there are parameters.
- Example is required. Not optional, not "see the cookbook."
- Expected output is required when the return value is not obvious.
- One method per section. Do not group methods into shared descriptions.
- No prose between the signature and the parameters table.

