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
	res := NsumTarget(nums, 3, 0, 0)

	return res
}
func NsumTarget(nums []int, n, start, target int) [][]int {
	res := [][]int{}

	size := len(nums)
	//处理双数之和
	if n == 2 {
		left := start
		right := size - 1
		for left < right {
			sum := nums[left] + nums[right]
			num_l := nums[left]
			num_r := nums[right]
			if sum < target {
				for left < right && nums[left] == num_l {
					left++
				}
			} else if sum > target {
				for left < right && nums[right] == num_r {
					right--
				}
			} else {
				arr := []int{nums[left], nums[right]}
				res = append(res, arr)
				for left < right && nums[left] == num_l {
					left++
				}
				for left < right && nums[right] == num_r {
					right--
				}
			}
		}
	} else {
		//把n数只和 转化为 n-1 的问题
		for i := 0; i < size; i++ {
			sub := NsumTarget(nums, n-1, i+1, target-nums[i])
			for _, arr := range sub {
				arr = append(arr, nums[i])
				res = append(res, arr)
			}
			for i < size-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res

}

// @lc code=end
