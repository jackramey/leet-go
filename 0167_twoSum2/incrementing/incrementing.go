package incrementing

// TwoSum iterates through the list of numbers with two cursors i and j that iterate through the entire set of numbers
// until it reaches the target value. This is a brute-force solution that operates in O(n^2) time
func TwoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return nil
}
