package kge

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// StagedPlanPaths returns staged markdown plan paths under plans/active|archived,
// excluding any fixtures/ segment.
func StagedPlanPaths(staged []string) []string {
	var out []string
	for _, s := range staged {
		s = filepath.ToSlash(s)
		if (strings.HasPrefix(s, "plans/active/") || strings.HasPrefix(s, "plans/archived/")) &&
			strings.HasSuffix(strings.ToLower(s), ".md") &&
			!pathContainsFixturesSegment(s) {
			out = append(out, s)
		}
	}
	return out
}

func pathContainsFixturesSegment(rel string) bool {
	for _, seg := range strings.Split(rel, "/") {
		if seg == "fixtures" {
			return true
		}
	}
	return false
}

func planMetaByPath(index []PlanMeta) map[string]PlanMeta {
	out := map[string]PlanMeta{}
	for _, p := range index {
		out[filepath.ToSlash(p.Path)] = p
	}
	return out
}

func hasOptOut(text, trailer string) bool {
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == trailer {
			return true
		}
	}
	return false
}

func displayStatus(s string) string {
	if s == "" {
		return "<missing>"
	}
	return s
}

// PlanTreeFrontmatterRule requires sub/spike plans to declare parent,
// sub_plan_reason, required_for_completion (and delegation fields when enabled).
type PlanTreeFrontmatterRule struct{}

func (PlanTreeFrontmatterRule) ID() string { return "rule.plan_tree_frontmatter" }
func (PlanTreeFrontmatterRule) Kind() Kind { return KindValidation }

func (PlanTreeFrontmatterRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapPlanIndex}
}

func (PlanTreeFrontmatterRule) Validate(ctx Context) []Finding {
	if hasOptOut(ctx.CommitMsg, "[skip-plan-tree-frontmatter]") {
		return nil
	}
	byPath := planMetaByPath(ctx.PlanIndex)
	var violations []string
	for _, rel := range StagedPlanPaths(ctx.StagedPaths) {
		pf, ok := byPath[rel]
		if !ok || !pf.HasFrontmatter {
			continue
		}
		kind := pf.PlanKind
		if kind == "" || (kind != "sub" && kind != "spike") {
			continue
		}
		var missing []string
		if !pf.HasParentField || strings.TrimSpace(pf.Parent) == "" {
			missing = append(missing, "parent")
		}
		if !pf.HasReasonField || pf.SubPlanReason == "" {
			missing = append(missing, "sub_plan_reason (non-empty)")
		}
		if pf.RequiredForCompletion == nil {
			missing = append(missing, "required_for_completion")
		}
		if len(missing) > 0 {
			violations = append(violations, fmt.Sprintf("%s missing: %s", rel, strings.Join(missing, ", ")))
		}
		if pf.DelegationEnabled {
			var dm []string
			if !pf.DelegationHasGoal {
				dm = append(dm, "delegation.brief.goal")
			}
			if !pf.DelegationHasAcceptance {
				dm = append(dm, "delegation.brief.acceptance")
			}
			if !pf.DelegationHasVerification {
				dm = append(dm, "delegation.brief.verification")
			}
			if !pf.DelegationHasModes {
				dm = append(dm, "delegation.execution.modes")
			}
			if len(dm) > 0 {
				violations = append(violations, fmt.Sprintf("%s delegation enabled but missing: %s", rel, strings.Join(dm, ", ")))
			}
		}
	}
	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_tree_frontmatter",
		Severity: SeverityError,
		Code:     "plan_tree_frontmatter",
		Message: "plan-tree-frontmatter: sub/spike plan(s) missing required frontmatter fields:\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  Add `parent: <main-id>`, `sub_plan_reason: <non-empty>` and `required_for_completion: true|false` " +
			"(see plans/active/2026-06-02-1200-plan-tree-hierarchy-governance/01-frontmatter-schema.md)" +
			"\n  Opt-out (emergency only): standalone `[skip-plan-tree-frontmatter]` trailer.",
	}}
}

// PlanTreeArchiveOrderRule blocks archiving a main plan while required
// children are not completed.
type PlanTreeArchiveOrderRule struct{}

func (PlanTreeArchiveOrderRule) ID() string { return "rule.plan_tree_archive_order" }
func (PlanTreeArchiveOrderRule) Kind() Kind { return KindValidation }

func (PlanTreeArchiveOrderRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapPlanIndex}
}

func (PlanTreeArchiveOrderRule) Validate(ctx Context) []Finding {
	if hasOptOut(ctx.CommitMsg, "[skip-plan-tree-archive-order]") {
		return nil
	}
	byPath := planMetaByPath(ctx.PlanIndex)
	var archivedMains []PlanMeta
	for _, rel := range StagedPlanPaths(ctx.StagedPaths) {
		if !strings.HasPrefix(rel, "plans/archived/") {
			continue
		}
		pf, ok := byPath[rel]
		if !ok || !pf.HasFrontmatter || pf.PlanKind != "main" {
			continue
		}
		archivedMains = append(archivedMains, pf)
	}
	if len(archivedMains) == 0 {
		return nil
	}
	var violations []string
	for _, main := range archivedMains {
		if main.ID == "" {
			continue
		}
		for _, p := range ctx.PlanIndex {
			if !p.HasFrontmatter || p.Parent != main.ID {
				continue
			}
			if p.RequiredForCompletion == nil || !*p.RequiredForCompletion {
				continue
			}
			if p.Status == "completed" {
				continue
			}
			violations = append(violations,
				fmt.Sprintf("main %s (%s) blocked by required sub %s (status=%s)",
					main.ID, main.Path, p.Path, displayStatus(p.Status)))
		}
	}
	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_tree_archive_order",
		Severity: SeverityError,
		Code:     "plan_tree_archive_order",
		Message: "plan-tree-archive-order: cannot archive main plan(s) with unfinished required sub-plans:\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  Complete the required sub-plan(s) first or flip required_for_completion: false with rationale." +
			"\n  Opt-out (emergency only): standalone `[skip-plan-tree-archive-order]` trailer.",
	}}
}

// PlanTreeParentReferenceRule blocks dangling parent: pointers on staged sub/spike plans.
type PlanTreeParentReferenceRule struct{}

func (PlanTreeParentReferenceRule) ID() string { return "rule.plan_tree_parent_reference" }
func (PlanTreeParentReferenceRule) Kind() Kind { return KindValidation }

func (PlanTreeParentReferenceRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapPlanIndex}
}

func (PlanTreeParentReferenceRule) Validate(ctx Context) []Finding {
	if hasOptOut(ctx.CommitMsg, "[skip-plan-tree-parent-reference]") {
		return nil
	}
	byPath := planMetaByPath(ctx.PlanIndex)
	var stagedSubs []PlanMeta
	for _, rel := range StagedPlanPaths(ctx.StagedPaths) {
		pf, ok := byPath[rel]
		if !ok || !pf.HasFrontmatter {
			continue
		}
		if pf.PlanKind != "sub" && pf.PlanKind != "spike" {
			continue
		}
		if strings.TrimSpace(pf.Parent) == "" {
			continue
		}
		stagedSubs = append(stagedSubs, pf)
	}
	if len(stagedSubs) == 0 {
		return nil
	}
	known := map[string]bool{}
	for _, p := range ctx.PlanIndex {
		if p.HasFrontmatter && p.ID != "" {
			known[p.ID] = true
		}
	}
	var violations []string
	for _, p := range stagedSubs {
		if !known[p.Parent] {
			violations = append(violations,
				fmt.Sprintf("%s references parent: %q which does not resolve to any plan id", p.Path, p.Parent))
		}
	}
	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_tree_parent_reference",
		Severity: SeverityError,
		Code:     "plan_tree_parent_reference",
		Message: "plan-tree-parent-reference: dangling parent pointer(s) detected:\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  Either fix the parent id or create the referenced main plan first." +
			"\n  Opt-out (emergency only): standalone `[skip-plan-tree-parent-reference]` trailer.",
	}}
}

// PlanTreeUniqueIDRule blocks duplicate frontmatter ids that touch staged paths.
type PlanTreeUniqueIDRule struct{}

func (PlanTreeUniqueIDRule) ID() string { return "rule.plan_tree_unique_id" }
func (PlanTreeUniqueIDRule) Kind() Kind { return KindValidation }

func (PlanTreeUniqueIDRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapPlanIndex}
}

func (PlanTreeUniqueIDRule) Validate(ctx Context) []Finding {
	if hasOptOut(ctx.CommitMsg, "[skip-plan-tree-unique-id]") {
		return nil
	}
	byID := map[string][]string{}
	for _, p := range ctx.PlanIndex {
		if !p.HasFrontmatter || p.ID == "" {
			continue
		}
		byID[p.ID] = append(byID[p.ID], filepath.ToSlash(p.Path))
	}
	stagedSet := map[string]bool{}
	for _, s := range ctx.StagedPaths {
		stagedSet[filepath.ToSlash(s)] = true
	}
	var ids []string
	for id := range byID {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	var violations []string
	for _, id := range ids {
		paths := byID[id]
		if len(paths) < 2 {
			continue
		}
		touchesStage := false
		for _, p := range paths {
			if stagedSet[p] {
				touchesStage = true
				break
			}
		}
		if !touchesStage {
			continue
		}
		sort.Strings(paths)
		violations = append(violations, fmt.Sprintf("id %q appears in: %s", id, strings.Join(paths, ", ")))
	}
	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.plan_tree_unique_id",
		Severity: SeverityError,
		Code:     "plan_tree_unique_id",
		Message: "plan-tree-unique-id: duplicate frontmatter id(s) detected:\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  Plan ids must be globally unique across active + archived." +
			"\n  Opt-out (emergency only): standalone `[skip-plan-tree-unique-id]` trailer.",
	}}
}
