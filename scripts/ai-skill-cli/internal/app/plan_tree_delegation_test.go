package app

// Phase 2 of sub-plan 03 (delegation schema). These tests encode the six
// acceptance gates the user set when releasing Phase 2:
//
//  1. enabled=true  => the 4 required fields must be present/non-empty.
//  2. context / constraints are always optional.
//  3. undeclared delegation => 100% pre-existing behavior (byte-identical).
//  4. enabled:false has an explicit, unambiguous meaning == undeclared.
//  5. Consumer Exclusive: the portable planvalidate engine has ZERO knowledge of
//     delegation (mechanical lock: source grep + struct reflection).
//  6. (glossary note lives in knowledge/glossary; not a Go test.)

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/internal/planvalidate"
)

// subWithDelegation builds a fully-valid sub-plan frontmatter (parent / required
// / reason all present) plus an arbitrary delegation block, so any violation the
// validator reports is delegation-specific and nothing else.
func subWithDelegation(t *testing.T, tmp, rel, delegationBlock string) {
	t.Helper()
	fm := "---\n" +
		"id: sub-deleg\n" +
		"plan_kind: sub\n" +
		"status: draft\n" +
		"owner: t\n" +
		"created: 2026-06-03\n" +
		"parent: main-x\n" +
		"required_for_completion: true\n" +
		"sub_plan_reason: r\n" +
		delegationBlock +
		"---"
	makePlan(t, tmp, rel, fm)
}

const delegationComplete = `delegation:
  enabled: true
  brief:
    goal: do the thing
    acceptance:
      - it works
    verification:
      - run the tests
  execution:
    modes:
      - human
`

// Gate 1 + Gate 2: enabled + all 4 required present, no context/constraints -> pass.
func TestDelegation_EnabledCompleteNoOptionals_Pass(t *testing.T) {
	tmp := t.TempDir()
	subWithDelegation(t, tmp, "plans/active/d.md", delegationComplete)
	got := validatePlanTreeFrontmatter("", []string{"plans/active/d.md"}, tmp)
	if got != "" {
		t.Fatalf("complete delegation (no context/constraints) should pass, got: %s", got)
	}
}

// Gate 1: missing brief.verification -> block, and names the missing field.
func TestDelegation_EnabledMissingVerification_Blocks(t *testing.T) {
	tmp := t.TempDir()
	block := `delegation:
  enabled: true
  brief:
    goal: do the thing
    acceptance:
      - it works
  execution:
    modes:
      - human
`
	subWithDelegation(t, tmp, "plans/active/d.md", block)
	got := validatePlanTreeFrontmatter("", []string{"plans/active/d.md"}, tmp)
	if !strings.Contains(got, "delegation.brief.verification") {
		t.Fatalf("expected missing delegation.brief.verification, got: %s", got)
	}
}

// Gate 1: empty execution.modes -> block.
func TestDelegation_EnabledEmptyModes_Blocks(t *testing.T) {
	tmp := t.TempDir()
	block := `delegation:
  enabled: true
  brief:
    goal: do the thing
    acceptance:
      - it works
    verification:
      - run the tests
  execution:
    modes: []
`
	subWithDelegation(t, tmp, "plans/active/d.md", block)
	got := validatePlanTreeFrontmatter("", []string{"plans/active/d.md"}, tmp)
	if !strings.Contains(got, "delegation.execution.modes") {
		t.Fatalf("expected missing delegation.execution.modes, got: %s", got)
	}
}

// Gate 1 robustness: acceptance/verification authored as scalars (not lists)
// still satisfy the non-empty requirement.
func TestDelegation_ScalarBriefFields_Pass(t *testing.T) {
	tmp := t.TempDir()
	block := `delegation:
  enabled: true
  brief:
    goal: do the thing
    acceptance: it works
    verification: run the tests
  execution:
    modes: human
`
	subWithDelegation(t, tmp, "plans/active/d.md", block)
	got := validatePlanTreeFrontmatter("", []string{"plans/active/d.md"}, tmp)
	if got != "" {
		t.Fatalf("scalar brief fields should pass, got: %s", got)
	}
}

// Gate 3: an undeclared delegation must produce byte-identical output to a plan
// that has no delegation block at all (here: both clean -> "").
func TestDelegation_Undeclared_IdenticalToBaseline(t *testing.T) {
	tmp := t.TempDir()
	makeSub(t, tmp, "plans/active/plain.md", "sub-plain", "main-x", "draft", "reason", true)
	got := validatePlanTreeFrontmatter("", []string{"plans/active/plain.md"}, tmp)
	if got != "" {
		t.Fatalf("plain sub (no delegation) baseline should pass, got: %s", got)
	}
}

// Gate 4: enabled:false is defined as identical to undeclared. Proven by output
// equality between (a) a sub with delegation.enabled:false and NO brief, and
// (b) the same sub with no delegation block. Neither may add any violation.
func TestDelegation_EnabledFalse_EqualsUndeclared(t *testing.T) {
	tmp := t.TempDir()
	// (a) enabled:false, deliberately no brief / modes at all.
	subWithDelegation(t, tmp, "plans/active/off.md", "delegation:\n  enabled: false\n")
	gotOff := validatePlanTreeFrontmatter("", []string{"plans/active/off.md"}, tmp)

	// (b) same valid sub, no delegation block.
	subWithDelegation(t, tmp, "plans/active/none.md", "")
	gotNone := validatePlanTreeFrontmatter("", []string{"plans/active/none.md"}, tmp)

	if gotOff != gotNone {
		t.Fatalf("enabled:false must equal undeclared: off=%q none=%q", gotOff, gotNone)
	}
	if gotOff != "" {
		t.Fatalf("enabled:false with no brief must not add any violation, got: %s", gotOff)
	}
}

// Gate 5a (Consumer Exclusive — source grep): the portable planvalidate engine
// package must contain no reference to "delegation" anywhere. This is the
// mechanical proof that deleting the consumer-layer delegation code cannot force
// an engine recompile — the engine simply never mentions it.
func TestDelegation_ConsumerExclusive_EngineSourceHasNoDelegation(t *testing.T) {
	engineDir := filepath.Join("..", "planvalidate")
	entries, err := os.ReadDir(engineDir)
	if err != nil {
		t.Fatalf("read planvalidate dir: %v", err)
	}
	sawGoFile := false
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".go") {
			continue
		}
		sawGoFile = true
		data, err := os.ReadFile(filepath.Join(engineDir, e.Name()))
		if err != nil {
			t.Fatalf("read %s: %v", e.Name(), err)
		}
		if strings.Contains(strings.ToLower(string(data)), "delegation") {
			t.Fatalf("planvalidate engine file %s references 'delegation' — the portable engine must stay unaware of the consumer-layer delegation feature", e.Name())
		}
	}
	if !sawGoFile {
		t.Fatal("expected to scan at least one planvalidate .go file")
	}
}

// Gate 5b (Consumer Exclusive — reflection): the engine's plan models must carry
// no delegation field. A second, structural lock independent of source text.
func TestDelegation_ConsumerExclusive_NormalizedModelHasNoDelegationField(t *testing.T) {
	for _, typ := range []reflect.Type{
		reflect.TypeOf(planvalidate.NormalizedPlanModel{}),
		reflect.TypeOf(planvalidate.RawPlan{}),
	} {
		for i := 0; i < typ.NumField(); i++ {
			if strings.Contains(strings.ToLower(typ.Field(i).Name), "delegation") {
				t.Fatalf("%s carries a delegation-related field %q — engine must be delegation-agnostic", typ.Name(), typ.Field(i).Name)
			}
		}
	}
}

// Reverse-direction lock (user Phase-2 review item 4): the CONSUMER must not
// start understanding workflow/tool semantics. The `execution` block is workflow,
// not validation input — so its schema is frozen to exactly {Modes, Constraints}.
// This mechanically blocks anyone growing execution.worktree / execution.agent /
// execution.<tool> that the validator could then interpret.
func TestDelegation_ExecutionSchemaLockedToModesAndConstraints(t *testing.T) {
	execField, ok := reflect.TypeOf(DelegationSpec{}).FieldByName("Execution")
	if !ok {
		t.Fatal("DelegationSpec has no Execution field")
	}
	execType := execField.Type
	allowed := map[string]bool{"Modes": true, "Constraints": true}
	for i := 0; i < execType.NumField(); i++ {
		if name := execType.Field(i).Name; !allowed[name] {
			t.Fatalf("execution schema grew field %q beyond {Modes, Constraints} — consumer must not encode tool/workflow semantics", name)
		}
	}
	if execType.NumField() != 2 {
		t.Fatalf("execution must have exactly 2 fields (Modes, Constraints), got %d", execType.NumField())
	}
}

// Reverse-direction lock (behavioral): the validator must never interpret the
// CONTENT of execution.constraints — it is opaque tool/adapter detail. Arbitrary,
// tool-looking, even non-existent-path constraints must still pass; validation
// only ever knows that an execution block exists (via modes), never what tools it
// names.
func TestDelegation_ConstraintsContentNeverInterpreted(t *testing.T) {
	tmp := t.TempDir()
	block := `delegation:
  enabled: true
  brief:
    goal: g
    acceptance: [a]
    verification: [v]
  execution:
    modes: [agent]
    constraints:
      - worktree=/nonexistent/path
      - requires-tool: some-tool-that-does-not-exist
      - arbitrary garbage the validator must not parse
`
	subWithDelegation(t, tmp, "plans/active/c.md", block)
	got := validatePlanTreeFrontmatter("", []string{"plans/active/c.md"}, tmp)
	if got != "" {
		t.Fatalf("constraints content must never be interpreted/validated, got: %s", got)
	}
}

// stripDelegationLines removes any output line mentioning delegation — the test
// analogue of deleting the validator's delegation branch. Non-delegation header /
// trailer / findings contain no "delegation" token and survive untouched.
func stripDelegationLines(s string) string {
	if s == "" {
		return ""
	}
	var kept []string
	for _, ln := range strings.Split(s, "\n") {
		if strings.Contains(strings.ToLower(ln), "delegation") {
			continue
		}
		kept = append(kept, ln)
	}
	return strings.Join(kept, "\n")
}

// Capability Removal Test (the user's single merge-blocking gate): removing the
// delegation capability must remove ONLY delegation findings and change nothing
// else. Two identical plan sets differing solely by the delegation block:
//   - the other three plan-tree validators (parent / archive-order / unique-id)
//     must be byte-identical (delegation is invisible to them);
//   - frontmatter output with delegation findings stripped must reproduce the
//     no-delegation output exactly.
// A shared NON-delegation finding (empty sub_plan_reason) is present in BOTH sets
// so the header/trailer exist on both sides and the delta is purely the
// delegation bullet — a faithful "delete the branch" simulation, not a vacuous
// both-empty comparison. (Engine / schema-compat non-impact is covered by the
// Consumer Exclusive locks above.)
func TestDelegation_CapabilityRemoval_OnlyDelegationFindingsChange(t *testing.T) {
	build := func(delegBlock string) string {
		tmp := t.TempDir()
		makeMain(t, tmp, "plans/active/main.md", "main-x", "draft")
		// sub is valid EXCEPT an empty sub_plan_reason (shared non-delegation
		// finding), plus the optional delegation block under test.
		fm := "---\n" +
			"id: sub-x\n" +
			"plan_kind: sub\n" +
			"status: draft\n" +
			"owner: t\n" +
			"created: 2026-06-03\n" +
			"parent: main-x\n" +
			"required_for_completion: true\n" +
			"sub_plan_reason: \"\"\n" +
			delegBlock +
			"---"
		makePlan(t, tmp, "plans/active/sub.md", fm)
		return tmp
	}
	staged := []string{"plans/active/main.md", "plans/active/sub.md"}

	// A: delegation enabled but incomplete (missing verification) -> a delegation
	// finding on top of the shared reason finding. B: no delegation block.
	withDeleg := "delegation:\n  enabled: true\n  brief:\n    goal: g\n    acceptance: [a]\n  execution:\n    modes: [human]\n"
	tmpA := build(withDeleg)
	tmpB := build("")

	// The other three validators must be byte-identical across A and B.
	for _, tc := range []struct {
		name string
		fn   func(string, []string, string) string
	}{
		{"parentReference", validatePlanTreeParentReference},
		{"uniqueID", validatePlanTreeUniqueID},
		{"archiveOrder", validatePlanTreeArchiveOrder},
	} {
		if a, b := tc.fn("", staged, tmpA), tc.fn("", staged, tmpB); a != b {
			t.Fatalf("%s changed by delegation presence (must be invisible): A=%q B=%q", tc.name, a, b)
		}
	}

	fa := validatePlanTreeFrontmatter("", staged, tmpA)
	fb := validatePlanTreeFrontmatter("", staged, tmpB)
	// Guard against a vacuous test: A must actually carry a delegation finding,
	// and B must carry the shared non-delegation finding.
	if !strings.Contains(strings.ToLower(fa), "delegation") {
		t.Fatalf("expected A to contain a delegation finding, got: %q", fa)
	}
	if !strings.Contains(fb, "sub_plan_reason") {
		t.Fatalf("expected B to keep the shared non-delegation finding, got: %q", fb)
	}
	// Removing delegation findings from A must reproduce B exactly — nothing else
	// (parent / reason / required / header / trailer) may shift.
	if got := stripDelegationLines(fa); got != fb {
		t.Fatalf("removing delegation must equal the no-delegation baseline:\n  A-stripped=%q\n  B=%q", got, fb)
	}
}
