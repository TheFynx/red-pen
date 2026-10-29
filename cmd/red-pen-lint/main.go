package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

type violation struct {
	Path     string
	Line     int
	Rule     string
	Match    string
	Message  string
	Severity string // "violation" (must fix) or "finding" (suggestion)
}

func wordPattern(word string) *regexp.Regexp {
	return regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(word) + `\b`)
}

// Sentence limits from SKILL.md: 20 words for procedural steps, 25 for everything else.
const (
	maxProceduralWords  = 20
	maxDescriptiveWords = 25
)

var (
	emDash = regexp.MustCompile(`\x{2014}|(^|[^-])--([^-]|$)`)
	enDash = regexp.MustCompile(`\x{2013}`)
	// An en dash between two numbers is a range ("pages 1–5"), not a sentence dash.
	enDashRange = regexp.MustCompile(`\d\s?\x{2013}\s?\d`)
	semicolon   = regexp.MustCompile(`;`)

	// **Term:** definition pattern (bolded-term-colon-definition). The colon can
	// sit inside the bold (**Term:** def) or after it (**Term**: def).
	boldColonDef = regexp.MustCompile(`\*\*[^*]+(?:\*\*:\s|:\*\*\s)`)

	// Contractions. Docs voice writes every one out. The personal voice mode
	// skips this check. o'clock is a set phrase, not a contraction, so it stays
	// off the list. Possessive apostrophes ("the server's port") stay off too:
	// a general 's pattern would flag them.
	contractionRe = regexp.MustCompile(`(?i)\b(can't|won't|don't|doesn't|isn't|aren't|wasn't|weren't|wouldn't|shouldn't|couldn't|mustn't|needn't|shan't|daren't|oughtn't|mightn't|hasn't|haven't|hadn't|didn't|it's|that's|there's|here's|what's|who's|where's|when's|why's|how's|let's|we're|they're|you're|there're|we've|they've|you've|would've|should've|could've|must've|might've|who've|i'm|he's|she's|we'll|they'll|you'll|i'll|he'll|she'll|i'd|you'd|he'd|she'd|we'd|they'd|it'd|who'd|that'd|there'd|y'all|gonna|wanna|gotta|kinda|sorta|dunno|lemme|gimme|outta)\b`)

	numberedItem = regexp.MustCompile(`^\d+[.)]\s+`)
	bulletItem   = regexp.MustCompile(`^[-*+]\s+`)

	// Sentence boundaries: end punctuation, optional closing marks, then whitespace.
	sentenceSplit = regexp.MustCompile(`[.!?]["')*_\]]*\s+`)
	// Abbreviations whose period does not end a sentence.
	abbreviation = regexp.MustCompile(`(?i)\b(e\.g|i\.e|vs|cf|approx|mr|mrs|ms|dr|fig|no)\.`)

	// Content stripped before any rule runs.
	inlineCode  = regexp.MustCompile("``[^`]*``|`[^`]*`")
	htmlComment = regexp.MustCompile(`<!--.*?-->`)
	htmlTag     = regexp.MustCompile(`</?[a-zA-Z][^>]*>`)
	htmlEntity  = regexp.MustCompile(`&(#[0-9]+|#x[0-9a-fA-F]+|[a-zA-Z]+);`)
	fenceOpen   = regexp.MustCompile("^(`{3,}|~{3,})")

	// <!-- red-pen-ignore-next-line -->, <!-- red-pen-ignore-start -->, <!-- red-pen-ignore-end -->
	ignoreDirective = regexp.MustCompile(`^<!--\s*red-pen-ignore-(next-line|start|end)\s*-->$`)
)

var bannedModals = []string{"should", "may", "might", "could", "would"}

var minimizingWords = []string{"just", "simply", "easily", "merely", "straightforward"}

var formulaicTransitions = []string{
	"moreover",
	"furthermore",
	"additionally",
	"it is worth noting",
	"it is important to note",
	"it should be noted",
}

var compulsiveSummaries = []string{
	"in summary",
	"in conclusion",
	"to sum up",
	"to summarize",
}

var cheerleading = []string{
	"great job",
	"awesome",
	"fantastic",
	"well done",
	"congratulations",
	"you're well on your way",
	"exciting",
}

var slopWords = []string{
	"delve",
	"tapestry",
	"landscape",
	"foster",
	"testament",
	"underscore",
	"multifaceted",
	"realm",
	"embark",
	"beacon",
	"spearhead",
	"groundbreaking",
	"cornerstone",
	"underpinning",
	"game-changer",
	"synergy",
	"paradigm shift",
	"deep dive",
	"unpack",
	"double-edged sword",
	"low-hanging fruit",
	"move the needle",
	"north star",
	"pain point",
}

var wordSwapVerbs = []string{
	"utilize",
	"leverage",
	"facilitate",
	"optimize",
	"enhance",
	"streamline",
	"empower",
	"initiate",
	"commence",
	"encompass",
	"necessitate",
	"ascertain",
	"endeavor",
	"remediate",
}

var wordSwapAdjectives = []string{
	"robust",
	"seamless",
	"comprehensive",
	"cutting-edge",
	"state-of-the-art",
	"innovative",
	"powerful",
	"elegant",
	"intuitive",
	"holistic",
	"pivotal",
	"crucial",
}

var wordSwapDeleteAdverbs = []string{
	"essentially",
	"basically",
	"effectively",
	"incredibly",
	"extremely",
	"truly",
	"certainly",
	"definitely",
	"absolutely",
	"notably",
	"importantly",
	"interestingly",
}

var wordSwapPhrases = []string{
	"in order to",
	"in terms of",
	"with respect to",
	"in regard to",
	"at the end of the day",
	"at this point in time",
	"for the purpose of",
	"in the event that",
	"a wide range of",
	"in a timely manner",
	"take into consideration",
	"is capable of",
	"has the ability to",
	"due to the fact that",
	"serves as a testament to",
	"plays a vital role in",
	"is a key component of",
	"provides a foundation for",
	"for all intents and purposes",
	"the vast majority of",
}

var promotionalWords = []string{
	"feature-rich",
	"best-in-class",
	"world-class",
	"enterprise-grade",
	"production-ready",
	"battle-tested",
	"blazing fast",
	"lightning fast",
}

type wordRule struct {
	re       *regexp.Regexp
	rule     string
	match    string
	message  string
	severity string
}

// Word and phrase rules, compiled once.
var wordRules = buildWordRules()

func buildWordRules() []wordRule {
	var rules []wordRule
	add := func(words []string, rule, format, severity string) {
		for _, w := range words {
			rules = append(rules, wordRule{wordPattern(w), rule, w, fmt.Sprintf(format, w), severity})
		}
	}
	add(bannedModals, "banned-modal", "Banned modal \"%s\". Use can, will, or must instead.", "violation")
	add(minimizingWords, "minimizing", "Minimizing word \"%s\". Delete it.", "violation")
	add(formulaicTransitions, "formulaic-transition", "Formulaic transition \"%s\". Start with the fact.", "violation")
	add(compulsiveSummaries, "compulsive-summary", "Compulsive summary \"%s\". Delete or rewrite as a fact.", "violation")
	add(cheerleading, "cheerleading", "Cheerleading \"%s\". Delete it.", "violation")
	add(slopWords, "ai-slop", "AI slop word \"%s\". See word-swaps.md for replacement.", "violation")
	// Word swaps are findings. A flagged word can name the technical thing, and
	// only the writer can tell, so the linter reports it and does not enforce.
	add(wordSwapVerbs, "word-swap", "Verbose word \"%s\". See word-swaps.md for replacement.", "finding")
	add(wordSwapAdjectives, "word-swap", "Verbose word \"%s\". See word-swaps.md for replacement.", "finding")
	add(wordSwapDeleteAdverbs, "word-swap", "Filler adverb \"%s\". Delete it.", "finding")
	add(wordSwapPhrases, "word-swap", "Verbose phrase \"%s\". See word-swaps.md for replacement.", "finding")
	add(promotionalWords, "promotional", "Promotional language \"%s\". State what it does instead.", "violation")
	return rules
}

// Rules the personal voice leaves to the author. See "Personal voice" in SKILL.md.
var personalVoiceRelaxed = map[string]bool{
	"contraction":     true,
	"banned-modal":    true,
	"minimizing":      true,
	"semicolon":       true,
	"em-dash":         true,
	"sentence-length": true,
}

// Stamp these at release time: ldflags -X main.version=... (goreleaser).
var version = "dev"

func main() {
	voice := flag.String("voice", "docs", "rule set: docs (all rules) or personal (keeps the author's voice)")
	versionFlag := flag.Bool("version", false, "print the build version and exit")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: red-pen-lint [--voice docs|personal] <file> [file...]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *versionFlag {
		fmt.Printf("red-pen-lint %s\n", version)
		os.Exit(0)
	}

	if flag.NArg() < 1 || (*voice != "docs" && *voice != "personal") {
		flag.Usage()
		os.Exit(2)
	}

	var violations, findings []violation

	for _, path := range flag.Args() {
		vs, err := lintFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s: %v\n", path, err)
			os.Exit(2)
		}
		for _, v := range filterVoice(vs, *voice) {
			if v.Severity == "finding" {
				findings = append(findings, v)
			} else {
				violations = append(violations, v)
			}
		}
	}

	printList(violations)

	if len(findings) > 0 {
		fmt.Printf("\n%d finding(s). Findings are suggestions, not failures:\n", len(findings))
		printList(findings)
		fmt.Println()
		fmt.Println("A finding names a word-swap candidate. Keep the word when it names the technical thing here, such as compiler or SQL optimization. Otherwise apply the swap from references/word-swaps.md.")
		fmt.Println("To drop a false positive from the report, put <!-- red-pen-ignore-next-line --> on the line above it, or wrap a region in <!-- red-pen-ignore-start --> and <!-- red-pen-ignore-end -->.")
	}

	switch {
	case len(violations) > 0:
		fmt.Printf("\n%d violation(s) found.\n", len(violations))
		os.Exit(1)
	case len(findings) > 0:
		fmt.Printf("\nNo violations. %d finding(s) to review.\n", len(findings))
		os.Exit(0)
	default:
		fmt.Println("Clean.")
		os.Exit(0)
	}
}

func printList(vs []violation) {
	for _, v := range vs {
		fmt.Printf("%s:%d: [%s] %s\n", v.Path, v.Line, v.Rule, v.Message)
	}
}

func filterVoice(violations []violation, voice string) []violation {
	if voice != "personal" {
		return violations
	}
	var kept []violation
	for _, v := range violations {
		if !personalVoiceRelaxed[v.Rule] {
			kept = append(kept, v)
		}
	}
	return kept
}

func lintFile(path string) ([]violation, error) {
	// #nosec G304 -- opening the file paths the user passes is the entire point
	// of this tool. os.Root scoping would defeat a multi-file linter.
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Fprintf(os.Stderr, "warning: %s: close failed: %v\n", path, cerr)
		}
	}()

	lines, err := readLines(f)
	if err != nil {
		return nil, err
	}
	vs := lintLines(lines)
	for i := range vs {
		vs[i].Path = path
	}
	return vs, nil
}

func readLines(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func lintLines(lines []string) []violation {
	text := prose(lines)
	var violations []violation

	for i, line := range text {
		lineNum := i + 1
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// --- Structural checks (apply to all prose lines) ---

		if emDash.MatchString(line) {
			violations = append(violations, violation{
				Line:     lineNum,
				Rule:     "em-dash",
				Message:  "Em dash found. Use a period, comma, or conjunction instead.",
				Severity: "violation",
			})
		} else if enDash.MatchString(enDashRange.ReplaceAllString(line, "")) {
			violations = append(violations, violation{
				Line:     lineNum,
				Rule:     "em-dash",
				Message:  "En dash used as a sentence dash. Use a period, comma, or conjunction instead. En dashes are fine in number ranges.",
				Severity: "violation",
			})
		}

		if !isTableSeparator(trimmed) && semicolon.MatchString(line) {
			violations = append(violations, violation{
				Line:     lineNum,
				Rule:     "semicolon",
				Message:  "Semicolon found. Write two sentences.",
				Severity: "violation",
			})
		}

		if boldColonDef.MatchString(line) && !isLabelShape(trimmed) {
			violations = append(violations, violation{
				Line:     lineNum,
				Rule:     "bold-colon-def",
				Message:  "Bolded-term-colon-definition pattern. Use a table, a bullet list, or work the definition into the sentence.",
				Severity: "violation",
			})
		}

		for _, m := range contractionRe.FindAllString(line, -1) {
			violations = append(violations, violation{
				Line:     lineNum,
				Rule:     "contraction",
				Match:    m,
				Message:  fmt.Sprintf("Contraction \"%s\" found. Write it out.", m),
				Severity: "violation",
			})
		}

		// Headings name topics, so they can contain flagged words.
		if strings.HasPrefix(trimmed, "#") {
			continue
		}

		// --- Word and phrase checks (prose and table cells) ---

		for _, r := range wordRules {
			if r.re.MatchString(line) {
				violations = append(violations, violation{
					Line:     lineNum,
					Rule:     r.rule,
					Match:    r.match,
					Message:  r.message,
					Severity: r.severity,
				})
			}
		}
	}

	return append(violations, checkSentenceLength(text)...)
}

// prose returns one entry per source line, with everything the rules do not
// apply to blanked out: frontmatter, fenced and indented code, HTML blocks and
// comments, inline code, and ignored regions. Indexes stay aligned with source lines.
func prose(lines []string) []string {
	out := make([]string, len(lines))
	inFrontmatter := len(lines) > 0 && strings.TrimSpace(lines[0]) == "---"
	fence := "" // the open fence marker, empty outside a fenced block
	inComment, inIgnore, skipNext := false, false, false
	inIndentedCode, inList := false, false
	prevBlank := true

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		blank := trimmed == ""
		wasBlank := prevBlank
		prevBlank = blank

		if inFrontmatter {
			if i > 0 && (trimmed == "---" || trimmed == "...") {
				inFrontmatter = false
			}
			continue
		}

		// Fenced code: ``` or ~~~, closed by the same character at least as long.
		if fence != "" {
			if m := fenceOpen.FindString(trimmed); m != "" && m[0] == fence[0] &&
				len(m) >= len(fence) && strings.TrimSpace(trimmed[len(m):]) == "" {
				fence = ""
			}
			continue
		}
		if m := fenceOpen.FindString(trimmed); m != "" {
			fence = m
			continue
		}

		// Indented code: four spaces or a tab after a blank line, outside a list.
		indented := strings.HasPrefix(line, "    ") || strings.HasPrefix(line, "\t")
		if inIndentedCode {
			if blank || indented {
				continue
			}
			inIndentedCode = false
		}
		if indented && !blank && wasBlank && !inList {
			inIndentedCode = true
			continue
		}

		// A list runs until a non-indented line that is not a list item.
		if bulletItem.MatchString(trimmed) || numberedItem.MatchString(trimmed) {
			inList = true
		} else if !blank && !indented {
			inList = false
		}

		if m := ignoreDirective.FindStringSubmatch(trimmed); m != nil {
			switch m[1] {
			case "next-line":
				skipNext = true
			case "start":
				inIgnore = true
			case "end":
				inIgnore = false
			}
			continue
		}
		if inIgnore {
			continue
		}
		if skipNext {
			skipNext = false
			continue
		}

		// HTML comments, including ones that span lines.
		if inComment {
			end := strings.Index(line, "-->")
			if end < 0 {
				continue
			}
			line = line[end+3:]
			inComment = false
		}
		line = htmlComment.ReplaceAllString(line, "")
		if start := strings.Index(line, "<!--"); start >= 0 {
			line = line[:start]
			inComment = true
		}

		// A line that opens with a tag is part of an HTML block, not prose.
		if strings.HasPrefix(strings.TrimSpace(line), "<") {
			continue
		}

		line = inlineCode.ReplaceAllString(line, "code")
		line = htmlTag.ReplaceAllString(line, " ")
		line = htmlEntity.ReplaceAllString(line, " ")
		// Editors often type curly apostrophes. Normalize so contractions match.
		line = strings.ReplaceAll(line, "’", "'")
		out[i] = line
	}
	return out
}

func checkSentenceLength(lines []string) []violation {
	var violations []violation

	check := func(text string, lineNum, max int, kind string) {
		for _, s := range splitSentences(text) {
			if wc := wordCount(s); wc > max {
				violations = append(violations, violation{
					Line:     lineNum,
					Rule:     "sentence-length",
					Message:  fmt.Sprintf("Sentence has %d words (max %d %s): \"%s\"", wc, max, kind, truncate(s, 80)),
					Severity: "violation",
				})
			}
		}
	}

	var current strings.Builder
	currentStart := 0
	flush := func() {
		if current.Len() > 0 {
			check(current.String(), currentStart, maxDescriptiveWords, "descriptive")
			current.Reset()
		}
	}

	for i, line := range lines {
		lineNum := i + 1
		trimmed := strings.TrimSpace(line)

		switch {
		case trimmed == "" || strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "!["):
			flush()
		case strings.HasPrefix(trimmed, "|"):
			flush()
			if !isTableSeparator(trimmed) {
				for _, cell := range strings.Split(trimmed, "|") {
					check(cell, lineNum, maxDescriptiveWords, "descriptive")
				}
			}
		case numberedItem.MatchString(trimmed):
			// Numbered items are steps, so the procedural limit applies.
			flush()
			check(numberedItem.ReplaceAllString(trimmed, ""), lineNum, maxProceduralWords, "procedural")
		case bulletItem.MatchString(trimmed):
			flush()
			check(bulletItem.ReplaceAllString(trimmed, ""), lineNum, maxDescriptiveWords, "descriptive")
		case strings.HasPrefix(trimmed, ">"):
			flush()
			check(strings.TrimSpace(strings.TrimLeft(trimmed, ">")), lineNum, maxDescriptiveWords, "descriptive")
		default:
			if current.Len() == 0 {
				currentStart = lineNum
			} else {
				current.WriteString(" ")
			}
			current.WriteString(trimmed)
		}
	}
	flush()

	return violations
}

func splitSentences(text string) []string {
	// Drop periods that do not end a sentence (e.g., i.e., vs.) before splitting.
	text = abbreviation.ReplaceAllStringFunc(text, func(m string) string {
		return strings.ReplaceAll(m, ".", "")
	})
	var sentences []string
	for _, p := range sentenceSplit.Split(text, -1) {
		if p = strings.TrimSpace(p); p != "" {
			sentences = append(sentences, p)
		}
	}
	return sentences
}

func wordCount(s string) int {
	return len(strings.Fields(s))
}

func truncate(s string, max int) string {
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	runes := []rune(s)
	return string(runes[:max]) + "..."
}

func isTableSeparator(line string) bool {
	if !strings.HasPrefix(line, "|") {
		return false
	}
	cleaned := strings.ReplaceAll(line, "|", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, ":", "")
	cleaned = strings.TrimSpace(cleaned)
	return cleaned == ""
}

// The definition-pattern check skips notice blockquotes and bold list labels.
// Both are sanctioned forms. See Editing guardrails and Formatting in SKILL.md.
func isLabelShape(trimmed string) bool {
	return strings.HasPrefix(trimmed, ">") ||
		bulletItem.MatchString(trimmed) ||
		numberedItem.MatchString(trimmed)
}
