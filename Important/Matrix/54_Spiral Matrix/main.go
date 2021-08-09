package main

import (
	"fmt"
)

// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]

func main() {
	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}

	x := spiralOrder(matrix)
	fmt.Print(x)
}
func spiralOrder(matrix [][]int) []int {
	// ans := make([]int, 0)

	// fmt.Println(len(matrix))
	// for i := 0; i < len(matrix[0]); i++ {
	// 	for j := 0; j < len(matrix); j++ {
	// 		if matrix[i][j] != 0 {
	// 			ans = append(ans, Res{i, j, matrix[i][j]})
	// 		}
	// 	}
	// }
	// //ans = append(ans, matrix[1][1])
	// return ans
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	size := len(matrix) * len(matrix[0])

	for len(res) != size {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if len(res) == size {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
