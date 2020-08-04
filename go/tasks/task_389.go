package tasks

func findTheDifference(s string, t string) byte {
	var origStringMap map[rune]int = map[rune]int{}
	for _, c := range s {
		if _, ok := origStringMap[c]; !ok {
			origStringMap[c] = 0
		}

		origStringMap[c]++
	}

	for _, c := range t {
		if _, ok := origStringMap[c]; !ok || origStringMap[c] < 1 {
			return byte(c)
		}

		origStringMap[c]--
	}

	return ' '
}
