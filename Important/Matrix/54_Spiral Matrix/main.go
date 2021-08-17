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

// func spiralOrder(matrix [][]int) []int {
// 	if len(matrix) == 0 {
// 		return []int{}
// 	}
// 	res := []int{}

// 	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

// 	size := len(matrix) * len(matrix[0])

// 	for len(res) != size {
// 		for i := left; i <= right; i++ {
// 			res = append(res, matrix[top][i])
// 		}
// 		top++
// 		for i := top; i <= bottom; i++ {
// 			res = append(res, matrix[i][right])
// 		}
// 		right--
// 		if len(res) == size {
// 			break
// 		}
// 		for i := right; i >= left; i-- {
// 			res = append(res, matrix[bottom][i])
// 		}
// 		bottom--
// 		for i := bottom; i >= top; i-- {
// 			res = append(res, matrix[i][left])
// 		}
// 		left++
// 	}
// 	return res
// }

// func spiralOrder(matrix [][]int) []int {
// 	//如果matrix長度是0,返回0
// 	if len(matrix) == 0 {
// 		return []int{}
// 	}

// 	res := []int{}
// 	//設定上下左右的範圍
// 	//上設為0,下設為matrix的長度-1,左設為0,右設為matrix第一行的長度-1
// 	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
// 	//設定size為len(matrix) * len(matrix[0])
// 	size := len(matrix) * len(matrix[0])

// 	//循環執行把所有長度比對完
// 	for len(res) != size {
// 		//先將第一個橫列的數迴圈執行完
// 		//如果右邊大於左邊,將matrix[top][i]寫入
// 		for i := left; i <= right; i++ {

// 			res = append(res, matrix[top][i])
// 			fmt.Println("左往右")
// 			fmt.Println(res)
// 		}
// 		//第一個橫列的數迴圈執行完,top的範圍往下加1
// 		top++
// 		//比對最右邊直列的樹
// 		//如果最下面大於上面,將matrix[i][right]寫入
// 		for i := top; i <= bottom; i++ {

// 			res = append(res, matrix[i][right])
// 			fmt.Println("上往下")
// 			fmt.Println(res)
// 		}
// 		right--
// 		if len(res) == size {
// 			break
// 		}
// 		//比對最下面橫的數
// 		//如果最右邊大於左,將matrix[bottom][i]寫入
// 		for i := right; i >= left; i-- {

// 			res = append(res, matrix[bottom][i])
// 			fmt.Println("右往左")
// 			fmt.Println(res)
// 		}
// 		bottom--
// 		//比對最左直的數
// 		//如果最下面大於上面,將matrixmatrix[i][left]寫入
// 		for i := bottom; i >= top; i-- {

// 			res = append(res, matrix[i][left])
// 			fmt.Println("下往上")
// 			fmt.Println(res)
// 		}
// 		left++
// 	}

// 	return res
// }
