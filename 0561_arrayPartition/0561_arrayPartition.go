package arrayPartition

import "slices"

func arrayPairSum(nums []int) int {
	slices.Sort(nums)
	sum := 0
	for i := 0; i+1 < len(nums); i = i + 2 {
		sum += min(nums[i], nums[i+1])
	}
	return sum
}
