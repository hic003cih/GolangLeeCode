package main

import (
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(strs))
}
func longestCommonPrefix(strs []string) (prefix string) {
	strlen := len(strs)
	if strlen == 0 {
		return ""
	}
	var index int
	prefix = strs[0]
	for index < strlen {
		prefix = findPrefix(prefix, strs[index])
		if len(prefix) < 0 {
			break
		}
		index++
	}
	return prefix
}
func findPrefix(prefix string, str1 string) (findPrefix string) {

	min := findMin(prefix, str1)

	var index int

	for (index < len(min)) && (prefix[index] == str1[index]) {
		index++
	}
	return min[:index]
}

func findMin(prefix string, str1 string) (minStr string) {
	if len(str1) > len(prefix) {
		return prefix
	}
	return str1
}

//正確版
// func longestCommonPrefix(strs []string) string {

// 	//如果傳入array長度等於0,直接回傳空
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	//先預設array的第一個為prefix
// 	prefix := strs[0]
// 	//計算array長度
// 	count := len(strs)
// 	//對array進行迴圈,從第二個開始
// 	for i := 1; i < count; i++ {
// 		//進行prefix的比對,並回傳prefix
// 		prefix = lcp(prefix, strs[i])

// 		//如果prefix長度是0,跳出迴圈
// 		if len(prefix) == 0 {
// 			break
// 		}
// 	}
// 	return prefix
// }

// //進行prefix的比對,並回傳prefix
// func lcp(str1, str2 string) string {
// 	//比對字串的長度
// 	//取出最短的字串來當
// 	length := min(len(str1), len(str2))
// 	index := 0
// 	//如果兩個字串的index的字母相同
// 	//則index+1繼續比對
// 	//然後迴圈比對到最小的字串長度結束
// 	for index < length && str1[index] == str2[index] {
// 		index++
// 	}
// 	//返回最後比對出來的prefix,只取到index以前的字母
// 	return str1[:index]
// }

// //比對字串大小
// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
