package kge

import (
	"strconv"
	"strings"
)

// BootstrapEntryPaths are AI-tool entry files that must stay thin pointers
// to canonical bootstrap sources (runtime/bootstrap-entry-points.yaml).
var BootstrapEntryPaths = []string{
	"CLAUDE.md",
	".cursor/rules/ai-skill-bootstrap.mdc",
	".roomodes",
	"AGENTS.md",
}

// BootstrapEntryThinnessRule blocks bloated tool-entry files.
//
// Opt-out: standalone `[skip-bootstrap-thinness]` trailer in CommitMsg.
type BootstrapEntryThinnessRule struct {
	MaxLines int // default 30
}

func (BootstrapEntryThinnessRule) ID() string { return "rule.bootstrap_entry_thinness" }
func (BootstrapEntryThinnessRule) Kind() Kind { return KindValidation }

func (BootstrapEntryThinnessRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapStagedPaths, CapStagedContent, CapCommitMsg}
}

func (r BootstrapEntryThinnessRule) maxLines() int {
	if r.MaxLines <= 0 {
		return 30
	}
	return r.MaxLines
}

var bootstrapForbiddenSubs = []string{
	"FAST/NORMAL/DEEP/FORENSIC/RECOVERY",
	"INDEX_ONLY/SUMMARY_FIRST/CHECKLIST_FIRST/SOURCE_BACKED/GRAPH_ASSISTED",
	"LIGHT/STANDARD/STRICT/LOCKDOWN",
	"NONE/EPISODIC/DECISION_REPLAY/FAILURE_REPLAY/PROJECT_CONTEXT",
	"Bootstrap: rules=✓ phase=phase.bootstrap obligations=",
	"| 維度 | 值 | 理由 |",
}

func (r BootstrapEntryThinnessRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-bootstrap-thinness]" {
			return nil
		}
	}

	entrySet := make(map[string]bool, len(BootstrapEntryPaths))
	for _, p := range BootstrapEntryPaths {
		entrySet[p] = true
	}

	maxLines := r.maxLines()
	var out []Finding
	for _, path := range ctx.StagedPaths {
		if !entrySet[path] {
			continue
		}
		content, ok := ctx.FileContents[path]
		if !ok {
			// Staged for deletion or adapter could not read — ignore.
			continue
		}
		lineCount := countLines(content)
		if lineCount > maxLines {
			out = append(out, Finding{
				RuleID:   r.ID(),
				Severity: SeverityError,
				Code:     "bootstrap_entry_thinness_lines",
				Path:     path,
				Message: "bootstrap-entry-thinness: " + path + " is " + strconv.Itoa(lineCount) +
					" lines (max " + strconv.Itoa(maxLines) + "); move obligation content to CORE_BOOTSTRAP.md or ai-tools/agent/<tool>.md per runtime/bootstrap-entry-points.yaml.",
			})
			continue
		}
		for _, sub := range bootstrapForbiddenSubs {
			if strings.Contains(content, sub) {
				out = append(out, Finding{
					RuleID:   r.ID(),
					Severity: SeverityError,
					Code:     "bootstrap_entry_thinness_content",
					Path:     path,
					Message: "bootstrap-entry-thinness: " + path + " contains canonical content fragment '" +
						sub + "'; this belongs in CORE_BOOTSTRAP.md / models/cognitive-modes/, not in tool entries.",
				})
				break
			}
		}
	}
	return out
}
