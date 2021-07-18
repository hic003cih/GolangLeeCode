package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}

	//subsets(nums)
	//fmt.Println(mask>>i&1 > 0)
	fmt.Print(subsetsWithDup(nums))
	//fmt.Println(2>>0&1 > 0)
	//fmt.Print(x)
}

var ans[][]int

func subsetsWithDup(nums []int) [][]int {

	beging :=make([]int,0)

	sort.Ints(nums)

	ans =make([][]int,0)

	dsf(0,beging,nums)

	return ans
}

func dsf(start int,set []int,nums []int)  {

		temp :=make([]int,len(set))

		copy(temp,set)

		ans=append(ans,temp)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i]==nums[i-1]{
			continue
		}
		set = append(set,nums[i])

		dsf(i+1,set,nums)

		set = set[:len(set)-1]
	}



}












//var ans [][]int
//
//func subsetsWithDup(nums []int) [][]int {
//
//	begin := make([]int, 0)
//
//	ans = make([][]int, 0)
//
//	sort.Ints(nums)
//
//	dfs(begin, 0, nums)
//	return ans
//}
//
//func dfs(set []int, start int, nums []int) {
//
//	temp := make([]int, len(set))
//
//	copy(temp, set)
//
//	ans = append(ans, temp)
//
//	for i := start; i < len(nums); i++ {
//		if i > start && nums[i] == nums[i-1] {
//			continue
//		}
//		set = append(set, nums[i])
//		dfs(set, i+1, nums)
//
//		set = set[:len(set)-1]
//	}
//
//}

//正確版
//回溯算法
//建立答案子集res
//var res [][]int
//
//func subsetsWithDup(nums []int) [][]int {
//	//建立一個答案子集res
//	res = make([][]int, 0)
//	//將array排序
//	sort.Ints(nums)
//	//執行第一次dfs
//	dfs([]int{}, nums, 0)
//	return res
//}
//func dfs(temp, num []int, start int) {
//	//建立一個長度為已取出的array數量的長度
//	tmp := make([]int, len(temp))
//	//將以取出的已取出的array複製到tmp
//	copy(tmp, temp)
//
//	fmt.Println(temp)
//	//將tmp存到res答案子集中
//	res = append(res, tmp)
//	//迴圈尋找array中的每個數
//	for i := start; i < len(num); i++ {
//		//如果i >不是自己本身(start),然後又等於上一層的數字(num[i-1])
//		//退出本次迴圈直接執行下次的迴圈
//		if i > start && num[i] == num[i-1] {
//			//退出本次迴圈直接執行下次的迴圈
//			continue
//		}
//		//將迴圈在array中數字存到temp中
//		temp = append(temp, num[i])
//		//再將temp去執行dfs,繼續往下一層找
//		dfs(temp, num, i+1)
//		//每個比對完以後沒有值
//		//將temp的最後一個位拿掉
//		//回到上一層繼續比對
//
//		fmt.Println("dfs下一層後的", temp)
//		temp = temp[:len(temp)-1]
//	}
//}

//自己寫的
// func subsetsWithDup(nums []int) [][]int {
// 	//宣告一個存結果的變數
// 	ans := [][]int{}

// 	sort.Ints(nums)

// 	dsf(0, []int{}, nums, ans)
// 	return ans
// }
// func dsf(i int, list []int, nums []int, ans [][]int) (list2 []int, ans2 [][]int) {

// 	n := len(nums)
// 	if i == n {
// 		set := make([]int, n)
// 		copy(set, list)
// 		ans := append(ans, set)
// 		return list, ans
// 	}
// 	list = append(list, nums[i])

// 	dsf(i+1, list, nums, ans)

// 	list = list[:len(list)-1]

// 	dsf(i+1, list, nums, ans)

// 	return list, ans
// }

//正確法
// var res[][]int
// func subsetsWithDup(nums []int)[][]int {
// 	res=make([][]int,0)
// 	 sort.Ints(nums)
// 	dfs([]int{},nums,0)
// 	return res
// }
// func dfs(temp, num []int, start int)  {
// 	tmp:=make([]int,len(temp))
// 	copy(tmp,temp)

// 	res=append(res,tmp)
// 	for i:=start;i<len(num);i++{
// 		if i>start&&num[i]==num[i-1]{
// 			continue
// 		}
// 		temp=append(temp,num[i])
// 		dfs(temp,num,i+1)
// 		temp=temp[:len(temp)-1]
// 	}
// }
