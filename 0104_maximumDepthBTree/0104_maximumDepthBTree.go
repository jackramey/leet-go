package maximumDepthBTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
