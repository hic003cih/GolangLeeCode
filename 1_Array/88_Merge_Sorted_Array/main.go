package main

import "fmt"

const (
	//var nums1 = []int{1, 2, 3, 0, 0, 0}
	m = 3
	//var nums2 := []int{2, 5, 6}
	n = 2
)

func main() {
	var nums1 = []int{1, 2, 3, 0, 5, 1}
	var nums2 = []int{2, 5, 6}
	merge(nums1, m, nums2, n)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	// //时间复杂度 : O((n + m)\log(n + m))
	// //空间复杂度 : O(1)O(1)。
	// //兩組合併以後
	// sumNums := append(nums1[:m], nums2[:n]...)
	// //利用sort包內的Ints自動排序
	// sort.Ints(sumNums)

	index1 := m - 1
	index2 := n - 1
	tail := m + n - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[tail] = nums1[index1]
			index1--
		} else {
			nums1[tail] = nums2[index2]
			index2--
		}
		tail--
		//fmt.Println(tail)
		fmt.Println(index1)
		fmt.Println(nums1)
	}
	for tail >= 0 && index1 >= 0 {
		nums1[tail] = nums1[index1]
		index1--
		tail--
	}
	for tail >= 0 && index2 >= 0 {
		nums1[tail] = nums2[index2]
		index2--
		tail--
	}

	//return leftPoint + 1

}
