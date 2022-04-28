package main

import "log"

var lists = []*ListNode{}

func main() {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	node2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	node3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 6,
		},
	}
	lists = append(lists, node1, node2, node3)
	for _, node := range lists {
		addNode(node)
	}
	res := &ListNode{}
	ret := res
	for {
		val := delNode()
		if val == nil {
			break
		}
		res.Next = val
		res = res.Next
	}
	for {
		if ret == nil {
			break
		}
		log.Println(ret.Val)
		ret = ret.Next
	}
}
