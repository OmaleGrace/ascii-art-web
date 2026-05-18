package main 

import (
	"fmt"
	"strings"
)
func upN(words []string, n int) []string {
	for i := 0; i < len(words); i++ {
		if i >= n {
			words[i] = strings.ToUpper(words[i])
		}
	}
	return words
}
func main() {
	s :=[]string{"this", "is", "so", "exciting"}
	fmt.Println(upN(s,2))
}