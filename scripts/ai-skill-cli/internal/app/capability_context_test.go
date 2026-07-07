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

func TestValidateCapabilityInvokeNoStanceRequiredPass(t *testing.T) {
	repo := t.TempDir()
	registryYAML := `registry_version: capability-registry/v1
status: candidate
owner_layer: knowledge/runtime
canonical_contract: governance/cognitive-stance.md
schema: metadata/capability-context-schema.md
stance_enum:
  standardized:
    - fault_finding
  implicit:
    - default
capabilities:
  - id: docs-only-edit
    status: active
    summary: Documentation wording fix; default stance suffices.
    artifact: none
`
	if err := os.MkdirAll(filepath.Join(repo, "knowledge", "runtime"), 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(repo, "knowledge", "runtime", "capability-registry.yaml"), registryYAML)

	registry, err := readCapabilityRegistry(repo)
	if err != nil {
		t.Fatal(err)
	}
	for _, stance := range []string{"", "default", "fault_finding"} {
		v := validateCapabilityInvoke(registry, "docs-only-edit", stance)
		if len(v.Warnings) != 0 {
			t.Fatalf("stance=%q: expected no warnings for capability without requires_context, got %#v", stance, v)
		}
	}
}

func TestCapabilityContextContractRegressionCases(t *testing.T) {
	repo := repoRootForTest(t)
	registry, err := readCapabilityRegistry(repo)
	if err != nil {
		t.Fatal(err)
	}

	cases := []struct {
		name     string
		capID    string
		stance   string
		wantWarn int
	}{
		{
			name:     "case1_required_stance_provided",
			capID:    "code-review",
			stance:   "fault_finding",
			wantWarn: 0,
		},
		{
			name:     "case2_required_stance_missing",
			capID:    "code-review",
			stance:   "",
			wantWarn: 1,
		},
		{
			name:     "case3_no_stance_required_empty",
			capID:    "docs-only-edit",
			stance:   "",
			wantWarn: 0,
		},
		{
			name:     "case4_invoke_stance_mismatch",
			capID:    "code-review",
			stance:   "creative",
			wantWarn: 1,
		},
	}

	// case3 uses synthetic capability without requires_context
	synthetic := capabilityRegistryEntry{
		ID:      "docs-only-edit",
		Status:  "active",
		Summary: "synthetic for regression",
	}
	registry.Capabilities = append(registry.Capabilities, synthetic)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := validateCapabilityInvoke(registry, tc.capID, tc.stance)
			if len(v.Warnings) != tc.wantWarn {
				t.Fatalf("got warnings=%d want=%d: %#v", len(v.Warnings), tc.wantWarn, v)
			}
		})
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
