package main

import "fmt"

func main() {
	nums := []int{3, 2, 3, 2, 3}
	val := 3
	//removeElement(nums, val)
	fmt.Println(removeElement(nums, val))
}

func removeElement(nums []int, val int) int {

	leftPoint := 0
	//左指針從0開始,右指針也從0開始
	for rightPoint, v := range nums {
		//如果右指針的值不等於要消去的值
		//把右邊的值往左移
		//左指針加一
		if v != val {
			//fmt.Println(nums)
			//如果左指針的值不等於要消去的值
			//
			nums[leftPoint] = nums[rightPoint]
			leftPoint++
		}
	}
	return leftPoint

	//正確版
	/* leftPoint := 0

	for rightPoint, v := range nums {

		if v != val {
			//fmt.Println(nums)
			nums[leftPoint] = nums[rightPoint]
			leftPoint++
		}
	}
	return leftPoint */

	//removeNums := map[int]int{}
	//removeNums := make([]int, 0)
	//length := 0
	// for i, v := range nums {

	// 	if v != val {
	// 		//fmt.Println(v)
	// 		//length++
	// 		//removeNums = append(removeNums, v)
	// 		nums = append(nums[:i], nums[i+1:]...)
	// 	}
	// }
	//nums = removeNums

	//return len(nums)

	// for i := 0; i < len(nums); {
	// 	if nums[i] == val {
	// 		//取到i位置之前的數組,並把i+1後的全部數組
	// 		nums = append(nums[:i], nums[i+1:]...)
	// 	} else {
	// 		i++
	// 	}
	// }
	// return len(nums)

	// 0->{} ->{2,2,3}
	// 1->{2,2,3}
	// 2->{2,2,3}
	// 3->{2,2} -> {2,2}

	// length := len(nums)
	// if length == 0 {
	// 	return 0
	// }

	// i := 0
	// j := 0
	// for j < length {
	// 	if nums[j] == val {
	// 		// 去找一个不是 val 的值
	// 		j++
	// 	} else {
	// 		// 赋值
	// 		nums[i] = nums[j]
	// 		i++
	// 		j++
	// 	}
	// }

	// return length - (j - i)

}
