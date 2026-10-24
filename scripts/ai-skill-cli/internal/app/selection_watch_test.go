package app

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestSelectionWatchFailureAndUsage(t *testing.T) {
	for _, kind := range []string{"ready", "timeout", "invalid", "usage-zero", "usage-invalid", "preexisting"} {
		t.Run(kind, func(t *testing.T) {
			dir := t.TempDir()
			request := filepath.Join(dir, "request.json")
			response := filepath.Join(dir, "response.json")
			usage := filepath.Join(dir, "usage.json")
			writeFile(t, request, `{"system":"Select direct only","query":"probe","candidates":[{"id":"a","title":"A","summary":"body"}]}`)
			deadline := "2s"
			if kind == "timeout" {
				deadline = "30ms"
			}
			args := []string{"selection-watch", "--request", request, "--response", response, "--runner-handle", "test-child", "--timeout", deadline, "--json"}
			// Producer waits for the watcher's explicit readiness marker before publishing.
			if kind != "timeout" && kind != "preexisting" {
				go func() {
					for i := 0; i < 200; i++ {
						if _, err := os.Stat(response + ".ready"); err == nil {
							break
						}
						time.Sleep(time.Millisecond)
					}
					raw := `{"selected":["a"],"assessments":[{"id":"a","relation":"direct","reason":"matches","need_to_read":"confirm conditions"}],"unknown_action":"read missing evidence"}`
					if kind == "invalid" {
						raw = `{"selected":["foreign"]}`
					}
					os.WriteFile(response+".tmp", []byte(raw), 0600)
					os.Rename(response+".tmp", response)
				}()
			}
			if kind == "preexisting" {
				writeFile(t, response, `{}`)
			}
			if kind == "usage-zero" || kind == "usage-invalid" {
				raw := `{"input_tokens":0,"output_tokens":0}`
				if kind == "usage-invalid" {
					raw = `{"input_tokens":-1,"output_tokens":2}`
				}
				writeFile(t, usage, raw)
				args = append(args, "--usage-file", usage)
			}
			var out, errout bytes.Buffer
			exit := runRuntime(args, &out, &errout)
			var got map[string]any
			if err := json.Unmarshal(out.Bytes(), &got); err != nil {
				t.Fatalf("missing receipt: exit=%d %s %s", exit, out.String(), errout.String())
			}
			receipt, ok := got["selection_watch"].(map[string]any)
			if !ok {
				t.Fatalf("missing watcher receipt %s", out.String())
			}
			if kind == "ready" || kind == "usage-zero" {
				if exit != 0 || receipt["state"] != "response_ready" || receipt["interrupt_required"] != false {
					t.Fatalf("ready %s", out.String())
				}
				usage := receipt["usage"].(map[string]any)
				want := "unavailable"
				if kind == "usage-zero" {
					want = "available"
				}
				if usage["availability"] != want {
					t.Fatalf("usage %s", out.String())
				}
			} else {
				if exit == 0 || receipt["interrupt_required"] != true {
					t.Fatalf("failure not handed off %s", out.String())
				}
			}
		})
	}
}

func TestSelectionWatchRequestChanged(t *testing.T) {
	dir := t.TempDir()
	request := filepath.Join(dir, "request.json")
	response := filepath.Join(dir, "response.json")
	writeFile(t, request, `{"system":"fixed","query":"probe","candidates":[]}`)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for i := 0; i < 1000; i++ {
			if _, err := os.Stat(response + ".ready"); err == nil {
				os.WriteFile(request, []byte(`{"changed":true}`), 0600)
				return
			}
			time.Sleep(time.Millisecond)
		}
	}()
	var out, errout bytes.Buffer
	exit := runRuntime([]string{"selection-watch", "--request", request, "--response", response, "--runner-handle", "test-child", "--timeout", "2s", "--json"}, &out, &errout)
	<-done
	var r Result
	if err := json.Unmarshal(out.Bytes(), &r); err != nil {
		t.Fatal(err)
	}
	if exit == 0 || r.SelectionWatch == nil || r.SelectionWatch.State != "request_changed" || !r.SelectionWatch.InterruptRequired {
		t.Fatalf("changed request accepted %s %s", out.String(), errout.String())
	}
}
