package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	ben, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	grace := strings.ReplaceAll(string(ben), "\r", "\n")
	line := strings.Split(grace, "\n")
	tit := make(map[rune][]string)
	lit := rune(32)
	for i := 0; i < len(line); i += 9 {
		if i+8 < len(line) {
			tit[lit] = line[i+1 : i+9]
			lit++
		}
		if lit > 126 {
			break
		}
		if len(tit) == 0 {
			return nil, errors.New("error")
		}
	}
	return tit, err
}
