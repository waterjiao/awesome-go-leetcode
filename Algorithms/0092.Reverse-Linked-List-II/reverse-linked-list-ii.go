package _092_Reverse_Linked_List_II

import . "waterjiao.com/awesome-go-leetcode/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	curr := head
	index := 1
	var p *ListNode
	for curr != nil {
		if index == left {
			for index <= right {
				temp := curr.Next
				curr.Next = p
				p = curr
				curr = temp
				index++
			}
			pre.Next.Next = curr
			pre.Next = p
		} else {
			pre = pre.Next
			curr = curr.Next
			index++
		}
	}
	return dummy.Next
}

// 一次过，记录下思路，引入dummy节点，再引入pre指针指向dummy，curr指针指向head
// 当index值与left相等时，说明此时curr往后的链表，到right下标时，应该要翻转
// 引入单链表翻转，先引入一个空指针p，引入temp指向curr的下一个，curr的下一个指向p，p移动到curr，curr移动到temp
// 此时单链表翻转完，将pre的下一个指向p，pre的下一个的下一个指向curr，这里不好理解的时pre的next的next，涉及到单链表翻转后
// 与原来的链表还没有断链，这里需要注意
