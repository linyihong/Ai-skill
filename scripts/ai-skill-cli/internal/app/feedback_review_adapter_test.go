package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const stagedReviewLesson = "feedback/history/demo/common/2026-10-26_010101-example.md"
const stagedReviewTarget = "intelligence/demo/decision.md"

func stagedReviewFixture() string {
	return `Status: candidate
#### One-line Summary
Reusable boundary decision
#### Evidence
Sanitized prior observation
#### Generalized Lesson
Check the consumer contract first
#### Agent Action
Compare input and known boundary
#### Goal / Action / Validation
- Goal: avoid invalid attribution
#### Applies When
Typed boundaries
#### Does Not Apply When
No known contract
#### Validation
Compare valid and counterexample
#### Promotion Target
Intelligence target below
#### Required Linked Updates
Source, target and index updated
#### Abstraction Check
- Swap-project test: PASS — independent typed consumer can use the decision
- Project evidence kept in: not applicable — synthetic fixture
- Product/cabinet/sample IDs absent from Generalized Lesson: PASS
#### Transfer Contract
- Inputs: typed boundary and known request
- Mechanism: representation can change at the adapter
- Decision: inspect the contract before claiming absence
- Verification: compare valid input and rejected mismatch
- Counterexample: known direct consumer needs no expansion
- Transfer case: unrelated settings reader
#### Intelligence Review
- Decision: extract
- Reason: changes reusable diagnostic choice
- Target: intelligence/demo/decision.md
- Discovery: knowledge/summaries/decision.md
`
}
func TestFeedbackReviewUsesIndexSnapshot(t *testing.T) {
	for _, mode := range []string{"worktree lesson repair", "untracked target", "worktree backlink repair", "worktree routing repair", "worktree discovery repair", "indexed complete"} {
		t.Run(mode, func(t *testing.T) {
			root := t.TempDir()
			runGit(t, root, "init")
			body := stagedReviewFixture()
			target := "---\nid: intelligence.demo.decision\ntype: intelligence\nlayer: intelligence\nstatus: candidate\n---\n[Source](../../" + stagedReviewLesson + ")\n"
			writeFile(t, filepath.Join(root, stagedReviewLesson), body)
			writeFile(t, filepath.Join(root, "feedback/history/demo/common/README.md"), "[Lesson](2026-10-26_010101-example.md)")
			writeFile(t, filepath.Join(root, stagedReviewTarget), target)
			writeFile(t, filepath.Join(root, "intelligence/demo/README.md"), "[Decision](decision.md)")
			if mode == "worktree lesson repair" {
				writeFile(t, filepath.Join(root, stagedReviewLesson), strings.Split(body, "#### Transfer Contract")[0])
			}
			if mode == "worktree backlink repair" {
				writeFile(t, filepath.Join(root, stagedReviewTarget), "---\nid: demo\ntype: intelligence\nlayer: intelligence\nstatus: candidate\n---")
			}
			writeFile(t, filepath.Join(root, "knowledge/summaries/decision.md"), "[Source](../../intelligence/demo/decision.md)")
			writeFile(t, filepath.Join(root, "knowledge/runtime/routing-registry.yaml"), "records:\n  - id: route.intelligence.demo\n    primary_source: intelligence/demo/decision.md\n")
			if mode == "worktree routing repair" {
				writeFile(t, filepath.Join(root, "knowledge/runtime/routing-registry.yaml"), "records: []")
			}
			if mode == "worktree discovery repair" {
				writeFile(t, filepath.Join(root, "knowledge/summaries/decision.md"), "[Wrong](../../intelligence/demo/other.md)")
			}
			runGit(t, root, "add", "knowledge")
			writeFile(t, filepath.Join(root, "knowledge/summaries/decision.md"), "[Source](../../intelligence/demo/decision.md)")
			writeFile(t, filepath.Join(root, "knowledge/runtime/routing-registry.yaml"), "records:\n  - id: route.intelligence.demo\n    primary_source: intelligence/demo/decision.md\n")

			runGit(t, root, "add", "feedback")
			if mode != "untracked target" {
				runGit(t, root, "add", "intelligence")
			}
			writeFile(t, filepath.Join(root, stagedReviewLesson), body)
			writeFile(t, filepath.Join(root, stagedReviewTarget), target)
			if mode == "indexed complete" {
				if err := os.Remove(filepath.Join(root, stagedReviewLesson)); err != nil {
					t.Fatal(err)
				}
			}
			paths, _ := gitLines(root, "diff", "--cached", "--name-only")
			msg := runKGEFeedbackLessonClosure("feedback: fixture", paths, root)
			ctx, err := buildKGEWorkspaceContext(root, false)
			if err != nil {
				t.Fatal(err)
			}
			if mode == "indexed complete" {
				if msg != "" {
					t.Fatal(msg)
				}
				if ctx.FileContents[stagedReviewLesson] != body {
					t.Fatal("CLI must use indexed lesson even when worktree file is absent")
				}
			} else if msg == "" {
				t.Fatal("unstaged repair must not satisfy commit gate")
			}
			if mode == "worktree lesson repair" && strings.Contains(ctx.FileContents[stagedReviewLesson], "#### Transfer Contract") {
				t.Fatal("CLI used unstaged repair")
			}
			if mode == "untracked target" && ctx.ExistingPaths[stagedReviewTarget] {
				t.Fatal("CLI attested untracked target")
			}
		})
	}
}
