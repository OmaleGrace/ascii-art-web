package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	w := SplitInput(input)
	for _, q := range w {
		if q != "" {
			break
		}
	}
	var result strings.Builder
	for _, g := range w {
		if g == "" {
			result.WriteString("\n")
		} else {
			rend := RenderLine(g, banner)
			for _, l := range rend {
				result.WriteString(l + "\n")
			}
		}
	}
	return result.String()
}
