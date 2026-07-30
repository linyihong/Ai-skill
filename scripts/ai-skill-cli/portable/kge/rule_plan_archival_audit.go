package kge

import (
	"fmt"
	"path/filepath"
	"strings"
)

// PlanArchivalAuditRule blocks archiving plans that still contain unchecked
// `- [ ]` items unless the commit body justifies deferral.
//
// Opt-out: standalone `[skip-plan-archival-audit]` trailer in CommitMsg.
type PlanArchivalAuditRule struct{}

func (PlanArchivalAuditRule) ID() string { return "rule.plan_archival_audit" }
func (PlanArchivalAuditRule) Kind() Kind { return KindValidation }

func (PlanArchivalAuditRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapStagedContent}
}

func (PlanArchivalAuditRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-plan-archival-audit]" {
			return nil
		}
	}
	archived := FindArchivedPlanPaths(ctx.StagedPaths)
	if len(archived) == 0 {
		return nil
	}
	if BodyJustifiesUnchecked(ctx.CommitMsg) {
		return nil
	}
	var violations []string
	for _, rel := range archived {
		content, ok := ctx.FileContents[rel]
		if !ok {
			content, ok = ctx.FileContents[filepath.ToSlash(rel)]
		}
		if !ok {
			continue
		}
		n := CountUncheckedCheckboxes(content)
		if n == 0 {
			continue
		}
		violations = append(violations,
			fmt.Sprintf("%s has %d unchecked item(s) with no body justification", rel, n),
		)
	}
	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_archival_audit",
		Severity: SeverityError,
		Code:     "plan_archival_audit",
		Message: "plan-archival-audit: archiving plan(s) with unresolved `- [ ]` items:\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  Either justify in commit body (deferred / non-goal / scope reduced / handover / 延後 / 拆分)" +
			" or add a standalone `[skip-plan-archival-audit]` trailer for emergency archives.",
	}}
}

func FindArchivedPlanPaths(staged []string) []string {
	var result []string
	for _, s := range staged {
		s = filepath.ToSlash(s)
		if strings.HasPrefix(s, "plans/archived/") && strings.HasSuffix(s, ".md") {
			result = append(result, s)
		}
	}
	return result
}

// BodyJustifiesUnchecked reports whether the commit body justifies leaving
// unchecked items in an archived plan.
func BodyJustifiesUnchecked(body string) bool {
	keywords := []string{
		"deferred", "non-goal", "scope reduced", "handover", "延後", "拆分",
	}
	lower := strings.ToLower(body)
	for _, kw := range keywords {
		if strings.Contains(lower, strings.ToLower(kw)) {
			return true
		}
	}
	return false
}

// CountUncheckedCheckboxes counts Markdown task-list lines starting with `- [ ]`.
func CountUncheckedCheckboxes(content string) int {
	n := 0
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- [ ]") {
			n++
		}
	}
	return n
}
