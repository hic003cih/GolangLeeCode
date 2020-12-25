package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	res := make([]int, 0)
	for _, v := range nums {

		for y := 0; y < len(nums); y++ {
			if nums[y] != v {

				fmt.Printf("nums[:y] : %v\n", nums[:y])
				//fmt.Println(nums[y+1:])

				fmt.Printf("nums[y+1:] : %v\n", nums[y+1:])
				res = append(res, nums[y])
				fmt.Println(res)
			}

			// fmt.Printf("y : %v\n", y)
			// fmt.Printf("v : %v\n", v)
			// fmt.Printf("[y] : %v\n", nums[y])
			// if nums[y] == v {

			// 	fmt.Printf("nums[:y] : %v\n", nums[:y])
			// 	//fmt.Println(nums[y+1:])

			// 	fmt.Printf("nums[y+1:] : %v\n", nums[y+1:])
			// 	nums = append(nums[:y], nums[y+1:]...)
			// 	fmt.Println(nums)
			// }
		}
	}
	return len(nums)
}
