package kge

import (
	"fmt"
	"path/filepath"
	"strings"
)

// PlanEvidenceConventionRule enforces evidence/ README + index + folder main
// (_plan.md; no flat sibling) when evidence/ is in play for the commit.
//
// Opt-out: standalone `[skip-plan-evidence]` trailer in CommitMsg.
type PlanEvidenceConventionRule struct{}

func (PlanEvidenceConventionRule) ID() string { return "rule.plan_evidence_convention" }
func (PlanEvidenceConventionRule) Kind() Kind { return KindValidation }

func (PlanEvidenceConventionRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapRepoFS, CapStagedContent}
}

func (PlanEvidenceConventionRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-plan-evidence]" {
			return nil
		}
	}
	planDirs := collectPlanDirsForEvidence(ctx)
	if len(planDirs) == 0 {
		return nil
	}

	var violations []string
	for planDir := range planDirs {
		mainRel := planDir + "/_plan.md"
		if !pathExists(ctx, mainRel) {
			violations = append(violations, fmt.Sprintf(
				"%s: missing _plan.md — when evidence/ is used, the main plan must live in the folder as _plan.md (not a sibling <slug>.md)",
				planDir))
		}
		flatSibling := planDir + ".md"
		if pathExists(ctx, flatSibling) {
			violations = append(violations, fmt.Sprintf(
				"%s: flat sibling %s still exists — move it to %s/_plan.md before using evidence/ (ai-skill plans folderize or git mv)",
				planDir, flatSibling, planDir))
		}

		readmeRel := planDir + "/evidence/README.md"
		readmeBody, readmeOK := fileContent(ctx, readmeRel)
		if !readmeOK {
			violations = append(violations, fmt.Sprintf("%s: missing evidence/README.md (required when using evidence/)", planDir))
			continue
		}
		if miss := readmeHasRequiredSections(readmeBody); len(miss) > 0 {
			violations = append(violations, fmt.Sprintf("%s: evidence/README.md missing: %s", planDir, strings.Join(miss, ", ")))
		}
		files := dirListing(ctx, planDir+"/evidence")
		for _, name := range files {
			if strings.EqualFold(name, "README.md") {
				continue
			}
			if !strings.Contains(readmeBody, name) {
				violations = append(violations, fmt.Sprintf("%s: evidence/%s not listed in evidence/README.md Run 索引", planDir, name))
			}
		}
	}

	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_evidence_convention",
		Severity: SeverityError,
		Code:     "plan_evidence_convention",
		Message: "plan-evidence-convention: evidence/ directory rule violation(s):\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  See governance/lifecycle/plan-evidence.md and plans/templates/plan-evidence/README.md." +
			"\n  Opt-out (emergency only): standalone `[skip-plan-evidence]` trailer.",
	}}
}

func collectPlanDirsForEvidence(ctx Context) map[string]bool {
	planDirs := map[string]bool{}
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if isPlanEvidencePath(s) && strings.HasSuffix(strings.ToLower(s), ".md") {
			if dir, ok := planDirFromEvidencePath(s); ok {
				planDirs[dir] = true
			}
		}
	}
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if !strings.HasPrefix(s, "plans/active/") && !strings.HasPrefix(s, "plans/archived/") {
			continue
		}
		if !strings.HasSuffix(strings.ToLower(s), ".md") {
			continue
		}
		parts := strings.Split(s, "/")
		if len(parts) != 3 {
			continue
		}
		planDir := strings.TrimSuffix(s, ".md")
		if pathExists(ctx, planDir+"/evidence") {
			planDirs[planDir] = true
		}
	}
	return planDirs
}

func isPlanEvidencePath(rel string) bool {
	rel = filepath.ToSlash(rel)
	if !strings.HasPrefix(rel, "plans/active/") && !strings.HasPrefix(rel, "plans/archived/") {
		return false
	}
	return strings.Contains(rel, "/evidence/")
}

func planDirFromEvidencePath(rel string) (string, bool) {
	rel = filepath.ToSlash(rel)
	idx := strings.Index(rel, "/evidence/")
	if idx < 0 {
		return "", false
	}
	return rel[:idx], true
}

func pathExists(ctx Context, rel string) bool {
	rel = filepath.ToSlash(rel)
	if ctx.ExistingPaths != nil && ctx.ExistingPaths[rel] {
		return true
	}
	if ctx.FileContents != nil {
		if _, ok := ctx.FileContents[rel]; ok {
			return true
		}
	}
	if ctx.DirListings != nil {
		if _, ok := ctx.DirListings[rel]; ok {
			return true
		}
	}
	return false
}

func fileContent(ctx Context, rel string) (string, bool) {
	rel = filepath.ToSlash(rel)
	if ctx.FileContents == nil {
		return "", false
	}
	body, ok := ctx.FileContents[rel]
	return body, ok
}

func dirListing(ctx Context, dir string) []string {
	dir = filepath.ToSlash(dir)
	if ctx.DirListings == nil {
		return nil
	}
	return ctx.DirListings[dir]
}

func readmeHasRequiredSections(body string) (missing []string) {
	if !strings.Contains(body, "引用規則") && !strings.Contains(body, "Citation rule") {
		missing = append(missing, "引用規則 section")
	}
	if !strings.Contains(body, "Run 索引") && !strings.Contains(body, "## Index") {
		missing = append(missing, "Run 索引 table")
	}
	return missing
}
