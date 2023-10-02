package sameTree

// Given the roots of two binary trees p and q, write a function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSameTree is a recursive solution to determine if two trees are equivalent. It does this by asserting if the current
// node is in both trees is either a leaf (q && p are both nil) or if they are not equivalent. If we are not on a node
// then we recursively call this function on the left and right children and compare the results.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	// this check is effectively an XOR at this point because we've already returned if both p and q are nil. If true
	// that means that one of them is nil and the other isn't and therefore p and q are not equivalent
	if p == nil || q == nil {
		return false
	}

	if q == nil && p != nil {
		return false
	}

	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}

	return false
}
