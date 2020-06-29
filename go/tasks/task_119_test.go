package tasks

import "testing"

func arraysEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestGetRow(t *testing.T) {
	var results map[int][]int = map[int][]int{
		0: []int{1},
		3: []int{1, 3, 3, 1},
		1: []int{1, 1},
		2: []int{1, 2, 1},
		8: []int{1, 8, 28, 56, 70, 56, 28, 8, 1},
	}
	for k, expected := range results {
		actualResult := getRow(k)
		if !arraysEqual(expected, actualResult) {
			t.Errorf("Row %d: Expected %v, but got %v", k, expected, actualResult)
		}
	}
}
