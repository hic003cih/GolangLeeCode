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
var ans [][]int

func subsets(nums []int) [][]int {
	ans = [][]int{}
	set := []int{}

	dfs(set, nums, 0)

	return ans
}

func dfs(set []int, nums []int, start int) {
	temp := make([]int, len(set))
	copy(temp, set)
	ans = append(ans, temp)
	if len(set) == len(nums) {
		return
	}
	for i := start; i < len(nums); i++ {
		// fmt.Println(i)
		set = append(set, nums[i])
		// fmt.Println(set)
		dfs(set, nums, i+1)
		//不用
		//i--
		set = set[:len(set)-1]
		// fmt.Println("刪除")
	}
}

// var ans [][]int
// func  subsets(nums []int) [][]int {

// 	set :=[]int{}
// 	ans = [][]int{}
// 	//ans = make([][]int,0)

// 	dsf(0,set,nums)

// 	return ans
// }

// func  dsf(start int,set []int,nums []int)  {

// 	//temp :=[]int{}
// 	temp := make([]int, len(set))
// 	copy(temp,set)
// 	ans = append(ans,temp)
// 	//if start == pth{
// 	//	return
// 	//}

// 	for i := start; i < len(nums); i++ {
// 		set = append(set,nums[i])
// 		dsf(i+1,set,nums)
// 		set = set[:len(set)-1]
// 	}

// }

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

//正確版
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

//官方正確版
//func subsets(nums []int) [][]int {
//	//宣告一個存結果的變數
//	set := []int{}
//	////宣告一個dfs func
//	var dfs func(i int, list []int)
//	//建立一個dfs func
//	dfs = func(i int, list []int) {
//
//		//fmt.Println("第幾次呼叫", i)
//
//		//fmt.Println(list)
//
//		//如果i等於array長度
//
//		//當dfs func循環的長度等於array的長度時(當是1,2,3時)
//		//把子集存到答案中,跳出當前循環
//		//並將子集中的尾數去除(當是1,2,3時,去除3,變回1,2)
//		if i == len(nums) {
//			//建立list長度的array
//			tmp := make([]int, len(list))
//			//將子集存入tmp,給後面append到答案中用
//			copy(tmp, list)
//			//append到答案中用
//			res = append(res, tmp)
//			//跳出該迴圈去執行子集中的尾數去除,回到上層繼續查找可能子集
//			return
//		}
//		//將array中的數字加到子集中
//		list = append(list, nums[i])
//		//重複呼叫自己本身,一直往下找有可能的子集(ex: 1往下找1,2,在往下找1,2,3)
//		dfs(i+1, list)
//
//		//將子集中的尾數去除(當是1,2,3時,去除3,變回1,2)
//		list = list[:len(list)-1]
//		//去除以後繼續進入dfs func循環中
//		//因為去除尾數就代表回到上一層(從1,2開始),繼續找子集
//		//但傳入的次數會繼續,例如:上面的i是3,跳出後,從1,2,3去掉最尾變1,2,但跳出後這邊的i是2,所以下面的i+1是3
//		//執行dfs後又會馬上跳出,1,2去掉最尾變1,跳出後這邊的i是1,所以下面的i+1是2
//		//執行dfs後,變成1,3,後面又在執行dfs,i+1是3,跳出,1,3去掉最尾變1
//		dfs(i+1, list)
//	}
//	//執行第一次的dfs func,傳入一個空子集
//	dfs(0, []int{})
//
//	return res
//}
