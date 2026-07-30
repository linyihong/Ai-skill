package app

import (
	"strings"
	"testing"
)

func TestRunKGECognitiveCost_Parity(t *testing.T) {
	ok := map[string]string{
		"execution_mode": "NORMAL",
		"context_mode":   "SUMMARY_FIRST",
		"cognitive_cost": "LOW",
	}
	if got := validateCognitiveCost(ok); got != "" {
		t.Fatalf("want pass, got %q", got)
	}
	bad := map[string]string{
		"execution_mode": "DEEP",
		"context_mode":   "SOURCE_BACKED",
		"cognitive_cost": "LOW",
	}
	if got := validateCognitiveCost(bad); got == "" || !strings.Contains(got, "cognitive_cost mismatch") {
		t.Fatalf("want mismatch via KGE, got %q", got)
	}
}

func TestRunKGECLIDocSync_SkipAndAdapter(t *testing.T) {
	root := t.TempDir()
	if got := validateCLIDocSync("feat: x", []string{"README.md"}, root); got != "" {
		t.Fatalf("want pass for unrelated staged, got %q", got)
	}
	msg := "feat: x\n\n[skip-cli-doc-sync]\n"
	staged := []string{"scripts/ai-skill-cli/internal/app/hooks.go"}
	if got := validateCLIDocSync(msg, staged, root); got != "" {
		t.Fatalf("want skip trailer pass, got %q", got)
	}
}
