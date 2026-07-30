package kge

import (
	"path/filepath"
	"strings"
)

// RuntimeYamlProjectsRule requires runtime/*.yaml to declare
// runtime_projection.enabled: true and target_key.
//
// Opt-out: standalone `[skip-runtime-yaml-projection]` trailer in CommitMsg.
type RuntimeYamlProjectsRule struct{}

func (RuntimeYamlProjectsRule) ID() string { return "rule.runtime_yaml_projects" }
func (RuntimeYamlProjectsRule) Kind() Kind { return KindValidation }

func (RuntimeYamlProjectsRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapStagedPaths, CapStagedContent, CapCommitMsg}
}

func (RuntimeYamlProjectsRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-runtime-yaml-projection]" {
			return nil
		}
	}
	var out []Finding
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if !strings.HasPrefix(s, "runtime/") || !strings.HasSuffix(s, ".yaml") {
			continue
		}
		content, ok := ctx.FileContents[s]
		if !ok {
			continue
		}
		hasProjection := strings.Contains(content, "runtime_projection:") &&
			(strings.Contains(content, "enabled: true") || strings.Contains(content, "enabled:true"))
		hasTargetKey := strings.Contains(content, "target_key:")
		if !hasProjection || !hasTargetKey {
			out = append(out, Finding{
				RuleID:   "rule.runtime_yaml_projects",
				Severity: SeverityError,
				Code:     "runtime_yaml_projects",
				Path:     s,
				Message:  "runtime-yaml-projects: " + s + " missing runtime_projection.enabled:true or target_key. Default rule: runtime/*.yaml must project to runtime.db. If intentional deferral, declare §Deferred Runtime Projection in plan + use [skip-runtime-yaml-projection].",
			})
		}
	}
	return out
}
