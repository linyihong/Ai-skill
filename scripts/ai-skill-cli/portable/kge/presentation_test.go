package kge

import (
	"strings"
	"testing"
)

func TestDocumentSizingRule_Hard(t *testing.T) {
	body := strings.Repeat("line\n", 320)
	eng := NewEngine(DocumentSizingRule{})
	ctx := Context{
		FileContents: map[string]string{"docs/big.md": body},
		Provided:     map[CapabilityID]bool{CapStagedContent: true},
	}
	f := eng.Run(ctx)
	if len(f) != 1 || f[0].Code != "document_sizing_hard" || f[0].Severity != SeverityWarning {
		t.Fatalf("want hard advisory, got %#v", f)
	}
	if Blocking(f) {
		t.Fatal("document sizing must not block")
	}
}

func TestFormatCheckReport_Summary(t *testing.T) {
	findings := []Finding{
		{RuleID: "rule.document_sizing", Severity: SeverityWarning, Message: "too long"},
		{RuleID: "rule.x", Severity: SeverityWarning, Message: "b"},
		{RuleID: "rule.y", Severity: SeverityWarning, Message: "c"},
		{RuleID: "rule.z", Severity: SeverityWarning, Message: "d"},
	}
	out := FormatCheckReport(findings, 3)
	if !strings.Contains(out, "Ready to push") || !strings.Contains(out, "4 recommendation") || !strings.Contains(out, "and 1 more") {
		t.Fatalf("unexpected check report:\n%s", out)
	}
}

func TestFormatCommitSummary(t *testing.T) {
	s := FormatCommitSummary(2, "ai-skill kge validate --advisory")
	if !strings.Contains(s, "2 advisory") || !strings.Contains(s, "validate --advisory") {
		t.Fatalf("got %q", s)
	}
}

func TestFormatIDEDiagnosticsJSON(t *testing.T) {
	findings := []Finding{
		{RuleID: "rule.document_sizing", Severity: SeverityWarning, Path: "docs/big.md", Code: "document_sizing_hard", Message: "too long"},
		{RuleID: "rule.x", Severity: SeverityError, Message: "blocked"},
	}
	out, err := FormatIDEDiagnosticsJSON(findings)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, `"path": "docs/big.md"`) || !strings.Contains(out, `"severity": "error"`) {
		t.Fatalf("unexpected diagnostics JSON:\n%s", out)
	}
}

func TestRunAvailable_SkipsMissingCaps(t *testing.T) {
	eng := NewEngine(CognitiveCostRule{}, DocumentSizingRule{})
	ctx := Context{
		FileContents: map[string]string{"a.md": strings.Repeat("x\n", 200)},
		Provided:     map[CapabilityID]bool{CapStagedContent: true},
	}
	f := eng.RunAvailable(ctx)
	if len(f) != 1 || f[0].RuleID != "rule.document_sizing" {
		t.Fatalf("want only document_sizing, got %#v", f)
	}
}
