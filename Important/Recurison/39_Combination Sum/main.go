package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 5}

	target := 8
	//subsets(nums)
	//fmt.Println(mask>>i&1 > 0)
	fmt.Print(combinationSum(nums, target))
	//fmt.Println(2>>0&1 > 0)
	//fmt.Print(x)
}

var ans [][]int

func combinationSum(candidates []int, target int) [][]int {
	begin := make([]int, 0)

	ans = make([][]int, 0)

	dfs(0, 0, begin, candidates, target)

	return ans
}
func dfs(sum int, start int, set []int, nums []int, target int) {

	if sum == target {
		temp := make([]int, len(set))

		copy(temp, set)

		ans = append(ans, temp)

		return
	}

	if sum > target {
		return
	}
	for i := start; i < len(nums); i++ {
		set = append(set, nums[i])

		sum += nums[i]

		dsf(sum, i, set, nums, target)

		sum -= nums[i]

		set = set[:len(set)-1]
	}
}

// Input: candidates = [2,3,5], target = 8
// Output: [[2,2,2,2],[2,3,3],[3,5]]

// func combinationSum(candidates []int, target int) [][]int {
// 	begin := make([]int, 0)

// 	ans = make([][]int, 0)

// 	dfs(0, 0, begin, candidates, target)

// 	return ans
// }
// func dfs(sum int, start int, set []int, nums []int, target int) {

// 	if sum == target {
// 		temp := make([]int, len(set))

// 		copy(temp, set)

// 		ans = append(ans, temp)

// 		return
// 	}
// 	if sum > target {
// 		return
// 	}
// 	for i := start; i < len(nums); i++ {
// 		set = append(set, nums[i])

// 		sum += nums[i]
//              //因為自己本身也可以重複使用,所以要傳i
// 		dfs(sum, i, set, nums, target)

// 		set = set[:len(set)-1]
// 		sum -= nums[i]
// 	}

// }

//正確版
// func combinationSum(candidates []int, target int) [][]int {
// 	var trcak []int
// 	var res [][]int
// 	backtracking(0,0,target,candidates,trcak,&res)
// 	return res
//     }
//     func backtracking(startIndex,sum,target int,candidates,trcak []int,res *[][]int){
// 	//终止条件
// 	if sum==target{
// 	    tmp:=make([]int,len(trcak))
// 	    copy(tmp,trcak)//拷贝
// 	    *res=append(*res,tmp)//放入结果集
// 	    return
// 	}
// 	if sum>target{return}
// 	//回溯
// 	for i:=startIndex;i<len(candidates);i++{
// 	    //更新路径集合和sum
// 	    trcak=append(trcak,candidates[i])
// 	    sum+=candidates[i]
// 	    //递归
// 	    backtracking(i,sum,target,candidates,trcak,res)
// 	    //回溯
// 	    trcak=trcak[:len(trcak)-1]
// 	    sum-=candidates[i]
// 	}

//     }
