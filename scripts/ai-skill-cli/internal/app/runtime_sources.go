package app

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

const runtimeIntelligenceDir = "intelligence/engineering/execution/validation-reasoning"

// RuntimeSourceDetail describes provenance, never adoption or authority.
type RuntimeSourceDetail struct {
	Checksum        string   `json:"checksum"`
	CanonicalStatus string   `json:"canonical_status"`
	Domains         []string `json:"domains"`
	Scope           any      `json:"scope"`
	Supersession    any      `json:"replaces"`
	Freshness       string   `json:"freshness"`
	Applicability   string   `json:"applicability"`
	Readback        string   `json:"readback"`
}

func runtimeCanonicalDetail(path, relative string) (RuntimeSourceDetail, error) {
	d := RuntimeSourceDetail{CanonicalStatus: "unknown", Domains: []string{}, Scope: "unknown", Supersession: "unknown", Freshness: "current", Applicability: "unverified; read canonical source and check conditions", Readback: relative}
	b, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		d.Freshness = "missing"
		return d, nil
	}
	if err != nil {
		return d, err
	}
	d.Checksum = fmt.Sprintf("%x", sha256.Sum256(b))
	text := string(b)
	fields := map[string]any{}
	if strings.HasPrefix(text, "---\n") {
		if end := strings.Index(text[4:], "\n---"); end >= 0 {
			if err = yaml.Unmarshal([]byte(text[4:4+end]), &fields); err != nil {
				return d, err
			}
		}
	} else if strings.HasSuffix(path, ".yaml") {
		if err = yaml.Unmarshal(b, &fields); err != nil {
			return d, err
		}
	}
	if v, ok := fields["status"].(string); ok && v != "" {
		d.CanonicalStatus = v
	} else {
		re := regexp.MustCompile("(?mi)^(?:\\*\\*)?Status(?:\\*\\*)?:\\s*`?([^`\\r\\n]+)")
		if m := re.FindStringSubmatch(text); len(m) > 1 {
			d.CanonicalStatus = strings.TrimSpace(m[1])
		}
	}
	if v, ok := fields["scope"]; ok {
		d.Scope = v
	}
	if v, ok := fields["replaces"]; ok {
		d.Supersession = v
	}
	switch v := fields["domains"].(type) {
	case []any:
		for _, x := range v {
			d.Domains = append(d.Domains, fmt.Sprint(x))
		}
	case string:
		if v != "" {
			d.Domains = append(d.Domains, v)
		}
	}
	return d, nil
}

func runtimeIntelligenceRecords(repo string) ([]runtimeIndexAtom, error) {
	paths, err := filepath.Glob(filepath.Join(repo, filepath.FromSlash(runtimeIntelligenceDir), "*.md"))
	if err != nil {
		return nil, err
	}
	records := []runtimeIndexAtom{}
	for _, path := range paths {
		if filepath.Base(path) == "README.md" {
			continue
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		rel, err := filepath.Rel(repo, path)
		if err != nil {
			return nil, err
		}
		rel = filepath.ToSlash(rel)
		d, err := runtimeCanonicalDetail(path, rel)
		if err != nil {
			return nil, err
		}
		summary := ""
		for _, paragraph := range strings.Split(string(b), "\n\n") {
			paragraph = strings.TrimSpace(paragraph)
			if paragraph != "" && !strings.HasPrefix(paragraph, "#") && !strings.HasPrefix(paragraph, "**Status") && !strings.HasPrefix(paragraph, "---") {
				summary = paragraph
				break
			}
		}
		records = append(records, runtimeIndexAtom{ID: "source:" + rel, SourcePath: rel, Layer: "intelligence", Type: "intelligence", Status: d.CanonicalStatus, Priority: "P2", Confidence: "unknown", ContextCost: "unknown", Title: runtimeTitleFromMarkdown(string(b)), Summary: summary, Body: string(b)})
	}
	return records, nil
}

// Snapshot the effective generator inputs, including membership. No timestamps:
// rebuilds with identical canonical inputs have identical fingerprints.
func runtimeIndexInputSnapshot(repo string) (map[string]string, error) {
	records, err := runtimeIndexRecords(repo)
	if err != nil {
		return nil, err
	}
	inputs := map[string]string{"generator:version": "source-body-v1"}
	for _, r := range records {
		b, err := json.Marshal(r)
		if err != nil {
			return nil, err
		}
		inputs["atom:"+r.ID] = fmt.Sprintf("%x", sha256.Sum256(b))
		if r.SourcePath != "" {
			path := filepath.Join(repo, filepath.FromSlash(r.SourcePath))
			b, err = os.ReadFile(path)
			if err != nil {
				if os.IsNotExist(err) {
					inputs["source:"+r.SourcePath] = "missing"
					continue
				}
				info, e := os.Stat(path)
				if e == nil && info.IsDir() {
					inputs["source:"+r.SourcePath] = "directory"
					continue
				}
				return nil, err
			}
			inputs["source:"+r.SourcePath] = fmt.Sprintf("%x", sha256.Sum256(b))
		}
	}
	for _, glob := range []string{"knowledge/graphs/*.yaml", "knowledge/glossary/*.md"} {
		paths, err := filepath.Glob(filepath.Join(repo, filepath.FromSlash(glob)))
		if err != nil {
			return nil, err
		}
		for _, path := range paths {
			b, err := os.ReadFile(path)
			if err != nil {
				return nil, err
			}
			rel, _ := filepath.Rel(repo, path)
			inputs["file:"+filepath.ToSlash(rel)] = fmt.Sprintf("%x", sha256.Sum256(b))
		}
	}
	return inputs, nil
}

func runtimeIndexFreshness(repo, path string) error {
	current, err := runtimeIndexInputSnapshot(repo)
	if err != nil {
		return err
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Query("SELECT source_path, checksum FROM index_inputs")
	if err != nil {
		return fmt.Errorf("index input manifest unavailable; rebuild: %w", err)
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var p, s string
		if err = rows.Scan(&p, &s); err != nil {
			return err
		}
		if current[p] != s {
			return fmt.Errorf("stale index input: %s", p)
		}
		count++
	}
	if err = rows.Err(); err != nil {
		return err
	}
	if count != len(current) {
		return fmt.Errorf("index input membership changed: indexed=%d current=%d", count, len(current))
	}
	return nil
}

func runtimeSourceDetail(db *sql.DB, path string) (*RuntimeSourceDetail, error) {
	var raw string
	if err := db.QueryRow("SELECT metadata FROM source_metadata WHERE source_path=?", path).Scan(&raw); err != nil {
		if err == sql.ErrNoRows {
			return &RuntimeSourceDetail{CanonicalStatus: "unknown", Domains: []string{}, Scope: "unknown", Supersession: "unknown", Freshness: "missing", Applicability: "unverified", Readback: path}, nil
		}
		return nil, err
	}
	var d RuntimeSourceDetail
	if err := json.Unmarshal([]byte(raw), &d); err != nil {
		return nil, err
	}
	return &d, nil
}
