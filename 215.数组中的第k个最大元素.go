package main

import "log"

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	m := &minHeap{}
	for _, num := range nums {
		m.push(num)

	}
	return (*m)[0]
}
func main() {
	a := []int{12, -34, 11, 67, 15, 63, 25, 17, 1, 56, 33, 28, 99}
	m := &minHeap{}
	for _, num := range a {
		m.push(num)
	}
	log.Println(m)
	for m.len() > 1 {
		m.pop()
		log.Println(m)
	}

}

type minHeap []int

func (m minHeap) len() int {
	return len(m)
}
func (m minHeap) less(i, j int) bool {
	return m[i] < m[j]
}
func (m minHeap) swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *minHeap) push(val int) {
	*m = append(*m, val)
	if m.len() == 1 {
		return
	}
	m.up(m.len() - 1)
}
func (m *minHeap) pop() int {
	if m.len() == 0 {
		return -1
	}
	val := (*m)[0]
	//第一个换到最后一个
	m.swap(0, m.len()-1)
	//把最后一个删除
	*m = (*m)[0 : m.len()-1]
	m.down(0)
	return val
}

func (m minHeap) parent(idx int) int {
	parent := (idx - 1) / 2
	log.Println("parent", idx, parent)
	return parent
}
func (m minHeap) left(idx int) int {
	return 2*idx + 1
}
func (m minHeap) right(idx int) int {
	return 2*idx + 2
}
func (m minHeap) up(idx int) {
	if idx == 0 {
		return
	}
	//获取父节点
	parent := m.parent(idx)
	//当前节点小于父节点 继续上浮
	log.Println("less", m[idx], m[parent])
	if m.less(idx, parent) {
		//和父节点交换 继续上浮
		m.swap(parent, idx)
		m.up(parent)
	}
}

// 下沉
func (m *minHeap) down(idx int) {
	if idx > m.len() {
		return
	}
	left := m.left(idx)
	//左边的节点超过数组最大长度
	if left > m.len()-1 {
		return
	}
	//最小的假设是左边的
	min := left
	right := m.right(idx)
	//找个左右节点最大的节点
	if right < m.len() && (*m)[right] < (*m)[left] {
		min = right
	}
	//如果当前节点比最小节点还要大
	if m.less(min, idx) {
		m.swap(min, idx)
		//还完位置以后让这个值继续下沉
		m.down(min)
	}

}

// @lc code=end
