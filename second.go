package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func grace() {
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
			var num int
			for {
				fmt.Println("Enter index: ")
				fmt.Scan(&num)

				if num >= 0 && num < len(input) {
					input = input[:num] + input[num+1:]
					fmt.Println(input)
					break
				} else {
					fmt.Println("Invalid Index restarting .....")
					continue
				}
			}
		default:
			fmt.Println("invalid operation")
		}
	}
}

func main() {
	grace()
}
