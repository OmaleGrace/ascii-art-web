package main

import "strings"

func SplitInput(input string) []string {
	line := strings.Split(input, "\\n")
	return line
}
