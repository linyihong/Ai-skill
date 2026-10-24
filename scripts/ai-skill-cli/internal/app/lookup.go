package app

import (
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

//go:embed lookup_selector_prompt.txt
var lookupSelectorPrompt string

type LookupCandidate struct {
	ID     string      `json:"id"`
	Title  string      `json:"title"`
	Result QueryResult `json:"result"`
}
type LookupBundle struct {
	Version    int               `json:"version"`
	Repo       string            `json:"repo"`
	Index      string            `json:"index"`
	IndexSHA   string            `json:"index_sha256"`
	BinarySHA  string            `json:"binary_sha256"`
	Query      string            `json:"query"`
	Layer      string            `json:"layer"`
	Status     string            `json:"status"`
	RequestSHA string            `json:"request_sha256"`
	Candidates []LookupCandidate `json:"candidates"`
}
type LookupReceipt struct {
	Phase        string                 `json:"phase"`
	BundlePath   string                 `json:"bundle_path,omitempty"`
	BundleSHA    string                 `json:"bundle_sha256"`
	RequestPath  string                 `json:"request_path,omitempty"`
	RequestSHA   string                 `json:"request_sha256"`
	ResponseSHA  string                 `json:"response_sha256,omitempty"`
	Payload      *retrievalBlindPayload `json:"payload,omitempty"`
	Deferred     []string               `json:"deferred_sources,omitempty"`
	SourceChecks []string               `json:"source_checks,omitempty"`
	Usage        *SelectionUsage        `json:"usage,omitempty"`
}

func lookupSHA(b []byte) string { return fmt.Sprintf("%x", sha256.Sum256(b)) }
func lookupFileSHA(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return lookupSHA(b), nil
}
func runtimeCandidateExpression(query string) string {
	words := regexp.MustCompile(`[\p{L}\p{N}_]+`).FindAllString(query, -1)
	terms := []string{}
	seen := map[string]bool{}
	for _, w := range words {
		key := strings.ToLower(w)
		if seen[key] {
			continue
		}
		seen[key] = true
		terms = append(terms, runtimeFTSMatchLiteral(w))
		if len(terms) == 64 {
			break
		}
	}
	if len(terms) == 0 {
		return runtimeFTSMatchLiteral(query)
	}
	return strings.Join(terms, " OR ")
}
func lookupRequest(bundle LookupBundle) selectionRequest {
	r := selectionRequest{System: lookupSelectorPrompt, Query: bundle.Query}
	r.Candidates = make([]struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Summary string `json:"summary"`
	}, 0, len(bundle.Candidates))
	for _, c := range bundle.Candidates {
		summary := c.Result.Summary
		if strings.TrimSpace(summary) == "" {
			summary = "Indexed summary unavailable; read the source to determine relevance."
		}
		r.Candidates = append(r.Candidates, struct {
			ID      string `json:"id"`
			Title   string `json:"title"`
			Summary string `json:"summary"`
		}{c.ID, c.Title, summary})
	}
	return r
}
func lookupWriteNew(path string, v any) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	b = append(b, '\n')
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}
	_, writeErr := f.Write(b)
	closeErr := f.Close()
	if writeErr != nil {
		return "", writeErr
	}
	return lookupSHA(b), closeErr
}
func runLookup(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: runtime lookup <index|prepare|complete> [flags]")
		return ExitInvalidUsage
	}
	phase := args[0]
	fs := newFlagSet("runtime lookup "+phase, stderr)
	var repo, db, query, layer, status, out, bundlePath, bundleSHA, response, receipt string
	var jsonOut bool
	fs.StringVar(&repo, "repo", ".", "canonical repo root")
	fs.StringVar(&db, "db", "", "derived index path")
	fs.StringVar(&query, "keyword", "", "task query")
	fs.StringVar(&layer, "layer", "", "candidate layer")
	fs.StringVar(&status, "status", "", "canonical status, retaining unknown")
	fs.StringVar(&out, "out", "", "new attempt directory")
	fs.StringVar(&bundlePath, "bundle", "", "prepare bundle.json")
	fs.StringVar(&bundleSHA, "bundle-sha", "", "externally captured prepare bundle SHA")
	fs.StringVar(&response, "response", "", "raw selector response")
	fs.StringVar(&receipt, "receipt", "", "successful production watcher receipt")
	fs.BoolVar(&jsonOut, "json", false, "JSON output")
	if err := fs.Parse(args[1:]); err != nil {
		return ExitInvalidUsage
	}
	r := Result{Command: "runtime lookup " + phase, Mode: "native", Status: "success", ExitCode: ExitSuccess, Checks: []Check{}, PlannedActions: []string{}, Mutations: []string{}}
	var err error
	switch phase {
	case "index":
		repo, err = filepath.Abs(repo)
		if err == nil {
			if strings.TrimSpace(db) == "" {
				err = fmt.Errorf("--db required")
			} else {
				db = runtimeQueryDBPath(repo, db)
				err = buildNativeRuntimeSQLiteIndex(repo, db)
				if err == nil {
					r.Mutations = append(r.Mutations, db)
				}
			}
		}
	case "prepare":
		r.Lookup, err = lookupPrepare(repo, db, query, layer, status, out)
	case "complete":
		r.Lookup, err = lookupComplete(bundlePath, bundleSHA, response, receipt)
	default:
		err = fmt.Errorf("unsupported lookup mode %q", phase)
	}
	if len(fs.Args()) > 0 {
		err = fmt.Errorf("unexpected positional arguments")
	}
	if err != nil {
		r.Lookup = nil
		r.Status = "blocked"
		r.ExitCode = ExitValidationFailed
		r.Error = &CommandError{Code: "lookup_" + phase + "_failed", Message: err.Error(), Remediation: "Preserve the failed attempt; resolve input/freshness/runner failure explicitly. No fallback or silent retry."}
	}
	if jsonOut {
		_ = writeJSON(stdout, r)
	} else {
		_ = writePlain(stdout, r)
	}
	return r.ExitCode
}
func lookupPrepare(repo, db, query, layer, status, out string) (*LookupReceipt, error) {
	if strings.TrimSpace(query) == "" || out == "" {
		return nil, fmt.Errorf("--keyword and new --out required")
	}
	repo, err := filepath.Abs(repo)
	if err != nil {
		return nil, err
	}
	db = runtimeQueryDBPath(repo, db)
	q := buildRuntimeQueryResult(runtimeOptions{repoPath: repo, dbPath: db, keyword: query, layer: layer, statusFilter: status, sourceLevel: true, candidateTerms: true, limit: 64})
	if q.ExitCode != 0 {
		return nil, fmt.Errorf("candidate query: %s", q.Error.Message)
	}
	exe, err := os.Executable()
	if err != nil {
		return nil, err
	}
	binarySHA, err := lookupFileSHA(exe)
	if err != nil {
		return nil, err
	}
	indexSHA, err := lookupFileSHA(db)
	if err != nil {
		return nil, err
	}
	b := LookupBundle{Version: 1, Repo: repo, Index: db, IndexSHA: indexSHA, BinarySHA: binarySHA, Query: query, Layer: layer, Status: status, Candidates: []LookupCandidate{}}
	sqlDB, err := sql.Open("sqlite", db)
	if err != nil {
		return nil, err
	}
	defer sqlDB.Close()
	for _, r := range q.Results {
		var title string
		if err = sqlDB.QueryRow("SELECT title FROM atoms WHERE id=?", r.ID).Scan(&title); err != nil {
			return nil, err
		}
		if strings.TrimSpace(title) == "" {
			title = filepath.Base(r.SourcePath)
		}
		id := lookupSHA([]byte("lookup-v1:" + query + ":" + r.SourcePath))[:16]
		b.Candidates = append(b.Candidates, LookupCandidate{ID: id, Title: title, Result: r})
	}
	sort.Slice(b.Candidates, func(i, j int) bool { return b.Candidates[i].ID < b.Candidates[j].ID })
	if err = runtimeIndexFreshness(repo, db); err != nil {
		return nil, err
	}
	out, err = filepath.Abs(out)
	if err != nil {
		return nil, err
	}
	if err = os.Mkdir(out, 0700); err != nil {
		return nil, err
	}
	success := false
	defer func() {
		if !success {
			os.RemoveAll(out)
		}
	}()
	requestPath := filepath.Join(out, "request.json")
	requestSHA, err := lookupWriteNew(requestPath, lookupRequest(b))
	if err != nil {
		return nil, err
	}
	b.RequestSHA = requestSHA
	bundlePath := filepath.Join(out, "bundle.json")
	bundleSHA, err := lookupWriteNew(bundlePath, b)
	if err != nil {
		return nil, err
	}
	success = true
	return &LookupReceipt{Phase: "prepared", BundlePath: bundlePath, BundleSHA: bundleSHA, RequestPath: requestPath, RequestSHA: requestSHA}, nil
}
func lookupComplete(bundlePath, bundleSHA, responsePath, receiptPath string) (*LookupReceipt, error) {
	if bundleSHA == "" {
		return nil, fmt.Errorf("--bundle-sha required")
	}
	raw, err := selectionReadBounded(bundlePath)
	if err != nil {
		return nil, err
	}
	if lookupSHA(raw) != bundleSHA {
		return nil, fmt.Errorf("bundle SHA mismatch")
	}
	var b LookupBundle
	if err = selectionDecode(raw, &b); err != nil {
		return nil, err
	}
	if b.Version != 1 {
		return nil, fmt.Errorf("unsupported bundle")
	}
	exe, err := os.Executable()
	if err != nil {
		return nil, err
	}
	exeSHA, err := lookupFileSHA(exe)
	if err != nil || exeSHA != b.BinarySHA {
		return nil, fmt.Errorf("executable changed")
	}
	indexSHA, err := lookupFileSHA(b.Index)
	if err != nil || indexSHA != b.IndexSHA {
		return nil, fmt.Errorf("index changed")
	}
	if err = runtimeIndexFreshness(b.Repo, b.Index); err != nil {
		return nil, err
	}
	requestRaw, err := selectionReadBounded(filepath.Join(filepath.Dir(bundlePath), "request.json"))
	if err != nil {
		return nil, err
	}
	expected, _ := json.MarshalIndent(lookupRequest(b), "", "  ")
	expected = append(expected, '\n')
	if !bytesEqual(requestRaw, expected) || lookupSHA(requestRaw) != b.RequestSHA {
		return nil, fmt.Errorf("request does not match bundle")
	}
	responseRaw, err := selectionReadBounded(responsePath)
	if err != nil {
		return nil, err
	}
	receiptRaw, err := selectionReadBounded(receiptPath)
	if err != nil {
		return nil, err
	}
	var watched Result
	if err = selectionDecode(receiptRaw, &watched); err != nil {
		return nil, err
	}
	w := watched.SelectionWatch
	if watched.Command != "runtime selection-watch" || watched.ExitCode != 0 || watched.Status != "success" || w == nil || w.State != "response_ready" || w.InterruptRequired || w.RequestSHA != b.RequestSHA || w.ResponseSHA != lookupSHA(responseRaw) {
		return nil, fmt.Errorf("missing or mismatched successful watcher receipt")
	}
	candidates := map[string]LookupCandidate{}
	ids := []string{}
	for _, c := range b.Candidates {
		if candidates[c.ID].ID != "" {
			return nil, fmt.Errorf("duplicate candidate")
		}
		candidates[c.ID] = c
		ids = append(ids, c.ID)
	}
	selected, err := retrievalParseBlind(responseRaw, ids)
	if err != nil {
		return nil, err
	}
	coverage := retrievalCoverage{Candidates: len(ids), Assessed: len(selected.Assessments), Unassessed: len(ids) - len(selected.Assessments)}
	needs := map[string]string{}
	for _, a := range selected.Assessments {
		needs[a.ID] = a.Need
		switch a.Relation {
		case "direct":
			coverage.Direct++
		case "unknown":
			coverage.ExplicitUnknown++
		case "background":
			coverage.Background++
		case "not-applicable":
			coverage.NotApplicable++
		}
	}
	coverage.Unknown = coverage.Unassessed + coverage.ExplicitUnknown
	switch {
	case coverage.Direct > 0:
		coverage.State = "direct-pending-verification"
	case coverage.Unknown > 0:
		coverage.State = "pending-assessment"
	case coverage.Candidates == 0:
		coverage.State = "no-candidates-retrieved"
	default:
		coverage.State = "no-direct-in-assessed-pool"
	}
	ranked := []retrievalBlindModelResult{}
	checks := []string{}
	for _, id := range selected.Selected {
		c := candidates[id]
		if c.Result.Source == nil || c.Result.Source.Checksum == "" {
			return nil, fmt.Errorf("selected source has no readable version: %s", c.Result.SourcePath)
		}
		sum, err := lookupFileSHA(filepath.Join(b.Repo, filepath.FromSlash(c.Result.SourcePath)))
		if err != nil || sum != c.Result.Source.Checksum {
			return nil, fmt.Errorf("selected source changed: %s", c.Result.SourcePath)
		}
		checks = append(checks, c.Result.SourcePath)
		ranked = append(ranked, retrievalBlindModelResult{Source: c.Result.SourcePath, Kind: c.Result.Type, Status: c.Result.Source.CanonicalStatus, Summary: c.Result.Summary, Need: needs[id]})
	}
	payload, deferred, err := lookupPack("lookup/"+b.RequestSHA[:16], coverage, ranked, 2000)
	if err != nil {
		return nil, err
	}
	if lookupPayloadRunes(payload) > 2000 {
		return nil, fmt.Errorf("complete payload exceeds budget")
	}
	if err = runtimeIndexFreshness(b.Repo, b.Index); err != nil {
		return nil, err
	}
	return &LookupReceipt{Phase: "completed-unverified-applicability", BundleSHA: bundleSHA, RequestSHA: b.RequestSHA, ResponseSHA: lookupSHA(responseRaw), Payload: &payload, Deferred: deferred, SourceChecks: checks, Usage: &w.Usage}, nil
}
func bytesEqual(a, b []byte) bool { return string(a) == string(b) }
