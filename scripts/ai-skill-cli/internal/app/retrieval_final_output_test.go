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

type retrievalModelResult struct {
	Source  string `json:"source"`
	Kind    string `json:"kind"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
}

type retrievalModelPayload struct {
	Guard      string                 `json:"guard"`
	State      string                 `json:"state"`
	Pool       int                    `json:"pool"`
	Unassessed int                    `json:"unassessed"`
	Unknown    int                    `json:"unknown"`
	Deferred   int                    `json:"deferred"`
	Trace      string                 `json:"trace"`
	Results    []retrievalModelResult `json:"results"`
}

// This is an offline prototype, not a production workflow consumer.
func retrievalPackModelPayload(caseID string, coverage retrievalCoverage, ranked []QueryResult, budget int) (retrievalModelPayload, []string, error) {
	p := retrievalModelPayload{Guard: "Candidates only; read sources and verify applicability; empty does not establish absence.", State: coverage.State, Pool: coverage.Candidates, Unassessed: coverage.Unassessed, Unknown: coverage.Unknown, Trace: "final-output/" + caseID, Results: []retrievalModelResult{}}
	for _, r := range ranked {
		p.Results = append(p.Results, retrievalModelResult{r.SourcePath, r.Type, r.Status, r.Summary})
	}
	keep := len(p.Results)
	for {
		p.Deferred = len(ranked) - keep
		p.Results = p.Results[:keep]
		if retrievalV3PayloadRunes(p) <= budget {
			break
		}
		if keep == 0 {
			return p, nil, fmt.Errorf("required state alone exceeds budget")
		}
		keep--
	}
	deferred := []string{}
	for _, r := range ranked[keep:] {
		deferred = append(deferred, r.SourcePath)
	}
	return p, deferred, nil
}

func TestRetrievalModelPayloadBudgetAndUnknown(t *testing.T) {
	ranked := []QueryResult{{SourcePath: "first", Type: "lesson", Status: "candidate", Summary: strings.Repeat("知", 300)}, {SourcePath: "second", Summary: strings.Repeat("x", 300)}}
	c := retrievalCoverage{Candidates: 9, Unassessed: 6, Unknown: 7, Direct: 2, State: "direct-pending-verification"}
	p, d, err := retrievalPackModelPayload("T", c, ranked, 700)
	if err != nil || len(p.Results) != 1 || len(d) != 1 || d[0] != "second" || p.Deferred != 1 || p.Unknown != 7 || p.Unassessed != 6 || p.Results[0].Summary != ranked[0].Summary || p.Results[0].Status != "candidate" || retrievalV3PayloadRunes(p) > 700 {
		t.Fatalf("bad payload %+v %v %v", p, d, err)
	}
	if _, _, err = retrievalPackModelPayload("T", c, ranked, 10); err == nil {
		t.Fatal("must not hide necessary state to fit")
	}
	p, _, err = retrievalPackModelPayload("T", retrievalCoverage{Candidates: 2, Unknown: 2, Unassessed: 2, State: "pending-assessment"}, nil, 2000)
	if err != nil || len(p.Results) != 0 || p.Unknown != 2 || p.Guard == "" || p.State != "pending-assessment" {
		t.Fatal("empty must retain boundary")
	}
}

func TestRetrievalFinalOutputExport(t *testing.T) {
	output := os.Getenv("AI_SKILL_RETRIEVAL_FINAL_REPORT")
	if output == "" {
		t.Skip("explicit offline report required")
	}
	b, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatal(err)
	}
	root := strings.TrimSpace(string(b))
	evidence := filepath.Join(root, "plans", "archived", "2026-10-23-2301-knowledge-retrieval-workflow-closure", "evidence")
	inputs := map[string]string{}
	read := func(name string, v any) {
		data, err := os.ReadFile(filepath.Join(evidence, name))
		if err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal(data, v); err != nil {
			t.Fatal(err)
		}
		inputs[name], err = retrievalFileSHA(filepath.Join(evidence, name))
		if err != nil {
			t.Fatal(err)
		}
	}
	var frozen struct {
		Cards map[string]retrievalV3Card `json:"cards"`
		Pools map[string]retrievalV3Pool `json:"pools"`
	}
	read("phase2-v3-fixed-pools.json", &frozen)
	var decisions struct {
		Cases []struct {
			ID            string                                      `json:"id"`
			Order         []string                                    `json:"selected_order"`
			Assessments   []struct{ Source, Relation, Reason string } `json:"assessments"`
			Needed        string                                      `json:"information_needed"`
			UnknownAction string                                      `json:"unknown_disposition"`
		} `json:"cases"`
	}
	read("phase2-final-output-decisions.json", &decisions)
	var audit struct {
		Sources []struct {
			Source, Revision string
			Checksum         string `json:"readback_checksum"`
			Expected         string `json:"expected_checksum"`
			Current          string `json:"current_checksum"`
			Scope            string `json:"scope_review"`
		} `json:"sources"`
	}
	read("phase2-final-output-readback.json", &audit)
	audits := map[string]int{}
	for i, a := range audit.Sources {
		if _, ok := audits[a.Source]; ok {
			t.Fatal("duplicate audit")
		}
		audits[a.Source] = i
	}
	cases := map[string]retrievalReplayCase{}
	for _, name := range []string{"baseline-cases.json", "overlap-controls-v1.json"} {
		var set retrievalReplaySet
		read(name, &set)
		for _, c := range set.Cases {
			cases[c.ID] = c
		}
	}
	seenCases := map[string]bool{}
	observations := []map[string]any{}
	for _, d := range decisions.Cases {
		c, ok := cases[d.ID]
		if !ok || seenCases[d.ID] {
			t.Fatal("unknown/duplicate case")
		}
		seenCases[d.ID] = true
		first := map[string]retrievalV3Card{}
		sources := []string{}
		for _, m := range frozen.Pools[d.ID].Members {
			card := frozen.Cards[m.ID]
			if _, ok := first[card.Result.SourcePath]; !ok {
				first[card.Result.SourcePath] = card
				sources = append(sources, card.Result.SourcePath)
			}
		}
		rawResponse := retrievalAgentResponse{Ranking: sources}
		// Existing strict parser applies to the sparse explicit decisions.
		responseJSON := map[string]any{"ranking": sources, "assessments": []map[string]string{}}
		ar := []map[string]string{}
		for _, a := range d.Assessments {
			ar = append(ar, map[string]string{"id": a.Source, "relation": a.Relation, "reason": a.Reason})
		}
		responseJSON["assessments"] = ar
		encoded, _ := json.Marshal(responseJSON)
		rawResponse, err = retrievalValidateAgentRanking(encoded, sources)
		if err != nil {
			t.Fatal(d.ID, err)
		}
		coverage := retrievalCoverageFor(rawResponse)
		if coverage.State == "no-direct-in-assessed-pool" {
			coverage.State = "no-direct-in-reviewed-candidates"
		}
		direct := map[string]bool{}
		for _, id := range retrievalDispositionDirectIDs(rawResponse) {
			direct[id] = true
		}
		if len(d.Order) != len(direct) {
			t.Fatal("direct/order mismatch")
		}
		ranked := []QueryResult{}
		selected := map[string]bool{}
		for _, source := range d.Order {
			if !direct[source] || selected[source] {
				t.Fatal("foreign/duplicate/non-direct selection")
			}
			selected[source] = true
			card := first[source]
			idx, ok := audits[source]
			if !ok {
				t.Fatal("selected source lacks review")
			}
			a := audit.Sources[idx]
			blob, err := exec.Command("git", "-C", root, "show", card.Revision+":"+source).Output()
			if err != nil {
				t.Fatal(err)
			}
			if fmt.Sprintf("%x", sha256.Sum256(blob)) != card.Checksum || a.Checksum != card.Checksum || a.Expected != card.Checksum || a.Revision != card.Revision || a.Scope == "" {
				t.Fatal("readback identity/scope mismatch")
			}
			current, err := retrievalFileSHA(filepath.Join(root, source))
			if err != nil || current != a.Current {
				t.Fatal("source changed since review")
			}
			ranked = append(ranked, card.Result)
		}
		payload, deferred, err := retrievalPackModelPayload(d.ID, coverage, ranked, 2000)
		if err != nil {
			t.Fatal(err)
		}
		actual := []string{}
		for _, r := range payload.Results {
			actual = append(actual, r.Source)
		}
		hits, rr := retrievalReplayScore(c.Expected, actual)
		confusers, _ := retrievalReplayScore(c.Excluded, actual)
		retrieved, _ := retrievalReplayScore(c.Excluded, sources)
		observations = append(observations, map[string]any{"id": d.ID, "trace": payload.Trace, "query": c.Query, "coverage": coverage, "model_payload": payload, "model_payload_runes": retrievalV3PayloadRunes(payload), "selected_sources": actual, "budget_deferred_sources": deferred, "retrieved_confuser_hits": retrieved, "hits": hits, "rr": rr, "selected_confuser_hits": confusers, "information_needed": d.Needed, "unknown_disposition": d.UnknownAction, "assessments": d.Assessments, "source_pool_ref": "phase2-v3-fixed-pools.json#pools/" + d.ID, "adoption": "not evaluated"})
	}
	if len(seenCases) != len(cases) {
		t.Fatal("incomplete case coverage")
	}
	sha, err := retrievalFileSHA(filepath.Join(root, "scripts", "ai-skill-cli", "internal", "app", "retrieval_final_output_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	revision, err := retrievalGitRevision(root)
	if err != nil {
		t.Fatal(err)
	}
	protocol, err := retrievalFileSHA(filepath.Join(evidence, "phase2-final-output-protocol.md"))
	if err != nil {
		t.Fatal(err)
	}
	retrievalV3Write(t, output, map[string]any{"kind": "current-agent known-gold exposed diagnostic, not controlled inference or production integration", "scorer_revision": revision, "scorer_sha256": sha, "protocol_sha256": protocol, "inputs_sha256": inputs, "observations": observations, "actual_model_tokens": "not evaluated", "controlled_inference_time": "not evaluated", "phase2": "incomplete", "adoption": "not evaluated"})
}
