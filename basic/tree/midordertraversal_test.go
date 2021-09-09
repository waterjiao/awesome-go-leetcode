package tree

import (
	"fmt"
	"testing"
	"waterjiao.com/awesome-go-leetcode/kit"
)

func TestMidOrderTraversal(t *testing.T) {
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

	midorderTraversal(root)
	fmt.Print(midOrderTraversal(root))
}
