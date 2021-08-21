package _142_Linked_List_Cycle_II

import (
	. "waterjiao.com/awesome-go-leetcode/kit"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func detectCycle(head *ListNode) *ListNode {
	var p, curr *ListNode = nil, head
	nodes := make(map[*ListNode]struct{})

	for curr != nil {
		if _, ok := nodes[curr]; ok {
			p = curr
			break
		}
		nodes[curr] = struct{}{}
		curr = curr.Next
	}
	if curr == nil {
		return curr
	}
	return p
}
