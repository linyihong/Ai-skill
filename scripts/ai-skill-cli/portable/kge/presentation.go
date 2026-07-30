package kge

import (
	"fmt"
	"strings"
)

// DefaultRules is the first-party pack for Ai-skill / copyable demos.
func DefaultRules() []Rule {
	return []Rule{
		CognitiveCostRule{},
		CLIDocSyncRule{},
		DocumentSizingRule{},
	}
}

// Partition splits findings by severity class for Adapter Presentation Policy.
func Partition(findings []Finding) (errors, advisories, other []Finding) {
	for _, f := range findings {
		switch f.Severity {
		case SeverityError:
			errors = append(errors, f)
		case SeverityWarning:
			advisories = append(advisories, f)
		default:
			other = append(other, f)
		}
	}
	return errors, advisories, other
}

// FormatCommitSummary is D9 commit-msg presentation: ≤3 lines, count + pointer.
func FormatCommitSummary(advisoryCount int, validateHint string) string {
	if advisoryCount <= 0 {
		return ""
	}
	if validateHint == "" {
		validateHint = "ai-skill kge validate --advisory"
	}
	return fmt.Sprintf("KGE: %d advisory finding(s).\nRun: %s", advisoryCount, validateHint)
}

// FormatCheckReport is D9 kge check / pre-push presentation.
// maxAdvisory caps body lines (default 3).
func FormatCheckReport(findings []Finding, maxAdvisory int) string {
	if maxAdvisory <= 0 {
		maxAdvisory = 3
	}
	errs, adv, _ := Partition(findings)
	var b strings.Builder
	if len(errs) > 0 {
		b.WriteString("KGE validation failed:\n")
		for _, f := range errs {
			b.WriteString("  - ")
			b.WriteString(f.Message)
			b.WriteByte('\n')
		}
	} else {
		b.WriteString("Ready to push (validation ok).\n")
	}
	if len(adv) == 0 {
		return strings.TrimRight(b.String(), "\n")
	}
	b.WriteString(fmt.Sprintf("%d recommendation(s):\n", len(adv)))
	limit := maxAdvisory
	if limit > len(adv) {
		limit = len(adv)
	}
	for i := 0; i < limit; i++ {
		b.WriteString("  - ")
		b.WriteString(shortLabel(adv[i]))
		b.WriteByte('\n')
	}
	if len(adv) > limit {
		b.WriteString(fmt.Sprintf("  … and %d more. Run: ai-skill kge validate --advisory\n", len(adv)-limit))
	} else {
		b.WriteString("Run: ai-skill kge validate --advisory\n")
	}
	return strings.TrimRight(b.String(), "\n")
}

// FormatAdvisoryFull lists all advisory (and info) findings.
func FormatAdvisoryFull(findings []Finding) string {
	_, adv, other := Partition(findings)
	all := append(append([]Finding{}, adv...), other...)
	if len(all) == 0 {
		return "No advisory findings."
	}
	var b strings.Builder
	b.WriteString(fmt.Sprintf("Advisory (%d):\n", len(all)))
	for _, f := range all {
		b.WriteString("  [")
		b.WriteString(f.RuleID)
		b.WriteString("] ")
		b.WriteString(f.Message)
		b.WriteByte('\n')
	}
	return strings.TrimRight(b.String(), "\n")
}

func shortLabel(f Finding) string {
	if f.RuleID != "" {
		msg := f.Message
		if len(msg) > 80 {
			msg = msg[:77] + "..."
		}
		return f.RuleID + ": " + msg
	}
	return f.Message
}
