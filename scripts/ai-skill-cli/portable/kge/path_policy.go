package kge

import (
	"path/filepath"
	"strings"
)

// RequiresDeepStrictCognitiveMode reports whether a staged path forces
// DEEP + STRICT (or higher) per cognitive-modes-phase-integration floors.
func RequiresDeepStrictCognitiveMode(path string) bool {
	path = filepath.ToSlash(path)
	return strings.HasPrefix(path, "runtime/") ||
		strings.HasPrefix(path, "scripts/ai-skill-cli/") ||
		strings.HasPrefix(path, "governance/") ||
		strings.HasPrefix(path, "enforcement/") ||
		strings.HasPrefix(path, "validation/") ||
		path == "knowledge/runtime/routing-registry.yaml" ||
		(strings.HasPrefix(path, "workflow/") && strings.HasSuffix(path, ".yaml")) ||
		strings.HasPrefix(path, "plans/active/")
}
