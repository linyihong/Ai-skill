package app

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func incTestEnv(t *testing.T) string {
	t.Helper()
	setHookTempDir(t, t.TempDir()) // isolates state file and shadow ledger
	return filepath.Join(t.TempDir(), "session.jsonl")
}

func incUserLine(text string) string {
	b, _ := json.Marshal(map[string]interface{}{"type": "user", "message": map[string]interface{}{"role": "user", "content": text}})
	return string(b) + "\n"
}

func incAssistantLine(text string) string {
	b, _ := json.Marshal(map[string]interface{}{"type": "assistant", "message": map[string]interface{}{"role": "assistant", "content": []map[string]string{{"type": "text", "text": text}}}})
	return string(b) + "\n"
}

func incWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func incAppend(t *testing.T, path, content string) {
	t.Helper()
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		t.Fatal(err)
	}
}

// incParse runs the incremental parser to EOF and checks it against the full parse.
func incParse(t *testing.T, path string) incrementalParseResult {
	t.Helper()
	res, err := parseIncrementalTranscript(path, fileSizeOrNegative(path))
	if err != nil {
		t.Fatalf("incremental parse: %v", err)
	}
	full := extractTranscriptMessages(path)
	if !reflect.DeepEqual(normalizeDetectorMessages(full), normalizeDetectorMessages(res.Messages)) {
		t.Fatalf("incremental != full\nfull=%#v\ninc=%#v", full, res.Messages)
	}
	return res
}

func incBigPrefix(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteString(incUserLine(fmt.Sprintf("message %03d %s", i, strings.Repeat("x", 200))))
	}
	return b.String()
}

func TestIncrementalTranscriptResumesOnAppend(t *testing.T) {
	path := incTestEnv(t)
	incWrite(t, path, incUserLine("first")+incAssistantLine("reply"))
	if r := incParse(t, path); r.Mode != "rebuild" || r.Reason != "no_state" {
		t.Fatalf("first parse should rebuild with no_state, got %s/%s", r.Mode, r.Reason)
	}
	added := incUserLine("second") + `{"type":"attachment","x":1}` + "\n" + incAssistantLine("reply 2")
	incAppend(t, path, added)
	r := incParse(t, path)
	if r.Mode != "incremental" || r.NewBytes != int64(len(added)) {
		t.Fatalf("expected incremental over %d new bytes, got %s new=%d reason=%s", len(added), r.Mode, r.NewBytes, r.Reason)
	}
	if len(r.Messages) != 4 {
		t.Fatalf("expected 4 messages, got %d", len(r.Messages))
	}
}

func TestIncrementalTranscriptFallbacks(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(t *testing.T, path string)
		reason string
	}{
		{"truncated", func(t *testing.T, path string) {
			incWrite(t, path, incUserLine("only"))
		}, "transcript_truncated"},
		{"overwritten head", func(t *testing.T, path string) {
			data, _ := os.ReadFile(path)
			incWrite(t, path, strings.Replace(string(data), "message 000", "MESSAGE 000", 1))
		}, "head_hash_mismatch"},
		{"processed region modified", func(t *testing.T, path string) {
			data, _ := os.ReadFile(path)
			incWrite(t, path, strings.Replace(string(data), "message 059", "MESSAGE 059", 1))
		}, "boundary_hash_mismatch"},
		{"state corrupt", func(t *testing.T, path string) {
			incWrite(t, incrementalStatePath(path), "{not json")
		}, "state_decode_failed"},
		{"state version", func(t *testing.T, path string) {
			editIncState(t, path, func(st *incrementalTranscriptState) { st.Version = 99 })
		}, "state_version_mismatch"},
		{"state for another transcript", func(t *testing.T, path string) {
			editIncState(t, path, func(st *incrementalTranscriptState) { st.Transcript = "/elsewhere.jsonl" })
		}, "transcript_path_mismatch"},
		{"cursor invalid", func(t *testing.T, path string) {
			editIncState(t, path, func(st *incrementalTranscriptState) { st.Offset = -5 })
		}, "cursor_invalid"},
		{"state removed", func(t *testing.T, path string) {
			_ = os.Remove(incrementalStatePath(path))
		}, "no_state"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			path := incTestEnv(t)
			incWrite(t, path, incBigPrefix(60)) // > 2 hash windows, so head and boundary differ
			incParse(t, path)
			tc.mutate(t, path)
			r := incParse(t, path)
			if r.Mode != "rebuild" || r.Reason != tc.reason {
				t.Fatalf("expected rebuild/%s, got %s/%s", tc.reason, r.Mode, r.Reason)
			}
			// After a rebuild the next append resumes incrementally.
			incAppend(t, path, incUserLine("after"))
			if r := incParse(t, path); r.Mode != "incremental" {
				t.Fatalf("expected incremental after rebuild, got %s/%s", r.Mode, r.Reason)
			}
		})
	}
}

func editIncState(t *testing.T, path string, f func(*incrementalTranscriptState)) {
	t.Helper()
	data, err := os.ReadFile(incrementalStatePath(path))
	if err != nil {
		t.Fatal(err)
	}
	var st incrementalTranscriptState
	if err := json.Unmarshal(data, &st); err != nil {
		t.Fatal(err)
	}
	f(&st)
	if err := saveIncrementalState(path, st); err != nil {
		t.Fatal(err)
	}
}

func TestIncrementalTranscriptUnterminatedTailAndMalformedLines(t *testing.T) {
	path := incTestEnv(t)
	tail := strings.TrimSuffix(incUserLine("tail not yet terminated"), "\n")
	incWrite(t, path, incUserLine("a")+"{broken json\n\n"+tail)
	r := incParse(t, path)
	if len(r.Messages) != 2 {
		t.Fatalf("expected the unterminated tail to be parsed for this result, got %d messages", len(r.Messages))
	}
	incAppend(t, path, "\n"+incAssistantLine("b"))
	if r := incParse(t, path); r.Mode != "incremental" || len(r.Messages) != 3 {
		t.Fatalf("expected incremental with 3 messages, got %s/%s n=%d", r.Mode, r.Reason, len(r.Messages))
	}
}

func TestIncrementalTranscriptParsesLinesOver2MB(t *testing.T) {
	path := incTestEnv(t)
	huge := `{"type":"user","message":{"role":"user","content":"` + strings.Repeat("y", 3*1024*1024) + `"}}` + "\n"
	incWrite(t, path, incUserLine("before")+huge+incUserLine("after"))
	if r := incParse(t, path); len(r.Messages) != 3 {
		t.Fatalf("expected all 3 messages, including those after the long line, got %d", len(r.Messages))
	}
	incAppend(t, path, incUserLine("later"))
	if r := incParse(t, path); r.Mode != "incremental" || len(r.Messages) != 4 {
		t.Fatalf("expected incremental with 4 messages, got %s/%s n=%d", r.Mode, r.Reason, len(r.Messages))
	}
}

func readShadowLedger(t *testing.T) []incrementalShadowRecord {
	t.Helper()
	f, err := os.Open(incrementalShadowLedgerPath())
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var out []incrementalShadowRecord
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var rec incrementalShadowRecord
		if err := json.Unmarshal(sc.Bytes(), &rec); err != nil {
			t.Fatalf("ledger line: %v", err)
		}
		if strings.Contains(sc.Text(), "secret-transcript-text") {
			t.Fatal("shadow ledger must not contain transcript text")
		}
		out = append(out, rec)
	}
	return out
}

func TestIncrementalShadowLedgerOutcomes(t *testing.T) {
	path := incTestEnv(t)
	incWrite(t, path, incUserLine("secret-transcript-text"))
	size := fileSizeOrNegative(path)
	full := extractTranscriptMessagesUpTo(path, size)

	if rec := runIncrementalShadow(path, full, size); rec.Outcome != "equal" || rec.Mode != "rebuild" {
		t.Fatalf("expected equal/rebuild, got %+v", rec)
	}
	// Lines appended after the snapshot (while the hook runs) are outside both
	// parses, so the comparison still counts.
	incAppend(t, path, incUserLine("appended during the hook"))
	if rec := runIncrementalShadow(path, full, size); rec.Outcome != "equal" {
		t.Fatalf("expected equal over the snapshot when the transcript grew, got %+v", rec)
	}
	if rec := runIncrementalShadow(path, full, fileSizeOrNegative(path)+10); rec.Outcome != "not_counted" || rec.Note != "transcript shrank during compare" {
		t.Fatalf("expected not_counted when the file is shorter than the snapshot, got %+v", rec)
	}
	if rec := runIncrementalShadow(path, full, -1); rec.Outcome != "not_counted" {
		t.Fatalf("expected not_counted when the size is unavailable, got %+v", rec)
	}
	// A differing authoritative result is reported, not repaired.
	fake := []DetectorMessage{{Role: "user", Text: "something else"}}
	if rec := runIncrementalShadow(path, fake, size); rec.Outcome != "mismatch" || rec.FirstDiff != 0 {
		t.Fatalf("expected mismatch at 0, got %+v", rec)
	}
	t.Setenv(incrementalShadowDisableEnv, "0")
	if rec := runIncrementalShadow(path, full, size); rec.Outcome != "not_counted" || rec.Note != "disabled" {
		t.Fatalf("expected disabled, got %+v", rec)
	}
	if got := len(readShadowLedger(t)); got != 6 {
		t.Fatalf("expected 6 ledger records, got %d", got)
	}
	info, err := os.Stat(incrementalStatePath(path))
	if err != nil {
		t.Fatal(err)
	}
	// Windows has no Unix permission bits (Go reports 0666); the per-user temp
	// directory's ACL protects the file there.
	if runtime.GOOS != "windows" && info.Mode().Perm() != 0o600 {
		t.Fatalf("state file must be 0600, got %v", info.Mode().Perm())
	}
}

// The shadow must not feed decisions: loadPreToolUseRouting's transcript is the
// full parse even when the incremental state is poisoned.
func TestIncrementalShadowCannotChangeRoutingInput(t *testing.T) {
	path := incTestEnv(t)
	incWrite(t, path, incUserLine("real"))
	incParse(t, path)
	editIncState(t, path, func(st *incrementalTranscriptState) {
		st.Messages = []DetectorMessage{{Role: "user", Text: "poisoned"}}
	})
	r := loadPreToolUseRouting(path, t.TempDir()) // no registry → not ok, but parsed
	if r.parsed {
		t.Fatal("routing without a registry should not parse")
	}
	full := extractTranscriptMessages(path)
	if len(full) != 1 || full[0].Text != "real" {
		t.Fatalf("authoritative parse changed: %#v", full)
	}
	rec := runIncrementalShadow(path, full, fileSizeOrNegative(path))
	if rec.Outcome != "mismatch" {
		t.Fatalf("poisoned incremental state must surface as a mismatch, got %+v", rec)
	}
}

// setHookTempDir redirects os.TempDir (and so hookRuntimePath) for one test.
// os.TempDir reads TMPDIR on Unix but TMP / TEMP on Windows, so all three are set.
func setHookTempDir(t *testing.T, dir string) {
	t.Helper()
	for _, k := range []string{"TMPDIR", "TMP", "TEMP"} {
		t.Setenv(k, dir)
	}
	if got := os.TempDir(); got != dir {
		t.Fatalf("os.TempDir() = %q, want %q", got, dir)
	}
}

func TestIncrementalShadowSnapshotParseIgnoresLaterAppends(t *testing.T) {
	path := incTestEnv(t)
	incWrite(t, path, incUserLine("before snapshot"))
	size := fileSizeOrNegative(path)
	incAppend(t, path, incUserLine("after snapshot"))
	if got := extractTranscriptMessagesUpTo(path, size); len(got) != 1 {
		t.Fatalf("expected only the message inside the snapshot, got %#v", got)
	}
	if got := extractTranscriptMessages(path); len(got) != 2 {
		t.Fatalf("unbounded parse should still read everything, got %d", len(got))
	}
}

func incRawLine(v map[string]any) string {
	b, _ := json.Marshal(v)
	return string(b) + "\n"
}

func incToolUseRow(msgID, toolID string) string {
	return incRawLine(map[string]any{"type": "assistant", "message": map[string]any{"id": msgID, "content": []map[string]any{
		{"type": "tool_use", "id": toolID, "name": "Bash", "input": map[string]any{"command": "ls"}},
	}}})
}

func TestIncrementalShadowRecordsCoverageFields(t *testing.T) {
	path := incTestEnv(t)
	long := incUserLine(strings.Repeat("L", incrementalLongLine))
	incWrite(t, path,
		incUserLine("start")+
			incRawLine(map[string]any{"type": "system", "subtype": "compact_boundary", "content": "Conversation compacted"})+
			"{not json\n"+
			long+
			incToolUseRow("msg_a", "t1")+incToolUseRow("msg_a", "t2")+incToolUseRow("msg_a", "t3")+
			incToolUseRow("msg_b", "t4")+
			strings.TrimSuffix(incUserLine("unfinished"), "\n"))
	size := fileSizeOrNegative(path)
	rec := runIncrementalShadow(path, extractTranscriptMessagesUpTo(path, size), size)
	if rec.Outcome != "equal" {
		t.Fatalf("expected equal, got %+v", rec)
	}
	if rec.CompactBoundaries != 1 || rec.UndecodableLines != 1 || rec.LongLines != 1 || rec.MaxParallelToolUse != 3 || !rec.UnterminatedTail {
		t.Fatalf("coverage fields wrong: %+v", rec)
	}
	if rec.ConcurrentHooks != 0 {
		t.Fatalf("expected no concurrent hooks, got %d", rec.ConcurrentHooks)
	}

	// Stats carry across incremental calls; the unfinished line is counted only
	// once it is terminated and consumed.
	incAppend(t, path, "\n"+incToolUseRow("msg_c", "t5")+incToolUseRow("msg_c", "t6"))
	size = fileSizeOrNegative(path)
	rec = runIncrementalShadow(path, extractTranscriptMessagesUpTo(path, size), size)
	if rec.Mode != "incremental" || rec.CompactBoundaries != 1 || rec.UndecodableLines != 1 || rec.MaxParallelToolUse != 3 || rec.UnterminatedTail {
		t.Fatalf("expected carried stats on incremental call, got %+v", rec)
	}
}

func TestIncrementalShadowCountsConcurrentRuns(t *testing.T) {
	path := incTestEnv(t)
	incWrite(t, path, incUserLine("x"))
	prefix := "ai-skill-transcript-incremental-" + md5Short(path) + ".inflight-"
	fresh := hookRuntimePath(prefix + "111-1")
	stale := hookRuntimePath(prefix + "222-2")
	incWrite(t, fresh, "")
	incWrite(t, stale, "")
	old := time.Now().Add(-2 * incrementalInflightTTL)
	if err := os.Chtimes(stale, old, old); err != nil {
		t.Fatal(err)
	}
	size := fileSizeOrNegative(path)
	rec := runIncrementalShadow(path, extractTranscriptMessagesUpTo(path, size), size)
	if rec.ConcurrentHooks != 1 {
		t.Fatalf("expected 1 concurrent run (stale marker ignored), got %d", rec.ConcurrentHooks)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatal("stale in-flight marker should be removed")
	}
	matches, _ := filepath.Glob(hookRuntimePath(prefix + "*"))
	if len(matches) != 1 || matches[0] != fresh {
		t.Fatalf("this run's marker should be removed after the run, left %v", matches)
	}
}
