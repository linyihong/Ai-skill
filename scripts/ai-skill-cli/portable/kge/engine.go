package kge

// Engine runs registered rules against a Context.
type Engine struct {
	rules []Rule
}

// NewEngine constructs an engine with the given rule set (order preserved).
func NewEngine(rules ...Rule) *Engine {
	return &Engine{rules: append([]Rule(nil), rules...)}
}

// Run validates capability requirements then dispatches each rule.
// Missing required capability → Finding code capability_missing (error).
func (e *Engine) Run(ctx Context) []Finding {
	var out []Finding
	for _, r := range e.rules {
		if missing := missingCaps(ctx, r.RequiredCapabilities()); len(missing) > 0 {
			for _, id := range missing {
				out = append(out, Finding{
					RuleID:   r.ID(),
					Severity: SeverityError,
					Code:     "capability_missing",
					Message:  "capability_missing: " + string(id) + " required by rule " + r.ID(),
				})
			}
			continue
		}
		out = append(out, r.Validate(ctx)...)
	}
	return out
}

// RunAvailable skips rules whose required capabilities are not provided
// (no capability_missing noise). Used by `kge check` over a partial Context.
func (e *Engine) RunAvailable(ctx Context) []Finding {
	var out []Finding
	for _, r := range e.rules {
		if len(missingCaps(ctx, r.RequiredCapabilities())) > 0 {
			continue
		}
		out = append(out, r.Validate(ctx)...)
	}
	return out
}

func missingCaps(ctx Context, need []CapabilityID) []CapabilityID {
	var miss []CapabilityID
	for _, id := range need {
		if !ctx.Has(id) {
			miss = append(miss, id)
		}
	}
	return miss
}

// Blocking reports whether any finding should fail a validation adapter.
func Blocking(findings []Finding) bool {
	for _, f := range findings {
		if f.Severity == SeverityError {
			return true
		}
	}
	return false
}
