package kge

import (
	"path/filepath"
	"strings"
)

// PlanCheckboxSyncRule requires referenced active plans to be staged with a
// `[ ]` → `[x]` transition when the commit also stages code / governance work.
//
// Opt-out: standalone `[skip-plan-checkbox-sync]` trailer in CommitMsg.
//
// Requires CapStagedDiff via PathDiffs (per-plan unified diffs from the adapter).
type PlanCheckboxSyncRule struct{}

func (PlanCheckboxSyncRule) ID() string { return "rule.plan_checkbox_sync" }
func (PlanCheckboxSyncRule) Kind() Kind { return KindValidation }

func (PlanCheckboxSyncRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapStagedDiff}
}

func (PlanCheckboxSyncRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-plan-checkbox-sync]" {
			return nil
		}
	}
	planRefs := FindActivePlanRefs(ctx.CommitMsg)
	if len(planRefs) == 0 {
		return nil
	}
	hasCodeWork := false
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if strings.HasSuffix(s, ".go") ||
			strings.HasPrefix(s, "validation/scenarios/") ||
			strings.HasPrefix(s, "runtime/") ||
			strings.HasPrefix(s, "governance/") ||
			strings.HasPrefix(s, "enforcement/") {
			hasCodeWork = true
			break
		}
	}
	if !hasCodeWork {
		return nil
	}
	stagedSet := map[string]bool{}
	for _, s := range ctx.StagedPaths {
		stagedSet[s] = true
	}
	seen := map[string]bool{}
	var violations []string
	for _, ref := range planRefs {
		clean := strings.TrimRight(ref, "),]\"")
		if seen[clean] {
			continue
		}
		seen[clean] = true
		if !stagedSet[clean] {
			violations = append(violations, clean+" referenced but not staged")
			continue
		}
		diff := ""
		if ctx.PathDiffs != nil {
			diff = ctx.PathDiffs[clean]
		}
		if diff == "" {
			// Adapter failed to supply this path's diff — same as legacy git error: skip.
			continue
		}
		if !PlanDiffFlipsCheckbox(diff) {
			violations = append(violations, clean+" staged but no `[ ]` → `[x]` transition detected in staged diff")
		}
	}
	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_checkbox_sync",
		Severity: SeverityError,
		Code:     "plan_checkbox_sync",
		Message: "plan-checkbox-sync: commit references plans/active/* and stages code / scenario / governance work, but plan progress did not advance:\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  Flip the corresponding `- [ ]` task to `- [x]` in the same commit (cite this commit hash), or add a standalone `[skip-plan-checkbox-sync]` trailer line if this commit intentionally does not advance a plan phase (hotfix / refactor / cross-plan reference).",
	}}
}

// PlanDiffFlipsCheckbox returns true iff the unified diff text contains a
// line added (prefix `+`) whose content (after the marker and any
// indentation) starts with `- [x]`.
func PlanDiffFlipsCheckbox(diff string) bool {
	for _, line := range strings.Split(diff, "\n") {
		if len(line) == 0 || line[0] != '+' {
			continue
		}
		if strings.HasPrefix(line, "+++") {
			continue
		}
		body := strings.TrimLeft(line[1:], " \t")
		if strings.HasPrefix(body, "- [x]") || strings.HasPrefix(body, "- [X]") {
			return true
		}
	}
	return false
}
