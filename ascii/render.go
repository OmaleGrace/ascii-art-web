package ascii

func RenderLine(s string, banner map[rune][]string) []string {
	result := (make([]string, 8))
	for i := 0; i <= 7; i++ {
		for _, l := range s {
			result[i] += banner[l][i]
		}
	}
	return result
}
