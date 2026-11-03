package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		want    ticketConfig
		wantErr bool
	}{
		{
			name: "defaults from an empty file",
			text: "",
			want: ticketConfig{autodetect: true, sections: map[string][]section{}},
		},
		{
			name: "agent-facing keys are ignored",
			text: "tracker: linear\n",
			want: ticketConfig{autodetect: true, sections: map[string][]section{}},
		},
		{
			name: "comments and blank lines are skipped",
			text: "# repo config\n\nticket-detect: off # checked by CI\n",
			want: ticketConfig{autodetect: false, sections: map[string][]section{}},
		},
		{
			name: "true and false work like yaml booleans",
			text: "ticket-detect: false\n",
			want: ticketConfig{autodetect: false, sections: map[string][]section{}},
		},
		{
			name: "title cap from config",
			text: "max-title-words: 8\n",
			want: ticketConfig{autodetect: true, maxTitleWords: 8, sections: map[string][]section{}},
		},
		{
			name:    "a bad detect value is an error",
			text:    "ticket-detect: sometimes\n",
			wantErr: true,
		},
		{
			name:    "a bad title cap is an error",
			text:    "max-title-words: ten\n",
			wantErr: true,
		},
		{
			name:    "a title cap under 1 is an error",
			text:    "max-title-words: 0\n",
			wantErr: true,
		},
		{
			name: "custom sections per type",
			text: "story-sections: Problem, Outcome, Test notes\nbug-sections: Steps\n",
			want: ticketConfig{
				autodetect: true,
				sections: map[string][]section{
					"story": {{"Problem", []string{"problem"}}, {"Outcome", []string{"outcome"}}, {"Test notes", []string{"test notes"}}},
					"bug":   {{"Steps", []string{"steps"}}},
				},
			},
		},
		{
			name:    "sections for an unknown type are an error",
			text:    "saga-sections: Problem\n",
			wantErr: true,
		},
		{
			name:    "an empty section name is an error",
			text:    "story-sections: Problem, , Outcome\n",
			wantErr: true,
		},
		{
			name:    "a line without a value is an error",
			text:    "tracker:\n",
			wantErr: true,
		},
		{
			name:    "a block list is not supported",
			text:    "story-sections:\n- Problem\n",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfig(tt.text)
			if tt.wantErr {
				if err == nil {
					t.Errorf("parseConfig(%q): want an error, got %+v", tt.text, got)
				}
				return
			}
			if err != nil {
				t.Errorf("parseConfig(%q): %v", tt.text, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseConfig(%q) = %+v, want %+v", tt.text, got, tt.want)
			}
		})
	}
}

func TestConfigFor(t *testing.T) {
	t.Run("no config file gives the defaults", func(t *testing.T) {
		cfg, err := configFor(filepath.Join(t.TempDir(), "doc.md"))
		if err != nil {
			t.Fatalf("configFor: %v", err)
		}
		if !cfg.autodetect || len(cfg.sections) != 0 {
			t.Errorf("configFor without a file = %+v, want the defaults", cfg)
		}
	})

	t.Run("the nearest config file wins over one further up", func(t *testing.T) {
		root, sub := t.TempDir(), filepath.Join(t.TempDir(), "docs")
		if err := os.MkdirAll(sub, 0o755); err != nil {
			t.Fatal(err)
		}
		write := func(dir, text string) {
			t.Helper()
			if err := os.WriteFile(filepath.Join(dir, configName), []byte(text), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		write(root, "ticket-detect: off\n")
		// The sub file does not repeat the keys above it. It replaces the
		// root file entirely, so its unset keys are the defaults.
		write(sub, "max-title-words: 6\n")
		cfg, err := configFor(filepath.Join(sub, "ticket.md"))
		if err != nil {
			t.Fatalf("configFor: %v", err)
		}
		if !cfg.autodetect || cfg.maxTitleWords != 6 {
			t.Errorf("configFor into %q = %+v", sub, cfg)
		}
		rootCfg, err := configFor(filepath.Join(root, "doc.md"))
		if err != nil {
			t.Fatalf("configFor: %v", err)
		}
		if rootCfg.autodetect || rootCfg.maxTitleWords != 0 {
			t.Errorf("configFor into %q = %+v", root, rootCfg)
		}
	})

	t.Run("a bad config file is an error", func(t *testing.T) {
		root := t.TempDir()
		if err := os.WriteFile(filepath.Join(root, configName), []byte("ticket-detect: maybe\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		if _, err := configFor(filepath.Join(root, "doc.md")); err == nil {
			t.Error("configFor: want an error for a bad value")
		}
	})
}

func TestCustomSectionsInLint(t *testing.T) {
	cfg, err := parseConfig("story-sections: Problem, Test notes\nmax-title-words: 5\n")
	if err != nil {
		t.Fatal(err)
	}
	input := strings.Join([]string{
		"# Export rows and rows", // 5 words: at the cap, not over.
		"",
		"## Problem",
		"",
		"Support copies rows by hand.",
		"",
		"## Test notes",
		"",
		"- [ ] The Export button downloads a CSV.",
	}, "\n")

	// The custom list replaces the built-in list, so no Outcome violation.
	got := hits(lintTicket(strings.Split(input, "\n"), "", cfg))
	want := []string{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("lintTicket with custom sections = %v, want %v", got, want)
	}

	// The custom cap applies: 6 words would fail even though 10 is the default.
	over := strings.Replace(input, "# Export rows and rows", "# Export rows and rows and rows", 1)
	if got := hits(lintTicket(strings.Split(over, "\n"), "", cfg)); len(got) != 1 || !strings.Contains(got[0], "ticket-title") {
		t.Errorf("lintTicket with a custom cap = %v, want one ticket-title hit", got)
	}

	// A file with one of two custom sections is not confidently the team's
	// ticket, so auto-detect stays quiet. Forcing the type reports the miss
	// by the team's own name.
	thin := strings.Join([]string{"# Export rows", "", "## Problem", "", "Support copies rows by hand.", ""}, "\n")
	got = hits(lintTicket(strings.Split(thin, "\n"), "", cfg))
	if len(got) != 0 {
		t.Errorf("lintTicket on a thin story, auto = %v, want no checks", got)
	}
	got = hits(lintTicket(strings.Split(thin, "\n"), "story", cfg))
	want = []string{"1:ticket-section"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("lintTicket on a thin story, forced = %v, want %v", got, want)
	}
}

func TestConfigOffKeepsTicketChecksQuiet(t *testing.T) {
	cfg, err := parseConfig("ticket-detect: off\n")
	if err != nil {
		t.Fatal(err)
	}
	story := "# Export rows\n\n## Problem\n\nSupport copies rows by hand.\n\n## Outcome\n\nThe table exports to CSV.\n\n## Acceptance criteria\n\n- [ ] The Export button downloads a CSV.\n"
	if got := hits(lintTicket(strings.Split(story, "\n"), "", cfg)); len(got) != 0 {
		t.Errorf("auto-detect off: %v, want no checks", got)
	}
	// An explicit --ticket flag still forces the checks. The type flag wins.
	if got := hits(lintTicket(strings.Split(story, "\n"), "story", cfg)); len(got) != 0 {
		t.Errorf("--ticket with detect off: %v, want a clean pass", got)
	}
}
