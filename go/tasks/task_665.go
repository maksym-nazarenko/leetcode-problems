package tasks

func checkPossibility(nums []int) bool {
	failedIndex := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if failedIndex >= 0 {
				return false
			}

			failedIndex = i
		}
	}

	return failedIndex <= 0 || failedIndex >= len(nums)-2 || nums[failedIndex-1] <= nums[failedIndex+1] || nums[failedIndex] <= nums[failedIndex+2]
}
