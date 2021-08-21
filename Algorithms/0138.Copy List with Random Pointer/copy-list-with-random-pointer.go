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
	origin, nodes := []*Node{}, []*Node{}
	random := make(map[*Node]int)
	for curr := head; curr != nil; curr = curr.Next {
		origin = append(origin, curr)
		nodes = append(nodes, &Node{Val: curr.Val})
	}
	for _, c := range origin {
		for index, q := range origin {
			if c.Random == q {
				random[c] = index
				break
			} else if c.Random == nil {
				random[c] = -1
				break
			}
		}
	}
	for i:= 0; i< len(origin)-1; i++ {
		nodes[i].Next = nodes[i+1]
		r := origin[i].Random
		v := random[r]
		if v == -1 {
			nodes[i].Random = nil
			continue
		}
		nodes[i].Random = nodes[v]
	}
	r := origin[len(origin)-1].Random
	v := random[r]
	if v == -1 {
		nodes[len(origin)-1].Random = nil
	} else {
		nodes[len(origin)-1].Random = nodes[v]
	}
	nodes[len(origin)-1].Next = nil
	return nodes[0]
}
