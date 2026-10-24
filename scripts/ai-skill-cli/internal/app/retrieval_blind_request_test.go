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
)

func TestRetrievalBlindRequestsExport(t *testing.T) {
	transport := os.Getenv("AI_SKILL_BLIND_TRANSPORT")
	if transport == "" {
		t.Skip("explicit request export required")
	}
	b, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatal(err)
	}
	root := strings.TrimSpace(string(b))
	e := filepath.Join(root, "plans", "archived", "2026-10-23-2301-knowledge-retrieval-workflow-closure", "evidence")
	read := func(name string, v any) []byte {
		b, err := os.ReadFile(filepath.Join(e, name))
		if err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal(b, v); err != nil {
			t.Fatal(err)
		}
		return b
	}
	var pool struct {
		Cards map[string]retrievalV3Card `json:"cards"`
		Pools map[string]retrievalV3Pool `json:"pools"`
	}
	poolBytes := read("phase2-v3-fixed-pools.json", &pool)
	prompt, err := os.ReadFile(filepath.Join(e, "phase2-blind-selector-prompt.txt"))
	if err != nil {
		t.Fatal(err)
	}
	requests := map[string]any{}
	if err = os.MkdirAll(transport, 0700); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"baseline-cases.json", "overlap-controls-v1.json"} {
		var set retrievalReplaySet
		read(name, &set)
		for _, c := range set.Cases {
			first := map[string]retrievalV3Card{}
			for _, m := range pool.Pools[c.ID].Members {
				card := pool.Cards[m.ID]
				if _, ok := first[card.Result.SourcePath]; !ok {
					first[card.Result.SourcePath] = card
				}
			}
			mapping := map[string]string{}
			for source := range first {
				id := fmt.Sprintf("%x", sha256.Sum256([]byte("blind-selector-v1:"+c.ID+":"+source)))[:16]
				if _, ok := mapping[id]; ok {
					t.Fatal("opaque collision")
				}
				mapping[id] = source
			}
			ids := []string{}
			for id := range mapping {
				ids = append(ids, id)
			}
			sort.Strings(ids)
			cards := []map[string]string{}
			for _, id := range ids {
				card := first[mapping[id]]
				cards = append(cards, map[string]string{"id": id, "title": card.Title, "summary": card.Result.Summary})
			}
			request := map[string]any{"system": string(prompt), "query": c.Query, "candidates": cards}
			data, err := json.MarshalIndent(request, "", "  ")
			if err != nil {
				t.Fatal(err)
			}
			data = bytes.ReplaceAll(data, []byte(`\\`), []byte(`\u005c`))
			data = append(data, '\n')
			sha := fmt.Sprintf("%x", sha256.Sum256(data))
			file := "phase2-blind-request-" + sha[:16] + ".json"
			for _, path := range []string{filepath.Join(e, file), filepath.Join(transport, sha[:16]+".json")} {
				if _, err = os.Stat(path); !os.IsNotExist(err) {
					t.Fatal("refuse overwrite", path)
				}
				if err = os.WriteFile(path, data, 0600); err != nil {
					t.Fatal(err)
				}
			}
			requests[c.ID] = map[string]any{"file": file, "sha256": sha, "bytes": len(data), "source_count": len(mapping), "opaque_to_source": mapping, "transport_file": sha[:16] + ".json", "response_file": sha[:16] + ".raw.txt"}
		}
	}
	sha, err := retrievalFileSHA(filepath.Join(root, "scripts", "ai-skill-cli", "internal", "app", "retrieval_blind_request_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	retrievalV3Write(t, filepath.Join(e, "phase2-blind-selector-manifest.json"), map[string]any{"kind": "query and fixed source-deduplicated cards only; exposed dataset, gold-blind execution", "pool_file_sha256": fmt.Sprintf("%x", sha256.Sum256(poolBytes)), "builder_sha256": sha, "requests": requests})
}
