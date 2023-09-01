package removeDuplicatesSorted

func removeDuplicates(nums []int) int {
	lastSeen := -101
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != lastSeen {
			nums[idx] = nums[i]
			lastSeen = nums[i]
			idx++
		}
	}
	return idx
}
