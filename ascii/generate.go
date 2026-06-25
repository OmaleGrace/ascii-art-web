package ascii

import (
	"strings"
)

func GenerateArt(s string, banner map[rune][]string) string {
	if s == "" {
		return ""
	}

	line := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")

	for _, rn := range s {
		if (rn < 32 || rn > 126) && rn != '\n' && rn != '\r' {
			return "Error: Invalid Character"
		}
	}

	var result strings.Builder
	for _, st := range line {
		if st == "" {
			result.WriteString("\n")
		} else {
			rend := RenderLine(st, banner)
			for _, l := range rend {
				result.WriteString(l)
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
