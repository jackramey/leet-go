package mergeSortedArray

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
//   representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To
//   accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be
//   merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
//

// merge compares the values of the arrays in reverse to populate values in the buffer region of the nums1 array. In
// order to do this we keep track of 1 read cursor per slice (i, j) and 1 write cursor (k). By iterating in reverse
// we compare the read cursor values and assign the maximum to the write cursor location. If either write cursor reach
// zero then we no longer need to compare and can write the remainder to the write cursor location. This requires a
// second loop only if the read cursor for nums2 still has values because we are storing in place for nums1.
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
