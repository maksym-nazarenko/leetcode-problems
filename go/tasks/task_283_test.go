package tasks

import (
	"fmt"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var expected = []struct {
		inArray  []int
		outArray []int
	}{
		{
			inArray:  []int{0, 1, 0, 3, 12},
			outArray: []int{1, 3, 12, 0, 0},
		},
		{
			inArray:  []int{1, 2, 3, 4},
			outArray: []int{1, 2, 3, 4},
		},
		{
			inArray:  []int{0, 0, 0, 1},
			outArray: []int{1, 0, 0, 0},
		},
		{
			inArray:  []int{0},
			outArray: []int{0},
		},
		{
			inArray:  []int{2},
			outArray: []int{2},
		},
		{
			inArray:  []int{0, 2},
			outArray: []int{2, 0},
		},
	}
	for _, testCase := range expected {
		var inArray = testCase.inArray
		moveZeroes(inArray)
		if !arraysEqual(testCase.outArray, inArray) {
			fmt.Printf("%v != %v\n", testCase.outArray, inArray)
			t.Error("Arrays are not equal")
		}
	}
}
