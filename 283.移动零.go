/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	s, f := 0, 0
	for f < len(nums) {
		if nums[f] != 0 {
			nums[s] = nums[f]
			s++
		}
		f++
	}
	for s < len(nums) {
		nums[s] = 0
		s++
	}
	return
}

// @lc code=end

