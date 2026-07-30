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
