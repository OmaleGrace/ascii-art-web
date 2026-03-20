package main

import (
	"strconv"
	"strings"
)

// CONVERTS HEXADECIMAL TO AND BINARY TO DECIMAL
func baseConversion(text string) string {
	word := strings.Fields(text)
	var result []string
	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" && i > 0 {
			hexval := word[i-1]
			dec, err := strconv.ParseInt(hexval, 16, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(dec, 10)
			}
			continue
		}
		if word[i] == "(bin)" && i > 0 {
			hexval := word[i-1]
			dec, err := strconv.ParseInt(hexval, 2, 64)
			if err == nil {
				result[len(result)-1] = strconv.FormatInt(dec, 10)
			}
			continue
		}
		result = append(result, word[i])
	}
	return strings.Join(result, " ")
}

// HANDLES (UP), (LOW) AND (CAP) CASE
func applyCaseTransform(text string) string {
	word := strings.Fields(text)
	var result []string
	for i := 0; i < len(word); i++ {
		c := word[i]
		if strings.HasPrefix(c, "(") {
			n := c
			for !strings.HasSuffix(n, ")") && i+1 < len(word) {
				i++
				n += " " + word[i]
			}
			n = strings.TrimSuffix(strings.TrimPrefix(n, "("), ")")
			pt := strings.Split(n, ",")
			action := strings.TrimSpace(pt[0])
			count := 1
			if len(pt) == 2 {
				if j, err := strconv.Atoi(strings.TrimSpace(pt[1])); err == nil {
					count = j
				}
			}
			for g := len(result) - count; g < len(result); g++ {
				if g >= 0 {
					switch action {
					case "up":
						result[g] = strings.ToUpper(result[g])
					case "low":
						result[g] = strings.ToLower(result[g])
					case "cap":
						result[g] = capitalizeword(result[g])
					}
				}
			}
			continue
		}
		result = append(result, c)
	}
	return strings.Join(result, " ")
}
func capitalizeword(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}
