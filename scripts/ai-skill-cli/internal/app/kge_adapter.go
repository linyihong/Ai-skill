package app

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
)

// kgeFindingsMessage flattens engine findings into the legacy commit-msg
// validator string shape (empty = pass).
func kgeFindingsMessage(findings []kge.Finding) string {
	if len(findings) == 0 {
		return ""
	}
	parts := make([]string, 0, len(findings))
	for _, f := range findings {
		if f.Message != "" {
			parts = append(parts, f.Message)
		}
	}
	return strings.Join(parts, "\n  - ")
}

// runKGECognitiveCost adapts commit-msg modes into a portable KGE Context
// and runs Rule A (cognitive_cost). Git-free.
func runKGECognitiveCost(modes map[string]string) string {
	ctx := kge.Context{
		Modes: modes,
		Provided: map[kge.CapabilityID]bool{
			kge.CapModes: true,
		},
	}
	eng := kge.NewEngine(kge.CognitiveCostRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// stagedDiffCached runs `git diff --cached` for paths. On failure returns
// empty diff (legacy validateCLIDocSync treated git errors as pass / no match).
func stagedDiffCached(root string, pathspecs ...string) string {
	args := append([]string{"-C", root, "diff", "--cached", "--"}, pathspecs...)
	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

// runKGECLIDocSync adapts commit-msg + staged paths into KGE Context.
// Git stays in this adapter; the portable rule never execs git.
func runKGECLIDocSync(text string, staged []string, root string) string {
	diff := stagedDiffCached(root, "scripts/ai-skill-cli/internal/app/")
	ctx := kge.Context{
		RepoRoot:    root,
		CommitMsg:   text,
		StagedPaths: staged,
		StagedDiff:  diff,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapStagedPaths: true,
			kge.CapStagedDiff:  true,
		},
	}
	eng := kge.NewEngine(kge.CLIDocSyncRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEMarkdownYamlSync adapts CapRepoFS (sibling exists) for the portable rule.
func runKGEMarkdownYamlSync(text string, staged []string, root string) string {
	existing := map[string]bool{}
	paths := make([]string, 0, len(staged))
	for _, s := range staged {
		p := filepath.ToSlash(s)
		paths = append(paths, p)
		if strings.HasSuffix(p, ".md") {
			sibling := strings.TrimSuffix(p, ".md") + ".yaml"
			if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(sibling))); err == nil {
				existing[sibling] = true
			}
		}
	}
	ctx := kge.Context{
		CommitMsg:     text,
		StagedPaths:   paths,
		ExistingPaths: existing,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapStagedPaths: true,
			kge.CapRepoFS:      true,
		},
	}
	eng := kge.NewEngine(kge.MarkdownYamlSyncRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEBootstrapEntryThinness loads staged entry-file contents for the portable rule.
func runKGEBootstrapEntryThinness(text string, staged []string, root string) string {
	entrySet := make(map[string]bool, len(kge.BootstrapEntryPaths))
	for _, p := range kge.BootstrapEntryPaths {
		entrySet[p] = true
	}
	paths := make([]string, 0, len(staged))
	contents := map[string]string{}
	hasEntry := false
	for _, s := range staged {
		p := filepath.ToSlash(s)
		paths = append(paths, p)
		if !entrySet[p] {
			continue
		}
		hasEntry = true
		body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(p)))
		if err != nil {
			continue
		}
		contents[p] = string(body)
	}
	if !hasEntry {
		return ""
	}
	ctx := kge.Context{
		CommitMsg:    text,
		StagedPaths:  paths,
		FileContents: contents,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:     true,
			kge.CapStagedPaths:   true,
			kge.CapStagedContent: true,
		},
	}
	eng := kge.NewEngine(kge.BootstrapEntryThinnessRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEGlossaryRetroOwn adapts staged paths for glossary co-change.
func runKGEGlossaryRetroOwn(text string, staged []string) string {
	paths := make([]string, 0, len(staged))
	for _, s := range staged {
		paths = append(paths, filepath.ToSlash(s))
	}
	ctx := kge.Context{
		CommitMsg:   text,
		StagedPaths: paths,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapStagedPaths: true,
		},
	}
	eng := kge.NewEngine(kge.GlossaryRetroOwnRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGERuntimeYamlProjects loads staged runtime/*.yaml contents for the portable rule.
func runKGERuntimeYamlProjects(text string, staged []string, root string) string {
	paths := make([]string, 0, len(staged))
	contents := map[string]string{}
	hasRuntimeYAML := false
	for _, s := range staged {
		p := filepath.ToSlash(s)
		paths = append(paths, p)
		if !strings.HasPrefix(p, "runtime/") || !strings.HasSuffix(p, ".yaml") {
			continue
		}
		hasRuntimeYAML = true
		body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(p)))
		if err != nil {
			continue
		}
		contents[p] = string(body)
	}
	if !hasRuntimeYAML {
		return ""
	}
	ctx := kge.Context{
		CommitMsg:    text,
		StagedPaths:  paths,
		FileContents: contents,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:     true,
			kge.CapStagedPaths:   true,
			kge.CapStagedContent: true,
		},
	}
	eng := kge.NewEngine(kge.RuntimeYamlProjectsRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGETokenBudget adapts modes + commit message for the portable rule.
func runKGETokenBudget(modes map[string]string, text string) string {
	ctx := kge.Context{
		CommitMsg: text,
		Modes:     modes,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg: true,
			kge.CapModes:     true,
		},
	}
	eng := kge.NewEngine(kge.TokenBudgetRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

func normalizeStagedPaths(staged []string) []string {
	paths := make([]string, 0, len(staged))
	for _, s := range staged {
		paths = append(paths, filepath.ToSlash(s))
	}
	return paths
}

// runKGEExecutionModeFloors adapts modes + staged paths for the portable rule.
func runKGEExecutionModeFloors(modes map[string]string, staged []string) string {
	ctx := kge.Context{
		Modes:       modes,
		StagedPaths: normalizeStagedPaths(staged),
		Provided: map[kge.CapabilityID]bool{
			kge.CapModes:       true,
			kge.CapStagedPaths: true,
		},
	}
	eng := kge.NewEngine(kge.ExecutionModeFloorsRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEGovernanceModeConsistency adapts modes + staged + commit msg.
func runKGEGovernanceModeConsistency(modes map[string]string, staged []string, text string) string {
	ctx := kge.Context{
		CommitMsg:   text,
		Modes:       modes,
		StagedPaths: normalizeStagedPaths(staged),
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapModes:       true,
			kge.CapStagedPaths: true,
		},
	}
	eng := kge.NewEngine(kge.GovernanceModeConsistencyRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEMemoryModeSubdir adapts modes + staged paths for the portable rule.
func runKGEMemoryModeSubdir(modes map[string]string, staged []string) string {
	ctx := kge.Context{
		Modes:       modes,
		StagedPaths: normalizeStagedPaths(staged),
		Provided: map[kge.CapabilityID]bool{
			kge.CapModes:       true,
			kge.CapStagedPaths: true,
		},
	}
	eng := kge.NewEngine(kge.MemoryModeSubdirRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEActivationSignals loads known discovery signals into Context.
func runKGEActivationSignals(c commitMsgCtx) string {
	known := readKnownCognitiveSignals(c.root)
	ctx := kge.Context{
		CommitMsg:    c.text,
		Modes:        c.modes,
		KnownSignals: known,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:    true,
			kge.CapModes:        true,
			kge.CapKnownSignals: true,
		},
	}
	eng := kge.NewEngine(kge.ActivationSignalsRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGECapabilitySnippet adapts modes + commit message for the portable rule.
func runKGECapabilitySnippet(modes map[string]string, text string) string {
	ctx := kge.Context{
		CommitMsg: text,
		Modes:     modes,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg: true,
			kge.CapModes:     true,
		},
	}
	eng := kge.NewEngine(kge.CapabilitySnippetRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEAdaptiveTriggers adapts modes + commit message for the portable rule.
func runKGEAdaptiveTriggers(modes map[string]string, text string) string {
	ctx := kge.Context{
		CommitMsg: text,
		Modes:     modes,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg: true,
			kge.CapModes:     true,
		},
	}
	eng := kge.NewEngine(kge.AdaptiveTriggersRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEEvidenceHierarchy adapts commit message + staged paths for the portable rule.
func runKGEEvidenceHierarchy(text string, staged []string) string {
	ctx := kge.Context{
		CommitMsg:   text,
		StagedPaths: staged,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapStagedPaths: true,
		},
	}
	eng := kge.NewEngine(kge.EvidenceHierarchyRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEPlanStatusSync adapts commit message + staged paths for the portable rule.
func runKGEPlanStatusSync(text string, staged []string) string {
	ctx := kge.Context{
		CommitMsg:   text,
		StagedPaths: staged,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapStagedPaths: true,
		},
	}
	eng := kge.NewEngine(kge.PlanStatusSyncRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEPlanCheckboxSync loads per-plan staged diffs into PathDiffs.
func runKGEPlanCheckboxSync(text string, staged []string, root string) string {
	pathDiffs := map[string]string{}
	for _, ref := range kge.FindActivePlanRefs(text) {
		clean := strings.TrimRight(ref, "),]\"")
		if clean == "" {
			continue
		}
		diff := stagedDiffCached(root, clean)
		if diff != "" {
			pathDiffs[clean] = diff
		}
	}
	ctx := kge.Context{
		CommitMsg:   text,
		StagedPaths: staged,
		PathDiffs:   pathDiffs,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:   true,
			kge.CapStagedPaths: true,
			kge.CapStagedDiff:  true,
		},
	}
	eng := kge.NewEngine(kge.PlanCheckboxSyncRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEPlanArchivalAudit loads staged archived plan contents for the portable rule.
func runKGEPlanArchivalAudit(text string, staged []string, root string) string {
	contents := map[string]string{}
	for _, s := range staged {
		p := filepath.ToSlash(s)
		if !strings.HasPrefix(p, "plans/archived/") || !strings.HasSuffix(p, ".md") {
			continue
		}
		body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(p)))
		if err != nil {
			continue
		}
		contents[p] = string(body)
	}
	ctx := kge.Context{
		CommitMsg:    text,
		StagedPaths:  staged,
		FileContents: contents,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:     true,
			kge.CapStagedPaths:   true,
			kge.CapStagedContent: true,
		},
	}
	eng := kge.NewEngine(kge.PlanArchivalAuditRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGENoNewShellScripts adapts git Added paths for the portable rule.
func runKGENoNewShellScripts(text string, root string) string {
	added, err := gitLines(root, "diff", "--cached", "--diff-filter=A", "--name-only")
	if err != nil {
		return "" // fail-open: don't block on git error
	}
	ctx := kge.Context{
		CommitMsg:  text,
		AddedPaths: added,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:  true,
			kge.CapAddedPaths: true,
		},
	}
	eng := kge.NewEngine(kge.NoNewShellScriptsRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGERuntimeTriggerWiring loads PathDiffs + discovery/Go corpus + routing-registry body.
func runKGERuntimeTriggerWiring(text string, staged []string, root string) string {
	pathDiffs := map[string]string{}
	for _, s := range staged {
		p := filepath.ToSlash(s)
		if p == "knowledge/runtime/routing-registry.yaml" ||
			(strings.HasPrefix(p, "runtime/") && strings.HasSuffix(p, ".yaml")) {
			diff := stagedDiffCached(root, p)
			if diff != "" {
				pathDiffs[p] = diff
			}
		}
	}
	var corpus strings.Builder
	if b, err := os.ReadFile(filepath.Join(root, "runtime", "cognitive-modes-discovery.yaml")); err == nil {
		corpus.Write(b)
		corpus.WriteByte('\n')
	}
	appendGoCorpus(&corpus, filepath.Join(root, "scripts", "ai-skill-cli"))
	contents := map[string]string{}
	if b, err := os.ReadFile(filepath.Join(root, "knowledge", "runtime", "routing-registry.yaml")); err == nil {
		contents["knowledge/runtime/routing-registry.yaml"] = string(b)
	}
	ctx := kge.Context{
		CommitMsg:    text,
		StagedPaths:  staged,
		PathDiffs:    pathDiffs,
		FileContents: contents,
		SearchCorpus: corpus.String(),
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:     true,
			kge.CapStagedPaths:   true,
			kge.CapStagedDiff:    true,
			kge.CapSearchCorpus:  true,
			kge.CapStagedContent: true,
		},
	}
	eng := kge.NewEngine(kge.RuntimeTriggerWiringRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

func appendGoCorpus(b *strings.Builder, dir string) {
	_ = filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(p, ".go") {
			return nil
		}
		data, rerr := os.ReadFile(p)
		if rerr != nil {
			return nil
		}
		b.Write(data)
		b.WriteByte('\n')
		return nil
	})
}

// runKGEEnforcementRuleRegistrySync loads BoundPaths from the registry snapshot
// and candidate YAML contents for the portable rule.
func runKGEEnforcementRuleRegistrySync(text string, staged []string, root string) string {
	regPath := filepath.Join(root, "enforcement", "enforcement-registry.yaml")
	snap, err := loadRegistrySnapshotFromPath(regPath)
	if err != nil {
		return ""
	}
	bound := map[string]bool{}
	for _, rc := range snap.RuleClasses {
		for _, sf := range rc.SourceFiles {
			bound[normalizeSourcePath(sf)] = true
		}
	}
	contents := map[string]string{}
	for _, p := range staged {
		rel := filepath.ToSlash(strings.TrimSpace(p))
		if !strings.HasPrefix(rel, "enforcement/") {
			continue
		}
		if !strings.HasSuffix(rel, ".yaml") && !strings.HasSuffix(rel, ".yml") {
			continue
		}
		if rel == "enforcement/enforcement-registry.yaml" {
			continue
		}
		body, readErr := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
		if readErr != nil {
			continue
		}
		contents[rel] = string(body)
	}
	ctx := kge.Context{
		CommitMsg:    text,
		StagedPaths:  staged,
		FileContents: contents,
		BoundPaths:   bound,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:     true,
			kge.CapStagedPaths:   true,
			kge.CapStagedContent: true,
			kge.CapBoundPaths:    true,
		},
	}
	eng := kge.NewEngine(kge.EnforcementRuleRegistrySyncRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// runKGEPlanEvidenceConvention loads CapRepoFS listings + README contents.
func runKGEPlanEvidenceConvention(text string, staged []string, root string) string {
	existing := map[string]bool{}
	listings := map[string][]string{}
	contents := map[string]string{}

	planDirs := collectPlanDirsForEvidenceConvention(staged, root)
	for planDir := range planDirs {
		mainRel := planDir + "/_plan.md"
		if _, ok := readFileString(root, mainRel); ok {
			existing[mainRel] = true
		}
		flatSibling := planDir + ".md"
		if st, err := os.Stat(filepath.Join(root, flatSibling)); err == nil && !st.IsDir() {
			existing[flatSibling] = true
		}
		evDir := planDir + "/evidence"
		if st, err := os.Stat(filepath.Join(root, evDir)); err == nil && st.IsDir() {
			existing[evDir] = true
		}
		readmeRel := planDir + "/evidence/README.md"
		if body, ok := readFileString(root, readmeRel); ok {
			existing[readmeRel] = true
			contents[readmeRel] = body
		}
		files, err := evidenceFilesInDir(root, planDir)
		if err == nil {
			listings[evDir] = files
		}
	}

	ctx := kge.Context{
		CommitMsg:     text,
		StagedPaths:   staged,
		ExistingPaths: existing,
		DirListings:   listings,
		FileContents:  contents,
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:     true,
			kge.CapStagedPaths:   true,
			kge.CapRepoFS:        true,
			kge.CapStagedContent: true,
		},
	}
	eng := kge.NewEngine(kge.PlanEvidenceConventionRule{})
	return kgeFindingsMessage(eng.Run(ctx))
}

// countKGEAdvisories runs advisory-only rules (D9 commit-msg count path).
// Does not run validation or discovery rules.
func countKGEAdvisories(root string, staged []string) int {
	ctx, err := buildKGEWorkspaceContext(root)
	if err != nil {
		return 0
	}
	if len(staged) > 0 {
		ctx.StagedPaths = make([]string, len(staged))
		for i, p := range staged {
			ctx.StagedPaths[i] = filepath.ToSlash(p)
		}
		contents := map[string]string{}
		for _, p := range ctx.StagedPaths {
			lower := strings.ToLower(p)
			if !strings.HasSuffix(lower, ".md") && !strings.HasSuffix(lower, ".markdown") {
				continue
			}
			body, readErr := os.ReadFile(filepath.Join(root, filepath.FromSlash(p)))
			if readErr != nil {
				continue
			}
			contents[p] = string(body)
		}
		ctx.FileContents = contents
		if ctx.Provided == nil {
			ctx.Provided = map[kge.CapabilityID]bool{}
		}
		if len(contents) > 0 {
			ctx.Provided[kge.CapStagedContent] = true
		} else {
			delete(ctx.Provided, kge.CapStagedContent)
		}
	}
	eng := kge.NewEngine(kge.AdvisoryRules()...)
	findings := eng.RunAvailable(ctx)
	_, adv, _ := kge.Partition(findings)
	return len(adv)
}

// attachKGEAdvisoryCount adds D9 commit-msg count-only warning (non-blocking).
func attachKGEAdvisoryCount(result Result, root string, staged []string) Result {
	n := countKGEAdvisories(root, staged)
	if n <= 0 {
		return result
	}
	result.Checks = append(result.Checks, Check{
		Name:    "kge_advisory",
		Status:  "warning",
		Message: kge.FormatCommitSummary(n, "ai-skill kge validate --advisory"),
	})
	return result
}

// attachKGECheck runs the full default pack with D9 check presentation.
// Blocks only on validation (error) findings.
func attachKGECheck(result Result, root string) Result {
	ctx, err := buildKGEWorkspaceContext(root)
	if err != nil {
		result.Checks = append(result.Checks, Check{
			Name:    "kge_check",
			Status:  "warning",
			Message: "kge check skipped: " + err.Error(),
		})
		return result
	}
	eng := kge.NewEngine(kge.DefaultRules()...)
	findings := eng.RunAvailable(ctx)
	report := kge.FormatCheckReport(findings, 3)
	if kge.Blocking(findings) {
		result.Status = "blocked"
		result.ExitCode = ExitValidationFailed
		result.Checks = append(result.Checks, Check{Name: "kge_check", Status: "error", Message: report})
		result.Error = &CommandError{
			Code:        "kge_check_failed",
			Message:     report,
			Remediation: "Fix validation findings, or inspect advisories with `ai-skill kge validate --advisory`.",
		}
		return result
	}
	status := "ok"
	if strings.Contains(report, "recommendation") {
		status = "warning"
	}
	result.Checks = append(result.Checks, Check{Name: "kge_check", Status: status, Message: report})
	return result
}
