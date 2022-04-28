package main

import "log"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
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
	return ret.Next
}

var arrays []*ListNode

func addNode(num *ListNode) {
	if num == nil {
		return
	}
	arrays = append(arrays, num)
	lenNum := len(arrays)
	if lenNum == 1 {
		return
	}
	swimNode(lenNum - 1)
}
func delNode() *ListNode {
	if len(arrays) == 0 {
		return nil
	}
	val := arrays[0]
	arrays[0], arrays[len(arrays)-1] = arrays[len(arrays)-1], arrays[0]
	arrays = arrays[0 : len(arrays)-1]
	sinkNode(0)
	//把当前这个节点的下一个值添加到数组中
	addNode(val.Next)
	return val
}
func parentNode(num int) int {
	parent := (num - 1) / 2
	log.Println(num, parent, "parent")
	return parent
}
func leftNode(num int) int {
	return 2*num + 1
}
func rightNode(num int) int {
	return 2*num + 2
}
func swimNode(idx int) {
	if idx == 0 {
		return
	}
	//获取当前元素的父节点
	parentNum := parentNode(idx)
	//和父节点对比 如果小于父节点 继续往上对比
	if arrays[parentNum].Val > arrays[idx].Val {
		arrays[parentNum], arrays[idx] = arrays[idx], arrays[parentNum]
		swimNode(parentNum)
	}
}
func sinkNode(idx int) {
	if idx >= len(arrays) {
		return
	}
	//下沉的代码 需要当前节点和左右两个节点分别对比大小 选一个最小的上来
	leftIdx := leftNode(idx)
	if leftIdx >= len(arrays) {
		return
	}
	older := leftIdx
	rightIdx := rightNode(idx)
	if rightIdx < len(arrays) && arrays[older].Val > arrays[rightIdx].Val {
		older = rightIdx
	}
	if arrays[older].Val < arrays[idx].Val {
		arrays[older], arrays[idx] = arrays[idx], arrays[older]
		sinkNode(older)
	}

}

// @lc code=end
