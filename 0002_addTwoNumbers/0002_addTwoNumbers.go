package addTwoNumbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNodes(l1, l2, 0)
}

func addNodes(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	var result *ListNode
	if carry != 0 {
		result = &ListNode{
			Val: carry,
		}
	}
	if l1 == nil && l2 == nil {
		return result
	}
	if l1 == nil {
		l1 = &ListNode{}
	}
	if l2 == nil {
		l2 = &ListNode{}
	}

	sum := l1.Val + l2.Val + carry

	digit := sum % 10
	carry = sum / 10
	return &ListNode{
		Val:  digit,
		Next: addNodes(l1.Next, l2.Next, carry),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
