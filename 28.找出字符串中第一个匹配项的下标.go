package main

func main() {
	s := "hello world"
	subS := "world"
	println(strStr(s, subS))
}

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 找出字符串中第一个匹配项的下标
 */

// @lc code=start
func strStr(haystack string, needle string) int {

	l, r := 0, 0
	subIdx := 0
	for r < len(haystack) {
		//如果右指针等于needle当前字符串 代表匹配 同时右移
		if haystack[r] == needle[subIdx] {
			r++
			subIdx++
		} else {
			//不匹配的话 l右移 r改为l
			l++
			r = l
			subIdx = 0
		}
		//如果 idx 匹配了长度 代表完整匹配 返回当前l
		if subIdx == len(needle) {
			return l
		}

	}
	return -1
}

// @lc code=end
