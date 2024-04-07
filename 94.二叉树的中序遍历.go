package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var method func(node *TreeNode)
	method = func(node *TreeNode) {
		if node == nil {
			return
		}
		method(node.Left)
		res = append(res, node.Val)
		method(node.Right)
	}
	method(root)
	return res
}

// @lc code=end
