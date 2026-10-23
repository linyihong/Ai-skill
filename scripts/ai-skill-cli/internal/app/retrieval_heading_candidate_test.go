package app

import "testing"

func TestRetrievalHeadingCandidate(t *testing.T) {
	for _, tc := range []struct{ in, title, summary string }{
		{"# Title\n## One-line Summary\n\nfirst\ncontinued\n\nbody", "Title", "first continued"},
		{"### Legacy\n#### One-line Summary\nvalue", "Legacy", "value"},
		{"# Title\n## One-line Summary\n\n## Evidence\nwrong", "Title", ""},
		{"```md\n# Fake\n```\n## Real\n### One-line Summary\nvalid", "Real", "valid"},
	} {
		title, summary := runtimeFeedbackHeadingFields(tc.in)
		if title != tc.title || summary != tc.summary {
			t.Errorf("got %q / %q, want %q / %q", title, summary, tc.title, tc.summary)
		}
	}
}
