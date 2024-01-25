package attempt2

func findDuplicates(nums []int) []int {
	results := []int{}
	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] < 0 {
			results = append(results, index+1)
		} else {
			nums[index] = -nums[index]
		}
	}
	return results
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
