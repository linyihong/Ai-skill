package app

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
	"gopkg.in/yaml.v3"
)

// Phase 4.5 of plans/archived/2026-05-31-2100-mechanical-enforcement-registry.md
// — Registry Self-Governance.
//
// This file implements the engine that powers both:
//   1. `ai-skill enforcement transition-check` CLI subcommand (public surface,
//      used by validation scenarios for cross-platform fixture-based testing)
//   2. `validateEnforcementRegistryTransition` commit-msg hook validator (Go
//      registry entry obligation.commit.enforcement_registry_transition;
//      blocks commits that violate R1/R2/R3 self-governance lint rules)
//
// Both surfaces share portable kge.EnforcementRegistryTransitionRule. This
// package only loads YAML snapshots, resolves ADR/symbol IO, and presents
// findings.

// ─────────────────────────────────────────────────────────────────────
// Violation model
// ─────────────────────────────────────────────────────────────────────

type registryTransitionViolation struct {
	Code      string // R1_*, R2_*, R3_*
	RuleClass string // empty for R1 commit-msg-level violations
	From      string // empty for R1
	To        string // empty for R1
	Detail    string // human readable (or full formatted line from kge)
}

func (v registryTransitionViolation) String() string {
	if strings.HasPrefix(v.Detail, "[") {
		return v.Detail
	}
	parts := []string{"[" + v.Code + "]"}
	if v.RuleClass != "" {
		parts = append(parts, "rule_class="+v.RuleClass)
	}
	if v.From != "" || v.To != "" {
		parts = append(parts, fmt.Sprintf("transition=%s→%s", v.From, v.To))
	}
	parts = append(parts, v.Detail)
	return strings.Join(parts, " ")
}

// transitionInput bundles the three artifacts every R1/R2/R3 check needs.
type transitionInput struct {
	repoRoot  string // for ADR file resolution + executor symbol grep
	oldSnap   *registrySnapshot
	newSnap   *registrySnapshot
	commitMsg string
	// stagedPaths optional: when non-nil, CapStagedPaths is provided so the
	// rule gates on staged enforcement-registry.yaml (commit-msg path).
	stagedPaths []string
}

const transitionOptOutMarker = "[skip-registry-transition]"

func toKGERegistrySnapshot(snap *registrySnapshot) *kge.RegistrySnapshotMeta {
	if snap == nil {
		return nil
	}
	out := &kge.RegistrySnapshotMeta{
		HelperAllowlist:      append([]string(nil), snap.InternalHelperAllowlist.Symbols...),
		BindingRequiredKinds: append([]string(nil), snap.ExecutorKindSpec.BindingRequiredFor...),
	}
	for _, rc := range snap.RuleClasses {
		meta := kge.RegistryClassMeta{
			ID:           rc.ID,
			Coverage:     rc.Coverage,
			AdrReference: rc.AdrReference,
		}
		for _, ex := range rc.Executors {
			meta.Executors = append(meta.Executors, kge.RegistryExecutorMeta{
				File:         ex.File,
				Symbol:       ex.Symbol,
				ExecutorKind: ex.ExecutorKind,
			})
		}
		out.RuleClasses = append(out.RuleClasses, meta)
	}
	return out
}

func buildTransitionSymbolIndex(repo string, snap *registrySnapshot) map[string]map[string]bool {
	out := map[string]map[string]bool{}
	if snap == nil {
		return out
	}
	seen := map[string]bool{}
	for _, rc := range snap.RuleClasses {
		for _, ex := range rc.Executors {
			rel := filepath.ToSlash(strings.TrimSpace(ex.File))
			if rel == "" || !strings.HasSuffix(rel, ".go") || seen[rel] {
				continue
			}
			seen[rel] = true
			full := filepath.Join(repo, filepath.FromSlash(rel))
			data, err := os.ReadFile(full)
			syms := map[string]bool{}
			if err == nil {
				for _, m := range goFuncDeclPattern.FindAllStringSubmatch(string(data), -1) {
					if len(m) >= 2 {
						syms[m[1]] = true
					}
				}
			}
			out[rel] = syms
		}
	}
	return out
}

func buildTransitionExistingPaths(repo string, snap *registrySnapshot) map[string]bool {
	out := map[string]bool{}
	if snap == nil {
		return out
	}
	for _, rc := range snap.RuleClasses {
		adr := strings.TrimSpace(rc.AdrReference)
		if adr == "" {
			continue
		}
		rel := filepath.ToSlash(adr)
		full := filepath.Join(repo, filepath.FromSlash(rel))
		if _, err := os.Stat(full); err == nil {
			out[rel] = true
		}
	}
	return out
}

func checkRegistryTransitions(in transitionInput) []registryTransitionViolation {
	if strings.Contains(in.commitMsg, transitionOptOutMarker) {
		return nil
	}
	if in.oldSnap == nil || in.newSnap == nil {
		return nil
	}
	ctx := kge.Context{
		RepoRoot:      in.repoRoot,
		CommitMsg:     in.commitMsg,
		RegistryOld:   toKGERegistrySnapshot(in.oldSnap),
		RegistryNew:   toKGERegistrySnapshot(in.newSnap),
		ExistingPaths: buildTransitionExistingPaths(in.repoRoot, in.newSnap),
		FileSymbols:   buildTransitionSymbolIndex(in.repoRoot, in.newSnap),
		Provided: map[kge.CapabilityID]bool{
			kge.CapCommitMsg:         true,
			kge.CapRegistrySnapshots: true,
			kge.CapRepoFS:            true,
			kge.CapSymbolIndex:       true,
		},
	}
	if in.stagedPaths != nil {
		ctx.StagedPaths = in.stagedPaths
		ctx.Provided[kge.CapStagedPaths] = true
	}
	eng := kge.NewEngine(kge.EnforcementRegistryTransitionRule{})
	findings := eng.Run(ctx)
	var violations []registryTransitionViolation
	for _, f := range findings {
		violations = append(violations, registryTransitionViolation{
			Code:   f.Code,
			Detail: f.Message,
		})
	}
	return violations
}

// ─────────────────────────────────────────────────────────────────────
// `ai-skill enforcement transition-check` CLI subcommand
// ─────────────────────────────────────────────────────────────────────

type enforcementTransitionOptions struct {
	repo            string
	old             string
	newPath         string
	commitMsgFile   string
	commitMsgInline string
	expectViolation string
	jsonOutput      bool
	plainOutput     bool
}

func runEnforcementTransitionCheck(args []string, stdout io.Writer, stderr io.Writer) int {
	opts := enforcementTransitionOptions{}
	fs := newFlagSet("enforcement transition-check", stderr)
	fs.StringVar(&opts.repo, "repo", ".", "repo root (used to resolve ADR + executor file paths)")
	fs.StringVar(&opts.old, "old", "", "path to old (HEAD) enforcement-registry.yaml")
	fs.StringVar(&opts.newPath, "new", "", "path to new (staged) enforcement-registry.yaml")
	fs.StringVar(&opts.commitMsgFile, "commit-msg-file", "", "path to commit message text file")
	fs.StringVar(&opts.commitMsgInline, "commit-msg", "", "inline commit message string (alternative to --commit-msg-file)")
	fs.StringVar(&opts.expectViolation, "expect-violation", "", "assertion mode: exit 0 if any violation code contains this substring, exit 30 otherwise")
	fs.BoolVar(&opts.jsonOutput, "json", false, "write JSON output")
	fs.BoolVar(&opts.plainOutput, "plain", false, "write plain text output (default)")
	if err := fs.Parse(args); err != nil {
		return ExitInvalidUsage
	}
	if opts.jsonOutput && opts.plainOutput {
		_, _ = fmt.Fprintln(stderr, "--json and --plain are mutually exclusive")
		return ExitInvalidUsage
	}
	if strings.TrimSpace(opts.old) == "" || strings.TrimSpace(opts.newPath) == "" {
		_, _ = fmt.Fprintln(stderr, "--old and --new are required")
		return ExitInvalidUsage
	}

	root, err := resolveEnforcementRepo(opts.repo)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "resolve repo: %v\n", err)
		return ExitInvalidUsage
	}
	oldSnap, err := loadRegistrySnapshotFromPath(opts.old)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "load --old: %v\n", err)
		return ExitValidationFailed
	}
	newSnap, err := loadRegistrySnapshotFromPath(opts.newPath)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "load --new: %v\n", err)
		return ExitValidationFailed
	}
	msg := opts.commitMsgInline
	if strings.TrimSpace(opts.commitMsgFile) != "" {
		data, err := os.ReadFile(opts.commitMsgFile)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "read --commit-msg-file: %v\n", err)
			return ExitInvalidUsage
		}
		msg = string(data)
	}

	violations := checkRegistryTransitions(transitionInput{
		repoRoot:  root,
		oldSnap:   oldSnap,
		newSnap:   newSnap,
		commitMsg: msg,
	})

	result := buildTransitionCheckResult(opts, violations)
	if opts.jsonOutput {
		if err := writeJSON(stdout, result); err != nil {
			_, _ = fmt.Fprintf(stderr, "write output: %v\n", err)
			return ExitGeneralFailure
		}
		return result.ExitCode
	}
	if err := writePlain(stdout, result); err != nil {
		_, _ = fmt.Fprintf(stderr, "write output: %v\n", err)
		return ExitGeneralFailure
	}
	return result.ExitCode
}

func buildTransitionCheckResult(opts enforcementTransitionOptions, violations []registryTransitionViolation) Result {
	result := Result{
		Command:        "enforcement transition-check",
		Mode:           "report",
		Status:         "success",
		ExitCode:       ExitSuccess,
		Checks:         []Check{},
		PlannedActions: []string{},
		Mutations:      []string{},
	}
	result.Checks = append(result.Checks, Check{Name: "old_registry", Status: "ok", Message: opts.old})
	result.Checks = append(result.Checks, Check{Name: "new_registry", Status: "ok", Message: opts.newPath})
	result.Checks = append(result.Checks, Check{Name: "violations_total", Status: "ok", Message: fmt.Sprintf("%d", len(violations))})
	for _, v := range violations {
		result.Checks = append(result.Checks, Check{
			Name:    "violation." + v.Code,
			Status:  "failed",
			Message: v.String(),
		})
	}
	// Assertion mode.
	if strings.TrimSpace(opts.expectViolation) != "" {
		result.Mode = "assert"
		want := strings.ToLower(strings.TrimSpace(opts.expectViolation))
		matched := 0
		for _, v := range violations {
			if strings.Contains(strings.ToLower(v.Code), want) || strings.Contains(strings.ToLower(v.Detail), want) {
				matched++
			}
		}
		if matched > 0 {
			result.Checks = append(result.Checks, Check{Name: "assertion", Status: "ok", Message: fmt.Sprintf("matched %d violation(s)", matched)})
			return result
		}
		result.Status = "failed"
		result.ExitCode = ExitValidationFailed
		result.Checks = append(result.Checks, Check{Name: "assertion", Status: "failed", Message: "no violation matched --expect-violation substring"})
		result.Error = &CommandError{
			Code:    "enforcement_transition_assertion_unmet",
			Message: fmt.Sprintf("--expect-violation %q matched 0 violations (of %d total)", opts.expectViolation, len(violations)),
		}
		return result
	}
	// Regular mode.
	if len(violations) > 0 {
		result.Status = "failed"
		result.ExitCode = ExitValidationFailed
		result.Error = &CommandError{
			Code:        "enforcement_registry_transition_blocked",
			Message:     fmt.Sprintf("%d transition violation(s) — see violation.* checks above", len(violations)),
			Remediation: "Add [registry-status-change] trailer + rationale: line (R1); attach adr_reference for demotions (R2); ensure executor symbol exists in declared file before promoting to mechanical (R3). Opt-out via [skip-registry-transition] in commit body.",
		}
	}
	return result
}

// ─────────────────────────────────────────────────────────────────────
// commit-msg hook validator: validateEnforcementRegistryTransition
// ─────────────────────────────────────────────────────────────────────

// validateEnforcementRuleRegistrySync is the Phase 5 commit-msg validator
// (obligation.commit.enforcement_rule_registry_sync — 21st per_commit
// validator). It blocks commits that stage an enforcement/*.yaml or
// enforcement/*.md rule file without simultaneously registering it in
// enforcement-registry.yaml (either by an existing source_files binding
// or by staging the registry itself with a new rule_class entry).
//
// Relationship with Phase 3 compile-time orphan_rule lint: dual-gate.
// orphan_rule still fails compile-time on the entire registry; this
// commit-msg validator catches the same drift earlier with a tighter
// per-staged-file scope, so the failure surfaces at commit not at the
// next `ai-skill runtime compile` (which may be hours later).
//
// Opt-out: [skip-enforcement-registry-sync] trailer in commit body.
//
// Scope choice: only enforcement/ subtree for now. runtime/ and
// governance/ rule yamls are out of scope because they are typically
// edited alongside their owning module (their orphan_rule check at
// compile time is the primary gate). Expanding scope is a Phase 6+
// extension if the failure pattern surfaces there.
func validateEnforcementRuleRegistrySync(text string, staged []string, root string) string {
	return runKGEEnforcementRuleRegistrySync(text, staged, root)
}

// validateEnforcementRegistryTransition is the commit-msg validator that
// enforces Phase 4.5 R1/R2/R3 at commit time. Returns empty string on
// pass; non-empty error description blocks the commit.
//
// Triggers only when enforcement/enforcement-registry.yaml is staged.
// Reads HEAD's version via `git show HEAD:enforcement/enforcement-registry.yaml`
// and the staged version via the working-tree file (already updated by the
// developer before commit-msg fires).
//
// Opt-out: include [skip-registry-transition] in the commit body.
func validateEnforcementRegistryTransition(text string, staged []string, root string) string {
	return runKGEEnforcementRegistryTransition(text, staged, root)
}

// formatRegistryTransitionFindings preserves the legacy commit-msg block shape.
func formatRegistryTransitionFindings(findings []kge.Finding) string {
	if len(findings) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "enforcement_registry_transition: %d violation(s) — fix or use [skip-registry-transition] opt-out:\n", len(findings))
	for _, f := range findings {
		fmt.Fprintf(&b, "  - %s\n", f.Message)
	}
	return strings.TrimRight(b.String(), "\n")
}

// loadStagedRegistrySnapshots loads HEAD + working-tree registry YAML when
// enforcement-registry.yaml is staged. Returns ok=false when not applicable.
func loadStagedRegistrySnapshots(text string, staged []string, root string) (oldSnap, newSnap *registrySnapshot, errMsg string, ok bool) {
	registryRel := "enforcement/enforcement-registry.yaml"
	stagedHit := false
	for _, p := range staged {
		if filepath.ToSlash(strings.TrimSpace(p)) == registryRel {
			stagedHit = true
			break
		}
	}
	if !stagedHit {
		return nil, nil, "", false
	}
	if strings.Contains(text, transitionOptOutMarker) {
		return nil, nil, "", false
	}
	oldData, err := exec.Command("git", "-C", root, "show", "HEAD:"+registryRel).Output()
	if err != nil {
		oldData = []byte("schema_version: 2\nrule_classes: []\n")
	}
	newData, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(registryRel)))
	if err != nil {
		return nil, nil, fmt.Sprintf("enforcement_registry_transition: cannot read staged %s: %v", registryRel, err), false
	}
	var old, neu registrySnapshot
	if err := yaml.Unmarshal(oldData, &old); err != nil {
		return nil, nil, fmt.Sprintf("enforcement_registry_transition: cannot parse HEAD %s: %v", registryRel, err), false
	}
	if err := yaml.Unmarshal(newData, &neu); err != nil {
		return nil, nil, fmt.Sprintf("enforcement_registry_transition: cannot parse staged %s: %v", registryRel, err), false
	}
	return &old, &neu, "", true
}
