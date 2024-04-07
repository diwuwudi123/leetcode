package main

import "log"

func main() {
	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i-1] * res[i-1]
	}
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = r * res[i]
		r *= nums[i]
	}
	return res
}

// @lc code=end
