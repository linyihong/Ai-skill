package kge

import (
	"fmt"
	"path/filepath"
	"strings"
)

// DocumentSizingRule emits advisory findings when Markdown files grow large.
// Thresholds align with governance/document-sizing.md (150 soft / 300 hard warn).
// Never SeverityError — split judgment stays human (registry behavioral_only).
type DocumentSizingRule struct {
	SoftLines int // default 150
	HardLines int // default 300
}

func (DocumentSizingRule) ID() string { return "rule.document_sizing" }
func (DocumentSizingRule) Kind() Kind { return KindAdvisory }

func (DocumentSizingRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapStagedContent}
}

func (r DocumentSizingRule) soft() int {
	if r.SoftLines <= 0 {
		return 150
	}
	return r.SoftLines
}

func (r DocumentSizingRule) hard() int {
	if r.HardLines <= 0 {
		return 300
	}
	return r.HardLines
}

func (r DocumentSizingRule) Validate(ctx Context) []Finding {
	var out []Finding
	soft, hard := r.soft(), r.hard()
	for path, body := range ctx.FileContents {
		if !isMarkdownPath(path) {
			continue
		}
		n := countLines(body)
		switch {
		case n >= hard:
			out = append(out, Finding{
				RuleID:   r.ID(),
				Severity: SeverityWarning,
				Code:     "document_sizing_hard",
				Path:     path,
				Message:  fmt.Sprintf("document sizing: %s is %d lines (≥ %d caution); consider splitting into a folder with a single-purpose index (governance/document-sizing.md)", path, n, hard),
			})
		case n >= soft:
			out = append(out, Finding{
				RuleID:   r.ID(),
				Severity: SeverityWarning,
				Code:     "document_sizing_soft",
				Path:     path,
				Message:  fmt.Sprintf("document sizing: %s is %d lines (≥ %d); check whether topics should stay single-purpose or split", path, n, soft),
			})
		}
	}
	return out
}

func isMarkdownPath(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".md" || ext == ".markdown"
}

func countLines(s string) int {
	if s == "" {
		return 0
	}
	n := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			n++
		}
	}
	if strings.HasSuffix(s, "\n") && n > 0 {
		n--
	}
	return n
}
