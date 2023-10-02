package invertBinaryTree

// invertTree takes a binary tree and inverts the nodes s/t what used to be left nodes are now right nodes and vice versa.
// It does this by swapping the left and right nodes in place and then recursively inverting each child node as long as
// they are not nil
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
