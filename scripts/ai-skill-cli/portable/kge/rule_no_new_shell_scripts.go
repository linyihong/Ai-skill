package kge

import (
	"path/filepath"
	"strings"
)

// ForbiddenNonGoScriptExts are newly-added script/language entrypoints
// blocked by Go-first automation policy. Documentation/config (.md/.yaml)
// and Go sources are not in this list.
var ForbiddenNonGoScriptExts = []string{
	".sh", ".bash", ".zsh", ".fish",
	".py", ".rb", ".pl", ".ps1",
	".js", ".mjs", ".cjs", ".ts",
}

// NoNewShellScriptsRule blocks newly added non-Go script entrypoints.
// Name kept for history; scope is Go-first (not only .sh).
// Modifications to existing forbidden files are allowed (migration window).
//
// Opt-out: standalone `[skip-go-migration]` trailer in CommitMsg.
type NoNewShellScriptsRule struct{}

func (NoNewShellScriptsRule) ID() string { return "rule.no_new_shell_scripts" }
func (NoNewShellScriptsRule) Kind() Kind { return KindValidation }

func (NoNewShellScriptsRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapAddedPaths, CapCommitMsg}
}

func isForbiddenNonGoScript(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	for _, banned := range ForbiddenNonGoScriptExts {
		if ext == banned {
			return true
		}
	}
	return false
}

func (NoNewShellScriptsRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-go-migration]" {
			return nil
		}
	}
	var banned []string
	for _, f := range ctx.AddedPaths {
		if isForbiddenNonGoScript(f) {
			banned = append(banned, f)
		}
	}
	if len(banned) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.no_new_shell_scripts",
		Severity: SeverityError,
		Code:     "no_new_non_go_scripts",
		Message: "new non-Go script entrypoint(s) staged: " + strings.Join(banned, ", ") +
			" — Go-first policy requires ai-skill CLI (Go) instead of .sh/.py/.rb/.js/…",
	}}
}
