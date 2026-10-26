package app

import (
	"gopkg.in/yaml.v3"
	"os/exec"
	"path"
	"strings"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
)

func feedbackIndexContent(root, path string) (string, error) {
	b, err := exec.Command("git", "-C", root, "show", ":"+path).Output()
	return string(b), err
}

// Targets/evidence must be tracked in the effective commit tree. Worktree-only
// repairs and untracked candidate atoms cannot satisfy a staged lesson's claim.
func loadFeedbackReviewReferences(root string, contents map[string]string, existing map[string]bool) {
	refs := map[string]bool{}
	for p, body := range contents {
		if !kge.IsFeedbackLessonPath(p) {
			continue
		}
		refs[path.Join(path.Dir(p), "README.md")] = true
		for _, ref := range kge.FeedbackReviewReferences(body) {
			refs[ref] = true
		}
	}
	for ref := range refs {
		body, err := feedbackIndexContent(root, ref)
		if err != nil {
			delete(existing, ref)
			delete(contents, ref)
			continue
		}
		contents[ref] = body
		existing[ref] = strings.TrimSpace(body) != ""
	}
}

// Only parsed registry records attest routing membership; comments do not.
func feedbackRoutedTargets(root string) map[string]bool {
	targets := map[string]bool{}
	body, err := feedbackIndexContent(root, "knowledge/runtime/routing-registry.yaml")
	if err != nil {
		return targets
	}
	var registry runtimeRoutingRegistry
	if yaml.Unmarshal([]byte(body), &registry) != nil {
		return targets
	}
	for _, r := range registry.Records {
		targets[r.PrimarySource] = true
		for _, p := range r.CandidateSources {
			targets[p] = true
		}
		for _, p := range r.RequiredDependencies {
			targets[p] = true
		}
	}
	return targets
}
