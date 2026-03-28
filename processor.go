package main

import (
	"strconv"
	"strings"
)

// PROCESS TEXT: base conversion + case transforms
func processText(text string) string {
	words := strings.Fields(text)
	var result []string

	for i := 0; i < len(words); i++ {
		w := words[i]

		// Handle base conversions first
		if w == "(hex)" && i > 0 {
			val := words[i-1]
			if dec, err := strconv.ParseInt(val, 16, 64); err == nil {
				result[len(result)-1] = strconv.FormatInt(dec, 10)
			}
			continue
		}

		if w == "(bin)" && i > 0 {
			val := words[i-1]
			if dec, err := strconv.ParseInt(val, 2, 64); err == nil {
				result[len(result)-1] = strconv.FormatInt(dec, 10)
			}
			continue
		}

		// Handle case transformations
		if strings.HasPrefix(w, "(") {
			marker := w
			for !strings.HasSuffix(marker, ")") && i+1 < len(words) {
				i++
				marker += " " + words[i]
			}
			marker = strings.TrimSuffix(strings.TrimPrefix(marker, "("), ")")
			parts := strings.Split(marker, ",")
			action := strings.TrimSpace(parts[0])
			count := 1
			if len(parts) == 2 {
				if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					count = n
				}
			}
			for j := len(result) - count; j < len(result); j++ {
				if j >= 0 {
					switch action {
					case "up":
						result[j] = strings.ToUpper(result[j])
					case "low":
						result[j] = strings.ToLower(result[j])
					case "cap":
						result[j] = capitalizeWord(result[j])
					}
				}
			}
			continue
		}

		// Regular words → append to result
		result = append(result, w)
	}

	return strings.Join(result, " ")
}

// HELPER FUNCTION: capitalize first letter of a word
func capitalizeWord(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
