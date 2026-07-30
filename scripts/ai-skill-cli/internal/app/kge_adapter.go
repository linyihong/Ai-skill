package app

import (
	"os"
	"os/exec"
	"path/filepath"
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

// countKGEAdvisories runs advisory-only rules (D9 commit-msg count path).
// Does not run validation or discovery rules.
func countKGEAdvisories(root string, staged []string) int {
	ctx, err := buildKGEWorkspaceContext(root)
	if err != nil {
		return 0
	}
	if len(staged) > 0 {
		ctx.StagedPaths = make([]string, len(staged))
		for i, p := range staged {
			ctx.StagedPaths[i] = filepath.ToSlash(p)
		}
		contents := map[string]string{}
		for _, p := range ctx.StagedPaths {
			lower := strings.ToLower(p)
			if !strings.HasSuffix(lower, ".md") && !strings.HasSuffix(lower, ".markdown") {
				continue
			}
			body, readErr := os.ReadFile(filepath.Join(root, filepath.FromSlash(p)))
			if readErr != nil {
				continue
			}
			contents[p] = string(body)
		}
		ctx.FileContents = contents
		if ctx.Provided == nil {
			ctx.Provided = map[kge.CapabilityID]bool{}
		}
		if len(contents) > 0 {
			ctx.Provided[kge.CapStagedContent] = true
		} else {
			delete(ctx.Provided, kge.CapStagedContent)
		}
	}
	eng := kge.NewEngine(kge.AdvisoryRules()...)
	findings := eng.RunAvailable(ctx)
	_, adv, _ := kge.Partition(findings)
	return len(adv)
}

// attachKGEAdvisoryCount adds D9 commit-msg count-only warning (non-blocking).
func attachKGEAdvisoryCount(result Result, root string, staged []string) Result {
	n := countKGEAdvisories(root, staged)
	if n <= 0 {
		return result
	}
	result.Checks = append(result.Checks, Check{
		Name:    "kge_advisory",
		Status:  "warning",
		Message: kge.FormatCommitSummary(n, "ai-skill kge validate --advisory"),
	})
	return result
}

// attachKGECheck runs the full default pack with D9 check presentation.
// Blocks only on validation (error) findings.
func attachKGECheck(result Result, root string) Result {
	ctx, err := buildKGEWorkspaceContext(root)
	if err != nil {
		result.Checks = append(result.Checks, Check{
			Name:    "kge_check",
			Status:  "warning",
			Message: "kge check skipped: " + err.Error(),
		})
		return result
	}
	eng := kge.NewEngine(kge.DefaultRules()...)
	findings := eng.RunAvailable(ctx)
	report := kge.FormatCheckReport(findings, 3)
	if kge.Blocking(findings) {
		result.Status = "blocked"
		result.ExitCode = ExitValidationFailed
		result.Checks = append(result.Checks, Check{Name: "kge_check", Status: "error", Message: report})
		result.Error = &CommandError{
			Code:        "kge_check_failed",
			Message:     report,
			Remediation: "Fix validation findings, or inspect advisories with `ai-skill kge validate --advisory`.",
		}
		return result
	}
	status := "ok"
	if strings.Contains(report, "recommendation") {
		status = "warning"
	}
	result.Checks = append(result.Checks, Check{Name: "kge_check", Status: status, Message: report})
	return result
}
