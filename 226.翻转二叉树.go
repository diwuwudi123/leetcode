/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	revers(root)
	return root
}
func revers(node *TreeNode) {
	if node == nil {
		return
	}
	tmp := node.Left
	node.Left = node.Right
	node.Right = tmp
	revers(node.Left)
	revers(node.Right)
}

// @lc code=end

