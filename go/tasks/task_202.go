package tasks

import (
	"strconv"
)

func isHappy(n int) bool {
	processed := make(map[int]bool)
	cycle := false
	newN := n
	prevN := n

	for newN != 1 && !cycle {
		digits := strconv.Itoa(prevN)
		newN = 0
		for _, d := range digits {
			newN += int(byte(d-48) * byte(d-48))
		}

		processed[prevN] = true
		prevN = newN
		if _, ok := processed[newN]; ok {
			cycle = true
		}
	}

	return newN == 1 && !cycle
}
