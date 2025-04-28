package main

import "log"

func main() {
	log.Println(climbStairs(5))
}

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	o, t, s := 0, 0, 1
	for i := 1; i <= n; i++ {
		o = t
		t = s
		s = o + t
	}
	return s
}

// @lc code=end
