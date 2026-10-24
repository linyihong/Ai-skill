package app

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// Admission changes task selection only; it never changes source authority.
func retrievalDispositionDirectIDs(response retrievalAgentResponse) []string {
	direct := map[string]bool{}
	for _, assessment := range response.Assessments {
		direct[assessment.ID] = assessment.Relation == "direct"
	}
	selected := []string{}
	for _, id := range response.Ranking {
		if direct[id] {
			selected = append(selected, id)
		}
	}
	return selected
}

// Score already-captured outputs only; this test never invokes an agent or API.
func TestRetrievalDispositionScore(t *testing.T) {
	output := os.Getenv("AI_SKILL_RETRIEVAL_DISPOSITION_REPORT")
	if output == "" {
		t.Skip("explicit agent pilot report required")
	}
	rootBytes, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatal(err)
	}
	root := string(bytes.TrimSpace(rootBytes))
	evidence := filepath.Join(root, "plans", "active", "2026-10-23-2301-knowledge-retrieval-workflow-closure", "evidence")
	read := func(name string, v any) []byte {
		data, err := os.ReadFile(filepath.Join(evidence, name))
		if err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal(data, v); err != nil {
			t.Fatal(err)
		}
		return data
	}
	var frozen struct {
		Cards map[string]retrievalV3Card `json:"cards"`
		Pools map[string]retrievalV3Pool `json:"pools"`
	}
	poolBytes := read("phase2-v3-fixed-pools.json", &frozen)
	var manifest struct {
		PoolFileSHA string `json:"fixed_pool_file_sha256"`
		Requests    map[string]struct {
			File    string            `json:"file"`
			FileSHA string            `json:"file_sha256"`
			Mapping map[string]string `json:"opaque_to_source"`
		} `json:"requests"`
	}
	manifestBytes := read("phase2-v3-semantic-manifest.json", &manifest)
	if fmt.Sprintf("%x", sha256.Sum256(poolBytes)) != manifest.PoolFileSHA {
		t.Fatal("pool file changed")
	}
	cases := map[string]retrievalReplayCase{}
	for _, name := range []string{"baseline-cases.json", "overlap-controls-v1.json"} {
		var set retrievalReplaySet
		read(name, &set)
		for _, c := range set.Cases {
			cases[c.ID] = c
		}
	}
	observations := []map[string]any{}
	permutations := []map[string]any{}
	var run struct {
		Attempts map[string]struct {
			RawFile    string `json:"raw_file"`
			RawSHA     string `json:"raw_sha256"`
			RequestSHA string `json:"request_sha256"`
		} `json:"attempts"`
	}
	runBytes := read("phase2-v3-agent-run.json", &run)
	for _, caseID := range []string{"R03", "R08", "O02", "O03", "O04", "O06"} {
		item, ok := manifest.Requests[caseID]
		if !ok {
			t.Fatal("missing request", caseID)
		}
		var request struct {
			Candidates []struct {
				ID string `json:"id"`
			} `json:"candidates"`
		}
		requestBytes := read(item.File, &request)
		requestSHA := fmt.Sprintf("%x", sha256.Sum256(requestBytes))
		if requestSHA != item.FileSHA || requestSHA != run.Attempts[caseID].RequestSHA {
			t.Fatal("request hash changed", caseID)
		}
		ids := []string{}
		for _, card := range request.Candidates {
			ids = append(ids, card.ID)
			if _, ok := item.Mapping[card.ID]; !ok {
				t.Fatal("missing opaque mapping")
			}
		}
		raw, err := os.ReadFile(filepath.Join(evidence, run.Attempts[caseID].RawFile))
		if err != nil {
			t.Fatal(err)
		}
		rawSHA := fmt.Sprintf("%x", sha256.Sum256(raw))
		if rawSHA != run.Attempts[caseID].RawSHA {
			t.Fatal("raw output changed", caseID)
		}
		response, err := retrievalValidateAgentRanking(raw, ids)
		if err != nil {
			t.Fatal("invalid raw ranking", caseID, err)
		}
		firstBySource := map[string]QueryResult{}
		for _, member := range frozen.Pools[caseID].Members {
			r := frozen.Cards[member.ID].Result
			r.Rank = member.Rank
			if _, exists := firstBySource[r.SourcePath]; !exists {
				firstBySource[r.SourcePath] = r
			}
		}
		if len(firstBySource) != len(ids) || len(item.Mapping) != len(ids) {
			t.Fatal("source membership mismatch", caseID)
		}
		ranked := []QueryResult{}
		sourceSet := map[string]bool{}
		for _, id := range response.Ranking {
			source := item.Mapping[id]
			r, ok := firstBySource[source]
			if !ok || sourceSet[source] {
				t.Fatal("unknown/duplicate source", caseID)
			}
			sourceSet[source] = true
			ranked = append(ranked, r)
		}
		c := cases[caseID]
		retrievedConfusers, _ := retrievalReplayScore(c.Excluded, retrievalV3Sources(ranked))
		decisions := []map[string]any{}
		admitted := []QueryResult{}
		directIDs := map[string]bool{}
		for _, id := range retrievalDispositionDirectIDs(response) {
			directIDs[id] = true
		}
		for i, id := range response.Ranking {
			relation, reason, explicit := "unknown", "not assessed", false
			for _, a := range response.Assessments {
				if a.ID == id {
					relation, reason, explicit = a.Relation, a.Reason, true
				}
			}
			action := "pending-confirmation"
			switch relation {
			case "direct":
				action = "queue-readback"
				if directIDs[id] {
					admitted = append(admitted, ranked[i])
				}
			case "background":
				action = "optional-background-only"
			case "not-applicable":
				action = "exclude-this-task"
			}
			decisions = append(decisions, map[string]any{"opaque_id": id, "source": ranked[i].SourcePath, "relation": relation, "reason": reason, "explicit_assessment": explicit, "action": action, "verification": "pending workflow verification", "adoption": "not evaluated"})
		}
		permutations = append(permutations, map[string]any{"id": caseID, "retrieved_confuser_hits": retrievedConfusers, "decisions": decisions, "readback_queue": retrievalV3Sources(admitted)})
		ranked = admitted
		sources := retrievalV3Sources(ranked)
		top5 := sources
		if len(top5) > 5 {
			top5 = top5[:5]
		}
		selectionHits, selectionRR := retrievalReplayScore(c.Expected, top5)
		expectedRanks := map[string]int{}
		confuserRanks := map[string]int{}
		for i, source := range sources {
			for _, expected := range c.Expected {
				if source == expected {
					expectedRanks[source] = i + 1
				}
			}
			for _, excluded := range c.Excluded {
				if source == excluded {
					confuserRanks[source] = i + 1
				}
			}
		}
		permutations = append(permutations, map[string]any{"id": caseID, "request_sha256": requestSHA, "raw_sha256": rawSHA, "candidate_sources": len(ids), "valid_complete_permutation": true, "ranked_sources": sources, "expected_source_ranks": expectedRanks, "confuser_source_ranks": confuserRanks, "assessments": response.Assessments})
		for _, compact := range []bool{false, true} {
			results, shortened := retrievalV3Pack(ranked, compact, 2000)
			actual := retrievalV3Sources(results)
			hits, rr := retrievalReplayScore(c.Expected, actual)
			confusers, _ := retrievalReplayScore(c.Excluded, actual)
			packer := "P0"
			if compact {
				packer = "P1"
			}
			observations = append(observations, map[string]any{"id": caseID, "selection": "S3-D1-direct-preview", "packer": packer, "pool_sha256": retrievalV3Hash(frozen.Pools[caseID]), "selection_hits": selectionHits, "selection_rr": selectionRR, "consumer_results": results, "sources": actual, "hits": hits, "rr": rr, "confuser_hits": confusers, "payload_runes": retrievalV3PayloadRunes(results), "shortened_sources": shortened, "information_sufficiency": "not evaluated", "adoption": "not evaluated"})
		}
	}
	protocolSHA, err := retrievalFileSHA(filepath.Join(evidence, "phase2-v3-disposition-protocol.md"))
	if err != nil {
		t.Fatal(err)
	}
	d0SHA, err := retrievalFileSHA(filepath.Join(evidence, "phase2-v3-agent-results.json"))
	if err != nil {
		t.Fatal(err)
	}
	scorerSHA, err := retrievalFileSHA(filepath.Join(root, "scripts", "ai-skill-cli", "internal", "app", "retrieval_disposition_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	scorerRevision, err := retrievalGitRevision(root)
	if err != nil {
		t.Fatal(err)
	}
	retrievalV3Write(t, output, map[string]any{"protocol_sha256": protocolSHA, "d0_results_sha256": d0SHA, "scorer_source_sha256": scorerSHA, "scorer_revision": scorerRevision, "schema_version": 1, "kind": "existing raw response replay; direct-only admission; selected preview pending workflow verification", "pool_file_sha256": manifest.PoolFileSHA, "semantic_manifest_sha256": fmt.Sprintf("%x", sha256.Sum256(manifestBytes)), "agent_run_sha256": fmt.Sprintf("%x", sha256.Sum256(runBytes)), "rank_field": "original v1 bm25 score retained; array order is semantic ranking, no semantic score invented", "model_snapshot": "unavailable", "token_usage": "unavailable", "timing_gate": "not passed: continuation gap in outer observations; inference latency unavailable", "verified_context": "not evaluated", "background_use_overrides": []string{}, "permutations": permutations, "observations": observations})
}

func TestRetrievalDispositionNoFallback(t *testing.T) {
	for _, tc := range []struct {
		raw  string
		want string
	}{
		{`{"ranking":["a","b"]}`, "[]"},
		{`{"ranking":["a","b"],"assessments":[{"id":"a","relation":"background","reason":"useful context"},{"id":"b","relation":"unknown","reason":"needs readback"}]}`, "[]"},
		{`{"ranking":["b","a"],"assessments":[{"id":"a","relation":"direct","reason":"direct focus"},{"id":"b","relation":"direct","reason":"direct focus"}]}`, `["b","a"]`},
		{`{"ranking":["a","b"],"assessments":[{"id":"a","relation":"not-applicable","reason":"excluded premise"},{"id":"b","relation":"direct","reason":"direct focus"}]}`, `["b"]`},
	} {
		r, err := retrievalValidateAgentRanking([]byte(tc.raw), []string{"a", "b"})
		if err != nil {
			t.Fatal(err)
		}
		before, _ := json.Marshal(r)
		got, _ := json.Marshal(retrievalDispositionDirectIDs(r))
		after, _ := json.Marshal(r)
		if string(got) != tc.want || !bytes.Equal(before, after) {
			t.Fatalf("admission or input mutation: %s", got)
		}
	}
}
