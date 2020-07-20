package tasks

func matrixReshape(nums [][]int, r int, c int) [][]int {
	var origRows = len(nums)
	var origCols = len(nums[0])

	if r < 1 || c < 1 || r*c != origRows*origCols {
		return nums
	}

	var ret [][]int = make([][]int, r)
	for el := range ret {
		ret[el] = make([]int, c)
	}

	for i := 0; i < r*c; i++ {
		ret[i/c][i%c] = nums[i/origCols][i%origCols]
	}

	return ret
}
