package kge

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	enforcementRegistryRel  = "enforcement/enforcement-registry.yaml"
	enforcementCompanionRel = "enforcement/enforcement-registry.md"
)

var topLevelIDRE = regexp.MustCompile(`(?m)^id:\s*([^\s#]+)`)

// EnforcementRuleRegistrySyncRule requires staged enforcement rule YAMLs with
// a top-level `id:` to be bound in enforcement-registry.yaml (or that the
// registry is staged in the same commit).
//
// Opt-out: commit body contains `[skip-enforcement-registry-sync]`.
type EnforcementRuleRegistrySyncRule struct{}

func (EnforcementRuleRegistrySyncRule) ID() string { return "rule.enforcement_rule_registry_sync" }
func (EnforcementRuleRegistrySyncRule) Kind() Kind { return KindValidation }

func (EnforcementRuleRegistrySyncRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapStagedContent, CapBoundPaths}
}

func (EnforcementRuleRegistrySyncRule) Validate(ctx Context) []Finding {
	if strings.Contains(ctx.CommitMsg, "[skip-enforcement-registry-sync]") {
		return nil
	}
	registryStaged := false
	var candidates []string
	for _, p := range ctx.StagedPaths {
		rel := filepath.ToSlash(strings.TrimSpace(p))
		if rel == "" {
			continue
		}
		if rel == enforcementRegistryRel {
			registryStaged = true
			continue
		}
		if rel == enforcementCompanionRel {
			continue
		}
		if !strings.HasPrefix(rel, "enforcement/") {
			continue
		}
		if !strings.HasSuffix(rel, ".yaml") && !strings.HasSuffix(rel, ".yml") {
			continue
		}
		candidates = append(candidates, rel)
	}
	if len(candidates) == 0 {
		return nil
	}

	bound := ctx.BoundPaths
	if bound == nil {
		bound = map[string]bool{}
	}
	var unbound []string
	for _, rel := range candidates {
		content := ""
		if ctx.FileContents != nil {
			content = ctx.FileContents[rel]
		}
		if ExtractTopLevelID(content) == "" {
			continue
		}
		if bound[rel] {
			continue
		}
		unbound = append(unbound, rel)
	}
	if len(unbound) == 0 || registryStaged {
		return nil
	}
	var b strings.Builder
	fmt.Fprintf(&b, "enforcement_rule_registry_sync: %d staged enforcement rule yaml(s) declare top-level id: but are not bound by enforcement-registry.yaml, AND enforcement-registry.yaml is not staged. Either:\n", len(unbound))
	for _, p := range unbound {
		fmt.Fprintf(&b, "  - %s\n", p)
	}
	b.WriteString("  Add a rule_class entry to enforcement/enforcement-registry.yaml (source_files: [<path>]) in this same commit, or use [skip-enforcement-registry-sync] opt-out.")
	return []Finding{{
		RuleID:   "rule.enforcement_rule_registry_sync",
		Severity: SeverityError,
		Code:     "enforcement_rule_registry_sync",
		Message:  strings.TrimRight(b.String(), "\n"),
	}}
}

// ExtractTopLevelID returns the first YAML `id:` value in content.
func ExtractTopLevelID(content string) string {
	m := topLevelIDRE.FindStringSubmatch(content)
	if len(m) < 2 {
		return ""
	}
	return strings.Trim(m[1], `"'`)
}
