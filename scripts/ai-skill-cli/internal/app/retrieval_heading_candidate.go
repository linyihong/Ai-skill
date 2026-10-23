package app

import "strings"

// runtimeFeedbackHeadingFields accepts ATX heading levels from historical
// lesson formats. Summary ends at the next heading; it never borrows body
// text from a later section. Fenced examples cannot become document fields.
func runtimeFeedbackHeadingFields(content string) (title, summary string) {
	fenced := false
	fence := ""
	inSummary := false
	lines := []string{}
	for _, raw := range strings.Split(content, "\n") {
		line := strings.TrimSpace(raw)
		if strings.HasPrefix(line, "```") || strings.HasPrefix(line, "~~~") {
			marker := line[:3]
			if !fenced {
				fenced = true
				fence = marker
			} else if marker == fence {
				fenced = false
			}
			continue
		}
		if fenced {
			continue
		}
		level := 0
		for level < len(line) && line[level] == '#' {
			level++
		}
		heading := level > 0 && level <= 6 && level < len(line) && line[level] == ' '
		if heading {
			if inSummary {
				break
			}
			text := strings.TrimSpace(strings.TrimRight(strings.TrimSpace(line[level:]), "#"))
			if title == "" {
				title = text
			}
			if strings.EqualFold(text, "One-line Summary") {
				inSummary = true
			}
			continue
		}
		if inSummary {
			if line == "" {
				if len(lines) > 0 {
					break
				}
				continue
			}
			lines = append(lines, line)
		}
	}
	return title, strings.Join(lines, " ")
}
