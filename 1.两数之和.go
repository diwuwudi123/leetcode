package main

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	numsMap := map[int][]int{}
	for idx, v := range nums {
		if _, ok := numsMap[v]; ok {
			numsMap[v] = append(numsMap[v], idx)
		} else {
			numsMap[v] = []int{idx}
		}
	}
	for k, v := range numsMap {
		numA := target - k
		if arr, ok := numsMap[numA]; ok {
			if numA == k && len(arr) >= 2 {
				return arr[0:2]
			}
			if numA != k {
				return []int{v[0], arr[0]}
			}
			continue
		}
	}
	return []int{}
}

// @lc code=end
