package twoSum2

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such
//   that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2]
//   where 1 <= index1 < index2 < numbers.length.
//
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of
//   length 2.
//
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
//
// Your solution must use only constant extra space.

// twoSum iterates through the list of numbers with two cursors i and j that iterate through the entire set of numbers
// until it reaches the target value. This is a brute-force solution that operates in O(n^2) time
func twoSum(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}

	return nil
}
