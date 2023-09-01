package march

// Rotate - Brute force approach, rotate the array by 1 each time and rotate k times. Runs in O(n*k)
func Rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		rotateOne(nums)
	}
}

func rotateOne(nums []int) {
	stash := nums[0]
	for i := 0; i < len(nums); i++ {
		nextIndex := getIndex(i, 1, len(nums))
		tmp := nums[nextIndex]
		nums[nextIndex] = stash
		stash = tmp
	}
}

func getIndex(i, k, length int) int {
	return (i + k) % length
}
