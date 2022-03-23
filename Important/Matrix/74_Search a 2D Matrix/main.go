package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}
	target := 13

	x := searchMatrix(matrix, target)
	fmt.Print(x)
}

//因為Matrix是依大小作排序,所以
// 规定左下为起始点
// 若点(i,j)等于target，直接返回true
// 若点(i,j)小于target，向右走，及j++
// 若点(i,j)大于target，向上走，及i--
// 若已经到达了边界，跳出循环，说明未找到

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	//衡的
	// row := len(matrix)
	// //直的
	// col := len(matrix[0])
	// //1. 规定起始点
	// i, j := row-1, 0

	//将「二维矩阵」当做「一维矩阵」来做
	//m存matrix的row的長度
	//low存0
	//high存matrix的總長度,因為是用來做index,所以要-1
	m, low, high := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	fmt.Println(high)
	//執行完全部matrix的總長度
	for low <= high {
		//尋找中位數
		//全部matrix的總長度減掉已經找過的數量
		mid := (high-low)/2 + low
		//x := matrix[mid/n][mid%n]
		// if matrix[i][j] == target { //2.找到则返回
		// 	return true
		// 	//
		// } else if matrix[i][j] < target { //3.小于target，向右查找
		// 	j++
		// } else { //4.大于target，向上查找
		// 	i--
		// }
		//2.找到则返回
		if matrix[mid/m][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] > target { //3.大于target，向左查找
			high = mid - 1
		} else { //4.小于target，向下查找
			low = mid + 1
		}
	}

	return false
}

// func searchMatrix(matrix [][]int, target int) bool {
// 	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
// 	if row < 0 {
// 		return false
// 	}
// 	col := sort.SearchInts(matrix[row], target)
// 	return col < len(matrix[row]) && matrix[row][col] == target
// }

//最左下往上和往左尋找法
// func searchMatrix(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 {
// 		return false
// 	}
// 	row := len(matrix)
// 	col := len(matrix[0])
// 	//1. 规定起始点
// 	i, j := row-1, 0
// 	for i >= 0 && j < col {
// 		if matrix[i][j] == target { //2.找到则返回
// 			return true
// 		} else if matrix[i][j] < target { //3.小于target，向右查找
// 			j++
// 		} else { //4.大于target，向上查找
// 			i--
// 		}
// 	}

// 	return false
// }

// func searchMatrix(matrix [][]int, target int) bool {
// 	//m為column長度
// 	//n為row長度
// 	m, n := len(matrix), len(matrix[0])
// 	//
// 	i := sort.Search(m*n, func(i int) bool { return matrix[i/n][i%n] >= target })
// 	return i < m*n && matrix[i/n][i%n] == target
// }
