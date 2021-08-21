package main

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
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next!= nil && fast.Next.Next != nil{
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var pre, curr *ListNode = nil, head
	for curr != nil {
		temp := curr.Next
		curr.Next = pre
		pre = curr
		curr = temp
	}
	return pre
}

func mergeList(l1, l2 *ListNode) *ListNode {
	curr := l1
	for l1 != nil && l2 != nil {
		temp1 := l1.Next
		temp2 := l2.Next
		l1.Next = l2
		l2.Next = temp1
		l1 = temp1
		l2 = temp2
	}
	return curr
}

func recorder1(head *ListNode) {
	if head == nil {
		return
	}
	mid := middleNode(head)
	l1 := head
	l2 := mid.Next
	mid.Next = nil
	l2 = reverseList(l2)
	head = mergeList(l1, l2)
}


func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
