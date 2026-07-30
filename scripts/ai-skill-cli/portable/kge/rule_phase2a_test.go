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
