package majorityElement

func majorityElement(nums []int) int {
	elements := make(map[int]int)
	for _, x := range nums {
		elements[x]++
	}

	maxKey := 0
	maxVal := 0
	for key, val := range elements {
		if val > maxVal {
			maxKey = key
			maxVal = val
		}
	}
	return maxKey
}
