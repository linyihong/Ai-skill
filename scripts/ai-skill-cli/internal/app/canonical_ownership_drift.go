package app

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// reviewArchitectureDebtPayoffCanon — Phase 2.3 active canonical sources scanned
// for Ownership / Path / Semantic / Artifact drift (ADR-013 debt payoff).
var reviewArchitectureDebtPayoffCanon = []string{
	"workflow/software-delivery/validation.md",
	"workflow/software-delivery/intake.md",
	"workflow/software-delivery/perf-risk-gate.md",
	"workflow/software-delivery/templates/README.md",
	"workflow/cross-cutting/review/self-review.md",
}

type canonicalOwnershipRule struct {
	debtClass string
	id        string
	pattern   *regexp.Regexp
	owner     string
}

var canonicalOwnershipRules = []canonicalOwnershipRule{
	{
		debtClass: "A",
		id:        "validation_owns_review_report",
		pattern:   regexp.MustCompile(`(?i)(Validate\s+完成後|validation phase output|Validate phase).{0,80}review-report-template`),
		owner:     "code-review capability (cross-cutting/review/)",
	},
	{
		debtClass: "A",
		id:        "validation_output_template_review_report",
		pattern:   regexp.MustCompile(`(?i)輸出模板.*Validate.*review-report-template`),
		owner:     "code-review capability (cross-cutting/review/)",
	},
	{
		debtClass: "A",
		id:        "readme_defines_stance_enum",
		pattern:   regexp.MustCompile(`(?i)(stance_enum:|reserved_policy:|conservative enum|Do not reserve placeholder)`),
		owner:     "governance/cognitive-stance.md",
	},
}

var pathDriftReviewChecklistPattern = regexp.MustCompile(`workflow/software-delivery/review-checklist\.md`)

var pathDriftReviewChecklistAllowed = regexp.MustCompile(`(?i)(stub|redirect|遷移|legacy|舊路徑|inbound|stable inbound|已遷移|Canonical source)`)

var semanticDriftReviewFlowPattern = regexp.MustCompile(`(?i)\breview\s+flow\b`)

func scanCanonicalOwnershipDrift(content, relPath string) []string {
	var violations []string
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if reviewArchitectureDriftNegation.MatchString(line) {
			continue
		}
		for _, rule := range canonicalOwnershipRules {
			if rule.id == "readme_defines_stance_enum" && relPath != "workflow/software-delivery/README.md" {
				continue
			}
			if (rule.id == "validation_owns_review_report" || rule.id == "validation_output_template_review_report") &&
				relPath != "workflow/software-delivery/validation.md" {
				continue
			}
			if rule.pattern.MatchString(line) {
				violations = append(violations, fmt.Sprintf(
					"%s:%d: ownership drift (class %s/%s) — canonical owner is %s (ADR-013)",
					relPath, i+1, rule.debtClass, rule.id, rule.owner))
			}
		}
		if semanticDriftReviewFlowPattern.MatchString(line) && !reviewArchitectureDriftNegation.MatchString(line) {
			violations = append(violations, fmt.Sprintf(
				"%s:%d: semantic drift (class C/review_flow) — use capability invoke, not review flow (ADR-013)",
				relPath, i+1))
		}
		if relPath != "workflow/software-delivery/review-checklist.md" &&
			pathDriftReviewChecklistPattern.MatchString(line) &&
			!pathDriftReviewChecklistAllowed.MatchString(line) {
			violations = append(violations, fmt.Sprintf(
				"%s:%d: path drift (class B/review_checklist_stub) — canonical path is workflow/cross-cutting/review/checklist.md",
				relPath, i+1))
		}
	}
	return violations
}

func reviewArchitectureScenarioPaths(repo string) []string {
	root := filepath.Join(repo, "validation", "scenarios", "software-delivery")
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil
	}
	var paths []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".yaml") {
			continue
		}
		paths = append(paths, filepath.Join("validation", "scenarios", "software-delivery", e.Name()))
	}
	return paths
}

func validateCanonicalOwnershipFiles(repo string, relPaths []string) []string {
	var violations []string
	for _, rel := range relPaths {
		content, err := os.ReadFile(filepath.Join(repo, rel))
		if err != nil {
			violations = append(violations, fmt.Sprintf("%s: read failed: %v", rel, err))
			continue
		}
		violations = append(violations, scanCanonicalOwnershipDrift(string(content), rel)...)
	}
	return violations
}

func nativeCanonicalOwnershipDriftValidation(repo string) Check {
	paths := append([]string{}, reviewArchitectureDebtPayoffCanon...)
	paths = append(paths, reviewArchitectureNavigationCanon...)
	paths = append(paths, reviewArchitectureScenarioPaths(repo)...)
	seen := map[string]bool{}
	var unique []string
	for _, p := range paths {
		if seen[p] {
			continue
		}
		seen[p] = true
		unique = append(unique, p)
	}
	violations := validateCanonicalOwnershipFiles(repo, unique)
	if len(violations) > 0 {
		return Check{
			Name:    "canonical_ownership_drift",
			Status:  "failed",
			Message: strings.Join(violations, "; "),
		}
	}
	return Check{
		Name:   "canonical_ownership_drift",
		Status: "ok",
		Message: fmt.Sprintf(
			"debt payoff canon ok: %d files scanned (ownership/path/semantic/artifact drift)",
			len(unique),
		),
	}
}

func validateCanonicalOwnershipDriftStaged(repo string, staged []string) string {
	stagedSet := map[string]bool{}
	for _, s := range staged {
		stagedSet[s] = true
	}
	var toScan []string
	for _, p := range reviewArchitectureDebtPayoffCanon {
		if stagedSet[p] {
			toScan = append(toScan, p)
		}
	}
	for _, p := range reviewArchitectureNavigationCanon {
		if stagedSet[p] {
			toScan = append(toScan, p)
		}
	}
	for _, s := range staged {
		if strings.HasPrefix(s, "validation/scenarios/software-delivery/") && strings.HasSuffix(s, ".yaml") {
			toScan = append(toScan, s)
		}
	}
	if len(toScan) == 0 {
		return ""
	}
	violations := validateCanonicalOwnershipFiles(repo, toScan)
	if len(violations) == 0 {
		return ""
	}
	return "canonical-ownership-drift: architectural debt detected in staged canonical sources:\n  - " +
		strings.Join(violations, "\n  - ") +
		"\n  See ADR-013 debt classes A–D. Opt-out: [skip-canonical-ownership-drift] on its own line."
}

func canonicalOwnershipDriftOptOut(text string) bool {
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == "[skip-canonical-ownership-drift]" {
			return true
		}
	}
	return false
}
