package app

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTranscriptLineReaderMatchesScanLines(t *testing.T) {
	long := strings.Repeat("z", 3*1024*1024)
	r := newTranscriptLineReader(strings.NewReader("a\r\n\nb\n" + long + "\nlast-no-newline"))
	var got []string
	for r.Scan() {
		got = append(got, r.Text())
	}
	want := []string{"a", "", "b", long, "last-no-newline"}
	if len(got) != len(want) {
		t.Fatalf("got %d lines, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("line %d: got %q (len %d), want len %d", i, got[i][:min(len(got[i]), 20)], len(got[i]), len(want[i]))
		}
	}
	if newTranscriptLineReader(strings.NewReader("")).Scan() {
		t.Fatal("empty input must yield no lines")
	}
}

// Regression: every transcript reader used a bufio.Scanner capped at 2 MB (one
// at 8 MB) and silently ignored all lines after a longer one. Real
// transcripts had lines of 2.2 MB and 7.7 MB.
func TestTranscriptReadersSeeLinesAfterAnOversizeLine(t *testing.T) {
	dir := t.TempDir()
	base := writeRuntimeReceiptTranscript(t, dir,
		"scripts/ai-skill-cli/bin/ai-skill-darwin-arm64 runtime receipt --repo .", runtimeReceiptOK, false, canonicalBootstrapReads)
	huge, _ := json.Marshal(map[string]any{"type": "user", "message": map[string]any{"content": []map[string]any{
		{"type": "tool_result", "tool_use_id": "big", "content": strings.Repeat("q", 9*1024*1024)},
	}}})
	ack, _ := json.Marshal(map[string]any{"type": "assistant", "message": map[string]any{"content": []map[string]any{
		{"type": "text", "text": "Bootstrap: rules=✓ phase=phase.bootstrap obligations=2 gates=2"},
	}}})
	toolUse, _ := json.Marshal(map[string]any{"type": "assistant", "message": map[string]any{"content": []map[string]any{
		{"type": "tool_use", "id": "t9", "name": "Bash", "input": map[string]any{"command": "ls"}},
	}}})
	path := filepath.Join(dir, "with-oversize.jsonl")
	writeFile(t, path, string(huge)+"\n"+mustReadTranscript(t, base)+string(ack)+"\n"+string(toolUse)+"\n")

	msgs := extractTranscriptMessages(path)
	if len(msgs) == 0 || !strings.Contains(msgs[len(msgs)-1].Text, "Bootstrap:") {
		t.Fatalf("extractTranscriptMessages missed lines after the oversize line: %#v", msgs)
	}
	found := false
	for _, s := range extractAssistantTexts(path) {
		found = found || strings.Contains(s, "Bootstrap:")
	}
	if !found {
		t.Fatal("extractAssistantTexts missed lines after the oversize line")
	}
	if ok, missing := transcriptHasRequiredBootstrapReads(path, bootstrapRequiredReadSuffixes); !ok {
		t.Fatalf("transcriptHasRequiredBootstrapReads missed reads after the oversize line: %v", missing)
	}
	if !transcriptHasRuntimeReceiptEvidence(path) {
		t.Fatal("transcriptHasRuntimeReceiptEvidence missed the receipt after the oversize line")
	}
	if !transcriptHasBootstrapAcknowledgment(path, 20) {
		t.Fatal("transcriptHasBootstrapAcknowledgment missed the receipt after the oversize line")
	}
	if !transcriptLatestAssistantHasToolUse(path) {
		t.Fatal("transcriptLatestAssistantHasToolUse missed the tool_use after the oversize line")
	}
}

func mustReadTranscript(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
