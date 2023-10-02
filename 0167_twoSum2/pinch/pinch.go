package pinch

// TwoSum iterates through the list of numbers with two cursors i and j that iterate through the set of numbers from
// opposite ends until it reaches the target value. If the sum is less than the target than we increment the low end
// cursor and if the sum is greater than the target than we decrement the high end cursor.
func TwoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		}
		if sum < target {
			l++
		}
		if sum > target {
			r--
		}
	}

	return nil
}
