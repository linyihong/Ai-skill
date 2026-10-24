package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type retrievalBlindAssessment struct {
	ID       string `json:"id"`
	Relation string `json:"relation"`
	Reason   string `json:"reason"`
	Need     string `json:"need_to_read"`
}
type retrievalBlindResponse struct {
	Selected      []string                   `json:"selected"`
	Assessments   []retrievalBlindAssessment `json:"assessments"`
	UnknownAction string                     `json:"unknown_action"`
}

func retrievalParseBlind(raw []byte, ids []string) (retrievalBlindResponse, error) {
	var r retrievalBlindResponse
	d := json.NewDecoder(bytes.NewReader(raw))
	d.DisallowUnknownFields()
	if err := d.Decode(&r); err != nil {
		return r, err
	}
	var extra any
	if d.Decode(&extra) != io.EOF {
		return r, fmt.Errorf("trailing output")
	}
	if r.Selected == nil || r.Assessments == nil || len(r.Assessments) > 8 || strings.TrimSpace(r.UnknownAction) == "" {
		return r, fmt.Errorf("missing fields or too many assessments")
	}
	allowed := map[string]bool{}
	for _, id := range ids {
		if id == "" || allowed[id] {
			return r, fmt.Errorf("invalid request IDs")
		}
		allowed[id] = true
	}
	seen := map[string]bool{}
	direct := map[string]bool{}
	for _, a := range r.Assessments {
		if !allowed[a.ID] || seen[a.ID] || strings.TrimSpace(a.Reason) == "" || strings.TrimSpace(a.Need) == "" {
			return r, fmt.Errorf("invalid assessment")
		}
		seen[a.ID] = true
		switch a.Relation {
		case "direct":
			direct[a.ID] = true
		case "background", "not-applicable", "unknown":
		default:
			return r, fmt.Errorf("unknown relation")
		}
	}
	selected := map[string]bool{}
	for _, id := range r.Selected {
		if !direct[id] || selected[id] {
			return r, fmt.Errorf("foreign/duplicate/non-direct selection")
		}
		selected[id] = true
	}
	if len(direct) != len(selected) {
		return r, fmt.Errorf("direct missing from selection")
	}
	return r, nil
}
