package kge

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	registryTransitionOptOut    = "[skip-registry-transition]"
	registryStatusChangeTrailer = "[registry-status-change]"
	enforcementRegistryPath     = "enforcement/enforcement-registry.yaml"
)

var rationaleLineRE = regexp.MustCompile(`(?im)^\s*rationale\s*:\s*\S`)

var demotionTable = map[string]map[string]bool{
	"mechanical": {
		"behavioral_only": true, "not_mechanizable": true,
		"research_required": true, "pending_implementation": true,
	},
	"pending_implementation": {
		"behavioral_only": true, "research_required": true,
	},
	"behavioral_only": {
		"not_mechanizable": true,
	},
}

// EnforcementRegistryTransitionRule enforces Phase 4.5 R1/R2/R3 when the
// enforcement registry coverage changes.
//
// Opt-out: commit body contains `[skip-registry-transition]`.
type EnforcementRegistryTransitionRule struct{}

func (EnforcementRegistryTransitionRule) ID() string {
	return "rule.enforcement_registry_transition"
}
func (EnforcementRegistryTransitionRule) Kind() Kind { return KindValidation }

func (EnforcementRegistryTransitionRule) RequiredCapabilities() []CapabilityID {
	// CapStagedPaths is optional: commit-msg adapters supply it to gate on
	// staged registry; CLI transition-check omits it and always evaluates.
	return []CapabilityID{CapCommitMsg, CapRegistrySnapshots, CapRepoFS, CapSymbolIndex}
}

func (EnforcementRegistryTransitionRule) Validate(ctx Context) []Finding {
	if strings.Contains(ctx.CommitMsg, registryTransitionOptOut) {
		return nil
	}
	if ctx.Has(CapStagedPaths) {
		staged := false
		for _, p := range ctx.StagedPaths {
			if filepath.ToSlash(strings.TrimSpace(p)) == enforcementRegistryPath {
				staged = true
				break
			}
		}
		if !staged {
			return nil
		}
	}
	if ctx.RegistryOld == nil || ctx.RegistryNew == nil {
		return nil
	}

	oldByID := map[string]string{}
	for _, rc := range ctx.RegistryOld.RuleClasses {
		oldByID[rc.ID] = rc.Coverage
	}
	type transition struct {
		id   string
		from string
		to   string
		rc   RegistryClassMeta
	}
	var transitions []transition
	for _, rc := range ctx.RegistryNew.RuleClasses {
		from, existed := oldByID[rc.ID]
		if !existed {
			from = "(new)"
		}
		if from == rc.Coverage {
			continue
		}
		transitions = append(transitions, transition{id: rc.ID, from: from, to: rc.Coverage, rc: rc})
	}

	var out []Finding
	emit := func(code, classID, from, to, detail string) {
		parts := []string{"[" + code + "]"}
		if classID != "" {
			parts = append(parts, "rule_class="+classID)
		}
		if from != "" || to != "" {
			parts = append(parts, fmt.Sprintf("transition=%s→%s", from, to))
		}
		parts = append(parts, detail)
		out = append(out, Finding{
			RuleID:   "rule.enforcement_registry_transition",
			Severity: SeverityError,
			Code:     code,
			Message:  strings.Join(parts, " "),
			Path:     enforcementRegistryPath,
		})
	}

	if len(transitions) > 0 {
		if !hasOptOut(ctx.CommitMsg, registryStatusChangeTrailer) {
			emit("R1_missing_trailer", "", "", "", "commit body must include [registry-status-change] trailer when staged diff changes rule_class coverage")
		}
		if !rationaleLineRE.MatchString(ctx.CommitMsg) {
			emit("R1_missing_rationale", "", "", "", "commit body must include a `rationale: <text>` line explaining the status change")
		}
	}
	for _, t := range transitions {
		if isDemotionCoverage(t.from, t.to) {
			for _, m := range checkR2(ctx, t.id, t.from, t.to, t.rc) {
				emit(m.code, t.id, t.from, t.to, m.detail)
			}
		}
		if isPromotionToMechanicalCoverage(t.from, t.to) {
			for _, m := range checkR3(ctx, t.id, t.from, t.to, t.rc) {
				emit(m.code, t.id, t.from, t.to, m.detail)
			}
		}
	}
	return out
}

func isDemotionCoverage(from, to string) bool {
	if m, ok := demotionTable[from]; ok {
		return m[to]
	}
	return false
}

func isPromotionToMechanicalCoverage(from, to string) bool {
	if to != "mechanical" {
		return false
	}
	switch from {
	case "(new)", "pending_implementation", "research_required", "behavioral_only":
		return true
	}
	return false
}

type transitionMsg struct {
	code   string
	detail string
}

func checkR2(ctx Context, classID, from, to string, rc RegistryClassMeta) []transitionMsg {
	adr := strings.TrimSpace(rc.AdrReference)
	if adr == "" {
		return []transitionMsg{{
			code:   "R2_demotion_missing_adr",
			detail: "demotion requires adr_reference field on the rule_class pointing to constitution/ADR-*.md",
		}}
	}
	if !strings.HasPrefix(adr, "constitution/ADR-") || !strings.HasSuffix(adr, ".md") {
		return []transitionMsg{{
			code:   "R2_demotion_invalid_adr_format",
			detail: fmt.Sprintf("adr_reference %q must match constitution/ADR-*.md", adr),
		}}
	}
	if ctx.ExistingPaths == nil || !ctx.ExistingPaths[filepath.ToSlash(adr)] {
		return []transitionMsg{{
			code:   "R2_demotion_adr_unresolved",
			detail: fmt.Sprintf("adr_reference %q does not resolve to an existing file under <repo>", adr),
		}}
	}
	return nil
}

func checkR3(ctx Context, classID, from, to string, rc RegistryClassMeta) []transitionMsg {
	allow := map[string]bool{}
	if ctx.RegistryNew != nil {
		for _, s := range ctx.RegistryNew.HelperAllowlist {
			allow[s] = true
		}
	}
	requiredKinds := map[string]bool{}
	if ctx.RegistryNew != nil {
		for _, k := range ctx.RegistryNew.BindingRequiredKinds {
			requiredKinds[k] = true
		}
	}
	var out []transitionMsg
	for _, ex := range rc.Executors {
		if ex.ExecutorKind == "runtime_state_machine_phase" {
			continue
		}
		if ex.ExecutorKind != "" && !requiredKinds[ex.ExecutorKind] {
			continue
		}
		if ex.Symbol == "" {
			out = append(out, transitionMsg{
				code:   "R3_promotion_missing_executor",
				detail: fmt.Sprintf("promotion to mechanical requires symbol_exists; symbol %q not found in %s", "(empty)", ex.File),
			})
			continue
		}
		if allow[ex.Symbol] {
			continue
		}
		if !strings.HasSuffix(ex.File, ".go") {
			continue
		}
		syms := map[string]bool{}
		if ctx.FileSymbols != nil {
			syms = ctx.FileSymbols[filepath.ToSlash(ex.File)]
		}
		if syms[ex.Symbol] {
			continue
		}
		out = append(out, transitionMsg{
			code:   "R3_promotion_missing_executor",
			detail: fmt.Sprintf("promotion to mechanical requires symbol_exists; symbol %q not found in %s", ex.Symbol, ex.File),
		})
	}
	return out
}
