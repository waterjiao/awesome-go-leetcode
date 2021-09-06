package tree

import (
	"fmt"
	"waterjiao.com/awesome-go-leetcode/kit"
)

func postorderTraversal(root *kit.TreeNode) {
	if root == nil {
		return
	}
	postorderTraversal(root.Left)
	postorderTraversal(root.Right)
	fmt.Print(root.Val, " ")
}

func postOrderTraversal(root *kit.TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*kit.TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}
