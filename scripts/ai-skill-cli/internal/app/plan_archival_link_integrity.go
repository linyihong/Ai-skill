package app

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
)

// planRename is one staged archive event detected via
// `git diff --cached --find-renames=90% --name-status`.
type planRename struct {
	OldPath string
	NewPath string
}

// linkFinding is one broken or stale reference discovered during a plan
// archive event. Severity / Category drive output formatting; the
// SuggestedReplacement carries the rewrite hint when known.
type linkFinding struct {
	Severity             string
	Category             string
	File                 string
	Line                 int
	Column               int
	Target               string
	SuggestedReplacement string
}

// validatePlanArchivalLinkIntegrity returns BLOCK-severity findings only.
// Portable scan lives in kge.PlanArchivalLinkIntegrityRule; this package
// supplies git rename detection, staged-first content, and legacy banners.
func validatePlanArchivalLinkIntegrity(text string, staged []string, root string) string {
	return formatArchivalKGEFindings(runKGEPlanArchivalLinkIntegrity(text, root), kge.SeverityError)
}

// warnPlanArchivalLinkIntegrity returns WARNING-severity findings (advisory).
func warnPlanArchivalLinkIntegrity(text string, staged []string, root string) string {
	return formatArchivalKGEFindings(runKGEPlanArchivalLinkIntegrity(text, root), kge.SeverityWarning)
}

func runKGEPlanArchivalLinkIntegrity(text, root string) []kge.Finding {
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == "[skip-plan-archival-link-integrity]" {
			return nil
		}
	}
	renames := detectPlanArchivalRenames(root)
	if len(renames) == 0 {
		return nil
	}
	ctx := buildPlanArchivalLinkContext(text, root, renames)
	eng := kge.NewEngine(kge.PlanArchivalLinkIntegrityRule{})
	return eng.Run(ctx)
}

func buildPlanArchivalLinkContext(text, root string, renames []planRename) kge.Context {
	metas := make([]kge.PathRenameMeta, 0, len(renames))
	movedNew := make(map[string]bool, len(renames))
	for _, r := range renames {
		metas = append(metas, kge.PathRenameMeta{OldPath: r.OldPath, NewPath: r.NewPath})
		movedNew[r.NewPath] = true
	}

	mdFiles := listRepoMarkdown(root)
	inSet := make(map[string]bool, len(mdFiles))
	for _, m := range mdFiles {
		inSet[m] = true
	}
	for newPath := range movedNew {
		if !inSet[newPath] {
			mdFiles = append(mdFiles, newPath)
			inSet[newPath] = true
		}
	}

	contents := map[string]string{}
	existing := map[string]bool{}
	for _, p := range mdFiles {
		existing[p] = true
	}
	// Attest all tracked paths so existence checks match prior Stat-on-index behavior.
	for _, p := range listRepoAllTracked(root) {
		existing[p] = true
	}
	for _, mdPath := range mdFiles {
		data, err := readFileForScan(root, mdPath, movedNew)
		if err != nil {
			continue
		}
		contents[mdPath] = string(data)
	}

	return kge.Context{
		CommitMsg:     text,
		PathRenames:   metas,
		FileContents:  contents,
		ExistingPaths: existing,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:     true,
			kge.CapPathRenames:   true,
			kge.CapStagedContent: true,
			kge.CapRepoFS:        true,
		},
	}
}

func formatArchivalKGEFindings(findings []kge.Finding, severity kge.Severity) string {
	var sel []kge.Finding
	for _, f := range findings {
		if f.Severity == severity {
			sel = append(sel, f)
		}
	}
	if len(sel) == 0 {
		return ""
	}
	var out strings.Builder
	switch severity {
	case kge.SeverityError:
		out.WriteString("plan-archival-link-integrity: archive breaks markdown reference(s):")
	default:
		out.WriteString("plan-archival-link-integrity (advisory): stale textual reference(s):")
	}
	for _, f := range sel {
		out.WriteString("\n    - " + f.Message)
	}
	if severity == kge.SeverityError {
		out.WriteString("\n  Update each reference, or add a standalone `[skip-plan-archival-link-integrity]` trailer for emergency archives.")
	} else {
		out.WriteString("\n  Update the reference or add a `<!-- archival-provenance -->` marker. Advisory only — does NOT block the commit.")
	}
	return out.String()
}

// readFileForScan reads content for the inbound / textual scan. The
// staged blob is the canonical commit candidate (TD-1).
func readFileForScan(root, mdPath string, _ map[string]bool) ([]byte, error) {
	if data, err := readStagedFileContent(root, mdPath); err == nil {
		return data, nil
	}
	return os.ReadFile(filepath.Join(root, filepath.FromSlash(mdPath)))
}

func detectPlanArchivalRenames(root string) []planRename {
	cmd := exec.Command("git", "-C", root, "diff", "--cached", "--find-renames=90%", "--name-status")
	out, err := cmd.Output()
	if err != nil {
		return nil
	}
	return parsePlanRenames(string(out))
}

// parsePlanRenames wraps kge.ParsePlanArchivalRenames for unit tests.
func parsePlanRenames(diffOut string) []planRename {
	metas := kge.ParsePlanArchivalRenames(diffOut)
	if len(metas) == 0 {
		return nil
	}
	out := make([]planRename, 0, len(metas))
	for _, m := range metas {
		out = append(out, planRename{OldPath: m.OldPath, NewPath: m.NewPath})
	}
	return out
}

func isPlanArchiveMove(oldPath, newPath string) bool {
	return kge.IsPlanArchiveMove(oldPath, newPath)
}

func readStagedFileContent(root, repoPath string) ([]byte, error) {
	cmd := exec.Command("git", "-C", root, "show", ":"+repoPath)
	return cmd.Output()
}

func resolveRepoPath(fromFile, linkTarget string) string {
	return kge.ResolveRepoPath(fromFile, linkTarget)
}

func pathExistsInRepo(root, repoPath string) bool {
	_, err := os.Stat(filepath.Join(root, filepath.FromSlash(repoPath)))
	return err == nil
}

func listRepoMarkdown(root string) []string {
	cmd := exec.Command("git", "-C", root, "ls-files", "*.md")
	out, err := cmd.Output()
	if err != nil {
		return nil
	}
	var files []string
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line != "" {
			files = append(files, line)
		}
	}
	return files
}

func listRepoAllTracked(root string) []string {
	cmd := exec.Command("git", "-C", root, "ls-files")
	out, err := cmd.Output()
	if err != nil {
		return nil
	}
	var files []string
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line != "" {
			files = append(files, line)
		}
	}
	return files
}

func suggestReplacement(fromFile, newPath, originalTarget string) string {
	return kge.SuggestArchivalReplacement(fromFile, newPath, originalTarget)
}

func posixRel(fromDir, toPath string) string {
	return kge.PosixRel(fromDir, toPath)
}

// scanBareTextualReferences wraps the portable scanner for unit tests.
func scanBareTextualReferences(mdPath string, content []byte, renames []planRename) []linkFinding {
	metas := make([]kge.PathRenameMeta, 0, len(renames))
	for _, r := range renames {
		metas = append(metas, kge.PathRenameMeta{OldPath: r.OldPath, NewPath: r.NewPath})
	}
	findings := kge.ScanBareTextualReferencesForTest(mdPath, string(content), metas)
	if len(findings) == 0 {
		return nil
	}
	out := make([]linkFinding, 0, len(findings))
	for _, f := range findings {
		sev := "warning"
		switch f.Severity {
		case kge.SeverityInfo:
			sev = "info"
		case kge.SeverityError:
			sev = "block"
		}
		out = append(out, findingToLinkFinding(f, sev))
	}
	return out
}

func findingToLinkFinding(f kge.Finding, sev string) linkFinding {
	// Message: path:line:col [category] target="..."  → suggested: "..."
	lf := linkFinding{
		Severity: sev,
		Category: f.Code,
		File:     f.Path,
	}
	msg := f.Message
	if i := strings.Index(msg, " target="); i >= 0 {
		rest := msg[i+len(" target="):]
		if strings.HasPrefix(rest, "\"") {
			rest = rest[1:]
			if j := strings.Index(rest, "\""); j >= 0 {
				lf.Target = rest[:j]
				rest = rest[j+1:]
			}
		}
		const sug = "  → suggested: \""
		if k := strings.Index(rest, sug); k >= 0 {
			rest = rest[k+len(sug):]
			if j := strings.Index(rest, "\""); j >= 0 {
				lf.SuggestedReplacement = rest[:j]
			}
		}
	}
	// Line/Column from "file:line:col"
	if f.Path != "" && strings.HasPrefix(msg, f.Path+":") {
		tail := msg[len(f.Path)+1:]
		var line, col int
		if _, err := fmt.Sscanf(tail, "%d:%d", &line, &col); err == nil {
			lf.Line = line
			lf.Column = col
		}
	}
	return lf
}

// formatFindingsBySeverity retained for any callers that still use
// linkFinding slices (tests). Prefer formatArchivalKGEFindings for hooks.
func formatFindingsBySeverity(findings []linkFinding, severity string) string {
	var sel []linkFinding
	for _, f := range findings {
		if f.Severity == severity {
			sel = append(sel, f)
		}
	}
	if len(sel) == 0 {
		return ""
	}
	var out strings.Builder
	switch severity {
	case "block":
		out.WriteString("plan-archival-link-integrity: archive breaks markdown reference(s):")
	default:
		out.WriteString("plan-archival-link-integrity (advisory): stale textual reference(s):")
	}
	for _, f := range sel {
		out.WriteString("\n    - " + formatFindingLine(f))
	}
	if severity == "block" {
		out.WriteString("\n  Update each reference, or add a standalone `[skip-plan-archival-link-integrity]` trailer for emergency archives.")
	} else {
		out.WriteString("\n  Update the reference or add a `<!-- archival-provenance -->` marker. Advisory only — does NOT block the commit.")
	}
	return out.String()
}

func formatFindingLine(f linkFinding) string {
	loc := fmt.Sprintf("%s:%d:%d", f.File, f.Line, f.Column)
	msg := fmt.Sprintf("%s [%s] target=%q", loc, f.Category, f.Target)
	if f.SuggestedReplacement != "" {
		msg += fmt.Sprintf("  → suggested: %q", f.SuggestedReplacement)
	}
	return msg
}
