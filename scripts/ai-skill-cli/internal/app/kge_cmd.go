package app

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
)

// runKge dispatches `ai-skill kge <check|validate|diagnose>`.
// Presentation follows plan D9 Adapter Presentation Policy.
func runKge(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		_, _ = fmt.Fprintln(stderr, "usage: ai-skill kge <check|validate|diagnose> [--root PATH] [--advisory]")
		return ExitInvalidUsage
	}
	cmd := args[0]
	rest := args[1:]
	root := ""
	advisory := false
	for i := 0; i < len(rest); i++ {
		switch rest[i] {
		case "--root":
			if i+1 >= len(rest) {
				_, _ = fmt.Fprintln(stderr, "--root requires a path")
				return ExitInvalidUsage
			}
			root = rest[i+1]
			i++
		case "--advisory":
			advisory = true
		case "-h", "--help":
			_, _ = fmt.Fprintln(stdout, "usage: ai-skill kge <check|validate|diagnose> [--root PATH] [--advisory]")
			_, _ = fmt.Fprintln(stdout, "  check     validation + advisory summary (push watershed; advisory does not fail)")
			_, _ = fmt.Fprintln(stdout, "  validate  validation only; add --advisory for full advisory list")
			_, _ = fmt.Fprintln(stdout, "  diagnose  IDE/MCP JSON diagnostics (full findings, all severities)")
			return ExitSuccess
		default:
			_, _ = fmt.Fprintf(stderr, "unknown flag or arg: %s\n", rest[i])
			return ExitInvalidUsage
		}
	}
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "getwd: %v\n", err)
			return ExitValidationFailed
		}
	}
	ctx, err := buildKGEWorkspaceContext(root)
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "kge context: %v\n", err)
		return ExitValidationFailed
	}
	eng := kge.NewEngine(kge.DefaultRules()...)
	findings := eng.RunAvailable(ctx)

	switch cmd {
	case "check":
		_, _ = fmt.Fprintln(stdout, kge.FormatCheckReport(findings, 3))
		if kge.Blocking(findings) {
			return ExitValidationFailed
		}
		return ExitSuccess
	case "validate":
		errs, _, _ := kge.Partition(findings)
		if len(errs) > 0 {
			_, _ = fmt.Fprintln(stdout, "Validation failed:")
			for _, f := range errs {
				_, _ = fmt.Fprintf(stdout, "  - %s\n", f.Message)
			}
			if advisory {
				_, _ = fmt.Fprintln(stdout, kge.FormatAdvisoryFull(findings))
			}
			return ExitValidationFailed
		}
		_, _ = fmt.Fprintln(stdout, "Validation ok.")
		if advisory {
			_, _ = fmt.Fprintln(stdout, kge.FormatAdvisoryFull(findings))
		}
		return ExitSuccess
	case "diagnose":
		payload, encErr := kge.FormatIDEDiagnosticsJSON(findings)
		if encErr != nil {
			_, _ = fmt.Fprintf(stderr, "encode diagnostics: %v\n", encErr)
			return ExitGeneralFailure
		}
		_, _ = fmt.Fprintln(stdout, payload)
		if kge.Blocking(findings) {
			return ExitValidationFailed
		}
		return ExitSuccess
	default:
		_, _ = fmt.Fprintf(stderr, "unknown kge subcommand: %s\n", cmd)
		return ExitInvalidUsage
	}
}

func buildKGEWorkspaceContext(root string) (kge.Context, error) {
	staged, err := gitLines(root, "diff", "--cached", "--name-only", "--diff-filter=ACMR")
	if err != nil || len(staged) == 0 {
		// Fallback: unstaged changed files for local check convenience
		if unstaged, err2 := gitLines(root, "diff", "--name-only", "--diff-filter=ACMR"); err2 == nil {
			staged = unstaged
		}
	}
	for i := range staged {
		staged[i] = filepath.ToSlash(staged[i])
	}
	contents := map[string]string{}
	for _, p := range staged {
		if !strings.HasSuffix(strings.ToLower(p), ".md") && !strings.HasSuffix(strings.ToLower(p), ".markdown") {
			continue
		}
		full := filepath.Join(root, p)
		body, readErr := os.ReadFile(full)
		if readErr != nil {
			continue
		}
		contents[filepath.ToSlash(p)] = string(body)
	}
	diff := stagedDiffCached(root, "scripts/ai-skill-cli/internal/app/")
	provided := map[kge.CapabilityID]bool{
		kge.CapStagedPaths: true,
		kge.CapCommitMsg:   true,
		kge.CapStagedDiff:  true,
	}
	if len(contents) > 0 {
		provided[kge.CapStagedContent] = true
	}
	return kge.Context{
		RepoRoot:     root,
		CommitMsg:    "",
		StagedPaths:  staged,
		StagedDiff:   diff,
		FileContents: contents,
		Provided:     provided,
	}, nil
}
