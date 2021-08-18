package _021_Merge_Two_Sorted_Lists

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

func mergeTwoLists_(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists_(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists_(l1, l2.Next)
		return l2
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newl := &ListNode{}
	p := newl
	p1 := l1
	p2 := l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p = p2
			p2 = p2.Next
		}
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return newl.Next
}
