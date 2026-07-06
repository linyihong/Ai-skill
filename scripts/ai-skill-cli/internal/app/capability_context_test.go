package app

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateCapabilityRegistryOK(t *testing.T) {
	repo := repoRootForTest(t)
	check := nativeCapabilityRegistryValidation(repo)
	if check.Status != "ok" {
		t.Fatalf("expected ok registry, got %#v", check)
	}
}

func TestValidateCapabilityRegistryRejectsImplicitStance(t *testing.T) {
	repo := t.TempDir()
	src := repoRootForTest(t)
	orig := filepath.Join(src, "knowledge", "runtime", "capability-registry.yaml")
	content, err := os.ReadFile(orig)
	if err != nil {
		t.Fatal(err)
	}
	bad := strings.Replace(string(content),
		"requires_context:\n      stance:\n        - fault_finding\n    artifact: review-report",
		"requires_context:\n      stance:\n        - default\n    artifact: review-report", 1)
	if err := os.MkdirAll(filepath.Join(repo, "knowledge", "runtime"), 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(repo, "knowledge", "runtime", "capability-registry.yaml"), bad)

	check := nativeCapabilityRegistryValidation(repo)
	if check.Status != "failed" || !strings.Contains(check.Message, "implicit") {
		t.Fatalf("expected implicit stance rejection, got %#v", check)
	}
}

func TestValidateCapabilityInvokeMissingStanceWarning(t *testing.T) {
	repo := repoRootForTest(t)
	registry, err := readCapabilityRegistry(repo)
	if err != nil {
		t.Fatal(err)
	}
	v := validateCapabilityInvoke(registry, "code-review", "")
	if len(v.Warnings) != 1 || !strings.Contains(v.Warnings[0], "fault_finding") {
		t.Fatalf("expected missing stance warning, got %#v", v)
	}
}

func TestValidateCapabilityInvokeMismatchWarning(t *testing.T) {
	repo := repoRootForTest(t)
	registry, err := readCapabilityRegistry(repo)
	if err != nil {
		t.Fatal(err)
	}
	v := validateCapabilityInvoke(registry, "code-review", "creative")
	if len(v.Warnings) != 1 || !strings.Contains(v.Warnings[0], "requires context.stance") {
		t.Fatalf("expected mismatch warning, got %#v", v)
	}
}

func TestValidateCapabilityInvokeOK(t *testing.T) {
	repo := repoRootForTest(t)
	registry, err := readCapabilityRegistry(repo)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"code-review", "security-audit", "incident-analysis"} {
		v := validateCapabilityInvoke(registry, id, "fault_finding")
		if len(v.Warnings) != 0 {
			t.Fatalf("%s: expected no warnings, got %#v", id, v)
		}
	}
}

func TestRuntimeCapabilityInvokeCLIWarningsExitZero(t *testing.T) {
	repo := repoRootForTest(t)
	var stdout, stderr bytes.Buffer
	code := Run([]string{
		"runtime", "capability-invoke",
		"--repo", repo,
		"--capability", "code-review",
	}, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("expected exit 0 on warning, got %d stderr=%s stdout=%s", code, stderr.String(), stdout.String())
	}
	if !strings.Contains(stderr.String(), "warning:") {
		t.Fatalf("expected warning on stderr, got stderr=%q stdout=%q", stderr.String(), stdout.String())
	}
}

func TestRuntimeCapabilityInvokeCLIOK(t *testing.T) {
	repo := repoRootForTest(t)
	var stdout, stderr bytes.Buffer
	code := Run([]string{
		"runtime", "capability-invoke",
		"--repo", repo,
		"--capability", "security-audit",
		"--stance", "fault_finding",
	}, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("expected success, got %d stderr=%s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "capability invoke ok") {
		t.Fatalf("expected ok message, got stdout=%q stderr=%q", stdout.String(), stderr.String())
	}
}
