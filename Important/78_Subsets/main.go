package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	//subsets(nums)
	//fmt.Println(mask>>i&1 > 0)
	fmt.Print(subsets(nums))
	//fmt.Println(2>>0&1 > 0)
	//fmt.Print(x)
}

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// func subsets(nums []int) [][]int {

// 	returnNums := [][]int{{}}
// 	for i, v := range nums {
// 		returnNums[0][i] = append(returnNums[0])

// 	}
// }

//正確版
//迭代法实现子集枚举
// func subsets(nums []int) (ans [][]int) {
// 	//計算nums長度
// 	n := len(nums)

// 	//因為是直接回傳ans和func的ans一樣,會自動偵測到,所以不用再宣告一個ans
// 	//ans = [][]int{}

// 	//1<<n 的意思 (1 * 2 的n次)
// 	//mask表示能夠找出來子集數量
// 	//可以找出2的n次
// 	//index從0開始,循環
// 	for mask := 0; mask < 1<<n; mask++ {

// 		//預設一組set
// 		set := []int{}
// 		fmt.Println("mask:", mask)
// 		//迴圈執行nums中的數字

// 		for i, v := range nums {
// 			//mask>>i 的意思 mask /(除以) (2 的i次)
// 			//如果mask /  (2 的i次) 結果是大於 0
// 			//將數字存入append到set中
// 			fmt.Println(i)
// 			fmt.Println(mask>>i&1 > 0)

// 			if mask>>i&1 > 0 {
// 				//將nums中的數字append到set
// 				//append後
// 				set = append(set, v)
// 			}

// 		}
// 		//fmt.Println([]int(nil))
// 		//fmt.Println(ans)
// 		//ans = [][]int{}
// 		//將上面的append完的set,append到[]int(nil)
// 		//會先變成[1,2,3]
// 		//再append一次到ans[][]
// 		//會變成[[1,2,3][1,2]]
// 		ans = append(ans, append([]int(nil), set...))
// 	}
// 	return
// }

//回溯算法
//宣告一個存結果的變數
// var res [][]int

// //
// func subsets(nums []int) [][]int {
// 	//建立一個存結果的空變數
// 	res = make([][]int, 0)
// 	//將nums作排序
// 	sort.Ints(nums)
// 	//
// 	Dfs([]int{}, nums, 0)
// 	return res
// }

// //計算Dfs的
// func Dfs(temp, nums []int, start int) {
// 	//建立temp長度的array
// 	fmt.Println(start)
// 	tmp := make([]int, len(temp))
// 	fmt.Println(temp)
// 	//將傳入的temp複製到tmp
// 	copy(tmp, temp)
// 	//將tmp append到res
// 	res = append(res, tmp)
// 	//迴圈執行,對傳入的nums做循環
// 	for i := start; i < len(nums); i++ {
// 		//將每個nums append 到temp
// 		temp = append(temp, nums[i])
// 		//再執行一次的nums[i+1]
// 		Dfs(temp, nums, i+1)
// 		//temp取到len(temp)-1的值
// 		temp = temp[:len(temp)-1]
// 	}
// }

func subsets(nums []int) [][]int {
	//宣告一個存結果的變數
	res := [][]int{}
	////宣告一個dfs func
	var dfs func(i int, list []int)
	//建立一個dfs func
	dfs = func(i int, list []int) {
		//如果i等於array長度

		if i == len(nums) {
			//建立list長度的array
			tmp := make([]int, len(list))
			//將傳入的list複製到tmp
			copy(tmp, list)
			//將tmp append到res
			res = append(res, tmp)
			return
		}

		list = append(list, nums[i])
		dfs(i+1, list)

		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})

	return res
}
