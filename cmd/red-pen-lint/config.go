package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// .red-pen.yml is optional repo config. The linter reads only the lint keys:
// ticket-detect, max-title-words, and the per-type section lists, such as
// story-sections. Every other key, such as tracker, is agent-facing and
// ignored here. The nearest file walking up from the linted file wins.
// Command-line flags win over config.

const configName = ".red-pen.yml"

// ticketConfig holds what the linter reads from a config file. maxTitleWords
// at 0 means the built-in default. sections is a custom required list per
// ticket type, and is empty when the repo uses the built-in templates.
type ticketConfig struct {
	autodetect    bool
	maxTitleWords int
	sections      map[string][]section
}

var defaultTicketConfig = ticketConfig{
	autodetect: true,
	sections:   map[string][]section{},
}

func parseConfig(text string) (ticketConfig, error) {
	// A fresh map for every file. defaultTicketConfig must never gain keys,
	// or one repo's custom sections would leak into every other lint run.
	cfg := ticketConfig{autodetect: true, sections: map[string][]section{}}
	for i, raw := range strings.Split(text, "\n") {
		line := strings.TrimSpace(stripConfigComment(raw))
		if line == "" {
			continue
		}
		m := configLine.FindStringSubmatch(line)
		if m == nil {
			return cfg, fmt.Errorf("line %d: expected \"key: value\"", i+1)
		}
		if err := cfg.apply(m[1], m[2], i+1); err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}

// stripConfigComment cuts a comment at the first # so values stay clean.
func stripConfigComment(line string) string {
	if i := strings.Index(line, "#"); i >= 0 {
		line = line[:i]
	}
	return line
}

var configLine = regexp.MustCompile(`^([A-Za-z][A-Za-z-]*)\s*:\s*(\S.*)$`)

func (c *ticketConfig) apply(key, val string, lineNum int) error {
	switch key {
	case "ticket-detect":
		switch val {
		case "auto", "true":
			c.autodetect = true
		case "off", "false":
			c.autodetect = false
		default:
			return fmt.Errorf("line %d: ticket-detect must be auto or off, not %q", lineNum, val)
		}
	case "max-title-words":
		n, err := strconv.Atoi(val)
		if err != nil || n < 1 {
			return fmt.Errorf("line %d: max-title-words must be a number of 1 or more, not %q", lineNum, val)
		}
		c.maxTitleWords = n
	default:
		if t, ok := strings.CutSuffix(key, "-sections"); ok {
			return c.applySections(t, val, lineNum)
		}
	}
	return nil
}

// applySections records a custom section list for one ticket type, so a team
// can pin the headings its own tickets need. Names are comma separated. The
// headings match by normalized text, and the built-in aliases stop applying
// to that type.
func (c *ticketConfig) applySections(ticketType, val string, lineNum int) error {
	if _, ok := ticketTypes[ticketType]; !ok {
		return fmt.Errorf("line %d: unknown ticket type %q: use story, bug, task, spike, or epic", lineNum, ticketType)
	}
	var secs []section
	for _, name := range strings.Split(val, ",") {
		name = strings.TrimSpace(name)
		if name == "" {
			return fmt.Errorf("line %d: %s-sections has an empty section name", lineNum, ticketType)
		}
		norm := normalizeHeading(name)
		secs = append(secs, section{name: name, aliases: []string{norm}})
	}
	c.sections[ticketType] = secs
	return nil
}

// configFor returns the config for the repo that owns path. Without a config
// file anywhere up the tree it returns the defaults.
func configFor(path string) (ticketConfig, error) {
	dir := filepath.Dir(path)
	for {
		f := filepath.Join(dir, configName)
		if s, err := os.Stat(f); err == nil && !s.IsDir() {
			// #nosec G304 -- reading this repo's own config file is the
			// point. os.Root scoping would defeat the walk up the tree.
			text, rerr := os.ReadFile(f)
			if rerr != nil {
				return defaultTicketConfig, rerr
			}
			cfg, perr := parseConfig(string(text))
			if perr != nil {
				return defaultTicketConfig, fmt.Errorf("%s: %w", f, perr)
			}
			return cfg, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return defaultTicketConfig, nil
		}
		dir = parent
	}
}
