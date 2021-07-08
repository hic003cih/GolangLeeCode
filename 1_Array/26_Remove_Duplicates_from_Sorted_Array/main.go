package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {

	//長度小於1,直接回傳nums長度
	if len(nums) < 1 {
		return len(nums)
	} else {
		//左指針
		leftPoint := 0
		//右指針循環
		//從1開始
		for rightPoint := 1; rightPoint < len(nums); rightPoint++ {
			//如果左指針的值不等於右指針
			//不一樣的值的長度+1
			if nums[leftPoint] != nums[rightPoint] {
				//保留左邊一樣的數字,指針往右移1,把新的數字加在右邊
				leftPoint++
				//指針往右移1,把新的數字加在右邊
				nums[leftPoint] = nums[rightPoint]
			}
		}
		//最後返回指針往右移動的次數,最後要加上1(原本還有一次)
		return leftPoint + 1
	}

	//最正確版本
	/* leftPoint := 0
	for rightPoint := 1; rightPoint < len(nums); rightPoint++ {
		if nums[leftPoint] != nums[rightPoint] {
			leftPoint++
			nums[leftPoint] = nums[rightPoint]
		}
	}
	return leftPoint + 1
	*/
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
