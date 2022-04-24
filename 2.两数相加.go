package main

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := false
	l1_len := 0
	l2_len := 0
	var base *ListNode
	var normal *ListNode
	t1 := l1
	t2 := l2
	for {
		if l1.Next != nil {
			l1 = l1.Next
			l1_len += 1
		} else {
			break

		}
	}
	for {
		if l2.Next != nil {
			l2 = l2.Next
			l2_len += 1
		} else {
			break

		}
	}
	if l1_len > l2_len {
		base = t1
		normal = t2
	} else {
		base = t2
		normal = t1
	}
	bas1 := base
	for {
		val2 := 0
		if base == nil {

			break
		}
		if normal == nil {
			if add == false {
				break
			}

		} else {
			val2 = normal.Val
		}
		if add == true {
			base.Val += 1
			add = false
		}

		val1 := base.Val

		val3 := val1 + val2
		if val3 > 9 {
			add = true
			val3 = val3 - 10
		}
		base.Val = val3
		if base.Next == nil && add == true {
			base.Next = &ListNode{
				Val:  1,
				Next: nil,
			}
			add = false
		}
		base = base.Next
		if normal == nil || normal.Next == nil {
			normal = nil
		} else {
			normal = normal.Next

		}
	}
	return bas1
}

// @lc code=end
