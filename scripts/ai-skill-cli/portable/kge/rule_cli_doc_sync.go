package kge

import (
	"path/filepath"
	"regexp"
	"strings"
)

const commandContractPath = "scripts/ai-skill-cli/docs/command-contract.md"

// CLIDocSyncRule is Phase 0 Mini Spike Rule B — doc/co-change proxy.
// Requires staged paths + staged diff from the adapter; never execs git.
type CLIDocSyncRule struct{}

func (CLIDocSyncRule) ID() string { return "rule.cli_doc_sync" }
func (CLIDocSyncRule) Kind() Kind { return KindValidation }

func (CLIDocSyncRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapStagedPaths, CapStagedDiff, CapCommitMsg}
}

func (CLIDocSyncRule) Validate(ctx Context) []Finding {
	for _, line := range strings.Split(ctx.CommitMsg, "\n") {
		if strings.TrimSpace(line) == "[skip-cli-doc-sync]" {
			return nil
		}
	}

	cliSourceStaged := false
	docStaged := false
	for _, s := range ctx.StagedPaths {
		if strings.HasPrefix(s, "scripts/ai-skill-cli/internal/app/") && strings.HasSuffix(s, ".go") {
			cliSourceStaged = true
		}
		if s == commandContractPath || filepath.ToSlash(s) == commandContractPath {
			docStaged = true
		}
	}
	if !cliSourceStaged || docStaged {
		return nil
	}

	diff := ctx.StagedDiff
	patterns := []string{
		`+\tcase "run `,
		`+\tcase "obligations"`,
		`+func runCommitMsgHook`,
		`+func runPrePushHook`,
		`+func runPreCommitHook`,
		`+func buildRuntimeObligationsResult`,
	}
	for _, p := range patterns {
		if strings.Contains(diff, p) {
			return []Finding{{
				RuleID:   "rule.cli_doc_sync",
				Severity: SeverityError,
				Code:     "cli_doc_sync",
				Message:  "cli-doc-sync: CLI source change adds subcommand dispatch / hook handler but scripts/ai-skill-cli/docs/command-contract.md is not staged. Per runtime/cli-modification-policy.yaml gate.cli.command_contract_synced. Use [skip-cli-doc-sync] for non-contract-affecting refactors.",
			}}
		}
	}
	if regexp.MustCompile(`(?m)^\+func run[A-Z][A-Za-z]+Hook\b`).MatchString(diff) {
		return []Finding{{
			RuleID:   "rule.cli_doc_sync",
			Severity: SeverityError,
			Code:     "cli_doc_sync",
			Message:  "cli-doc-sync: CLI source change adds new runXxxHook function but command-contract.md not staged. See runtime/cli-modification-policy.yaml.",
		}}
	}
	return nil
}
