package tree

import (
	"fmt"
	"waterjiao.com/awesome-go-leetcode/kit"
)

func midorderTraversal(root *kit.TreeNode) {
	if root == nil {
		return
	}
	midorderTraversal(root.Left)
	fmt.Print(root.Val, " ")
	midorderTraversal(root.Right)
}

func midOrderTraversal(root *kit.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*kit.TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		root = node.Right
	}
	return result
}
