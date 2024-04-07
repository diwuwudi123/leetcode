package main

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {

	res := [][]int{}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		//获取当前节点
		cur := intervals[i]
		last := res[len(res)-1]
		if cur[0] <= last[1] {
			last[1] = max(last[1], cur[1])
		} else {
			res = append(res, cur)
		}

	}
	return res
}

// @lc code=end
