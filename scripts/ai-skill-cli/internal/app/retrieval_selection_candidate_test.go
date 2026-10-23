package app

import (
	"database/sql"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestRetrievalSelectionPhraseAndFilters(t *testing.T) {
	path := filepath.Join(t.TempDir(), "selection.sqlite")
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE atoms (id TEXT, source_path TEXT, layer TEXT, type TEXT, status TEXT, priority TEXT, confidence TEXT, context_cost TEXT, summary TEXT);
 CREATE VIRTUAL TABLE fts USING fts5(id, source_path, title, summary, tags, when_to_read, validation_signal);`)
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range []struct{ id, title, hint, status string }{
		{"exact", "quartz lamp", "", "deprecated"},
		{"loose", "quartz other lamp", "", "validated"},
		{"hint-only", "other", strings.Repeat("quartz lamp ", 40), "validated"},
	} {
		if _, err = db.Exec(`INSERT INTO atoms VALUES (?,?,'knowledge','summary',?,'P2','medium','low','summary')`, row.id, row.id+".md", row.status); err != nil {
			t.Fatal(err)
		}
		if _, err = db.Exec(`INSERT INTO fts VALUES (?, ?, ?, 'summary', '', '', ?)`, row.id, row.id+".md", row.title, row.hint); err != nil {
			t.Fatal(err)
		}
	}
	exact, err := runRuntimeIndexQuery(path, "quartz lamp", runtimeOptions{limit: 5})
	if err != nil {
		t.Fatal(err)
	}
	if len(exact) != 1 || exact[0].ID != "exact" || !strings.HasPrefix(exact[0].MatchReason, "exact phrase") {
		t.Fatal("phrase or content weighting", exact)
	}
	filtered, err := runRuntimeIndexQuery(path, "quartz lamp", runtimeOptions{limit: 5, statusFilter: "validated"})
	if err != nil {
		t.Fatal(err)
	}
	if len(filtered) != 1 || filtered[0].Status != "validated" {
		t.Fatal("filter lost", filtered)
	}
	relaxed, err := runtimeIndexSelectionCandidates(path, "quartz absent", runtimeOptions{limit: 5})
	if err != nil {
		t.Fatal(err)
	}
	if len(relaxed) < 2 || !strings.HasPrefix(relaxed[0].MatchReason, "lexical fallback") {
		t.Fatal("fallback missing", relaxed)
	}
	absent, err := runRuntimeIndexQuery(path, "nonce-unfindable-xx", runtimeOptions{limit: 5})
	if err != nil || len(absent) != 0 {
		t.Fatal("no-answer result", absent, err)
	}
}

func TestRetrievalSelectionBudgetPreservesIdentity(t *testing.T) {
	candidate := QueryResult{ID: "identity", SourcePath: "knowledge/example.md", Status: "deprecated", Summary: strings.Repeat("中文😀<&>\"", 1000), MatchReason: "exact phrase; applicability unverified"}
	result := runtimeBoundRetrievalResults([]QueryResult{candidate, candidate}, 5, 2000)
	if len(result) != 1 || result[0].ID != candidate.ID || result[0].SourcePath != candidate.SourcePath || result[0].Status != candidate.Status {
		t.Fatal("identity changed", result)
	}
	payload, err := json.Marshal(result)
	if err != nil || !utf8.Valid(payload) || utf8.RuneCount(payload) > 2000 || !strings.HasSuffix(result[0].Summary, "…") {
		t.Fatal("invalid budget output", len(payload), err)
	}
	candidate.SourcePath = strings.Repeat("x", 3000)
	if got := runtimeBoundRetrievalResults([]QueryResult{candidate}, 5, 2000); len(got) != 0 {
		t.Fatal("oversized identity leaked")
	}
	if got := runtimeBoundRetrievalResults(result, 0, 2000); len(got) != 0 {
		t.Fatal("zero limit ignored")
	}
}
