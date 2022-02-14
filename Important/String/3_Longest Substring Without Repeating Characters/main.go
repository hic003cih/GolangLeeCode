package main

import (
	"fmt"
)

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

func main() {
	nums := "dvdf"
	x := lengthOfLongestSubstring(nums)
	fmt.Print(x)
}

// func lengthOfLongestSubstring(s string) int {
// 	//如果小於1直接返回0
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	if len(s) == 1 {
// 		return 1
// 	}
// 	hashMap := map[byte]int{}
// 	//儲存答案的ans
// 	ans := 0
// 	//給新的答案的temp
// 	temp := 0
// 	//將string 轉換成byte,方便迴圈使用
// 	byteString := []byte(s)
// 	//循環整個String
// 	for i, v := range byteString {
// 		//如果字母已經有存在當前的map
// 		//就清空temp和map,重新計算新的子串
// 		if _, ok := hashMap[v]; ok {
// 			//新的答案的temp的長度又大於舊的ans
// 			//就替代舊的ans
// 			if temp > ans {
// 				ans = temp
// 			}
// 			//將存新的答案的temp青空
// 			temp = 0
// 			//將map清空
// 			hashMap = map[byte]int{}
// 		}
// 		//如果在map沒有重複的,則將temp長度加一
// 		temp++
// 		fmt.Println(string(v))
// 		//並將該字母存入map中,用來檢查是否重複
// 		hashMap[v] = i
// 	}
// 	if temp > 0 && temp > ans {
// 		ans = temp
// 	}
// 	return ans
// }
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}

	left, right, ans := 0, 0, 0
	//
	m := map[byte]int{}
	for right < strLen {
		if _, ok := m[s[right]]; !ok {
			m[s[right]] = right
		} else {
			if m[s[right]]+1 >= left {
				left = m[s[right]] + 1
			}
			m[s[right]] = right
		}
		ans = max(right-left+1, ans)
		right++
	}

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// //正式版
// func lengthOfLongestSubstring(s string) int {
// 	// 哈希集合，记录每个字符是否出现过
// 	m := map[byte]int{}
// 	n := len(s)
// 	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
// 	rk, ans := -1, 0
// 	for i := 0; i < n; i++ {
// 		fmt.Println(i)
// 		if i != 0 {
// 			// 左指针向右移动一格，移除一个字符
// 			delete(m, s[i-1])
// 		}
// 		for rk+1 < n && m[s[rk+1]] == 0 {
// 			fmt.Println(m[s[rk+1]])
// 			// 不断地移动右指针
// 			m[s[rk+1]]++
// 			rk++
// 		}
// 		// 第 i 到 rk 个字符是一个极长的无重复字符子串
// 		ans = max(ans, rk-i+1)
// 	}
// 	return ans
// }

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }
