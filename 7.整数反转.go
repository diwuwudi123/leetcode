/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	res := 0
	for {
		mod := x % 10

		if res > 214748364 || (res == 214748364 && mod > 7) {
			return 0
		}
		//判断是否 小于 最小32位整数
		if res < -214748364 || (res == -214748364 && mod < -8) {
			return 0
		}

		x = x / 10

		res = res*10 + mod
		if x == 0 {
			break
		}
	}

	return res
}

// @lc code=end

