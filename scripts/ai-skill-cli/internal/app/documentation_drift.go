package app

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// reviewArchitectureNavigationCanon lists navigation-layer sources that must
// describe review only as capability invoke — never as a workflow phase/slice.
// Phase 2.2 Documentation Drift Lock (ADR-013).
var reviewArchitectureNavigationCanon = []string{
	"workflow/software-delivery/execution-flow.md",
	"workflow/software-delivery/README.md",
	"governance/cognitive-slice-taxonomy.md",
	"knowledge/runtime/routing-registry.yaml",
}

type reviewArchitectureDriftPattern struct {
	id      string
	pattern *regexp.Regexp
}

var reviewArchitectureDriftPatterns = []reviewArchitectureDriftPattern{
	{
		id:      "review_phase",
		pattern: regexp.MustCompile(`(?i)\breview\s+phase\b`),
	},
	{
		id:      "review_workflow",
		pattern: regexp.MustCompile(`(?i)\breview\s+workflow\b`),
	},
	{
		id:      "review_slice",
		pattern: regexp.MustCompile(`(?i)\breview\s+slice\b`),
	},
	{
		id:      "sd_review_slice",
		pattern: regexp.MustCompile(`(?i)\bsd-review\b`),
	},
}

var reviewArchitectureDriftNegation = regexp.MustCompile(`(?i)(不新增|不做|reject|forbidden|forbid|非|not a|not.*sd-review|No ` + "`" + `sd-review)`)

func scanReviewArchitectureNavigationDrift(content, relPath string) []string {
	var violations []string
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") && strings.Contains(trimmed, "Forbidden") {
			continue
		}
		if reviewArchitectureDriftNegation.MatchString(line) {
			continue
		}
		for _, p := range reviewArchitectureDriftPatterns {
			if p.pattern.MatchString(line) {
				violations = append(violations, fmt.Sprintf(
					"%s:%d: forbidden %q — review is capability invoke, not workflow phase/slice (ADR-013)",
					relPath, i+1, p.id))
			}
		}
	}
	return violations
}

func validateReviewArchitectureNavigationFiles(repo string, relPaths []string) []string {
	var violations []string
	for _, rel := range relPaths {
		content, err := os.ReadFile(filepath.Join(repo, rel))
		if err != nil {
			violations = append(violations, fmt.Sprintf("%s: read failed: %v", rel, err))
			continue
		}
		violations = append(violations, scanReviewArchitectureNavigationDrift(string(content), rel)...)
	}
	return violations
}

func nativeReviewArchitectureDocDriftValidation(repo string) Check {
	violations := validateReviewArchitectureNavigationFiles(repo, reviewArchitectureNavigationCanon)
	if len(violations) > 0 {
		return Check{
			Name:    "review_architecture_doc_drift",
			Status:  "failed",
			Message: strings.Join(violations, "; "),
		}
	}
	return Check{
		Name:   "review_architecture_doc_drift",
		Status: "ok",
		Message: fmt.Sprintf(
			"navigation canon ok: %d files scanned for review phase/workflow/slice drift",
			len(reviewArchitectureNavigationCanon),
		),
	}
}

func validateReviewArchitectureNavigationDriftStaged(repo string, staged []string) string {
	stagedSet := map[string]bool{}
	for _, s := range staged {
		stagedSet[s] = true
	}
	var toScan []string
	for _, canon := range reviewArchitectureNavigationCanon {
		if stagedSet[canon] {
			toScan = append(toScan, canon)
		}
	}
	if len(toScan) == 0 {
		return ""
	}
	violations := validateReviewArchitectureNavigationFiles(repo, toScan)
	if len(violations) == 0 {
		return ""
	}
	return "review-architecture-doc-drift: navigation layer must describe review as capability invoke only:\n  - " +
		strings.Join(violations, "\n  - ") +
		"\n  See governance/cognitive-stance.md and ADR-013. Opt-out: [skip-review-architecture-doc-drift] on its own line."
}

func reviewArchitectureDocDriftOptOut(text string) bool {
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == "[skip-review-architecture-doc-drift]" {
			return true
		}
	}
	return false
}
