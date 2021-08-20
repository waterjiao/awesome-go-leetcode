package main

import (
	"fmt"
	. "waterjiao.com/awesome-go-leetcode/kit"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}
	var pre *ListNode
	origin := &ListNode{}
	originHeader := origin
	for head != nil {
		//origin.Next = head
		//origin = origin.Next
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	origin = originHeader.Next
	for origin.Val != pre.Val {
		o1 := origin.Next
		o2 := pre.Next
		origin.Next = pre
		pre.Next = o1
		origin = o1
		pre = o2
	}
	origin.Next = nil
	head = originHeader.Next
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	reorderList(head)
	fmt.Println(head)
}
