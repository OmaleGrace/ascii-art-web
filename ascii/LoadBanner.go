package ascii

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	ban, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	fit := strings.ReplaceAll(string(ban), "\r", "\n")
	line := strings.Split(fit, "\n")
	mp := make(map[rune][]string)
	rn := rune(32)
	for i := 0; i < len(line); i += 9 {
		if i+8 < len(line) {
			mp[rn] = line[i+1 : i+9]
			rn++
		}

		if rn > 126 {
			break
		}

		if len(mp) == 0 {
			return nil, errors.New("error")
		}
	}
	return mp, err
}
