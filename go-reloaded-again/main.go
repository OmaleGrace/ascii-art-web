package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage :go run . input.txt ooutput.txt")
		return
	}
	input := os.Args[1]
	output := os.Args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	text := string(data)
	text = HexToDec(text)
	text = BinToDec(text)
	text = TransformCases(text)
	text = FixPunctuation(text)
	text = FixQuotes(text)
	// text = fixArticle(text)

	err = os.WriteFile(output, []byte(text+"\n"), 0644)
	if err != nil {
		panic(err)
	}
}
