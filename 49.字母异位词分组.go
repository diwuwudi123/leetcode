package main

import "log"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {

	res := map[[26]int][]string{}
	for _, str := range strs {
		code := encode(str)
		res[code] = append(res[code], str)

	}
	resArr := [][]string{}
	for _, val := range res {
		resArr = append(resArr, val)
	}
	return resArr
}
func encode(s string) [26]int {
	kmap := [26]int{}

	for i := 0; i < len(s); i++ {

		idx := s[i] - 'a'
		kmap[idx]++

	}
	return kmap
}

// @lc code=end
