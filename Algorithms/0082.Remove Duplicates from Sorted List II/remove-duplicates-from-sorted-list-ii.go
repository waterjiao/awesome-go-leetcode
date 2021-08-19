package _082_Remove_Duplicates_from_Sorted_List_II

import . "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	num := 0
	for  p.Next != nil && p.Next.Next != nil{
		if p.Next.Val == p.Next.Next.Val {
			num = p.Next.Val
			p.Next = p.Next.Next.Next
			for p.Next != nil {
				if p.Next.Val == num {
					p.Next = p.Next.Next
				} else {
					break
				}
			}
		}  else {
			p = p.Next
		}
	}
	return dummy.Next
}
