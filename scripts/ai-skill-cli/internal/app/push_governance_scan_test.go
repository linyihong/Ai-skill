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
	runGit(t, repo, "checkout", "-b", "main")
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
	runGit(t, repo, "checkout", "-b", "main")
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
	runGit(t, repo, "checkout", "-b", "main")
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
