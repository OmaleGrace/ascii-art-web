package main

import (
	"fmt"
	"strings"
)

func Article(w string) string {
	first := string(w[0])
	vowel := "aeiouhAEIOUH"
	if strings.ContainsAny(first, vowel) {
		return "an " + w
	}
	return "a " + w
}

func main() {

	fmt.Println(cap("oy"))

}
