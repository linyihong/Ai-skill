package app

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
)

const pushGovernanceReplayRemediation = `One or more commits being pushed likely bypassed local hook validation (e.g. git commit --no-verify).
Remediation:
  1. Identify failing commit(s) below and reset: git reset --soft @{u}
  2. Fix the reported issue(s)
  3. Re-commit WITHOUT --no-verify (hooks must run)
  4. git push (WITHOUT --no-verify)
Note: git push --no-verify also skips this pre-push replay; use normal git push. Server-side branch protection + CI remain the third line of defense.`

func validatePushGovernanceReplay(root string) (string, []Check) {
	commits, changed, rangeLabel, err := pushReplayScope(root)
	if err != nil {
		return "", []Check{{Name: "push_governance_replay", Status: "warning", Message: "could not resolve push range: " + err.Error()}}
	}
	if len(commits) == 0 {
		return "", []Check{{Name: "push_governance_replay", Status: "skipped", Message: "no commits to push (" + rangeLabel + ")"}}
	}

	checks := []Check{{Name: "push_governance_replay", Status: "ok", Message: fmt.Sprintf("replaying pre-commit + commit-msg checks for %d commit(s) (%s)", len(commits), rangeLabel)}}

	if msg := validateSanitizationAtHEAD(root, changed); msg != "" {
		return msg, checks
	}

	for _, sha := range commits {
		text, err := commitMessageBody(root, sha)
		if err != nil {
			return "push-governance-replay: could not read commit message for " + shortSHA(sha) + ": " + err.Error(), checks
		}
		files, err := commitChangedFiles(root, sha)
		if err != nil {
			return "push-governance-replay: could not list files for " + shortSHA(sha) + ": " + err.Error(), checks
		}
		added, err := filesAddedInCommit(root, sha)
		if err != nil {
			return "push-governance-replay: could not list added files for " + shortSHA(sha) + ": " + err.Error(), checks
		}
		if msg := validateNewShellScriptsInCommit(root, sha, text, added); msg != "" {
			return fmt.Sprintf("push-governance-replay (pre-commit shell policy, %s):\n%s", shortSHA(sha), msg), checks
		}
		if cmdErr := validateCommitMessageGovernance(root, text, files); cmdErr != nil {
			return fmt.Sprintf("push-governance-replay (commit-msg %s):\n%s", shortSHA(sha), cmdErr.Message), checks
		}
	}

	return "", checks
}

func pushReplayScope(root string) (commits []string, changed []string, rangeLabel string, err error) {
	upstreamOutput, upstreamErr := exec.Command("git", "-C", root, "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}").Output()
	if upstreamErr != nil {
		rangeLabel = "HEAD^..HEAD"
		changed, err = gitLines(root, "diff", "--name-only", "HEAD^...HEAD")
		if err != nil {
			return nil, nil, rangeLabel, err
		}
		commits, err = gitLines(root, "rev-list", "--reverse", rangeLabel)
		return commits, changed, rangeLabel, err
	}
	upstream := strings.TrimSpace(string(upstreamOutput))
	rangeLabel = upstream + "..HEAD"
	changed, err = gitLines(root, "diff", "--name-only", upstream+"...HEAD")
	if err != nil {
		return nil, nil, rangeLabel, err
	}
	commits, err = gitLines(root, "rev-list", "--reverse", rangeLabel)
	return commits, changed, rangeLabel, err
}

func commitMessageBody(root, sha string) (string, error) {
	out, err := exec.Command("git", "-C", root, "log", "-1", "--format=%B", sha).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func commitChangedFiles(root, sha string) ([]string, error) {
	return gitLines(root, "diff-tree", "--no-commit-id", "--name-only", "-r", sha)
}

func filesAddedInCommit(root, sha string) ([]string, error) {
	return gitLines(root, "diff-tree", "--no-commit-id", "--diff-filter=A", "--name-only", "-r", sha)
}

// validateNewShellScriptsInCommit replays the pre-commit Go-first gate for
// one pushed commit using the same portable rule and context builder, so the
// replay and pre-commit share one forbidden set (extensions + shebang).
func validateNewShellScriptsInCommit(root, sha, text string, added []string) string {
	return kgeFindingsMessage(kge.NewEngine(kge.NoNewShellScriptsRule{}).Run(noNewScriptsContext(root, sha, text, added)))
}

func shortSHA(sha string) string {
	if len(sha) > 8 {
		return sha[:8]
	}
	return sha
}
