/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	i := 0
	flag := -1
	for idx := 0; idx < len(s); idx++ {
		if i == numRows-1 || i == 0 {
			flag = 0 - flag
		}

		res[i] += s[idx : idx+1]
		i += flag
	}
	return strings.Join(res, "")
}

// @lc code=end

