package main

import "fmt"

func main() {
	var prices = []int{8, 8, 7, 7, 7}
	//target := 3
	fmt.Println(majorityElement(prices))

}
func majorityElement(nums []int) int {
	hashTable := map[int]int{}
	maxTable := map[int]int{}

	if len(nums) <= 2 {
		return nums[0]
	}
	for i, x := range nums {
		//fmt.Println(i)
		if _, ok := hashTable[x]; ok {
			maxTable[x] = maxTable[x] + 1
			// if x > max {
			// 	max = x
			// }
			//fmt.Println(maxTable)
		}
		hashTable[x] = i
		fmt.Println(maxTable)
	}
	max := 0
	maxIndex := 0
	for i, _ := range maxTable {
		//fmt.Println(i)
		//fmt.Println(maxTable[i])
		if maxTable[i] > max {
			max = maxTable[i]
			maxIndex = i
			fmt.Println(i)
		}
	}
	return maxIndex
}
