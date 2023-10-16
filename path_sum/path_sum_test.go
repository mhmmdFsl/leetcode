package path_sum

import (
	"github.com/mhmmdFsl/leetcode/data"
	"testing"
)

func TestPathSum(t *testing.T) {

	tree := &data.TreeNode{
		Val: 5,
		Left: &data.TreeNode{
			Val: 4,
			Left: &data.TreeNode{
				Val: 11,
				Left: &data.TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: &data.TreeNode{
					Val:   2,
					Right: nil,
					Left:  nil,
				},
			},
			Right: nil,
		},
		Right: &data.TreeNode{
			Val: 8,
			Left: &data.TreeNode{
				Val:   13,
				Left:  nil,
				Right: nil,
			},
			Right: &data.TreeNode{
				Val:  4,
				Left: nil,
				Right: &data.TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	e := true
	r := HasSum(tree, 22)

	if e != r {
		t.Fail()
	}
}
