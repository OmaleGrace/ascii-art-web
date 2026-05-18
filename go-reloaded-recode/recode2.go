package main 

import (
	"fmt"
	"strings"
)

func cap(words string) string {
	return strings.ToUpper(string(words[0])) + strings.ToLower(words[1:])
}
func main() {
	fmt.Println(cap("hELLO"))
	fmt.Println(cap("WORLD"))
}