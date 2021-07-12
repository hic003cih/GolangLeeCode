package main

import "fmt"

func main()  {
	nums := []int{-1,0,3,5,9,12,13}
	target := 13


	fmt.Println(search(nums,target))
}
func search(nums []int, target int) int {
	//計算總長度
	length := len(nums)
	//如果長度是0,直接return -1
	if length==0 {
		return -1
	}
	//計算中間的index
	mid :=length /2
	//如果是直接等於中間值,就不用迴圈計算了,直接回傳
	//不用特別提出來,放在下面比較也執行時間也是差不多
	/*if target==nums[mid]{
		return mid
	}*/
	//如果target大或等於,就比對右半邊的
	if target>=nums[mid] {
		for i := mid; i < length; i++ {
			if nums[i]==target {
				return i
			}
		}
	}else {
		//如果target小,就比對左半邊的
		for i := 0; i < mid; i++ {
			if nums[i]==target {
				return i
			}
		}
	}
	//比對沒有找到回傳－１
 return -1
}


//正確版

/*func search(nums []int, target int) int {
	high := len(nums)-1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid-1
		} else {
			low = mid+1
		}
	}
	return -1
}*/