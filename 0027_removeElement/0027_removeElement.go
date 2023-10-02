package removeElement

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the
//   elements may be changed. Then return the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the
//   following things:
//     - Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The
//         remaining elements of nums are not important as well as the size of nums.
//     - Return k.
//

// removeElement iterates through an unsorted array and removes elements that match val. We do this by keeping 2 cursors
// a read cursor i and a write cursor j. The read cursor iterates through the slice normally but the write cursor
// is used to replace values that match the input and so is not incremented if there is a match.
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
