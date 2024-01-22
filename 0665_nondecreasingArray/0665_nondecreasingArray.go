package nondecreasingArray

func checkPossibility(nums []int) bool {
	peakIndex := -1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}

		if peakIndex >= 0 {
			return false
		}

		if i == 0 || nums[i+1] >= nums[i-1] {
			nums[i] = nums[i+1]
		} else {
			nums[i+1] = nums[i]
		}
		peakIndex = i
	}

	return true
}

func checkPossibility_old(nums []int) bool {
	j := -1

	for i := 1; i < len(nums); i++ {
		l, r := nums[i-1], nums[i]
		if r <= l {
			if j >= 0 {
				return false
			}
			j = i - 1
		}
	}

	return true
}
