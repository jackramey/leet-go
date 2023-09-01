package stash

// Rotate - This algorithm is what I'm calling "stash Rotate". It tries to bounce around the array stashing the
// nums[i+k%len] value and replacing it in place with the nums[i] value. The idea is that instead of iterating to i++
// you move to i = i + k so that the stashed value isn't lost. The problem here is identifying how many times to run
// the rotate algorithm and figuring out what the offset is. I haven't been able to deduce that successfully and I have
// a sinking feeling it's going to take some weird discrete math...
func Rotate(nums []int, k int) {
	// edge cases
	if len(nums) == 1 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}

	stashRotateHelper(nums, k, 0)
	if k != 1 && len(nums)/k > 0 && len(nums)%k == 0 {
		for i := 1; i < k; i++ {
			stashRotateHelper(nums, k, i)
		}
	}
}

func stashRotateHelper(nums []int, k int, startIndex int) {
	stash := nums[startIndex]
	i := startIndex
	for {
		nextIdx := getIndex(i, k, len(nums))
		tmp := nums[nextIdx]
		nums[nextIdx] = stash
		stash = tmp
		i = nextIdx
		if i == startIndex {
			break
		}
	}
}

func getIndex(i, k, length int) int {
	return (i + k) % length
}
