package kge

import (
	"regexp"
	"strconv"
	"strings"
)

var tokenEstimateRE = regexp.MustCompile(`(?i)Token\s+Estimate:\s*(\d+)`)

// TokenBudgetRule enforces declared Token Estimate against mode-tuple budgets.
//
// Opt-out: standalone `[skip-token-budget]` trailer in CommitMsg.
// No estimate declared → no-op (opt-in for the turn).
type TokenBudgetRule struct{}

func (TokenBudgetRule) ID() string { return "rule.token_budget" }
func (TokenBudgetRule) Kind() Kind { return KindValidation }

func (TokenBudgetRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes, CapCommitMsg}
}

func (TokenBudgetRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-token-budget]" {
			return nil
		}
	}
	match := tokenEstimateRE.FindStringSubmatch(ctx.CommitMsg)
	if len(match) < 2 {
		return nil
	}
	estimate, err := strconv.Atoi(match[1])
	if err != nil {
		return nil
	}

	modes := ctx.Modes
	if modes == nil {
		modes = map[string]string{}
	}
	exec := modes["execution_mode"]
	cmode := modes["context_mode"]
	gov := modes["governance_mode"]
	mem := modes["memory_mode"]

	exactBudgets := map[string]int{
		"FAST|INDEX_ONLY|LIGHT|NONE":                    1000,
		"NORMAL|SUMMARY_FIRST|STANDARD|EPISODIC":        5000,
		"DEEP|SOURCE_BACKED|STRICT|DECISION_REPLAY":     20000,
		"FORENSIC|GRAPH_ASSISTED|STRICT|FAILURE_REPLAY": 50000,
	}
	execDefaults := map[string]int{
		"FAST":     1000,
		"NORMAL":   5000,
		"DEEP":     20000,
		"FORENSIC": 50000,
		"RECOVERY": 50000,
	}

	key := exec + "|" + cmode + "|" + gov + "|" + mem
	budget, ok := exactBudgets[key]
	if !ok {
		budget, ok = execDefaults[exec]
		if !ok {
			return nil
		}
	}
	if estimate <= budget {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.token_budget",
		Severity: SeverityError,
		Code:     "token_budget",
		Message: "token_budget: declared Token Estimate=" + strconv.Itoa(estimate) +
			" exceeds budget=" + strconv.Itoa(budget) +
			" for mode tuple (execution_mode=" + exec +
			", context_mode=" + cmode +
			", governance_mode=" + gov +
			", memory_mode=" + mem +
			"). Downgrade context_mode (GRAPH_ASSISTED → SOURCE_BACKED → CHECKLIST_FIRST → SUMMARY_FIRST → INDEX_ONLY) or split the work. Use [skip-token-budget] only if exceptional.",
	}}
}
