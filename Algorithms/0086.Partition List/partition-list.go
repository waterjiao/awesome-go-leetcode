package _086_Partition_List

import . "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	smallHead := small
	large := &ListNode{}
	largeHead := large

	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	small.Next = largeHead.Next
	large.Next = nil
	return smallHead.Next
}
