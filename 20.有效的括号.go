package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	kMap := map[byte]byte{']': '[', '}': '{', ')': '('}
	kArr := []byte{}
	for i := 0; i < len(s); i++ {
		key := s[i]
		if len(kArr) == 0 {
			kArr = append(kArr, key)
			continue
		}
		//如果key 是右括号
		if _, ok := kMap[key]; ok {
			//获取左边的括号看看是不是匹配
			pre := kArr[len(kArr)-1]
			if pre == kMap[key] {
				//匹配pop掉左括号
				kArr = kArr[:len(kArr)-1]
				continue
			} else {
				//不匹配 代表有问题
				return false
			}
		} else {
			kArr = append(kArr, key)
		}

	}
	return len(kArr) == 0
}

// @lc code=end
