package kge

// ExecutionModeFloorsRule enforces execution_mode floor requirements
// (cognitive-modes-phase-integration.yaml).
type ExecutionModeFloorsRule struct{}

func (ExecutionModeFloorsRule) ID() string { return "rule.execution_mode_floors" }
func (ExecutionModeFloorsRule) Kind() Kind { return KindValidation }

func (ExecutionModeFloorsRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes, CapStagedPaths}
}

func (ExecutionModeFloorsRule) Validate(ctx Context) []Finding {
	modes := ctx.Modes
	if modes == nil {
		modes = map[string]string{}
	}
	exec := modes["execution_mode"]
	gov := modes["governance_mode"]
	cmode := modes["context_mode"]
	mem := modes["memory_mode"]

	if exec == "FAST" {
		for _, f := range ctx.StagedPaths {
			if RequiresDeepStrictCognitiveMode(f) {
				return []Finding{{
					RuleID:   "rule.execution_mode_floors",
					Severity: SeverityError,
					Code:     "execution_mode_floors",
					Path:     f,
					Message:  "execution_mode=FAST forbidden when staged files touch runtime/routing/workflow-contract/active-plan/governance-critical paths (auto-escalation rule per cognitive-modes-phase-integration.yaml). File: " + f,
				}}
			}
		}
	}

	if exec == "NORMAL" {
		for _, f := range ctx.StagedPaths {
			if RequiresDeepStrictCognitiveMode(f) {
				return []Finding{{
					RuleID:   "rule.execution_mode_floors",
					Severity: SeverityError,
					Code:     "execution_mode_floors",
					Path:     f,
					Message:  "execution_mode=NORMAL insufficient when staged files touch runtime/routing/workflow-contract/active-plan/governance-critical paths; use DEEP or higher. File: " + f,
				}}
			}
		}
	}

	if exec == "DEEP" || exec == "FORENSIC" || exec == "RECOVERY" {
		if gov != "STRICT" && gov != "LOCKDOWN" {
			return []Finding{{
				RuleID:   "rule.execution_mode_floors",
				Severity: SeverityError,
				Code:     "execution_mode_floors",
				Message:  "execution_mode=" + exec + " requires governance_mode ≥ STRICT (declared: " + gov + ")",
			}}
		}
	}

	if exec == "DEEP" && cmode != "SOURCE_BACKED" && cmode != "GRAPH_ASSISTED" {
		return []Finding{{
			RuleID:   "rule.execution_mode_floors",
			Severity: SeverityError,
			Code:     "execution_mode_floors",
			Message:  "execution_mode=DEEP requires context_mode ≥ SOURCE_BACKED (declared: " + cmode + ")",
		}}
	}
	if exec == "FORENSIC" && cmode != "GRAPH_ASSISTED" {
		return []Finding{{
			RuleID:   "rule.execution_mode_floors",
			Severity: SeverityError,
			Code:     "execution_mode_floors",
			Message:  "execution_mode=FORENSIC requires context_mode=GRAPH_ASSISTED (declared: " + cmode + ")",
		}}
	}
	if exec == "RECOVERY" {
		if cmode != "CHECKLIST_FIRST" {
			return []Finding{{
				RuleID:   "rule.execution_mode_floors",
				Severity: SeverityError,
				Code:     "execution_mode_floors",
				Message:  "execution_mode=RECOVERY requires context_mode=CHECKLIST_FIRST (declared: " + cmode + ")",
			}}
		}
		if mem != "FAILURE_REPLAY" {
			return []Finding{{
				RuleID:   "rule.execution_mode_floors",
				Severity: SeverityError,
				Code:     "execution_mode_floors",
				Message:  "execution_mode=RECOVERY requires memory_mode=FAILURE_REPLAY (declared: " + mem + ")",
			}}
		}
	}
	return nil
}
