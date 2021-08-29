package problem0206

type ListNode struct {
	value string
	next  *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		next := curr.next
		curr.next = pre
		pre = curr
		curr = next
	}
	return pre
}
