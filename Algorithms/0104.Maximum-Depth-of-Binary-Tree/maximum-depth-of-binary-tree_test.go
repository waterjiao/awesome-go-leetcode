package _104_Maximum_Depth_of_Binary_Tree

import (
	"fmt"
	"testing"
	"waterjiao.com/awesome-go-leetcode/kit"
)

func TestMax(t *testing.T) {
	root := &kit.TreeNode{
		Val: 1,
		Left: &kit.TreeNode{
			Val: 2,
			Left: &kit.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &kit.TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &kit.TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Print(maxDepth(root))
}
