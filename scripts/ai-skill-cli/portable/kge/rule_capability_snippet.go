package kge

import "strings"

// CapabilitySnippetRule requires a Capability summary section for high-risk modes.
type CapabilitySnippetRule struct{}

func (CapabilitySnippetRule) ID() string { return "rule.capability_snippet" }
func (CapabilitySnippetRule) Kind() Kind { return KindValidation }

func (CapabilitySnippetRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes, CapCommitMsg}
}

func (CapabilitySnippetRule) Validate(ctx Context) []Finding {
	modes := ctx.Modes
	if modes == nil {
		modes = map[string]string{}
	}
	exec := modes["execution_mode"]
	gov := modes["governance_mode"]
	highRisk := exec == "DEEP" || exec == "FORENSIC" || exec == "RECOVERY" || gov == "STRICT" || gov == "LOCKDOWN"
	if !highRisk {
		return nil
	}
	if strings.Contains(ctx.CommitMsg, "Capability summary:") {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.capability_snippet",
		Severity: SeverityError,
		Code:     "capability_snippet_missing",
		Message:  "capability snippet missing: high-risk Cognitive Contract modes require a Capability summary section",
	}}
}
