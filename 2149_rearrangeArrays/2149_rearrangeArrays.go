package rearrangeArrays

// Strategy is to separate the values into 2 arrays based on their sign and then recombine. This preserves order
// and ensures that we only have to loop through the values twice resulting in an O(N) time complexity and O(N) space
// complexity.
func rearrangeArray(nums []int) []int {
	pos := []int{}
	neg := []int{}
	for _, num := range nums {
		if num < 0 {
			neg = append(neg, num)
		} else {
			pos = append(pos, num)
		}
	}

	for i := 0; i < len(pos); i++ {
		idx := i * 2
		nums[idx] = pos[i]
		nums[idx+1] = neg[i]
	}

	return nums
}
