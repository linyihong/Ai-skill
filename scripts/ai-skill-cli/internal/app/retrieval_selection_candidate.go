package app

import (
	"encoding/json"
	"unicode/utf8"
)

// Experimental v2 bounds do not assert relevance or applicability.
const runtimeRetrievalPreviewRunes = 2000
const runtimeRetrievalCandidateLimit = 64

func runtimeIndexSelectionCandidates(dbPath, keyword string, opts runtimeOptions) ([]QueryResult, error) {
	opts.limit = runtimeRetrievalCandidateLimit
	candidates, err := runRuntimeIndexMatches(dbPath, keyword, runtimeFTSMatchLiteral(keyword), opts)
	stage := "exact phrase; applicability unverified"
	if err != nil {
		return nil, err
	}
	if len(candidates) == 0 {
		candidates, err = runRuntimeIndexMatches(dbPath, keyword, runtimeFTSLexicalQuery(keyword), opts)
		stage = "lexical fallback; applicability unverified"
	}
	for i := range candidates {
		candidates[i].MatchReason = stage
	}
	return candidates, err
}

func runRuntimeIndexQuery(dbPath, keyword string, opts runtimeOptions) ([]QueryResult, error) {
	candidates, err := runtimeIndexSelectionCandidates(dbPath, keyword, opts)
	if err != nil {
		return nil, err
	}
	return runtimeBoundRetrievalResults(candidates, opts.limit, runtimeRetrievalPreviewRunes), nil
}

// Measure the same compact JSON results replay scores. Never cut JSON bytes
// or identity fields. Summary is shortened on rune boundaries with ellipsis.
func runtimeBoundRetrievalResults(candidates []QueryResult, limit, budget int) []QueryResult {
	output := []QueryResult{}
	if limit <= 0 || budget < 2 {
		return output
	}
	seen := map[string]bool{}
	for _, candidate := range candidates {
		if seen[candidate.SourcePath] {
			continue
		}
		seen[candidate.SourcePath] = true
		if runtimeRetrievalPayloadRunes([]QueryResult{candidate}) > budget {
			summary := []rune(candidate.Summary)
			candidate.Summary = ""
			if runtimeRetrievalPayloadRunes([]QueryResult{candidate}) > budget {
				continue
			}
			low, high := 0, len(summary)
			for low < high {
				middle := (low + high + 1) / 2
				candidate.Summary = string(summary[:middle]) + "…"
				if runtimeRetrievalPayloadRunes([]QueryResult{candidate}) <= budget {
					low = middle
				} else {
					high = middle - 1
				}
			}
			if low > 0 {
				candidate.Summary = string(summary[:low]) + "…"
			} else {
				candidate.Summary = ""
			}
		}
		output = append(output, candidate)
		// Conservative top-one shortlist, fixed before seeing replay outcomes.
		break
	}
	return output
}
func runtimeRetrievalPayloadRunes(results []QueryResult) int {
	payload, _ := json.Marshal(results)
	return utf8.RuneCount(payload)
}
