package main

import (
	"fmt"
	"os"
	"strings"
)

// READ THE BANNER FILE
func readBanner(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(content), "\n")
	return lines, nil
}

// GET A CHARTACTERS ART & LINE
func getLines(bannerLines []string, ch rune) []string {
	start := (int(ch) - 32) * 9
	return bannerLines[start+1 : start+9]
}

// PRINT ASCII ART
func printArt(lines []string, bannerLines []string) {

	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		for i := 0; i <= 7; i++ {
			for _, ch := range line {
				charArt := getLines(bannerLines, ch)
				fmt.Print(charArt[i])
			}
			fmt.Println()
		}
	}
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . [string]")
		return
	}
	input := os.Args[1]
	lines := strings.Split(input, "\\n")
	bannerLines, err := readBanner("banner/standard.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	printArt(lines, bannerLines)
}
