package tree

import "waterjiao.com/awesome-go-leetcode/kit"

func levelOrderTraversal(root *kit.TreeNode) [][]int {
	// 使用二维切片，存放每一层的数据
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	// 使用队列，每次将当前node的左右节点存到队列中
	queue := make([]*kit.TreeNode, 0)
	queue = append(queue, root)
	// 遍历队列，直到队列为空
	for len(queue) > 0 {
		list := make([]int, 0)
		// 取当前新的一层时，队列的长度，后续添加的等到下一层计算时使用
		l := len(queue)
		for i := 0; i < l; i++ {
			// 取队列的头，取出后截取队列，符合队列操作
			level := queue[i]
			queue = queue[1:]
			list = append(list, level.Val)

			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
		}
		result = append(result, list)
	}
	return result
}
