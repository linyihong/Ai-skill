package kge

import (
	"path/filepath"
	"strings"
)

// MarkdownYamlSyncRule enforces sibling-pair markdown/YAML co-change:
// if a staged .md has a sibling .yaml that exists on disk (CapRepoFS),
// that .yaml must also be staged. Cross-path companions are out of scope.
//
// Opt-out: standalone `[skip-markdown-yaml-sync]` trailer in CommitMsg.
type MarkdownYamlSyncRule struct{}

func (MarkdownYamlSyncRule) ID() string { return "rule.markdown_yaml_sync" }
func (MarkdownYamlSyncRule) Kind() Kind { return KindValidation }

func (MarkdownYamlSyncRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapStagedPaths, CapRepoFS, CapCommitMsg}
}

func (MarkdownYamlSyncRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-markdown-yaml-sync]" {
			return nil
		}
	}

	stagedSet := make(map[string]bool, len(ctx.StagedPaths))
	for _, s := range ctx.StagedPaths {
		stagedSet[filepath.ToSlash(s)] = true
	}

	var out []Finding
	for _, s := range ctx.StagedPaths {
		s = filepath.ToSlash(s)
		if !strings.HasSuffix(s, ".md") {
			continue
		}
		sibling := strings.TrimSuffix(s, ".md") + ".yaml"
		if !ctx.ExistingPaths[sibling] {
			continue
		}
		if !stagedSet[sibling] {
			out = append(out, Finding{
				RuleID:   "rule.markdown_yaml_sync",
				Severity: SeverityError,
				Code:     "markdown_yaml_sync",
				Path:     s,
				Message:  "markdown-yaml-sync: " + s + " staged but sibling companion " + sibling + " not staged. Canonical .md edits typically need YAML companion update. If markdown-only change is intentional, use [skip-markdown-yaml-sync].",
			})
		}
	}
	return out
}
