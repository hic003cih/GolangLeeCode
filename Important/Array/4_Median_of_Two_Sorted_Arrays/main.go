package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 2, 3, 4}
	nums2 := []int{6, 7, 8}
	//findMedianSortedArrays(nums1, nums2)
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

//递归消除法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var m = len(nums1)
	var n = len(nums2)
	if n < m { // 确保nums1比nums2短，即确保m比n小
		return findMedianSortedArrays(nums2, nums1)
	}
	//midM為num1的中位數
	var midM = (m - 1) / 2
	//midN為num2的中位數
	var midN = (n - 1) / 2

	//处理长度为0的情况
	//如果短的num已經沒有長度可以刪除了,直接處理長的num找中位數
	if m == 0 {
		//長的num無法被偶數整除的話,代表還剩奇數個數量,直接返回中位數
		if n%2 == 1 {
			return float64(nums2[midN])
		}
		//長的num可以被偶數整除的話,代表還剩偶數個數量,要相加以後除2才能得出中位數
		return float64(nums2[midN]+nums2[midN+1]) / 2
	}

	//最终到达边界的情况下，nums1的长度可能是1或者2。对于这两种情况，我们只需拿nums1剩下的数和nums2中间几位数进行排序，就能得到中位数。

	if m == 1 || m == 2 { // 边界条件
		if n < 3 { // n小于3的情况下，取nums2所有元素和nums1的元素进行排序
			for i := 0; i < n; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else if n%2 == 1 { // n大于2且为奇数的情况下，取nums2中间3位和nums1的元素进行排序
			for i := midN - 1; i < midN+2; i++ {
				nums1 = append(nums1, nums2[i])
			}
		} else { // 其他情况下，取nums2的中间4位和nums1的元素进行排序
			for i := midN - 1; i < midN+3; i++ {
				nums1 = append(nums1, nums2[i])
			}
		}
		sort.Ints(nums1)
		m = len(nums1)
		midM = (m - 1) / 2

		if len(nums1)%2 == 1 {
			return float64(nums1[midM])
		} else {
			return float64(nums1[midM]+nums1[midM+1]) / 2
		}
	}

	//抓取要刪除的數量,兩個nums都要刪到一樣的長度,每次刪的數量是短的nums的一半
	// n为奇数时，midNP==midN。
	//奇數的時候不用+1,算出來的中位數直接就可以用
	var midNP = midN
	//n为偶数时，midNP==midN+1。
	//偶数的時候要+1才能找到中位數
	if n%2 == 0 {
		midNP++
	}

	//兩個nums的中位數相同時,两值相等的情况下，其值就是中位数
	if nums1[midM] == nums2[midNP] {
		return float64(nums1[midM])
	}
	//如果較長的nums的中位數大於短的nums的中位數
	if nums1[midM] < nums2[midNP] {
		//消除nums1数组0至midM-1下标的元素，和nums2数组n-midM下标之后的元素
		return findMedianSortedArrays(nums1[midM:], nums2[:n-midM])
	}
	//如果較短的nums的中位數大於長的nums的中位數
	//消除nums2数组0至midM-1下标的元素，和nums1数组n-midM下标之后的元素
	return findMedianSortedArrays(nums2[midM:], nums1[:m-midM])
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	var midNum1 = len(nums1) / 2
// 	var midNum2 = len(nums2) / 2

// 	if len(nums1) == 0 {
// 		return findMid(nums2, midNum2)
// 	}
// 	if len(nums2) == 0 {
// 		return findMid(nums1, midNum1)
// 	}

// 	mergedNums := append(nums1, nums2...)

// 	sort.Ints(mergedNums)

// 	var midMergedNums = len(mergedNums) / 2
// 	return findMid(mergedNums, midMergedNums)
// }

// func findMid(nums []int, mid int) float64 {
// 	if len(nums)%2 == 0 {
// 		return (float64(nums[mid]) + float64(nums[mid-1])) / 2
// 	} else {
// 		return float64(nums[mid])
// 	}
// }

//將findMid拆出重複使用
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	var midNum1 = len(nums1) / 2
// 	var midNum2 = len(nums2) / 2

// 	if len(nums1) == 0 {
// 		return findMid(nums2, midNum2)
// 	}
// 	if len(nums2) == 0 {
// 		return findMid(nums1, midNum1)
// 	}

// 	mergedNums := append(nums1, nums2...)

// 	sort.Ints(mergedNums)

// 	var midMergedNums = len(mergedNums) / 2
// 	return findMid(mergedNums, midMergedNums)
// }

// func findMid(nums []int, mid int) float64 {
// 	if len(nums)%2 == 0 {
// 		return (float64(nums[mid]) + float64(nums[mid-1])) / 2
// 	} else {
// 		return float64(nums[mid])
// 	}
// }

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
