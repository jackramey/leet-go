package mergeSortedArray

// Compare the values of the arrays in reverse to populate values in the buffer region of the nums1 array.
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Define our value cursors to be the end of both arrays and the populate cursor to be the end of the nums1 array.
	i, j, k := m-1, n-1, m+n-1
	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i-- // Decrement nums1 cursor
		} else {
			nums1[k] = nums2[j]
			j-- // Decrement nums2 cursor
		}
	}

	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j-- // Decrement nums2 cursor
	}
}
