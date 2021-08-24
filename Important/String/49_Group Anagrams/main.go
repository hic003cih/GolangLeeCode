package main

import (
	"fmt"
	"sort"
)

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

func main() {
	nums := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	x := groupAnagrams(nums)
	fmt.Print(x)
}
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}

	ans := [][]string{}

	for _, v := range strs {
		s := []byte(v)
		sort.Slice(s, func(i int, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], v)
	}
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// func groupAnagrams(strs []string) [][]string {

// 	mp := map[string][]string{}

// 	ans := [][]string{}

// 	for _, str := range strs {
// 		//將字串轉為byte
// 		s := []byte(str)
// 		//使用sort.Slice方法,將array傳入,比大小作排序
// 		sort.Slice(s, func(i, j int) bool {
// 			return s[i] < s[j]
// 		})
// 		//將排序過的字母byte轉回string
// 		sortedStr := string(s)
// 		//將排序過一樣的字串,存到同一個索引map
// 		//排過的sortedStr當索引,原始的str當值
// 		mp[sortedStr] = append(mp[sortedStr], str)
// 		fmt.Println(mp[sortedStr])
// 	}
// 	//循環將map的值存到ans
// 	for _, v := range mp {
// 		ans = append(ans, v)
// 	}
// 	return ans
// }

//正確版
// func groupAnagrams(strs []string) [][]string {

// 	mp := map[string][]string{}
// 	for _, str := range strs {
// 		s := []byte(str)
// 		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
// 		sortedStr := string(s)
// 		mp[sortedStr] = append(mp[sortedStr], str)
// 	}
// 	ans := make([][]string, 0, len(mp))
// 	for _, v := range mp {
// 		ans = append(ans, v)
// 	}
// 	return ans
// }
