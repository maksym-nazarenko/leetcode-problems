package tasks

import "testing"

func TestMaxProfit(t *testing.T) {
	var expected = []struct {
		prices []int
		result int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			result: 7,
		},
		{
			prices: []int{1, 2, 3, 4, 5},
			result: 4,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			result: 0,
		},
		{
			prices: []int{1, 1, 1, 10, 10, 10},
			result: 9,
		},
	}

	for _, testCase := range expected {
		result := maxProfit(testCase.prices)
		if testCase.result != result {
			t.Errorf("Array: %v\nExpected: %d, got %d\n", testCase.prices, testCase.result, result)
		}
	}
}
