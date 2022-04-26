/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	start := 0
	end := len(str) - 1
	for {
		if start == end || end < start {
			break
		}
		if str[start] == str[end] {
			start++
			end--
			continue
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

