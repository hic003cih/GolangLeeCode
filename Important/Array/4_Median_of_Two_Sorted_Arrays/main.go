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

//正確版
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	mergedNums := append(nums1, nums2...)

// 	if len(mergedNums) < 2 {
// 		return float64(mergedNums[0])
// 	}
// 	if len(mergedNums)%2 == 0 {

// 	}

// 	fmt.Println(mergedNums)
// 	fmt.Println(findMax(mergedNums))

// 	return 0.0
// }

// func findMax(nums []int) int {
// 	sort.Ints(nums)

// 	fmt.Println(nums)

// 	return nums[len(nums)-1]
// }

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
