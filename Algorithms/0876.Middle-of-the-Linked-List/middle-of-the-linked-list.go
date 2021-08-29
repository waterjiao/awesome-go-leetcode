package _876_Middle_of_the_Linked_List

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode2(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func middleNode(head *ListNode) *ListNode {
	length := 1
	p := head
	for p.Next != nil {
		p = p.Next
		length++
	}
	mid := length / 2
	for i := 0; i < mid; i++ {
		head = head.Next
	}
	return head
}
