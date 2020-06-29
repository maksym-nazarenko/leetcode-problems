// Given a non-negative index k where k ≤ 33,
// return the kth index row of the Pascal's triangle.
// Note that the row index starts from 0.

// Example:
// Input: 3
// Output: [1,3,3,1]

package tasks

func getRow(rowIndex int) []int {
	var ret = make([]int, rowIndex+1)

	for i := 0; i < rowIndex+1; i++ {
		ret[i] = getItem(rowIndex, i)
	}

	return ret
}

func getItem(k, i int) int {
	if i == 0 || i == k {
		return 1
	}

	if i == 1 || i == k-1 {
		return k
	}

	return getItem(k-1, i-1) + getItem(k-1, i)
}
