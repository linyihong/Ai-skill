package kge

import (
	"path/filepath"
	"strings"
)

const glossaryPath = "knowledge/glossary/ai-skill.md"

// GlossaryRetroOwnRule requires glossary staging when framework vocabulary
// surfaces change (cognitive-modes / economics / ecosystem).
//
// Opt-out: standalone `[skip-glossary-retro-own]` trailer in CommitMsg.
type GlossaryRetroOwnRule struct{}

func (GlossaryRetroOwnRule) ID() string { return "rule.glossary_retro_own" }
func (GlossaryRetroOwnRule) Kind() Kind { return KindValidation }

func (GlossaryRetroOwnRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapStagedPaths, CapCommitMsg}
}

func (GlossaryRetroOwnRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-glossary-retro-own]" {
			return nil
		}
	}
	framework := false
	glossary := false
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if strings.HasPrefix(s, "runtime/cognitive-modes") && strings.HasSuffix(s, ".yaml") {
			framework = true
		}
		if strings.HasPrefix(s, "runtime/economics/") {
			framework = true
		}
		if strings.HasPrefix(s, "ecosystem/") {
			framework = true
		}
		if s == glossaryPath {
			glossary = true
		}
	}
	if !framework || glossary {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.glossary_retro_own",
		Severity: SeverityError,
		Code:     "glossary_retro_own",
		Path:     glossaryPath,
		Message:  "glossary-retro-own: staged change touches framework cognitive vocabulary surface (runtime/cognitive-modes*.yaml, runtime/economics/, ecosystem/) but knowledge/glossary/ai-skill.md is not staged. Per plans/active/2026-05-25-1000-context-language-glossary-system.md Phase 6 and runtime/cli-modification-policy.yaml gate.glossary.retro_own_required, new framework terms must retro-own a canonical glossary entry. Use [skip-glossary-retro-own] (standalone trailer line) if this change is a typo / refactor / comment-only edit and introduces no new term.",
	}}
}
