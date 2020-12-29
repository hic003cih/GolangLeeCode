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

//只能用一次迴圈
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	mergeNums := merge(nums1, nums2)
	fmt.Println(mergeNums)
	if len(mergeNums) < 2 {
		fmt.Println(mergeNums)
		return float64(mergeNums[0])
	}
	//如果是奇數
	if len(mergeNums)%2 == 1 {
		i := int(len(mergeNums) / 2)
		fmt.Println(i)
		return (float64(mergeNums[i]))
	}
	i := len(mergeNums) / 2
	// for {
	// 	i++
	// 	j--
	// 	if i == j || (i+1) == j {
	// 		break
	// 	}
	// }
	return (float64(mergeNums[i-1]) + float64(mergeNums[i])) / 2.0

}
func merge(nums1 []int, nums2 []int) []int {
	nums1 = append(nums1, nums2...)

	leftPoint := 0
	for rightPoint := 1; rightPoint < len(nums1); rightPoint++ {
		if nums1[leftPoint] > nums1[rightPoint] {

			tempRight := nums1[rightPoint]
			nums1[rightPoint] = nums1[leftPoint]
			nums1[leftPoint] = tempRight
		}
		leftPoint++
	}

	return nums1
}
