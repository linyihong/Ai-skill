package kge

import "strings"

// GovernanceModeConsistencyRule enforces governance_mode vs staged path
// sensitivity and LOCKDOWN approval trailers.
type GovernanceModeConsistencyRule struct{}

func (GovernanceModeConsistencyRule) ID() string { return "rule.governance_mode_consistency" }
func (GovernanceModeConsistencyRule) Kind() Kind { return KindValidation }

func (GovernanceModeConsistencyRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes, CapStagedPaths, CapCommitMsg}
}

func (GovernanceModeConsistencyRule) Validate(ctx Context) []Finding {
	modes := ctx.Modes
	if modes == nil {
		modes = map[string]string{}
	}
	gov := modes["governance_mode"]
	if gov == "" {
		return []Finding{{
			RuleID:   "rule.governance_mode_consistency",
			Severity: SeverityError,
			Code:     "governance_mode_missing",
			Message:  "governance_mode missing from Cognitive Mode block",
		}}
	}

	if gov == "LIGHT" || gov == "STANDARD" {
		for _, f := range ctx.StagedPaths {
			if RequiresDeepStrictCognitiveMode(f) {
				return []Finding{{
					RuleID:   "rule.governance_mode_consistency",
					Severity: SeverityError,
					Code:     "governance_mode_consistency",
					Path:     f,
					Message:  "governance_mode=" + gov + " forbidden when staged files include runtime/routing/workflow-contract/active-plan/governance-critical paths; use STRICT or LOCKDOWN. File: " + f,
				}}
			}
		}
	}

	if gov == "LOCKDOWN" {
		hasApproval := false
		for _, line := range strings.Split(ctx.CommitMsg, "\n") {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "[approved-by:") && strings.HasSuffix(trimmed, "]") {
				hasApproval = true
				break
			}
		}
		if !hasApproval {
			return []Finding{{
				RuleID:   "rule.governance_mode_consistency",
				Severity: SeverityError,
				Code:     "governance_mode_lockdown_approval",
				Message:  "governance_mode=LOCKDOWN requires an [approved-by: <name>] trailer line in the commit body",
			}}
		}
	}
	return nil
}
