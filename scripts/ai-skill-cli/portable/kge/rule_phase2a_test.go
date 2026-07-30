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
