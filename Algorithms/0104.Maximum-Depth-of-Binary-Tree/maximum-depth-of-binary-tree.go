package _104_Maximum_Depth_of_Binary_Tree

import (
	"waterjiao.com/awesome-go-leetcode/kit"
)

func maxDepth(root *kit.TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l < r {
		return r + 1
	}
	return l + 1
}
