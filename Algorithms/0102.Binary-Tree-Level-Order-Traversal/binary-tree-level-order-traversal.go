package _102_Binary_Tree_Level_Order_Traversal

import "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *kit.TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*kit.TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		r := make([]int, 0)
		l := len(stack)
		for i := 0; i < l; i++ {
			t := stack[0]
			stack = stack[1:]
			r = append(r, t.Val)
			if t.Left != nil {
				stack = append(stack, t.Left)
			}
			if t.Right != nil {
				stack = append(stack, t.Right)
			}
		}
		result = append(result, r)
	}
	return result
}
