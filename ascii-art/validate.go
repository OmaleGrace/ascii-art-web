package main

import (
	"fmt"
)

func ValidateInput(s string) (rune, error) {
	for _, w := range s {
		if w < 32 || w > 126 {
			return w, fmt.Errorf("non ascii character : %c", w)
		}
	}
	return 0, nil
}
