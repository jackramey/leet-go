package majorityElement

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element
//   always exists in the array.

// majorityElement creates a map of all values and a counter value. Each time the value appears we increment the counter
// for that value. If that counter value is greater than n / 2 then we have found the majority element and we return.
func majorityElement(nums []int) int {
	elements := make(map[int]int)
	for _, x := range nums {
		elements[x]++
		if elements[x] > (len(nums) / 2) {
			return x
		}
	}

	return 0
}
