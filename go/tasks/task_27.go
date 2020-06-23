package tasks

func removeElement(nums []int, val int) int {
	// p1, p2 - indices for head and tail of an array
	// when p1 points to "val", shift right-part of the array after p1
	// one element left
	// repeat until p1 >= p2
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}

		return 1
	}

	p1 := 0
	ret := len(nums)

	for p1 < ret {
		if nums[p1] != val {
			p1++
			continue
		}

		if nums[p1] == val && p1 == ret-1 {
			ret--
			break
		}

		for i := p1 + 1; i < ret; i++ {
			nums[i-1] = nums[i]
		}

		ret--
	}

	return ret
}
