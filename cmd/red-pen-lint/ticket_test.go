package main

import (
	"reflect"
	"strings"
	"testing"
)

func lintTicketString(input, mode string) []string {
	return hits(lintTicket(strings.Split(input, "\n"), mode, defaultTicketConfig))
}

func TestLintTicket(t *testing.T) {
	story := "# Export filtered rows to CSV\n\n## Problem\n\nSupport copies rows by hand.\n\n## Outcome\n\nThe table exports to CSV.\n\n## Acceptance criteria\n\n- [ ] The Export button downloads the filtered rows as CSV.\n"

	tests := []struct {
		name  string
		input string
		mode  string
		want  []string
	}{
		{
			name:  "complete story passes",
			input: story,
			want:  []string{},
		},
		{
			name:  "one ticket heading does not trigger auto-detect",
			input: "# Guide\n\n## Environment\n\nRuns on Linux.\n",
			want:  []string{},
		},
		{
			name:  "two ticket headings do not trigger auto-detect",
			input: "# Data export design\n\n## Problem\n\nCustomers export by hand.\n\n## Proposal\n\nAdd an export button.\n\n## Outcome\n\nFewer support requests.\n",
			want:  []string{},
		},
		{
			name:  "three ticket headings trigger auto-detect",
			input: "# Crash on save\n\n## Steps to reproduce\n\n1. Click Save.\n\n## Expected\n\nThe file saves.\n\n## Actual\n\nIt crashes.\n",
			want:  []string{"1:ticket-section"},
		},
		{
			name:  "forced type reports every missing section",
			input: "# Rotate keys\n\nRotate the keys.\n",
			mode:  "task",
			want:  []string{"1:ticket-section", "1:ticket-section", "1:ticket-section"},
		},
		{
			name:  "forced with no headings cannot tell the type",
			input: "# Rotate keys\n\nRotate the keys.\n",
			mode:  "any",
			want:  []string{"1:ticket-type"},
		},
		{
			name:  "long title flagged",
			input: strings.Replace(story, "# Export filtered rows to CSV", "# "+words(11), 1),
			want:  []string{"1:ticket-title"},
		},
		{
			name:  "empty section flagged",
			input: story + "\n## Links\n\n---\n\n## Risks\n\nNone known.\n",
			want:  []string{"15:ticket-empty-section"},
		},
		{
			name:  "heading with subheadings is not empty",
			input: story + "\n## Links\n\n### Docs\n\nSee the export guide.\n",
			want:  []string{},
		},
		{
			name:  "placeholders flagged",
			input: story + "\n## Environment\n\nProduction or TBD.\n\nPick: [option A / option B]\n",
			want:  []string{"17:ticket-placeholder", "19:ticket-placeholder"},
		},
		{
			name:  "links, checkboxes, and alerts are not placeholders",
			input: story + "\n## Links\n\n- [x] See [the guide](https://example.com).\n\n> [!NOTE]\n> Read it first.\n",
			want:  []string{},
		},
		{
			name:  "vague criterion is a finding",
			input: strings.Replace(story, "downloads the filtered rows as CSV", "works correctly", 1),
			want:  []string{"13:ticket-vague", "13:ticket-vague"},
		},
		{
			name:  "bold-line headings count",
			input: "**Steps to Reproduce**\n\n1. Click Save.\n\n**Expected Behavior**\n\nIt saves.\n\n**Actual Behavior**\n\nIt crashes.\n\n**Environment**\n\nProduction.\n",
			want:  []string{},
		},
		{
			name:  "classic user story satisfies problem and outcome",
			input: "# Export rows\n\n## User story\n\nAs a support agent, I want to export rows, so that I stop copying by hand.\n\n## Acceptance criteria\n\n- [ ] The Export button downloads a CSV.\n",
			want:  []string{},
		},
		{
			name:  "true epic needs an end state",
			input: "# Billing v2\n\n## Problem\n\nInvoices are manual.\n\n## Outcome\n\nInvoices send on their own.\n\n## Stories\n\n- Generate invoices.\n",
			want:  []string{"1:ticket-section"},
		},
		{
			name:  "bucket epic needs no end state",
			input: "# Platform upkeep\n\n## Purpose\n\nSmall platform fixes.\n\n## In scope\n\n- Tool upgrades.\n\n## Out of scope\n\n- Product features.\n\n## Entry rule\n\nA problem and an owner.\n\n## Review cadence\n\nMonthly.\n",
			mode:  "epic",
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lintTicketString(tt.input, tt.mode)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormalizeHeading(t *testing.T) {
	tests := map[string]string{
		"**⚠️ Risks / Edge Cases**":                "risks edge cases",
		"Acceptance Criteria / Definition of Done": "acceptance criteria definition of done",
		"What are we doing?":                       "what are we doing",
		"Scope & Constraints":                      "scope and constraints",
	}
	for in, want := range tests {
		if got := normalizeHeading(in); got != want {
			t.Errorf("normalizeHeading(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestTicketFlag(t *testing.T) {
	var f ticketFlag
	if err := f.Set("true"); err != nil || f.mode != "any" {
		t.Errorf("bare --ticket: got %q, %v", f.mode, err)
	}
	if err := f.Set("Bug"); err != nil || f.mode != "bug" {
		t.Errorf("--ticket=Bug: got %q, %v", f.mode, err)
	}
	if err := f.Set("saga"); err == nil {
		t.Error("--ticket=saga: want an error")
	}
}
