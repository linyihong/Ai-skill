package kge

import (
	"path/filepath"
	"strings"
)

// EvidenceHierarchyRule requires an evidence citation when a commit claims
// completion and stages code / scenario / governance / runtime work.
//
// Opt-out: standalone `[skip-evidence-hierarchy]` trailer in CommitMsg.
type EvidenceHierarchyRule struct{}

func (EvidenceHierarchyRule) ID() string { return "rule.evidence_hierarchy" }
func (EvidenceHierarchyRule) Kind() Kind { return KindValidation }

func (EvidenceHierarchyRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths}
}

func (EvidenceHierarchyRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-evidence-hierarchy]" {
			return nil
		}
	}
	lower := strings.ToLower(ctx.CommitMsg)
	hasSuccessClaim := false
	for _, phrase := range []string{"complete", "completed", "done", "✅", "完成", "結案"} {
		if strings.Contains(lower, strings.ToLower(phrase)) {
			hasSuccessClaim = true
			break
		}
	}
	if !hasSuccessClaim {
		return nil
	}
	hasCodeWork := false
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if strings.HasSuffix(s, ".go") ||
			strings.HasPrefix(s, "validation/scenarios/") ||
			strings.HasPrefix(s, "runtime/") ||
			strings.HasPrefix(s, "governance/") ||
			strings.HasPrefix(s, "enforcement/") {
			hasCodeWork = true
			break
		}
	}
	if !hasCodeWork {
		return nil
	}
	evidenceMarkers := []string{
		"test pass", "tests pass", "tests green", "fixture", "go test", "all green",
		"exit 0", "scenario", "audit", "runtime validate", "validate pass",
		"commit ", "based on", "per commit", "evidence", "證據",
	}
	for _, m := range evidenceMarkers {
		if strings.Contains(lower, m) {
			return nil
		}
	}
	return []Finding{{
		RuleID:   "rule.evidence_hierarchy",
		Severity: SeverityError,
		Code:     "evidence_hierarchy",
		Message:  "evidence-hierarchy: commit body asserts task completion (e.g., 完成 / done / ✅) without citing evidence — required by enforcement.evidence_hierarchy.contract §confidence_integrity (source: enforcement/evidence-hierarchy.yaml). Add at least one evidence reference (test pass / fixture / scenario id / audit/validate output / commit hash). Use `[skip-evidence-hierarchy]` (standalone trailer) for recovery / rollback / pre-existing-evidence commits.",
	}}
}
