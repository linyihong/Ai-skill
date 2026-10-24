package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLookupPrepareIsProductionCandidateConsumer(t *testing.T) {
	repo := t.TempDir()
	writeFeedbackHistoryRouteFixture(t, repo)
	writeFile(t, filepath.Join(repo, "feedback/history/demo/legacy.md"), "# Legacy clue\nStatus: candidate\n## One-line Summary\n\nneedle legacy knowledge\n")
	db := filepath.Join(t.TempDir(), "index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
		t.Fatal(err)
	}
	var out, errout bytes.Buffer
	exit := runRuntime([]string{"lookup", "prepare", "--repo", repo, "--db", db, "--keyword", "needle differently phrased", "--out", filepath.Join(t.TempDir(), "attempt"), "--json"}, &out, &errout)
	var result Result
	if err := json.Unmarshal(out.Bytes(), &result); err != nil || exit != 0 {
		t.Fatalf("lookup absent or failed: %d %s %s %v", exit, out.String(), errout.String(), err)
	}
}

func TestLookupEndToEndAndFailureDispositions(t *testing.T) {
	repo := t.TempDir()
	writeFeedbackHistoryRouteFixture(t, repo)
	source := filepath.Join(repo, "feedback/history/demo/legacy.md")
	content := "# Legacy clue\nStatus: candidate\n## One-line Summary\n\nneedle legacy knowledge\n"
	writeFile(t, source, content)
	writeFile(t, filepath.Join(repo, "feedback/history/demo/other.md"), "# Other\nStatus: candidate\n## One-line Summary\nneedle background\n")
	db := filepath.Join(t.TempDir(), "index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, db); err != nil {
		t.Fatal(err)
	}
	dir := filepath.Join(t.TempDir(), "attempt")
	prepared, err := lookupPrepare(repo, db, "needle differently phrased", "", "", dir)
	if err != nil {
		t.Fatal(err)
	}
	var bundle LookupBundle
	b, err := os.ReadFile(prepared.BundlePath)
	if err != nil {
		t.Fatal(err)
	}
	if err = json.Unmarshal(b, &bundle); err != nil {
		t.Fatal(err)
	}
	if len(bundle.Candidates) != 2 {
		t.Fatalf("production pool not used %+v", bundle)
	}
	id := ""
	for _, c := range bundle.Candidates {
		if c.Result.SourcePath == "feedback/history/demo/legacy.md" {
			id = c.ID
		}
	}
	if id == "" {
		t.Fatal("legacy heading extraction lost source")
	}
	response := filepath.Join(dir, "response.json")
	raw := fmt.Sprintf(`{"selected":[%q],"assessments":[{"id":%q,"relation":"direct","reason":"query focus","need_to_read":"read trigger and verify project conditions"}],"unknown_action":"one candidate remains unassessed"}`, id, id)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 1000; i++ {
			if _, err := os.Stat(response + ".ready"); err == nil {
				os.WriteFile(response+".tmp", []byte(raw), 0600)
				os.Rename(response+".tmp", response)
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()
	var out, errout bytes.Buffer
	exit := runRuntime([]string{"selection-watch", "--request", prepared.RequestPath, "--response", response, "--runner-handle", "fixture-child", "--timeout", "2s", "--json"}, &out, &errout)
	<-done
	if exit != 0 {
		t.Fatalf("watch failed %s %s", out.String(), errout.String())
	}
	receipt := filepath.Join(dir, "watcher.json")
	writeFile(t, receipt, out.String())
	completed, err := lookupComplete(prepared.BundlePath, prepared.BundleSHA, response, receipt)
	if err != nil {
		t.Fatal(err)
	}
	if completed.Payload == nil || len(completed.Payload.Results) != 1 || completed.Payload.Unassessed != 1 || completed.Payload.Unknown != 1 || completed.Payload.Results[0].Summary != "needle legacy knowledge" || completed.Payload.Results[0].Status != "candidate" || lookupPayloadRunes(completed.Payload) > 2000 {
		t.Fatalf("bad payload %+v", completed)
	}
	if _, err = lookupComplete(prepared.BundlePath, "wrong-sha", response, receipt); err == nil {
		t.Fatal("changed bundle admitted")
	}
	originalRequest, err := os.ReadFile(prepared.RequestPath)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, prepared.RequestPath, "{}")
	if _, err = lookupComplete(prepared.BundlePath, prepared.BundleSHA, response, receipt); err == nil {
		t.Fatal("changed request admitted")
	}
	os.WriteFile(prepared.RequestPath, originalRequest, 0600)
	originalReceipt, err := os.ReadFile(receipt)
	if err != nil {
		t.Fatal(err)
	}
	var watch Result
	json.Unmarshal(originalReceipt, &watch)
	watch.SelectionWatch.State = "timeout"
	watch.SelectionWatch.InterruptRequired = true
	changed, _ := json.Marshal(watch)
	os.WriteFile(receipt, changed, 0600)
	if _, err = lookupComplete(prepared.BundlePath, prepared.BundleSHA, response, receipt); err == nil {
		t.Fatal("failed watcher treated as success")
	}
	os.WriteFile(receipt, originalReceipt, 0600)
	writeFile(t, response, `{"selected":[],"assessments":[],"unknown_action":"changed"}`)
	if _, err = lookupComplete(prepared.BundlePath, prepared.BundleSHA, response, receipt); err == nil {
		t.Fatal("changed raw admitted")
	}
	writeFile(t, response, raw)
	writeFile(t, source, content+"\nSource changed after selection.\n")
	if _, err = lookupComplete(prepared.BundlePath, prepared.BundleSHA, response, receipt); err == nil {
		t.Fatal("stale source adopted")
	}
}

func TestLookupPromptAndPackerPolicyUnchanged(t *testing.T) {
	root := repoRootForTest(t)
	b, err := os.ReadFile(filepath.Join(root, "plans/archived/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence/phase2-blind-selector-prompt.txt"))
	if err != nil || string(b) != lookupSelectorPrompt {
		t.Fatal("selector prompt drift", err)
	}
	ranked := []retrievalBlindModelResult{{Source: "first", Summary: strings.Repeat("字", 1100), Need: "read first source"}, {Source: "second", Summary: strings.Repeat("字", 1100), Need: "read second source"}}
	p, deferred, err := lookupPack("lookup/fixture", retrievalCoverage{Candidates: 10, Unassessed: 8, Unknown: 8, State: "direct-pending-verification"}, ranked, 2000)
	if err != nil || len(p.Results) != 1 || p.Results[0].Summary != ranked[0].Summary || p.Results[0].Need != ranked[0].Need || len(deferred) != 1 || deferred[0] != "second" || p.Unknown != 8 || lookupPayloadRunes(p) > 2000 {
		t.Fatalf("packer policy changed %+v %v %v", p, deferred, err)
	}
}
