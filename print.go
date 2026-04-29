package main

import (
	"strings"
	"fmt"
)

func PrintArt(input string, banner map[rune][]string) {
    if input == "\\n" || input == "\n" {
        fmt.Println()
        return
    }
    lines := strings.Split(input, "\\n")
    for _, line := range lines {
        if line == "" {
            fmt.Println()
            continue
        }
        asciiRows := RenderLine(line, banner)
        for _, row := range asciiRows {
            fmt.Println(row)
        }
    }
}
