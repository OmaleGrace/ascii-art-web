package ascii

func RenderLine(s string, ban map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i <=7; i++ {
		for _, l := range s {
			result[i] = ban[l][i]
		}
	}
	return result
}