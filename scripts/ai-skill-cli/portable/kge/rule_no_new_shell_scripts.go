package kge

import (
	"strings"
)

// NoNewShellScriptsRule blocks newly added `.sh` files (Go-first policy).
// Modifications to existing `.sh` files are allowed.
//
// Opt-out: standalone `[skip-go-migration]` trailer in CommitMsg.
type NoNewShellScriptsRule struct{}

func (NoNewShellScriptsRule) ID() string { return "rule.no_new_shell_scripts" }
func (NoNewShellScriptsRule) Kind() Kind { return KindValidation }

func (NoNewShellScriptsRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapAddedPaths, CapCommitMsg}
}

func (NoNewShellScriptsRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-go-migration]" {
			return nil
		}
	}
	var newShells []string
	for _, f := range ctx.AddedPaths {
		if strings.HasSuffix(f, ".sh") {
			newShells = append(newShells, f)
		}
	}
	if len(newShells) == 0 {
		return nil
	}
	return []Finding{{
		RuleID:   "rule.no_new_shell_scripts",
		Severity: SeverityError,
		Code:     "no_new_shell_scripts",
		Message: "new shell script(s) staged: " + strings.Join(newShells, ", ") +
			" — cross-platform policy requires Go implementation instead of .sh",
	}}
}
