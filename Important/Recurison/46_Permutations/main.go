package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	//subsets(nums)
	//fmt.Println(mask>>i&1 > 0)
	fmt.Print(permute(nums))
	//fmt.Println(2>>0&1 > 0)
	//fmt.Print(x)
}

var ans [][]int

func permute(nums []int) [][]int {
	begin := make([]int, 0)

	ans = make([][]int, 0)

	used := map[int]bool{}

	dsf(used, begin, nums)

	return ans
}
func dsf(used map[int]bool, set []int, nums []int) {

	if len(set) == len(nums) {
		temp := make([]int, len(set))
		copy(temp, set)
		ans = append(ans, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			set = append(set, nums[i])
			dsf(used, set, nums)
			set = set[:len(set)-1]
			used[i] = false
		}
	}
}

//func permute(nums []int) [][]int {
//
//	begin := make([]int, 0)
//
//	ans = make([][]int, 0)
//
//	used := map[int]bool{}
//
//	//因為不用去除已經找過的,所不用繼續往下找
//	//可以看出元素1在[1,2]中已经使用过了
//	//但是在[2,1]中还要在使用一次1，所以处理排列问题就不用使用startIndex了。???
//	//不懂為什麼不需要晚上回去看影片
//	dfs(begin, nums, used)
//
//	return ans
//}
//func dfs(set []int, nums []int, used map[int]bool) {
//
//	if len(set) == len(nums) {
//		temp := make([]int, len(set))
//
//		copy(temp, set)
//
//		ans = append(ans, temp)
//
//		return
//	}
//
//	for i := 0; i < len(nums); i++ {
//		if !used[i] {
//			used[i] = true
//			set = append(set, nums[i])
//
//			dfs(set, nums, used)
//
//			set = set[:len(set)-1]
//			used[i] = false
//		}
//
//	}
//}
//正確版
// var result [][]int
// func permute(nums []int) [][]int {

// 	begin := make([]int, 0)

// 	result = make([][]int, 0)

// 	visited := map[int]bool{}

// 	backtrack(nums, begin, visited)

// 	return result
// }
// func backtrack(nums,pathNums []int,used map[int]bool){
//     if len(nums)==len(pathNums){
//         tmp:=make([]int,len(nums))
//         copy(tmp,pathNums)
//         result=append(result,tmp)
//         //result=append(result,pathNums)
//         return
//     }
//     for i:=0;i<len(nums);i++{
//         if !used[i]{
//             used[i]=true
//             pathNums=append(pathNums,nums[i])
//             backtrack(nums,pathNums,used)
//             pathNums=pathNums[:len(pathNums)-1]
//             used[i]=false
//         }
//     }
// }

//正確的另一個版本
// func permute(nums []int) [][]int {
// 	res := [][]int{}
// 	visited := map[int]bool{}

// 	var dfs func(path []int)
// 	dfs = func(path []int) {
// 		if len(path) == len(nums) {
// 			temp := make([]int, len(path))
// 			copy(temp, path)
// 			res = append(res, temp)
// 			return
// 		}
// 		for _, n := range nums {
// 			if visited[n] {
// 				continue
// 			}
// 			path = append(path, n)
// 			visited[n] = true
// 			dfs(path)
// 			path = path[:len(path)-1]
// 			visited[n] = false
// 		}
// 	}

// 	dfs([]int{})
// 	return res
// }
