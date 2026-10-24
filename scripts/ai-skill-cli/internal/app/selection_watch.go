package app

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"time"
)

type SelectionUsage struct {
	Availability string `json:"availability"`
	InputTokens  *int64 `json:"input_tokens"`
	OutputTokens *int64 `json:"output_tokens"`
}
type SelectionWatchReceipt struct {
	State              string                  `json:"state"`
	RunnerHandle       string                  `json:"runner_handle"`
	RequestSHA         string                  `json:"request_sha256"`
	ResponseSHA        string                  `json:"response_sha256,omitempty"`
	ElapsedMS          int64                   `json:"elapsed_ms"`
	InterruptRequired  bool                    `json:"interrupt_required"`
	InterruptConfirmed bool                    `json:"interrupt_confirmed"`
	Usage              SelectionUsage          `json:"usage"`
	Selection          *retrievalBlindResponse `json:"selection,omitempty"`
}
type selectionRequest struct {
	System     string `json:"system"`
	Query      string `json:"query"`
	Candidates []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Summary string `json:"summary"`
	} `json:"candidates"`
}

func selectionReadBounded(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := io.ReadAll(io.LimitReader(f, 1<<20+1))
	if err != nil {
		return nil, err
	}
	if len(b) > 1<<20 {
		return nil, fmt.Errorf("file exceeds 1 MiB")
	}
	return b, nil
}
func selectionDecode(raw []byte, target any) error {
	d := json.NewDecoder(bytes.NewReader(raw))
	d.DisallowUnknownFields()
	if err := d.Decode(target); err != nil {
		return err
	}
	var extra any
	if d.Decode(&extra) != io.EOF {
		return fmt.Errorf("trailing JSON")
	}
	return nil
}

// This transport observes a child owned by the calling agent. It cannot invoke
// the host's interrupt tool. Every unsuccessful attempt requires that handoff;
// interrupt_confirmed is deliberately never inferred from timeout or a signal.
func runSelectionWatch(args []string, stdout, stderr io.Writer) int {
	fs := newFlagSet("runtime selection-watch", stderr)
	var request, response, handle, usageFile string
	var timeout time.Duration
	var jsonOut bool
	fs.StringVar(&request, "request", "", "fixed label-free request JSON")
	fs.StringVar(&response, "response", "", "new response path, atomically published after .ready appears")
	fs.StringVar(&handle, "runner-handle", "", "opaque adapter-owned child handle")
	fs.StringVar(&usageFile, "usage-file", "", "optional usage JSON supplied by adapter")
	fs.DurationVar(&timeout, "timeout", 2*time.Minute, "bounded observation deadline, at most 10 minutes")
	fs.BoolVar(&jsonOut, "json", false, "JSON receipt")
	if err := fs.Parse(args); err != nil {
		return ExitInvalidUsage
	}
	r := Result{Command: "runtime selection-watch", Mode: "native", Status: "blocked", ExitCode: ExitValidationFailed, Checks: []Check{}, PlannedActions: []string{}, Mutations: []string{}}
	receipt := &SelectionWatchReceipt{State: "invalid_request", RunnerHandle: handle, InterruptRequired: true, Usage: SelectionUsage{Availability: "unavailable"}}
	r.SelectionWatch = receipt
	started := time.Now()
	finish := func(state string, err error) int {
		receipt.State = state
		receipt.ElapsedMS = time.Since(started).Milliseconds()
		if err != nil {
			r.Error = &CommandError{Code: state, Message: err.Error(), Remediation: "Interrupt the adapter-owned child, confirm termination, preserve this receipt; do not retry or use fallback silently."}
		} else {
			r.Status = "success"
			r.ExitCode = ExitSuccess
			receipt.InterruptRequired = false
		}
		if jsonOut {
			_ = writeJSON(stdout, r)
		} else {
			writePlain(stdout, r)
			fmt.Fprintf(stdout, "runner state=%s interrupt_required=%t interrupt_confirmed=false\n", state, receipt.InterruptRequired)
		}
		return r.ExitCode
	}
	if strings.TrimSpace(handle) == "" || request == "" || response == "" || request == response || timeout <= 0 || timeout > 10*time.Minute || len(fs.Args()) != 0 {
		return finish("invalid_request", fmt.Errorf("request, distinct response, runner handle and 0 < timeout <= 10m required"))
	}
	raw, err := selectionReadBounded(request)
	if err != nil {
		return finish("invalid_request", err)
	}
	var input selectionRequest
	if err = selectionDecode(raw, &input); err != nil {
		return finish("invalid_request", err)
	}
	if strings.TrimSpace(input.System) == "" || strings.TrimSpace(input.Query) == "" || input.Candidates == nil || len(input.Candidates) > 256 {
		return finish("invalid_request", fmt.Errorf("system/query and at most 256 candidates required"))
	}
	ids := []string{}
	seen := map[string]bool{}
	for _, c := range input.Candidates {
		if strings.TrimSpace(c.ID) == "" || seen[c.ID] || strings.TrimSpace(c.Title) == "" || strings.TrimSpace(c.Summary) == "" {
			return finish("invalid_request", fmt.Errorf("invalid candidate card"))
		}
		ids = append(ids, c.ID)
		seen[c.ID] = true
	}
	receipt.RequestSHA = fmt.Sprintf("%x", sha256.Sum256(raw))
	if _, err = os.Lstat(response); !os.IsNotExist(err) {
		return finish("preexisting_response", fmt.Errorf("response must not exist at dispatch"))
	}
	ready, err := os.OpenFile(response+".ready", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return finish("readiness_failed", err)
	}
	_, err = ready.WriteString(receipt.RequestSHA + "\n")
	closeErr := ready.Close()
	defer os.Remove(response + ".ready")
	if err != nil {
		return finish("readiness_failed", err)
	}
	if closeErr != nil {
		return finish("readiness_failed", closeErr)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	timer := time.NewTimer(timeout)
	defer timer.Stop()
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return finish("canceled", ctx.Err())
		case <-timer.C:
			return finish("timeout", fmt.Errorf("response deadline exceeded; child termination is not confirmed"))
		case <-ticker.C:
		}
		// Check the monotonic deadline before accepting even an already-published file.
		if time.Since(started) >= timeout {
			return finish("timeout", fmt.Errorf("response deadline exceeded"))
		}
		current, err := selectionReadBounded(request)
		if err != nil || !bytes.Equal(raw, current) {
			return finish("request_changed", fmt.Errorf("fixed request changed during observation"))
		}
		responseRaw, err := selectionReadBounded(response)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return finish("response_failed", err)
		}
		receipt.ResponseSHA = fmt.Sprintf("%x", sha256.Sum256(responseRaw))
		selection, err := retrievalParseBlind(responseRaw, ids)
		if err != nil {
			return finish("invalid_output", err)
		}
		if usageFile != "" {
			usageRaw, err := selectionReadBounded(usageFile)
			if err != nil {
				return finish("usage_failed", err)
			}
			var tokens struct {
				Input  *int64 `json:"input_tokens"`
				Output *int64 `json:"output_tokens"`
			}
			if err = selectionDecode(usageRaw, &tokens); err != nil || tokens.Input == nil || tokens.Output == nil || *tokens.Input < 0 || *tokens.Output < 0 {
				return finish("usage_failed", fmt.Errorf("usage requires nonnegative explicit input_tokens/output_tokens"))
			}
			receipt.Usage = SelectionUsage{Availability: "available", InputTokens: tokens.Input, OutputTokens: tokens.Output}
		}
		if time.Since(started) >= timeout {
			return finish("timeout", fmt.Errorf("validation exceeded the observation deadline"))
		}
		current, err = selectionReadBounded(request)
		if err != nil || !bytes.Equal(raw, current) {
			return finish("request_changed", fmt.Errorf("request changed before admission"))
		}
		receipt.Selection = &selection
		return finish("response_ready", nil)
	}
}
