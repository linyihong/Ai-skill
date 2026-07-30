package kge

// Severity classifies a Finding. Only Error is default-blocking for adapters.
type Severity string

const (
	SeverityError    Severity = "error"    // validation block candidate
	SeverityWarning  Severity = "warning"  // advisory
	SeverityInfo     Severity = "info"     // discovery / telemetry-lite
)

// CapabilityID names a unit of context an adapter may supply.
type CapabilityID string

const (
	CapCommitMsg    CapabilityID = "cap.commit_msg"
	CapModes        CapabilityID = "cap.modes"
	CapStagedPaths  CapabilityID = "cap.staged_paths"
	CapStagedContent CapabilityID = "cap.staged_content"
	CapStagedDiff   CapabilityID = "cap.staged_diff"
	CapRepoFS       CapabilityID = "cap.repo_fs"
	CapPathCochange CapabilityID = "cap.path_cochange"
	CapKnownSignals CapabilityID = "cap.known_signals"
	CapAddedPaths   CapabilityID = "cap.added_paths"
	CapSearchCorpus CapabilityID = "cap.search_corpus"
)

// Context is the adapter-normalized input. No git handles.
type Context struct {
	RepoRoot     string
	CommitMsg    string
	Modes        map[string]string // cognitive modes if parsed
	StagedPaths  []string
	AddedPaths   []string          // CapAddedPaths: newly added paths (diff-filter=A)
	StagedDiff   string            // empty = not supplied (may fail rules that require CapStagedDiff)
	PathDiffs    map[string]string // CapStagedDiff: optional path → unified diff (per-file)
	FileContents map[string]string // path → content when CapStagedContent provided
	ExistingPaths map[string]bool  // CapRepoFS: paths attested to exist on disk
	KnownSignals  map[string]bool  // CapKnownSignals: discovery signal vocabulary
	SearchCorpus  string           // CapSearchCorpus: searchable text blob from adapter
	// Provided lists which capabilities the adapter filled.
	Provided map[CapabilityID]bool
}

// Has reports whether the adapter marked a capability as provided.
func (c Context) Has(id CapabilityID) bool {
	if c.Provided == nil {
		return false
	}
	return c.Provided[id]
}

// Finding is one rule result row.
type Finding struct {
	RuleID   string
	Severity Severity
	Message  string
	Code     string // stable machine code when useful
	Path     string // optional repo-relative path for IDE / MCP diagnostics
}

// Rule is a pure Validate(Context) → Findings contract.
type Rule interface {
	ID() string
	Kind() Kind
	RequiredCapabilities() []CapabilityID
	Validate(ctx Context) []Finding
}

// Kind mirrors plan taxonomy (validation block vs advisory remind).
type Kind string

const (
	KindValidation Kind = "validation"
	KindAdvisory   Kind = "advisory"
	KindDiscovery  Kind = "discovery"
)
