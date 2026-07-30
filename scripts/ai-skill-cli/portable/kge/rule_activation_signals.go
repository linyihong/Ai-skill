package kge

import (
	"sort"
	"strings"
)

// ParseActivationSignals extracts discovery signals from compact Sig= mode
// or activation_reason: list in the commit message body.
func ParseActivationSignals(text string, modes map[string]string) []string {
	if modes != nil {
		if sig := strings.TrimSpace(modes["activation_signal"]); sig != "" {
			return []string{sig}
		}
	}
	lines := strings.Split(text, "\n")
	inBlock := false
	var signals []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "activation_reason:" {
			inBlock = true
			continue
		}
		if !inBlock {
			continue
		}
		if trimmed == "" || strings.HasPrefix(trimmed, "##") || strings.HasPrefix(trimmed, "Capability summary:") {
			break
		}
		if strings.HasPrefix(trimmed, "- ") {
			sig := strings.TrimSpace(strings.TrimPrefix(trimmed, "- "))
			if idx := strings.Index(sig, "#"); idx >= 0 {
				sig = strings.TrimSpace(sig[:idx])
			}
			if sig != "" {
				signals = append(signals, sig)
			}
		}
	}
	return signals
}

// ActivationSignalsRule requires at least one known discovery signal.
type ActivationSignalsRule struct{}

func (ActivationSignalsRule) ID() string { return "rule.activation_signals" }
func (ActivationSignalsRule) Kind() Kind { return KindValidation }

func (ActivationSignalsRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes, CapCommitMsg, CapKnownSignals}
}

func (ActivationSignalsRule) Validate(ctx Context) []Finding {
	signals := ParseActivationSignals(ctx.CommitMsg, ctx.Modes)
	if len(signals) == 0 {
		return []Finding{{
			RuleID:   "rule.activation_signals",
			Severity: SeverityError,
			Code:     "activation_signals_missing",
			Message:  "activation_reason missing: Cognitive Contract v2 requires at least one discovery signal",
		}}
	}
	if len(ctx.KnownSignals) == 0 {
		return []Finding{{
			RuleID:   "rule.activation_signals",
			Severity: SeverityError,
			Code:     "activation_signals_vocab_unavailable",
			Message:  "activation_reason: known discovery signal list unavailable from runtime generated surface or YAML source",
		}}
	}
	var unknown []string
	for _, sig := range signals {
		if !ctx.KnownSignals[sig] {
			unknown = append(unknown, sig)
		}
	}
	if len(unknown) == 0 {
		return nil
	}
	sort.Strings(unknown)
	return []Finding{{
		RuleID:   "rule.activation_signals",
		Severity: SeverityError,
		Code:     "activation_signals_unknown",
		Message:  "activation_reason contains unknown discovery signal(s): " + strings.Join(unknown, ", "),
	}}
}
