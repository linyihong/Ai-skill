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
// Two detections, one invariant ("no new non-Go automation entrypoint"):
//   - by extension: ForbiddenNonGoScriptExts, always on;
//   - by content: an added file whose content starts with a shebang ("#!"),
//     whatever its name. Only when the adapter provides CapStagedContent with
//     that path in FileContents; without content the rule degrades to the
//     extension check (not required, so adapters without content keep
//     working).
//
// Exemption (content detection only): shebang files directly inside one of
// ctx.HookAdapterDirs are git hook adapters, which the Go-first policy allows
// as thin entry shells. The directory comes from repository configuration
// (core.hooksPath), not from an enumerated file list. Extension matches are
// not exempted anywhere.
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

// isShebangScript reports whether content starts with an interpreter line.
func isShebangScript(content string) bool {
	return strings.HasPrefix(content, "#!")
}

// inHookAdapterDir reports whether path is a direct child of one of dirs
// (git only runs hooks at the top level of the hooks directory).
func inHookAdapterDir(path string, dirs []string) bool {
	p := filepath.ToSlash(path)
	for _, d := range dirs {
		d = strings.TrimSuffix(filepath.ToSlash(strings.TrimSpace(d)), "/")
		if d == "" || d == "." {
			continue
		}
		if strings.HasPrefix(p, d+"/") && !strings.Contains(strings.TrimPrefix(p, d+"/"), "/") {
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
	contentAvailable := ctx.Has(CapStagedContent)
	for _, f := range ctx.AddedPaths {
		if isForbiddenNonGoScript(f) {
			banned = append(banned, f)
			continue
		}
		if !contentAvailable {
			continue
		}
		if body, ok := ctx.FileContents[f]; ok && isShebangScript(body) && !inHookAdapterDir(f, ctx.HookAdapterDirs) {
			banned = append(banned, f+" (shebang)")
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
