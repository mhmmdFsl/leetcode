package path_sum

import (
	"github.com/mhmmdFsl/leetcode/data"
)

func HasSum(root *data.TreeNode, targetSum int) bool {
	if root == nil && targetSum == 0 {
		return false
	}

	if root == nil && (targetSum > 0 || targetSum < 0) {
		return false
	}

	var t = false
	traverseAndSum(root, targetSum, &t)

	return t
}

func traverseAndSum(root *data.TreeNode, target int, t *bool) {

	if *t == true {
		*t = true
	}

	if root == nil {
		return
	}

	if root.Left != nil {
		root.Left.Val = root.Left.Val + root.Val
		traverseAndSum(root.Left, target, t)
	}

	if root.Right != nil {
		root.Right.Val = root.Right.Val + root.Val
		traverseAndSum(root.Right, target, t)
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == target {
			*t = true
		}
	}

}
