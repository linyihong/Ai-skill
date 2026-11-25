package app

import (
	"os/exec"
	"path/filepath"
	"strings"
)

// resolveGitPath returns the filesystem path git itself uses for rel inside
// the repository at root (git rev-parse --git-path). Unlike joining
// <root>/.git/<rel>, this works in linked worktrees, where .git is a file:
// shared entries such as info/* resolve to the common git dir, and
// per-worktree entries such as COMMIT_EDITMSG resolve to the worktree's own
// git dir. In a normal checkout the result equals <root>/.git/<rel>.
//
// If git cannot resolve the path (root is not a git repository, git is
// missing), it falls back to <root>/.git/<rel> so callers keep their
// previous not-found behaviour.
func resolveGitPath(root, rel string) string {
	fallback := filepath.Join(root, ".git", filepath.FromSlash(rel))
	output, err := exec.Command("git", "-C", root, "rev-parse", "--git-path", rel).Output()
	if err != nil {
		return fallback
	}
	resolved := strings.TrimSpace(string(output))
	if resolved == "" {
		return fallback
	}
	resolved = filepath.FromSlash(resolved)
	if !filepath.IsAbs(resolved) {
		resolved = filepath.Join(root, resolved)
	}
	return filepath.Clean(resolved)
}
