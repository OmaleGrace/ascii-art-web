package main

import (
	"fmt"
)

func main() {
	f, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Print(RenderLine("Hi", f))
}
package main

import (
	"fmt"
	//"os"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: go run . [string]")
	// 	return
	// }
	// // input := os.Args[1]
	// // // if input == "\\n" {
	// // // 	fmt.Println()
	// // // 	return
	// // // }
	// // // if input == "" {
	// // // 	return
	// // // }
	grace, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error loading file", err)
		return
	}
	g := RenderLine("hi", grace)
	for _, e := range g {
		fmt.Println(e)
	}
	// if !ValidateInput(input, grace) {
	// 	fmt.Println("Error: Input contains unsupported characters.")
	// 	return
	// }

	// PrintArt(input, grace)
}
