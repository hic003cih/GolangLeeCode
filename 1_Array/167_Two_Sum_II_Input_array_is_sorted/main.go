package main

import "fmt"

func main() {
	var prices = []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(prices, target))

}

func twoSum(numbers []int, target int) []int {
	// hashTable := map[int]int{}
	// for i, x := range numbers {
	// 	if p, ok := hashTable[target-x]; ok {
	// 		return []int{p, i}
	// 	}
	// }
	// hashTable[x] = i

	hashTable := map[int]int{}
	for i, x := range numbers {
		if p, ok := hashTable[target-x]; ok {
			fmt.Printf(string(p))
			return []int{p + 1, i + 1}
		}
		hashTable[x] = i
		fmt.Printf(string(hashTable[x]))
	}
	return nil
}
