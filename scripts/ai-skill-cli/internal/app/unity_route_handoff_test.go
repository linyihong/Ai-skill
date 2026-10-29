package app

import (
	"path/filepath"
	"testing"
)

func TestUnityDecisionRoutePositiveAndAPKFridaNegative(t *testing.T) {
	repo := discoveryRepoRoot(t)
	reg, err := readRuntimeRoutingRegistry(filepath.Join(repo, "knowledge/runtime/routing-registry.yaml"))
	if err != nil {
		t.Fatal(err)
	}

	contains := func(hits []DetectedRoute, id string) bool {
		for _, h := range hits {
			if h.RouteID == id {
				return true
			}
		}
		return false
	}

	t.Run("positive_unity_decision_signals", func(t *testing.T) {
		for _, msg := range []string{
			"Unity soft-enter GetContent still null",
			"acceptance_oracle for Unity soft",
			"IL2CPP contract mismatch on fixture",
		} {
			hits := DetectWorkflows(reg, []DetectorMessage{{Role: "user", Text: msg}}, nil)
			if !contains(hits, "route.analysis.unity") && !contains(hits, "route.analysis.unity-fx") {
				t.Fatalf("expected unity decision or unity-fx for %q; got %#v", msg, hits)
			}
		}
		fx := DetectWorkflows(reg, []DetectorMessage{{Role: "user", Text: "Animator.Play hit but visual gap m_Payloads"}}, nil)
		if !contains(fx, "route.analysis.unity-fx") {
			t.Fatalf("expected route.analysis.unity-fx; got %#v", fx)
		}
		regRec := routeByID(reg, "route.analysis.unity")
		if regRec == nil || regRec.PrimarySource != "analysis/unity/decision-workflow.md" {
			t.Fatalf("route.analysis.unity primary_source want decision-workflow.md, got %#v", regRec)
		}
	})

	t.Run("negative_generic_apk_frida_not_unity", func(t *testing.T) {
		for _, msg := range []string{
			"對這個 APK 用 Frida 抓包",
			"用 Frida hook 這個 app 的 HTTP API",
			"local proxy hook Netty pipeline",
		} {
			hits := DetectWorkflows(reg, []DetectorMessage{{Role: "user", Text: msg}}, nil)
			if contains(hits, "route.analysis.unity") || contains(hits, "route.analysis.unity-fx") {
				t.Fatalf("generic APK/Frida must not lock Unity routes for %q; got %#v", msg, hits)
			}
		}
	})
}

func routeByID(reg runtimeRoutingRegistry, id string) *runtimeRouteRecord {
	for i := range reg.Records {
		if reg.Records[i].ID == id {
			return &reg.Records[i]
		}
	}
	return nil
}
