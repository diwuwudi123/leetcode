package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode

	cur := head
	for cur != nil {
		//拿到下一个节点
		nextTmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTmp
	}
	return pre
}

// @lc code=end
