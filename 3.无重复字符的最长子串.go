package main

import "log"

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	dict := map[byte]int{}
	left, right := 0, 0
	res := 0
	for i := 0; i < len(s); i++ {
		d := s[i]
		dict[d]++
		right++
		for dict[d] > 1 {
			s := s[left]
			dict[s]--
			left++
		}
		res = max(res, right-left)
	}
	return res
}

// @lc code=end
