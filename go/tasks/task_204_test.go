package tasks

import "testing"

func TestCountPrimes(t *testing.T) {
	var testCases = map[int]int{
		10: 4,
		11: 4,
		12: 5,
		2:  0,
		1:  0,
	}

	for n, expected := range testCases {
		ret := countPrimes(n)
		if ret != expected {
			t.Errorf("n=%d: Expected %d, got %d\n", n, expected, ret)
		}
	}
}
