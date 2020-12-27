package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	x := twoSum(nums, target)
	fmt.Print(x)
}

func twoSum(nums []int, target int) []int {
	//暴力法
	//for i := 0; i < len(nums); i++ {
	//	y := i
	//	for y = (y + 1); y < len(nums); y++ {
	//		if target == nums[i]+nums[y] {
	//			return []int{i, y}
	//		}
	//	}
	//}
	//return nil
	//哈希表
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			fmt.Printf(string(p))
			return []int{p, i}
		}
		hashTable[x] = i
		fmt.Printf(string(hashTable[x]))
	}
	return nil
}
