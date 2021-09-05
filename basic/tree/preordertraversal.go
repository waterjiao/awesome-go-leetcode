package tree

import (
	"fmt"
	"waterjiao.com/awesome-go-leetcode/kit"
)

func preorderTraversal(root *kit.TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func preOrderTraversal(root *kit.TreeNode) []int{
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	stack := make([]*kit.TreeNode, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			result = append(result, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return result
}