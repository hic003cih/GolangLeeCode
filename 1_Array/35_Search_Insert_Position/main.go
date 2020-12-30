package main

import "fmt"

// import (
// 	"fmt"
// )

//const nums1=[]int{1,3,5,6}
//const tar=5

func main()  {
	//[1,3,5,6], 5
	nums1:=[]int{1,3,5,6}
	tar:=1
	fmt.Print(searchInsert(nums1,tar))
}

func searchInsert(nums []int, target int) int {
	//暴力法
	// hashTable :=map[int]int{}
	// for i,v := range nums{
	// 	hashTable[v] =i
	// }
	// if nums[0]>target{
	// 	return 0
	// }
	// if nums[len(nums)-1] < target{
	// 	return len(nums)
	// }
	// if p,ok:=hashTable[target];ok{
	// 	return p
	// }
	// for i := 0; i < len(nums); i++ {
	// 	if(i==(len(nums)-1)){
	// 		return i+1
	// 	}
		
	// 	if(nums[i]<=target&&nums[i+1]>target){
	// 		return i+1
	// 	}
	// }

	// return 0

	//二分法
	left:=0
	right:=len(nums)

	for left <right{
		mid :=int((left+right) /2)
		if(target<=nums[mid]){
			right = mid
		}else{
			left = mid+1
		}
	}
	return left
}