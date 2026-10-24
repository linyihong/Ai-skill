package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestRetrievalEngineeringExportIndex(t *testing.T) {
	p := os.Getenv("AI_SKILL_ENGINEERING_INDEX")
	if p == "" {
		t.Skip("explicit diagnostic export")
	}
	root, err := filepath.Abs("../../../..")
	if err != nil {
		t.Fatal(err)
	}
	if err = buildNativeRuntimeSQLiteIndex(root, p); err != nil {
		t.Fatal(err)
	}
}
func TestRuntimeIntelligenceBodyCoverage(t *testing.T) {
	repo := t.TempDir()
	writeFeedbackHistoryRouteFixture(t, repo)
	dir := filepath.Join(repo, "intelligence/engineering/execution/validation-reasoning")
	os.MkdirAll(dir, 0755)
	os.WriteFile(filepath.Join(dir, "probe.md"), []byte("# Probe\n\n**Status**: `candidate-intelligence`\n\nBody contains independent_observation.\n"), 0644)
	db := filepath.Join(t.TempDir(), "index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
		t.Fatal(err)
	}
	got, err := runRuntimeIndexQuery(db, "independent_observation", runtimeOptions{limit: 5})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].SourcePath != "intelligence/engineering/execution/validation-reasoning/probe.md" {
		t.Fatalf("body source absent: %+v", got)
	}
}

func TestRuntimeSourceLifecycleAndMetadata(t *testing.T) {
	repo := t.TempDir()
	writeFeedbackHistoryRouteFixture(t, repo)
	dir := filepath.Join(repo, filepath.FromSlash(runtimeIntelligenceDir))
	os.MkdirAll(dir, 0755)
	path := filepath.Join(dir, "first.md")
	writeFile(t, path, "---\nstatus: candidate\ndomains: [engineering]\nscope: local\nreplaces: old.md\n---\n# Probe\n\nneedle one\n")
	db := filepath.Join(t.TempDir(), "index.sqlite")
	build := func() {
		t.Helper()
		if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
			t.Fatal(err)
		}
	}
	query := func() Result {
		return buildRuntimeQueryResult(runtimeOptions{repoPath: repo, dbPath: db, keyword: "needle", sourceLevel: true, limit: 5})
	}
	build()
	r := query()
	if r.ExitCode != 0 || len(r.Results) != 1 {
		t.Fatalf("initial: %+v", r)
	}
	d := r.Results[0].Source
	if d == nil || d.CanonicalStatus != "candidate" || d.Scope != "local" || d.Supersession != "old.md" || len(d.Domains) != 1 || d.Checksum == "" {
		t.Fatalf("metadata %+v", d)
	}
	for _, action := range []string{"edit", "add", "move", "delete"} {
		switch action {
		case "edit":
			writeFile(t, path, "# Changed\n\nneedle two\n")
		case "add":
			writeFile(t, filepath.Join(dir, "second.md"), "# Second\n\nneedle three\n")
		case "move":
			if err := os.Rename(path, filepath.Join(dir, "moved.md")); err != nil {
				t.Fatal(err)
			}
			path = filepath.Join(dir, "moved.md")
		case "delete":
			os.Remove(path)
		}
		if r = query(); r.Error == nil || r.Error.Code != "stale_runtime_index" || len(r.Results) > 0 {
			t.Fatalf("%s failed to block stale: %+v", action, r)
		}
		build()
		if r = query(); r.ExitCode != 0 {
			t.Fatalf("%s refresh %+v", action, r)
		}
	}
	if len(r.Results) != 1 || r.Results[0].Source.CanonicalStatus != "unknown" || r.Results[0].Source.Scope != "unknown" {
		t.Fatalf("unknown metadata %+v", r)
	}
	first, _ := json.Marshal(r.Results)
	build()
	r = query()
	second, _ := json.Marshal(r.Results)
	if string(first) != string(second) {
		t.Fatal("non-deterministic source results")
	}
	before, _ := os.ReadFile(db)
	writeFile(t, filepath.Join(dir, "broken.md"), "---\nstatus: [\n---\n# Broken\n")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err == nil {
		t.Fatal("bad metadata accepted")
	}
	after, _ := os.ReadFile(db)
	if !bytes.Equal(before, after) {
		t.Fatal("failed build destroyed existing cache")
	}
	missing := buildRuntimeQueryResult(runtimeOptions{repoPath: repo, dbPath: filepath.Join(t.TempDir(), "missing.sqlite"), keyword: "needle", sourceLevel: true, limit: 5})
	if missing.Error == nil || missing.Error.Code != "missing_runtime_index" {
		t.Fatalf("missing: %+v", missing)
	}
}

func TestRuntimeSourceDedupBeforeLimit(t *testing.T) {
	repo := t.TempDir()
	writeFeedbackHistoryRouteFixture(t, repo)
	writeFile(t, filepath.Join(repo, "knowledge/runtime/routing-registry.yaml"), `records:
  - id: route.a
    primary_source: intelligence/engineering/execution/validation-reasoning/a.md
    task_intent: needle
  - id: route.b
    primary_source: intelligence/engineering/execution/validation-reasoning/a.md
    task_intent: needle
`)
	for _, n := range []string{"a", "b"} {
		writeFile(t, filepath.Join(repo, filepath.FromSlash(runtimeIntelligenceDir), n+".md"), "# needle\n\nneedle\n")
	}
	db := filepath.Join(t.TempDir(), "index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
		t.Fatal(err)
	}
	r := buildRuntimeQueryResult(runtimeOptions{repoPath: repo, dbPath: db, keyword: "needle", sourceLevel: true, limit: 2})
	if r.ExitCode != 0 || len(r.Results) != 2 || r.Results[0].SourcePath == r.Results[1].SourcePath {
		t.Fatalf("dedup %+v", r)
	}
}

func TestRuntimeGraphBoundedCycle(t *testing.T) {
	repo := t.TempDir()
	writeFile(t, filepath.Join(repo, "knowledge/graphs/a.yaml"), `id: cycle
source: a.md
edges:
 - type: references
   target: b.md
 - type: depends_on
   target: a.md
`)
	opts := runtimeOptions{repoPath: repo, graphQuery: true, source: "a.md", limit: 1}
	a := buildRuntimeQueryResult(opts)
	b := buildRuntimeQueryResult(opts)
	x, _ := json.Marshal(a.Results)
	y, _ := json.Marshal(b.Results)
	if len(a.Results) != 1 || a.Results[0].Depth != 1 || a.Results[0].EdgeType != "references" || !bytes.Equal(x, y) {
		t.Fatalf("bounded graph %+v", a)
	}
	opts.queryType = "depends_on"
	a = buildRuntimeQueryResult(opts)
	if len(a.Results) != 1 || a.Results[0].Target != "a.md" || a.Results[0].Depth != 1 {
		t.Fatalf("cycle expanded %+v", a)
	}
}

func TestRuntimeSourceCanonicalStatusNotRouteCompatibility(t *testing.T) {
	repo := t.TempDir()
	writeFile(t, filepath.Join(repo, "knowledge/runtime/routing-registry.yaml"), `records:
  - id: route.a
    primary_source: intelligence/engineering/execution/validation-reasoning/a.md
    task_intent: needle
    metadata:
      compatibility_state: active
`)
	writeFile(t, filepath.Join(repo, filepath.FromSlash(runtimeIntelligenceDir), "a.md"), "---\nstatus: deprecated\ndomains: []\nreplaces: [old.md]\n---\n# needle\n\nneedle\n")
	db := filepath.Join(t.TempDir(), "index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
		t.Fatal(err)
	}
	opts := runtimeOptions{repoPath: repo, dbPath: db, keyword: "needle", sourceLevel: true, limit: 5, statusFilter: "deprecated", queryType: "route"}
	r := buildRuntimeQueryResult(opts)
	if r.ExitCode != 0 || len(r.Results) != 1 || r.Results[0].Status != "active" || r.Results[0].Source.CanonicalStatus != "deprecated" {
		t.Fatalf("status conflated %+v", r)
	}
	opts.statusFilter = "active"
	r = buildRuntimeQueryResult(opts)
	if r.ExitCode != 0 || len(r.Results) != 0 {
		t.Fatalf("compatibility treated as canonical %+v", r)
	}
}

// Explicit binary/index inputs keep this diagnostic independent of whichever
// repo-local binary happens to be installed. No selection labels go to a runner.
func TestRetrievalEngineeringSourceComparison(t *testing.T) {
	binary, index, out := os.Getenv("AI_SKILL_ENGINEERING_QUERY_BINARY"), os.Getenv("AI_SKILL_ENGINEERING_QUERY_INDEX"), os.Getenv("AI_SKILL_ENGINEERING_OUTPUT")
	if binary == "" && index == "" && out == "" {
		t.Skip("explicit pinned comparison required")
	}
	if binary == "" || index == "" || out == "" {
		t.Fatal("binary, index and output all required")
	}
	root := repoRootForTest(t)
	e := filepath.Join(root, "plans/archived/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence")
	raw, err := os.ReadFile(filepath.Join(e, "phase2-intelligence-cases.json"))
	if err != nil {
		t.Fatal(err)
	}
	type probe struct {
		Query  string `json:"query"`
		Source string `json:"expected_source"`
		SHA    string `json:"source_sha256"`
	}
	var dataset struct {
		Cases []probe `json:"cases"`
	}
	if err = json.Unmarshal(raw, &dataset); err != nil {
		t.Fatal(err)
	}
	binaryRaw, err := os.ReadFile(binary)
	if err != nil {
		t.Fatal(err)
	}
	binarySHA := fmt.Sprintf("%x", sha256.Sum256(binaryRaw))
	indexRaw, err := os.ReadFile(index)
	if err != nil {
		t.Fatal(err)
	}
	sourceLevel := os.Getenv("AI_SKILL_ENGINEERING_SOURCE_LEVEL") == "1"
	type scored struct {
		Query   string   `json:"query"`
		Source  string   `json:"expected_source"`
		Sources []string `json:"sources_at_5"`
		Rank    int      `json:"expected_rank"`
		Raw     Result   `json:"raw_result"`
	}
	scores := []scored{}
	times := []float64{}
	hits := 0
	mrr := 0.0
	for _, c := range dataset.Cases {
		source, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(c.Source)))
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprintf("%x", sha256.Sum256(source)) != c.SHA {
			t.Fatal("frozen source changed", c.Source)
		}
		args := []string{"runtime", "query", "--repo", root, "--db", index, "--keyword", c.Query, "--limit", "64", "--json"}
		if sourceLevel {
			args = append(args, "--source-level")
		}
		var result Result
		for rep := 0; rep < 4; rep++ {
			started := time.Now()
			output, err := exec.Command(binary, args...).CombinedOutput()
			elapsed := float64(time.Since(started).Microseconds()) / 1000
			if err != nil {
				t.Fatalf("query: %v %s", err, output)
			}
			if err = json.Unmarshal(output, &result); err != nil {
				t.Fatal(err)
			}
			if rep > 0 {
				times = append(times, elapsed)
			}
		}
		normalize := func(s string) string {
			return strings.ReplaceAll(strings.ReplaceAll(s, index, "<INDEX_PATH>"), root, "<AI_SKILL_REPO>")
		}
		for i := range result.Checks {
			result.Checks[i].Message = normalize(result.Checks[i].Message)
		}
		for i := range result.PlannedActions {
			result.PlannedActions[i] = normalize(result.PlannedActions[i])
		}
		s := scored{Query: c.Query, Source: c.Source, Sources: []string{}, Raw: result}
		seen := map[string]bool{}
		for _, r := range result.Results {
			if seen[r.SourcePath] {
				continue
			}
			seen[r.SourcePath] = true
			s.Sources = append(s.Sources, r.SourcePath)
			if len(s.Sources) == 5 {
				break
			}
		}
		for i, p := range s.Sources {
			if p == c.Source {
				s.Rank = i + 1
				hits++
				mrr += 1 / float64(i+1)
			}
		}
		scores = append(scores, s)
	}
	sort.Float64s(times)
	p95 := times[(len(times)*95+99)/100-1]
	generatorFiles := map[string]string{}
	paths, _ := filepath.Glob(filepath.Join(root, "scripts/ai-skill-cli/internal/app/*.go"))
	for _, p := range paths {
		if strings.HasSuffix(p, "_test.go") {
			continue
		}
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		generatorFiles[filepath.Base(p)] = fmt.Sprintf("%x", sha256.Sum256(b))
	}
	result := map[string]any{"source_level": sourceLevel, "dataset_sha256": fmt.Sprintf("%x", sha256.Sum256(raw)), "query_binary_sha256": binarySHA, "index_sha256": fmt.Sprintf("%x", sha256.Sum256(indexRaw)), "generator_source_files": generatorFiles, "hits": hits, "denominator": len(scores), "mrr": mrr / float64(len(scores)), "warm_cli_samples_ms": times, "warm_cli_p95_ms": p95, "cases": scores}
	if err := retrievalEngineeringWriteNew(out, result); err != nil {
		t.Fatal(err)
	}
	after, err := os.ReadFile(binary)
	if err != nil || fmt.Sprintf("%x", sha256.Sum256(after)) != binarySHA {
		t.Fatal("query executable changed during comparison")
	}
}
func retrievalEngineeringWriteNew(path string, v any) error {
	raw, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(append(raw, '\n'))
	return err
}

func TestRuntimeSourceStatusRetainsUnknown(t *testing.T) {
	repo := t.TempDir()
	writeFeedbackHistoryRouteFixture(t, repo)
	writeFile(t, filepath.Join(repo, filepath.FromSlash(runtimeIntelligenceDir), "unknown.md"), "# needle\n\nneedle\n")
	writeFile(t, filepath.Join(repo, filepath.FromSlash(runtimeIntelligenceDir), "known.md"), "---\nstatus: deprecated\n---\n# needle\n\nneedle\n")
	db := filepath.Join(t.TempDir(), "index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
		t.Fatal(err)
	}
	r := buildRuntimeQueryResult(runtimeOptions{repoPath: repo, dbPath: db, keyword: "needle", sourceLevel: true, limit: 5, statusFilter: "candidate"})
	if r.ExitCode != 0 || len(r.Results) != 1 || r.Results[0].Source.CanonicalStatus != "unknown" {
		t.Fatalf("unknown silently filtered: %+v", r)
	}
}

func TestRuntimeFailedIndexPopulationPreservesCache(t *testing.T) {
	repo := t.TempDir()
	writeFeedbackHistoryRouteFixture(t, repo)
	db := filepath.Join(t.TempDir(), "index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
		t.Fatal(err)
	}
	before, err := os.ReadFile(db)
	if err != nil {
		t.Fatal(err)
	}
	// Identical duplicate IDs pass source reads but fail the second SQL INSERT,
	// after temporary DB population has begun. This exercises atomic replacement,
	// unlike a parse error that is discovered before any DB is opened.
	writeFile(t, filepath.Join(repo, "knowledge/runtime/routing-registry.yaml"), `records:
  - id: route.duplicate
    primary_source: docs/source.md
    task_intent: duplicate
  - id: route.duplicate
    primary_source: docs/source.md
    task_intent: duplicate
`)
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err == nil {
		t.Fatal("duplicate SQL primary key accepted")
	}
	after, err := os.ReadFile(db)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(before, after) {
		t.Fatal("failed SQL population replaced existing cache")
	}
	leftovers, err := filepath.Glob(filepath.Join(filepath.Dir(db), ".runtime-index-*.sqlite"))
	if err != nil || len(leftovers) != 0 {
		t.Fatalf("partial cache leaked: %v %v", leftovers, err)
	}
}
