package main

import "log"

func main() {
	nums := []int{-1, 4, 2, 1, 9, 10}
	res := firstMissingPositive(nums)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
func abs(a int) int {
	if a < 0 {
		return 0 - a
	}
	return a
}

/**
	size := len(nums)

	for i := 0; i < size; i++ {
		if nums[i] <= 0 {
			nums[i] = size + 1
		}
	}

	for i := 0; i < size; i++ {
		if abs(nums[i]) <= size {
			idx := abs(nums[i]) - 1
			nums[idx] = 0 - abs(nums[idx])
		}
	}

	for i := 0; i < size; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return size + 1

**/

// @lc code=end
