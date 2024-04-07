package main

import "log"

var array []int

func add(num int) {
	array = append(array, num)
	lenNum := len(array)
	if lenNum == 1 {
		return
	}
	swim(lenNum - 1)
}
func del() int {
	if len(array) == 0 {
		return -1
	}
	val := array[0]
	array[0], array[len(array)-1] = array[len(array)-1], array[0]
	array = array[0 : len(array)-1]
	sink(0)
	return val
}
func parent(num int) int {
	parent := (num - 1) / 2
	log.Println(num, parent, "parent")
	return parent
}
func left(num int) int {
	return 2*num + 1
}
func right(num int) int {
	return 2*num + 2
}
func swim(idx int) {
	if idx == 0 {
		return
	}
	//获取当前元素的父节点
	parentNum := parent(idx)
	//和父节点对比 如果小于父节点 继续往上对比
	if array[parentNum] > array[idx] {
		array[parentNum], array[idx] = array[idx], array[parentNum]
		swim(parentNum)
	}
}
func sink(idx int) {
	if idx >= len(array) {
		return
	}
	//下沉的代码 需要当前节点和左右两个节点分别对比大小 选一个最小的上来
	leftIdx := left(idx)
	if leftIdx >= len(array) {
		return
	}
	older := leftIdx
	rightIdx := right(idx)
	if rightIdx < len(array) && array[older] > array[rightIdx] {
		older = rightIdx
	}
	if array[older] < array[idx] {
		array[older], array[idx] = array[idx], array[older]
		sink(older)
	}

}
