package main

import "log"

func main() {
	a := []int{1, 1, 2, 3, 4, 5, 6, 7, 7}
	res := count(a, 8)
	log.Println(res)
}

func count(arr []int, sum int) int {
	//使用双指针来计算arr里面的和为sum的个数,arr存在重复值
	if len(arr) < 2 {
		return 0
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				res++
			}
		}
	}
	return res

}
