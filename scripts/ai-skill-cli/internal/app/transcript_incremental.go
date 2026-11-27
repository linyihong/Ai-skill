package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

// Incremental transcript parsing — SHADOW ONLY.
//
// Plan: plans/active/2026-11-25-0842-verification-architecture-contract
// (Phase 2a step 1). The full parse (extractTranscriptMessages) stays the only
// authoritative source for routing, gates and advisories. This file keeps its
// own cursor state across PreToolUse calls and is compared against the full
// parse; its result is never returned to a decision path.
//
// Independence: the incremental path opens the file itself and keeps its own
// state. It shares only the pure per-line text extractor (extractEntryText),
// so the shadow verifies the cursor, boundary and fallback mechanism — not the
// per-line extraction logic.

const (
	// Version 2: lines of any length are parsed (the full parse no longer
	// stops at lines over 2 MB), so version-1 state that stalled is rebuilt.
	// Version 3: state carries input statistics for the shadow ledger.
	incrementalTranscriptStateVersion = 3
	// incrementalLongLine is the old bufio.Scanner limit; lines at or above it
	// are counted so the ledger shows the long-line case was exercised.
	incrementalLongLine = 2 * 1024 * 1024
	// incrementalInflightTTL bounds how long an in-flight marker left by a
	// crashed hook can be mistaken for a concurrent run.
	incrementalInflightTTL = 60 * time.Second
	incrementalHashWindow  = 4096

	incrementalShadowDisableEnv = "AI_SKILL_PRETOOLUSE_INCREMENTAL_SHADOW"
)

type incrementalTranscriptState struct {
	Version      int                   `json:"version"`
	Transcript   string                `json:"transcript"`
	Offset       int64                 `json:"offset"`
	HeadHash     string                `json:"head_hash"`
	BoundaryHash string                `json:"boundary_hash"`
	Messages     []DetectorMessage     `json:"messages"`
	Stats        incrementalInputStats `json:"stats"`
}

// incrementalInputStats counts input shapes seen in the consumed part of the
// transcript, so the shadow ledger can show which evidence conditions were
// exercised. Counts only — never text. They are metadata, not compared.
type incrementalInputStats struct {
	CompactBoundaries   int    `json:"compact_boundaries"`
	UndecodableLines    int    `json:"undecodable_lines"`
	LongLines           int    `json:"long_lines"`
	MaxParallelToolUse  int    `json:"max_parallel_tool_use"`
	LastMessageID       string `json:"last_message_id"`
	LastMessageToolUses int    `json:"last_message_tool_uses"`
}

// incrementalParseResult describes one incremental parse.
type incrementalParseResult struct {
	Messages []DetectorMessage
	Mode     string // "incremental" or "rebuild"
	Reason   string // why a rebuild happened ("" for incremental)
	NewBytes int64
	Stats    incrementalInputStats
	// UnterminatedTail reports that the parsed range ended inside a line.
	UnterminatedTail bool
}

func incrementalStatePath(transcriptPath string) string {
	return hookRuntimePath("ai-skill-transcript-incremental-" + md5Short(transcriptPath) + ".json")
}

func hashFileRange(f *os.File, start, end int64) (string, error) {
	if end <= start {
		return "", nil
	}
	h := sha256.New()
	if _, err := io.Copy(h, io.NewSectionReader(f, start, end-start)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func hashHead(f *os.File, offset int64) (string, error) {
	start, end := headRange(offset)
	return hashFileRange(f, start, end)
}

func hashBoundary(f *os.File, offset int64) (string, error) {
	start, end := boundaryRange(offset)
	return hashFileRange(f, start, end)
}

func headRange(offset int64) (int64, int64) {
	end := offset
	if end > incrementalHashWindow {
		end = incrementalHashWindow
	}
	return 0, end
}

func boundaryRange(offset int64) (int64, int64) {
	start := offset - incrementalHashWindow
	if start < 0 {
		start = 0
	}
	return start, offset
}

// parseIncrementalTranscript parses transcriptPath up to limit bytes, resuming
// from the state saved by the previous call when that state can be proven to
// still describe the file's prefix. Any doubt rebuilds from byte 0.
func parseIncrementalTranscript(transcriptPath string, limit int64) (incrementalParseResult, error) {
	f, err := os.Open(transcriptPath)
	if err != nil {
		return incrementalParseResult{}, err
	}
	defer f.Close()

	state, reason := loadValidIncrementalState(f, transcriptPath, limit)
	res := incrementalParseResult{Mode: "incremental"}
	if reason != "" {
		state = incrementalTranscriptState{Version: incrementalTranscriptStateVersion, Transcript: transcriptPath}
		res.Mode, res.Reason = "rebuild", reason
	}

	if limit > state.Offset {
		buf := make([]byte, limit-state.Offset)
		if _, err := f.ReadAt(buf, state.Offset); err != nil && err != io.EOF {
			return incrementalParseResult{}, err
		}
		res.NewBytes = int64(len(buf))
		consumed, tail := consumeTranscriptLines(buf, &state.Messages, &state.Stats)
		state.Offset += consumed
		// A final line without its newline is parsed for this result only
		// (the full parse returns it too) but the cursor stays before it.
		if tail != nil {
			res.UnterminatedTail = true
			res.Messages = appendTranscriptLine(append([]DetectorMessage(nil), state.Messages...), tail, nil)
		}
	}
	if res.Messages == nil {
		res.Messages = state.Messages
	}
	res.Stats = state.Stats

	if state.HeadHash, err = hashHead(f, state.Offset); err != nil {
		return res, err
	}
	if state.BoundaryHash, err = hashBoundary(f, state.Offset); err != nil {
		return res, err
	}
	_ = saveIncrementalState(transcriptPath, state)
	return res, nil
}

// loadValidIncrementalState returns the saved state, or a non-empty reason
// explaining why it cannot be trusted.
func loadValidIncrementalState(f *os.File, transcriptPath string, limit int64) (incrementalTranscriptState, string) {
	data, err := os.ReadFile(incrementalStatePath(transcriptPath))
	if err != nil {
		return incrementalTranscriptState{}, "no_state"
	}
	var st incrementalTranscriptState
	if err := json.Unmarshal(data, &st); err != nil {
		return incrementalTranscriptState{}, "state_decode_failed"
	}
	switch {
	case st.Version != incrementalTranscriptStateVersion:
		return st, "state_version_mismatch"
	case st.Transcript != transcriptPath:
		return st, "transcript_path_mismatch"
	case st.Offset < 0:
		return st, "cursor_invalid"
	case limit < st.Offset:
		return st, "transcript_truncated"
	}
	head, err := hashHead(f, st.Offset)
	if err != nil {
		return st, "cursor_unverifiable"
	}
	if head != st.HeadHash {
		return st, "head_hash_mismatch"
	}
	boundary, err := hashBoundary(f, st.Offset)
	if err != nil {
		return st, "cursor_unverifiable"
	}
	if boundary != st.BoundaryHash {
		return st, "boundary_hash_mismatch"
	}
	return st, ""
}

// consumeTranscriptLines appends messages for every newline-terminated line in
// buf. It returns the bytes consumed and the unterminated tail (nil if none).
func consumeTranscriptLines(buf []byte, msgs *[]DetectorMessage, st *incrementalInputStats) (int64, []byte) {
	var consumed int64
	for len(buf) > 0 {
		i := bytes.IndexByte(buf, '\n')
		if i < 0 {
			return consumed, buf
		}
		*msgs = appendTranscriptLine(*msgs, buf[:i], st)
		consumed += int64(i + 1)
		buf = buf[i+1:]
	}
	return consumed, nil
}

// appendTranscriptLine appends the message for one line. st, when non-nil,
// accumulates input statistics for a consumed line (the unterminated tail is
// parsed with st == nil because the cursor does not move past it).
func appendTranscriptLine(msgs []DetectorMessage, raw []byte, st *incrementalInputStats) []DetectorMessage {
	if st != nil && len(raw) >= incrementalLongLine {
		st.LongLines++
	}
	line := strings.TrimSpace(strings.TrimSuffix(string(raw), "\r"))
	if line == "" {
		return msgs
	}
	var entry map[string]json.RawMessage
	if json.Unmarshal([]byte(line), &entry) != nil {
		if st != nil {
			st.UndecodableLines++
		}
		return msgs
	}
	roleField := entry["type"]
	if roleField == nil {
		roleField = entry["role"]
	}
	var role string
	if roleField != nil {
		_ = json.Unmarshal(roleField, &role)
	}
	if st != nil {
		recordTranscriptLineStats(st, role, entry)
	}
	if role != "user" && role != "assistant" {
		return msgs
	}
	if text := extractEntryText(entry); text != "" {
		msgs = append(msgs, DetectorMessage{Role: role, Text: text})
	}
	return msgs
}

func recordTranscriptLineStats(st *incrementalInputStats, role string, entry map[string]json.RawMessage) {
	switch role {
	case "system":
		var subtype string
		if raw, ok := entry["subtype"]; ok && json.Unmarshal(raw, &subtype) == nil && subtype == "compact_boundary" {
			st.CompactBoundaries++
		}
	case "assistant":
		// Claude Code writes each tool_use block of one assistant message as its
		// own row sharing message.id; several blocks mean parallel tool calls.
		var msg struct {
			ID      string `json:"id"`
			Content []struct {
				Type string `json:"type"`
			} `json:"content"`
		}
		raw, ok := entry["message"]
		if !ok || json.Unmarshal(raw, &msg) != nil {
			return
		}
		n := 0
		for _, c := range msg.Content {
			if c.Type == "tool_use" {
				n++
			}
		}
		if n == 0 {
			return
		}
		if msg.ID == "" || msg.ID != st.LastMessageID {
			st.LastMessageID, st.LastMessageToolUses = msg.ID, 0
		}
		st.LastMessageToolUses += n
		if st.LastMessageToolUses > st.MaxParallelToolUse {
			st.MaxParallelToolUse = st.LastMessageToolUses
		}
	}
}

func saveIncrementalState(transcriptPath string, st incrementalTranscriptState) error {
	data, err := json.Marshal(st)
	if err != nil {
		return err
	}
	path := incrementalStatePath(transcriptPath)
	tmp, err := os.CreateTemp(filepath.Dir(path), filepath.Base(path)+".tmp-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return err
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpName)
		return err
	}
	return os.Rename(tmpName, path)
}

// incrementalShadowRecord is one line of the shadow ledger. It carries counts,
// sizes and timings only — never transcript text.
type incrementalShadowRecord struct {
	Time          string  `json:"time"`
	Session       string  `json:"session"` // md5 of the transcript path
	Bytes         int64   `json:"bytes"`
	Outcome       string  `json:"outcome"` // equal | mismatch | not_counted | error
	Mode          string  `json:"mode,omitempty"`
	Reason        string  `json:"reason,omitempty"`
	NewBytes      int64   `json:"new_bytes,omitempty"`
	FullCount     int     `json:"full_count"`
	IncCount      int     `json:"inc_count"`
	FirstDiff     int     `json:"first_diff,omitempty"` // index of the first differing message, -1 if only lengths differ
	IncrementalMS float64 `json:"incremental_ms"`
	// Evidence-coverage fields (counts over the consumed transcript, no text).
	CompactBoundaries  int    `json:"compact_boundaries"`
	MaxParallelToolUse int    `json:"max_parallel_tool_use"`
	ConcurrentHooks    int    `json:"concurrent_hooks"` // other shadow runs on this transcript in flight at start
	UndecodableLines   int    `json:"undecodable_lines"`
	LongLines          int    `json:"long_lines"`
	UnterminatedTail   bool   `json:"unterminated_tail,omitempty"`
	Note               string `json:"note,omitempty"`
}

func incrementalShadowLedgerPath() string {
	return hookRuntimePath("ai-skill-pretooluse-incremental-shadow.jsonl")
}

// runIncrementalShadow compares the incremental parse with the authoritative
// full parse over the same bytes: snapshot is the size the full parse was
// limited to, and the incremental parse is limited to it too, so lines appended
// while the hook runs no longer make the comparison uncountable. Only a size
// that is unavailable, or a file that shrank below the snapshot, is recorded
// as not counted. It never returns anything a decision path could use, and it
// recovers from panics.
func runIncrementalShadow(transcriptPath string, full []DetectorMessage, snapshot int64) (rec incrementalShadowRecord) {
	rec = incrementalShadowRecord{
		Time:      time.Now().UTC().Format(time.RFC3339),
		Session:   md5Short(transcriptPath),
		Bytes:     snapshot,
		FullCount: len(full),
	}
	defer func() {
		if r := recover(); r != nil {
			rec.Outcome, rec.Note = "error", fmt.Sprintf("panic: %v", r)
		}
		appendIncrementalShadowRecord(rec)
	}()
	if os.Getenv(incrementalShadowDisableEnv) == "0" {
		rec.Outcome, rec.Note = "not_counted", "disabled"
		return rec
	}
	if snapshot < 0 {
		rec.Outcome, rec.Note = "not_counted", "transcript size unavailable"
		return rec
	}
	if now := fileSizeOrNegative(transcriptPath); now < snapshot {
		rec.Outcome, rec.Note = "not_counted", "transcript shrank during compare"
		return rec
	}
	var release func()
	rec.ConcurrentHooks, release = enterIncrementalInflight(transcriptPath)
	defer release()
	start := time.Now()
	res, err := parseIncrementalTranscript(transcriptPath, snapshot)
	rec.IncrementalMS = float64(time.Since(start).Microseconds()) / 1000
	if err != nil {
		rec.Outcome, rec.Note = "error", err.Error()
		return rec
	}
	rec.Mode, rec.Reason, rec.NewBytes = res.Mode, res.Reason, res.NewBytes
	rec.CompactBoundaries = res.Stats.CompactBoundaries
	rec.MaxParallelToolUse = res.Stats.MaxParallelToolUse
	rec.UndecodableLines = res.Stats.UndecodableLines
	rec.LongLines = res.Stats.LongLines
	rec.UnterminatedTail = res.UnterminatedTail
	rec.IncCount = len(res.Messages)
	if reflect.DeepEqual(normalizeDetectorMessages(full), normalizeDetectorMessages(res.Messages)) {
		rec.Outcome = "equal"
		return rec
	}
	rec.Outcome = "mismatch"
	rec.FirstDiff = -1
	for i := 0; i < len(full) && i < len(res.Messages); i++ {
		if full[i] != res.Messages[i] {
			rec.FirstDiff = i
			break
		}
	}
	return rec
}

// enterIncrementalInflight drops a marker for this shadow run and counts other
// fresh markers for the same transcript — concurrent PreToolUse hooks, as with
// parallel tool calls, which race on the state file. The returned func removes
// the marker.
func enterIncrementalInflight(transcriptPath string) (int, func()) {
	prefix := "ai-skill-transcript-incremental-" + md5Short(transcriptPath) + ".inflight-"
	mine := hookRuntimePath(fmt.Sprintf("%s%d-%d", prefix, os.Getpid(), time.Now().UnixNano()))
	if err := os.WriteFile(mine, nil, 0o600); err != nil {
		return 0, func() {}
	}
	others := 0
	matches, _ := filepath.Glob(hookRuntimePath(prefix + "*"))
	for _, m := range matches {
		if m == mine {
			continue
		}
		info, err := os.Stat(m)
		if err != nil {
			continue
		}
		if time.Since(info.ModTime()) > incrementalInflightTTL {
			_ = os.Remove(m) // left by a crashed run
			continue
		}
		others++
	}
	return others, func() { _ = os.Remove(mine) }
}

func normalizeDetectorMessages(m []DetectorMessage) []DetectorMessage {
	if len(m) == 0 {
		return nil
	}
	return m
}

func appendIncrementalShadowRecord(rec incrementalShadowRecord) {
	data, err := json.Marshal(rec)
	if err != nil {
		return
	}
	f, err := os.OpenFile(incrementalShadowLedgerPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return
	}
	defer f.Close()
	_, _ = f.Write(append(data, '\n'))
}

func fileSizeOrNegative(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		return -1
	}
	return info.Size()
}
