package app

// Plan evidence validators (2026-07-09; main-absorbed rule 2026-07-15).
//
// Enforces the evidence/ subdirectory convention documented in
// governance/lifecycle/plan-evidence.md:
//
//   - validatePlanEvidenceConvention     block    README + index sync for evidence/;
//                                                 evidence/ implies folder main (_plan.md);
//                                                 block flat sibling <slug>.md when evidence/ exists
//   - warnPlanEvidenceLineNumberCitations warning  discourage L123 line refs in plan folder
//
// See: plans/templates/plan-evidence/README.md

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var planEvidenceLineNumberRE = regexp.MustCompile(`\bL\d+\b`)

// isPlanEvidencePath reports whether rel is under plans/{active,archived}/<plan>/evidence/.
func isPlanEvidencePath(rel string) bool {
	rel = filepath.ToSlash(rel)
	if !strings.HasPrefix(rel, "plans/active/") && !strings.HasPrefix(rel, "plans/archived/") {
		return false
	}
	return strings.Contains(rel, "/evidence/")
}

// planDirFromEvidencePath returns plans/active/foo from plans/active/foo/evidence/bar.md.
func planDirFromEvidencePath(rel string) (string, bool) {
	rel = filepath.ToSlash(rel)
	idx := strings.Index(rel, "/evidence/")
	if idx < 0 {
		return "", false
	}
	return rel[:idx], true
}

func stagedPlanEvidencePaths(staged []string) []string {
	var out []string
	for _, s := range staged {
		s = filepath.ToSlash(s)
		if isPlanEvidencePath(s) && strings.HasSuffix(strings.ToLower(s), ".md") {
			out = append(out, s)
		}
	}
	return out
}

func readFileString(root, rel string) (string, bool) {
	abs := rel
	if root != "" {
		abs = filepath.Join(root, rel)
	}
	data, err := os.ReadFile(abs)
	if err != nil {
		return "", false
	}
	return string(data), true
}

// evidenceFilesInDir lists .md files directly under planDir/evidence/ (non-recursive).
func evidenceFilesInDir(root, planDir string) ([]string, error) {
	dir := filepath.Join(root, planDir, "evidence")
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.EqualFold(filepath.Ext(name), ".md") {
			files = append(files, name)
		}
	}
	return files, nil
}

func readmeHasRequiredSections(body string) (missing []string) {
	if !strings.Contains(body, "引用規則") && !strings.Contains(body, "Citation rule") {
		missing = append(missing, "引用規則 section")
	}
	if !strings.Contains(body, "Run 索引") && !strings.Contains(body, "## Index") {
		missing = append(missing, "Run 索引 table")
	}
	return missing
}

func readmeReferencesFile(readmeBody, filename string) bool {
	return strings.Contains(readmeBody, filename)
}

// collectPlanDirsForEvidenceConvention returns plan folders that must satisfy
// the evidence convention for this commit:
//   - any staged path under .../evidence/
//   - any staged top-level plans/{active|archived}/<slug>.md while <slug>/evidence/ exists
func collectPlanDirsForEvidenceConvention(staged []string, root string) map[string]bool {
	planDirs := map[string]bool{}
	for _, p := range stagedPlanEvidencePaths(staged) {
		if dir, ok := planDirFromEvidencePath(p); ok {
			planDirs[dir] = true
		}
	}
	for _, s := range staged {
		s = filepath.ToSlash(s)
		if !strings.HasPrefix(s, "plans/active/") && !strings.HasPrefix(s, "plans/archived/") {
			continue
		}
		if !strings.HasSuffix(strings.ToLower(s), ".md") {
			continue
		}
		parts := strings.Split(s, "/")
		// plans/active/<slug>.md only (top-level flat main)
		if len(parts) != 3 {
			continue
		}
		planDir := strings.TrimSuffix(s, ".md")
		evDir := filepath.Join(root, planDir, "evidence")
		if st, err := os.Stat(evDir); err == nil && st.IsDir() {
			planDirs[planDir] = true
		}
	}
	return planDirs
}

// validatePlanEvidenceConvention enforces evidence/README.md presence, index
// coverage, and main-plan absorption (_plan.md inside the folder; no flat
// sibling <slug>.md) when evidence/ is in play for the commit.
func validatePlanEvidenceConvention(text string, staged []string, root string) string {
	if hasOptOutTrailer(text, "[skip-plan-evidence]") {
		return ""
	}
	planDirs := collectPlanDirsForEvidenceConvention(staged, root)
	if len(planDirs) == 0 {
		return ""
	}

	var violations []string
	for planDir := range planDirs {
		mainRel := planDir + "/_plan.md"
		if _, ok := readFileString(root, mainRel); !ok {
			violations = append(violations, fmt.Sprintf(
				"%s: missing _plan.md — when evidence/ is used, the main plan must live in the folder as _plan.md (not a sibling <slug>.md)",
				planDir))
		}
		flatSibling := planDir + ".md"
		if st, err := os.Stat(filepath.Join(root, flatSibling)); err == nil && !st.IsDir() {
			violations = append(violations, fmt.Sprintf(
				"%s: flat sibling %s still exists — move it to %s/_plan.md before using evidence/ (ai-skill plans folderize or git mv)",
				planDir, flatSibling, planDir))
		}

		readmeRel := planDir + "/evidence/README.md"
		readmeBody, readmeOK := readFileString(root, readmeRel)
		if !readmeOK {
			violations = append(violations, fmt.Sprintf("%s: missing evidence/README.md (required when using evidence/)", planDir))
			continue
		}
		if miss := readmeHasRequiredSections(readmeBody); len(miss) > 0 {
			violations = append(violations, fmt.Sprintf("%s: evidence/README.md missing: %s", planDir, strings.Join(miss, ", ")))
		}
		files, err := evidenceFilesInDir(root, planDir)
		if err != nil {
			violations = append(violations, fmt.Sprintf("%s: cannot read evidence/: %v", planDir, err))
			continue
		}
		for _, name := range files {
			if strings.EqualFold(name, "README.md") {
				continue
			}
			if !readmeReferencesFile(readmeBody, name) {
				violations = append(violations, fmt.Sprintf("%s: evidence/%s not listed in evidence/README.md Run 索引", planDir, name))
			}
		}
	}

	if len(violations) == 0 {
		return ""
	}
	return "plan-evidence-convention: evidence/ directory rule violation(s):\n    - " +
		strings.Join(violations, "\n    - ") +
		"\n  See governance/lifecycle/plan-evidence.md and plans/templates/plan-evidence/README.md." +
		"\n  Opt-out (emergency only): standalone `[skip-plan-evidence]` trailer."
}

// warnPlanEvidenceLineNumberCitations returns a non-blocking warning when staged
// plan-folder markdown uses line-number citations (L123) instead of file paths.
func warnPlanEvidenceLineNumberCitations(_ string, staged []string, root string) string {
	planDirs := map[string]bool{}
	for _, rel := range staged {
		rel = filepath.ToSlash(rel)
		if !strings.HasPrefix(rel, "plans/active/") && !strings.HasPrefix(rel, "plans/archived/") {
			continue
		}
		if !strings.HasSuffix(strings.ToLower(rel), ".md") {
			continue
		}
		if strings.Contains(rel, "/evidence/") {
			if dir, ok := planDirFromEvidencePath(rel); ok {
				planDirs[dir] = true
			}
			continue
		}
		if dir, ok := planDirFromEvidencePath(rel); ok {
			planDirs[dir] = true
		} else {
			// plans/active/foo/_plan.md or plans/active/foo/01-kit.md
			parts := strings.Split(rel, "/")
			if len(parts) >= 3 {
				planDirs[strings.Join(parts[:3], "/")] = true
			}
		}
	}
	// Also warn when staging evidence paths touch a plan dir
	for _, rel := range stagedPlanEvidencePaths(staged) {
		if dir, ok := planDirFromEvidencePath(rel); ok {
			planDirs[dir] = true
		}
	}

	var warnings []string
	for planDir := range planDirs {
		evDir := filepath.Join(root, planDir, "evidence")
		if _, err := os.Stat(evDir); err != nil {
			continue
		}
		for _, rel := range staged {
			rel = filepath.ToSlash(rel)
			if !strings.HasPrefix(rel, planDir+"/") || !strings.HasSuffix(strings.ToLower(rel), ".md") {
				continue
			}
			body, ok := readFileString(root, rel)
			if !ok {
				continue
			}
			if planEvidenceLineNumberRE.MatchString(body) {
				warnings = append(warnings, fmt.Sprintf("%s: contains line-number citation (e.g. L123); prefer evidence/<file>.md paths per governance/lifecycle/plan-evidence.md", rel))
			}
		}
	}
	if len(warnings) == 0 {
		return ""
	}
	return "plan-evidence-line-citation (warning): prefer file-path citations over line numbers:\n    - " +
		strings.Join(warnings, "\n    - ")
}

// isUnderPlanEvidenceDir reports whether rel is a file inside .../evidence/ (not README naming check).
func isUnderPlanEvidenceDir(rel string) bool {
	return isPlanEvidencePath(rel)
}
