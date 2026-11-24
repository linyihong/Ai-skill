package app

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// commandContractRowRE matches a command-contract table row and captures the
// command text inside its first backticks, e.g. "ai-skill hooks run stop".
var commandContractRowRE = regexp.MustCompile("(?m)^\\| `(ai-skill [^`]*)`")

var braceGroupRE = regexp.MustCompile(`\{([^}]*)\}`)

// contractRowCommands expands each row into command word sequences: drops
// placeholders (<file>, [--flag], --flag) and expands {a,b,c} into one entry
// per alternative.
func contractRowCommands(doc string) [][]string {
	var out [][]string
	for _, m := range commandContractRowRE.FindAllStringSubmatch(doc, -1) {
		text := strings.TrimSpace(strings.TrimPrefix(m[1], "ai-skill "))
		variants := []string{text}
		if bm := braceGroupRE.FindStringSubmatchIndex(text); bm != nil {
			variants = nil
			for _, alt := range strings.Split(text[bm[2]:bm[3]], ",") {
				variants = append(variants, text[:bm[0]]+strings.TrimSpace(alt)+text[bm[1]:])
			}
		}
		for _, v := range variants {
			var words []string
			for _, w := range strings.Fields(v) {
				if strings.HasPrefix(w, "<") || strings.HasPrefix(w, "[") || strings.HasPrefix(w, "-") {
					break
				}
				words = append(words, w)
			}
			if len(words) > 0 {
				out = append(out, words)
			}
		}
	}
	return out
}

func readCommandContract(t *testing.T) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", "docs", "command-contract.md"))
	if err != nil {
		t.Fatalf("read command-contract: %v", err)
	}
	return string(data)
}

// TestCommandContractMatchesCommandTable is the state-based replacement for
// the diff-pattern cli_doc_sync lock: the declarative command table and
// docs/command-contract.md must describe the same command set, in both
// directions, regardless of how a change is written.
func TestCommandContractMatchesCommandTable(t *testing.T) {
	rows := contractRowCommands(readCommandContract(t))
	documented := map[string]bool{}
	for _, words := range rows {
		for i := 1; i <= len(words); i++ {
			documented[strings.Join(words[:i], " ")] = true
		}
	}

	// Table → contract: every command and every subcommand has a row.
	for _, c := range cliCommands {
		if len(c.Subcommands) == 0 {
			if !documented[c.Name] {
				t.Errorf("command %q is dispatchable but has no command-contract row", c.Name)
			}
			continue
		}
		for _, s := range c.Subcommands {
			if !documented[c.Name+" "+s] {
				t.Errorf("subcommand %q is dispatchable but has no command-contract row", c.Name+" "+s)
			}
		}
	}

	// Contract → table: every row names a dispatchable command.
	for _, words := range rows {
		c, ok := findCommand(words[0])
		if !ok {
			t.Errorf("command-contract row %q names an unknown command", strings.Join(words, " "))
			continue
		}
		if len(c.Subcommands) == 0 {
			continue
		}
		matched := false
		for _, s := range c.Subcommands {
			sw := strings.Fields(s)
			if len(words)-1 >= len(sw) && strings.Join(words[1:1+len(sw)], " ") == s {
				matched = true
				break
			}
		}
		if !matched {
			t.Errorf("command-contract row %q does not match any declared subcommand of %q", strings.Join(words, " "), c.Name)
		}
	}
}

// TestContractRowParsing covers the row parser on both directions of the
// shapes the contract uses, so a parser regression cannot make the contract
// test pass vacuously.
func TestContractRowParsing(t *testing.T) {
	doc := strings.Join([]string{
		"| `ai-skill scan-checkboxes <file>` | x |",
		"| `ai-skill kge validate [--advisory]` | x |",
		"| `ai-skill router proposals {list,gc}` | x |",
		"| `ai-skill hooks run commit-msg` validator `validateX` | x |",
		"prose mentioning `ai-skill fake thing` is not a row",
	}, "\n")
	got := contractRowCommands(doc)
	want := []string{"scan-checkboxes", "kge validate", "router proposals list", "router proposals gc", "hooks run commit-msg"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i, w := range want {
		if strings.Join(got[i], " ") != w {
			t.Fatalf("row %d: got %q, want %q", i, strings.Join(got[i], " "), w)
		}
	}
}

// TestCommandTableGatesDispatch checks the table is load-bearing: an
// undeclared command or subcommand is rejected with ExitInvalidUsage before
// any group logic runs.
func TestCommandTableGatesDispatch(t *testing.T) {
	cases := [][]string{
		{"no-such-command"},
		{"runtime", "no-such-sub"},
		{"hooks", "run", "no-such-hook"},
		{"hooks", "no-such-sub"},
		{"enforcement", "no-such-sub"},
		{"plans", "no-such-sub"},
		{"image-processing", "no-such-sub"},
		{"kge", "no-such-sub"},
		{"roo", "no-such-sub"},
		{"copilot", "no-such-sub"},
		{"glossary", "no-such-sub"},
	}
	for _, args := range cases {
		var stdout, stderr bytes.Buffer
		if code := Run(args, &stdout, &stderr); code != ExitInvalidUsage {
			t.Errorf("%v: exit %d, want ExitInvalidUsage (stderr=%q)", args, code, stderr.String())
		}
	}
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"help"}, &stdout, &stderr); code != ExitSuccess {
		t.Fatalf("help exit %d", code)
	}
	for _, c := range cliCommands {
		if !strings.Contains(stdout.String(), "  "+c.Name+"    ") {
			t.Errorf("help output missing command %q", c.Name)
		}
	}
}

// TestStandaloneBinariesInCatalog closes T3's cmd/ scope gap: every
// standalone binary under cmd/ (other than the ai-skill entry itself) must be
// registered in the tool-tree capability catalog.
func TestStandaloneBinariesInCatalog(t *testing.T) {
	readme, err := os.ReadFile(filepath.Join("..", "..", "README.md"))
	if err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(filepath.Join("..", "..", "cmd"))
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		if !e.IsDir() || e.Name() == "ai-skill" {
			continue
		}
		if !strings.Contains(string(readme), "`go run ./cmd/"+e.Name()+"`") {
			t.Errorf("cmd/%s is a standalone binary but has no capability catalog row in scripts/ai-skill-cli/README.md", e.Name())
		}
	}
}
