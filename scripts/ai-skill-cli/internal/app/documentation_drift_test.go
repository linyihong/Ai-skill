package app

import (
	"strings"
	"testing"
)

func TestScanReviewArchitectureNavigationDriftForbidden(t *testing.T) {
	content := "Add a Review Phase after implementation.\n"
	v := scanReviewArchitectureNavigationDrift(content, "workflow/software-delivery/README.md")
	if len(v) != 1 || !strings.Contains(v[0], "review_phase") {
		t.Fatalf("expected review_phase violation, got %#v", v)
	}
}

func TestScanReviewArchitectureNavigationDriftNegationAllowed(t *testing.T) {
	content := "不新增 sd-review lifecycle slice.\n"
	v := scanReviewArchitectureNavigationDrift(content, "governance/cognitive-slice-taxonomy.md")
	if len(v) != 0 {
		t.Fatalf("expected no violations for negation, got %#v", v)
	}
}

func TestNativeReviewArchitectureDocDriftValidationCanonOK(t *testing.T) {
	repo := repoRootForTest(t)
	check := nativeReviewArchitectureDocDriftValidation(repo)
	if check.Status != "ok" {
		t.Fatalf("expected canon navigation ok, got %#v", check)
	}
}

func TestDocumentationRegressionFourCaseDriftPatterns(t *testing.T) {
	cases := []struct {
		name    string
		content string
		want    int
	}{
		{"review_phase", "enter the review phase", 1},
		{"review_workflow", "load the review workflow", 1},
		{"review_slice", "create a review slice", 1},
		{"sd_review", "promote sd-review", 1},
		{"invoke_ok", "invoke code-review capability with fault_finding", 0},
		{"negation_ok", "Reject sd-review as primary model", 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := scanReviewArchitectureNavigationDrift(tc.content, "test.md")
			if len(v) != tc.want {
				t.Fatalf("got %d violations, want %d: %#v", len(v), tc.want, v)
			}
		})
	}
}
