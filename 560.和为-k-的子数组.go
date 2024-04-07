package main

import "log"

func main() {
	num := []int{1, 1, 1, 1}
	k := 2
	res := subarraySum(num, k)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		sum := 0
		for end := start; end < len(nums); end++ {
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// @lc code=end
