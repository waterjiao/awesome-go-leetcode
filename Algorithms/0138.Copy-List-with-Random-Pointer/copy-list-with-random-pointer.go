package _138_Copy_List_with_Random_Pointer

import (
	. "waterjiao.com/awesome-go-leetcode/kit"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	curr := head
	for curr != nil {
		newNode := &Node{Val: curr.Val}
		newNode.Next = curr.Next
		curr.Next = newNode
		curr = newNode.Next
	}
	curr = head
	for curr != nil {
		curr.Next.Random = curr.Random.Next
		curr = curr.Next.Next
	}
	curr = &Node{}
	dummy := curr
	p := head
	for p != nil {
		curr.Next = p.Next
		curr = curr.Next
		p.Next = curr.Next
		p = p.Next
	}
	return dummy.Next
}
