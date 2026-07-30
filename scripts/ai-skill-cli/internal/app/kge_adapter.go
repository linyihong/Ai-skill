package app

import (
	"os/exec"
	"strings"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
)

// kgeFindingsMessage flattens engine findings into the legacy commit-msg
// validator string shape (empty = pass).
func kgeFindingsMessage(findings []kge.Finding) string {
	if len(findings) == 0 {
		return ""
	}
	parts := make([]string, 0, len(findings))
	for _, f := range findings {
		if f.Message != "" {
			parts = append(parts, f.Message)
		}
	}
	return strings.Join(parts, "\n  - ")
}

// runKGECognitiveCost adapts commit-msg modes into a portable KGE Context
// and runs Rule A (cognitive_cost). Git-free.
func runKGECognitiveCost(modes map[string]string) string {
	ctx := kge.Context{
		Modes: modes,
		Provided: map[kge.CapabilityID]bool{
			kge.CapModes: true,
		},
	}
	eng := kge.NewEngine(kge.CognitiveCostRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// stagedDiffCached runs `git diff --cached` for paths. On failure returns
// empty diff (legacy validateCLIDocSync treated git errors as pass / no match).
func stagedDiffCached(root string, pathspecs ...string) string {
	args := append([]string{"-C", root, "diff", "--cached", "--"}, pathspecs...)
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

// runKGECLIDocSync adapts commit-msg + staged paths into KGE Context.
// Git stays in this adapter; the portable rule never execs git.
func runKGECLIDocSync(text string, staged []string, root string) string {
	diff := stagedDiffCached(root, "scripts/ai-skill-cli/internal/app/")
	ctx := kge.Context{
		RepoRoot:    root,
		CommitMsg:   text,
		StagedPaths: staged,
		StagedDiff:  diff,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapStagedPaths: true,
			kge.CapStagedDiff:  true,
		},
	}
	eng := kge.NewEngine(kge.CLIDocSyncRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}
