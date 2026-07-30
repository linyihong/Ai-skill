package kge

import "testing"

func TestCognitiveCostRule_OK(t *testing.T) {
	eng := NewEngine(CognitiveCostRule{})
	ctx := Context{
		Modes: map[string]string{
			"execution_mode": "NORMAL",
			"context_mode":   "SUMMARY_FIRST",
			"cognitive_cost": "LOW",
		},
		Provided: map[CapabilityID]bool{CapModes: true},
	}
	f := eng.Run(ctx)
	if len(f) != 0 {
		t.Fatalf("want no findings, got %#v", f)
	}
}

func TestCognitiveCostRule_Mismatch(t *testing.T) {
	eng := NewEngine(CognitiveCostRule{})
	ctx := Context{
		Modes: map[string]string{
			"execution_mode": "DEEP",
			"context_mode":   "SOURCE_BACKED",
			"cognitive_cost": "LOW",
		},
		Provided: map[CapabilityID]bool{CapModes: true},
	}
	f := eng.Run(ctx)
	if len(f) != 1 || f[0].Code != "cognitive_cost_mismatch" {
		t.Fatalf("want mismatch finding, got %#v", f)
	}
	if !Blocking(f) {
		t.Fatal("mismatch should block")
	}
}

func TestCognitiveCostRule_MissingCapability(t *testing.T) {
	eng := NewEngine(CognitiveCostRule{})
	f := eng.Run(Context{})
	if len(f) != 1 || f[0].Code != "capability_missing" {
		t.Fatalf("want capability_missing, got %#v", f)
	}
}

func TestCLIDocSyncRule_RequiresDiffFromAdapter(t *testing.T) {
	eng := NewEngine(CLIDocSyncRule{})
	ctx := Context{
		CommitMsg:   "feat: x",
		StagedPaths: []string{"scripts/ai-skill-cli/internal/app/hooks.go"},
		Provided: map[CapabilityID]bool{
			CapStagedPaths: true,
			CapCommitMsg:   true,
			// CapStagedDiff intentionally absent
		},
	}
	f := eng.Run(ctx)
	if len(f) != 1 || f[0].Code != "capability_missing" {
		t.Fatalf("want capability_missing for staged_diff, got %#v", f)
	}
}

func TestCLIDocSyncRule_NoGit_DetectsNewHook(t *testing.T) {
	eng := NewEngine(CLIDocSyncRule{})
	ctx := Context{
		CommitMsg:   "feat: add hook",
		StagedPaths: []string{"scripts/ai-skill-cli/internal/app/hooks.go"},
		StagedDiff:  "+func runSomethingHook(ctx context.Context) error {\n",
		Provided: map[CapabilityID]bool{
			CapStagedPaths: true,
			CapStagedDiff:  true,
			CapCommitMsg:   true,
		},
	}
	f := eng.Run(ctx)
	if len(f) != 1 || f[0].Code != "cli_doc_sync" {
		t.Fatalf("want cli_doc_sync finding, got %#v", f)
	}
}

func TestCLIDocSyncRule_SkipTrailer(t *testing.T) {
	eng := NewEngine(CLIDocSyncRule{})
	ctx := Context{
		CommitMsg:   "feat: x\n\n[skip-cli-doc-sync]\n",
		StagedPaths: []string{"scripts/ai-skill-cli/internal/app/hooks.go"},
		StagedDiff:  "+func runSomethingHook() {}\n",
		Provided: map[CapabilityID]bool{
			CapStagedPaths: true,
			CapStagedDiff:  true,
			CapCommitMsg:   true,
		},
	}
	f := eng.Run(ctx)
	if len(f) != 0 {
		t.Fatalf("skip trailer should silence, got %#v", f)
	}
}
