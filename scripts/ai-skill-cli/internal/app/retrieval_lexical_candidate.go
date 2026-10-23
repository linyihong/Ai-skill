package app

import (
	"regexp"
	"strings"
)

// runtimeFTSLexicalQuery treats user input as data, never as FTS syntax.
// This experimental candidate uses Unicode word runs without synonyms,
// language-specific segmentation, or benchmark-specific stopword lists.
func runtimeFTSLexicalQuery(keyword string) string {
	words := regexp.MustCompile(`[\p{L}\p{N}_]+`).FindAllString(keyword, -1)
	terms := []string{}
	seen := map[string]bool{}
	for _, word := range words {
		key := strings.ToLower(word)
		if seen[key] {
			continue
		}
		seen[key] = true
		terms = append(terms, runtimeFTSMatchLiteral(word))
		if len(terms) == 64 {
			break
		}
	}
	if len(terms) == 0 {
		return runtimeFTSMatchLiteral(keyword)
	}
	return strings.Join(terms, " OR ")
}
