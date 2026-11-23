package app

import (
	"os"
	"path/filepath"
	"testing"
)

// Executor-direction samples for orphan_executor hook-call coverage (F1 of
// plans/active/2026-11-23-0834-project-tooling-governance). Violation
// variants must be caught regardless of naming; legal look-alikes must not.

const hookCoverageRegistryHeader = `schema_version: 2
id: enforcement.enforcement-registry
enforcement_mode:
  orphan_rule: fail
  orphan_executor: fail
executor_kind_spec:
  binding_required_for: []
`

const hookCoverageHooksSrc = `package app

func runPreCommitHook() string {
	if msg := checkFoo(); msg != "" {
		return msg
	}
	return helperOnly()
}

func runCommitMsgHook() string { return "" }

func runPrePushHook() string { return "" }

var commitMsgValidatorRegistry = map[string]func() string{
	"obligation.commit.x": func() string { return validateBar() },
}

// notAHookEntry is not a root: what it calls stays out of coverage.
func notAHookEntry() string { return checkBaz() }
`

const hookCoverageOtherSrc = `package app

import "strings"

func checkFoo() string    { return strings.TrimSpace("") }
func validateBar() string { return "" }
func checkBaz() string    { return "" }
func helperOnly() string  { return "" }
`

func writeHookCoverageFixture(t *testing.T, registryTail, hooksSrc string) string {
	t.Helper()
	dir := writeRegistryFixture(t, hookCoverageRegistryHeader+registryTail)
	app := filepath.Join(dir, "scripts", "ai-skill-cli", "internal", "app")
	files := map[string]string{
		"hooks.go":     hooksSrc,
		"other.go":     hookCoverageOtherSrc,
		"only_test.go": "package app\n\nfunc checkOnlyInTest() string { return \"\" }\n",
	}
	for name, src := range files {
		if err := os.WriteFile(filepath.Join(app, name), []byte(src), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

func orphanSymbols(errs []EnforcementRegistryLintError) map[string]bool {
	out := map[string]bool{}
	for _, e := range errs {
		if e.Type != "orphan_executor" {
			continue
		}
		for _, f := range e.Fields {
			if f.Key == "symbol" {
				out[f.Value] = true
			}
		}
	}
	return out
}

func TestLintHookCallCoverage_UnboundCalleeDetectedRegardlessOfName(t *testing.T) {
	dir := writeHookCoverageFixture(t, `
internal_helper_allowlist:
  symbols: [helperOnly]
rule_classes: []
`, hookCoverageHooksSrc)
	errs, err := LintEnforcementRegistry(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := orphanSymbols(errs)
	// checkFoo has no run*/validate* prefix: only hook-call coverage finds it.
	// validateBar lives outside hooks.go and is reached via the registry literal.
	for _, want := range []string{"checkFoo", "validateBar"} {
		if !got[want] {
			t.Fatalf("expected orphan_executor for %s; got %v", want, got)
		}
	}
}

func TestLintHookCallCoverage_LegalLookAlikesNotFlagged(t *testing.T) {
	dir := writeHookCoverageFixture(t, `
internal_helper_allowlist:
  symbols: [helperOnly, validateBar]
rule_classes:
  - id: fixture_class
    coverage: mechanical
    source_files: []
    rationale: "fixture"
    executors:
      - file: scripts/ai-skill-cli/internal/app/other.go
        symbol: checkFoo
        executor_kind: pre_commit_validator
        block_or_warn: block
      - file: scripts/ai-skill-cli/internal/app/hooks.go
        symbol: runPreCommitHook
        executor_kind: hook_dispatcher_entry
        block_or_warn: block
      - file: scripts/ai-skill-cli/internal/app/hooks.go
        symbol: runCommitMsgHook
        executor_kind: hook_dispatcher_entry
        block_or_warn: block
      - file: scripts/ai-skill-cli/internal/app/hooks.go
        symbol: runPrePushHook
        executor_kind: hook_dispatcher_entry
        block_or_warn: block
`, hookCoverageHooksSrc)
	errs, err := LintEnforcementRegistry(dir)
	if err != nil {
		t.Fatal(err)
	}
	got := orphanSymbols(errs)
	// bound (checkFoo), allowlisted (validateBar, helperOnly), called only by a
	// non-root (checkBaz), stdlib selector calls (strings.TrimSpace), roots
	// themselves, and test-only declarations must all stay out.
	for _, sym := range []string{"checkFoo", "validateBar", "helperOnly", "checkBaz", "TrimSpace", "runPreCommitHook", "checkOnlyInTest"} {
		if got[sym] {
			t.Fatalf("did not expect orphan_executor for %s; got %v", sym, got)
		}
	}
	for _, e := range errs {
		if e.Type == "hook_call_coverage_root_missing" || e.Type == "hook_call_coverage_unavailable" {
			t.Fatalf("unexpected %s: %s", e.Type, e.Format())
		}
	}
}

func TestLintHookCallCoverage_RootMissingFails(t *testing.T) {
	src := `package app

func runPreCommitHook() string { return "" }
func runCommitMsgHook() string { return "" }
var commitMsgValidatorRegistry = map[string]func() string{}
`
	dir := writeHookCoverageFixture(t, `
internal_helper_allowlist:
  symbols: []
rule_classes: []
`, src)
	errs, err := LintEnforcementRegistry(dir)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, e := range errs {
		if e.Type == "hook_call_coverage_root_missing" && e.IsFail() {
			for _, f := range e.Fields {
				if f.Key == "root" && f.Value == "runPrePushHook" {
					found = true
				}
			}
		}
	}
	if !found {
		t.Fatalf("expected FAIL hook_call_coverage_root_missing for runPrePushHook; got %+v", errs)
	}
}

// TestLintHookCallCoverage_RealRepo locks the live registry: the coverage set
// reaches the 8 symbols the old hooks.go naming scan missed, and every hook
// callee is classified (bound or allowlisted).
func TestLintHookCallCoverage_RealRepo(t *testing.T) {
	root := repoRootForLint(t)
	calls, missing, err := hookCallCoverage(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(missing) > 0 {
		t.Fatalf("hook entry roots missing: %v", missing)
	}
	for _, sym := range []string{
		"warnSanitizationIncidentScore",
		"validateReviewArchitectureNavigationDriftStaged",
		"warnPlanArchivalLinkIntegrity",
		"warnPlanEvidenceLineNumberCitations",
		"validateEnforcementRegistryTransition",
		"validateEnforcementRuleRegistrySync",
		"validatePlanEvidenceConvention",
		"validatePushGovernanceReplay",
	} {
		if _, ok := calls[sym]; !ok {
			t.Fatalf("hook-call coverage should reach %s", sym)
		}
	}
	errs, err := LintEnforcementRegistry(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range errs {
		switch e.Type {
		case "orphan_executor", "hook_call_coverage_root_missing", "hook_call_coverage_unavailable":
			t.Fatalf("live registry finding: %s", e.Format())
		}
	}
}
