package kge

import (
	"path/filepath"
	"regexp"
	"strings"
)

var (
	addedRouteIDRE   = regexp.MustCompile(`^\+\s+-\s+id:\s+(route\.[\w.\-]+)\s*$`)
	manualActivationRE = regexp.MustCompile(`^\+\s+manual_activation:\s*$`)
	addedTargetKeyRE = regexp.MustCompile(`^\+\s+target_key:\s+(\S+)\s*$`)
)

const routingRegistryPath = "knowledge/runtime/routing-registry.yaml"

// RuntimeTriggerWiringRule blocks orphan routing-registry routes and
// runtime/*.yaml target_keys without discovery / consumer / manual_activation.
//
// Opt-out: standalone `[skip-runtime-trigger-wiring]` trailer in CommitMsg.
//
// Adapter supplies PathDiffs for staged routing/runtime YAML and SearchCorpus
// covering discovery YAML, routing-registry, and Go sources to search.
type RuntimeTriggerWiringRule struct{}

func (RuntimeTriggerWiringRule) ID() string { return "rule.runtime_trigger_wiring" }
func (RuntimeTriggerWiringRule) Kind() Kind { return KindValidation }

func (RuntimeTriggerWiringRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapStagedDiff, CapSearchCorpus, CapStagedContent}
}

func (RuntimeTriggerWiringRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-runtime-trigger-wiring]" {
			return nil
		}
	}
	hasRoutingDiff := false
	var runtimeYamls []string
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if s == routingRegistryPath {
			hasRoutingDiff = true
			continue
		}
		if strings.HasPrefix(s, "runtime/") && strings.HasSuffix(s, ".yaml") {
			runtimeYamls = append(runtimeYamls, s)
		}
	}
	if !hasRoutingDiff && len(runtimeYamls) == 0 {
		return nil
	}

	var violations []string
	// SearchCorpus must be discovery YAML + Go sources (not routing-registry),
	// matching legacy routeWiredInTree. Target keys also check FileContents
	// routing-registry (legacy targetKeyConsumedInTree).
	corpus := ctx.SearchCorpus
	registryBody := ""
	if ctx.FileContents != nil {
		registryBody = ctx.FileContents[routingRegistryPath]
	}

	if hasRoutingDiff {
		diff := ""
		if ctx.PathDiffs != nil {
			diff = ctx.PathDiffs[routingRegistryPath]
		}
		added := ParseAddedRouteIDs(diff)
		annotated := ParseManualAnnotatedRouteIDs(diff)
		for _, id := range added {
			if annotated[id] {
				continue
			}
			if strings.Contains(corpus, id) {
				continue
			}
			violations = append(violations, "new route `"+id+"` in routing-registry has no discovery signal, Go consumer, or manual_activation annotation")
		}
	}

	for _, yamlPath := range runtimeYamls {
		diff := ""
		if ctx.PathDiffs != nil {
			diff = ctx.PathDiffs[yamlPath]
		}
		for _, key := range ParseAddedTargetKeys(diff) {
			if strings.Contains(corpus, key) || strings.Contains(registryBody, key) {
				continue
			}
			violations = append(violations, "new target_key `"+key+"` in "+yamlPath+" has no consumer (no routing-registry / Go source reference)")
		}
	}

	if len(violations) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.runtime_trigger_wiring",
		Severity: SeverityError,
		Code:     "runtime_trigger_wiring",
		Message: "runtime-trigger-wiring: staged change introduces orphan runtime surface(s) per governance/lifecycle/system-upgrade-governance.yaml §define_runtime_trigger_flow:\n    - " +
			strings.Join(violations, "\n    - ") +
			"\n  Wire each new route to a discovery signal (runtime/cognitive-modes-discovery.yaml) OR a commit-msg validator (scripts/ai-skill-cli/internal/app/hooks.go) OR add a `manual_activation: { reason: <enum> }` annotation. For new target_keys, wire a routing-registry consumer or Go validator that queries the projection. Add `[skip-runtime-trigger-wiring]` (standalone trailer line) for doc-only / annotation-only / pre-existing-state edits.",
	}}
}

// ParseAddedRouteIDs returns route ids added in a unified diff.
func ParseAddedRouteIDs(diff string) []string {
	var out []string
	seen := map[string]bool{}
	for _, line := range strings.Split(diff, "\n") {
		m := addedRouteIDRE.FindStringSubmatch(line)
		if len(m) != 2 {
			continue
		}
		id := m[1]
		if seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, id)
	}
	return out
}

// ParseManualAnnotatedRouteIDs returns route ids whose added hunk includes
// a `manual_activation:` annotation.
func ParseManualAnnotatedRouteIDs(diff string) map[string]bool {
	out := map[string]bool{}
	currentID := ""
	for _, line := range strings.Split(diff, "\n") {
		if m := addedRouteIDRE.FindStringSubmatch(line); len(m) == 2 {
			currentID = m[1]
			continue
		}
		if currentID != "" && manualActivationRE.MatchString(line) {
			out[currentID] = true
		}
	}
	return out
}

// ParseAddedTargetKeys returns target_key values added in a unified diff.
func ParseAddedTargetKeys(diff string) []string {
	var out []string
	seen := map[string]bool{}
	for _, line := range strings.Split(diff, "\n") {
		m := addedTargetKeyRE.FindStringSubmatch(line)
		if len(m) != 2 {
			continue
		}
		key := m[1]
		if seen[key] {
			continue
		}
		seen[key] = true
		out = append(out, key)
	}
	return out
}
