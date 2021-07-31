package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 1, 0, 2}
	//subsets(nums)
	//fmt.Println(mask>>i&1 > 0)
	fmt.Print(trap(nums))
	//fmt.Println(2>>0&1 > 0)
	//fmt.Print(x)
}

// func trap(height []int) int {

// }

func trap(height []int) (ans int) {
	//找到array的長度
	n := len(height)
	//如果是0就返回
	if n == 0 {
		return
	}
	//建立左邊最高的array
	leftMax := make([]int, n)
	//預設左邊最高的為index 0
	leftMax[0] = height[0]
	//迴圈逐一比較array中每個數
	//當前的i要存的是,leftMax的index-1和height的當前數,做比較後較高的
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	//建立右邊最高的array
	rightMax := make([]int, n)
	//預設右邊最高的為index最後
	rightMax[n-1] = height[n-1]

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
	}
	return
}

//尋找最小的
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//尋找最大的
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// int trap(vector<int>& height)
// {
//     int ans = 0;
//     int size = height.size();
//     for (int i = 1; i < size - 1; i++) {
//         int max_left = 0, max_right = 0;
//         for (int j = i; j >= 0; j--) { //Search the left part for max bar size
//             max_left = max(max_left, height[j]);
//         }
//         for (int j = i; j < size; j++) { //Search the right part for max bar size
//             max_right = max(max_right, height[j]);
//         }
//         ans += min(max_left, max_right) - height[i];
//     }
//     return ans;
// }
