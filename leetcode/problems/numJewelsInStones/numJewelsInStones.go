package numJewelsInStones

func numJewelsInStones(J string, S string) int {
	jmap := map[rune]int{}
	for _, j := range J {
		jmap[j] = 0
	}

	for _, s := range S {
		if _, ok := jmap[s]; ok == true {
			jmap[s]++
		}
	}
	var total int
	for _, j := range jmap {
		total += j
	}
	return total
}
