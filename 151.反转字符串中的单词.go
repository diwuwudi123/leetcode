/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	s = " " + s + " "
	//l,r 两个指针都从右往左移动
	l, r := len(s)-1, len(s)-1
	strArr := []string{}
	for l >= 0 {
		//左右相等代表左指针要移动了
		if l == r {
			l--
			continue
		}

		if s[l] == ' ' {
			subS := s[l+1 : r]
			if subS != "" {
				strArr = append(strArr, subS)
			}
			r = l
			l--
			continue
		} else {
			l--
		}
	}

	return strings.Join(strArr, " ")
}

// @lc code=end
