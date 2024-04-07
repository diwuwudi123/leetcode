package main

import (
	"log"
	"math"
)

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	res := maxSubArray(arr)
	log.Println("res", res)

}

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	arr := make([]int, len(nums))
	arr[0] = nums[0]
	res := math.MinInt
	for i := 1; i < len(nums); i++ {
		arr[i] = max(nums[i], nums[i]+arr[i-1])
	}
	log.Println(arr)
	for _, num := range arr {
		res = max(res, num)
	}
	return res
}

// @lc code=end
