package app

import (
	"archive/tar"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"
)

const lookupReplayCorpus = "5f12fc3906051e40228959e9e6bfaed714e0658e"

type lookupReplayAttempt struct {
	ID       string        `json:"id"`
	Dir      string        `json:"dir"`
	Handle   string        `json:"handle"`
	Prepared LookupReceipt `json:"prepared"`
}
type lookupReplayManifest struct {
	Binary      string                `json:"binary"`
	BinarySHA   string                `json:"binary_sha256"`
	Corpus      string                `json:"corpus"`
	Revision    string                `json:"corpus_revision"`
	ArchiveSHA  string                `json:"archive_sha256"`
	IndexSHA    string                `json:"index_sha256"`
	DatasetSHAs map[string]string     `json:"dataset_sha256"`
	Attempts    []lookupReplayAttempt `json:"attempts"`
	WarmP95     float64               `json:"warm_prepare_p95_ms"`
}

func lookupReplayRoot(t *testing.T) (string, string, string) {
	t.Helper()
	dir := os.Getenv("AI_SKILL_LOOKUP_REPLAY")
	if dir == "" {
		t.Skip("explicit production replay required")
	}
	root := repoRootForTest(t)
	return dir, root, filepath.Join(root, "plans/active/2026-10-23-2301-knowledge-retrieval-workflow-closure/evidence")
}
func lookupReplayWrite(t *testing.T, p string, v any) {
	t.Helper()
	if _, err := lookupWriteNew(p, v); err != nil {
		t.Fatal(err)
	}
}
func lookupReplayCommand(t *testing.T, binary string, args ...string) Result {
	t.Helper()
	output, err := exec.Command(binary, args...).CombinedOutput()
	if err != nil {
		t.Fatalf("production CLI failed: %v %s", err, output)
	}
	var r Result
	if err = json.Unmarshal(output, &r); err != nil || r.ExitCode != 0 {
		t.Fatalf("invalid CLI result %v %s", err, output)
	}
	return r
}
func TestLookupReplayPrepare(t *testing.T) {
	dir, root, e := lookupReplayRoot(t)
	binary := os.Getenv("AI_SKILL_LOOKUP_BINARY")
	expectedSHA := os.Getenv("AI_SKILL_LOOKUP_BINARY_SHA")
	sha, err := lookupFileSHA(binary)
	if err != nil || expectedSHA == "" || sha != expectedSHA {
		t.Fatal("explicit binary SHA required", err)
	}
	if err = os.Mkdir(dir, 0700); err != nil {
		t.Fatal(err)
	}
	corpus := filepath.Join(dir, "corpus")
	os.Mkdir(corpus, 0700)
	cmd := exec.Command("git", "-C", root, "archive", lookupReplayCorpus)
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		t.Fatal(err)
	}
	h := sha256.New()
	reader := io.TeeReader(pipe, h)
	tr := tar.NewReader(reader)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		path := filepath.Join(corpus, filepath.FromSlash(header.Name))
		if !strings.HasPrefix(path, corpus+string(filepath.Separator)) {
			t.Fatal("unsafe archive path")
		}
		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(path, 0755)
		case tar.TypeReg:
			os.MkdirAll(filepath.Dir(path), 0755)
			var f *os.File
			f, err = os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.FileMode(header.Mode))
			if err == nil {
				_, err = io.Copy(f, tr)
				f.Close()
			}
		case tar.TypeSymlink:
			os.MkdirAll(filepath.Dir(path), 0755)
			err = os.Symlink(header.Linkname, path)
		}
		if err != nil {
			t.Fatal(err)
		}
	}
	io.Copy(io.Discard, reader)
	if err = cmd.Wait(); err != nil {
		t.Fatal(err)
	}
	index := filepath.Join(dir, "index.sqlite")
	lookupReplayCommand(t, binary, "runtime", "lookup", "index", "--repo", corpus, "--db", index, "--json")
	indexSHA, _ := lookupFileSHA(index)
	m := lookupReplayManifest{Binary: binary, BinarySHA: sha, Corpus: corpus, Revision: lookupReplayCorpus, ArchiveSHA: fmt.Sprintf("%x", h.Sum(nil)), IndexSHA: indexSHA, DatasetSHAs: map[string]string{}, Attempts: []lookupReplayAttempt{}}
	attemptsDir := filepath.Join(dir, "attempts")
	os.Mkdir(attemptsDir, 0700)
	warmDir := filepath.Join(dir, "warm")
	os.Mkdir(warmDir, 0700)
	times := []float64{}
	for _, name := range []string{"baseline-cases.json", "overlap-controls-v1.json"} {
		raw, err := os.ReadFile(filepath.Join(e, name))
		if err != nil {
			t.Fatal(err)
		}
		m.DatasetSHAs[name] = lookupSHA(raw)
		var set retrievalReplaySet
		if err = json.Unmarshal(raw, &set); err != nil {
			t.Fatal(err)
		}
		for _, c := range set.Cases {
			if c.ProvenanceSHA != "" {
				sum, err := lookupFileSHA(filepath.Join(corpus, filepath.FromSlash(c.Provenance)))
				if err != nil || sum != c.ProvenanceSHA {
					t.Fatal("frozen source mismatch", c.ID, err)
				}
			}
			key := lookupSHA([]byte(c.Query))[:16]
			attemptDir := filepath.Join(attemptsDir, key)
			args := []string{"runtime", "lookup", "prepare", "--repo", corpus, "--db", index, "--keyword", c.Query, "--layer", c.Layer, "--status", c.Status, "--out", attemptDir, "--json"}
			prepared := lookupReplayCommand(t, binary, args...).Lookup
			if prepared == nil {
				t.Fatal("missing prepare receipt")
			}
			m.Attempts = append(m.Attempts, lookupReplayAttempt{ID: c.ID, Dir: attemptDir, Handle: "/root/select_" + key, Prepared: *prepared})
			// Twenty warm prepare invocations, matching the frozen baseline repeat count.
			for rep := 0; rep < 20; rep++ {
				args[len(args)-2] = filepath.Join(warmDir, fmt.Sprintf("%s-%02d", key, rep))
				started := time.Now()
				lookupReplayCommand(t, binary, args...)
				times = append(times, float64(time.Since(started).Microseconds())/1000)
			}
		}
	}
	sort.Float64s(times)
	m.WarmP95 = times[(len(times)*95+99)/100-1]
	lookupReplayWrite(t, filepath.Join(dir, "manifest.json"), m)
	lookupReplayWrite(t, filepath.Join(dir, "warm.json"), times)
	after, _ := lookupFileSHA(binary)
	if after != sha {
		t.Fatal("binary changed")
	}
	t.Logf("prepared %d actual production pools, warm p95 %.3fms", len(m.Attempts), m.WarmP95)
}

type lookupReplayControl struct {
	Armed []string `json:"armed"`
	Stop  bool     `json:"stop"`
}

func TestLookupReplayMonitor(t *testing.T) {
	dir, _, _ := lookupReplayRoot(t)
	raw, err := os.ReadFile(filepath.Join(dir, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m lookupReplayManifest
	if err = json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	actual, _ := lookupFileSHA(m.Binary)
	if actual != m.BinarySHA {
		t.Fatal("binary changed")
	}
	log, err := os.OpenFile(filepath.Join(dir, "monitor.jsonl"), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer log.Close()
	enc := json.NewEncoder(log)
	started := time.Now()
	last := started
	startedCases := map[string]bool{}
	finished := map[string]bool{}
	children := map[string]*exec.Cmd{}
	files := map[string]*os.File{}
	completed := make(chan string, 26)
	events := []map[string]any{}
	defer func() {
		for id, cmd := range children {
			if !finished[id] && cmd.Process != nil {
				cmd.Process.Kill()
			}
		}
		for _, f := range files {
			f.Close()
		}
	}()
	for {
		now := time.Now()
		gap := now.Sub(last).Seconds()
		last = now
		enc.Encode(map[string]any{"utc": now.UTC().Format(time.RFC3339Nano), "elapsed_ms": now.Sub(started).Milliseconds(), "gap_seconds": gap, "started": len(startedCases), "finished": len(finished)})
		log.Sync()
		if gap > 30 {
			t.Fatal("continuous observation gap exceeded")
		}
		if now.Sub(started) > 3000*time.Second {
			t.Fatal("batch deadline exceeded; parent must interrupt remaining handles")
		}
		var control lookupReplayControl
		if raw, err = os.ReadFile(filepath.Join(dir, "control.json")); err == nil {
			if err = json.Unmarshal(raw, &control); err != nil {
				t.Fatal(err)
			}
		}
		for _, id := range control.Armed {
			if startedCases[id] {
				continue
			}
			if len(startedCases)-len(finished) >= 3 {
				break
			}
			var a *lookupReplayAttempt
			for i := range m.Attempts {
				if m.Attempts[i].ID == id {
					a = &m.Attempts[i]
					break
				}
			}
			if a == nil {
				t.Fatal("foreign control ID")
			}
			f, err := os.OpenFile(filepath.Join(a.Dir, "watcher.json"), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
			if err != nil {
				t.Fatal(err)
			}
			files[id] = f
			cmd := exec.Command(m.Binary, "runtime", "selection-watch", "--request", a.Prepared.RequestPath, "--response", filepath.Join(a.Dir, "raw.json"), "--runner-handle", a.Handle, "--timeout", "300s", "--json")
			cmd.Stdout = f
			cmd.Stderr = f
			if err = cmd.Start(); err != nil {
				t.Fatal(err)
			}
			children[id] = cmd
			startedCases[id] = true
			events = append(events, map[string]any{"id": id, "event": "watch_started", "utc": time.Now().UTC().Format(time.RFC3339Nano), "handle": a.Handle})
			go func(id string, cmd *exec.Cmd) { _ = cmd.Wait(); completed <- id }(id, cmd)
		}
		select {
		case id := <-completed:
			finished[id] = true
			files[id].Sync()
			files[id].Close()
			events = append(events, map[string]any{"id": id, "event": "watch_finished", "utc": time.Now().UTC().Format(time.RFC3339Nano)})
			fmt.Println("watch finished", id)
		default:
		}
		if len(finished) == len(m.Attempts) {
			lookupReplayWrite(t, filepath.Join(dir, "monitor-events.json"), events)
			return
		}
		if control.Stop {
			t.Fatal("parent stopped batch; raw and failures preserved")
		}
		time.Sleep(time.Second)
	}
}

func TestLookupReplayComplete(t *testing.T) {
	dir, _, _ := lookupReplayRoot(t)
	raw, err := os.ReadFile(filepath.Join(dir, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m lookupReplayManifest
	if err = json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	for _, a := range m.Attempts {
		r := lookupReplayCommand(t, m.Binary, "runtime", "lookup", "complete", "--bundle", a.Prepared.BundlePath, "--bundle-sha", a.Prepared.BundleSHA, "--response", filepath.Join(a.Dir, "raw.json"), "--receipt", filepath.Join(a.Dir, "watcher.json"), "--json")
		lookupReplayWrite(t, filepath.Join(a.Dir, "complete.json"), r)
	}
	sha, _ := lookupFileSHA(m.Binary)
	if sha != m.BinarySHA {
		t.Fatal("binary changed")
	}
}

// Export exactly the compact Go JSON object counted by the production packer,
// excluding the tool-only receipt envelope. This is the model message artifact.
func TestLookupReplayExportPayloads(t *testing.T) {
	dir, _, _ := lookupReplayRoot(t)
	b, err := os.ReadFile(filepath.Join(dir, "manifest.json"))
	if err != nil {
		t.Fatal(err)
	}
	var m lookupReplayManifest
	if err = json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	for _, a := range m.Attempts {
		raw, err := os.ReadFile(filepath.Join(a.Dir, "complete.json"))
		if err != nil {
			t.Fatal(err)
		}
		var r Result
		if err = json.Unmarshal(raw, &r); err != nil || r.Lookup == nil || r.Lookup.Payload == nil {
			t.Fatal("no successful payload", a.ID, err)
		}
		payload, err := json.Marshal(r.Lookup.Payload)
		if err != nil {
			t.Fatal(err)
		}
		f, err := os.OpenFile(filepath.Join(a.Dir, "payload.json"), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
		if err != nil {
			t.Fatal(err)
		}
		_, err = f.Write(payload)
		f.Close()
		if err != nil {
			t.Fatal(err)
		}
		if lookupPayloadRunes(r.Lookup.Payload) > 2000 {
			t.Fatal("message over budget", a.ID)
		}
	}
}
