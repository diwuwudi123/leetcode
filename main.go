package main

import (
	"log"
)

func main() {
	s := "dvdf"
	res := lengthOfLongestSubstring(s)
	log.Println(res)
}
func lengthOfLongestSubstring(s string) int {
	right := 0
	left := 0
	charMap := map[string]int{}
	max_len := 0
	for {
		right += 1
		if right > len(s) {
			break
		}
		char := s[right-1 : right]
		if _, ok := charMap[char]; ok {
			charMap[char]++
		} else {
			charMap[char] = 1
		}
		for {
			if charMap[char] > 1 {
				oldChar := s[left : left+1]
				charMap[oldChar]--
				left++
			} else {
				break
			}
		}
		if max_len < right-left {
			max_len = right - left
		}
	}
	return max_len
}
