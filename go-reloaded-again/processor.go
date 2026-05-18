package main

import (
	"regexp"
	"strconv"
	"strings"
)

// CONVERT HEXADECIMAL TO DECIMAL
func HexToDec(text string) string {
	word := strings.Fields(text)
	var result []string
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0 {
			hexVal := word[i-1]
			dec, err := strconv.ParseInt(hexVal, 16, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(dec, 10)
			}
			continue
		}
		result = append(result, word[i])
	}
	return strings.Join(result, " ")
}

// CONVERT BINARY TO DECIMAL
func BinToDec(text string) string {
	words := strings.Fields(text)
	var result []string
	for i := 0; i < len(words); i++ {
		if words[i] == "(bin)" && i > 0 {
			binVal := words[i-1]
			bin, err := strconv.ParseInt(binVal, 2, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(bin, 10)
			}
			continue
		}
		result = append(result, words[i])
	}
	return strings.Join(result, " ")
}

// HANDLE (UP) , (UP, N) & (LOW) , (LOW, N) & (CAP) , (CAP, N)
func TransformCases(text string) string {
	var result []string
	words := strings.Fields(text)
	for i := 0; i < len(words); i++ {
		w := words[i]
		if strings.HasPrefix(w, "(") {
			for !strings.HasSuffix(w, ")") {
				i++
				w += " " + words[i]
			}
			p := strings.Split(strings.Trim(w, "()"), ",")
			d := strings.TrimSpace(p[0])
			n, _ := strconv.Atoi(strings.TrimSpace(strings.Join(p[1:], "")))
			if n < 1 {
				n = 1
			}
			for j := len(result) - n; j < len(result); j++ {
				if j >= 0 {
					switch d {
					case "up":
						result[j] = strings.ToUpper(result[j])
					case "low":
						result[j] = strings.ToLower(result[j])
					case "cap":
						result[j] = strings.Title(result[j])
					}
				}
			}
			continue
		}
		result = append(result, w)
	}
	return strings.Join(result, " ")
}

// HANDLES PUNCTUATIONS
func FixPunctuation(text string) string {
	text = regexp.MustCompile(`\s([,";:.?!])`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`\s([,;:"?!])([,;:"?!])\s`).ReplaceAllString(text, "$1 $2")
	text = regexp.MustCompile(`\s.\s*\.\s*\.`).ReplaceAllString(text, "...")
	text = regexp.MustCompile(`\s([,";:?!])`).ReplaceAllStri ng(text, "$1 ")
	return text
}

// HANDLE QUOTES
func FixQuotes(text string) string {
	text = regexp.MustCompile(`'\s([^']+?)\s'`).ReplaceAllString(text, "'$1'")
	return text
}
