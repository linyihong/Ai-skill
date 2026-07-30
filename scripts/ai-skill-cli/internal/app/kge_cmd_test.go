package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunKgeCheck_DocumentSizingAdvisory(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	md := filepath.Join(root, "notes.md")
	var b strings.Builder
	for i := 0; i < 320; i++ {
		b.WriteString("line\n")
	}
	if err := os.WriteFile(md, []byte(b.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "notes.md")

	var stdout, stderr strings.Builder
	code := runKge([]string{"check", "--root", root}, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("check exit=%d stderr=%s stdout=%s", code, stderr.String(), stdout.String())
	}
	out := stdout.String()
	if !strings.Contains(out, "Ready to push") || !strings.Contains(out, "recommendation") {
		t.Fatalf("want advisory summary on check, got:\n%s", out)
	}
}

func TestRunKgeValidateAdvisory(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	md := filepath.Join(root, "notes.md")
	var b strings.Builder
	for i := 0; i < 160; i++ {
		b.WriteString("line\n")
	}
	if err := os.WriteFile(md, []byte(b.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "notes.md")

	var stdout, stderr strings.Builder
	code := runKge([]string{"validate", "--root", root, "--advisory"}, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("validate exit=%d stderr=%s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "document_sizing") && !strings.Contains(stdout.String(), "Advisory") {
		t.Fatalf("want advisory detail, got:\n%s", stdout.String())
	}
}

func TestRunKgeDiagnoseJSON(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	md := filepath.Join(root, "notes.md")
	var b strings.Builder
	for i := 0; i < 320; i++ {
		b.WriteString("line\n")
	}
	if err := os.WriteFile(md, []byte(b.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "notes.md")

	var stdout, stderr strings.Builder
	code := runKge([]string{"diagnose", "--root", root}, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("diagnose exit=%d stderr=%s", code, stderr.String())
	}
	out := stdout.String()
	if !strings.Contains(out, `"rule_id"`) || !strings.Contains(out, "document_sizing") {
		t.Fatalf("want IDE JSON diagnostics, got:\n%s", out)
	}
}

func TestCountKGEAdvisories(t *testing.T) {
	root := t.TempDir()
	runGit(t, root, "init")
	md := filepath.Join(root, "notes.md")
	var b strings.Builder
	for i := 0; i < 320; i++ {
		b.WriteString("line\n")
	}
	if err := os.WriteFile(md, []byte(b.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, root, "add", "notes.md")
	if n := countKGEAdvisories(root, []string{"notes.md"}); n != 1 {
		t.Fatalf("want 1 advisory, got %d", n)
	}
	result := attachKGEAdvisoryCount(Result{Checks: nil}, root, []string{"notes.md"})
	if len(result.Checks) != 1 || result.Checks[0].Name != "kge_advisory" {
		t.Fatalf("want kge_advisory warning check, got %#v", result.Checks)
	}
	if !strings.Contains(result.Checks[0].Message, "1 advisory") {
		t.Fatalf("want count-only summary, got %q", result.Checks[0].Message)
	}
	if strings.Contains(result.Checks[0].Message, "document sizing:") {
		t.Fatal("commit-msg must not expand advisory body")
	}
}
