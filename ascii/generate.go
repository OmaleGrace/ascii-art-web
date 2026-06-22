package ascii

import (
	"strings"
)

func GenerateArt(s string, banner map[rune][]string) string {
	if s == "" {
		return ""
	}

	for _, rn := range s {
		if rn < 32 || rn > 126 {
			return "Error: Invalid Character"
		}
	}
	line := strings.ReplaceAll(s, "\r", "")
	sp := strings.Split(line, "\\n")

	for _, ch := range sp {
		if ch != "" {
			break
		}
	}

	var result strings.Builder
	for _, st := range sp {
		if st == "" {
			result.WriteString("\n")
		} else {
			r := RenderLine(s, banner)
			for _, l := range r {
				result.WriteString(l)
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
