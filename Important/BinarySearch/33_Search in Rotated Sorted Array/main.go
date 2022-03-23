package main

func main() {
	nums := []int{
		4, 5, 6, 7, 0, 1, 2,
	}
	target := 0
	search(nums, target)
	//fmt.Print(x)
}

// func search(nums []int, target int) int {
// 	lens := len(nums)
// 	// 1. 找临界点确定区间
// 	index := 0
// 	for index <= (lens-2) && nums[index] < nums[index+1] { // 先判断范围，以免索引超出界限
// 		index++
// 	}

// 	// 2. 确定区间
// 	left, right := 0, lens-1 //初始化
// 	if target >= nums[0] {   // 在左边区间
// 		right = index
// 	} else { // 在右区间
// 		left = index + 1
// 	}

// 	// 二分
// 	for left <= right {
// 		mid := (left + right) / 2
// 		fmt.Println(mid)
// 		if nums[mid] == target {
// 			return mid
// 		} else if target < nums[mid] {
// 			right = mid - 1
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return -1
// }

// func search(nums []int, target int) int {
// 	//建立左右指標
// 	left, right := 0, len(nums)-1
// 	//如果右指針大於等於左指針
// 	for left <= right {
// 		//計算mid值
// 		//為什麼要+left
// 		mid := (right-left)/2 + left
// 		//如果取道target值,就返回
// 		if nums[mid] == target {
// 			return mid
// 		}
// 		//判斷mid位置
// 		//在左邊的區塊
// 		if nums[mid] >= nums[left] {
// 			//判断是否在 A 段
// 			if nums[mid] > target && target >= nums[left] {
// 				//右指針設為中間往左移一
// 				//因為mid已經確定不是了,所以不用用mid
// 				right = mid - 1
// 			} else {
// 				//左指針設為中間往右移一,所以不用用mid
// 				left = mid + 1
// 			}
// 		} else { //nums[mid] < nums[left],在右半部分
// 			//判断是否在 B 段
// 			if nums[mid] < target && target <= nums[right] {
// 				left = mid + 1
// 			} else {
// 				right = mid - 1
// 			}
// 		}
// 	}
// 	return -1
// }

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// nums[left] <= nums[mid]
// 此时前半部分是升序序列
// (1) 若nums[left] <= target <= nums[mid],则表示target在前半部分有序的范围内存在
// right = mid - 1
// (2) 若target < nums[left] <= nums[mid], 则表示target小于前半部分升序序列中最小值，不能再前半部分序列中查找，需要在后半部分序列中查找
// left = mid + 1
// (3) 若nums[left] <= nums[mid] < target, 则表示target大于前半部分升序序列中最大值，也不能再前半部分序列中查找，需要在后半部分序列中查找
// left = mid + 1

// nums[left] > nums[mid] 即：nums[mid] <= nums[right]
// 此时后半部分序列有序
// (1) 若nums[mid] <= target <= nums[right],则表示target在后半部分有序列的范围内存在
// left = mid + 1
// (2) 若target < nums[mid] <= nums[right],则表示target小于后半部分升序序列中最小值，不能再后半部分序列中查找，需要在前半部分序列中查找
// right = mid - 1
// (3) 若nums[mid] <= nums[right] < target, 则表示target大于后半部分升序序列中最大值，也不能再后半部分序列中查找，需要在前半部分序列中查找
// right = mid - 1
