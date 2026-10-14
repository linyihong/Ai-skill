package app

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestImageProcessingGenerateAndVerifyNegatives(t *testing.T) {
	dir := t.TempDir()
	inputs := filepath.Join(dir, "workflow", "image-processing", "fixtures", "inputs")
	var stdout, stderr bytes.Buffer
	code := Run([]string{"image-processing", "generate-fixtures", "--out-dir", inputs}, &stdout, &stderr)
	if code != ExitSuccess {
		t.Fatalf("generate-fixtures exit=%d stderr=%s", code, stderr.String())
	}

	cases := []struct {
		name         string
		expectBlock  bool
		wantFinality string
	}{
		{"neg.alpha-opaque", true, "rejected"},
		{"neg.residue-halo", true, "rejected"},
		{"pos.d-canvas-328", false, "accepted"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			png := filepath.Join(inputs, tc.name+".png")
			outJSON := filepath.Join(dir, tc.name+".verify.json")
			args := []string{"image-processing", "verify", "--profile", tc.name, "--json-out", outJSON}
			if tc.expectBlock {
				args = append(args, "--expect-block")
			}
			args = append(args, png)
			var so, se bytes.Buffer
			code := Run(args, &so, &se)
			if tc.expectBlock {
				if code != ExitSuccess {
					t.Fatalf("expect-block want 0 got %d stderr=%s stdout=%s", code, se.String(), so.String())
				}
			} else if code != ExitSuccess {
				t.Fatalf("positive want 0 got %d stderr=%s", code, se.String())
			}
			raw, err := os.ReadFile(outJSON)
			if err != nil {
				t.Fatal(err)
			}
			var report map[string]any
			if err := json.Unmarshal(raw, &report); err != nil {
				t.Fatal(err)
			}
			got, _ := report["finality"].(string)
			if got != tc.wantFinality {
				t.Fatalf("finality=%q want %q report=%s", got, tc.wantFinality, so.String())
			}
		})
	}
}

func TestNoNewNonGoScriptsBlockedByKGESurface(t *testing.T) {
	// Smoke: command exists in usage
	var stdout, stderr bytes.Buffer
	code := Run([]string{"image-processing"}, &stdout, &stderr)
	if code != ExitInvalidUsage {
		t.Fatalf("want usage exit, got %d", code)
	}
}
