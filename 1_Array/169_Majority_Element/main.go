package main

import "fmt"

func main() {
	var prices = []int{8, 8, 8, 7, 7}
	//target := 3
	fmt.Println(majorityElement(prices))

}
func majorityElement(nums []int) int {

	//最原始的map記錄法
	// hashTable := map[int]int{}
	// maxTable := map[int]int{}

	// if len(nums) <= 2 {
	// 	return nums[0]
	// }
	// for i, x := range nums {
	// 	//fmt.Println(i)
	// 	if _, ok := hashTable[x]; ok {
	// 		maxTable[x] = maxTable[x] + 1
	// 		// if x > max {
	// 		// 	max = x
	// 		// }
	// 		//fmt.Println(maxTable)
	// 	}
	// 	hashTable[x] = i
	// 	fmt.Println(maxTable)
	// }
	// max := 0
	// maxIndex := 0
	// for i, _ := range maxTable {
	// 	//fmt.Println(i)
	// 	//fmt.Println(maxTable[i])
	// 	if maxTable[i] > max {
	// 		max = maxTable[i]
	// 		maxIndex = i
	// 		fmt.Println(i)
	// 	}
	// }
	// return maxIndex

	//摩爾投票法
	//一定要出現的數超過一半才可以用
	//設定兩個變數,max為紀錄多數的數,count為記錄出現的次數,遇到不一樣的count減1,最後剩下的為贏家

	max := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if max != nums[i] {
			count -= 1
			if count == 0 {
				max = nums[i]
				count = 1
			}
		} else {
			count += 1
		}
	}
	return max

}
