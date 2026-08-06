package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Local-only author email allowlist (never committed).
//
// Config sources (merged; either alone is enough):
//  1. git config --local --get-all ai-skill.allowedAuthorEmail
//  2. optional file: .git/info/ai-skill-allowed-author-emails
//     (one email per line; # comments and blank lines ignored)
//
// Behavior:
//   - allowlist empty → skip (no restriction; other clones unaffected)
//   - allowlist set → GIT_AUTHOR_IDENT and GIT_COMMITTER_IDENT emails
//     must both be in the allowlist (case-insensitive)
const (
	allowedAuthorEmailConfigKey = "ai-skill.allowedAuthorEmail"
	allowedAuthorEmailsInfoFile = "ai-skill-allowed-author-emails"
)

// validateAuthorEmailAllowlist returns a non-empty error message when a
// local allowlist is configured and the impending commit author/committer
// email is not on it.
func validateAuthorEmailAllowlist(root string) string {
	allowed, err := loadAuthorEmailAllowlist(root)
	if err != nil {
		return "author-email-allowlist: failed to load local allowlist: " + err.Error()
	}
	if len(allowed) == 0 {
		return ""
	}

	authorIdent, err := gitVar(root, "GIT_AUTHOR_IDENT")
	if err != nil {
		return "author-email-allowlist: could not resolve GIT_AUTHOR_IDENT: " + err.Error()
	}
	committerIdent, err := gitVar(root, "GIT_COMMITTER_IDENT")
	if err != nil {
		return "author-email-allowlist: could not resolve GIT_COMMITTER_IDENT: " + err.Error()
	}

	authorEmail := emailFromGitIdent(authorIdent)
	committerEmail := emailFromGitIdent(committerIdent)
	if authorEmail == "" {
		return "author-email-allowlist: GIT_AUTHOR_IDENT has no email: " + authorIdent
	}
	if committerEmail == "" {
		return "author-email-allowlist: GIT_COMMITTER_IDENT has no email: " + committerIdent
	}

	var bad []string
	if !allowlistContains(allowed, authorEmail) {
		bad = append(bad, fmt.Sprintf("author <%s>", authorEmail))
	}
	if !allowlistContains(allowed, committerEmail) {
		bad = append(bad, fmt.Sprintf("committer <%s>", committerEmail))
	}
	if len(bad) == 0 {
		return ""
	}

	allowedList := make([]string, 0, len(allowed))
	for email := range allowed {
		allowedList = append(allowedList, email)
	}
	return fmt.Sprintf(
		"author-email-allowlist: %s not in local allowlist [%s].\n"+
			"Remediation:\n"+
			"  • Fix identity: git config --local user.email <allowed-email>\n"+
			"  • Or extend allowlist: git config --local --add %s <email>\n"+
			"  • Or edit local file: .git/info/%s (never commit this file)\n"+
			"  • Unset allowlist to disable: git config --local --unset-all %s",
		strings.Join(bad, " and "),
		strings.Join(allowedList, ", "),
		allowedAuthorEmailConfigKey,
		allowedAuthorEmailsInfoFile,
		allowedAuthorEmailConfigKey,
	)
}

func loadAuthorEmailAllowlist(root string) (map[string]struct{}, error) {
	allowed := map[string]struct{}{}

	cfgEmails, err := gitConfigGetAll(root, allowedAuthorEmailConfigKey)
	if err != nil {
		return nil, err
	}
	for _, email := range cfgEmails {
		addAllowlistEmail(allowed, email)
	}

	infoPath := filepath.Join(root, ".git", "info", allowedAuthorEmailsInfoFile)
	data, err := os.ReadFile(infoPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else {
		for _, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			addAllowlistEmail(allowed, line)
		}
	}

	return allowed, nil
}

func addAllowlistEmail(dst map[string]struct{}, raw string) {
	email := strings.ToLower(strings.TrimSpace(raw))
	email = strings.TrimPrefix(email, "<")
	email = strings.TrimSuffix(email, ">")
	email = strings.TrimSpace(email)
	if email == "" || !strings.Contains(email, "@") {
		return
	}
	dst[email] = struct{}{}
}

func allowlistContains(allowed map[string]struct{}, email string) bool {
	_, ok := allowed[strings.ToLower(strings.TrimSpace(email))]
	return ok
}

// emailFromGitIdent parses `Name <email> unix-seconds tz` from `git var`.
func emailFromGitIdent(ident string) string {
	ident = strings.TrimSpace(ident)
	start := strings.LastIndex(ident, "<")
	end := strings.LastIndex(ident, ">")
	if start < 0 || end <= start {
		return ""
	}
	return strings.TrimSpace(ident[start+1 : end])
}

func gitVar(root, name string) (string, error) {
	output, err := exec.Command("git", "-C", root, "var", name).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func gitConfigGetAll(root, key string) ([]string, error) {
	output, err := exec.Command("git", "-C", root, "config", "--local", "--get-all", key).CombinedOutput()
	if err != nil {
		// git exits 1 when the key is unset — treat as empty, not failure.
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return nil, nil
		}
		return nil, fmt.Errorf("%s: %s", err.Error(), strings.TrimSpace(string(output)))
	}
	var values []string
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			values = append(values, line)
		}
	}
	return values, nil
}
