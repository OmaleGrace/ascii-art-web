package main

import (
	"strings"
)

func RenderLine(text string, tit map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i <= 7; i++ {
		var sb strings.Builder
		for _, r := range text {
			sb.WriteString(tit[r][i])
		}
		result[i] = sb.String()
	}
	return result
}

package main

func RenderLine(text string, banner map[rune][]string) []string {
	word := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for _, r := range text {
			word[i] += banner[r][i]
		}
	}
	return word
}

