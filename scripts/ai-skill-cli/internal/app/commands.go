package app

import (
	"fmt"
	"io"
	"strings"
)

// cliCommand is one entry of the declarative CLI command table (the
// Parameters role for dispatch). Top-level dispatch, `ai-skill help`, each
// command group's usage line and its "unknown subcommand" gate all read this
// table, and TestCommandContractMatchesCommandTable checks it against
// docs/command-contract.md in both directions. Adding a command or
// subcommand therefore means adding it here: a case added only to a group's
// switch is rejected by the group's gate before it is reached.
//
// Subcommands holds the accepted first argument(s) after the command name.
// Nested forms are written with a space, e.g. "run pre-commit" for
// `ai-skill hooks run pre-commit` and "proposals list" for
// `ai-skill router proposals list`.
type cliCommand struct {
	Name        string
	Summary     string
	Subcommands []string
	Run         func(args []string, stdout io.Writer, stderr io.Writer) int
}

var cliCommands []cliCommand

// Populated in init() rather than in the var initializer: the run functions
// read the table (usage lines, gates), which would otherwise form an
// initialization cycle.
func init() {
	cliCommands = []cliCommand{
		{Name: "version", Summary: "print build version metadata", Run: runVersion},
		{Name: "doctor", Summary: "check local runtime and repository readiness", Run: runDoctor},
		{Name: "init-project", Summary: "plan project-local AI tool bootstrap files", Run: runInitProject},
		{Name: "goals", Summary: "inspect or plan project-local goal ledger changes", Run: runGoals,
			Subcommands: []string{"status", "init", "start", "update", "split", "pause", "complete", "cleanup"}},
		{Name: "close-loop", Summary: "inspect repository close-loop readiness", Run: runCloseLoop},
		{Name: "hooks", Summary: "inspect or plan Git hook installation", Run: runHooks,
			Subcommands: []string{
				"install",
				"run pre-commit", "run commit-msg", "run post-commit", "run pre-push",
				"run session-start", "run pre-tool-use", "run post-tool-use", "run user-prompt-submit", "run stop",
			}},
		{Name: "sync-cursor-bundle", Summary: "plan Cursor bundle mirror sync", Run: runSyncCursorBundle},
		{Name: "runtime", Summary: "wrap or inspect runtime tooling", Run: runRuntime,
			Subcommands: []string{"validate", "refresh", "compile", "query", "lookup", "selection-watch", "obligations", "receipt", "workflow-context", "audit", "capability-invoke"}},
		{Name: "roo", Summary: "manage guarded Roo Code settings", Run: runRoo,
			Subcommands: []string{"set-global-custom-instructions"}},
		{Name: "copilot", Summary: "generate GitHub Copilot guided bootstrap prompts", Run: runCopilot,
			Subcommands: []string{"start"}},
		{Name: "glossary", Summary: "validate knowledge/glossary/ entries", Run: runGlossary,
			Subcommands: []string{"validate"}},
		{Name: "scan-checkboxes", Summary: "scan a Markdown file for unchecked/checked task-list items", Run: runScanCheckboxes},
		{Name: "enforcement", Summary: "run enforcement-registry lint, coverage report or transition check", Run: runEnforcement,
			Subcommands: []string{"lint", "coverage", "transition-check"}},
		{Name: "plans", Summary: "render, validate or folderize the plan tree", Run: runPlans,
			Subcommands: []string{"tree", "validate", "folderize"}},
		{Name: "kge", Summary: "Knowledge Governance Engine: check / validate / diagnose (D9 presentation)", Run: runKge,
			Subcommands: []string{"check", "validate", "diagnose"}},
		{Name: "router", Summary: "manage workflow route-candidate proposals (Phase 6.1)", Run: runRouter,
			Subcommands: []string{"proposals list", "proposals record", "proposals promote", "proposals reject", "proposals gc"}},
		{Name: "image-processing", Summary: "verify images / generate synthetic fixtures (methods only)", Run: runImageProcessing,
			Subcommands: []string{"verify", "generate-fixtures"}},
	}
}

func findCommand(name string) (cliCommand, bool) {
	for _, c := range cliCommands {
		if c.Name == name {
			return c, true
		}
	}
	return cliCommand{}, false
}

// hasSubcommand reports whether sub (possibly nested, e.g. "run pre-commit")
// is declared for the command group.
func hasSubcommand(group, sub string) bool {
	c, ok := findCommand(group)
	if !ok {
		return false
	}
	for _, s := range c.Subcommands {
		if s == sub {
			return true
		}
	}
	return false
}

// hasSubcommandHead reports whether head is the first word of a declared
// subcommand of group (e.g. "run" for "run pre-commit").
func hasSubcommandHead(group, head string) bool {
	for _, s := range strings.Split(subcommandChoices(group, ""), "|") {
		if s == head {
			return true
		}
	}
	return false
}

// subcommandChoices renders the declared subcommands of group as "a|b|c".
// With prefix set (e.g. "run "), only nested subcommands under that prefix
// are listed, with the prefix stripped.
func subcommandChoices(group, prefix string) string {
	c, _ := findCommand(group)
	var out []string
	seen := map[string]bool{}
	for _, s := range c.Subcommands {
		if prefix != "" {
			if !strings.HasPrefix(s, prefix) {
				continue
			}
			s = strings.TrimPrefix(s, prefix)
		} else if i := strings.Index(s, " "); i >= 0 {
			s = s[:i]
		}
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return strings.Join(out, "|")
}

// usageLine renders "usage: ai-skill <group> <a|b|c> <suffix>".
func usageLine(group, suffix string) string {
	line := fmt.Sprintf("usage: ai-skill %s <%s>", group, subcommandChoices(group, ""))
	if suffix != "" {
		line += " " + suffix
	}
	return line
}
