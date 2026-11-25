package app

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// newRepoWithWorktree creates a real repository with one commit plus a linked
// worktree (where .git is a file, not a directory) and returns both roots.
func newRepoWithWorktree(t *testing.T) (main, worktree string) {
	t.Helper()
	base := t.TempDir()
	main = filepath.Join(base, "main")
	worktree = filepath.Join(base, "wt")
	if err := os.MkdirAll(main, 0o755); err != nil {
		t.Fatal(err)
	}
	runGit(t, main, "init")
	runGit(t, main, "config", "user.email", "allowed@example.invalid")
	runGit(t, main, "config", "user.name", "Allowed")
	runGit(t, main, "config", "commit.gpgsign", "false")
	runGit(t, main, "config", "core.hooksPath", "/dev/null")
	if err := os.WriteFile(filepath.Join(main, "README.md"), []byte("x\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	runGit(t, main, "add", "README.md")
	runGit(t, main, "commit", "-m", "init")
	runGit(t, main, "worktree", "add", "-b", "wt-branch", worktree)

	info, err := os.Stat(filepath.Join(worktree, ".git"))
	if err != nil {
		t.Fatal(err)
	}
	if info.IsDir() {
		t.Fatal("expected .git to be a file in a linked worktree")
	}
	return main, worktree
}

func writeAllowlistInfoFile(t *testing.T, repo, email string) {
	t.Helper()
	infoDir := filepath.Join(repo, ".git", "info")
	if err := os.MkdirAll(infoDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(infoDir, allowedAuthorEmailsInfoFile), []byte(email+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadAuthorEmailAllowlistInLinkedWorktree(t *testing.T) {
	main, worktree := newRepoWithWorktree(t)
	writeAllowlistInfoFile(t, main, "allowed@example.invalid")

	allowed, err := loadAuthorEmailAllowlist(worktree)
	if err != nil {
		t.Fatalf("expected allowlist to load in worktree, got: %v", err)
	}
	if _, ok := allowed["allowed@example.invalid"]; !ok {
		t.Fatalf("expected shared info-file entry in worktree allowlist, got: %v", allowed)
	}
	if msg := validateAuthorEmailAllowlist(worktree); msg != "" {
		t.Fatalf("expected allowlisted author to pass in worktree, got: %s", msg)
	}
}

func TestLoadAuthorEmailAllowlistInWorktreeWithoutInfoFile(t *testing.T) {
	_, worktree := newRepoWithWorktree(t)
	if _, err := loadAuthorEmailAllowlist(worktree); err != nil {
		t.Fatalf("missing info file must not fail in worktree, got: %v", err)
	}
}

func TestResolveGitPathNormalCheckoutUnchanged(t *testing.T) {
	main, _ := newRepoWithWorktree(t)
	for _, rel := range []string{"info/" + allowedAuthorEmailsInfoFile, "COMMIT_EDITMSG", "ai-skill-agent.lock", "info/exclude"} {
		got := resolveGitPath(main, rel)
		want := filepath.Join(main, ".git", filepath.FromSlash(rel))
		if !sameFilePath(t, got, want) {
			t.Fatalf("resolveGitPath(%q) = %q, want %q", rel, got, want)
		}
	}
}

func TestResolveGitPathWorktreeSplitsSharedAndPerWorktree(t *testing.T) {
	main, worktree := newRepoWithWorktree(t)
	commonInfo := filepath.Join(main, ".git", "info", "exclude")
	if got := resolveGitPath(worktree, "info/exclude"); !sameFilePath(t, got, commonInfo) {
		t.Fatalf("info/exclude should resolve to common dir %q, got %q", commonInfo, got)
	}
	msg := resolveGitPath(worktree, "COMMIT_EDITMSG")
	if !strings.Contains(filepath.ToSlash(msg), "/.git/worktrees/") {
		t.Fatalf("COMMIT_EDITMSG should resolve to per-worktree git dir, got %q", msg)
	}
}

func TestEnsureGoalGitExcludeInLinkedWorktree(t *testing.T) {
	main, worktree := newRepoWithWorktree(t)
	if err := ensureGoalGitExclude(worktree); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(main, ".git", "info", "exclude"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), ".agent-goals/") {
		t.Fatalf("expected .agent-goals/ in shared exclude, got: %q", data)
	}
}

func TestCloseLoopLockInLinkedWorktree(t *testing.T) {
	_, worktree := newRepoWithWorktree(t)
	release, err := acquireCloseLoopLock(worktree)
	if err != nil {
		t.Fatalf("expected lock acquire in worktree, got: %v", err)
	}
	if check := closeLoopLockCheck(worktree); check.Status != "blocked" {
		t.Fatalf("expected held lock to be visible, got %+v", check)
	}
	release()
	if check := closeLoopLockCheck(worktree); check.Status != "ok" {
		t.Fatalf("expected lock released, got %+v", check)
	}
}

// sameFilePath compares paths after resolving symlinks (macOS /var → /private/var).
func sameFilePath(t *testing.T, a, b string) bool {
	t.Helper()
	norm := func(p string) string {
		dir, base := filepath.Dir(p), filepath.Base(p)
		for {
			if resolved, err := filepath.EvalSymlinks(dir); err == nil {
				return filepath.Join(resolved, base)
			}
			parent := filepath.Dir(dir)
			if parent == dir {
				return p
			}
			base = filepath.Join(filepath.Base(dir), base)
			dir = parent
		}
	}
	return norm(a) == norm(b)
}
