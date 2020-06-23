package tasks

import (
	"testing"
)

type inputTestData struct {
	nums         []int
	expectedNums []int
	searchValue  int
}

func equalsN(a, b []int, n int) bool {
	minN := len(a)
	if len(b) < minN {
		minN = len(b)
	}

	for i := 0; i < minN; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestRemoveElement(t *testing.T) {
	var ret int
	var inputNums []int

	inputNums = []int{3, 2, 2, 3}
	ret = removeElement(inputNums, 3)
	if ret != 2 {
		t.Errorf("Wanted %d, got %d", 3, ret)
	}

	expected := []int{2, 2}
	if !equalsN(expected, inputNums, 2) {
		t.Errorf("Wanted %v, got %v", expected, inputNums)
	}
}

func TestRemoveElement2(t *testing.T) {
	var ret int
	var inputNums []int

	inputNums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	ret = removeElement(inputNums, 2)
	if ret != 5 {
		t.Errorf("Wanted %d, got %d", 5, ret)
	}

	expected := []int{0, 1, 3, 0, 4}
	if !equalsN(expected, inputNums, 2) {
		t.Errorf("Wanted %v, got %v", expected, inputNums)
	}
}
