package app

import "github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"

// Link is an inline markdown link (compat wrapper around portable kge.Link).
type Link struct {
	Target string
	Line   int
	Column int
}

// extractMarkdownLinks delegates to the portable KGE markdown link parser.
func extractMarkdownLinks(content []byte) []Link {
	src := kge.ExtractMarkdownLinks(content)
	if len(src) == 0 {
		return nil
	}
	out := make([]Link, len(src))
	for i, l := range src {
		out[i] = Link{Target: l.Target, Line: l.Line, Column: l.Column}
	}
	return out
}

func shouldKeepLinkTarget(t string) bool {
	return kge.ShouldKeepLinkTarget(t)
}

func stripLinkFragment(target string) string {
	return kge.StripLinkFragment(target)
}

func isLinkTargetContext(line string, pos int) bool {
	return kge.IsLinkTargetContext(line, pos)
}
