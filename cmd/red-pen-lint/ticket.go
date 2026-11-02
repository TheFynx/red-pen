package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Ticket mode checks a story, bug, task, spike, or epic against the templates
// in references/tickets.md. It runs when --ticket is passed, or when a file
// covers three sections of one ticket type.

// Ticket titles are scanned in a list view, so they stay short.
const maxTitleWords = 10

// Auto-detect waits until this many of a type's sections are present. Two
// built-in matches are not enough: "Problem" and "Outcome" also appear in
// ordinary design docs that are not tickets. A custom section list may hold
// fewer sections, and its own names are stronger signals, so the gate drops
// to the list length.
const minAutoSections = 3

type section struct {
	name    string   // the heading the template uses, shown in messages
	aliases []string // normalized heading texts that count as this section
}

var (
	// "User story" holds the "As a / I want / so that" sentence, which states
	// both the problem and the outcome.
	secProblem     = section{"Problem", []string{"problem", "problem statement", "summary", "user story"}}
	secOutcome     = section{"Outcome", []string{"outcome", "outcomes", "desired outcome", "desired outcomes", "user story"}}
	secCriteria    = section{"Acceptance criteria", []string{"acceptance criteria", "acceptance and validation criteria", "acceptance criteria definition of done", "definition of done", "done when"}}
	secSteps       = section{"Steps to reproduce", []string{"steps to reproduce", "reproduction steps", "repro steps"}}
	secExpected    = section{"Expected", []string{"expected", "expected behavior", "expected behaviour", "expected result"}}
	secActual      = section{"Actual", []string{"actual", "actual behavior", "actual behaviour", "actual result"}}
	secEnvironment = section{"Environment", []string{"environment"}}
	secWhat        = section{"What", []string{"what", "what changes", "what are we doing"}}
	secWhy         = section{"Why", []string{"why", "why are we doing it"}}
	secDone        = section{"Done when", []string{"done when", "definition of done", "acceptance criteria", "acceptance and validation criteria", "acceptance criteria definition of done"}}
	secQuestion    = section{"Question", []string{"question", "research question"}}
	secTimebox     = section{"Timebox", []string{"timebox", "time box"}}
	secDeliverable = section{"Deliverable", []string{"deliverable", "deliverables"}}
	secEndState    = section{"End state", []string{"end state"}}
	secStories     = section{"Stories", []string{"stories", "story breakdown"}}
	secPurpose     = section{"Purpose", []string{"purpose"}}
	secInScope     = section{"In scope", []string{"in scope", "scope"}}
	secOutScope    = section{"Out of scope", []string{"out of scope", "scope"}}
	secEntryRule   = section{"Entry rule", []string{"entry rule", "entry rules"}}
	secCadence     = section{"Review cadence", []string{"review cadence"}}
)

var ticketTypes = map[string][]section{
	"story": {secProblem, secOutcome, secCriteria},
	"bug":   {secSteps, secExpected, secActual, secEnvironment},
	"task":  {secWhat, secWhy, secDone},
	"spike": {secQuestion, secTimebox, secDeliverable, secDone},
	"epic":  {secProblem, secOutcome, secEndState, secStories},
}

// A bucket epic holds ad hoc work with no finish line, so it has no end state.
var bucketEpic = []section{secPurpose, secInScope, secOutScope, secEntryRule, secCadence}

// Any of these headings marks an epic as a bucket. Scope headings do not: a
// true epic can have a Scope section too.
var bucketMarkers = []section{secPurpose, secEntryRule, secCadence}

// detectType walks this list, so ties go to the earlier type.
var typeOrder = []string{"bug", "spike", "epic", "story", "task"}

// requiredFor returns the sections a ticket of this type must have. A custom
// list from .red-pen.yml replaces the built-in list, including the bucket
// epic handling.
func requiredFor(ticketType string, hs []heading, cfg ticketConfig) []section {
	if ss, ok := cfg.sections[ticketType]; ok {
		return ss
	}
	if ticketType == "epic" {
		if found, _ := present(bucketMarkers, hs); len(found) > 0 {
			return bucketEpic
		}
	}
	return ticketTypes[ticketType]
}

// Words that make an acceptance criterion untestable. A reviewer cannot pass
// or fail "works correctly". These are findings: the word can be fine in context.
var vagueCriteria = []string{
	"works", "working", "proper", "properly", "correctly", "reliable", "reliably", "as expected", "as needed",
	"appropriate", "appropriately", "user-friendly", "fast", "quickly",
	"better", "improved", "reasonable", "near zero", "seamlessly", "etc",
}

var vagueRules = func() []*regexp.Regexp {
	var out []*regexp.Regexp
	for _, w := range vagueCriteria {
		out = append(out, wordPattern(w))
	}
	return out
}()

var (
	atxHeading  = regexp.MustCompile(`^(#{1,6})\s+(.*?)\s*#*$`)
	boldHeading = regexp.MustCompile(`^(?:\*\*|__)([^*_]+)(?:\*\*|__):?$`)
	thematic    = regexp.MustCompile(`^([-*_])(\s*[-*_]){2,}$`)

	placeholderWord = regexp.MustCompile(`(?i)\b(TBD|TBA|TBC|to be (?:filled|defined|determined|decided|added|confirmed)(?: in)?|will be defined|lorem ipsum)\b`)
	// [describe the problem] or [option A / option B]. Skips links, checkboxes, footnotes,
	// alert markers ([!NOTE]), and reference definitions.
	placeholderBracket = regexp.MustCompile(`\[([^\]\[!^][^\]\[]+)\](?:[^(\[:]|$)`)
)

type heading struct {
	line  int    // 1-based
	level int    // 1-6 for # headings, 7 for a bold-only line
	text  string // normalized
}

// normalizeHeading lowercases the text and drops markup, emoji, and punctuation,
// so "### **⚠️ Risks / Edge Cases**" becomes "risks edge cases".
func normalizeHeading(s string) string {
	s = strings.ToLower(strings.ReplaceAll(s, "&", " and "))
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == '-':
			b.WriteRune(r)
		default:
			b.WriteRune(' ')
		}
	}
	return strings.Join(strings.Fields(b.String()), " ")
}

func parseHeading(trimmed string) (level int, text string, ok bool) {
	if m := atxHeading.FindStringSubmatch(trimmed); m != nil {
		return len(m[1]), m[2], true
	}
	if m := boldHeading.FindStringSubmatch(trimmed); m != nil {
		return 7, m[1], true
	}
	return 0, "", false
}

func headings(text []string) []heading {
	var out []heading
	for i, line := range text {
		if level, t, ok := parseHeading(strings.TrimSpace(line)); ok {
			out = append(out, heading{i + 1, level, normalizeHeading(t)})
		}
	}
	return out
}

func (s section) matches(h heading) bool {
	for _, a := range s.aliases {
		if h.text == a {
			return true
		}
	}
	return false
}

func present(sections []section, hs []heading) (found, missing []section) {
	for _, s := range sections {
		hit := false
		for _, h := range hs {
			if s.matches(h) {
				hit = true
				break
			}
		}
		if hit {
			found = append(found, s)
		} else {
			missing = append(missing, s)
		}
	}
	return found, missing
}

// detectType picks the type whose required sections the file covers best.
// It walks typeOrder, so ties go to the earlier type. It returns "" and a
// covered count of 0 when no ticket heading matches.
func detectType(hs []heading, cfg ticketConfig) (string, int) {
	best, bestScore := "", 0
	for _, t := range typeOrder {
		found, _ := present(requiredFor(t, hs, cfg), hs)
		if len(found) > bestScore {
			best, bestScore = t, len(found)
		}
	}
	return best, bestScore
}

// lintTicket runs the ticket checks. mode is "" for auto-detect, "any" for
// --ticket with no type, or a type name.
func lintTicket(lines []string, mode string, cfg ticketConfig) []violation {
	text := prose(lines)
	hs := headings(text)

	ticketType, covered := mode, 0
	if mode == "" || mode == "any" {
		if mode == "" && !cfg.autodetect {
			return nil
		}
		ticketType, covered = detectType(hs, cfg)
		gate := minAutoSections
		if ss, ok := cfg.sections[ticketType]; ok && len(ss) < gate {
			gate = len(ss)
		}
		// Auto-detect stays out of ordinary docs. --ticket ("any") always runs.
		if mode == "" && covered < gate {
			return nil
		}
	}

	titleCap := cfg.maxTitleWords
	if titleCap == 0 {
		titleCap = maxTitleWords
	}

	var vs []violation
	titleLine := 1
	for i, line := range text {
		if m := atxHeading.FindStringSubmatch(strings.TrimSpace(line)); m != nil && len(m[1]) == 1 {
			titleLine = i + 1
			if wc := wordCount(strings.NewReplacer("*", "", "_", "").Replace(m[2])); wc > titleCap {
				vs = append(vs, violation{
					Line:     titleLine,
					Rule:     "ticket-title",
					Message:  fmt.Sprintf("Ticket title has %d words (max %d). Name the change, not the background.", wc, titleCap),
					Severity: "violation",
				})
			}
			break
		}
	}

	if ticketType == "" {
		vs = append(vs, violation{
			Line:     titleLine,
			Rule:     "ticket-type",
			Message:  "Cannot tell the ticket type from the headings. Pass --ticket=story, bug, task, spike, or epic.",
			Severity: "violation",
		})
	} else {
		_, missing := present(requiredFor(ticketType, hs, cfg), hs)
		for _, s := range missing {
			vs = append(vs, violation{
				Line:     titleLine,
				Rule:     "ticket-section",
				Match:    s.name,
				Message:  fmt.Sprintf("%s is missing the \"%s\" section. See references/tickets.md.", ticketType, s.name),
				Severity: "violation",
			})
		}
	}

	vs = append(vs, emptySections(lines, hs)...)
	vs = append(vs, placeholders(text)...)
	vs = append(vs, vagueInCriteria(text, hs)...)
	return vs
}

// emptySections flags a heading with no content before the next heading at the
// same or a higher level. A heading followed by its own subheadings has content.
func emptySections(lines []string, hs []heading) []violation {
	var vs []violation
	for i, h := range hs {
		end := len(lines)
		for _, next := range hs[i+1:] {
			if next.level <= h.level {
				end = next.line - 1
				break
			}
		}
		content := false
		for j := h.line; j < end; j++ {
			raw := strings.TrimSpace(lines[j])
			if raw == "" || thematic.MatchString(raw) {
				continue
			}
			content = true
			break
		}
		if !content {
			vs = append(vs, violation{
				Line:     h.line,
				Rule:     "ticket-empty-section",
				Message:  "Empty section. Fill it or delete the heading.",
				Severity: "violation",
			})
		}
	}
	return vs
}

func placeholders(text []string) []violation {
	var vs []violation
	for i, line := range text {
		m := placeholderWord.FindString(line)
		if m == "" {
			if b := placeholderBracket.FindStringSubmatch(line); b != nil {
				m = "[" + b[1] + "]"
			}
		}
		if m != "" {
			vs = append(vs, violation{
				Line:     i + 1,
				Rule:     "ticket-placeholder",
				Match:    m,
				Message:  fmt.Sprintf("Template placeholder \"%s\" left in. Replace it with the real value or delete the line.", m),
				Severity: "violation",
			})
		}
	}
	return vs
}

// vagueInCriteria checks the lines under an acceptance criteria or done-when heading.
func vagueInCriteria(text []string, hs []heading) []violation {
	var vs []violation
	for i, h := range hs {
		if !secCriteria.matches(h) && !secDone.matches(h) {
			continue
		}
		end := len(text)
		if i+1 < len(hs) {
			end = hs[i+1].line - 1
		}
		for j := h.line; j < end; j++ {
			for k, re := range vagueRules {
				if re.MatchString(text[j]) {
					vs = append(vs, violation{
						Line:     j + 1,
						Rule:     "ticket-vague",
						Match:    vagueCriteria[k],
						Message:  fmt.Sprintf("Vague criterion \"%s\". Name the value or behavior a reviewer can check.", vagueCriteria[k]),
						Severity: "finding",
					})
				}
			}
		}
	}
	return vs
}

// ticketFlag is --ticket. Alone, it forces ticket mode and detects the type.
// --ticket=<type> also sets the type.
type ticketFlag struct{ mode string }

func (t *ticketFlag) String() string   { return t.mode }
func (t *ticketFlag) IsBoolFlag() bool { return true }

func (t *ticketFlag) Set(s string) error {
	s = strings.ToLower(s)
	if s == "true" {
		t.mode = "any"
		return nil
	}
	if _, ok := ticketTypes[s]; !ok {
		return fmt.Errorf("unknown ticket type %q: use story, bug, task, spike, or epic", s)
	}
	t.mode = s
	return nil
}
