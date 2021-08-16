package main

import (
	"fmt"
)

// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]

func main() {
	matrix := [][]int{
		{9, 9, 4}, {6, 6, 8}, {2, 1, 1},
	}

	x := longestIncreasingPath(matrix)
	fmt.Print(x)
}
func longestIncreasingPath(matrix [][]int) int {
	//思路：采用动态规划的算法，dp[i][j],代表以自身开头的最长的递增路径的长度
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 1
	//初始化
	//取matrix的長度
	dp := make([][]int, len(matrix))
	//建立matrix內的各個array
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		fmt.Println(dp[i])
	}

	//获取以某个点开始的最长的递增路径大小
	var getMaxDisPath func(i, j int) int

	getMaxDisPath = func(i, j int) int {
		//用途不清楚
		if dp[i][j] != 0 {
			return dp[i][j]
		}
		max := 1
		//可以向下
		//由上往下搜尋
		//如果i+1 < len(matrix),代表還沒找到最後一個,就繼續執行
		if i+1 < len(matrix) {
			//先從第一直行一個數開始往下找
			fmt.Println("matrix[i][j]", matrix[i][j])
			if matrix[i][j] < matrix[i+1][j] {
				//
				down := getMaxDisPath(i+1, j)

				fmt.Println(down)
				//如果
				if down+1 > max {
					max = down + 1
				}
			}
		}
		//可以向右
		if j+1 < len(matrix[0]) {
			if matrix[i][j] < matrix[i][j+1] {
				right := getMaxDisPath(i, j+1)
				if right+1 > max {
					max = right + 1
				}
			}
		}
		//可以向左
		if j-1 >= 0 {
			if matrix[i][j] < matrix[i][j-1] {
				left := getMaxDisPath(i, j-1)
				if left+1 > max {
					max = left + 1
				}
			}
		}
		//可以向上
		if i-1 >= 0 {
			if matrix[i][j] < matrix[i-1][j] {
				up := getMaxDisPath(i-1, j)
				if up+1 > max {
					max = up + 1
				}
			}
		}
		dp[i][j] = max

		return dp[i][j]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if res < getMaxDisPath(i, j) {
				res = getMaxDisPath(i, j)
			}
		}
	}
	return res
}

//最原始版
// func longestIncreasingPath(matrix [][]int) int {
// 	//思路：采用动态规划的算法，dp[i][j],代表以自身开头的最长的递增路径的长度
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return 0
// 	}
// 	res := 1
// 	//初始化
// 	dp := make([][]int, len(matrix))
// 	for i := 0; i < len(matrix); i++ {
// 		dp[i] = make([]int, len(matrix[0]))
// 	}

// 	//获取以某个点开始的最长的递增路径大小
// 	var getMaxDisPath func(i, j int) int
// 	getMaxDisPath = func(i, j int) int {
// 		if dp[i][j] != 0 {
// 			return dp[i][j]
// 		}
// 		max := 1
// 		//可以向下
// 		if i+1 < len(matrix) {
// 			if matrix[i][j] < matrix[i+1][j] {
// 				down := getMaxDisPath(i+1, j)
// 				if down+1 > max {
// 					max = down + 1
// 				}
// 			}
// 		}
// 		//可以向右
// 		if j+1 < len(matrix[0]) {
// 			if matrix[i][j] < matrix[i][j+1] {
// 				right := getMaxDisPath(i, j+1)
// 				if right+1 > max {
// 					max = right + 1
// 				}
// 			}
// 		}
// 		//可以向左
// 		if j-1 >= 0 {
// 			if matrix[i][j] < matrix[i][j-1] {
// 				left := getMaxDisPath(i, j-1)
// 				if left+1 > max {
// 					max = left + 1
// 				}
// 			}
// 		}
// 		//可以向上
// 		if i-1 >= 0 {
// 			if matrix[i][j] < matrix[i-1][j] {
// 				up := getMaxDisPath(i-1, j)
// 				if up+1 > max {
// 					max = up + 1
// 				}
// 			}
// 		}
// 		dp[i][j] = max

// 		return dp[i][j]
// 	}

// 	for i := 0; i < len(matrix); i++ {
// 		for j := 0; j < len(matrix[0]); j++ {
// 			if res < getMaxDisPath(i, j) {
// 				res = getMaxDisPath(i, j)
// 			}
// 		}
// 	}
// 	return res
// }
