// Given a non-negative index k where k ≤ 33,
// return the kth index row of the Pascal's triangle.
// Note that the row index starts from 0.

// Example:
// Input: 3
// Output: [1,3,3,1]

package tasks

func getRow(rowIndex int) []int {
	var ret = make([]int, rowIndex+1)
	ret[0] = 1

	for i := 1; i < rowIndex+1; i++ {
		ret[i] = ret[i-1] * (rowIndex - (i - 1)) / i
	}

	return ret
}
