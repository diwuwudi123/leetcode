package main

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	list := []*TreeNode{root}
	res := [][]int{}
	if root == nil {
		return res
	}
	for len(list) > 0 {
		sz := len(list)
		level := []int{}
		for i := 0; i < sz; i++ {
			cur := list[0]
			list = list[1:]

			level = append(level, cur.Val)
			if cur.Left != nil {
				list = append(list, cur.Left)
			}
			if cur.Right != nil {
				list = append(list, cur.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// @lc code=end
