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
	//如果matrix長度是0,返回0
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	//設定上下左右的範圍
	//上設為0,下設為matrix的長度-1,左設為0,右設為matrix第一行的長度-1
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	//計算矩陣長度
	size := len(matrix) * len(matrix[0])
	//如果存入的結果長度還沒跟矩陣長度一樣,不斷執行
	//先處理最外圍的一圈,在處理裡面的一圈
	for len(res) != size {
		//上面的橫列.每處裡一筆,left往右移
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		//處理完以後,top加一往下移
		top++
		//上面的橫列處理完,指標到達最右邊,右邊的直行往下處理.每處裡一筆,top往下移
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		//處理完以後,right減一往左縮
		right--
		//檢查是否已經完成
		if len(res) == size {
			break
		}
		//右邊的直行處理完後,指標到達最下面,往左處理.每處裡一筆,right往左移
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		//處理完以後,bottom減一往上移
		bottom--
		//下面的橫列處理完後,指標到達最左邊,往上處理.每處裡一筆,bottom往上移
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		//處理完以後,left加一往右移
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
