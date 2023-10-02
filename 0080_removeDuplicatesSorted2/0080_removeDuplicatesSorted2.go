package removeDuplicatesSorted2

// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique
//   element appears at most twice. The relative order of the elements should be kept the same.
//
// Since it is impossible to change the length of the array in some languages, you must instead have the result be
//   placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates,
//   then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first
//   k elements.
//
//  - Return k after placing the final result in the first k slots of nums.
//
//  Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
//

// removeDuplicates iterates through the sorted array and if the current value is equal to the previous value we know
// it is a duplicate. We keep track of 2 cursors, a read cursor (i) which iterates through the array normally and a
// write cursor (j) that keeps track of where values should be written when we come across a value that has not appeared
// twice already.
func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
