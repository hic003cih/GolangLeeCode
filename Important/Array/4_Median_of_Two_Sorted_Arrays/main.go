package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{-1, 5}
	//findMedianSortedArrays(nums1, nums2)
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var midM = len(nums1) / 2
	var midN = len(nums2) / 2

	if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			return (float64(nums2[midN]) + float64(nums2[midN-1])) / 2
		} else {
			return (float64(nums2[midN]))
		}
	}
	if len(nums2) == 0 {
		if len(nums1)%2 == 0 {
			return (float64(nums1[midM]) + float64(nums1[midM-1])) / 2
		} else {
			return (float64(nums1[midM]))
		}
	}

	mergedNums := append(nums1, nums2...)

	sort.Ints(mergedNums)

	if len(mergedNums) < 2 {
		return float64(mergedNums[0])
	}

	mid := len(mergedNums) / 2

	if len(mergedNums)%2 == 0 {
		return (float64(mergedNums[mid]) + float64(mergedNums[mid-1])) / 2
	} else {
		return (float64(mergedNums[mid]))
	}
}

//自己寫的版本,簡化版
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

// 	//先做長度小於1的array的判斷
// 	var midM = len(nums1) / 2
// 	var midN = len(nums2) / 2

// 	if len(nums1) == 0 {
// 		if len(nums2)%2 == 0 {
// 			return (float64(nums2[midN]) + float64(nums2[midN-1])) / 2
// 		} else {
// 			return (float64(nums2[midN]))
// 		}
// 	}
// 	if len(nums2) == 0 {
// 		if len(nums1)%2 == 0 {
// 			return (float64(nums1[midM]) + float64(nums1[midM-1])) / 2
// 		} else {
// 			return (float64(nums1[midM]))
// 		}
// 	}

// 	mergedNums := append(nums1, nums2...)

// 	sort.Ints(mergedNums)

// 	if len(mergedNums) < 2 {
// 		return float64(mergedNums[0])
// 	}

// 	mid := len(mergedNums) / 2

// 	if len(mergedNums)%2 == 0 {
// 		return (float64(mergedNums[mid]) + float64(mergedNums[mid-1])) / 2
// 	} else {
// 		return (float64(mergedNums[mid]))
// 	}
// }

//自己寫的版本,效能較差
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	mergedNums := append(nums1, nums2...)

// 	sort.Ints(mergedNums)

// 	fmt.Println(mergedNums)

// 	ans := 0.0

// 	mid := 0

// 	if len(mergedNums) < 2 {
// 		return float64(mergedNums[0])
// 	}
// 	//總長度除2如果餘0,代表示偶數
// 	if len(mergedNums)%2 == 0 {
// 		//總長度除2沒有餘1,就代表是偶數
// 		//中位數會有兩個,除出來的數和除出來的數減1
// 		//兩個相加除2就會變成中位數平均
// 		mid = len(mergedNums) / 2
// 		ans = (float64(mergedNums[mid]) + float64(mergedNums[mid-1])) / 2
// 		return ans
// 	} else {
// 		//總長度除2如果餘1,代表示奇數
// 		//中位數就是totalLength除以2
// 		//中位數是除出來的數
// 		mid = len(mergedNums) / 2
// 		return float64(mergedNums[mid])
// 	}

// 	//fmt.Println(mergedNums)
// 	//fmt.Println(findMax(mergedNums))

// 	return 0.0
// }

// //官方正確版,尋找K數,校能較好,但複雜
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	//先算出總長度
// 	totalLength := len(nums1) + len(nums2)
// 	//總長度除2如果餘1,代表示奇數
// 	if totalLength%2 == 1 {
// 		//中位數就是totalLength除以2
// 		//中位數是除出來的數
// 		midIndex := totalLength / 2
// 		fmt.Println(midIndex)
// 		return float64(getKthElement(nums1, nums2, midIndex+1))
// 	} else {
// 		//總長度除2沒有餘1,就代表是偶數
// 		//中位數會有兩個,除出來的數和除出來的數減1
// 		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
// 		fmt.Println(midIndex1)
// 		fmt.Println(midIndex2)
// 		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
// 	}
// 	return 0
// }

// func getKthElement(nums1, nums2 []int, k int) int {
// 	//建立兩個指針
// 	index1, index2 := 0, 0
// 	for {
// 		//如果左指針等於nums1的長度
// 		if index1 == len(nums1) {
// 			return nums2[index2+k-1]
// 		}
// 		//如果右指針等於nums2的長度
// 		if index2 == len(nums2) {
// 			return nums1[index1+k-1]
// 		}
// 		//如果k==1,直接尋找最小值
// 		if k == 1 {
// 			return min(nums1[index1], nums2[index2])
// 		}
// 		//尋找中位數
// 		half := k / 2
// 		//如果新的index
// 		newIndex1 := min(index1+half, len(nums1)) - 1

// 		newIndex2 := min(index2+half, len(nums2)) - 1

// 		//左指針右指針
// 		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
// 		//如果左指針小於右指針
// 		if pivot1 <= pivot2 {
// 			k -= (newIndex1 - index1 + 1)
// 			index1 = newIndex1 + 1
// 		} else {
// 			k -= (newIndex2 - index2 + 1)
// 			index2 = newIndex2 + 1
// 		}
// 	}
// 	return 0
// }

// //尋找最小值
// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }

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
