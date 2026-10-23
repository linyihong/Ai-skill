package app

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// The index artifact binds the generator identity separately from the query
// executable. Reading it never rebuilds the index with the current test code.
type retrievalIndexIdentity struct {
	Schema             int    `json:"schema_version"`
	CorpusRevision     string `json:"corpus_revision"`
	GeneratorRevision  string `json:"generator_revision"`
	GeneratorSourceSHA string `json:"generator_module_source_sha256"`
	CorpusSourcesSHA   string `json:"corpus_sources_sha256"`
	InputsSHA          string `json:"generator_inputs_sha256"`
	IndexSHA           string `json:"index_sha256"`
}

func retrievalFileSHA(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(b)), nil
}
func retrievalGitRevision(repo string) (string, error) {
	b, err := exec.Command("git", "-C", repo, "rev-parse", "HEAD").Output()
	return strings.TrimSpace(string(b)), err
}
func retrievalGeneratorSourceSHA(repo string) (string, error) {
	root := filepath.Join(repo, "scripts", "ai-skill-cli")
	files := map[string]string{}
	err := filepath.WalkDir(root, func(path string, e os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if e.IsDir() {
			if e.Name() == "bin" || e.Name() == "dist" {
				return filepath.SkipDir
			}
			return nil
		}
		if (strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go")) || e.Name() == "go.mod" || e.Name() == "go.sum" {
			sha, err := retrievalFileSHA(path)
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			files[filepath.ToSlash(rel)] = sha
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(files)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha256.Sum256(b)), nil
}
func retrievalBuildIndex(generatorRepo, corpus, index string) (retrievalIndexIdentity, error) {
	var id retrievalIndexIdentity
	id.Schema = 1
	if _, err := os.Stat(index); err == nil {
		return id, fmt.Errorf("refuse to overwrite index artifact %s", index)
	} else if !os.IsNotExist(err) {
		return id, err
	}
	var err error
	if err := retrievalRequireCommittedGenerator(generatorRepo); err != nil {
		return id, err
	}
	if err := retrievalRequireCleanCorpus(corpus); err != nil {
		return id, err
	}
	if id.CorpusRevision, err = retrievalGitRevision(corpus); err != nil {
		return id, err
	}
	if id.GeneratorRevision, err = retrievalGitRevision(generatorRepo); err != nil {
		return id, err
	}
	if id.GeneratorSourceSHA, err = retrievalGeneratorSourceSHA(generatorRepo); err != nil {
		return id, err
	}
	records, err := runtimeIndexRecords(corpus)
	if err != nil {
		return id, err
	}
	edges, err := runtimeIndexEdges(corpus)
	if err != nil {
		return id, err
	}
	sources, err := runtimeIndexSources(corpus, records)
	if err != nil {
		return id, err
	}
	sourceMap := map[string]string{}
	for _, source := range sources {
		sourceMap[source.sourcePath] = source.checksum
	}
	sourceBytes, err := json.Marshal(sourceMap)
	if err != nil {
		return id, err
	}
	id.CorpusSourcesSHA = fmt.Sprintf("%x", sha256.Sum256(sourceBytes))
	b, err := json.Marshal([]any{records, edges, sources})
	if err != nil {
		return id, err
	}
	id.InputsSHA = fmt.Sprintf("%x", sha256.Sum256(b))
	if err = buildNativeRuntimeSQLiteIndex(corpus, index); err != nil {
		return id, err
	}
	if err := retrievalRequireCleanCorpus(corpus); err != nil {
		return id, err
	}
	if err := retrievalRequireCommittedGenerator(generatorRepo); err != nil {
		return id, err
	}
	if sha, e := retrievalGeneratorSourceSHA(generatorRepo); e != nil || sha != id.GeneratorSourceSHA {
		return id, fmt.Errorf("generator source changed during snapshot: %v", e)
	}
	if id.IndexSHA, err = retrievalFileSHA(index); err != nil {
		return id, err
	}
	return id, nil
}
func retrievalReadIndex(index, manifest, corpusRevision string) (retrievalIndexIdentity, error) {
	var id retrievalIndexIdentity
	b, err := os.ReadFile(manifest)
	if err != nil {
		return id, err
	}
	if err = json.Unmarshal(b, &id); err != nil {
		return id, err
	}
	if id.Schema != 1 || id.GeneratorRevision == "" || len(id.GeneratorSourceSHA) != 64 || len(id.InputsSHA) != 64 || len(id.CorpusSourcesSHA) != 64 || id.CorpusRevision != corpusRevision {
		return id, fmt.Errorf("index identity / corpus mismatch")
	}
	sha, err := retrievalFileSHA(index)
	if err != nil {
		return id, err
	}
	if sha != id.IndexSHA {
		return id, fmt.Errorf("index SHA mismatch")
	}
	return id, nil
}
func retrievalPinnedBinary(path, expectedSHA string) (string, error) {
	if path == "" || len(expectedSHA) != 64 {
		return "", fmt.Errorf("explicit query binary path and expected SHA-256 required")
	}
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	sha, err := retrievalFileSHA(path)
	if err != nil {
		return "", err
	}
	if sha != expectedSHA {
		return "", fmt.Errorf("query binary SHA mismatch")
	}
	return path, nil
}
func retrievalValidateGroup(group, querySHA string, id retrievalIndexIdentity, referencePath string) error {
	var ref struct {
		Corpus string `json:"corpus_revision"`
		Query  string `json:"binary_sha256"`
		Inputs string `json:"generator_inputs_sha256"`
	}
	b, err := os.ReadFile(referencePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &ref); err != nil {
		return err
	}
	if ref.Corpus == "" || len(ref.Query) != 64 || len(ref.Inputs) != 64 || id.CorpusRevision != ref.Corpus {
		return fmt.Errorf("invalid reference or corpus changed")
	}
	oldQuery, oldIndex := querySHA == ref.Query, id.InputsSHA == ref.Inputs
	switch group {
	case "B0":
		if !oldQuery || !oldIndex {
			return fmt.Errorf("B0 requires old index and old query")
		}
	case "Q":
		if oldQuery || !oldIndex {
			return fmt.Errorf("Q requires old index and distinct new query binary")
		}
	case "E":
		if !oldQuery || oldIndex {
			return fmt.Errorf("E requires changed extraction inputs and old query binary")
		}
	case "QE":
		if oldQuery || oldIndex {
			return fmt.Errorf("QE requires changed extraction inputs and distinct new query binary")
		}
	default:
		return fmt.Errorf("explicit experiment group B0/Q/E/QE required")
	}
	return nil
}

// Build a persistent index+manifest before executing any group. The generator
// checkout and corpus checkout are intentionally independent.
func TestRetrievalIndexSnapshot(t *testing.T) {
	index := os.Getenv("AI_SKILL_RETRIEVAL_INDEX_OUT")
	if index == "" {
		t.Skip("explicit snapshot output required")
	}
	b, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatal(err)
	}
	generator := strings.TrimSpace(string(b))
	corpus := os.Getenv("AI_SKILL_RETRIEVAL_CORPUS")
	if corpus == "" {
		t.Fatal("explicit corpus required for index snapshot")
	}
	id, err := retrievalBuildIndex(generator, corpus, index)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.MarshalIndent(id, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(index+".manifest.json", append(data, '\n'), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Logf("index=%s manifest=%s", index, index+".manifest.json")
}
func TestRetrievalReplayRejectsWrongArtifacts(t *testing.T) {
	dir := t.TempDir()
	binary := filepath.Join(dir, "query")
	if err := os.WriteFile(binary, []byte("old query"), 0o644); err != nil {
		t.Fatal(err)
	}
	old, _ := retrievalFileSHA(binary)
	newSHA := strings.Repeat("b", 64)
	oldInputs := strings.Repeat("c", 64)
	newInputs := strings.Repeat("d", 64)
	if _, err := retrievalPinnedBinary("", old); err == nil {
		t.Fatal("implicit binary accepted")
	}
	if _, err := retrievalPinnedBinary(binary, newSHA); err == nil {
		t.Fatal("wrong binary accepted")
	}
	if _, err := retrievalPinnedBinary(binary, old); err != nil {
		t.Fatal(err)
	}
	ref := filepath.Join(dir, "baseline.json")
	data, _ := json.Marshal(map[string]string{"corpus_revision": "corpus", "binary_sha256": old, "generator_inputs_sha256": oldInputs})
	if err := os.WriteFile(ref, data, 0o644); err != nil {
		t.Fatal(err)
	}
	oldID := retrievalIndexIdentity{CorpusRevision: "corpus", InputsSHA: oldInputs}
	newID := oldID
	newID.InputsSHA = newInputs
	for _, tc := range []struct {
		group, query string
		id           retrievalIndexIdentity
		pass         bool
	}{
		{"B0", old, oldID, true}, {"Q", newSHA, oldID, true}, {"E", old, newID, true}, {"QE", newSHA, newID, true},
		{"Q", old, oldID, false}, {"E", newSHA, oldID, false}, {"QE", old, newID, false}, {"B0", old, newID, false},
	} {
		err := retrievalValidateGroup(tc.group, tc.query, tc.id, ref)
		if (err == nil) != tc.pass {
			t.Errorf("%s pass=%v err=%v", tc.group, tc.pass, err)
		}
	}
	index := filepath.Join(dir, "index.sqlite")
	if err := os.WriteFile(index, []byte("index"), 0o644); err != nil {
		t.Fatal(err)
	}
	id := retrievalIndexIdentity{Schema: 1, CorpusRevision: "corpus", GeneratorRevision: "generator", GeneratorSourceSHA: strings.Repeat("e", 64), CorpusSourcesSHA: strings.Repeat("f", 64), InputsSHA: oldInputs}
	id.IndexSHA, _ = retrievalFileSHA(index)
	data, _ = json.Marshal(id)
	manifest := index + ".json"
	if err := os.WriteFile(manifest, data, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := retrievalReadIndex(index, manifest, "corpus"); err != nil {
		t.Fatal(err)
	}
	if _, err := retrievalReadIndex(index, manifest, "other-corpus"); err == nil {
		t.Fatal("different corpus accepted")
	}
	if err := os.WriteFile(index, []byte("tampered index"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := retrievalReadIndex(index, manifest, "corpus"); err == nil {
		t.Fatal("tampered index accepted")
	}
}

func retrievalRequireCleanCorpus(repo string) error {
	b, err := exec.Command("git", "-C", repo, "status", "--porcelain", "--untracked-files=all").Output()
	if err != nil {
		return err
	}
	if strings.TrimSpace(string(b)) != "" {
		return fmt.Errorf("corpus checkout must be clean; use a frozen checkout, not the generator worktree")
	}
	return nil
}

type retrievalAblationReport struct {
	Group   string                 `json:"experiment_group"`
	Dataset string                 `json:"dataset_sha256"`
	Corpus  string                 `json:"corpus_revision"`
	Query   string                 `json:"binary_sha256"`
	Index   retrievalIndexIdentity `json:"index_identity"`
}

func retrievalValidateMatrix(groups map[string]retrievalAblationReport) error {
	if len(groups) != 4 {
		return fmt.Errorf("four experiment groups required")
	}
	b, ok := groups["B0"]
	if !ok {
		return fmt.Errorf("missing B0")
	}
	for _, name := range []string{"B0", "Q", "E", "QE"} {
		r, ok := groups[name]
		if !ok || r.Group != name || r.Dataset == "" || r.Corpus != b.Corpus || r.Index.CorpusRevision != r.Corpus || r.Dataset != b.Dataset {
			return fmt.Errorf("group/corpus/dataset mismatch: %s", name)
		}
		if len(r.Query) != 64 || len(r.Index.IndexSHA) != 64 || len(r.Index.InputsSHA) != 64 || len(r.Index.CorpusSourcesSHA) != 64 || r.Index.GeneratorRevision == "" || len(r.Index.GeneratorSourceSHA) != 64 {
			return fmt.Errorf("missing artifact identity: %s", name)
		}
		if r.Index.CorpusSourcesSHA != b.Index.CorpusSourcesSHA {
			return fmt.Errorf("corpus coverage changed in extraction-only experiment")
		}
	}
	q, e, qe := groups["Q"], groups["E"], groups["QE"]
	if b.Query != e.Query || q.Query != qe.Query || b.Query == q.Query {
		return fmt.Errorf("old/new query executables are not paired")
	}
	if b.Index != q.Index || e.Index != qe.Index || b.Index.InputsSHA == e.Index.InputsSHA || b.Index.GeneratorSourceSHA == e.Index.GeneratorSourceSHA {
		return fmt.Errorf("old/new index artifacts are not paired")
	}
	return nil
}
func TestRetrievalReplayMatrixPairing(t *testing.T) {
	old := retrievalIndexIdentity{Schema: 1, CorpusRevision: "corpus", GeneratorRevision: "old-generator", GeneratorSourceSHA: strings.Repeat("a", 64), CorpusSourcesSHA: strings.Repeat("b", 64), InputsSHA: strings.Repeat("c", 64), IndexSHA: strings.Repeat("d", 64)}
	newer := old
	newer.GeneratorRevision = "new-generator"
	newer.GeneratorSourceSHA = strings.Repeat("e", 64)
	newer.InputsSHA = strings.Repeat("f", 64)
	newer.IndexSHA = strings.Repeat("0", 64)
	groups := map[string]retrievalAblationReport{}
	for _, name := range []string{"B0", "Q", "E", "QE"} {
		r := retrievalAblationReport{name, "dataset", "corpus", strings.Repeat("1", 64), old}
		if name == "Q" || name == "QE" {
			r.Query = strings.Repeat("2", 64)
		}
		if name == "E" || name == "QE" {
			r.Index = newer
		}
		groups[name] = r
	}
	if err := retrievalValidateMatrix(groups); err != nil {
		t.Fatal(err)
	}
	original := groups["QE"]
	bad := original
	bad.Query = strings.Repeat("3", 64)
	groups["QE"] = bad
	if err := retrievalValidateMatrix(groups); err == nil {
		t.Fatal("different Q/QE binaries accepted")
	}
	bad = original
	bad.Index.IndexSHA = strings.Repeat("4", 64)
	groups["QE"] = bad
	if err := retrievalValidateMatrix(groups); err == nil {
		t.Fatal("different E/QE index artifacts accepted")
	}
	bad = original
	bad.Index.CorpusSourcesSHA = strings.Repeat("5", 64)
	groups["QE"] = bad
	if err := retrievalValidateMatrix(groups); err == nil {
		t.Fatal("expanded corpus hidden inside extraction-only accepted")
	}
	groups["QE"] = original
	bad = groups["Q"]
	bad.Dataset = "different-dataset"
	groups["Q"] = bad
	if err := retrievalValidateMatrix(groups); err == nil {
		t.Fatal("different datasets accepted")
	}
}

// After running the four groups, validate their report identities as one unit.
// The input is a JSON array of four absolute report paths, not shell commands.
func TestRetrievalAblationMatrix(t *testing.T) {
	path := os.Getenv("AI_SKILL_RETRIEVAL_MATRIX_REPORTS")
	if path == "" {
		t.Skip("explicit report list required")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var paths []string
	if err = json.Unmarshal(b, &paths); err != nil {
		t.Fatal(err)
	}
	groups := map[string]retrievalAblationReport{}
	for _, p := range paths {
		b, err = os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		var r retrievalAblationReport
		if err = json.Unmarshal(b, &r); err != nil {
			t.Fatal(err)
		}
		if _, ok := groups[r.Group]; ok {
			t.Fatal("duplicate experiment group")
		}
		groups[r.Group] = r
	}
	if err = retrievalValidateMatrix(groups); err != nil {
		t.Fatal(err)
	}
}

func retrievalRequireCommittedGenerator(repo string) error {
	tracked, err := exec.Command("git", "-C", repo, "diff", "HEAD", "--name-only", "-z", "--", "scripts/ai-skill-cli").Output()
	if err != nil {
		return err
	}
	untracked, err := exec.Command("git", "-C", repo, "ls-files", "--others", "--exclude-standard", "-z", "--", "scripts/ai-skill-cli").Output()
	if err != nil {
		return err
	}
	for _, path := range strings.Split(string(append(tracked, untracked...)), "\x00") {
		if strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") || strings.HasSuffix(path, "/go.mod") || strings.HasSuffix(path, "/go.sum") {
			return fmt.Errorf("generator production source must be committed before snapshot: %s", path)
		}
	}
	return nil
}
func TestRetrievalReplayRejectsDirtySourceIdentity(t *testing.T) {
	repo := initTempGitRepo(t)
	if err := retrievalRequireCleanCorpus(repo); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(repo, "scripts", "ai-skill-cli", "probe_test.go"), "package probe\n")
	if err := retrievalRequireCommittedGenerator(repo); err != nil {
		t.Fatal("test-only harness should not change generator identity", err)
	}
	if err := retrievalRequireCleanCorpus(repo); err == nil {
		t.Fatal("dirty corpus accepted")
	}
	writeFile(t, filepath.Join(repo, "scripts", "ai-skill-cli", "probe.go"), "package probe\n")
	if err := retrievalRequireCommittedGenerator(repo); err == nil {
		t.Fatal("uncommitted generator accepted")
	}
	runGit(t, repo, "add", "scripts/ai-skill-cli/probe.go")
	if err := retrievalRequireCommittedGenerator(repo); err == nil {
		t.Fatal("staged but uncommitted generator accepted")
	}
}

func retrievalCandidateReviews(expected, excluded, actual []string) map[string]string {
	reviews := map[string]string{}
	for _, source := range actual {
		reviews[source] = "unreviewed"
		for _, gold := range expected {
			if source == gold {
				reviews[source] = "prelabelled-relevant"
			}
		}
		for _, confuser := range excluded {
			if source == confuser {
				reviews[source] = "prelabelled-confuser"
			}
		}
	}
	return reviews
}
func TestRetrievalReplayCandidateReviewDoesNotInferRelevance(t *testing.T) {
	reviews := retrievalCandidateReviews([]string{"gold"}, []string{"confuser"}, []string{"gold", "confuser", "confuser", "new"})
	if len(reviews) != 3 || reviews["gold"] != "prelabelled-relevant" || reviews["confuser"] != "prelabelled-confuser" || reviews["new"] != "unreviewed" {
		t.Fatal("duplicate or unjudged candidate incorrectly scored", reviews)
	}
	reviews = retrievalCandidateReviews(nil, []string{"confuser"}, []string{"new"})
	if reviews["new"] != "unreviewed" {
		t.Fatal("incomplete gold labels must not make every new candidate a false positive")
	}
}
