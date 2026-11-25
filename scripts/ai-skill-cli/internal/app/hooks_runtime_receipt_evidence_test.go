package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

const runtimeReceiptOK = "ai-skill runtime receipt: success (exit 0)\n" +
	"- bootstrap_receipt: ok - Bootstrap: rules=✓ phase=phase.bootstrap obligations=2 gates=2\n"

// writeRuntimeReceiptTranscript writes Read tool_use entries for readPaths,
// then a Bash tool_use running `command` and a tool_result carrying `output`.
// No assistant text is written, modelling a client that has not persisted
// mid-turn text yet.
func writeRuntimeReceiptTranscript(t *testing.T, dir, command, output string, isError bool, readPaths []string) string {
	t.Helper()
	var lines []string
	add := func(v map[string]any) {
		buf, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		lines = append(lines, string(buf))
	}
	for i, p := range readPaths {
		add(map[string]any{"type": "assistant", "message": map[string]any{"content": []map[string]any{
			{"type": "tool_use", "id": fmt.Sprintf("read_%d", i), "name": "Read", "input": map[string]any{"file_path": p}},
		}}})
	}
	add(map[string]any{"type": "assistant", "message": map[string]any{"content": []map[string]any{
		{"type": "tool_use", "id": "bash_1", "name": "Bash", "input": map[string]any{"command": command}},
	}}})
	add(map[string]any{"type": "user", "message": map[string]any{"content": []map[string]any{
		{"type": "tool_result", "tool_use_id": "bash_1", "content": output, "is_error": isError},
	}}})
	path := filepath.Join(dir, "transcript.jsonl")
	writeFile(t, path, strings.Join(lines, "\n")+"\n")
	return path
}

var canonicalBootstrapReads = []string{"/repo/CORE_BOOTSTRAP.md", "/repo/runtime/core-bootstrap.yaml"}

func runGateOn(t *testing.T, tr string) (string, string) {
	t.Helper()
	setHookStdin(t, fmt.Sprintf(`{"tool_name":"Bash","transcript_path":%q}`, tr))
	var stdout, stderr bytes.Buffer
	if code := runPreToolUseHook(t.TempDir(), &stdout, &stderr); code != ExitSuccess {
		t.Fatalf("expected ExitSuccess; got %d; stderr=%s", code, stderr.String())
	}
	return stdout.String(), stderr.String()
}

func TestPreToolUseAcceptsRuntimeReceiptResultWithReads(t *testing.T) {
	tr := writeRuntimeReceiptTranscript(t, t.TempDir(),
		"scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime receipt --repo .", runtimeReceiptOK, false, canonicalBootstrapReads)
	_, stderr := runGateOn(t, tr)
	if !strings.Contains(stderr, "ALLOW_RECEIPT_FOUND_WITH_READS") {
		t.Fatalf("expected runtime receipt result to satisfy the gate; stderr:\n%s", stderr)
	}
}

func TestPreToolUseRuntimeReceiptStillRequiresReads(t *testing.T) {
	tr := writeRuntimeReceiptTranscript(t, t.TempDir(),
		"scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime receipt --repo .", runtimeReceiptOK, false, nil)
	stdout, stderr := runGateOn(t, tr)
	assertPreToolUseDeny(t, stdout)
	if !strings.Contains(stderr, "BLOCK_RECEIPT_WITHOUT_READS") {
		t.Fatalf("expected reads to stay required; stderr:\n%s", stderr)
	}
}

func TestPreToolUseRejectsReceiptMarkerFromOtherCommands(t *testing.T) {
	cases := map[string]struct {
		command string
		output  string
		isError bool
	}{
		"grep prints marker":  {"grep -rn bootstrap_receipt hooks.go", runtimeReceiptOK, false},
		"obligations command": {"bin/ai-skill-darwin-arm64 runtime obligations --repo .", runtimeReceiptOK, false},
		"receipt errored":     {"bin/ai-skill-darwin-arm64 runtime receipt --repo .", runtimeReceiptOK, true},
		"receipt failed":      {"bin/ai-skill-darwin-arm64 runtime receipt --repo .", "ai-skill runtime receipt: failed (exit 1)\n", false},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			tr := writeRuntimeReceiptTranscript(t, t.TempDir(), c.command, c.output, c.isError, canonicalBootstrapReads)
			stdout, stderr := runGateOn(t, tr)
			assertPreToolUseDeny(t, stdout)
			if !strings.Contains(stderr, "BLOCK_NO_RECEIPT") {
				t.Fatalf("expected BLOCK_NO_RECEIPT; stderr:\n%s", stderr)
			}
		})
	}
}
