package tasks

import "testing"

func TestIsHappy(t *testing.T) {
	testCases := map[int]bool{
		19: true,
	}

	for n, expected := range testCases {
		result := isHappy(n)
		if expected != result {
			t.Errorf("[%d]: expected %v, got %v\n", n, expected, result)
		}
	}
}
