package minimumdepthbynaritree

import (
	"github.com/mhmmdFsl/leetcode/data"
)

func minDepth(tree *data.TreeNode) int {
	if tree == nil {
		return 0
	}

	return max(minDepth(tree.Right), minDepth(tree.Left)) + 1
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
