package main

import (
	"log"
)

func main() {
	s := "aBAbaAabBbA"
	t := "bbA"
	res := minWindows(s, t)
	log.Println(res)
}
func minWindows(s, t string) string {
	left, right := 0, 0
	need := map[byte]int{}
	dict := map[byte]int{}
	vaild := 0
	res := ""
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	for right < len(s) {
		code := s[right]
		right++
		if _, ok := need[code]; ok {
			dict[code]++
			if dict[code] == need[code] {
				vaild += dict[code]
			}
		}
		for vaild == len(t) {
			d := s[left]
			if _, ok := need[d]; ok {
				if need[d] == dict[d] {
					if res == "" {
						res = s[left:right]
					} else if len(s[left:right]) < len(res) {
						res = s[left:right]
					}
					vaild -= dict[d]
				}
				dict[d]--
			}
			left++
		}
	}
	return res
}

/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
func minWindow(s string, t string) string {
	left, right := 0, 0
	need := map[byte]int{}
	dict := map[byte]int{}
	valid := 0
	res := ""
	//构造校验字典
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	//只要没到最后就一直遍历
	for right < len(s) {
		//拿到最右的key
		c := s[right]
		right++
		//判断当前key 是否在校验字典里 在的话加到dict里
		// 如果dict key == need key 代表当前key已经完全包含了 校验值+当前key的数量
		if _, ok := need[c]; ok {
			dict[c]++
			if dict[c] == need[c] {
				valid += dict[c]
			}
		}
		//如果校验值== t长度 代表当前字符串已经包含子串了 需要缩小窗口了
		//拿到最左边的key 来判断是不是在校验字典里 在的话先把当前子串存下来 然后缩小左边的窗口
		for valid == len(t) {
			d := s[left]
			if _, ok := need[d]; ok {
				if dict[d] == need[d] {
					if res == "" {
						res = s[left:right]
					} else if len(s[left:right]) < len(res) && res != "" {
						res = s[left:right]
					}
					//因为当前key 在校验字典里 去掉以后就不符合校验了 所以把当前key的数量减去
					valid -= dict[d]
				}
				//减少当前key的值
				dict[d]--
			}
			//收缩左边窗口
			left++
		}
	}
	return res
}

// @lc code=end
