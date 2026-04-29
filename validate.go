package main

func ValidateInput(text string, banner map[rune][]string) bool {
	for _, r := range text {
		if r == '\n' || r == '\r' {
			continue
		}
		if _, exists := banner[r]; !exists {
			return false
		}
	}
	return true
}
