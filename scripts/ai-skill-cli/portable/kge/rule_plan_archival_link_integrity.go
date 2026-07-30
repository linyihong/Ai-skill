package kge

import (
	"fmt"
	"path"
	"sort"
	"strings"
)

const (
	planArchivalLinkOptOut   = "[skip-plan-archival-link-integrity]"
	archivalProvenanceMarker = "<!-- archival-provenance -->"
)

// PlanArchivalLinkIntegrityRule detects broken markdown links and stale
// textual references when plans move between plans/active/ and
// plans/archived/.
//
// Adapter supplies CapPathRenames (git rename detection), CapStagedContent
// (staged-first file bodies), and CapRepoFS ExistingPaths (existence map).
type PlanArchivalLinkIntegrityRule struct{}

func (PlanArchivalLinkIntegrityRule) ID() string { return "rule.plan_archival_link_integrity" }
func (PlanArchivalLinkIntegrityRule) Kind() Kind { return KindValidation }

func (PlanArchivalLinkIntegrityRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapCommitMsg, CapPathRenames, CapStagedContent, CapRepoFS}
}

func (PlanArchivalLinkIntegrityRule) Validate(ctx Context) []Finding {
	if hasOptOut(ctx.CommitMsg, planArchivalLinkOptOut) {
		return nil
	}
	renames := filterPlanArchiveRenames(ctx.PathRenames)
	if len(renames) == 0 {
		return nil
	}

	renameMap := make(map[string]string, len(renames))
	movedNew := make(map[string]bool, len(renames))
	for _, r := range renames {
		renameMap[r.OldPath] = r.NewPath
		movedNew[r.NewPath] = true
	}

	var findings []Finding
	paths := make([]string, 0, len(ctx.FileContents))
	for mdPath := range ctx.FileContents {
		paths = append(paths, mdPath)
	}
	sort.Strings(paths)
	for _, mdPath := range paths {
		content := ctx.FileContents[mdPath]
		isMoved := movedNew[mdPath]
		kind := "broken_inbound_link"
		if isMoved {
			kind = "broken_outbound_link"
		}
		for _, link := range ExtractMarkdownLinks([]byte(content)) {
			f, ok := classifyArchivalLink(mdPath, link, renameMap, ctx.ExistingPaths, kind)
			if !ok {
				continue
			}
			if !isMoved && f.SuggestedReplacement == "" {
				continue
			}
			findings = append(findings, toArchivalFinding(f))
		}
		findings = append(findings, scanBareTextualReferences(mdPath, content, renames)...)
	}
	return findings
}

// ParsePlanArchivalRenames parses `git diff --name-status` rename lines.
// Each line is the tab-separated form: R100\told\tnew
func ParsePlanArchivalRenames(diffOut string) []PathRenameMeta {
	var renames []PathRenameMeta
	for _, line := range strings.Split(diffOut, "\n") {
		if !strings.HasPrefix(line, "R") {
			continue
		}
		fields := strings.Split(line, "\t")
		if len(fields) < 3 {
			continue
		}
		oldPath := strings.TrimSpace(fields[1])
		newPath := strings.TrimSpace(fields[2])
		if IsPlanArchiveMove(oldPath, newPath) {
			renames = append(renames, PathRenameMeta{OldPath: oldPath, NewPath: newPath})
		}
	}
	return renames
}

// IsPlanArchiveMove reports active ↔ archived plan path moves.
func IsPlanArchiveMove(oldPath, newPath string) bool {
	active := "plans/active/"
	archived := "plans/archived/"
	return (strings.HasPrefix(oldPath, active) && strings.HasPrefix(newPath, archived)) ||
		(strings.HasPrefix(oldPath, archived) && strings.HasPrefix(newPath, active))
}

func filterPlanArchiveRenames(in []PathRenameMeta) []PathRenameMeta {
	var out []PathRenameMeta
	for _, r := range in {
		if IsPlanArchiveMove(r.OldPath, r.NewPath) {
			out = append(out, r)
		}
	}
	return out
}

type archivalHit struct {
	Severity             Severity
	Category             string
	File                 string
	Line                 int
	Column               int
	Target               string
	SuggestedReplacement string
}

func toArchivalFinding(h archivalHit) Finding {
	return Finding{
		RuleID:   "rule.plan_archival_link_integrity",
		Severity: h.Severity,
		Code:     h.Category,
		Path:     h.File,
		Message:  formatArchivalFindingLine(h),
	}
}

func formatArchivalFindingLine(f archivalHit) string {
	loc := fmt.Sprintf("%s:%d:%d", f.File, f.Line, f.Column)
	msg := fmt.Sprintf("%s [%s] target=%q", loc, f.Category, f.Target)
	if f.SuggestedReplacement != "" {
		msg += fmt.Sprintf("  → suggested: %q", f.SuggestedReplacement)
	}
	return msg
}

func scanBareTextualReferences(mdPath, content string, renames []PathRenameMeta) []Finding {
	hit := false
	for _, r := range renames {
		if strings.Contains(content, r.OldPath) {
			hit = true
			break
		}
	}
	if !hit {
		return nil
	}
	lines := strings.Split(content, "\n")
	var findings []Finding
	for lineIdx, line := range lines {
		for _, r := range renames {
			old := r.OldPath
			start := 0
			for {
				pos := strings.Index(line[start:], old)
				if pos < 0 {
					break
				}
				abs := start + pos
				start = abs + len(old)
				if isLinkTargetContext(line, abs) {
					continue
				}
				provenance := strings.Contains(line, archivalProvenanceMarker) ||
					(lineIdx > 0 && strings.Contains(lines[lineIdx-1], archivalProvenanceMarker))
				sev := SeverityWarning
				category := "stale_textual_reference"
				if provenance {
					sev = SeverityInfo
					category = "historical_provenance_reference"
				}
				findings = append(findings, toArchivalFinding(archivalHit{
					Severity:             sev,
					Category:             category,
					File:                 mdPath,
					Line:                 lineIdx + 1,
					Column:               abs + 1,
					Target:               old,
					SuggestedReplacement: r.NewPath,
				}))
			}
		}
	}
	return findings
}

func isLinkTargetContext(line string, pos int) bool {
	return IsLinkTargetContext(line, pos)
}

// IsLinkTargetContext reports whether pos starts a markdown link target.
func IsLinkTargetContext(line string, pos int) bool {
	i := pos - 1
	for i >= 0 && (line[i] == ' ' || line[i] == '\t') {
		i--
	}
	if i < 0 || line[i] != '(' {
		return false
	}
	i--
	if i < 0 || line[i] != ']' {
		return false
	}
	return true
}

func classifyArchivalLink(fromFile string, link Link, renameMap map[string]string, existing map[string]bool, kind string) (archivalHit, bool) {
	target := stripLinkFragment(link.Target)
	if target == "" {
		return archivalHit{}, false
	}
	resolved := ResolveRepoPath(fromFile, target)
	if resolved == "" || resolved == ".." || strings.HasPrefix(resolved, "../") {
		return archivalHit{}, false
	}
	if newPath, ok := renameMap[resolved]; ok {
		return archivalHit{
			Severity:             SeverityError,
			Category:             kind,
			File:                 fromFile,
			Line:                 link.Line,
			Column:               link.Column,
			Target:               link.Target,
			SuggestedReplacement: SuggestArchivalReplacement(fromFile, newPath, link.Target),
		}, true
	}
	if existing != nil && existing[resolved] {
		return archivalHit{}, false
	}
	return archivalHit{
		Severity: SeverityError,
		Category: kind,
		File:     fromFile,
		Line:     link.Line,
		Column:   link.Column,
		Target:   link.Target,
	}, true
}

func stripLinkFragment(target string) string {
	return StripLinkFragment(target)
}

// StripLinkFragment removes a trailing #fragment from a link target.
func StripLinkFragment(target string) string {
	if idx := strings.Index(target, "#"); idx >= 0 {
		return target[:idx]
	}
	return target
}

// ResolveRepoPath joins a markdown link target against the source file's
// directory and cleans the result as a POSIX-style repo path.
func ResolveRepoPath(fromFile, linkTarget string) string {
	if linkTarget == "" {
		return ""
	}
	dir := path.Dir(fromFile)
	return path.Clean(path.Join(dir, linkTarget))
}

// SuggestArchivalReplacement computes a rewrite hint for a broken link.
func SuggestArchivalReplacement(fromFile, newPath, originalTarget string) string {
	if (strings.HasPrefix(originalTarget, "plans/active/") || strings.HasPrefix(originalTarget, "plans/archived/")) &&
		!strings.HasPrefix(originalTarget, "./") &&
		!strings.HasPrefix(originalTarget, "../") {
		return newPath
	}
	return PosixRel(path.Dir(fromFile), newPath)
}

// PosixRel is a POSIX-path equivalent of filepath.Rel.
func PosixRel(fromDir, toPath string) string {
	if fromDir == "." || fromDir == "" {
		return toPath
	}
	fromParts := strings.Split(fromDir, "/")
	toParts := strings.Split(toPath, "/")
	common := 0
	for common < len(fromParts) && common < len(toParts) && fromParts[common] == toParts[common] {
		common++
	}
	var parts []string
	for i := common; i < len(fromParts); i++ {
		parts = append(parts, "..")
	}
	parts = append(parts, toParts[common:]...)
	if len(parts) == 0 {
		return "."
	}
	return strings.Join(parts, "/")
}

// ScanBareTextualReferencesForTest exposes bare-text scanning for unit tests.
func ScanBareTextualReferencesForTest(mdPath, content string, renames []PathRenameMeta) []Finding {
	return scanBareTextualReferences(mdPath, content, renames)
}
