package minimumdepthbynaritree

import (
	"testing"

	"github.com/mhmmdFsl/leetcode/data"
)

func TestMinimumDepth(t *testing.T) {

	tc1 := &data.TreeNode{
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

	// tc2 := &data.TreeNode{
	// 	Val: 2,
	// 	Right: &data.TreeNode{
	// 		Val: 3,
	// 		Right: &data.TreeNode{
	// 			Val: 4,
	// 			Right: &data.TreeNode{
	// 				Val: 5,
	// 				Right: &data.TreeNode{
	// 					Val: 6,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	i := minDepth(tc1)
	// i2 := minDepth(tc2)

	if i != 2 {
		t.Fail()
	}
}
