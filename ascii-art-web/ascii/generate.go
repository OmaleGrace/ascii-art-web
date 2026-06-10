package ascii

import (
	"strings"
)

func GenerateArt(s string, ban map[rune][]string) string {
	if s == "" {
		return ""
	}

	for _, l := range s {
		if l < 32 || l > 126 {
			return ""
		}
	}

	g := SplitInput(s)
	for _, p := range g {
		if p != "" {
			break
		}
	}

	var result strings.Builder
	for _, l := range g {
		if l != "" {
			result.WriteString("\n")
		} else {
			a := RenderLine(l, ban)
			for _, at := range a {
				result.WriteString(at + "\n")
			}
		}
	}
	return result.String()
}