package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func peace() {
	for {
		fmt.Println("Welcome To String Processor")
		fmt.Println()
		time.Sleep(3 * time.Second)

		reader := bufio.NewReader(os.Stdin)
		fmt.Println()
		fmt.Println("Enter your name: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var operation string
		fmt.Println("Enter operation: ")
		fmt.Println("1. lastChar")
		fmt.Println("2. capitalize")
		fmt.Println("3. deleteIndex")
		fmt.Scanln(&operation)
		switch operation {
		case "1":
			for i := 0; i < len(input); i++ {
				if input[i] != ' ' && (i == len(input)-1 || input[i+1] == ' ') {
					fmt.Print(string(input[i]), " ")
				}
			}
		case "2":
			fmt.Println(strings.ToUpper(input))

		case "3":
			fmt.Println("Enter indices to delete (space separated, e.g. 2 3): ")

			reader := bufio.NewReader(os.Stdin)
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			parts := strings.Fields(line)

			// convert to integers
			var indices []int
			for _, p := range parts {
				var num int
				fmt.Sscanf(p, "%d", &num)

				if num >= 0 && num < len(input) {
					indices = append(indices, num)
				} else {
					fmt.Println("Invalid index:", num)
				}
			}

			// delete from right to left (VERY IMPORTANT)
			for i := len(indices) - 1; i >= 0; i-- {
				idx := indices[i]
				input = input[:idx] + input[idx+1:]
			}
			fmt.Println("Result:", input)
		default:
			fmt.Println("invalid operation")
		}
	}
}

func main() {
	peace()
}
