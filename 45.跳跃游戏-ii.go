package main

import "log"

func main() {
	log.Println(jump([]int{1, 2, 1, 1, 1}))
}

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	maxIdx := 0
	//本次跳跃最大重点 当前i == end 代表本次跳跃计算结束 需要进行下次跳跃计算了
	end := 0
	//跳跃次数
	jump := 0
	//计算每个各自最大可跳跃距离
	for i := 0; i < len(nums)-1; i++ {
		//本次最大跳跃距离
		maxIdx = max(i+nums[i], maxIdx)
		if i == end {
			end = maxIdx
			jump++
		}

	}
	return jump
}

// @lc code=end
