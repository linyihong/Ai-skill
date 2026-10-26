package app

import (
	"database/sql"
	"path/filepath"
	"strings"
	"testing"
)

func TestUnityFXDiscoveryWiring(t *testing.T) {
	repo := discoveryRepoRoot(t)
	load := func() runtimeRoutingRegistry {
		reg, err := readRuntimeRoutingRegistry(filepath.Join(repo, "knowledge/runtime/routing-registry.yaml"))
		if err != nil {
			t.Fatal(err)
		}
		return reg
	}
	summaries := LoadDiscoverySummaries(repo)
	if err := validateUnityFXDiscovery(repo, load(), summaries); err != nil {
		t.Fatal(err)
	}
	for _, mutation := range []string{"route", "trigger", "dependency", "summary"} {
		t.Run(mutation, func(t *testing.T) {
			reg := load()
			ss := append([]discoverySummary{}, summaries...)
			for i := range reg.Records {
				if reg.Records[i].ID == unityFXRoute {
					switch mutation {
					case "route":
						reg.Records = append(reg.Records[:i], reg.Records[i+1:]...)
					case "trigger":
						reg.Records[i].ActivationTriggers = runtimeRouteTriggers{}
					case "dependency":
						reg.Records[i].RequiredDependencies = nil
					}
					break
				}
			}
			if mutation == "summary" {
				for i := range ss {
					if ss[i].AtomID == "unity-fx" {
						ss[i].AtomID = "detached"
					}
				}
			}
			if err := validateUnityFXDiscovery(repo, reg, ss); err == nil {
				t.Fatal("missing discovery wiring must block")
			}
		})
	}
}

func TestUnityFXRetrievalWiring(t *testing.T) {
	repo := discoveryRepoRoot(t)
	path := filepath.Join(t.TempDir(), "runtime-index.sqlite")
	if err := buildNativeRuntimeSQLiteIndex(repo, path); err != nil {
		t.Fatal(err)
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if err := validateUnityFXRetrieval(repo, db); err != nil {
		t.Fatal(err)
	}
	// Leave the remaining corpus intact: unrelated hits must not mask missing source.
	if _, err := db.Exec("DELETE FROM fts WHERE id IN (SELECT id FROM atoms WHERE source_path=?)", unityFXSource); err != nil {
		t.Fatal(err)
	}
	if err := validateUnityFXRetrieval(repo, db); err == nil || !strings.Contains(err.Error(), "cannot retrieve") {
		t.Fatalf("missing canonical FTS hits must block: %v", err)
	}
}
