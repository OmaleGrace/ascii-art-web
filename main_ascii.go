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
