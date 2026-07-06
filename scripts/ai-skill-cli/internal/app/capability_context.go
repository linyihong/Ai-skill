package app

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// capabilityContextContractTargetKey is projected to generated_surfaces as
// runtime.capability_context.contract (runtime/capability-context.yaml).
const capabilityContextContractTargetKey = "runtime.capability_context.contract"

type capabilityRegistry struct {
	RegistryVersion   string                    `yaml:"registry_version"`
	Status            string                    `yaml:"status"`
	OwnerLayer        string                    `yaml:"owner_layer"`
	CanonicalContract string                    `yaml:"canonical_contract"`
	Schema            string                    `yaml:"schema"`
	StanceEnum        capabilityStanceEnum      `yaml:"stance_enum"`
	Capabilities      []capabilityRegistryEntry `yaml:"capabilities"`
}

type capabilityStanceEnum struct {
	Standardized   []string `yaml:"standardized"`
	Implicit       []string `yaml:"implicit"`
	ReservedPolicy string   `yaml:"reserved_policy"`
}

type capabilityRegistryEntry struct {
	ID                  string                    `yaml:"id"`
	Status              string                    `yaml:"status"`
	Summary             string                    `yaml:"summary"`
	RequiresContext     capabilityRequiresContext `yaml:"requires_context"`
	Artifact            string                    `yaml:"artifact"`
	TypicalCallerSlices []string                  `yaml:"typical_caller_slices"`
}

type capabilityRequiresContext struct {
	Stance []string `yaml:"stance"`
}

type capabilityInvokeValidation struct {
	CapabilityID string
	Stance       string
	Warnings     []string
}

func capabilityRegistryPath(repo string) string {
	return filepath.Join(repo, "knowledge", "runtime", "capability-registry.yaml")
}

func readCapabilityRegistry(repo string) (capabilityRegistry, error) {
	content, err := os.ReadFile(capabilityRegistryPath(repo))
	if err != nil {
		return capabilityRegistry{}, err
	}
	var registry capabilityRegistry
	if err := yaml.Unmarshal(content, &registry); err != nil {
		return capabilityRegistry{}, err
	}
	return registry, nil
}

func validateCapabilityRegistry(registry capabilityRegistry) error {
	if strings.TrimSpace(registry.RegistryVersion) == "" {
		return fmt.Errorf("capability registry missing registry_version")
	}
	if len(registry.Capabilities) == 0 {
		return fmt.Errorf("capability registry has no capabilities")
	}
	standardized := map[string]bool{}
	for _, v := range registry.StanceEnum.Standardized {
		standardized[strings.TrimSpace(v)] = true
	}
	implicit := map[string]bool{}
	for _, v := range registry.StanceEnum.Implicit {
		implicit[strings.TrimSpace(v)] = true
	}
	seen := map[string]bool{}
	for _, cap := range registry.Capabilities {
		id := strings.TrimSpace(cap.ID)
		if id == "" {
			return fmt.Errorf("capability entry missing id")
		}
		if seen[id] {
			return fmt.Errorf("duplicate capability id: %s", id)
		}
		seen[id] = true
		if strings.TrimSpace(cap.Status) == "" {
			return fmt.Errorf("%s missing status", id)
		}
		if strings.TrimSpace(cap.Summary) == "" {
			return fmt.Errorf("%s missing summary", id)
		}
		for _, stance := range cap.RequiresContext.Stance {
			s := strings.TrimSpace(stance)
			if s == "" {
				return fmt.Errorf("%s has empty requires_context.stance entry", id)
			}
			if implicit[s] {
				return fmt.Errorf("%s requires_context.stance must not list implicit value %q", id, s)
			}
			if !standardized[s] {
				return fmt.Errorf("%s requires_context.stance %q not in stance_enum.standardized", id, s)
			}
		}
	}
	return nil
}

func nativeCapabilityRegistryValidation(repo string) Check {
	registry, err := readCapabilityRegistry(repo)
	if err != nil {
		return Check{
			Name:    "capability_registry",
			Status:  "failed",
			Message: err.Error(),
		}
	}
	if err := validateCapabilityRegistry(registry); err != nil {
		return Check{
			Name:    "capability_registry",
			Status:  "failed",
			Message: err.Error(),
		}
	}
	active := 0
	withStance := 0
	for _, cap := range registry.Capabilities {
		if cap.Status == "active" {
			active++
		}
		if len(cap.RequiresContext.Stance) > 0 {
			withStance++
		}
	}
	return Check{
		Name:   "capability_registry",
		Status: "ok",
		Message: fmt.Sprintf(
			"capability registry ok: capabilities=%d active=%d requires_context.stance=%d standardized_stances=%v",
			len(registry.Capabilities),
			active,
			withStance,
			registry.StanceEnum.Standardized,
		),
	}
}

func validateCapabilityInvoke(registry capabilityRegistry, capabilityID, invokeStance string) capabilityInvokeValidation {
	result := capabilityInvokeValidation{
		CapabilityID: strings.TrimSpace(capabilityID),
		Stance:       strings.TrimSpace(invokeStance),
	}
	if result.CapabilityID == "" {
		result.Warnings = append(result.Warnings, "missing --capability")
		return result
	}

	var entry *capabilityRegistryEntry
	for i := range registry.Capabilities {
		if registry.Capabilities[i].ID == result.CapabilityID {
			entry = &registry.Capabilities[i]
			break
		}
	}
	if entry == nil {
		result.Warnings = append(result.Warnings,
			fmt.Sprintf("unknown capability %q (not in capability-registry.yaml)", result.CapabilityID))
		return result
	}
	required := entry.RequiresContext.Stance
	if len(required) == 0 {
		return result
	}

	invoke := result.Stance
	if invoke == "" || invoke == "default" {
		result.Warnings = append(result.Warnings, fmt.Sprintf(
			"capability %q requires context.stance %v but invoke omitted stance or used default (Phase 1.2: warning; no auto-fill)",
			result.CapabilityID, required))
		return result
	}

	for _, r := range required {
		if invoke == r {
			return result
		}
	}
	result.Warnings = append(result.Warnings, fmt.Sprintf(
		"capability %q requires context.stance %v but invoke has %q (Phase 1.2: warning)",
		result.CapabilityID, required, invoke))
	return result
}

func runRuntimeCapabilityInvoke(opts runtimeOptions, stdout io.Writer, stderr io.Writer) int {
	root, repoCheck := closeLoopRepoRoot(opts.repoPath)
	if repoCheck.Status != "ok" {
		_, _ = fmt.Fprintln(stderr, repoCheck.Message)
		return ExitInvalidUsage
	}
	capabilityID := strings.TrimSpace(opts.capabilityID)
	invokeStance := strings.TrimSpace(opts.invokeStance)

	registry, err := readCapabilityRegistry(root)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "capability registry: %v\n", err)
		return ExitValidationFailed
	}
	if err := validateCapabilityRegistry(registry); err != nil {
		_, _ = fmt.Fprintf(stderr, "capability registry invalid: %v\n", err)
		return ExitValidationFailed
	}

	validation := validateCapabilityInvoke(registry, capabilityID, invokeStance)
	if opts.jsonOutput {
		payload := map[string]any{
			"capability": capabilityID,
			"stance":     invokeStance,
			"warnings":   validation.Warnings,
			"ok":         len(validation.Warnings) == 0,
		}
		if err := json.NewEncoder(stdout).Encode(payload); err != nil {
			_, _ = fmt.Fprintf(stderr, "write output: %v\n", err)
			return ExitGeneralFailure
		}
		return ExitSuccess
	}

	if len(validation.Warnings) == 0 {
		_, _ = fmt.Fprintf(stdout, "capability invoke ok: %s stance=%q\n", capabilityID, invokeStance)
		return ExitSuccess
	}
	for _, w := range validation.Warnings {
		_, _ = fmt.Fprintf(stderr, "warning: %s\n", w)
	}
	_, _ = fmt.Fprintf(stdout, "capability invoke warnings=%d (Phase 1.2: exit 0)\n", len(validation.Warnings))
	return ExitSuccess
}
