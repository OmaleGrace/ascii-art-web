package main

import (
	"fmt"
)

func main() {
	var grace string
	fmt.Println("Enter your name:")
	fmt.Scan(&grace)
	if grace == "smallie" {
		fmt.Println("correct")
	} else {
		fmt.Print("incorrect")
	}
}
