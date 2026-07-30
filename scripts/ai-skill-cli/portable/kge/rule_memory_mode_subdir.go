package kge

import (
	"path/filepath"
	"strings"
)

// MemoryModeSubdirRule enforces that staged memory/ files match memory_mode.
type MemoryModeSubdirRule struct{}

func (MemoryModeSubdirRule) ID() string { return "rule.memory_mode_subdir" }
func (MemoryModeSubdirRule) Kind() Kind { return KindValidation }

func (MemoryModeSubdirRule) RequiredCapabilities() []CapabilityID {
	return []CapabilityID{CapModes, CapStagedPaths}
}

func (MemoryModeSubdirRule) Validate(ctx Context) []Finding {
	modes := ctx.Modes
	if modes == nil {
		modes = map[string]string{}
	}
	mem := modes["memory_mode"]
	allowedPrefix := ""
	switch mem {
	case "NONE":
		allowedPrefix = ""
	case "EPISODIC":
		allowedPrefix = "memory/episodic/"
	case "DECISION_REPLAY":
		allowedPrefix = "memory/decision/"
	case "FAILURE_REPLAY":
		allowedPrefix = "memory/failure/"
	case "PROJECT_CONTEXT":
		allowedPrefix = "memory/project/"
	default:
		return nil
	}

	isLayerDoc := func(f string) bool {
		return f == "memory/README.md" ||
			strings.HasPrefix(f, "memory/retrieval-governance/")
	}

	for _, f := range ctx.StagedPaths {
		f = filepath.ToSlash(f)
		if !strings.HasPrefix(f, "memory/") {
			continue
		}
		if isLayerDoc(f) {
			continue
		}
		if mem == "NONE" {
			return []Finding{{
				RuleID:   "rule.memory_mode_subdir",
				Severity: SeverityError,
				Code:     "memory_mode_subdir",
				Path:     f,
				Message:  "memory_mode=NONE but commit touches " + f + " (NONE forbids all memory/ writes per cognitive-modes-memory-integration.yaml)",
			}}
		}
		if !strings.HasPrefix(f, allowedPrefix) {
			return []Finding{{
				RuleID:   "rule.memory_mode_subdir",
				Severity: SeverityError,
				Code:     "memory_mode_subdir",
				Path:     f,
				Message:  "memory_mode=" + mem + " allows only " + allowedPrefix + " but commit touches " + f,
			}}
		}
	}
	return nil
}
