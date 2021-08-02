package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 4}
	target := 6
	x := twoSum(nums, target)
	fmt.Print(x)
}

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		//if x, ok := hasTable[target-v]; ok
		if x, ok := hashMap[target-nums[i]]; ok {
			return []int{x, i}
		}
		// _, ok := hashMap[target-nums[i]]
		// if ok == true {
		// 	return []int{hashMap[target-nums[i]], i}
		// }
		hashMap[nums[i]] = i

	}
	return []int{}
}

// func main() {
// 	nums := []int{3, 2, 4}
// 	target := 6
// 	x := twoSum(nums, target)
// 	fmt.Print(x)
// }

// func twoSum(nums []int, target int) []int {
// 	//暴力法
// 	if len(nums) > 1 {
// 	for i := 0; i < len(nums); i++ {
// 		for y := i + 1; y < len(nums); y++ {
// 			if (nums[i] + nums[y]) == target {
// 				return []int{i, y}
// 			}
// 		}
// 	}
// }
// return []int{}
// 	//哈希表

// 	hasTable := map[int]int{}

// 	for i, v := range nums {
// 		//x是map的值
// 		if x, ok := hasTable[target-v]; ok {
// 			return []int{x, i}
// 		}
// 		//map index存value,值存nums是第幾位置
// 		hasTable[v] = i
// 	}
// 	return nil
// }

// func searchInsert(nums []int, target int) int {
// 	left := 0

// 	right := len(nums)

// 	for left < right {

// 		mid := int((left + right) / 2)
// 		if target > nums[mid] {
// 			left = mid + 1
// 		} else {
// 			right = mid
// 		}

// 	}
// 	return left
// }
