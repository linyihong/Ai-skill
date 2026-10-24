package app

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Offline consumer-preview prototype only. No workflow calls this helper.
type retrievalCoverage struct {
	Candidates      int    `json:"candidates"`
	Assessed        int    `json:"assessed"`
	Unassessed      int    `json:"unassessed"`
	ExplicitUnknown int    `json:"explicit_unknown"`
	Unknown         int    `json:"unknown"`
	Direct          int    `json:"direct"`
	Background      int    `json:"background"`
	NotApplicable   int    `json:"not_applicable"`
	State           string `json:"state"`
}

func retrievalCoverageFor(response retrievalAgentResponse) retrievalCoverage {
	c := retrievalCoverage{Candidates: len(response.Ranking), Assessed: len(response.Assessments)}
	c.Unassessed = c.Candidates - c.Assessed
	for _, a := range response.Assessments {
		switch a.Relation {
		case "direct":
			c.Direct++
		case "background":
			c.Background++
		case "not-applicable":
			c.NotApplicable++
		case "unknown":
			c.ExplicitUnknown++
		}
	}
	c.Unknown = c.Unassessed + c.ExplicitUnknown
	switch {
	case c.Direct > 0:
		c.State = "direct-pending-verification"
	case c.Unknown > 0:
		c.State = "pending-assessment"
	case c.Candidates == 0:
		c.State = "no-candidates-in-fixed-pool"
	default:
		c.State = "no-direct-in-assessed-pool"
	}
	return c
}

func TestRetrievalCoverageEmptyDoesNotClaimAbsence(t *testing.T) {
	for _, tc := range []struct {
		raw, state          string
		unknown, unassessed int
	}{
		{`{"ranking":[]}`, "no-candidates-in-fixed-pool", 0, 0},
		{`{"ranking":["a"]}`, "pending-assessment", 1, 1},
		{`{"ranking":["a"],"assessments":[{"id":"a","relation":"unknown","reason":"needs source"}]}`, "pending-assessment", 1, 0},
		{`{"ranking":["a"],"assessments":[{"id":"a","relation":"not-applicable","reason":"wrong scope"}]}`, "no-direct-in-assessed-pool", 0, 0},
		{`{"ranking":["a","b"],"assessments":[{"id":"a","relation":"direct","reason":"matches focus"}]}`, "direct-pending-verification", 1, 1},
	} {
		var input retrievalAgentResponse
		if err := json.Unmarshal([]byte(tc.raw), &input); err != nil {
			t.Fatal(err)
		}
		validated, err := retrievalValidateAgentRanking([]byte(tc.raw), input.Ranking)
		if err != nil {
			t.Fatal(err)
		}
		c := retrievalCoverageFor(validated)
		if c.State != tc.state || c.Unknown != tc.unknown || c.Unassessed != tc.unassessed {
			t.Fatalf("coverage: %+v", c)
		}
		if c.Direct+c.Background+c.NotApplicable+c.Unknown != c.Candidates {
			t.Fatal("lost candidate")
		}
	}
}

func TestRetrievalPhase2ReadinessExport(t *testing.T) {
	output := os.Getenv("AI_SKILL_RETRIEVAL_READINESS_REPORT")
	if output == "" {
		t.Skip("explicit offline report required")
	}
	b, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		t.Fatal(err)
	}
	root := strings.TrimSpace(string(b))
	evidence := filepath.Join(root, "plans", "active", "2026-10-23-2301-knowledge-retrieval-workflow-closure", "evidence")
	inputs := map[string]string{}
	read := func(name string, v any) {
		data, err := os.ReadFile(filepath.Join(evidence, name))
		if err != nil {
			t.Fatal(err)
		}
		if err = json.Unmarshal(data, v); err != nil {
			t.Fatal(err)
		}
		sha, err := retrievalFileSHA(filepath.Join(evidence, name))
		if err != nil {
			t.Fatal(err)
		}
		inputs[name] = sha
	}
	var frozen struct {
		Cards map[string]retrievalV3Card `json:"cards"`
		Pools map[string]retrievalV3Pool `json:"pools"`
	}
	read("phase2-v3-fixed-pools.json", &frozen)
	var d1 struct {
		Observations []struct {
			ID      string        `json:"id"`
			Packer  string        `json:"packer"`
			Results []QueryResult `json:"consumer_results"`
		} `json:"observations"`
	}
	read("phase2-v3-disposition-results.json", &d1)
	var manifest struct {
		PoolSHA  string `json:"fixed_pool_file_sha256"`
		Requests map[string]struct {
			Mapping map[string]string `json:"opaque_to_source"`
		} `json:"requests"`
	}
	read("phase2-v3-semantic-manifest.json", &manifest)
	if manifest.PoolSHA != inputs["phase2-v3-fixed-pools.json"] {
		t.Fatal("pool changed")
	}
	var run struct {
		Attempts map[string]struct {
			Raw string `json:"raw_file"`
			SHA string `json:"raw_sha256"`
		} `json:"attempts"`
	}
	read("phase2-v3-agent-run.json", &run)
	rows := []map[string]any{}
	for _, name := range []string{"baseline-cases.json", "overlap-controls-v1.json"} {
		var set retrievalReplaySet
		read(name, &set)
		for _, c := range set.Cases {
			sources := map[string]bool{}
			for _, member := range frozen.Pools[c.ID].Members {
				sources[frozen.Cards[member.ID].Result.SourcePath] = true
			}
			response := retrievalAgentResponse{Ranking: []string{}}
			for source := range sources {
				response.Ranking = append(response.Ranking, source)
			}
			available := false
			if attempt, ok := run.Attempts[c.ID]; ok {
				data, err := os.ReadFile(filepath.Join(evidence, attempt.Raw))
				if err != nil {
					t.Fatal(err)
				}
				sha, err := retrievalFileSHA(filepath.Join(evidence, attempt.Raw))
				if err != nil || sha != attempt.SHA {
					t.Fatal("raw changed", c.ID)
				}
				inputs[attempt.Raw] = sha
				ids := []string{}
				for id, source := range manifest.Requests[c.ID].Mapping {
					if !sources[source] {
						t.Fatal("source outside pool")
					}
					ids = append(ids, id)
				}
				if len(ids) != len(sources) {
					t.Fatal("mapping cardinality")
				}
				response, err = retrievalValidateAgentRanking(data, ids)
				if err != nil {
					t.Fatal(err)
				}
				available = true
			}
			coverage := retrievalCoverageFor(response)
			results := []QueryResult{}
			if available {
				found := false
				for _, o := range d1.Observations {
					if o.ID == c.ID && o.Packer == "P0" {
						results = o.Results
						found = true
					}
				}
				if !found || len(results) != coverage.Direct {
					t.Fatal("P0/direct mismatch", c.ID)
				}
			}
			envelope := map[string]any{"coverage": coverage, "results": results, "packer": "P0", "verification": "pending workflow verification", "knowledge_absence": "not established"}
			trace := "not available: no semantic assessment captured"
			if available {
				trace = "phase2-v3-disposition-results.json#case=" + c.ID
			}
			rows = append(rows, map[string]any{"id": c.ID, "semantic_response_available": available, "coverage": coverage, "default_packer": "P0", "preview": envelope, "results_array_runes": retrievalV3PayloadRunes(results), "consumer_envelope_runes": retrievalV3PayloadRunes(envelope), "envelope_within_2000_runes": retrievalV3PayloadRunes(envelope) <= 2000, "decision_trace": trace, "information_sufficiency": "not evaluated", "adoption": "not evaluated"})
		}
	}
	sha, err := retrievalFileSHA(filepath.Join(root, "scripts", "ai-skill-cli", "internal", "app", "retrieval_readiness_test.go"))
	if err != nil {
		t.Fatal(err)
	}
	revision, err := retrievalGitRevision(root)
	if err != nil {
		t.Fatal(err)
	}
	retrievalV3Write(t, output, map[string]any{"kind": "offline coverage prototype; no workflow consumer integration", "scorer_revision": revision, "scorer_sha256": sha, "inputs_sha256": inputs, "cases": rows, "phase2": "incomplete", "verified_context": "not evaluated"})
}
