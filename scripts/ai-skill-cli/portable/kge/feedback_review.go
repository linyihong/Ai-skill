package kge

import (
	"path"
	"strings"
)

// FeedbackTransferFields is shared with canonical-schema parity tests.
func FeedbackTransferFields() []string {
	return []string{"inputs", "mechanism", "decision", "verification", "counterexample", "transfer case"}
}
func FeedbackIntelligenceDecisions() []string {
	return []string{"extract", "existing", "method-only", "deferred"}
}

// FeedbackReviewReferences lists the canonical repo-relative files needed to
// check a lesson. Adapters attest these from the index, never untracked files.
func FeedbackReviewReferences(body string) []string {
	fields := reviewFields(sectionBody(body, "Intelligence Review"))
	refs := []string{}
	if p := fields["target"]; canonicalReviewPath(p) {
		refs = append(refs, p, path.Join(path.Dir(p), "README.md"))
	}
	if p := fields["discovery"]; canonicalReviewPath(p) {
		refs = append(refs, p)
	}
	reuse := reviewFields(sectionBody(body, "Reuse Evidence"))
	if p := reuse["reference"]; canonicalReviewPath(p) {
		refs = append(refs, p)
	}
	return refs
}

func reviewFields(section string) map[string]string {
	fields := map[string]string{}
	for _, line := range strings.Split(section, "\n") {
		line = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "- "))
		if key, value, ok := strings.Cut(line, ":"); ok {
			key = strings.ToLower(strings.TrimSpace(key))
			if _, exists := fields[key]; exists {
				fields[key] = ""
				continue
			}
			fields[key] = strings.Trim(strings.TrimSpace(value), "`")
		}
	}
	return fields
}
func reviewValue(value string) bool {
	v := strings.ToLower(strings.TrimSpace(value))
	switch v {
	case "", "n/a", "none", "unknown", "pending", "todo", "tbd", "pass", "fail", "...", "待補", "待確認", "未定", "later", "以後再說", "之後再說":
		return false
	}
	return !(strings.Contains(v, "<") && strings.Contains(v, ">")) && !strings.HasPrefix(v, "todo") && !strings.HasPrefix(v, "tbd")
}
func canonicalReviewPath(p string) bool {
	return strings.HasSuffix(p, ".md") && !strings.HasPrefix(p, "/") && !strings.Contains(p, "\\") && !strings.Contains(p, ":") && path.Clean(p) == p && !strings.HasPrefix(p, "../") && !strings.Contains(p, "<")
}
func markdownRefersTo(from, body, target string) bool {
	for _, l := range ExtractMarkdownLinks([]byte(body)) {
		raw := strings.Split(l.Target, "#")[0]
		if path.Clean(path.Join(path.Dir(from), raw)) == target {
			return true
		}
	}
	return false
}
func feedbackReviewViolations(ctx Context, lesson, body string) []string {
	var errors []string
	transfer := reviewFields(sectionBody(body, "Transfer Contract"))
	for _, key := range FeedbackTransferFields() {
		if !reviewValue(transfer[key]) {
			errors = append(errors, "Transfer Contract requires concrete "+key)
		}
	}
	review := reviewFields(sectionBody(body, "Intelligence Review"))
	decision := review["decision"]
	if !reviewValue(review["reason"]) {
		errors = append(errors, "Intelligence Review requires a concrete Reason")
	}
	target := review["target"]
	switch decision {
	case "extract", "existing", "method-only":
		prefixes := []string{"intelligence/"}
		if decision == "method-only" {
			prefixes = []string{"analysis/", "workflow/", "enforcement/"}
		}
		allowed := false
		for _, prefix := range prefixes {
			if strings.HasPrefix(target, prefix) {
				allowed = true
			}
		}
		if !allowed || !canonicalReviewPath(target) || strings.HasSuffix(target, "/README.md") {
			errors = append(errors, "Intelligence Review "+decision+" requires a canonical target file in the selected layer")
			break
		}
		targetBody, exists := ctx.FileContents[target]
		if !exists || !ctx.ExistingPaths[target] {
			errors = append(errors, "Intelligence Review target is missing from the index: "+target)
			break
		}
		if !markdownRefersTo(target, targetBody, lesson) {
			errors = append(errors, "Intelligence Review target must link back to the source lesson: "+target)
		}
		index := path.Join(path.Dir(target), "README.md")
		if indexBody, ok := ctx.FileContents[index]; !ok || !ctx.ExistingPaths[index] || !markdownRefersTo(index, indexBody, target) {
			errors = append(errors, "Intelligence Review target requires a resolving parent README index: "+index)
		}
		if decision == "existing" {
			for _, added := range ctx.AddedPaths {
				if added == target {
					errors = append(errors, "existing intelligence target cannot be newly added; choose extract")
				}
			}
		}
		if decision != "method-only" && !ctx.FeedbackRoutedTargets[target] {
			errors = append(errors, "Intelligence Review target is not referenced by an indexed routing record: "+target)
		}
		if decision == "extract" {
			discovery := review["discovery"]
			summary, ok := ctx.FileContents[discovery]
			if !canonicalReviewPath(discovery) || !strings.HasPrefix(discovery, "knowledge/summaries/") || !ctx.ExistingPaths[discovery] || !ok || !markdownRefersTo(discovery, summary, target) {
				errors = append(errors, "extracted intelligence requires indexed Discovery summary linking to its target")
			}
			parts := strings.SplitN(strings.TrimPrefix(targetBody, "---\n"), "\n---", 2)
			metadata := reviewFields(parts[0])
			if !strings.HasPrefix(targetBody, "---\n") || len(parts) != 2 || metadata["type"] != "intelligence" || metadata["layer"] != "intelligence" || metadata["id"] == "" || !validAtomReviewStatus(metadata["status"]) {
				errors = append(errors, "extracted intelligence requires id/type/layer/status frontmatter")
			}
			if (lessonStatus(body) == "candidate" || lessonStatus(body) == "experimental") && metadata["status"] != "candidate" {
				errors = append(errors, "candidate lesson extraction cannot raise atom confidence above candidate")
			}
		}
	case "deferred":
		if target != "none" {
			errors = append(errors, "deferred review must use Target: none")
		}
		if !reviewValue(review["revisit"]) {
			errors = append(errors, "deferred review requires an actionable Revisit condition")
		}
	case "project-only":
		errors = append(errors, "project-only findings belong in project docs, not reusable feedback/history")
	default:
		errors = append(errors, "Intelligence Review Decision must be extract, existing, method-only, or deferred")
	}
	status := lessonStatus(body)
	if status == "validated" || status == "promoted" {
		reuse := reviewFields(sectionBody(body, "Reuse Evidence"))
		for _, key := range []string{"independent context", "result"} {
			if !reviewValue(reuse[key]) {
				errors = append(errors, "Reuse Evidence requires concrete "+key+"; otherwise keep candidate")
			}
		}
		ref := reuse["reference"]
		if !canonicalReviewPath(ref) || ref == lesson || !ctx.ExistingPaths[ref] || strings.TrimSpace(ctx.FileContents[ref]) == "" {
			errors = append(errors, "Reuse Evidence Reference must resolve to a separate indexed evidence record")
		}
		lower := strings.ToLower(sectionBody(body, "Reuse Evidence"))
		for _, denial := range []string{"尚未有獨立", "沒有獨立", "未有獨立", "尚未實測", "no independent", "not independently", "same project only"} {
			if strings.Contains(lower, denial) {
				errors = append(errors, "Reuse Evidence contradicts validated/promoted status; keep candidate")
				break
			}
		}
	}
	return errors
}
func validAtomReviewStatus(s string) bool {
	return s == "candidate" || s == "validated" || s == "stable"
}
