package main

import (
	"fmt"
	"regexp"
)

func fix(w string) string {
	re1 := regexp.MustCompile(`\s+'`)
	re2 := regexp.MustCompile(`'\s+`)
	re3 := regexp.MustCompile(`\s+"`)
	re4 := regexp.MustCompile(`"\s+`)

	w = re1.ReplaceAllString(w, "'")
	w = re2.ReplaceAllString(w, "'")
	w = re3.ReplaceAllString(w, `"`)
	w = re4.ReplaceAllString(w, `"`)

	return w

}

func main() {
	fmt.Println(fix("'    great     '"))
}
