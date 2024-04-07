package main

import "log"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	log.Println(nums)
}

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	if k == 0 || len(nums) < 2 || k%len(nums) == 0 {
		return
	}
	tmp := []int{}
	size := len(nums)
	k = k % size
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		if k > len(tmp) {
			nums[i] = nums[size-k+i]
		} else {
			nums[i] = tmp[0]
			tmp = tmp[1:]
		}
		tmp = append(tmp, t)
	}

	// copy(nums, append(nums[size-k:], nums[:size-k]...))

	// a := nums[size-k:]
	// b := nums[:size-k]
	// c := append(a, b...)
	// for i := 0; i < size; i++ {
	// 	nums[i] = c[i]
	// }

}

// @lc code=end
