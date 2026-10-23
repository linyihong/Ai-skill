package app

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

// This trace diagnoses the committed source's candidate selection, not a second
// CLI implementation. Final replay results are obtained from the pinned binary;
// all final fields must equal the source call before this diagnostic is accepted.
func TestRetrievalSelectionTrace(t *testing.T) {
	reportPath := os.Getenv("AI_SKILL_RETRIEVAL_TRACE_REPORT")
	if reportPath == "" {
		t.Skip("explicit trace report required")
	}
	index := os.Getenv("AI_SKILL_RETRIEVAL_INDEX")
	if _, err := retrievalReadIndex(index, os.Getenv("AI_SKILL_RETRIEVAL_INDEX_MANIFEST"), "5f12fc3906051e40228959e9e6bfaed714e0658e"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(os.Getenv("AI_SKILL_RETRIEVAL_CASES"))
	if err != nil {
		t.Fatal(err)
	}
	var set retrievalReplaySet
	if err = json.Unmarshal(data, &set); err != nil {
		t.Fatal(err)
	}
	data, err = os.ReadFile(os.Getenv("AI_SKILL_RETRIEVAL_FINAL_REPORT"))
	if err != nil {
		t.Fatal(err)
	}
	var replay struct {
		Observations []retrievalReplayObservation `json:"observations"`
	}
	if err = json.Unmarshal(data, &replay); err != nil {
		t.Fatal(err)
	}
	final := map[string]retrievalReplayObservation{}
	for _, o := range replay.Observations {
		final[o.ID] = o
	}
	trace := []map[string]any{}
	for _, c := range set.Cases {
		opts := runtimeOptions{limit: 5, layer: c.Layer, statusFilter: c.Status}
		raw, err := runtimeIndexSelectionCandidates(index, c.Query, opts)
		if err != nil {
			t.Fatal(err)
		}
		bounded, err := runRuntimeIndexQuery(index, c.Query, opts)
		if err != nil {
			t.Fatal(err)
		}
		observed := final[c.ID]
		if len(bounded) != len(observed.ConsumerResults) || len(bounded) > 0 && !reflect.DeepEqual(bounded, observed.ConsumerResults) {
			t.Fatalf("binary/source final mismatch: %s", c.ID)
		}
		rawSources := []string{}
		rawDetails := []map[string]any{}
		for i, r := range raw {
			rawSources = append(rawSources, r.SourcePath)
			if i < 5 {
				rawDetails = append(rawDetails, map[string]any{"source": r.SourcePath, "rank": r.Rank, "stage": r.MatchReason, "summary_runes": len([]rune(r.Summary))})
			}
		}
		rawTop := rawSources
		if len(rawTop) > 5 {
			rawTop = rawTop[:5]
		}
		hits, rr := retrievalReplayScore(c.Expected, rawTop)
		confusers, _ := retrievalReplayScore(c.Excluded, rawTop)
		expectedRanks := map[string]int{}
		for _, expected := range c.Expected {
			for rank, source := range rawSources {
				if source == expected {
					expectedRanks[expected] = rank + 1
					break
				}
			}
		}
		truncated := false
		for _, r := range bounded {
			for _, original := range raw {
				if r.ID == original.ID && r.Summary != original.Summary {
					truncated = true
				}
			}
		}
		trace = append(trace, map[string]any{"id": c.ID, "candidate_atoms": len(raw), "pre_bound_top5": rawDetails, "pre_bound_top5_hits": hits, "pre_bound_top5_rr": rr, "pre_bound_top5_confuser_hits": confusers, "expected_atom_ranks_before_bound": expectedRanks, "final_sources": observed.Sources, "final_hits": observed.Hits, "final_rr": observed.RR, "final_payload_runes": observed.PreviewEstimate, "summary_truncated": truncated, "binary_source_final_fields_equal": true})
	}
	payload, err := json.MarshalIndent(map[string]any{"diagnostic_kind": "same committed source candidate call; final fields checked against pinned CLI replay", "index_path": index, "observations": trace}, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err = os.WriteFile(reportPath, append(payload, '\n'), 0644); err != nil {
		t.Fatal(err)
	}
}
