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
