package main

import "log"

func main() {
	arr := []int{1, 0, 5, 0, 5, 2, 3, 9, 0, 3}
	moveZeroes(arr)
	log.Println(arr)
}

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}

}

// @lc code=end
