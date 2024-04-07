package main

import (
	"log"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
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
	res := []int{}
	arr := &maxArr{}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			arr.push(nums[i])
		} else {
			arr.push(nums[i])

			max := arr.max()
			//窗口滑动 往右移 最左边的会弹出来
			arr.pop(nums[i-k+1])
			res = append(res, max)
		}

	}
	return res
}

type maxArr struct {
	arr []int
}

func (m *maxArr) push(k int) {
	//遍历 把所有小于k的都删掉
	for len(m.arr) > 0 && m.arr[len(m.arr)-1] < k {
		m.arr = m.arr[:len(m.arr)-1]
	}
	m.arr = append(m.arr, k)
}
func (m *maxArr) max() int {
	return m.arr[0]
}
func (m *maxArr) pop(k int) {
	if k == m.arr[0] {
		m.arr = m.arr[1:]
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
