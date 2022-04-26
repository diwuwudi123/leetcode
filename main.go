package main

import "log"

func main() {
	input := 123
	output := reverse(input)
	log.Println(output)
	input = -123
	output = reverse(input)
	log.Println(output)
}
func reverse(x int) int {

	res := 0
	pre := 1
	for {
		mod := x % 10

		if res > 214748364 || (res == 214748364 && mod > 7) {
			return 0
		}
		//判断是否 小于 最小32位整数
		if res < -214748364 || (res == -214748364 && mod < -8) {
			return 0
		}

		x = x / 10

		res = res*10 + mod
		if x == 0 {
			break
		}
		pre = pre * 10
	}

	return res
}
