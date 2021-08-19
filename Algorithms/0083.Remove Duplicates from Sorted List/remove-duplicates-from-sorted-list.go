package _083_Remove_Duplicates_from_Sorted_List

import . "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	p := head
	for p != nil && p.Next != nil{
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	p := head
	pre := p
	m := make(map[int]bool)

	for p != nil {
		if m[p.Val] {
			pre.Next = p.Next
		} else {
			m[p.Val] = true
			pre = p
		}
		p = p.Next
	}
	return head
}
