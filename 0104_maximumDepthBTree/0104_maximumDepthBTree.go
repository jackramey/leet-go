package maximumDepthBTree

// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest
//   leaf node.
//

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth uses a DFS algorithm to recursively find the deepest path in a tree. If the node in question is a leaf then
// we return 0 which accounts for the trivial case of the root being nil. By returning zero on the current node that
// means we have to add 1 to the return value of the check on child nodes. Once we have the values of the maxDepth for
// the child nodes we compare them and return the maximum.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)
	if lDepth > rDepth {
		return lDepth + 1
	} else {
		return rDepth + 1
	}
}
