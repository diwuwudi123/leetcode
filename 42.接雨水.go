package main

import (
	"log"
)

func main() {
	arr := []int{3, 0, 3}
	// arr := []int{4, 2, 0, 3, 2, 5}

	res := trap(arr)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	size := len(height)
	l_max := make([]int, size)
	r_max := make([]int, size)
	res := 0
	l_max[0] = height[0]
	r_max[size-1] = height[size-1]
	for i := 1; i < size; i++ {
		l_max[i] = max(height[i], l_max[i-1])
	}
	for i := size - 2; i > 0; i-- {
		r_max[i] = max(height[i], r_max[i+1])
	}
	for i := 1; i < size; i++ {
		res += min(l_max[i], r_max[i]) - height[i]
	}
	return res

}

// @lc code=end

// 时间On² 空间o1
// size := len(height)
// res := 0
// for i := 1; i < size-1; i++ {
// 	l_max, r_max := 0, 0
// 	for j := 0; j < i; j++ {
// 		l_max = max(l_max, height[j])
// 	}
// 	for j := size - 1; j > i; j-- {
// 		r_max = max(r_max, height[j])
// 	}
// 	if l_max < height[i] || r_max < height[i] {
// 		continue
// 	}
// 	res += min(l_max, r_max) - height[i]
// }
// return res

// 时间复杂度on 空间复杂度on
// size := len(height)
// res := 0
// l_max := make([]int, size)
// r_max := make([]int, size)
// l_max[0] = height[0]
// r_max[size-1] = height[size-1]
// for i := 1; i < size; i++ {
// 	l_max[i] = max(height[i], l_max[i-1])
// }
// for i := size - 2; i >= 0; i-- {
// 	r_max[i] = max(height[i], r_max[i+1])
// }
// for i := 1; i < size; i++ {
// 	res += min(l_max[i], r_max[i]) - height[i]
// }
// return res

//时间 on 空间o1
// res := 0
// left, right := 0, len(height)-1
// l_max := 0
// r_max := 0
// for left < right {
// 	l_max = max(l_max, height[left])
// 	r_max = max(r_max, height[right])
// 	if l_max < r_max {
// 		res += l_max - height[left]
// 		left++
// 	} else {
// 		res += r_max - height[right]
// 		right--
// 	}
// }
// return res
