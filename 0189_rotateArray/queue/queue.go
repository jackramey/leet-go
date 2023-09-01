package queue

// Rotate using a queue to stash values. In order to do this we need to stash K values so that we don't overwrite them
// and then iterate through the array until the end. The loop for the queue could be a for loop that just checks to see
// if any values are still in the queue and in that case it would need to wrap around the end of the array. This runs in
// O(n+k) time with O(k) additional memory
func Rotate(nums []int, k int) {
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
