package kge

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var planTreeFolderNameRE = regexp.MustCompile(`^\d{2}-`)
var planFlatMainBasenameRE = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}-\d{4}-[a-z0-9-]+$`)

// PlanTreeFolderConventionRule emits advisory folder-shape findings
// (depth, NN- filename, flat multi-file clusters).
//
// Opt-out: standalone `[skip-plan-tree-folder-convention]` trailer.
//
// CapRepoFS DirListings may supply `plans/active` / `plans/archived` basenames
// for flat-cluster detection; when absent, only path-shape checks run.
type PlanTreeFolderConventionRule struct{}

func (PlanTreeFolderConventionRule) ID() string { return "rule.plan_tree_folder_convention" }
func (PlanTreeFolderConventionRule) Kind() Kind { return KindAdvisory }

func (PlanTreeFolderConventionRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths}
}

func (PlanTreeFolderConventionRule) Validate(ctx Context) []Finding {
	if hasOptOut(ctx.CommitMsg, "[skip-plan-tree-folder-convention]") {
		return nil
	}
	var warnings []string
	for _, rel := range StagedPlanPaths(ctx.StagedPaths) {
		segs := strings.Split(rel, "/")
		if len(segs) < 3 {
			continue
		}
		base := segs[len(segs)-1]
		inEvidence := false
		for i, seg := range segs {
			if i >= 2 && seg == "evidence" {
				inEvidence = true
				break
			}
		}
		depth := len(segs) - 2
		if depth >= 3 && !inEvidence {
			warnings = append(warnings, fmt.Sprintf("%s: nested depth %d (recommend < 3, consider splitting into independent main plan)", rel, depth))
		}
		if depth >= 2 {
			if inEvidence {
				if !strings.HasSuffix(strings.ToLower(base), ".md") {
					warnings = append(warnings, fmt.Sprintf("%s: evidence files should be .md", rel))
				}
				continue
			}
			if base != "_plan.md" && !planTreeFolderNameRE.MatchString(base) {
				warnings = append(warnings, fmt.Sprintf("%s: filename should be `_plan.md` or `NN-<slug>.md`", rel))
			}
		}
	}
	warnings = append(warnings, flatClusterWarningsFromListings(ctx.StagedPaths, ctx.DirListings)...)
	if len(warnings) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_tree_folder_convention",
		Severity: SeverityWarning,
		Code:     "plan_tree_folder_convention",
		Message: "plan-tree-folder-convention (warning): UI convention advisories (non-blocking):\n    - " +
			strings.Join(warnings, "\n    - ") +
			"\n  Folder shape is a recommendation — frontmatter `parent` is the source of truth." +
			"\n  Opt-out: standalone `[skip-plan-tree-folder-convention]` trailer.",
	}}
}

func flatClusterWarningsFromListings(staged []string, listings map[string][]string) []string {
	if len(staged) == 0 || listings == nil {
		return nil
	}
	stagedSet := map[string]bool{}
	for _, s := range staged {
		stagedSet[filepath.ToSlash(s)] = true
	}
	var warnings []string
	for _, loc := range []string{"active", "archived"} {
		dir := "plans/" + loc
		names := listings[dir]
		if len(names) == 0 {
			continue
		}
		for _, c := range scanFlatPlanClustersFromBasenames(loc, names) {
			touches := false
			for _, p := range append([]string{c.mainRel}, c.companionRels...) {
				if stagedSet[p] {
					touches = true
					break
				}
			}
			if !touches {
				continue
			}
			warnings = append(warnings, fmt.Sprintf(
				"flat multi-file cluster %s (%d files) — move to %s/_plan.md + NN-<suffix>.md (`ai-skill plans folderize --cluster %s --dry-run`)",
				c.base, 1+len(c.companionRels), c.folderRel, c.base,
			))
		}
	}
	return warnings
}

type flatClusterLite struct {
	base          string
	mainRel       string
	folderRel     string
	companionRels []string
}

func scanFlatPlanClustersFromBasenames(loc string, names []string) []flatClusterLite {
	set := map[string]bool{}
	var basenames []string
	for _, name := range names {
		if !strings.HasSuffix(strings.ToLower(name), ".md") {
			continue
		}
		base := strings.TrimSuffix(name, filepath.Ext(name))
		if !planFlatMainBasenameRE.MatchString(base) {
			continue
		}
		if !set[base] {
			basenames = append(basenames, base)
			set[base] = true
		}
	}
	sort.Strings(basenames)
	seen := map[string]bool{}
	var out []flatClusterLite
	for _, mainBase := range basenames {
		if seen[mainBase] {
			continue
		}
		prefix := mainBase + "-"
		var companions []string
		for other := range set {
			if other == mainBase {
				continue
			}
			if strings.HasPrefix(other, prefix) {
				companions = append(companions, filepath.ToSlash(filepath.Join("plans", loc, other+".md")))
			}
		}
		if len(companions) == 0 {
			continue
		}
		sort.Strings(companions)
		seen[mainBase] = true
		out = append(out, flatClusterLite{
			base:          mainBase,
			mainRel:       filepath.ToSlash(filepath.Join("plans", loc, mainBase+".md")),
			folderRel:     filepath.ToSlash(filepath.Join("plans", loc, mainBase)),
			companionRels: companions,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].mainRel < out[j].mainRel })
	return out
}
