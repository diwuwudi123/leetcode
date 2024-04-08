package main

import "log"

func main() {
	nums := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 7}
	res := removeDuplicates(nums)
	log.Println("remove duplicates", nums[0:res])
}

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	slow++
	return slow
}

// @lc code=end
