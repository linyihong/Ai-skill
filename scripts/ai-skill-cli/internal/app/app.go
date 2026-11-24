package app

import (
	"flag"
	"fmt"
	"io"
)

func Run(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		_, _ = fmt.Fprintln(stderr, "usage: ai-skill <command> [flags]")
		return ExitInvalidUsage
	}

	switch args[0] {
	case "help", "-h", "--help":
		printUsage(stdout)
		return ExitSuccess
	}
	if c, ok := findCommand(args[0]); ok {
		return c.Run(args[1:], stdout, stderr)
	}
	_, _ = fmt.Fprintf(stderr, "unknown command: %s\n", args[0])
	return ExitInvalidUsage
}

func printUsage(w io.Writer) {
	_, _ = fmt.Fprintln(w, "usage: ai-skill <command> [flags]")
	_, _ = fmt.Fprintln(w, "")
	_, _ = fmt.Fprintln(w, "commands:")
	for _, c := range cliCommands {
		_, _ = fmt.Fprintf(w, "  %s    %s\n", c.Name, c.Summary)
	}
}

func newFlagSet(name string, stderr io.Writer) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(stderr)
	return fs
}
