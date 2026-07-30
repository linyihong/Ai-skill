package kge

import (
	"strings"
	"testing"
)

func TestMarkdownYamlSyncRule(t *testing.T) {
	eng := NewEngine(MarkdownYamlSyncRule{})
	ctx := Context{
		CommitMsg:   "feat: x",
		StagedPaths: []string{"governance/foo.md"},
		ExistingPaths: map[string]bool{
			"governance/foo.yaml": true,
		},
		Provided: map[CapabilityID]bool{
			CapCommitMsg: true, CapStagedPaths: true, CapRepoFS: true,
		},
	}
	f := eng.Run(ctx)
	if len(f) != 1 || f[0].Code != "markdown_yaml_sync" {
		t.Fatalf("want sync violation, got %#v", f)
	}

	ctx.StagedPaths = []string{"governance/foo.md", "governance/foo.yaml"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want pass when both staged, got %#v", got)
	}

	ctx.StagedPaths = []string{"governance/orphan.md"}
	ctx.ExistingPaths = map[string]bool{}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want pass for orphan md, got %#v", got)
	}

	ctx.StagedPaths = []string{"governance/foo.md"}
	ctx.ExistingPaths = map[string]bool{"governance/foo.yaml": true}
	ctx.CommitMsg = "feat: x\n\n[skip-markdown-yaml-sync]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestBootstrapEntryThinnessRule(t *testing.T) {
	eng := NewEngine(BootstrapEntryThinnessRule{})
	thin := "# entry\n\nRead CORE_BOOTSTRAP.md. That's the canonical source.\n"
	ctx := Context{
		CommitMsg:   "feat: x",
		StagedPaths: []string{"CLAUDE.md"},
		FileContents: map[string]string{"CLAUDE.md": thin},
		Provided: map[CapabilityID]bool{
			CapCommitMsg: true, CapStagedPaths: true, CapStagedContent: true,
		},
	}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want pass for thin entry, got %#v", got)
	}

	ctx.FileContents["CLAUDE.md"] = thin + "\nMode values: FAST/NORMAL/DEEP/FORENSIC/RECOVERY\n"
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "bootstrap_entry_thinness_content" {
		t.Fatalf("want content violation, got %#v", got)
	}

	many := strings.Repeat("line\n", 40)
	ctx.FileContents["CLAUDE.md"] = many
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "bootstrap_entry_thinness_lines" {
		t.Fatalf("want lines violation, got %#v", got)
	}

	ctx.CommitMsg = "feat: x\n\n[skip-bootstrap-thinness]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}

	ctx.CommitMsg = "feat: x"
	ctx.StagedPaths = []string{"README.md"}
	ctx.FileContents = map[string]string{"README.md": many}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want no enforcement on non-entry, got %#v", got)
	}
}

func TestGlossaryRetroOwnRule(t *testing.T) {
	eng := NewEngine(GlossaryRetroOwnRule{})
	ctx := Context{
		CommitMsg:   "feat: x",
		StagedPaths: []string{"runtime/cognitive-modes-discovery.yaml"},
		Provided:    map[CapabilityID]bool{CapCommitMsg: true, CapStagedPaths: true},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "glossary_retro_own" {
		t.Fatalf("want glossary violation, got %#v", got)
	}
	ctx.StagedPaths = append(ctx.StagedPaths, "knowledge/glossary/ai-skill.md")
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want pass with glossary staged, got %#v", got)
	}
	ctx.StagedPaths = []string{"runtime/cognitive-modes-discovery.yaml"}
	ctx.CommitMsg = "feat: x\n\n[skip-glossary-retro-own]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestRuntimeYamlProjectsRule(t *testing.T) {
	eng := NewEngine(RuntimeYamlProjectsRule{})
	ctx := Context{
		CommitMsg:    "feat: x",
		StagedPaths:  []string{"runtime/bad.yaml"},
		FileContents: map[string]string{"runtime/bad.yaml": "runtime_projection:\n  enabled: false\n"},
		Provided:     map[CapabilityID]bool{CapCommitMsg: true, CapStagedPaths: true, CapStagedContent: true},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "runtime_yaml_projects" {
		t.Fatalf("want projection violation, got %#v", got)
	}
	ctx.FileContents["runtime/bad.yaml"] = "runtime_projection:\n  enabled: true\ntarget_key: runtime.test.contract\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want pass for projected yaml, got %#v", got)
	}
	ctx.FileContents["runtime/bad.yaml"] = "runtime_projection:\n  enabled: false\n"
	ctx.CommitMsg = "feat: x\n\n[skip-runtime-yaml-projection]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestTokenBudgetRule(t *testing.T) {
	eng := NewEngine(TokenBudgetRule{})
	modes := map[string]string{
		"execution_mode": "NORMAL", "context_mode": "SUMMARY_FIRST",
		"governance_mode": "STANDARD", "memory_mode": "NONE",
	}
	ctx := Context{
		CommitMsg: "feat: x\n\nToken Estimate: 3000\n",
		Modes:     modes,
		Provided:  map[CapabilityID]bool{CapCommitMsg: true, CapModes: true},
	}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want pass under default NORMAL budget, got %#v", got)
	}
	ctx.CommitMsg = "feat: x\n\nToken Estimate: 9999\n"
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "token_budget" {
		t.Fatalf("want budget violation, got %#v", got)
	}
	ctx.CommitMsg = "feat: x\n\nToken Estimate: 999999\n\n[skip-token-budget]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestExecutionModeFloorsRule(t *testing.T) {
	eng := NewEngine(ExecutionModeFloorsRule{})
	ctx := Context{
		Modes:       map[string]string{"execution_mode": "FAST"},
		StagedPaths: []string{"enforcement/foo.md"},
		Provided:    map[CapabilityID]bool{CapModes: true, CapStagedPaths: true},
	}
	if got := eng.Run(ctx); len(got) != 1 {
		t.Fatalf("want FAST violation, got %#v", got)
	}
	ctx.Modes = map[string]string{"execution_mode": "DEEP", "governance_mode": "STRICT", "context_mode": "SOURCE_BACKED"}
	ctx.StagedPaths = nil
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want DEEP ok, got %#v", got)
	}
	ctx.Modes = map[string]string{"execution_mode": "RECOVERY", "governance_mode": "STRICT", "context_mode": "CHECKLIST_FIRST", "memory_mode": "EPISODIC"}
	if got := eng.Run(ctx); len(got) != 1 {
		t.Fatalf("want RECOVERY memory violation, got %#v", got)
	}
}

func TestGovernanceModeConsistencyRule(t *testing.T) {
	eng := NewEngine(GovernanceModeConsistencyRule{})
	ctx := Context{
		CommitMsg:   "feat: x",
		Modes:       map[string]string{"governance_mode": "LIGHT"},
		StagedPaths: []string{"enforcement/x.md"},
		Provided:    map[CapabilityID]bool{CapCommitMsg: true, CapModes: true, CapStagedPaths: true},
	}
	if got := eng.Run(ctx); len(got) != 1 {
		t.Fatalf("want LIGHT violation, got %#v", got)
	}
	ctx.Modes = map[string]string{"governance_mode": "LOCKDOWN"}
	ctx.StagedPaths = nil
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "governance_mode_lockdown_approval" {
		t.Fatalf("want LOCKDOWN approval violation, got %#v", got)
	}
	ctx.CommitMsg = "feat: critical\n\n[approved-by: alice]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want LOCKDOWN ok with approval, got %#v", got)
	}
}

func TestMemoryModeSubdirRule(t *testing.T) {
	eng := NewEngine(MemoryModeSubdirRule{})
	ctx := Context{
		Modes:       map[string]string{"memory_mode": "NONE"},
		StagedPaths: []string{"memory/episodic/foo.md"},
		Provided:    map[CapabilityID]bool{CapModes: true, CapStagedPaths: true},
	}
	if got := eng.Run(ctx); len(got) != 1 {
		t.Fatalf("want NONE violation, got %#v", got)
	}
	ctx.Modes = map[string]string{"memory_mode": "EPISODIC"}
	ctx.StagedPaths = []string{"memory/episodic/foo.md"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want EPISODIC ok, got %#v", got)
	}
	ctx.Modes = map[string]string{"memory_mode": "NONE"}
	ctx.StagedPaths = []string{"memory/README.md", "memory/retrieval-governance/activation-thresholds.md"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want layer doc exemption, got %#v", got)
	}
}

func TestActivationSignalsRule(t *testing.T) {
	eng := NewEngine(ActivationSignalsRule{})
	known := map[string]bool{"file_diff_runtime_schema": true, "user_keyword_deep": true}
	ctx := Context{
		CommitMsg: "feat: x\n\nactivation_reason:\n  - file_diff_runtime_schema\n",
		Modes:     map[string]string{},
		KnownSignals: known,
		Provided: map[CapabilityID]bool{
			CapCommitMsg: true, CapModes: true, CapKnownSignals: true,
		},
	}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want known signal pass, got %#v", got)
	}
	ctx.CommitMsg = "feat: x\n\nactivation_reason:\n  - made_up_signal\n"
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "activation_signals_unknown" {
		t.Fatalf("want unknown signal, got %#v", got)
	}
	ctx.CommitMsg = "feat: x\n"
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "activation_signals_missing" {
		t.Fatalf("want missing signal, got %#v", got)
	}
	ctx.Modes = map[string]string{"activation_signal": "user_keyword_deep"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want compact Sig pass, got %#v", got)
	}
	ctx.KnownSignals = map[string]bool{}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "activation_signals_vocab_unavailable" {
		t.Fatalf("want vocab unavailable, got %#v", got)
	}
}

func TestCapabilitySnippetRule(t *testing.T) {
	eng := NewEngine(CapabilitySnippetRule{})
	ctx := Context{
		CommitMsg: "feat: x\n",
		Modes:     map[string]string{"execution_mode": "DEEP", "governance_mode": "STRICT"},
		Provided:  map[CapabilityID]bool{CapCommitMsg: true, CapModes: true},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "capability_snippet_missing" {
		t.Fatalf("want snippet missing, got %#v", got)
	}
	ctx.CommitMsg = "feat: x\n\nCapability summary:\n  DEEP -> source-backed\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want snippet present pass, got %#v", got)
	}
	ctx.Modes = map[string]string{"execution_mode": "NORMAL", "governance_mode": "STANDARD"}
	ctx.CommitMsg = "feat: x\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want low-risk skip, got %#v", got)
	}
}

func TestAdaptiveTriggersRule(t *testing.T) {
	eng := NewEngine(AdaptiveTriggersRule{})
	modes := map[string]string{
		"execution_mode": "NORMAL", "context_mode": "SUMMARY_FIRST",
		"governance_mode": "STANDARD", "memory_mode": "EPISODIC",
	}
	ctx := Context{
		CommitMsg: "feat: reconcile contradict plans/active/a.md vs constitution/ADR-001.md",
		Modes:     modes,
		Provided:  map[CapabilityID]bool{CapCommitMsg: true, CapModes: true},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "adaptive_contradiction_risk" {
		t.Fatalf("want contradiction_risk, got %#v", got)
	}
	ctx.Modes = map[string]string{
		"execution_mode": "DEEP", "context_mode": "SOURCE_BACKED",
		"governance_mode": "STRICT", "memory_mode": "DECISION_REPLAY",
	}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want STRICT+SOURCE_BACKED pass, got %#v", got)
	}

	ctx.Modes = modes
	ctx.CommitMsg = "fix: address enforcement/failure-patterns/foo.md and enforcement/failure-patterns/bar.md"
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "adaptive_repeated_failure" {
		t.Fatalf("want repeated_failure, got %#v", got)
	}
	ctx.Modes = map[string]string{
		"execution_mode": "RECOVERY", "context_mode": "CHECKLIST_FIRST",
		"governance_mode": "STRICT", "memory_mode": "FAILURE_REPLAY",
	}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want RECOVERY+FAILURE_REPLAY pass, got %#v", got)
	}

	ctx.Modes = modes
	ctx.CommitMsg = "feat: medium work\n\nToken Estimate: 4500\n"
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "adaptive_budget_near_ceiling" {
		t.Fatalf("want near-ceiling, got %#v", got)
	}

	ctx.CommitMsg = "feat: reconcile contradict plans/a.md vs constitution/ADR-001.md\n\n[skip-adaptive]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestEvidenceHierarchyRule(t *testing.T) {
	eng := NewEngine(EvidenceHierarchyRule{})
	ctx := Context{
		CommitMsg:   "feat(x): Phase 1 完成",
		StagedPaths: []string{"scripts/ai-skill-cli/internal/app/x.go"},
		Provided:    map[CapabilityID]bool{CapCommitMsg: true, CapStagedPaths: true},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "evidence_hierarchy" {
		t.Fatalf("want evidence hierarchy violation, got %#v", got)
	}
	ctx.CommitMsg = "feat(x): done — tests pass and fixture green"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want evidence pass, got %#v", got)
	}
	ctx.CommitMsg = "feat(x): Phase 1 complete\n\n[skip-evidence-hierarchy]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
	ctx.CommitMsg = "docs: Phase 1 complete"
	ctx.StagedPaths = []string{"README.md"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want doc-only bypass, got %#v", got)
	}
}

func TestPlanStatusSyncRule(t *testing.T) {
	eng := NewEngine(PlanStatusSyncRule{})
	body := "feat: Phase 3 完成\n\nSee plans/active/2026-05-22-1629-runtime-cognitive-modes-system.md"
	ctx := Context{
		CommitMsg:   body,
		StagedPaths: []string{"scripts/ai-skill-cli/internal/app/hooks.go"},
		Provided:    map[CapabilityID]bool{CapCommitMsg: true, CapStagedPaths: true},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "plan_status_sync" {
		t.Fatalf("want plan status violation, got %#v", got)
	}
	ctx.StagedPaths = []string{"plans/active/2026-05-22-1629-runtime-cognitive-modes-system.md"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want staged plan pass, got %#v", got)
	}
	ctx.CommitMsg = "docs: see plans/active/foo.md for context\n\nPhase 3 context"
	ctx.StagedPaths = nil
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want no completion trigger, got %#v", got)
	}
	ctx.CommitMsg = body + "\n\n[skip-plan-status-sync]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestPlanCheckboxSyncRule(t *testing.T) {
	eng := NewEngine(PlanCheckboxSyncRule{})
	plan := "plans/active/foo.md"
	ctx := Context{
		CommitMsg:   "feat: work on " + plan,
		StagedPaths: []string{"scripts/ai-skill-cli/internal/app/x.go"},
		PathDiffs:   map[string]string{},
		Provided: map[CapabilityID]bool{
			CapCommitMsg: true, CapStagedPaths: true, CapStagedDiff: true,
		},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "plan_checkbox_sync" {
		t.Fatalf("want not-staged violation, got %#v", got)
	}
	ctx.StagedPaths = []string{plan, "scripts/ai-skill-cli/internal/app/x.go"}
	ctx.PathDiffs = map[string]string{plan: "+some other change\n"}
	if got := eng.Run(ctx); len(got) != 1 {
		t.Fatalf("want no-flip violation, got %#v", got)
	}
	ctx.PathDiffs = map[string]string{plan: "+- [x] done item\n"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want flip pass, got %#v", got)
	}
	if !PlanDiffFlipsCheckbox("+  - [X] Done\n") {
		t.Fatal("want PlanDiffFlipsCheckbox true for + - [X]")
	}
	ctx.CommitMsg = ctx.CommitMsg + "\n\n[skip-plan-checkbox-sync]\n"
	ctx.PathDiffs = map[string]string{plan: "+no flip\n"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestPlanArchivalAuditRule(t *testing.T) {
	eng := NewEngine(PlanArchivalAuditRule{})
	path := "plans/archived/foo.md"
	ctx := Context{
		CommitMsg:   "chore: archive plan",
		StagedPaths: []string{path},
		FileContents: map[string]string{
			path: "# plan\n\n- [ ] leftover\n- [x] done\n",
		},
		Provided: map[CapabilityID]bool{
			CapCommitMsg: true, CapStagedPaths: true, CapStagedContent: true,
		},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "plan_archival_audit" {
		t.Fatalf("want archival audit violation, got %#v", got)
	}
	ctx.CommitMsg = "chore: archive plan — remaining items deferred"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want deferred justification pass, got %#v", got)
	}
	ctx.CommitMsg = "chore: archive plan"
	ctx.FileContents[path] = "# plan\n\n- [x] all done\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want all-checked pass, got %#v", got)
	}
	ctx.FileContents[path] = "# plan\n\n- [ ] leftover\n"
	ctx.CommitMsg = "chore: archive\n\n[skip-plan-archival-audit]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestNoNewShellScriptsRule(t *testing.T) {
	eng := NewEngine(NoNewShellScriptsRule{})
	ctx := Context{
		CommitMsg:  "feat: add helper",
		AddedPaths: []string{"scripts/foo.sh", "scripts/bar.go"},
		Provided:   map[CapabilityID]bool{CapCommitMsg: true, CapAddedPaths: true},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "no_new_shell_scripts" {
		t.Fatalf("want shell block, got %#v", got)
	}
	ctx.AddedPaths = []string{"scripts/existing.go"}
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want no .sh pass, got %#v", got)
	}
	ctx.AddedPaths = []string{"scripts/foo.sh"}
	ctx.CommitMsg = "feat: x\n\n[skip-go-migration]\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}

func TestRuntimeTriggerWiringRule(t *testing.T) {
	eng := NewEngine(RuntimeTriggerWiringRule{})
	reg := "knowledge/runtime/routing-registry.yaml"
	ctx := Context{
		CommitMsg:   "feat: add route",
		StagedPaths: []string{reg},
		PathDiffs: map[string]string{
			reg: "+  - id: route.orphan.test\n",
		},
		SearchCorpus: "unrelated content",
		FileContents: map[string]string{reg: "routes: []\n"},
		Provided: map[CapabilityID]bool{
			CapCommitMsg: true, CapStagedPaths: true, CapStagedDiff: true,
			CapSearchCorpus: true, CapStagedContent: true,
		},
	}
	if got := eng.Run(ctx); len(got) != 1 || got[0].Code != "runtime_trigger_wiring" {
		t.Fatalf("want orphan route, got %#v", got)
	}
	ctx.SearchCorpus = "mentions route.orphan.test in discovery"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want wired via corpus, got %#v", got)
	}
	ctx.SearchCorpus = ""
	ctx.PathDiffs[reg] = "+  - id: route.manual.only\n+    manual_activation:\n+      reason: docs\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want manual_activation pass, got %#v", got)
	}
	yamlPath := "runtime/foo.yaml"
	ctx.StagedPaths = []string{yamlPath}
	ctx.PathDiffs = map[string]string{
		yamlPath: "+  target_key: runtime.orphan.key\n",
	}
	ctx.FileContents = map[string]string{reg: "no key here\n"}
	if got := eng.Run(ctx); len(got) != 1 {
		t.Fatalf("want orphan target_key, got %#v", got)
	}
	ctx.FileContents[reg] = "uses runtime.orphan.key\n"
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want registry consumer pass, got %#v", got)
	}
	ctx.CommitMsg = "feat: x\n\n[skip-runtime-trigger-wiring]\n"
	ctx.FileContents[reg] = ""
	if got := eng.Run(ctx); len(got) != 0 {
		t.Fatalf("want opt-out pass, got %#v", got)
	}
}
