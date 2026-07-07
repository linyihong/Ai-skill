package app

import "testing"

func TestNativeCanonicalOwnershipDriftValidationOK(t *testing.T) {
	repo := repoRootForTest(t)
	check := nativeCanonicalOwnershipDriftValidation(repo)
	if check.Status != "ok" {
		t.Fatalf("expected debt payoff canon ok, got %#v", check)
	}
}

func TestScanCanonicalOwnershipDriftValidationOwnsReport(t *testing.T) {
	content := "> **輸出模板**：Validate 完成後，使用 review-report-template.md 記錄審查報告。\n"
	v := scanCanonicalOwnershipDrift(content, "workflow/software-delivery/validation.md")
	if len(v) == 0 {
		t.Fatal("expected ownership drift for validation owning review report")
	}
}

func TestScanCanonicalOwnershipDriftPathReviewChecklist(t *testing.T) {
	content := "See workflow/software-delivery/review-checklist.md for details.\n"
	v := scanCanonicalOwnershipDrift(content, "workflow/software-delivery/intake.md")
	if len(v) == 0 {
		t.Fatal("expected path drift for review-checklist canonical path")
	}
}

func TestScanCanonicalOwnershipDriftPathStubAllowed(t *testing.T) {
	content := "舊路徑 review-checklist.md 為 stub redirect。\n"
	v := scanCanonicalOwnershipDrift(content, "workflow/software-delivery/README.md")
	if len(v) != 0 {
		t.Fatalf("expected stub redirect allowed, got %#v", v)
	}
}

func TestScanCanonicalOwnershipDriftStanceEnumInREADME(t *testing.T) {
	content := "stance_enum:\n  standardized:\n    - fault_finding\n"
	v := scanCanonicalOwnershipDrift(content, "workflow/software-delivery/README.md")
	if len(v) == 0 {
		t.Fatal("expected stance enum ownership drift in README")
	}
}
