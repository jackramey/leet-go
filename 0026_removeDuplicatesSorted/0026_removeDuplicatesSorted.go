package removeDuplicatesSorted

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique
//   element appears only once. The relative order of the elements should be kept the same. Then return the number of
//   unique elements in nums.
//
// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
//    - Change the array nums such that the first k elements of nums contain the unique elements in the order they were
//      present in nums initially. The remaining elements of nums are not important as well as the size of nums.
//    - Return k.
//

// removeDuplicates iterates through the sorted array and if the current value is equal to the previous value we know
// it is a duplicate. We keep track of 2 cursors, a read cursor (i) which iterates through the array normally and a
// write cursor (j) that keeps track of where unique values should be stored when we come across them.
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
