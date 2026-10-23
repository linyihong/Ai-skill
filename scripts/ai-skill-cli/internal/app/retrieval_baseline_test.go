package app

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
)

type retrievalReplayCase struct {
	ID             string   `json:"id"`
	Group          string   `json:"group"`
	Split          string   `json:"split"`
	ProvenanceKind string   `json:"provenance_kind"`
	Provenance     string   `json:"provenance"`
	ProvenanceSHA  string   `json:"provenance_sha256"`
	Query          string   `json:"query"`
	Expected       []string `json:"expected_sources"`
	Layer          string   `json:"layer"`
	Status         string   `json:"status"`
	ReviewRequired bool     `json:"review_required"`
	Excluded       []string `json:"excluded_sources"`
}
type retrievalReplaySet struct {
	Schema    int                   `json:"schema_version"`
	SuiteKind string                `json:"suite_kind"`
	Limit     int                   `json:"limit"`
	Repeats   int                   `json:"repeats"`
	Cases     []retrievalReplayCase `json:"cases"`
}
type retrievalReplayObservation struct {
	ID               string            `json:"id"`
	Split            string            `json:"split"`
	Sources          []string          `json:"sources"`
	Expected         int               `json:"expected"`
	Hits             int               `json:"hits"`
	RR               float64           `json:"reciprocal_rank"`
	MissReasons      map[string]string `json:"miss_reasons"`
	FirstMS          float64           `json:"first_invocation_ms"`
	WarmP95MS        float64           `json:"warm_cli_p95_ms"`
	PreviewEstimate  int               `json:"preview_token_estimate"`
	ReviewRequired   bool              `json:"review_required"`
	CandidateReviews map[string]string `json:"candidate_reviews,omitempty"`
	ConsumerResults  []QueryResult     `json:"consumer_results,omitempty"`
}
type retrievalReplaySummary struct {
	PositiveCases      int     `json:"positive_cases"`
	Expected           int     `json:"expected_sources"`
	Hits               int     `json:"hit_sources"`
	Recall             float64 `json:"recall_at_5"`
	MRR                float64 `json:"mrr_at_5"`
	NegativeCases      int     `json:"negative_cases"`
	FalseSuggestions   int     `json:"negative_false_suggestions"`
	ReviewCases        int     `json:"review_cases"`
	UnjudgedCandidates int     `json:"unjudged_candidates"`
	KnownConfuserHits  int     `json:"known_confuser_hits"`
}

// Score sources rather than duplicate summary/route atoms.
func retrievalReplayScore(expected, actual []string) (int, float64) {
	want, seen := map[string]bool{}, map[string]bool{}
	for _, p := range expected {
		want[p] = true
	}
	hits, rank, rr := 0, 0, 0.0
	for _, p := range actual {
		if seen[p] {
			continue
		}
		seen[p] = true
		rank++
		if want[p] {
			hits++
			if rr == 0 {
				rr = 1 / float64(rank)
			}
		}
	}
	return hits, rr
}
func retrievalReplayP95(v []float64) float64 {
	copyV := append([]float64(nil), v...)
	sort.Float64s(copyV)
	return copyV[int(math.Ceil(.95*float64(len(copyV))))-1]
}
func TestRetrievalReplaySourceScoring(t *testing.T) {
	h, rr := retrievalReplayScore([]string{"b", "c"}, []string{"a", "a", "b", "b"})
	if h != 1 || rr != .5 {
		t.Fatalf("duplicate source scoring: %d %v", h, rr)
	}
	h, rr = retrievalReplayScore(nil, []string{"a"})
	if h != 0 || rr != 0 {
		t.Fatal("negative scoring")
	}
}
func TestRetrievalReplayP95(t *testing.T) {
	v := make([]float64, 20)
	for i := range v {
		v[i] = float64(20 - i)
	}
	if retrievalReplayP95(v) != 19 || v[0] != 20 {
		t.Fatal("nearest-rank P95 or input mutation")
	}
}

// Opt-in real-corpus replay using the production generator and committed CLI.
// Normal CI skips this integration experiment, keeping corpus separate from
// deterministic unit tests. The experiment measures retrieval, not adoption.
func TestRetrievalBaselineReplay(t *testing.T) {
	datasetPath, reportPath := os.Getenv("AI_SKILL_RETRIEVAL_CASES"), os.Getenv("AI_SKILL_RETRIEVAL_REPORT")
	if datasetPath == "" {
		t.Skip("explicit AI_SKILL_RETRIEVAL_CASES required")
	}
	if reportPath == "" {
		t.Fatal("AI_SKILL_RETRIEVAL_REPORT required")
	}
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatal(err)
	}
	generatorRepo := strings.TrimSpace(string(out))
	repo := generatorRepo
	if corpus := os.Getenv("AI_SKILL_RETRIEVAL_CORPUS"); corpus != "" {
		repo, err = filepath.Abs(corpus)
		if err != nil {
			t.Fatal(err)
		}
	}
	binary, err := retrievalPinnedBinary(os.Getenv("AI_SKILL_RETRIEVAL_QUERY_BINARY"), os.Getenv("AI_SKILL_RETRIEVAL_QUERY_SHA256"))
	if err != nil {
		t.Fatal(err)
	}
	querySHA, err := retrievalFileSHA(binary)
	if err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(datasetPath)
	if err != nil {
		t.Fatal(err)
	}
	var set retrievalReplaySet
	if err = json.Unmarshal(data, &set); err != nil {
		t.Fatal(err)
	}
	validSuite := set.Schema == 1 && len(set.Cases) >= 18 || set.Schema == 2 && set.SuiteKind == "overlap-negative-v1" && len(set.Cases) >= 4
	if !validSuite || set.Limit != 5 || set.Repeats < 20 {
		t.Fatal("invalid replay contract")
	}
	groups, ids := map[string]string{}, map[string]bool{}
	for _, c := range set.Cases {
		if ids[c.ID] || c.Group == "" || (c.Split != "diagnostic" && c.Split != "holdout") {
			t.Fatal("invalid id/group/split", c.ID)
		}
		ids[c.ID] = true
		if set.SuiteKind == "overlap-negative-v1" && (c.ProvenanceKind != "synthetic-overlap-control" || !c.ReviewRequired || len(c.Excluded) == 0) {
			t.Fatal("overlap suite requires review and prelabelled confusers", c.ID)
		}
		gold := map[string]bool{}
		for _, source := range c.Expected {
			if gold[source] {
				t.Fatal("duplicate expected source", c.ID)
			}
			gold[source] = true
		}
		for _, source := range c.Excluded {
			if gold[source] {
				t.Fatal("contradictory relevance labels", c.ID)
			}
		}
		if old := groups[c.Group]; old != "" && old != c.Split {
			t.Fatal("group leak", c.Group)
		}
		groups[c.Group] = c.Split
		if c.ProvenanceKind == "historical-lesson-replay" {
			b, e := os.ReadFile(filepath.Join(repo, c.Provenance))
			if e != nil {
				t.Fatal(e)
			}
			if fmt.Sprintf("%x", sha256.Sum256(b)) != c.ProvenanceSHA {
				t.Fatal("provenance changed", c.ID)
			}
			if len(c.Expected) == 0 {
				t.Fatal("missing expected source", c.ID)
			}
		} else if c.ProvenanceKind == "synthetic-overlap-control" && set.SuiteKind == "overlap-negative-v1" && c.ReviewRequired && len(c.Excluded) > 0 {
			// Overlap controls require candidate-level review; presence alone is not an automatic false-positive verdict.
		} else if c.ProvenanceKind != "synthetic-negative-control" || c.ReviewRequired || len(c.Expected) != 0 {
			t.Fatal("unknown provenance", c.ID)
		}
		for _, p := range append(append([]string{}, c.Expected...), c.Excluded...) {
			if _, e := os.Stat(filepath.Join(repo, p)); e != nil {
				t.Fatal(e)
			}
		}
	}
	if err := retrievalRequireCleanCorpus(repo); err != nil {
		t.Fatal(err)
	}
	revision, err := retrievalGitRevision(repo)
	if err != nil {
		t.Fatal(err)
	}
	index := os.Getenv("AI_SKILL_RETRIEVAL_INDEX")
	group := os.Getenv("AI_SKILL_RETRIEVAL_GROUP")
	var identity retrievalIndexIdentity
	if index != "" {
		identity, err = retrievalReadIndex(index, os.Getenv("AI_SKILL_RETRIEVAL_INDEX_MANIFEST"), revision)
		if err == nil {
			err = retrievalValidateGroup(group, querySHA, identity, os.Getenv("AI_SKILL_RETRIEVAL_REFERENCE"))
		}
	} else {
		if group != "" {
			t.Fatal("an experiment group requires a prebuilt index, manifest and reference; implicit rebuild forbidden")
		}
		index = filepath.Join(t.TempDir(), "replay.sqlite")
		identity, err = retrievalBuildIndex(generatorRepo, repo, index)
	}
	if err != nil {
		t.Fatal(err)
	}
	db, err := sql.Open("sqlite", index)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := nativeRuntimeIndexChecksumsCheck(repo, db); err != nil {
		t.Fatal(err)
	}
	versionContext, cancelVersion := context.WithTimeout(context.Background(), 10*time.Second)
	versionPayload, err := exec.CommandContext(versionContext, binary, "version", "--json").Output()
	cancelVersion()
	if err != nil || !json.Valid(versionPayload) {
		t.Fatalf("query binary version identity unavailable: %v", err)
	}
	binaryPath, err := filepath.Rel(generatorRepo, binary)
	if err != nil {
		// A Windows artifact may be on another volume; display paths must
		// not prevent execution of a correctly pinned binary.
		binaryPath = binary
	}
	summaries := map[string]*retrievalReplaySummary{"diagnostic": {}, "holdout": {}}
	observations := []retrievalReplayObservation{}
	allWarm := []float64{}
	maxPreview := 0
	for _, c := range set.Cases {
		args := []string{"runtime", "query", "--repo", repo, "--db", index, "--keyword", c.Query, "--limit", "5", "--json"}
		if c.Layer != "" {
			args = append(args, "--layer", c.Layer)
		}
		if c.Status != "" {
			args = append(args, "--status", c.Status)
		}
		var first Result
		warm := []float64{}
		firstMS := 0.0
		for n := 0; n <= set.Repeats; n++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			cmd := exec.CommandContext(ctx, binary, args...)
			start := time.Now()
			payload, e := cmd.Output()
			elapsed := float64(time.Since(start).Microseconds()) / 1000
			cancel()
			if e != nil {
				t.Fatalf("query %s: %v", c.ID, e)
			}
			var result Result
			if e = json.Unmarshal(payload, &result); e != nil {
				t.Fatal(e)
			}
			if result.ExitCode != 0 || result.Command != "runtime query" || result.Status != "success" {
				t.Fatalf("query status %s: %#v", c.ID, result.Error)
			}
			if n == 0 {
				first = result
				firstMS = elapsed
			} else {
				if !reflect.DeepEqual(first.Results, result.Results) {
					t.Fatal("non-deterministic", c.ID)
				}
				warm = append(warm, elapsed)
				allWarm = append(allWarm, elapsed)
			}
		}
		actual := []string{}
		for _, r := range first.Results {
			actual = append(actual, r.SourcePath)
		}
		previewBytes, _ := json.Marshal(first.Results)
		estimate := utf8.RuneCount(previewBytes)
		if estimate > maxPreview {
			maxPreview = estimate
		}
		hits, rr := retrievalReplayScore(c.Expected, actual)
		miss := map[string]string{}
		for _, p := range c.Expected {
			found := false
			for _, a := range actual {
				if a == p {
					found = true
				}
			}
			if found {
				continue
			}
			var count int
			if e := db.QueryRow("SELECT COUNT(*) FROM sources WHERE source_path=?", p).Scan(&count); e != nil {
				t.Fatal(e)
			}
			reason := "indexed_source_not_returned_cause_not_attributed"
			if count == 0 {
				reason = "source_not_indexed"
			}
			miss[p] = reason
		}
		s := summaries[c.Split]
		reviews := map[string]string{}
		needsReview := c.ReviewRequired || os.Getenv("AI_SKILL_RETRIEVAL_REVIEW_ALL") == "1" && c.Split == "diagnostic" && len(c.Expected) > 0
		if needsReview {
			s.ReviewCases++
			reviews = retrievalCandidateReviews(c.Expected, c.Excluded, actual)
			for _, verdict := range reviews {
				if verdict == "unreviewed" {
					s.UnjudgedCandidates++
				}
				if verdict == "prelabelled-confuser" {
					s.KnownConfuserHits++
				}
			}
		}
		if len(c.Expected) == 0 && !c.ReviewRequired {
			s.NegativeCases++
			if len(actual) > 0 {
				s.FalseSuggestions++
			}
		} else if len(c.Expected) > 0 {
			s.PositiveCases++
			s.Expected += len(c.Expected)
			s.Hits += hits
			s.MRR += rr
		}
		var consumerResults []QueryResult
		if os.Getenv("AI_SKILL_RETRIEVAL_CAPTURE_RESULTS") == "1" {
			consumerResults = first.Results
		}
		observations = append(observations, retrievalReplayObservation{c.ID, c.Split, actual, len(c.Expected), hits, rr, miss, firstMS, retrievalReplayP95(warm), estimate, needsReview, reviews, consumerResults})
	}
	for _, s := range summaries {
		if s.Expected > 0 {
			s.Recall = float64(s.Hits) / float64(s.Expected)
		}
		if s.PositiveCases > 0 {
			s.MRR /= float64(s.PositiveCases)
		}
	}
	if err := nativeRuntimeIndexChecksumsCheck(repo, db); err != nil {
		t.Fatal(err)
	}
	if err := retrievalRequireCleanCorpus(repo); err != nil {
		t.Fatal(err)
	}
	if sha, e := retrievalFileSHA(binary); e != nil || sha != querySHA {
		t.Fatal("query binary changed during replay", e)
	}
	if sha, e := retrievalFileSHA(index); e != nil || sha != identity.IndexSHA {
		t.Fatal("index artifact changed during replay", e)
	}
	report := map[string]any{
		"schema_version": 2, "corpus_revision": revision, "dataset_sha256": fmt.Sprintf("%x", sha256.Sum256(data)),
		"generator_inputs_sha256": identity.InputsSHA, "binary_sha256": querySHA,
		"index_identity": identity, "query_binary_path": filepath.ToSlash(binaryPath), "query_build_info": json.RawMessage(versionPayload), "experiment_group": group,
		"miss_classification": "source-row presence plus actual binary results only; no host-side FTS/query reimplementation",
		"environment":         runtime.GOOS + "/" + runtime.GOARCH, "go_version": runtime.Version(), "cases": len(set.Cases), "warm_runs_per_case": set.Repeats,
		"first_invocation":     "per-case first CLI invocation; OS cache not controlled, not cold-start proof",
		"preview_token_method": "UTF-8 rune count of result JSON; estimate only, not model tokens or cap enforcement",
		"warm_cli_p95_ms":      retrievalReplayP95(allWarm), "max_preview_token_estimate": maxPreview,
		"summary": summaries, "observations": observations, "adoption_metrics": "not evaluated", "applicability_distribution": "not evaluated: source qualification not wired",
		"route_discovery": "not evaluated: Bridge unchanged; workflow knowledge lookup only",
	}
	encoded, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(reportPath, append(encoded, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Logf("report=%s diagnostic=%+v holdout=%+v", reportPath, *summaries["diagnostic"], *summaries["holdout"])
}

// Controlled probes distinguish current FTS phrase/token behavior from semantic
// retrieval. They are artificial controls, not additional historical cases.
func TestRetrievalReplayFTSQueryControls(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)
	if _, err = db.Exec("CREATE VIRTUAL TABLE probe USING fts5(content); INSERT INTO probe(content) VALUES ('host pipeline status'), ('模型載入失敗')"); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		query string
		want  int
	}{{"host pipeline", 1}, {"pipeline host", 0}, {"模型", 0}, {"模型載入失敗", 1}} {
		var got int
		if err = db.QueryRow("SELECT count(*) FROM probe WHERE probe MATCH ?", runtimeFTSMatchLiteral(tc.query)).Scan(&got); err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("%q got=%d want=%d", tc.query, got, tc.want)
		}
	}
}

func TestRetrievalReplayStaleAndHistoricalControls(t *testing.T) {
	repo := t.TempDir()
	writeRuntimeNativeReportSourceFixture(t, repo)
	writeFile(t, filepath.Join(repo, "README.md"), "# baseline\n")
	writeFile(t, filepath.Join(repo, "CORE_BOOTSTRAP.md"), "# bootstrap\n")
	writeFile(t, filepath.Join(repo, "feedback", "history", "README.md"), "# feedback\n")
	index := filepath.Join(t.TempDir(), "controls.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, index); err != nil {
		t.Fatal(err)
	}
	db, err := sql.Open("sqlite", index)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	// Synthetic retirement represents a historical result that remains relevant
	// to an audit query, but must not silently become a current-adoption result.
	if _, err = db.Exec("UPDATE atoms SET status='deprecated' WHERE id='test.summary'"); err != nil {
		t.Fatal(err)
	}
	historical, err := runRuntimeIndexQuery(index, "Test summary text", runtimeOptions{limit: 5})
	if err != nil {
		t.Fatal(err)
	}
	if len(historical) == 0 || historical[0].ID != "test.summary" || historical[0].Status != "deprecated" {
		t.Fatal("history hidden", historical)
	}
	current, err := runRuntimeIndexQuery(index, "Test summary text", runtimeOptions{limit: 5, statusFilter: "validated"})
	if err != nil {
		t.Fatal(err)
	}
	if len(current) != 0 {
		t.Fatal("retired result leaked through explicit current-status filter")
	}
	if err = nativeRuntimeIndexChecksumsCheck(repo, db); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(repo, "README.md"), "# changed after lookup\n")
	if err = nativeRuntimeIndexChecksumsCheck(repo, db); err == nil {
		t.Fatal("stale source accepted by replay freshness check")
	}
}
