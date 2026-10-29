package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

// words returns n filler words that trigger no rule.
func words(n int) string {
	return strings.TrimSpace(strings.Repeat("word ", n))
}

// hits formats violations as sorted "line:rule" strings for comparison.
func hits(vs []violation) []string {
	out := []string{}
	for _, v := range vs {
		out = append(out, fmt.Sprintf("%d:%s", v.Line, v.Rule))
	}
	sort.Strings(out)
	return out
}

func TestLintLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		// Content that is skipped.
		{
			name:  "frontmatter skipped",
			input: "---\ndescription: You should just use it.\n---\n\nYou should restart.",
			want:  []string{"5:banned-modal"},
		},
		{
			name:  "backtick fence skipped",
			input: "```\nyou should not lint this\n```\nYou should lint this.",
			want:  []string{"4:banned-modal"},
		},
		{
			name:  "tilde fence skipped",
			input: "~~~sh\nyou should not lint this\n~~~\nYou should lint this.",
			want:  []string{"4:banned-modal"},
		},
		{
			name:  "longer fence holds shorter fence",
			input: "````md\n```\nyou should not lint this\n```\n````\nYou should lint this.",
			want:  []string{"6:banned-modal"},
		},
		{
			name:  "indented code after blank line skipped",
			input: "Run this:\n\n    you should not lint this\n\nYou should lint this.",
			want:  []string{"5:banned-modal"},
		},
		{
			name:  "indented list continuation is prose",
			input: "- First item.\n\n    You should lint this.",
			want:  []string{"3:banned-modal"},
		},
		{
			name:  "inline code stripped",
			input: "Set `should` and `just` in the config.",
			want:  []string{},
		},
		{
			name:  "word after inline code still flagged",
			input: "Run `adjust` and then just restart it.",
			want:  []string{"1:minimizing"},
		},
		{
			name:  "html block skipped",
			input: "<td style=\"color: red; width: 10%\">you should not lint</td>\nYou should lint this.",
			want:  []string{"2:banned-modal"},
		},
		{
			name:  "inline html and entities stripped",
			input: "Press <kbd>Ctrl</kbd> &amp; <kbd>C</kbd> to stop.",
			want:  []string{},
		},
		{
			name:  "multi-line html comment skipped",
			input: "<!-- you should\nnot lint this; at all -->\nYou should lint this.",
			want:  []string{"3:banned-modal"},
		},

		// Ignore directives.
		{
			name:  "ignore next line",
			input: "<!-- red-pen-ignore-next-line -->\nYou should skip this.\nYou should flag this.",
			want:  []string{"3:banned-modal"},
		},
		{
			name:  "ignore region",
			input: "<!-- red-pen-ignore-start -->\nJust skip this; all of it.\n<!-- red-pen-ignore-end -->\nYou should flag this.",
			want:  []string{"4:banned-modal"},
		},
		{
			name:  "directive in backticks does nothing",
			input: "Write `<!-- red-pen-ignore-start -->` to start.\nYou should flag this.",
			want:  []string{"2:banned-modal"},
		},

		// Punctuation.
		{
			name:  "em dash flagged",
			input: "The build fails — check the log.",
			want:  []string{"1:em-dash"},
		},
		{
			name:  "double hyphen dash flagged",
			input: "The build fails -- check the log.",
			want:  []string{"1:em-dash"},
		},
		{
			name:  "en dash in number range allowed",
			input: "Pages 1–5 cover setup.",
			want:  []string{},
		},
		{
			name:  "en dash as sentence dash flagged",
			input: "The build fails – check the log.",
			want:  []string{"1:em-dash"},
		},
		{
			name:  "table separator is not a semicolon",
			input: "| A | B |\n|---|:---:|\n| x | y |",
			want:  []string{},
		},

		// Words.
		{
			name:  "curly apostrophe contraction flagged",
			input: "Don’t restart the server.",
			want:  []string{"1:contraction"},
		},
		{
			name:  "word inside a longer word not flagged",
			input: "Make an adjustment to the awesomeness setting.",
			want:  []string{},
		},
		{
			name:  "table cells checked",
			input: "| Option | Description |\n|---|---|\n| port | You should set this. |",
			want:  []string{"3:banned-modal"},
		},
		{
			name:  "headings skip word checks",
			input: "# Why you should leverage this",
			want:  []string{},
		},

		// Sentence length.
		{
			name:  "21-word step flagged",
			input: "1. " + words(21) + ".",
			want:  []string{"1:sentence-length"},
		},
		{
			name:  "20-word step allowed",
			input: "1. " + words(20) + ".",
			want:  []string{},
		},
		{
			name:  "21-word paragraph allowed",
			input: words(21) + ".",
			want:  []string{},
		},
		{
			name:  "26-word paragraph flagged",
			input: words(26) + ".",
			want:  []string{"1:sentence-length"},
		},
		{
			name:  "26-word table cell flagged",
			input: "| A |\n|---|\n| " + words(26) + ". |",
			want:  []string{"3:sentence-length"},
		},

		// Word swaps come out as findings, not violations.
		{
			name:  "verbose words are findings",
			input: "Utilize it to optimize the query.",
			want:  []string{"1:word-swap", "1:word-swap"},
		},

		// Patterns.
		{
			name:  "colon after bold flagged",
			input: "**Term**: a plain definition",
			want:  []string{"1:bold-colon-def"},
		},
		{
			name:  "colon inside bold flagged",
			input: "**Idempotent:** safe on a rerun",
			want:  []string{"1:bold-colon-def"},
		},
		{
			name:  "bold label with period not flagged",
			input: "**Step one.** Then run the check.",
			want:  []string{},
		},
		{
			name:  "notice blockquote not flagged",
			input: "> **Warning:** The rewrite keeps this notice.",
			want:  []string{},
		},
		{
			name:  "bold list label not flagged",
			input: "- **TTS engines:** XTTSv2 and Bark.",
			want:  []string{},
		},
		{
			name:  "expanded contractions flagged",
			input: "You mustn't run it. It'd break the build.",
			want:  []string{"1:contraction", "1:contraction"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hits(lintLines(strings.Split(tt.input, "\n")))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitSentences(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"Pick a port, e.g. 8080, for the server. Then start it.", 2},
		{"Use one format, i.e. YAML. Then commit it.", 2},
		{"**Do this first.** Then do that.", 2},
		{"Run it (as root.) Then check it.", 2},
		{"Version 1.5 is out. Upgrade now!", 2},
	}
	for _, tt := range tests {
		if got := splitSentences(tt.input); len(got) != tt.want {
			t.Errorf("splitSentences(%q) = %q, want %d sentences", tt.input, got, tt.want)
		}
	}
}

func TestFilterVoice(t *testing.T) {
	vs := lintLines([]string{"I'm told you should just delve in — it's fine."})

	docs := hits(filterVoice(vs, "docs"))
	wantDocs := []string{"1:ai-slop", "1:banned-modal", "1:contraction", "1:contraction", "1:em-dash", "1:minimizing"}
	if !reflect.DeepEqual(docs, wantDocs) {
		t.Errorf("docs voice: got %v, want %v", docs, wantDocs)
	}

	personal := hits(filterVoice(vs, "personal"))
	wantPersonal := []string{"1:ai-slop"}
	if !reflect.DeepEqual(personal, wantPersonal) {
		t.Errorf("personal voice: got %v, want %v", personal, wantPersonal)
	}
}

func TestFindingsSeverity(t *testing.T) {
	vs := lintLines([]string{"Utilize it to optimize the query."})
	if len(vs) != 2 {
		t.Fatalf("got %d hits, want 2 word-swap hits", len(vs))
	}
	for _, v := range vs {
		if v.Rule != "word-swap" || v.Severity != "finding" {
			t.Errorf("got %q severity %q, want a word-swap finding", v.Rule, v.Severity)
		}
	}
}

func TestViolationSeverity(t *testing.T) {
	vs := lintLines([]string{"You should restart the server."})
	if len(vs) != 1 {
		t.Fatalf("got %d hits, want 1", len(vs))
	}
	if vs[0].Severity != "violation" {
		t.Errorf("got severity %q, want violation", vs[0].Severity)
	}
}
