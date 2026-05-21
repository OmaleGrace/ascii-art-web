package main 

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter a sentence:")
	read := bufio.NewReader(os.Stdin)
	in, _ := read.ReadString('\n')
	gr := strings.Split(in, " ")
	fmt.Println("No of words: ",len(gr))
	fit := strings.Join(gr, "")
	lin := strings.TrimSuffix(fit, "\n")
	fmt.Println("No of characters: ",len(lin))
}