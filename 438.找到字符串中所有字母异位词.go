package main

import "log"

func main() {
	s := "cbaebabacd"
	p := "abc"
	res := findAnagrams(s, p)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	need, dict := map[byte]int{}, map[byte]int{}
	valid := 0
	left, right := 0, 0
	res := []int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	for right < len(s) {
		c := s[right]

		if _, ok := need[c]; ok {

			dict[c]++
			if dict[c] == need[c] {
				valid++
			}
		}
		if right-left >= len(p)-1 {
			if valid == len(p) {
				res = append(res, left)
			}
			d := s[left]

			left++
			if _, ok := need[d]; ok {
				if dict[d] == need[d] {
					valid--
				}
				dict[d]--
			}
		}
		log.Println(string(c), right-left, valid, left, right)
		right++

	}
	return res
}

// @lc code=end
