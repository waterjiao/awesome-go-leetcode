package _019_Remove_Nth_Node_From_End_of_List

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//定义dummy节点后，先向前移动再遍历
	p1 := &ListNode{Next: head}
	p2 := p1
	p3 := p1
	for i := 0; i < n+1; i++ {
		p2 = p2.Next
	}
	for p2 != nil {
		p2 = p2.Next
		p3 = p3.Next
	}
	p3.Next = p3.Next.Next
	return p1.Next
}
