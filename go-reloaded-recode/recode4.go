package main  

import "fmt"
func punc(word string) bool {
	if word == "," || word == "!" {
		return true
	} else {
		return false
	}
}
func main() {
fmt.Println(punc(","))
fmt.Println(punc("!"))
fmt.Println(punc("x"))
}