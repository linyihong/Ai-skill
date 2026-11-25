package app

import (
	"fmt"
	"os"
	"testing"
)

// TestMain points TMPDIR at a per-run directory so hook tests never write to
// the real hook runtime files (bootstrap caches, logs, and the incremental
// transcript shadow ledger whose records count as Phase 2a evidence in
// plans/active/2026-11-25-0842-verification-architecture-contract).
func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "ai-skill-app-test-")
	if err != nil {
		fmt.Fprintln(os.Stderr, "TestMain: create temp dir:", err)
		os.Exit(1)
	}
	if err := os.Setenv("TMPDIR", dir); err != nil {
		fmt.Fprintln(os.Stderr, "TestMain: set TMPDIR:", err)
		os.Exit(1)
	}
	code := m.Run()
	_ = os.RemoveAll(dir)
	os.Exit(code)
}
