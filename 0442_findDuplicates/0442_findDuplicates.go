package findDuplicates

func findDuplicates(nums []int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	values := map[int]bool{}
	duplicates := []int{}
	for _, num := range nums {
		_, ok := values[num]
		if ok {
			duplicates = append(duplicates, num)
		} else {
			values[num] = true
		}
	}
	return duplicates
}

func findDuplicatesInPlace(nums []int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	results := []int{}
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		v := nums[idx]
		if v < 0 {
			results = append(results, abs(v))
		} else {
			nums[idx] = -v
		}
	}

	return results
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
