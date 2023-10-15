package balancedbinarytree

import "github.com/mhmmdFsl/leetcode/data"

func isBalanced(root *data.TreeNode) bool {
	if root == nil {
		return true
	}

	dif := maxDepth(root.Left) - maxDepth(root.Right)
	if dif <= 1 && dif >= -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func maxDepth(node *data.TreeNode) int {
	if node == nil {
		return 0
	}
	return max(maxDepth(node.Left), maxDepth(node.Right)) + 1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
