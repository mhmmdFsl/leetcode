package balancedbinarytree

import (
	"testing"

	"github.com/mhmmdFsl/leetcode/data"
)

func TestBalancedBinaryTree(t *testing.T) {

	tree := data.TreeNode{
		Val: 3,
		Left: &data.TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &data.TreeNode{
			Val: 20,
			Left: &data.TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &data.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	b := isBalanced(&tree)
	if !b {
		t.Fail()
	}
}
