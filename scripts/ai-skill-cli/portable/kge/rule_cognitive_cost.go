package kge

// CognitiveCostRule is Phase 0 Mini Spike Rule A — no projection, needs modes.
type CognitiveCostRule struct{}

func (CognitiveCostRule) ID() string   { return "rule.cognitive_cost" }
func (CognitiveCostRule) Kind() Kind   { return KindValidation }

func (CognitiveCostRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes}
}

func (CognitiveCostRule) Validate(ctx Context) []Finding {
	modes := ctx.Modes
	if modes == nil {
		return []Finding{{
			RuleID:   "rule.cognitive_cost",
			Severity: SeverityError,
			Code:     "cognitive_cost_missing",
			Message:  "cognitive_cost missing from Cognitive Contract v2 block",
		}}
	}
	declared := modes["cognitive_cost"]
	if declared == "" {
		return []Finding{{
			RuleID:   "rule.cognitive_cost",
			Severity: SeverityError,
			Code:     "cognitive_cost_missing",
			Message:  "cognitive_cost missing from Cognitive Contract v2 block",
		}}
	}
	derived := deriveCognitiveCost(modes["execution_mode"], modes["context_mode"])
	if derived == "" {
		return []Finding{{
			RuleID:   "rule.cognitive_cost",
			Severity: SeverityError,
			Code:     "cognitive_cost_underivable",
			Message:  "cognitive_cost: cannot derive cost for execution_mode=" + modes["execution_mode"] + " context_mode=" + modes["context_mode"],
		}}
	}
	if declared != derived {
		return []Finding{{
			RuleID:   "rule.cognitive_cost",
			Severity: SeverityError,
			Code:     "cognitive_cost_mismatch",
			Message:  "cognitive_cost mismatch: declared=" + declared + " derived=" + derived + " for execution_mode=" + modes["execution_mode"] + " context_mode=" + modes["context_mode"],
		}}
	}
	return nil
}

// deriveCognitiveCost mirrors app.deriveCognitiveCost (kept local so kge
// does not import the hook monolith). Keep in sync when changing cost matrix.
func deriveCognitiveCost(executionMode, contextMode string) string {
	switch executionMode {
	case "FAST":
		if contextMode == "INDEX_ONLY" {
			return "LOW"
		}
		return "MEDIUM"
	case "NORMAL":
		if contextMode == "INDEX_ONLY" || contextMode == "SUMMARY_FIRST" {
			return "LOW"
		}
		if contextMode == "CHECKLIST_FIRST" || contextMode == "SOURCE_BACKED" || contextMode == "GRAPH_ASSISTED" {
			return "MEDIUM"
		}
	case "DEEP":
		return "HIGH"
	case "FORENSIC", "RECOVERY":
		return "VERY_HIGH"
	}
	return ""
}
