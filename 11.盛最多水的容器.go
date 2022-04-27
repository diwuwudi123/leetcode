/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	res := 0
	for {
		if left >= right {
			break
		}
		min := getMinByTwoNum(height[left], height[right])
		res = getMaxByTwoNum(res, min*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
func getMinByTwoNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func getMaxByTwoNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

