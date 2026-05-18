package main

import (
	"fmt"
	"strconv"
)

func ConvertToDecimal(s string, base int) (int64, error) {
	a, _ := strconv.ParseInt("10", 2, 64)
	b, _ := strconv.ParseInt("1010", 2, 64)
	c, _ := strconv.ParseInt("11111111", 2, 64)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	return strconv.ParseInt(s, base, 64)
}
func main() {
	fmt.Println(ConvertToDecimal("1E", 16))
}
