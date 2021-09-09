package _701_Insert_into_a_Binary_Search_Tree

import "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *kit.TreeNode, val int) *kit.TreeNode {
	if root == nil {
		return &kit.TreeNode{Val: val}
	}
	var left, right *kit.TreeNode
	result := root
	for root != nil {
		left, right = root, root
		if val > root.Val {
			root = root.Right
			left = nil
			continue
		}
		root = root.Left
		right = nil
	}
	if left != nil {
		left.Left = &kit.TreeNode{Val: val}
	}
	if right != nil {
		right.Right = &kit.TreeNode{Val: val}
	}
	return result
}
