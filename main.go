package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage; go run . <input> <output>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	text := string(data)
	text = baseConversion(text)
	text = applyCaseTransform(text)
	err = os.WriteFile(outputFile, []byte(text+"\n"), 0644)
	if err != nil {
		panic(err)
	}
}
