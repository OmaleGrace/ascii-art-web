package ascii 

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(fielname string) (map[rune][]string, error) {
	ban, err := os.ReadFile(fielname)
	if err != nil {
		return nil, err
	}
	gin := strings.ReplaceAll(string(ban), "\r", "\n")
	alt := strings.Split(gin, "\n")
	mp := make(map[rune][]string)
	rn := rune(32)
	for i := 0; i < len(alt); i += 9 {
		if i+8 < len(alt) {
			mp[rn] = alt[rn+1 : rn+9]
			rn++
		}
		if rn > 126 {
			break
		}
		if len(mp) == 0 {
			return nil, errors.New("Error found")
		}
	}
	return mp, err
}
