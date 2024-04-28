package main

import (
	"log"
)

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	res := maxSlidingWindow(nums, k)
	log.Println(res)

}

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	arr := &maxArr{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		arr.push(nums[i])
		if i >= k-1 {
			max := arr.max()
			res = append(res, max)
			arr.pop(nums[i+1-k])
		}
	}
	return res
}

type maxArr []int

func (m *maxArr) push(num int) {
	for len(*m) > 0 && (*m)[len(*m)-1] < num {
		(*m) = (*m)[:len(*m)-1]
	}
	(*m) = append(*m, num)
}
func (m *maxArr) max() int {
	return (*m)[0]
}
func (m *maxArr) pop(num int) {
	if (*m)[0] == num {
		(*m) = (*m)[1:]
	}

}

// @lc code=end

//最简单的办法 两层遍历
// res := []int{}

// for i := 0; i < len(nums)-(k-1); i++ {
// 	maxNum := math.MinInt
// 	for j := 0; j < k; j++ {
// 		maxNum = max(maxNum, nums[j+i])
// 	}
// 	res = append(res, maxNum)
// }
// return res
