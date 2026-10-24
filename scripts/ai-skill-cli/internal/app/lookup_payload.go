package app

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

func lookupPayloadRunes(v any) int { b, _ := json.Marshal(v); return utf8.RuneCount(b) }

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

type retrievalBlindModelResult struct {
	Source  string `json:"source"`
	Kind    string `json:"kind"`
	Status  string `json:"status"`
	Summary string `json:"summary"`
	Need    string `json:"need_to_read"`
}
type retrievalBlindPayload struct {
	Guard      string                      `json:"guard"`
	State      string                      `json:"state"`
	Pool       int                         `json:"pool"`
	Unassessed int                         `json:"unassessed"`
	Unknown    int                         `json:"unknown"`
	Deferred   int                         `json:"deferred"`
	Trace      string                      `json:"trace"`
	Results    []retrievalBlindModelResult `json:"results"`
}

func retrievalBlindPack(caseID string, coverage retrievalCoverage, ranked []retrievalBlindModelResult, budget int) (retrievalBlindPayload, []string, error) {
	return lookupPack("blind-selector/"+caseID, coverage, ranked, budget)
}

func lookupPack(trace string, coverage retrievalCoverage, ranked []retrievalBlindModelResult, budget int) (retrievalBlindPayload, []string, error) {
	p := retrievalBlindPayload{Guard: "Candidates only; read sources and verify applicability; empty does not establish absence.", State: coverage.State, Pool: coverage.Candidates, Unassessed: coverage.Unassessed, Unknown: coverage.Unknown, Trace: trace, Results: append([]retrievalBlindModelResult{}, ranked...)}
	keep := len(ranked)
	for {
		p.Deferred = len(ranked) - keep
		p.Results = p.Results[:keep]
		if lookupPayloadRunes(p) <= budget {
			break
		}
		if keep == 0 {
			return p, nil, fmt.Errorf("required state over budget")
		}
		keep--
	}
	deferred := []string{}
	for _, r := range ranked[keep:] {
		deferred = append(deferred, r.Source)
	}
	return p, deferred, nil
}
