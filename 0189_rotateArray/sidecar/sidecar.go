package sidecar

// Rotate uses a "sidecar" approach which creates a second array of equal length and adds the offset values to it. Once
// complete with the rotation it copies the rotated array back to the original array. This runs in O(N) time with O(N)
// additional memory used.
func Rotate(nums []int, k int) {
	rot := make([]int, len(nums))
	for i, num := range nums {
		rot[getIndex(i, k, len(nums))] = num
	}

	copy(nums, rot)
}

func getIndex(i, k, length int) int {
	return (i + k) % length
}
