package app

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPushGovernanceReplayBlocksSanitizationBypass(t *testing.T) {
	repo := initTempGitRepo(t)
	seedSanitizationRuntimeDB(t, repo, map[string]bool{"plans/": true}, []derivedMatchToken{
		{
			Token:                "SecretProject",
			CanonicalToken:       "SecretProject",
			EntityName:           "Secret Project",
			Kind:                 "codename",
			OwningProjectID:      "secret-project",
			SourceMetadataPath:   ".ai-skill-project.yaml",
			SuggestedPlaceholder: "<SECRET_PROJECT>",
		},
	})
	rel := "plans/active/leak.md"
	writeFile(t, filepath.Join(repo, rel), "# leak\n\nSecretProject detail\n")
	runGit(t, repo, "add", rel)
	runGit(t, repo, "commit", "-qm", "bad: leak private token\n\n[skip-cognitive-mode]", "--no-verify")

	msg, _ := validatePushGovernanceReplay(repo)
	if msg == "" {
		t.Fatal("expected push governance replay to block sanitization leak committed with --no-verify")
	}
	if !strings.Contains(msg, "sanitization-scan") || !strings.Contains(msg, "SecretProject") {
		t.Fatalf("unexpected message:\n%s", msg)
	}
}

func TestPushGovernanceReplayBlocksMissingCognitiveMode(t *testing.T) {
	repo := initTempGitRepo(t)
	seedSanitizationRuntimeDB(t, repo, map[string]bool{"plans/": true}, nil)
	rel := "plans/active/ok.md"
	writeFile(t, filepath.Join(repo, rel), "# ok\n\ngeneric lesson\n")
	runGit(t, repo, "add", rel)
	runGit(t, repo, "commit", "-qm", "docs: no cognitive block", "--no-verify")

	msg, _ := validatePushGovernanceReplay(repo)
	if msg == "" || !strings.Contains(msg, "Cognitive Contract v2 block") {
		t.Fatalf("expected missing cognitive block, got:\n%s", msg)
	}
}

func TestPushGovernanceReplayPassesValidCommit(t *testing.T) {
	repo := initTempGitRepo(t)
	seedSanitizationRuntimeDB(t, repo, map[string]bool{"plans/": true}, nil)
	rel := "plans/active/ok.md"
	writeFile(t, filepath.Join(repo, rel), "# ok\n\ngeneric lesson\n")
	runGit(t, repo, "add", rel)
	runGit(t, repo, "commit", "-qm", "docs: ok\n\n[skip-cognitive-mode]", "--no-verify")

	msg, checks := validatePushGovernanceReplay(repo)
	if msg != "" {
		t.Fatalf("expected pass, got block:\n%s", msg)
	}
	if !hasCheckStatus(checks, "push_governance_replay", "ok") {
		t.Fatalf("expected ok replay check, got %#v", checks)
	}
}

// F6: pre-commit and pre-push replay share one forbidden set. Before F6 the
// replay only matched ".sh", so a --no-verify commit adding .py or an
// extensionless shebang script passed it.
func TestPushGovernanceReplayBlocksNonGoScriptsLikePreCommit(t *testing.T) {
	for _, tc := range []struct{ rel, body string }{
		{"scripts/tool.py", "print(1)\n"},
		{"scripts/deploy", "#!/usr/bin/env bash\necho hi\n"},
	} {
		repo := initTempGitRepo(t)
		seedSanitizationRuntimeDB(t, repo, map[string]bool{"plans/": true}, nil)
		writeFile(t, filepath.Join(repo, tc.rel), tc.body)
		runGit(t, repo, "add", tc.rel)
		runGit(t, repo, "commit", "-qm", "feat: add script\n\n[skip-cognitive-mode]", "--no-verify")
		msg, _ := validatePushGovernanceReplay(repo)
		if !strings.Contains(msg, "pre-commit shell policy") || !strings.Contains(msg, tc.rel) {
			t.Fatalf("%s: expected replay to block, got:\n%s", tc.rel, msg)
		}
	}
}

func TestPushGovernanceReplayAllowsHookAdapterInHooksPath(t *testing.T) {
	repo := initTempGitRepo(t)
	seedSanitizationRuntimeDB(t, repo, map[string]bool{"plans/": true}, nil)
	runGit(t, repo, "config", "core.hooksPath", ".githooks")
	rel := ".githooks/pre-commit"
	writeFile(t, filepath.Join(repo, rel), "#!/bin/sh\nexec ai-skill hooks run pre-commit \"$@\"\n")
	runGit(t, repo, "add", rel)
	runGit(t, repo, "commit", "-qm", "chore: add hook adapter\n\n[skip-cognitive-mode]", "--no-verify")
	if msg, _ := validatePushGovernanceReplay(repo); msg != "" {
		t.Fatalf("expected hook adapter in core.hooksPath to pass, got:\n%s", msg)
	}
}

// Pre-commit adapter reads the staged blob (index), not the work tree.
func TestRunKGENoNewShellScriptsDetectsStagedShebang(t *testing.T) {
	repo := initTempGitRepo(t)
	rel := "scripts/deploy"
	writeFile(t, filepath.Join(repo, rel), "#!/usr/bin/env bash\n")
	runGit(t, repo, "add", rel)
	if msg := runKGENoNewShellScripts("feat: x", repo); !strings.Contains(msg, rel+" (shebang)") {
		t.Fatalf("expected staged shebang block, got %q", msg)
	}
	writeFile(t, filepath.Join(repo, "notes"), "plain text\n")
	runGit(t, repo, "rm", "-q", "--cached", rel)
	runGit(t, repo, "add", "notes")
	if msg := runKGENoNewShellScripts("feat: x", repo); msg != "" {
		t.Fatalf("expected plain extensionless file to pass, got %q", msg)
	}
}
