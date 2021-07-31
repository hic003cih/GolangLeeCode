package main

import (
	"fmt"
)

func main() {
	//nums := []int{1, 2, 2}

	//subsets(nums)
	//fmt.Println(mask>>i&1 > 0)
	fmt.Print(combine(1, 1))
	//fmt.Println(2>>0&1 > 0)
	//fmt.Print(x)
}

var ans [][]int
func combine(n int, k int) [][]int {

	begin :=make([]int,0)
	ans = make([][]int,0)
	dfs(1,begin,k,n)
	return ans
}
func dfs(start int, set []int, setLen int, rang int) {
	if len(set) ==setLen{
		temp :=make([]int,len(set))

		copy(temp,set)

		ans = append(ans,temp)
	}
	//如果長度
	if len(set)+rang-start+1 < setLen {
		 	return
	}
	for i := start; i <= rang; i++ {
		set = append(set,i)

		dfs(i+1,set,setLen,rang)

		set = set[:len(set)-1]
	}

}

//自己寫的
//var ans [][]int
//
//func combine(n int, k int) [][]int {
//	begin := make([]int, 0)
//
//	ans = make([][]int, 0)
//	//不要空的子集,所以從1開始
//	//dfs(0, begin, k, n)
//	dfs(1, begin, k, n)
//	return ans
//}
//func dfs(start int, set []int, setLen int, rang int) {
//
//	//限制長度為2,當長度為2時,copy到答案
//	if len(set) == setLen {
//		temp := make([]int, len(set))
//
//		copy(temp, set)
//
//		ans = append(ans, temp)
//	}
//
//	//剪枝優化
//	// if len(set)+r-start+1 < setLen {
//	// 	return
//	// }
//
//	//nums := make([]int, 0)
//	//nums = append(nums, r)
//
//
//	//查找的範圍等於限制的n
//	for i := start; i <= rang; i++ {
//		//處理節點
//		set = append(set, i)
//		//循環
//		dfs(i+1, set, setLen, rang)
//		//回朔
//		set = set[:len(set)-1]
//	}
//
//}

//正確版
// var res [][]int

// func combine(n int, k int) [][]int {
// 	res = [][]int{}
// 	if n <= 0 || k <= 0 || k > n {
// 		return res
// 	}
// 	backtrack(n, k, 1, []int{})
// 	return res
// }
// func backtrack(n, k, start int, track []int) {
// 	if len(track) == k {
// 		temp := make([]int, k)
// 		copy(temp, track)
// 		res = append(res, temp)
// 	}
// 	if len(track)+n-start+1 < k {
// 		return
// 	}
// 	for i := start; i <= n; i++ {
// 		track = append(track, i)
// 		backtrack(n, k, i+1, track)
// 		track = track[:len(track)-1]
// 	}
// }

// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]
