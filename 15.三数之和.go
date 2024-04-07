package main

import (
	"log"
	"sort"
)

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4, -1}
	res := threeSum(arr)
	log.Println(res)
}

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := sss(nums, 3, 0, 0)

	return res
}
func NsumTarget(nums []int, n, start, target int) [][]int {
	res := [][]int{}
	size := len(nums)
	if size < 2 || size < n {
		return res
	}
	if n == 2 {
		l, h := start, size-1
		for l < h {
			num := nums[l] + nums[h]
			left, right := nums[l], nums[h]
			//num 小于目标 代表左边的值小了 需要左指针往右移
			if num < target {
				//左下标 小于 右下标 就一直循环
				//左值=右值就跳过
				for l < h && nums[l] == left {
					l++
				}
			} else if num > target {
				for l < h && nums[h] == right {
					h--
				}
			} else {
				//把当前解加到返回结果,并且开始左右指针收缩
				res = append(res, []int{left, right})
				for l < h && nums[l] == left {
					l++
				}
				for l < h && nums[h] == right {
					h--
				}
			}
		}
	} else {
		for i := start; i < size; i++ {
			sub := NsumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			//跳过重复的值
			for i < size-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res

}

// @lc code=end

func sss(nums []int, n, target, start int) [][]int {
	size := len(nums)
	res := [][]int{}

	if n == 2 {
		left := start
		right := size - 1
		for left < right {
			num := nums[left] + nums[right]
			if num < target {
				left++
			} else if num > target {
				right--
			} else {
				res = append(res, []int{nums[left], nums[right]})
				left++
				right--
			}
		}
	} else {
		for i := 0; i < size; i++ {
			subs := sss(nums, n-1, target-nums[i], i+1)
			for _, arr := range subs {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
		}
	}
	return res
}
