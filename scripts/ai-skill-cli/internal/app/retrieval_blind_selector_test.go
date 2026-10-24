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

func TestRetrievalBlindSelectionControls(t *testing.T) {
	for _, tc := range []struct {
		raw   string
		valid bool
	}{
		{`{"selected":[],"assessments":[],"unknown_action":"pool empty"}`, true},
		{`{"selected":["a"],"assessments":[{"id":"a","relation":"direct","reason":"focus","need_to_read":"check scope"}],"unknown_action":"verify"}`, true},
		{`{"selected":["a"],"assessments":[{"id":"a","relation":"unknown","reason":"scope unclear","need_to_read":"identify provider"}],"unknown_action":"do not apply"}`, false},
		{`{"selected":[],"assessments":[{"id":"a","relation":"direct","reason":"focus","need_to_read":"check scope"}],"unknown_action":"verify"}`, false},
		{`{"selected":["x"],"assessments":[],"unknown_action":"verify"}`, false},
		{`{"selected":[],"assessments":[],"unknown_action":""}`, false},
		{`{"selected":[],"unknown_action":"missing assessments"}`, false},
		{`{"selected":["a"],"assessments":[{"id":"a","relation":"direct","reason":"focus","need_to_read":""}],"unknown_action":"verify"}`, false},
	} {
		_, err := retrievalParseBlind([]byte(tc.raw), []string{"a"})
		if (err == nil) != tc.valid {
			t.Fatalf("%s: %v", tc.raw, err)
		}
	}
	ranked := []retrievalBlindModelResult{{Source: "a", Summary: strings.Repeat("字", 300), Need: "inspect missing prerequisite"}, {Source: "b", Summary: strings.Repeat("z", 300), Need: "verify"}}
	p, d, err := retrievalBlindPack("T", retrievalCoverage{Candidates: 5, Unknown: 3, Unassessed: 3, State: "direct-pending-verification"}, ranked, 700)
	if err != nil || len(p.Results) != 1 || len(d) != 1 || p.Results[0].Summary != ranked[0].Summary || p.Results[0].Need != ranked[0].Need || p.Unknown != 3 || retrievalV3PayloadRunes(p) > 700 {
		t.Fatal("lost handoff or budget", p, d, err)
	}
}

func TestRetrievalBlindSelectorScore(t *testing.T) {
	output := os.Getenv("AI_SKILL_BLIND_REPORT")
	if output == "" {
		t.Skip("explicit captured-run score required")
	}
	b, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatal(err)
	}
	root := strings.TrimSpace(string(b))
	e := filepath.Join(root, "plans", "archived", "2026-10-23-2301-knowledge-retrieval-workflow-closure", "evidence")
	inputs := map[string]string{}
	read := func(name string, v any) {
		b, err := os.ReadFile(filepath.Join(e, name))
		if err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal(b, v); err != nil {
			t.Fatal(err)
		}
		inputs[name], err = retrievalFileSHA(filepath.Join(e, name))
		if err != nil {
			t.Fatal(err)
		}
	}
	var frozen struct {
		Cards map[string]retrievalV3Card `json:"cards"`
		Pools map[string]retrievalV3Pool `json:"pools"`
	}
	read("phase2-v3-fixed-pools.json", &frozen)
	var manifest struct {
		PoolSHA  string `json:"pool_file_sha256"`
		Requests map[string]struct {
			File, SHA string
			FileSHA   string            `json:"sha256"`
			Mapping   map[string]string `json:"opaque_to_source"`
		} `json:"requests"`
	}
	read("phase2-blind-selector-manifest.json", &manifest)
	if manifest.PoolSHA != inputs["phase2-v3-fixed-pools.json"] {
		t.Fatal("pool changed")
	}
	var run struct {
		Attempts map[string]struct {
			Raw        string `json:"raw_file"`
			SHA        string `json:"raw_sha256"`
			RequestSHA string `json:"request_sha256"`
			Status     string `json:"status"`
		} `json:"attempts"`
	}
	read("phase2-blind-selector-run.json", &run)
	observations := []map[string]any{}
	for _, name := range []string{"baseline-cases.json", "overlap-controls-v1.json"} {
		var set retrievalReplaySet
		read(name, &set)
		for _, c := range set.Cases {
			i, ok := manifest.Requests[c.ID]
			if !ok {
				t.Fatal("missing request")
			}
			var request struct {
				Candidates []struct {
					ID string `json:"id"`
				} `json:"candidates"`
			}
			read(i.File, &request)
			if inputs[i.File] != i.FileSHA {
				t.Fatal("request changed")
			}
			ids := []string{}
			for _, card := range request.Candidates {
				ids = append(ids, card.ID)
			}
			a, ok := run.Attempts[c.ID]
			if !ok {
				t.Fatal("unaccounted attempt")
			}
			if a.RequestSHA != i.FileSHA {
				t.Fatal("attempt request mismatch")
			}
			if a.Status != "raw-captured" {
				observations = append(observations, map[string]any{"id": c.ID, "status": a.Status, "parser_valid": false, "quality": "not evaluated"})
				continue
			}
			raw, err := os.ReadFile(filepath.Join(e, a.Raw))
			if err != nil {
				t.Fatal(err)
			}
			sha, err := retrievalFileSHA(filepath.Join(e, a.Raw))
			if err != nil || sha != a.SHA {
				t.Fatal("raw changed")
			}
			inputs[a.Raw] = sha
			r, err := retrievalParseBlind(raw, ids)
			if err != nil {
				observations = append(observations, map[string]any{"id": c.ID, "parser_valid": false, "parser_error": err.Error(), "quality": "not evaluated"})
				continue
			}
			first := map[string]retrievalV3Card{}
			sources := []string{}
			for _, m := range frozen.Pools[c.ID].Members {
				card := frozen.Cards[m.ID]
				if _, ok := first[card.Result.SourcePath]; !ok {
					first[card.Result.SourcePath] = card
					sources = append(sources, card.Result.SourcePath)
				}
			}
			if len(first) != len(ids) || len(i.Mapping) != len(ids) {
				t.Fatal("source membership mismatch")
			}
			for _, id := range ids {
				if _, ok := first[i.Mapping[id]]; !ok {
					t.Fatal("unknown source")
				}
			}
			coverage := retrievalCoverage{Candidates: len(ids), Assessed: len(r.Assessments), Unassessed: len(ids) - len(r.Assessments)}
			needs := map[string]string{}
			for _, a := range r.Assessments {
				needs[a.ID] = a.Need
				switch a.Relation {
				case "direct":
					coverage.Direct++
				case "background":
					coverage.Background++
				case "not-applicable":
					coverage.NotApplicable++
				case "unknown":
					coverage.ExplicitUnknown++
				}
			}
			coverage.Unknown = coverage.Unassessed + coverage.ExplicitUnknown
			switch {
			case coverage.Direct > 0:
				coverage.State = "direct-pending-verification"
			case coverage.Unknown > 0:
				coverage.State = "pending-assessment"
			case coverage.Candidates == 0:
				coverage.State = "no-candidates-in-fixed-pool"
			default:
				coverage.State = "no-direct-in-reviewed-candidates"
			}
			ranked := []retrievalBlindModelResult{}
			queue := []map[string]any{}
			for _, id := range r.Selected {
				card := first[i.Mapping[id]]
				v := card.Result
				ranked = append(ranked, retrievalBlindModelResult{v.SourcePath, v.Type, v.Status, v.Summary, needs[id]})
				blob, err := exec.Command("git", "-C", root, "show", card.Revision+":"+v.SourcePath).Output()
				if err != nil || fmt.Sprintf("%x", sha256.Sum256(blob)) != card.Checksum {
					t.Fatal("fixed source readback mismatch", v.SourcePath)
				}
				currentSHA, currentErr := retrievalFileSHA(filepath.Join(root, v.SourcePath))
				currentState := "same-version"
				if currentErr != nil || currentSHA != card.Checksum {
					currentState = "requires-revalidation"
				}
				queue = append(queue, map[string]any{"source": v.SourcePath, "revision": card.Revision, "checksum": card.Checksum, "need_to_read": needs[id], "fixed_blob_verified": true, "full_source_bytes": len(blob), "current_checksum": currentSHA, "current_version_state": currentState})
			}
			payload, deferred, err := retrievalBlindPack(c.ID, coverage, ranked, 2000)
			if err != nil {
				t.Fatal(err)
			}
			actual := []string{}
			for _, v := range payload.Results {
				actual = append(actual, v.Source)
			}
			hits, rr := retrievalReplayScore(c.Expected, actual)
			confusers, _ := retrievalReplayScore(c.Excluded, actual)
			retrieved, _ := retrievalReplayScore(c.Excluded, sources)
			observations = append(observations, map[string]any{"id": c.ID, "parser_valid": true, "request_sha256": i.FileSHA, "raw_sha256": sha, "coverage": coverage, "model_payload": payload, "model_payload_runes": retrievalV3PayloadRunes(payload), "selected_sources": actual, "readback_queue": queue, "budget_deferred_sources": deferred, "retrieved_confuser_hits": retrieved, "selected_confuser_hits": confusers, "hits": hits, "rr": rr, "unknown_action": r.UnknownAction, "assessments": r.Assessments, "adoption": "not evaluated"})
		}
	}
	sha, err := retrievalFileSHA(filepath.Join(root, "scripts", "ai-skill-cli", "internal", "app", "retrieval_blind_selector_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	revision, err := retrievalGitRevision(root)
	if err != nil {
		t.Fatal(err)
	}
	retrievalV3Write(t, output, map[string]any{"kind": "gold-blind execution on exposed diagnostic data; fresh existing runner, no manual decisions supplied", "inputs_sha256": inputs, "scorer_sha256": sha, "scorer_revision": revision, "observations": observations, "actual_tokens": "unavailable", "production_integration": "not evaluated", "phase2": "incomplete"})
}
