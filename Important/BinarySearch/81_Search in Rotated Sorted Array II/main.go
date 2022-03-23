package main

func main() {
	nums := []int{
		3, 1, 2, 3, 3, 3, 3,
	}
	target := 2
	search(nums, target)
	//fmt.Print(x)
}

// func search(nums []int, target int) bool {
// 	n := len(nums)
// 	left := 0
// 	right := n - 1
// 	// 恢复二段性
// 	for left < right && nums[0] == nums[right] {
// 		right--
// 	}

// 	// 第一次二分，找旋转点
// 	for left < right {
// 		mid := (left + right + 1) >> 1
// 		if nums[mid] >= nums[0] {
// 			left = mid
// 		} else {
// 			right = mid - 1
// 		}
// 	}

// 	idx := n
// 	if nums[right] >= nums[0] && right+1 < n {
// 		idx = right + 1
// 	}
// 	// 第二次二分，找目标值
// 	ans := find(nums, 0, idx-1, target)
// 	if ans != -1 {
// 		return true
// 	}
// 	ans = find(nums, idx, n-1, target)
// 	return ans != -1
// }
// func find(nums []int, l int, r int, t int) int {
// 	for l < r {
// 		mid := (l + r) >> 1
// 		if nums[mid] >= t {
// 			r = mid
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	if nums[r] == t {
// 		return t
// 	}
// 	return -1
// }

// 对于数组中有重复元素的情况，二分查找时可能会有 a[l]=a[{mid}]=a[r]，此时无法判断区间 [l,mid] 和区间 [mid+1,r] 哪个是有序的。

// 例如  nums=[3,1,2,3,3,3,3]， target=2，首次二分时无法判断区间 [0,3] 和区间 [4,6] 哪个是有序的。

// 对于这种情况，我们只能将当前二分区间的左边界加一，右边界减一，然后在新区间上继续二分查找。

//编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

func search(nums []int, target int) bool {
	n := len(nums)
	//如果長度0,直接跳出
	if n == 0 {
		return false
	}
	//如果長度只有1,直接傳index 0的值
	if n == 1 {
		return nums[0] == target
	}
	//設定左指針為0,右指針為長度-1(因為要取的是index)
	l, r := 0, n-1
	//指針未重和,持續尋找至重和
	for l <= r {
		//尋找中間值
		mid := (l + r) / 2
		//如果中間值直接就是target,直接返回
		if nums[mid] == target {
			return true
		}
		//如果有重复元素,無法分辨哪个是有序的
		//只能将当前二分区间的左边界加一，右边界减一，然后在新区间上继续二分查找。
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l++
			r--
		} else if nums[l] <= nums[mid] {
			//若nums[left] <= target <= nums[mid],则表示target在前半部分有序的范围内存在
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				//若target < nums[left] <= nums[mid], 则表示target小于前半部分升序序列中最小值，不能再前半部分序列中查找，需要在后半部分序列中查找
				//若nums[left] <= nums[mid] < target, 则表示target大于前半部分升序序列中最大值，也不能再前半部分序列中查找，需要在后半部分序列中查找
				l = mid + 1
			}
		} else { // nums[left] > nums[mid] 即：nums[mid] <= nums[right]

			//若nums[mid] <= target <= nums[right],则表示target在后半部分有序列的范围内存在
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				//若target < nums[mid] <= nums[right],则表示target小于后半部分升序序列中最小值，不能再后半部分序列中查找，需要在前半部分序列中查找
				//若nums[mid] <= nums[right] < target, 则表示target大于后半部分升序序列中最大值，也不能再后半部分序列中查找，需要在前半部分序列中查找
				r = mid - 1
			}
		}
	}
	return false
}
