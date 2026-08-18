package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

// TestIsAiSkillBootstrapCLICommand covers the narrow allowlist that breaks the
// gate.bootstrap.receipt_present deadlock: only this repo's own repo-local
// binary running a read-only bootstrap subcommand qualifies.
func TestIsAiSkillBootstrapCLICommand(t *testing.T) {
	allowed := []string{
		`./scripts/ai-skill-cli/bin/ai-skill-windows-amd64.exe runtime receipt --repo C:/repo`,
		`/repo/scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime obligations --repo /repo`,
		`scripts/ai-skill-cli/bin/ai-skill-linux-amd64 runtime   receipt`,
		`"C:\repo\scripts\ai-skill-cli\bin\ai-skill-windows-amd64.exe" hooks run commit-msg`,
		`AI-SKILL-LINUX-ARM64 RUNTIME RECEIPT`,
	}
	for _, cmd := range allowed {
		if !isAiSkillBootstrapCLICommand(cmd) {
			t.Errorf("expected allowlisted bootstrap CLI command, got reject: %q", cmd)
		}
	}

	rejected := []string{
		"",
		`git status --short --branch`,
		// right binary, non-bootstrap subcommand
		`./scripts/ai-skill-cli/bin/ai-skill-linux-amd64 plans validate`,
		`./scripts/ai-skill-cli/bin/ai-skill-windows-amd64.exe runtime compile`,
		// bootstrap subcommand, but not the repo-local platform binary
		`ai-skill runtime receipt --repo /repo`,
		`ai-skill-windows-386.exe runtime receipt`,
		// "hooks" without "run"
		`./scripts/ai-skill-cli/bin/ai-skill-linux-amd64 hooks install`,
		// substring of a longer word must not match `hooks run`
		`./scripts/ai-skill-cli/bin/ai-skill-linux-amd64 hooks runtime-thing`,
	}
	for _, cmd := range rejected {
		if isAiSkillBootstrapCLICommand(cmd) {
			t.Errorf("expected non-allowlisted command to be rejected: %q", cmd)
		}
	}
}

func TestBashCommandFromPreToolUsePayload(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		want    string
	}{
		{"tool_input", `{"tool_input":{"command":"git status"}}`, "git status"},
		{"arguments", `{"arguments":{"command":"ls"}}`, "ls"},
		{"input", `{"input":{"command":"pwd"}}`, "pwd"},
		{"no command key", `{"tool_input":{"file_path":"/x"}}`, ""},
		{"no payload key", `{"tool_name":"Bash"}`, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var payload map[string]json.RawMessage
			if err := json.Unmarshal([]byte(tc.payload), &payload); err != nil {
				t.Fatalf("unmarshal payload: %v", err)
			}
			if got := bashCommandFromPreToolUsePayload(payload); got != tc.want {
				t.Fatalf("bashCommandFromPreToolUsePayload = %q, want %q", got, tc.want)
			}
		})
	}
}

// TestPreToolUseHookAllowsBootstrapCLIQuery is the regression test for the
// deadlock: a transcript with no Receipt and no bootstrap Reads must still let
// the `runtime receipt` CLI call through, because that call is the only way to
// obtain the Receipt's numbers.
func TestPreToolUseHookAllowsBootstrapCLIQuery(t *testing.T) {
	dir := t.TempDir()
	tr := writeBootstrapTranscript(t, dir, "No receipt yet.", nil)
	cmd := `./scripts/ai-skill-cli/bin/ai-skill-windows-amd64.exe runtime receipt --repo C:/repo`
	payload := fmt.Sprintf(`{"tool_name":"Bash","transcript_path":%q,"tool_input":{"command":%q}}`, tr, cmd)
	setHookStdin(t, payload)

	var stdout, stderr bytes.Buffer
	code := runPreToolUseHook(dir, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("expected ExitSuccess; got %d; stderr=%s", code, stderr.String())
	}
	if !strings.Contains(stderr.String(), "ALLOW_BOOTSTRAP_CLI_QUERY") {
		t.Fatalf("expected ALLOW_BOOTSTRAP_CLI_QUERY in stderr; got:\n%s", stderr.String())
	}
	if strings.Contains(stdout.String(), "deny") {
		t.Fatalf("expected no deny decision on stdout; got:\n%s", stdout.String())
	}
}

// The exemption must stay narrow: an ordinary Bash command with the same
// missing-Receipt transcript is still blocked.
func TestPreToolUseHookStillBlocksNonBootstrapBash(t *testing.T) {
	dir := t.TempDir()
	tr := writeBootstrapTranscript(t, dir, "No receipt yet.", nil)
	payload := fmt.Sprintf(`{"tool_name":"Bash","transcript_path":%q,"tool_input":{"command":"git commit -m x"}}`, tr)
	setHookStdin(t, payload)

	var stdout, stderr bytes.Buffer
	code := runPreToolUseHook(dir, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("expected ExitSuccess (deny carried by JSON); got %d; stderr=%s", code, stderr.String())
	}
	assertPreToolUseDeny(t, stdout.String())
	if !strings.Contains(stderr.String(), "BLOCK_NO_RECEIPT") {
		t.Fatalf("expected BLOCK_NO_RECEIPT in stderr; got:\n%s", stderr.String())
	}
}
