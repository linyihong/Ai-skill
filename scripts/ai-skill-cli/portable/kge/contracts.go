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
)

// Context is the adapter-normalized input. No git handles.
type Context struct {
	RepoRoot     string
	CommitMsg    string
	Modes        map[string]string // cognitive modes if parsed
	StagedPaths  []string
	StagedDiff   string            // empty = not supplied (may fail rules that require CapStagedDiff)
	FileContents map[string]string // path → content when CapStagedContent provided
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
}

// Rule is a pure Validate(Context) → Findings contract.
type Rule interface {
	ID() string
	RequiredCapabilities() []CapabilityID
	Validate(ctx Context) []Finding
}

// Kind mirrors plan taxonomy (Phase 1 focuses on validation).
type Kind string

const (
	KindValidation Kind = "validation"
	KindAdvisory   Kind = "advisory"
)
