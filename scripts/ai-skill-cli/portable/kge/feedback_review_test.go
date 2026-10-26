package kge

import (
	"strings"
	"testing"
)

const reviewLessonPath = "feedback/history/demo/common/2026-10-26_000000-review.md"
const reviewAtomPath = "intelligence/demo/decision.md"

func reviewFixture() (Context, string) {
	body := `Status: candidate
#### Transfer Contract
- Inputs: observed typed call and known-good request
- Mechanism: representation mismatch can hide valid data
- Decision: verify boundary before claiming absence
- Verification: compare known-good and offline rejected input
- Counterexample: complete validated contract requires no conversion
- Transfer case: unrelated configuration reader with typed keys
#### Intelligence Review
- Decision: extract
- Reason: changes a reusable diagnostic decision
- Target: intelligence/demo/decision.md
- Discovery: knowledge/summaries/decision.md
`
	atom := "---\nid: intelligence.demo.decision\ntype: intelligence\nlayer: intelligence\nstatus: candidate\n---\n[Source](../../" + reviewLessonPath + ")\n"
	ctx := Context{FileContents: map[string]string{reviewAtomPath: atom, "intelligence/demo/README.md": "[Decision](decision.md)", "knowledge/summaries/decision.md": "[Source](../../intelligence/demo/decision.md)"}, FeedbackRoutedTargets: map[string]bool{reviewAtomPath: true}, ExistingPaths: map[string]bool{reviewAtomPath: true, "intelligence/demo/README.md": true, "knowledge/summaries/decision.md": true}}
	return ctx, body
}
func TestFeedbackReviewContractMutations(t *testing.T) {
	ctx, body := reviewFixture()
	if got := feedbackReviewViolations(ctx, reviewLessonPath, body); len(got) != 0 {
		t.Fatal(got)
	}
	for _, tc := range []struct{ name, old, new string }{
		{"missing transfer", "#### Transfer Contract", "#### Old Contract"},
		{"empty mechanism", "Mechanism: representation mismatch can hide valid data", "Mechanism: TODO"},
		{"missing disposition", "Decision: extract", "Decision: "},
		{"invalid disposition", "Decision: extract", "Decision: promote-everything"},
		{"wrong target layer", "Target: intelligence/demo/decision.md", "Target: feedback/history/demo/decision.md"},
		{"traversal", "Target: intelligence/demo/decision.md", "Target: intelligence/../../private.md"},
		{"deferred needs revisit", "Decision: extract", "Decision: deferred"},
		{"project only", "Decision: extract", "Decision: project-only"},
		{"duplicate decision", "- Decision: extract", "- Decision: extract\n- Decision: existing"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := feedbackReviewViolations(ctx, reviewLessonPath, strings.Replace(body, tc.old, tc.new, 1)); len(got) == 0 {
				t.Fatal("mutation must block")
			}
		})
	}
	for _, name := range []string{"missing target", "missing backlink", "missing index", "missing metadata", "missing route", "missing discovery", "raised confidence", "unclosed metadata"} {
		t.Run(name, func(t *testing.T) {
			c, b := reviewFixture()
			switch name {
			case "missing target":
				delete(c.ExistingPaths, reviewAtomPath)
			case "missing backlink":
				c.FileContents[reviewAtomPath] = "---\nid: x\ntype: intelligence\nlayer: intelligence\nstatus: candidate\n---"
			case "missing index":
				c.FileContents["intelligence/demo/README.md"] = "# No links"
			case "missing route":
				delete(c.FeedbackRoutedTargets, reviewAtomPath)
			case "missing discovery":
				delete(c.ExistingPaths, "knowledge/summaries/decision.md")
			case "raised confidence":
				c.FileContents[reviewAtomPath] = strings.Replace(c.FileContents[reviewAtomPath], "status: candidate", "status: validated", 1)
			case "unclosed metadata":
				c.FileContents[reviewAtomPath] = strings.Replace(c.FileContents[reviewAtomPath], "\n---\n[Source]", "\n[Source]", 1)
			case "missing metadata":
				c.FileContents[reviewAtomPath] = "[Source](../../" + reviewLessonPath + ")"
			}
			if got := feedbackReviewViolations(c, reviewLessonPath, b); len(got) == 0 {
				t.Fatal("mutation must block")
			}
		})
	}
}
func TestFeedbackReviewDispositionAndEvidence(t *testing.T) {
	ctx, body := reviewFixture()
	deferred := strings.Replace(body, "Decision: extract", "Decision: deferred", 1)
	deferred = strings.Replace(deferred, "Target: "+reviewAtomPath, "Target: none\n- Revisit: when a second consumer has the same diagnostic choice", 1)
	if got := feedbackReviewViolations(ctx, reviewLessonPath, deferred); len(got) != 0 {
		t.Fatal(got)
	}
	method := strings.Replace(body, "Decision: extract", "Decision: method-only", 1)
	method = strings.Replace(method, "Target: "+reviewAtomPath, "Target: analysis/demo/method.md", 1)
	ctx.FileContents["analysis/demo/method.md"] = "[Source](../../" + reviewLessonPath + ")"
	ctx.FileContents["analysis/demo/README.md"] = "[Method](method.md)"
	ctx.ExistingPaths["analysis/demo/method.md"] = true
	ctx.ExistingPaths["analysis/demo/README.md"] = true
	if got := feedbackReviewViolations(ctx, reviewLessonPath, method); len(got) != 0 {
		t.Fatal(got)
	}
	existing := strings.Replace(body, "Decision: extract", "Decision: existing", 1)
	if got := feedbackReviewViolations(ctx, reviewLessonPath, existing); len(got) != 0 {
		t.Fatal(got)
	}
	ctx.AddedPaths = []string{reviewAtomPath}
	if got := feedbackReviewViolations(ctx, reviewLessonPath, existing); len(got) == 0 {
		t.Fatal("new atom cannot claim existing disposition")
	}
	ctx.AddedPaths = nil
	validated := strings.Replace(body, "Status: candidate", "Status: validated", 1) + "\n#### Reuse Evidence\n- Independent context: unrelated consumer owned separately\n- Result: rule verified in that consumer\n- Reference: evidence/reuse.md\n"
	if got := feedbackReviewViolations(ctx, reviewLessonPath, validated); len(got) == 0 {
		t.Fatal("unresolved evidence must block")
	}
	ctx.ExistingPaths["evidence/reuse.md"] = true
	ctx.FileContents["evidence/reuse.md"] = "Sanitized independent context and result"
	if got := feedbackReviewViolations(ctx, reviewLessonPath, validated); len(got) != 0 {
		t.Fatal(got)
	}
	if got := feedbackReviewViolations(ctx, reviewLessonPath, validated+"尚未有獨立重用證據"); len(got) == 0 {
		t.Fatal("contradictory reuse claim must block")
	}
}
