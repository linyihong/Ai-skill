package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidatePlanEvidenceConvention_MissingReadme(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+"/_plan.md", "# main\n")
	writePlanFile(t, tmp, planDir+"/evidence/run-a.md", "# run\n")
	got := validatePlanEvidenceConvention("", []string{planDir + "/evidence/run-a.md"}, tmp)
	if got == "" || !strings.Contains(got, "missing evidence/README.md") {
		t.Fatalf("want missing README block, got: %q", got)
	}
}

func TestValidatePlanEvidenceConvention_UnindexedFile(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+"/_plan.md", "# main\n")
	writePlanFile(t, tmp, planDir+"/evidence/README.md", "# Index\n\n## 引用規則\n\nx\n\n## Run 索引\n\n| a | b |\n")
	writePlanFile(t, tmp, planDir+"/evidence/run-a.md", "# run\n")
	got := validatePlanEvidenceConvention("", []string{planDir + "/evidence/run-a.md"}, tmp)
	if got == "" || !strings.Contains(got, "run-a.md not listed") {
		t.Fatalf("want unindexed violation, got: %q", got)
	}
}

func TestValidatePlanEvidenceConvention_Pass(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+"/_plan.md", "# main\n")
	writePlanFile(t, tmp, planDir+"/evidence/README.md", "# Index\n\n## 引用規則\n\nx\n\n## Run 索引\n\n| Run | 檔案 |\n| 2d′ | [run-a.md](run-a.md) |\n")
	writePlanFile(t, tmp, planDir+"/evidence/run-a.md", "# run\n")
	got := validatePlanEvidenceConvention("", []string{planDir + "/evidence/run-a.md"}, tmp)
	if got != "" {
		t.Fatalf("want pass, got: %q", got)
	}
}

func TestValidatePlanEvidenceConvention_MissingPlanMain(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+"/evidence/README.md", "# Index\n\n## 引用規則\n\nx\n\n## Run 索引\n\n| Run | 檔案 |\n| a | [run-a.md](run-a.md) |\n")
	writePlanFile(t, tmp, planDir+"/evidence/run-a.md", "# run\n")
	got := validatePlanEvidenceConvention("", []string{planDir + "/evidence/run-a.md"}, tmp)
	if got == "" || !strings.Contains(got, "missing _plan.md") {
		t.Fatalf("want missing _plan.md block, got: %q", got)
	}
}

func TestValidatePlanEvidenceConvention_FlatSiblingBlocked(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+"/_plan.md", "# main in folder\n")
	writePlanFile(t, tmp, planDir+".md", "# flat sibling still here\n")
	writePlanFile(t, tmp, planDir+"/evidence/README.md", "# Index\n\n## 引用規則\n\nx\n\n## Run 索引\n\n| Run | 檔案 |\n| a | [run-a.md](run-a.md) |\n")
	writePlanFile(t, tmp, planDir+"/evidence/run-a.md", "# run\n")
	got := validatePlanEvidenceConvention("", []string{planDir + "/evidence/run-a.md"}, tmp)
	if got == "" || !strings.Contains(got, "flat sibling") {
		t.Fatalf("want flat sibling block, got: %q", got)
	}
}

func TestValidatePlanEvidenceConvention_StagingFlatMainWithEvidenceDir(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+".md", "# still flat\n")
	writePlanFile(t, tmp, planDir+"/evidence/README.md", "# Index\n\n## 引用規則\n\nx\n\n## Run 索引\n\n| Run | 檔案 |\n| a | [run-a.md](run-a.md) |\n")
	writePlanFile(t, tmp, planDir+"/evidence/run-a.md", "# run\n")
	got := validatePlanEvidenceConvention("", []string{planDir + ".md"}, tmp)
	if got == "" || !strings.Contains(got, "missing _plan.md") || !strings.Contains(got, "flat sibling") {
		t.Fatalf("want block when staging flat main while evidence/ exists, got: %q", got)
	}
}

func TestValidatePlanEvidenceConvention_SkipTrailer(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+"/evidence/run-a.md", "# run\n")
	got := validatePlanEvidenceConvention("fix\n\n[skip-plan-evidence]\n", []string{planDir + "/evidence/run-a.md"}, tmp)
	if got != "" {
		t.Fatalf("want skip, got: %q", got)
	}
}

func TestWarnPlanEvidenceLineNumberCitations(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	ev := filepath.Join(tmp, planDir, "evidence")
	if err := os.MkdirAll(ev, 0o755); err != nil {
		t.Fatal(err)
	}
	writePlanFile(t, tmp, planDir+"/evidence/README.md", "## 引用規則\n\n## Run 索引\n")
	writePlanFile(t, tmp, planDir+"/_plan.md", "see kit L449 for details\n")
	w := warnPlanEvidenceLineNumberCitations("", []string{planDir + "/_plan.md"}, tmp)
	if w == "" || !strings.Contains(w, "line-number citation") {
		t.Fatalf("want line citation warning, got: %q", w)
	}
}

func TestValidatePlanTreeFolderConvention_EvidencePathExempt(t *testing.T) {
	tmp := t.TempDir()
	planDir := "plans/active/2026-07-08-test-plan"
	writePlanFile(t, tmp, planDir+"/evidence/my-run-record.md", "# evidence\n")
	got := validatePlanTreeFolderConvention("", []string{planDir + "/evidence/my-run-record.md"}, tmp)
	if strings.Contains(got, "filename should be") {
		t.Fatalf("evidence file should be exempt from NN- prefix, got: %s", got)
	}
	if strings.Contains(got, "nested depth") {
		t.Fatalf("evidence path should be exempt from depth warning, got: %s", got)
	}
}
