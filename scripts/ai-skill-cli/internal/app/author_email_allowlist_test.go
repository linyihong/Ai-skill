package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateAuthorEmailAllowlistSkipsWhenUnset(t *testing.T) {
	repo := t.TempDir()
	runGit(t, repo, "init")
	runGit(t, repo, "config", "user.email", "anyone@example.invalid")
	runGit(t, repo, "config", "user.name", "Anyone")

	if msg := validateAuthorEmailAllowlist(repo); msg != "" {
		t.Fatalf("expected skip when allowlist unset, got: %s", msg)
	}
}

func TestValidateAuthorEmailAllowlistBlocksNonAllowlistedEmail(t *testing.T) {
	repo := t.TempDir()
	runGit(t, repo, "init")
	runGit(t, repo, "config", "user.email", "blocked@example.invalid")
	runGit(t, repo, "config", "user.name", "Blocked")
	runGit(t, repo, "config", "--add", allowedAuthorEmailConfigKey, "allowed@example.invalid")

	msg := validateAuthorEmailAllowlist(repo)
	if msg == "" {
		t.Fatal("expected block for non-allowlisted email")
	}
	if !strings.Contains(msg, "blocked@example.invalid") {
		t.Fatalf("expected rejected email in message, got: %s", msg)
	}
	if !strings.Contains(msg, "allowed@example.invalid") {
		t.Fatalf("expected allowlist contents in message, got: %s", msg)
	}
}

func TestValidateAuthorEmailAllowlistAllowsConfiguredEmail(t *testing.T) {
	repo := t.TempDir()
	runGit(t, repo, "init")
	runGit(t, repo, "config", "user.email", "allowed@example.invalid")
	runGit(t, repo, "config", "user.name", "Allowed")
	runGit(t, repo, "config", "--add", allowedAuthorEmailConfigKey, "allowed@example.invalid")

	if msg := validateAuthorEmailAllowlist(repo); msg != "" {
		t.Fatalf("expected allow, got: %s", msg)
	}
}

func TestValidateAuthorEmailAllowlistReadsInfoFile(t *testing.T) {
	repo := t.TempDir()
	runGit(t, repo, "init")
	runGit(t, repo, "config", "user.email", "allowed@example.invalid")
	runGit(t, repo, "config", "user.name", "Allowed")

	infoDir := filepath.Join(repo, ".git", "info")
	if err := os.MkdirAll(infoDir, 0o755); err != nil {
		t.Fatal(err)
	}
	content := "# local only\nallowed@example.invalid\n"
	if err := os.WriteFile(filepath.Join(infoDir, allowedAuthorEmailsInfoFile), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	if msg := validateAuthorEmailAllowlist(repo); msg != "" {
		t.Fatalf("expected allow via info file, got: %s", msg)
	}

	runGit(t, repo, "config", "user.email", "other@example.invalid")
	msg := validateAuthorEmailAllowlist(repo)
	if msg == "" {
		t.Fatal("expected block when info-file allowlist rejects email")
	}
}

func TestEmailFromGitIdent(t *testing.T) {
	got := emailFromGitIdent("Allowed <allowed@example.invalid> 1710000000 +0900")
	if got != "allowed@example.invalid" {
		t.Fatalf("got %q", got)
	}
}
