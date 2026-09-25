package kge

import (
	"path/filepath"
	"strings"
)

// FeedbackLessonClosureRule keeps feedback/history entries usable as durable
// lessons. New lessons need an index entry; existing lessons are checked again
// whenever they are touched, which provides an incremental repair path.
//
// Opt-out: standalone [skip-feedback-lesson-closure] trailer in CommitMsg.
type FeedbackLessonClosureRule struct{}

func (FeedbackLessonClosureRule) ID() string { return "rule.feedback_lesson_closure" }
func (FeedbackLessonClosureRule) Kind() Kind { return KindValidation }

func (FeedbackLessonClosureRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapStagedPaths, CapStagedContent, CapAddedPaths}
}

func (FeedbackLessonClosureRule) Validate(ctx Context) []Finding {
	if hasStandaloneTrailer(ctx.CommitMsg, "[skip-feedback-lesson-closure]") {
		return nil
	}
	added := make(map[string]bool, len(ctx.AddedPaths))
	for _, p := range ctx.AddedPaths {
		added[filepath.ToSlash(p)] = true
	}
	staged := make(map[string]bool, len(ctx.StagedPaths))
	for _, p := range ctx.StagedPaths {
		staged[filepath.ToSlash(p)] = true
	}

	var out []Finding
	for _, raw := range ctx.StagedPaths {
		path := filepath.ToSlash(raw)
		if !IsFeedbackLessonPath(path) {
			continue
		}
		body, ok := ctx.FileContents[path]
		if !ok { // deletion or unavailable staged blob
			continue
		}
		missing := missingLessonSections(body)
		if len(missing) > 0 {
			out = append(out, feedbackClosureFinding(path, "missing required section(s): "+strings.Join(missing, ", ")))
			continue
		}
		status := lessonStatus(body)
		if !validLessonStatus(status) {
			out = append(out, feedbackClosureFinding(path, "Status must be experimental, candidate, validated, promoted, or deprecated"))
			continue
		}
		if sectionIsEmpty(body, "Required Linked Updates") {
			out = append(out, feedbackClosureFinding(path, "Required Linked Updates must state completed updates or an explicit not-applicable reason"))
		}
		if (status == "validated" || status == "promoted") && sectionIsEmpty(body, "Reuse Evidence") {
			out = append(out, feedbackClosureFinding(path, "Status "+status+" requires Reuse Evidence; otherwise keep the lesson candidate"))
		}
		if status == "promoted" && sectionIsEmpty(body, "Promotion Record") {
			out = append(out, feedbackClosureFinding(path, "Status promoted requires Promotion Record with target path and validation/commit reference"))
		}
		if added[path] {
			index := filepath.ToSlash(filepath.Join(filepath.Dir(path), "README.md"))
			if !staged[index] {
				out = append(out, feedbackClosureFinding(path, "new lesson requires its category README index in the same commit: "+index))
			}
		}
	}
	return out
}

func IsFeedbackLessonPath(path string) bool {
	if !strings.HasPrefix(path, "feedback/history/") || !strings.HasSuffix(path, ".md") {
		return false
	}
	if strings.HasSuffix(path, "/README.md") {
		return false
	}
	return len(strings.Split(path, "/")) >= 5
}

func missingLessonSections(body string) []string {
	required := []string{"One-line Summary", "Evidence", "Generalized Lesson", "Agent Action", "Goal / Action / Validation", "Applies / Does Not Apply", "Validation", "Promotion Target", "Required Linked Updates"}
	var missing []string
	for _, heading := range required {
		if !hasHeading(body, heading) {
			missing = append(missing, heading)
		}
	}
	return missing
}

func lessonStatus(body string) string {
	for _, line := range strings.Split(body, "\n") {
		line = strings.TrimSpace(line)
		if len(line) >= len("Status:") && strings.EqualFold(line[:len("Status:")], "Status:") {
			return strings.ToLower(strings.TrimSpace(line[len("Status:"):]))
		}
	}
	return ""
}

func validLessonStatus(status string) bool {
	switch status {
	case "experimental", "candidate", "validated", "promoted", "deprecated":
		return true
	default:
		return false
	}
}

func hasHeading(body, heading string) bool {
	if heading == "Applies / Does Not Apply" {
		return strings.Contains(body, "#### Applies / Does Not Apply") ||
			(strings.Contains(body, "#### Applies When") && strings.Contains(body, "#### Does Not Apply When"))
	}
	return strings.Contains(body, "#### "+heading) || strings.Contains(body, "### "+heading) || strings.Contains(body, "## "+heading)
}

func sectionIsEmpty(body, heading string) bool {
	idx, needleLen := -1, 0
	for _, prefix := range []string{"#### ", "### ", "## "} {
		if pos := strings.Index(body, prefix+heading); pos >= 0 && (idx < 0 || pos < idx) {
			idx, needleLen = pos, len(prefix)+len(heading)
		}
	}
	if idx < 0 {
		return true
	}
	rest := body[idx+needleLen:]
	if next := strings.Index(rest, "\n#"); next >= 0 {
		rest = rest[:next]
	}
	return strings.TrimSpace(rest) == ""
}

func hasStandaloneTrailer(text, marker string) bool {
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == marker {
			return true
		}
	}
	return false
}

func feedbackClosureFinding(path, detail string) Finding {
	return Finding{RuleID: "rule.feedback_lesson_closure", Severity: SeverityError, Code: "feedback_lesson_closure", Path: path, Message: "feedback-lesson-closure: " + path + " " + detail + ". See feedback/feedback-lessons.md; use [skip-feedback-lesson-closure] only for an intentional retrospective exception."}
}
