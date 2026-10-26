package app

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const unityFXRoute = "route.analysis.unity-fx"
const unityFXSource = "analysis/apk/workflows/unity-fx-driver-trace.md"

// Domain wiring invariant, reached by runtime refresh/validate and their hooks.
// Small fixture repos without the APK domain are outside this invariant.
func unityFXDomainPresent(repo string) bool {
	_, err := os.Stat(filepath.Join(repo, "analysis/apk/README.md"))
	return err == nil
}

func validateUnityFXDiscovery(repo string, registry runtimeRoutingRegistry, summaries []discoverySummary) error {
	if !unityFXDomainPresent(repo) {
		return nil
	}
	var route *runtimeRouteRecord
	for i := range registry.Records {
		if registry.Records[i].ID == unityFXRoute {
			route = &registry.Records[i]
			break
		}
	}
	if route == nil || route.PrimarySource != unityFXSource || route.ActivationMode != "auto-detect" {
		return fmt.Errorf("Unity FX discovery: missing active canonical route %s", unityFXRoute)
	}
	for _, dep := range []string{"analysis/apk/README.md", "analysis/apk/workflows/README.md", "workflow/apk-analysis/execution-flow.md", "workflow/apk-analysis/execution-flow.yaml"} {
		found := false
		for _, got := range route.RequiredDependencies {
			if got == dep {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("Unity FX discovery: missing dependency %s", dep)
		}
	}
	for path, anchor := range map[string]string{
		"analysis/apk/README.md":                    "unity-fx-driver-trace.md",
		"analysis/apk/workflows/README.md":          "unity-fx-driver-trace.md",
		"workflow/apk-analysis/README.md":           "unity-fx-driver-trace.md",
		"workflow/apk-analysis/execution-flow.md":   "unity-fx-driver-trace.md",
		"workflow/apk-analysis/execution-flow.yaml": "gate.apk_analysis.unity_fx_driver_evidence",
		"knowledge/indexes/README.md":               "unity-fx-driver-trace.md",
		"knowledge/summaries/unity-fx.md":           "../../" + unityFXSource,
	} {
		b, err := os.ReadFile(filepath.Join(repo, path))
		if err != nil || !strings.Contains(string(b), anchor) {
			return fmt.Errorf("Unity FX discovery: %s missing %s", path, anchor)
		}
	}
	foundSummary := false
	for _, s := range summaries {
		if s.AtomID == "unity-fx" {
			foundSummary = true
		}
	}
	if !foundSummary {
		return fmt.Errorf("Unity FX discovery: missing unity-fx summary identity")
	}
	for _, msg := range []string{"Animator.Play", "m_Payloads", "PlayableDirector", "Unity 特效"} {
		found := false
		for _, d := range DetectWorkflows(registry, []DetectorMessage{{Role: "user", Text: msg}}, nil) {
			if d.RouteID == unityFXRoute {
				found = true
			}
		}
		if !found {
			return fmt.Errorf("Unity FX discovery: detector lost signal %s", msg)
		}
	}
	// Exercise the real scoring path with a detector-miss style artifact request.
	// It must remain a proposal; this does not activate or lock any workflow.
	input := DiscoveryInput{UserMessage: "幫我 review docs/fx-driver.md 的 visual gap 與 payload writer 分析", Basenames: []string{"fx-driver.md"}, Paths: []string{"docs/fx-driver.md"}, Extensions: []string{".md"}}
	cfg := LoadDiscoveryConfig(filepath.Join(repo, "runtime/runtime.db"))
	if reason, ok := EligibilityCheck(input, cfg); !ok {
		return fmt.Errorf("Unity FX discovery: bridge ineligible: %s", reason)
	}
	candidates, _ := RunLightDiscovery(input, registry, summaries, cfg, repo)
	for _, c := range candidates {
		if c.Route == unityFXRoute {
			return nil
		}
	}
	return fmt.Errorf("Unity FX discovery: bridge no longer proposes %s", unityFXRoute)
}

func validateUnityFXRetrieval(repo string, db *sql.DB) error {
	if !unityFXDomainPresent(repo) {
		return nil
	}
	for _, term := range []string{"m_Payloads", "PlayableDirector"} {
		var n int
		err := db.QueryRow(`SELECT COUNT(*) FROM fts JOIN atoms ON atoms.id=fts.id WHERE fts MATCH ? AND atoms.source_path=?`, runtimeFTSMatchLiteral(term), unityFXSource).Scan(&n)
		if err != nil {
			return fmt.Errorf("Unity FX RAG: %w", err)
		}
		if n == 0 {
			return fmt.Errorf("Unity FX RAG: %s cannot retrieve canonical source", term)
		}
	}
	return nil
}
