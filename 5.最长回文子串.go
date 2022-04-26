/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	res := ""
	//遍历字符串 然后在以每个i为下标 左右扩展 寻找最长回文字符串
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}
func palindrome(s string, l, r int) string {
	for {
		if l < 0 || r >= len(s) {
			break
		}
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	return s[l+1 : r]
}

// @lc code=end

