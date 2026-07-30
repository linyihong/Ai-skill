package app

// Plan-tree validators (Phase 2 of 2026-06-02-1200-plan-tree-hierarchy-governance).
//
// Five commit-msg validators that mechanically enforce the plan tree frontmatter
// schema established in 01-frontmatter-schema.md:
//
//   - validatePlanTreeFrontmatter        block    sub/spike missing required fields
//   - validatePlanTreeArchiveOrder       block    main archive blocked by pending required child
//   - validatePlanTreeParentReference    block    parent: <id> must resolve to a real plan
//   - validatePlanTreeUniqueID           block    no two plans may share frontmatter id
//   - validatePlanTreeFolderConvention   warning  folder shape advisory (depth/_plan.md/NN- prefix)
//
// All validators are pre-existing-plan friendly: a plan file without YAML
// frontmatter is silently skipped (the Phase 4 migration sub-plan handles
// retro-fitting old plans). This avoids a one-shot break of the existing
// repository when these validators land.
//
// Files inside any path segment named "fixtures" are excluded from cross-plan
// scans (uniqueness / parent-existence indexes), so example/testdata files
// shipped alongside the schema docs don't collide with real plans.
//
// See: plans/active/2026-06-02-1200-plan-tree-hierarchy-governance/02-validator-implementation.md

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// PlanFrontmatter captures the minimal-governance frontmatter fields used by
// the five plan-tree validators. Unknown fields are tolerated and ignored.
type PlanFrontmatter struct {
	Path                  string // repo-relative path (e.g. "plans/active/foo.md")
	HasFrontmatter        bool   // true if a YAML frontmatter block was found
	ID                    string
	PlanKind              string // "main" | "sub" | "spike" | ""
	Status                string // "draft" | "in-progress" | "completed" | ""
	Parent                string
	HasParentField        bool // distinguish parent: null from missing parent field
	RequiredForCompletion *bool
	HasReasonField        bool
	SubPlanReason         string // raw trimmed value; empty string = block
	SchemaVersion         string // declared schema_version (quotes stripped); "" = absent
	// Delegation is the optional nested `delegation` object (sub-plan 03).
	// nil = undeclared. Ai-skill consumer-layer ONLY — the portable planvalidate
	// engine deliberately has zero knowledge of this field (see Consumer Exclusive
	// test). enabled:false is treated identically to undeclared.
	Delegation *DelegationSpec
}

// flexStrings tolerates a YAML value that is either a scalar (`acceptance: done`)
// or a sequence (`acceptance: [a, b]`), so the delegation brief can be authored
// either way. An absent/null value decodes to nil.
type flexStrings []string

func (f *flexStrings) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode {
		if strings.TrimSpace(value.Value) == "" || value.Tag == "!!null" {
			*f = nil
			return nil
		}
		*f = flexStrings{value.Value}
		return nil
	}
	var s []string
	if err := value.Decode(&s); err != nil {
		return err
	}
	*f = flexStrings(s)
	return nil
}

func (f flexStrings) hasContent() bool {
	for _, x := range f {
		if strings.TrimSpace(x) != "" {
			return true
		}
	}
	return false
}

// DelegationSpec mirrors the capability-first delegation schema (sub-plan 03,
// Phase 1). `brief` is the portable, tool-neutral capability description;
// `execution` is the workflow (paths + optional constraints). Required-when-
// enabled set = brief.goal, brief.acceptance, brief.verification, execution.modes.
// brief.context and execution.constraints are always optional.
type DelegationSpec struct {
	Enabled bool `yaml:"enabled"`
	Brief   struct {
		Goal         string      `yaml:"goal"`
		Acceptance   flexStrings `yaml:"acceptance"`
		Verification flexStrings `yaml:"verification"`
		Context      struct {
			Required flexStrings `yaml:"required"`
		} `yaml:"context"`
	} `yaml:"brief"`
	Execution struct {
		Modes       flexStrings `yaml:"modes"`
		Constraints flexStrings `yaml:"constraints"`
	} `yaml:"execution"`
}

var (
	planTreeFrontmatterDelim = []byte("---")
	planTreeYAMLKeyValueRE   = regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)\s*:\s*(.*)$`)
)

// parsePlanFrontmatterFromBytes extracts the frontmatter fields from a plan
// markdown body. Returns a zero PlanFrontmatter with HasFrontmatter=false if
// the file does not start with a `---` frontmatter block.
func parsePlanFrontmatterFromBytes(path string, data []byte) PlanFrontmatter {
	pf := PlanFrontmatter{Path: path}
	text := string(data)
	// Tolerate UTF-8 BOM (some Windows editors prepend it) and leading whitespace.
	text = strings.TrimLeft(text, "\ufeff \t\r\n")
	if !strings.HasPrefix(text, "---") {
		return pf
	}
	idx := strings.Index(text, "---")
	if idx < 0 {
		return pf
	}
	rest := text[idx+3:]
	// Find the closing "---" on a line by itself.
	lines := strings.Split(rest, "\n")
	var body []string
	closed := false
	for i, line := range lines {
		if i == 0 && strings.TrimSpace(line) == "" {
			continue
		}
		if strings.TrimSpace(line) == "---" {
			closed = true
			break
		}
		body = append(body, line)
	}
	if !closed {
		return pf
	}
	pf.HasFrontmatter = true

	// Walk lines; tolerate folded scalars (`>` / `|`) for sub_plan_reason.
	var current string
	var foldingKey string
	var folded []string
	flushFolded := func() {
		if foldingKey == "" {
			return
		}
		val := strings.TrimSpace(strings.Join(folded, " "))
		assignField(&pf, foldingKey, val)
		foldingKey = ""
		folded = nil
	}
	for _, raw := range body {
		line := strings.TrimRight(raw, "\r")
		if foldingKey != "" {
			// Continuation if indented; otherwise flush and re-parse.
			if strings.HasPrefix(line, "  ") || strings.HasPrefix(line, "\t") {
				folded = append(folded, strings.TrimSpace(line))
				continue
			}
			flushFolded()
		}
		m := planTreeYAMLKeyValueRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		key := m[1]
		val := strings.TrimSpace(m[2])
		current = key
		_ = current
		if val == ">" || val == ">-" || val == "|" || val == "|-" {
			foldingKey = key
			folded = nil
			continue
		}
		// Strip surrounding quotes.
		if (strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") && len(val) >= 2) ||
			(strings.HasPrefix(val, "'") && strings.HasSuffix(val, "'") && len(val) >= 2) {
			val = val[1 : len(val)-1]
		}
		assignField(&pf, key, val)
	}
	flushFolded()

	// Nested `delegation` object (sub-plan 03): the flat line parser above cannot
	// express nesting, so decode the whole frontmatter body once with yaml.v3 into
	// a delegation-only wrapper. On ANY unmarshal error we leave pf.Delegation nil
	// (== undeclared), preserving the pre-existing zero-behavior-change guarantee
	// for every plan the flat parser already tolerates.
	var dw struct {
		Delegation *DelegationSpec `yaml:"delegation"`
	}
	if err := yaml.Unmarshal([]byte(strings.Join(body, "\n")), &dw); err == nil {
		pf.Delegation = dw.Delegation
	}
	return pf
}

func assignField(pf *PlanFrontmatter, key, val string) {
	switch key {
	case "id":
		pf.ID = val
	case "plan_kind":
		pf.PlanKind = val
	case "status":
		pf.Status = val
	case "parent":
		pf.HasParentField = true
		if val == "null" || val == "~" || val == "" {
			pf.Parent = ""
		} else {
			pf.Parent = val
		}
	case "required_for_completion":
		b := strings.ToLower(strings.TrimSpace(val)) == "true"
		f := strings.ToLower(strings.TrimSpace(val)) == "false"
		if b {
			t := true
			pf.RequiredForCompletion = &t
		} else if f {
			t := false
			pf.RequiredForCompletion = &t
		}
	case "sub_plan_reason":
		pf.HasReasonField = true
		pf.SubPlanReason = strings.TrimSpace(val)
	case "schema_version":
		// Quotes already stripped above ("1" -> 1), satisfying the Q3 loader
		// requirement; carried into RawPlan.SchemaVersion for the compat layer.
		pf.SchemaVersion = strings.TrimSpace(val)
	}
}

// parsePlanFrontmatterFile reads a path and parses its frontmatter.
func parsePlanFrontmatterFile(absPath string) (PlanFrontmatter, error) {
	data, err := os.ReadFile(absPath)
	if err != nil {
		return PlanFrontmatter{}, err
	}
	rel := absPath
	// Best-effort rel path: walk-up logic happens at call site.
	pf := parsePlanFrontmatterFromBytes(rel, data)
	return pf, nil
}

// scanAllPlanFrontmatter walks both plans/active and plans/archived under root
// and returns every parsed frontmatter. Files under any "fixtures" segment are
// excluded (those are documentation testdata, not real plans). Files without
// frontmatter are returned with HasFrontmatter=false so callers can choose to
// skip them or include them in coverage stats.
func scanAllPlanFrontmatter(root string) []PlanFrontmatter {
	var out []PlanFrontmatter
	for _, sub := range []string{"plans/active", "plans/archived"} {
		base := sub
		if root != "" {
			base = filepath.Join(root, sub)
		}
		_ = filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				return nil
			}
			if !strings.HasSuffix(strings.ToLower(d.Name()), ".md") {
				return nil
			}
			rel, _ := filepath.Rel(root, path)
			rel = filepath.ToSlash(rel)
			if pathContainsFixturesSegment(rel) {
				return nil
			}
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				return nil
			}
			pf := parsePlanFrontmatterFromBytes(rel, data)
			pf.Path = rel
			out = append(out, pf)
			return nil
		})
	}
	return out
}

func pathContainsFixturesSegment(rel string) bool {
	for _, seg := range strings.Split(rel, "/") {
		if seg == "fixtures" {
			return true
		}
	}
	return false
}

// stagedPlanPaths returns the subset of staged paths under plans/active or
// plans/archived that look like plan markdown.
func stagedPlanPaths(staged []string) []string {
	var out []string
	for _, s := range staged {
		s = filepath.ToSlash(s)
		if (strings.HasPrefix(s, "plans/active/") || strings.HasPrefix(s, "plans/archived/")) &&
			strings.HasSuffix(strings.ToLower(s), ".md") &&
			!pathContainsFixturesSegment(s) {
			out = append(out, s)
		}
	}
	return out
}

// readStagedPlan parses the staged plan path against the working tree (post-
// stage, pre-commit reflects what will be committed).
func readStagedPlan(root, rel string) (PlanFrontmatter, bool) {
	abs := rel
	if root != "" {
		abs = filepath.Join(root, rel)
	}
	data, err := os.ReadFile(abs)
	if err != nil {
		return PlanFrontmatter{}, false
	}
	pf := parsePlanFrontmatterFromBytes(rel, data)
	pf.Path = rel
	return pf, true
}

// ---------------------------------------------------------------------------
// Validator 1: validatePlanTreeFrontmatter (block)
//
// Sub or spike plans must declare: parent (non-empty), sub_plan_reason
// (non-empty string), required_for_completion (bool). Main plans require
// parent (may be null) but the schema permits an absent parent field on main;
// we accept either explicit `parent: null` or no parent field.
// ---------------------------------------------------------------------------
func validatePlanTreeFrontmatter(text string, staged []string, root string) string {
	return runKGEPlanTreeFrontmatter(text, staged, root)
}

// ---------------------------------------------------------------------------
// Validator 2: validatePlanTreeArchiveOrder (block)
//
// When a main plan is being archived (its _plan.md or top-level .md moved
// into plans/archived/), every sub-plan declaring parent == <main>.id with
// required_for_completion: true must be status: completed (location-agnostic:
// still-active OR already-archived both qualify as long as status==completed).
// ---------------------------------------------------------------------------
func validatePlanTreeArchiveOrder(text string, staged []string, root string) string {
	return runKGEPlanTreeArchiveOrder(text, staged, root)
}

func displayStatus(s string) string {
	if s == "" {
		return "<missing>"
	}
	return s
}

// ---------------------------------------------------------------------------
// Validator 3: validatePlanTreeParentReference (block)
//
// Every sub/spike plan in the staged set whose parent field is non-empty must
// reference an id that exists somewhere in the repository (active or archived,
// excluding fixtures/). Prevents dangling parent pointers.
// ---------------------------------------------------------------------------
func validatePlanTreeParentReference(text string, staged []string, root string) string {
	return runKGEPlanTreeParentReference(text, staged, root)
}

// ---------------------------------------------------------------------------
// Validator 4: validatePlanTreeUniqueID (block)
//
// No two plans (across active + archived, excluding fixtures/) may share an
// `id:` frontmatter value. Fires when staged plans introduce or modify the id.
// ---------------------------------------------------------------------------
func validatePlanTreeUniqueID(text string, staged []string, root string) string {
	return runKGEPlanTreeUniqueID(text, staged, root)
}

// ---------------------------------------------------------------------------
// Validator 5: validatePlanTreeFolderConvention (warning)
//
// Warning-only advisory checks against the UI convention:
//   - A sub-plan folder (one that contains any sub-plan files) should contain a _plan.md.
//   - Files inside a plan folder should match `^\d{2}-` (NN- prefix) or be `_plan.md`.
//   - Path depth under plans/active or plans/archived should be < 3 levels.
// Returns warnings prefixed with `plan-tree-folder-convention (warning):` —
// hooks.go renders warnings without blocking.
// ---------------------------------------------------------------------------
func validatePlanTreeFolderConvention(text string, staged []string, root string) string {
	return runKGEPlanTreeFolderConvention(text, staged, root)
}

// hasOptOutTrailer returns true if the commit message body contains the given
// trailer on a line by itself (case-sensitive, whitespace-trimmed).
func hasOptOutTrailer(text, trailer string) bool {
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == trailer {
			return true
		}
	}
	return false
}
