package main

import (
	"fmt"
	"sort"
)

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	x := threeSum(nums)
	fmt.Print(x)
}
func threeSum(nums []int) [][]int {
	//先排序
	sort.Ints(nums)
	//生成一個map
	res := [][]int{}
	//迴圈比對全部數
	for i := 0; i < len(nums); i++ {

		//先選定一個數為n1
		n1 := nums[i]
		// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
		if n1 > 0 {
			break
		}
		//如果i不是index第一個數,並且值和前一個值一樣,則繼續執行
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		//選定好n1後,設定左指針和右指針並取出值做比對
		//左指針設定為i+1
		//右指針設定為總長度減1(設定在最右邊)
		left, right := i+1, len(nums)-1
		//持續比對
		for left < right {
			// 去重复逻辑如果放在这里，0，0，0 的情况，可能直接导致 right<=left 了，从而漏掉了 0,0,0 这种三元组
			/*
			   while (right > left && nums[right] == nums[right - 1]) right--;
			   while (right > left && nums[left] == nums[left + 1]) left++;
			*/
			//把右指針和左指針的值取出來
			n2, n3 := nums[left], nums[right]
			//如果三個值加起來為0
			//將三個值加入res矩陣中
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 因為不能有重複的,所以需要去重
				// 有去重逻辑应该放在找到一个三元组之后
				//如果右指針大於左指針,並且左指針的值等於當前的值,則左指針往右移
				for left < right && nums[left] == n2 {
					left++
				}
				//如果右指針大於左指針,並且右指針的值等於當前的值,則右指針往左移
				for left < right && nums[right] == n3 {
					right--
				}
				//如果三個數加起來小於0,則左指針往右移,繼續尋找
			} else if n1+n2+n3 < 0 {
				left++
				//其他情況則右指針往內縮
			} else {
				right--
			}

		}
	}
	return res
}

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	res := [][]int{}
// 	for i := 0; i < len(nums); i++ {
// 		n1 := nums[i]
// 		if n1 > 0 {
// 			break
// 		}
// 		if i > 0 && n1 == nums[i-1] {
// 			continue
// 		}
// 		left, right := i+1, len(nums)-1
// 		for left < right {
// 			n2, n3 := nums[left], nums[right]
// 			if n1+n2+n3 == 0 {
// 				res = append(res, []int{n1, n2, n3})
// 				for left < right && nums[left] == n2 {
// 					left++
// 				}
// 			} else if n1+n2+n3 < 0 {
// 				left++
// 			} else {
// 				right--
// 			}
// 		}
// 	}
// 	return res
// }
//
//
// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	res := [][]int{}
// 	for i := 0; i < len(nums); i++ {

// 		n1 := nums[i]
// 		if n1 > 0 {
// 			break
// 		}
// 		if i > 0 && n1 == nums[i-1] {
// 			continue
// 		}
// 		left, right := i+1, len(nums)-1

// 		for left < right {
// 			n2, n3 := nums[left], nums[right]

// 			if n1+n2+n3 == 0 {
// 				res = append(res, []int{n1, n2, n3})
// 				for left < right && nums[left] == n2 {
// 					left++
// 				}
// 				////不用也可以過
// 				// for left < right && nums[right] == n3 {
// 				// 	right--
// 				// }
// 			} else if n1+n2+n3 < 0 {
// 				left++
// 			} else {
// 				right--
// 			}
// 		}

// 	}
// 	return res
// }

// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums)
// 	res := [][]int{}

// 	for i := 0; i < len(nums)-2; i++ {
// 		n1 := nums[i]
// 		if n1 > 0 {
// 			break
// 		}
// 		if i > 0 && n1 == nums[i-1] {
// 			continue
// 		}
// 		l, r := i+1, len(nums)-1
// 		for l < r {
// 			n2, n3 := nums[l], nums[r]
// 			if n1+n2+n3 == 0 {
// 				res = append(res, []int{n1, n2, n3})
// 				for l < r && nums[l] == n2 {
// 					l++
// 				}
// 				for l < r && nums[r] == n3 {
// 					r--
// 				}
// 			} else if n1+n2+n3 < 0 {
// 				l++
// 			} else {
// 				r--
// 			}
// 		}
// 	}
// 	return res
// }

//用dsf未完成
// var ans [][]int

// func threeSum(nums []int) [][]int {
// 	//hashMap := map[int]int{}
// 	ans = [][]int{}

// 	set := make([]int, 0)

// 	used := map[int]bool{}

// 	dfs(0, nums, set, 0, used)

// 	return ans
// }

// func dfs(start int, nums []int, set []int, sum int, used map[int]bool) {

// 	if len(set) == 3 {
// 		if sum == 0 {
// 			ans = append(ans, set)
// 		}
// 		return
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		if !used[i] {
// 			used[i] = true
// 			set = append(set, nums[i])
// 			sum += nums[i]
// 			dfs(i+1, nums, set, sum, used)
// 			sum -= nums[i]
// 			set = set[:len(set)-1]
// 		}
// 	}
// }

// 外层循环：指针 i 遍历数组。
// 内层循环：用双指针，去寻找满足三数之和 == 0 的元素
// 先排序的意义
// 便于跳过重复元素，如果当前元素和前一个元素相同，跳过。

// 双指针的移动时，避免出现重复解
// 找到一个解后，左右指针同时向内收缩，为了避免指向重复的元素，需要：

// 左指针在保证left < right的前提下，一直右移，直到指向不重复的元素
// 右指针在保证left < right的前提下，一直左移，直到指向不重复的元素
// 小优化
// 排序后，如果外层遍历的数已经大于0，则另外两个数一定大于0，sum不会等于0，直接break。

// func threeSum(nums []int) [][]int {
// 先排序的意义
// 便于跳过重复元素，如果当前元素和前一个元素相同，跳过。
// 	sort.Ints(nums)
// 	res := [][]int{}
//指针 i 遍历数组。
//第一個n1要找遍所有的值
//n2,n3用左右指針來尋找
// 	for i := 0; i < len(nums)-2; i++ {
// 		n1 := nums[i]
//		//排序后，如果外层遍历的数已经大于0，则另外两个数一定大于0，sum不会等于0，直接break。
// 		if n1 > 0 {
// 			break
// 		}
// 		if i > 0 && n1 == nums[i-1] {
// 			continue
// 		}
//建立左右指針
//做指針的i+1,右指針為最後一個數
// 		l, r := i+1, len(nums)-1
// 		for l < r {
// 			n2, n3 := nums[l], nums[r]
//如果三個數相加為0
// 			if n1+n2+n3 == 0 {
//  				//附加到res中
// 				res = append(res, []int{n1, n2, n3})
//				//當左指針小於右指針
//				// 左指針又等於n2,將指針往右加
// 				for l < r && nums[l] == n2 {
// 					l++
// 				}
////				不用也可以過
//				//當左指針小於右指針
//				// 右指針又等於n3,將指針往右加
// 				for l < r && nums[r] == n3 {
// 					r--
// 				}
//			// 如果 nums[i] + nums[left] + nums[right] < 0 说明 此时 三数之和小了，left 就向右移动，才能让三数之和大一些，直到left与right相遇为止。
// 			} else if n1+n2+n3 < 0 {
// 				l++
//			// 如果nums[i] + nums[left] + nums[right] > 0 就说明 此时三数之和大了，因为数组是排序后了，所以right下表就应该向左移动，这样才能让三数之和小一些。
// 			} else {
// 				r--
// 			}
// 		}
// 	}
// 	return res
// }

//自己寫的,會抓到重複的
// func threeSum(nums []int) [][]int {
// 	//hashMap := map[int]int{}
// 	ans := [][]int{}

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					ans = append(ans, []int{nums[i], nums[j], nums[k]})
// 				}
// 			}
// 		}
// 	}
// 	return ans
// }
