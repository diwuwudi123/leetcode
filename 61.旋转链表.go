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
	log.Println(listLen)
	//判断实际需要移动的位置
	if k > listLen {
		k = k % listLen
	} else if k == listLen {
		return head
	}
	log.Println(k)
	//如果k==0 代表不需要移动
	if k == 0 {
		return head
	}
	//如果k >0 位移实际的长度
	// 1-2-3-4-5
	// 右移两位 其实约等于 把链表拆分为 1-2-3   4-5 然后把head->4  把5->1
	oldHead := head
	newhead := &ListNode{}
	idx := 1
	//从指定位置把链表拆分两个
	//需要idx = 链表长度 - 位移的数量比如 链表长度5 位移2位 就是从第三位开始断开链表
	for head.Next != nil {
		if idx == listLen-k {
			newhead = head.Next
			head.Next = nil
			continue
		}
		idx++
		head = head.Next
	}
	head = newhead
	for {
		if newhead.Next == nil {
			newhead.Next = oldHead
			return head
		}
		newhead = newhead.Next

	}
	return head
}

// @lc code=end
