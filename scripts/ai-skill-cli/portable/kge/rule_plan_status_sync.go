package kge

import (
	"regexp"
	"strings"
)

var (
	planPathRE     = regexp.MustCompile(`plans/active/[^\s)"\]]+\.md`)
	phaseMentionRE = regexp.MustCompile(`(?i)Phase\s+\d+(?:\.\d+)?(?:[\.-][A-Za-z]+)?`)
)

var completionPhrases = []string{
	"complete", "completed", "completes", "done", "finish", "finished",
	"完成", "結案", "結束", "✅",
}

// PlanStatusSyncRule requires referenced active plan files to be staged when
// the commit body claims phase/milestone completion.
//
// Opt-out: standalone `[skip-plan-status-sync]` trailer in CommitMsg.
type PlanStatusSyncRule struct{}

func (PlanStatusSyncRule) ID() string { return "rule.plan_status_sync" }
func (PlanStatusSyncRule) Kind() Kind { return KindValidation }

func (PlanStatusSyncRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths}
}

func (PlanStatusSyncRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-plan-status-sync]" {
			return nil
		}
	}

	text := ctx.CommitMsg
	hasCompletion := false
	lowered := strings.ToLower(text)
	for _, phrase := range completionPhrases {
		if strings.Contains(lowered, strings.ToLower(phrase)) {
			hasCompletion = true
			break
		}
	}
	if !hasCompletion {
		return nil
	}
	if !phaseMentionRE.MatchString(text) {
		return nil
	}
	planRefs := planPathRE.FindAllString(text, -1)
	if len(planRefs) == 0 {
		return nil
	}

	stagedSet := make(map[string]bool, len(ctx.StagedPaths))
	for _, s := range ctx.StagedPaths {
		stagedSet[s] = true
	}
	var missing []string
	seen := map[string]bool{}
	for _, ref := range planRefs {
		clean := strings.TrimRight(ref, "),]\"")
		if seen[clean] {
			continue
		}
		seen[clean] = true
		if !stagedSet[clean] {
			missing = append(missing, clean)
		}
	}
	if len(missing) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_status_sync",
		Severity: SeverityError,
		Code:     "plan_status_sync",
		Message: "plan-status-sync: commit body claims phase completion and references " +
			strings.Join(missing, ", ") +
			" but the plan file is not in the staged set. Update the plan's Phase section in the same commit (runtime/plan-status-sync-enforcement.yaml). Use a [skip-plan-status-sync] trailer for retrospective references.",
	}}
}
