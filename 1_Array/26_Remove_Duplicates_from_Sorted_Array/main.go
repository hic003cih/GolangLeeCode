package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {

	leftPoint := 0
	for rightPoint := 1; rightPoint < len(nums); rightPoint++ {
		if nums[leftPoint] != nums[rightPoint] {
			leftPoint++
			nums[leftPoint] = nums[rightPoint]
		}
	}
	return leftPoint + 1

	//res := make([]int, 0)
	//for i, v := range nums {
	//	if nums[i] == nums[i+1]{
	//		continue
	//	}
	//
	//    y:=0
	//	for y = 0; y < len(nums); y++ {
	//		if nums[y] == v {
	//
	//			//fmt.Printf("nums[:y] : %v\n", nums[:y])
	//			//fmt.Println(nums[y+1:])
	//
	//			//fmt.Printf("nums[y+1:] : %v\n", nums[y+1:])
	//			res = append(res, nums[y])
	//			fmt.Println(res)
	//		}
	//	}
	//}
	//return len(nums)
}
