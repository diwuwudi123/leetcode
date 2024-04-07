package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxNum := 0
	for right > left {
		m := min(height[left], height[right])
		maxNum = max(m*(right-left), maxNum)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return maxNum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
