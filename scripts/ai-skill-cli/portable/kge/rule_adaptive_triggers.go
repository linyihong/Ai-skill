package kge

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	contradictionKeywordsRE = regexp.MustCompile(`(?i)contradict\w*|conflict\w*|mismatch\w*|discrepan\w*|衝突|矛盾|不一致`)
	failurePatternRefRE     = regexp.MustCompile(`enforcement/failure-patterns/[\w-]+\.md`)
	revertHotfixRE          = regexp.MustCompile(`(?i)\b(revert|hotfix|retry)\b`)
	sourceClassRE           = regexp.MustCompile(`(plans/|constitution/|decisions/)[^\s)"\]]+`)
)

// AdaptiveTriggersRule implements runtime/cognitive-modes-adaptive.yaml
// commit-msg-detectable adaptive triggers.
//
// Opt-out: standalone `[skip-adaptive]` trailer in CommitMsg.
type AdaptiveTriggersRule struct{}

func (AdaptiveTriggersRule) ID() string { return "rule.adaptive_triggers" }
func (AdaptiveTriggersRule) Kind() Kind { return KindValidation }

func (AdaptiveTriggersRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes, CapCommitMsg}
}

func (AdaptiveTriggersRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-adaptive]" {
			return nil
		}
	}

	modes := ctx.Modes
	if modes == nil {
		modes = map[string]string{}
	}
	exec := modes["execution_mode"]
	cmode := modes["context_mode"]
	gov := modes["governance_mode"]
	mem := modes["memory_mode"]
	text := ctx.CommitMsg

	var out []Finding

	if contradictionKeywordsRE.MatchString(text) {
		refs := sourceClassRE.FindAllString(text, -1)
		distinct := map[string]bool{}
		for _, r := range refs {
			distinct[r] = true
		}
		if len(distinct) >= 2 {
			govOK := gov == "STRICT" || gov == "LOCKDOWN"
			ctxOK := cmode == "SOURCE_BACKED" || cmode == "GRAPH_ASSISTED"
			if !govOK || !ctxOK {
				out = append(out, Finding{
					RuleID:   "rule.adaptive_triggers",
					Severity: SeverityError,
					Code:     "adaptive_contradiction_risk",
					Message: "adaptive: contradiction risk detected (cross-source keywords + ≥2 distinct source refs) but governance_mode=" +
						gov + " / context_mode=" + cmode +
						" below required floor (governance_mode≥STRICT, context_mode in {SOURCE_BACKED, GRAPH_ASSISTED}). Upgrade modes per runtime/cognitive-modes-adaptive.yaml §contradiction_risk.",
				})
			}
		}
	}

	failureRefs := len(failurePatternRefRE.FindAllString(text, -1))
	revertHits := len(revertHotfixRE.FindAllString(text, -1))
	if failureRefs >= 2 || revertHits >= 2 {
		if exec != "RECOVERY" || mem != "FAILURE_REPLAY" {
			out = append(out, Finding{
				RuleID:   "rule.adaptive_triggers",
				Severity: SeverityError,
				Code:     "adaptive_repeated_failure",
				Message: "adaptive: repeated failure signal (failure-pattern refs=" + strconv.Itoa(failureRefs) +
					", revert/hotfix/retry hits=" + strconv.Itoa(revertHits) +
					") requires execution_mode=RECOVERY and memory_mode=FAILURE_REPLAY (declared: execution_mode=" +
					exec + ", memory_mode=" + mem +
					"). Adjust mode tuple per runtime/cognitive-modes-adaptive.yaml §repeated_failure.",
			})
		}
	}

	match := tokenEstimateRE.FindStringSubmatch(text)
	if len(match) >= 2 {
		estimate, err := strconv.Atoi(match[1])
		if err == nil {
			exactBudgets := map[string]int{
				"FAST|INDEX_ONLY|LIGHT|NONE":                    1000,
				"NORMAL|SUMMARY_FIRST|STANDARD|EPISODIC":        5000,
				"DEEP|SOURCE_BACKED|STRICT|DECISION_REPLAY":     20000,
				"FORENSIC|GRAPH_ASSISTED|STRICT|FAILURE_REPLAY": 50000,
			}
			execDefaults := map[string]int{
				"FAST": 1000, "NORMAL": 5000, "DEEP": 20000, "FORENSIC": 50000, "RECOVERY": 50000,
			}
			key := exec + "|" + cmode + "|" + gov + "|" + mem
			budget, ok := exactBudgets[key]
			if !ok {
				budget, ok = execDefaults[exec]
			}
			if ok && estimate >= (budget*80/100) && estimate <= budget {
				out = append(out, Finding{
					RuleID:   "rule.adaptive_triggers",
					Severity: SeverityError,
					Code:     "adaptive_budget_near_ceiling",
					Message: "adaptive[warning]: Token Estimate=" + strconv.Itoa(estimate) +
						" is ≥80% of budget=" + strconv.Itoa(budget) +
						"; consider downgrading context_mode one step along the downgrade_path (GRAPH_ASSISTED → SOURCE_BACKED → CHECKLIST_FIRST → SUMMARY_FIRST → INDEX_ONLY) OR split the work. Suppress this notice with [skip-adaptive].",
				})
			}
		}
	}
	return out
}
