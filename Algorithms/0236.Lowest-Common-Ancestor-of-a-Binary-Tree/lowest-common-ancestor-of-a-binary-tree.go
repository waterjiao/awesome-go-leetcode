package _236_Lowest_Common_Ancestor_of_a_Binary_Tree

import "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
分治法
模版
1。退出条件
2。左右划分
3。合并

套用此题
显示判断root是否为nil，如果为nil，直接返回
再判断root是否与p和q相等，如果相等，直接返回，因为一旦找到p或q，那么说明p和q在一定程度上，属于祖父节点

递归查找左右子树
判断left和right

如果当前结点root 等于 NULL，则直接返回 NULL
如果root等于p或者q，那这颗树一定返回p或者q
然后递归左右子树，用left和right表述
若left为空，最终结果要看right，反之看left
如果都为空，则返回空
如果都非空，则返回root，左右分别含有p和q

*/
func lowestCommonAncestor(root, p, q *kit.TreeNode) *kit.TreeNode {
	if root == nil {
		return root
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
