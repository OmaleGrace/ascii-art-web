package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Punc(token []string) string {
	word := strings.Join(token, " ")

	re := regexp.MustCompile(`\s+([\p{P}])`)
	return re.ReplaceAllString(word, "$1")
}

func main() {
	fmt.Println(Punc([]string{"this", ",", "is", "  !"}))
}
