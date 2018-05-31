package uniqueMorseRepresentations

var codes = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	cmap := map[string]int{}
	for _, w := range words {
		var key string
		for _, r := range w {
			key += codes[r-'a']
		}
		cmap[key] = 0
	}
	return len(cmap)
}
