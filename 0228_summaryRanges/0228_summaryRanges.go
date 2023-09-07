package summaryRanges

import "strconv"

const sep = "->"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		return addRange([]string{}, nums[0], nums[0])
	}

	var ranges []string
	i, j := 0, 1
	for {
		if nums[j] == nums[j-1]+1 {
			j++
		} else {
			ranges = addRange(ranges, nums[i], nums[j-1])
			i, j = j, j+1
		}
		if j >= len(nums) {
			ranges = addRange(ranges, nums[i], nums[j-1])
			return ranges
		}
	}
}

func addRange(ranges []string, x, y int) []string {
	var r string
	if x == y {
		r = strconv.Itoa(x)
	} else {
		r = strconv.Itoa(x) + sep + strconv.Itoa(y)
	}
	return append(ranges, r)
}
