package _098_Validate_Binary_Search_Tree

import "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *kit.TreeNode) bool {
	result := make([]int, 0)
	midOrder(root, &result)
	return bst(result)
}

func bst(result []int) bool {
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func midOrder(root *kit.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	midOrder(root.Left, result)
	*result = append(*result, root.Val)
	midOrder(root.Right, result)
}
