package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	res := rotateRight(list, 2)
	for {
		log.Println(res.Val)
		if res.Next == nil {

			break
		}
		res = res.Next
	}
}

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	//向右旋转链表
	//如果k小于5 右移实际的长度
	//如果k> 5  右移实际长度= k%len(head)
	if head == nil || head.Next == nil {
		return head
	}
	listLen := 1
	list := head
	//计算链表长度
	for list.Next != nil {
		listLen++
		list = list.Next
	}
	//判断实际需要移动的位置
	if k > listLen {
		k = k % listLen
	} else if k == listLen {
		return head
	}
	//如果k==0 代表不需要移动
	if k == 0 {
		return head
	}
	//如果k >0 位移实际的长度
	// 1-2-3-4-5
	// 右移两位 其实约等于 把链表拆分为 1-2-3   4-5 然后把head->4  把5->1
	idx := listLen - k
	//把链表变成环 最后一个连到head上然后开始计算需要断开的位置 就直接断开
	list.Next = head
	for idx > 0 {
		list = list.Next
		idx--

	}
	//从这个位置断开链接
	newHead := list.Next
	list.Next = nil

	return newHead
}

// @lc code=end
