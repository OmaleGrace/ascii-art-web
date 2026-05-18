package main

import (
	"fmt"
	"strings"
)

func cap(w string) string {
	f := strings.Fields(w)
	vowel := "aeiouhAEIOUH"

	for i := 0; i < len(f)-1; i++ {
		first := string(f[i+1][0])

		if f[i] == "a" || f[i] == "an" {
			if strings.ContainsAny(first, vowel) {
				f[i] = "an"
			} else {
				f[i] = "a"
			}
		}
		if f[i] == "A" || f[i] == "An" {
			if strings.ContainsAny(first, vowel) {
				f[i] = "An"
			} else {
				f[i] = "A"
			}
		}
	}
	return strings.Join(f, " ")

}
func main() {

	fmt.Println(cap("An boy a apple an goat a orange"))

}