package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-1, 3}
	//findMedianSortedArrays(nums1, nums2)
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sumLen := len(nums1) + len(nums2)
	if len(sumLen) < 2 {
		if len(nums1) != 0 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}
	if len(sumLen)%2 == 1 {
		numfindMax()
	}
}

func findMax(nums []int) int {
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i+1] > max {
			max = nums[i+1]
		}

	}
	return max
}

//只能用一次迴圈
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	mergeNums := merge(nums1, nums2)
// 	fmt.Println(mergeNums)
// 	if len(mergeNums) < 2 {
// 		fmt.Println(mergeNums)
// 		return float64(mergeNums[0])
// 	}
// 	//如果是奇數
// 	if len(mergeNums)%2 == 1 {
// 		i := int(len(mergeNums) / 2)
// 		fmt.Println(i)
// 		return (float64(mergeNums[i]))
// 	}
// 	i := len(mergeNums) / 2
// 	return (float64(mergeNums[i-1]) + float64(mergeNums[i])) / 2.0

// }
// func merge(nums1 []int, nums2 []int) []int {
// 	nums1 = append(nums1, nums2...)

// 	leftPoint := 0

// 	for rightPoint := 1; rightPoint < len(nums); rightPoint++ {
// 		if nums[rightPoint] > nums[leftPoint] {
// 			tempRight := nums1[rightPoint]
// 			nums1[rightPoint] = nums1[leftPoint]
// 			nums1[leftPoint] = tempRight
// 			leftPoint++
// 			nums[leftPoint] = nums[rightPoint]
// 		}
// 	}
// 	return leftPoint + 1

// 	for rightPoint := 1; rightPoint < len(nums1); rightPoint++ {
// 		if nums1[leftPoint] > nums1[rightPoint] {

// 			tempRight := nums1[rightPoint]
// 			nums1[rightPoint] = nums1[leftPoint]
// 			nums1[leftPoint] = tempRight
// 		}
// 		leftPoint++
// 	}
// 	// for rightPoint := 1; rightPoint < len(nums1); rightPoint++ {
// 	// 	if nums1[leftPoint] > nums1[rightPoint] {

// 	// 		tempRight := nums1[rightPoint]
// 	// 		nums1[rightPoint] = nums1[leftPoint]
// 	// 		nums1[leftPoint] = tempRight
// 	// 	}
// 	// 	leftPoint++
// 	// }

// 	return nums1
// }
