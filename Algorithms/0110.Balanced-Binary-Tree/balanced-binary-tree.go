package _110_Balanced_Binary_Tree

import "waterjiao.com/awesome-go-leetcode/kit"

// 与返回二叉树高度结合起来，通过判断左右子树的高度来确定是否为高度平衡二叉树

func isBalanced(root *kit.TreeNode) bool {
	if maxDepth(root) == -1 {
		return false
	}
	return true
}

func maxDepth(root *kit.TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
