package rotateArray

func rotate(nums []int, k int) {
	if len(nums) == 1 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}

	q := queue{}
	for i := len(nums) - k; i < len(nums); i++ {
		idx := i % len(nums)
		q.push(nums[idx])
	}
	for i := 0; i < len(nums); i++ {
		q.push(nums[i])
		nums[i] = q.pop()
		if len(q) <= 0 {
			return
		}

	}
}

type queue []int

func (q *queue) push(i int) {
	*q = append(*q, i)
}

func (q *queue) pop() int {
	h := *q
	var el int
	l := len(h)
	el, *q = h[0], h[1:l]
	return el
}
