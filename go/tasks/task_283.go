package tasks

func moveZeroes(nums []int) {
	var i = 0
	var j = 0
	var n = len(nums)

	for i < n && j < n {
		for i < n && nums[i] != 0 {
			i++
		}
		j = i + 1
		for j < n && nums[j] == 0 {
			j++
		}

		if j >= n {
			break
		}

		nums[i], nums[j] = nums[j], 0

		i++
	}
}
