package main

import (
	"fmt"
)

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.

// Each column must contain the digits 1-9 without repetition.

// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

//Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '3', '.', '.', '7', '.', '.', '.', '.'},
	}
	//[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]

	x := isValidSudoku(board)
	fmt.Print(x)
}

// // 解法一 暴力遍历，时间复杂度 O(n^3)
// func isValidSudoku(board [][]byte) bool {
// 	// 判断橫的row
// 	for i := 0; i < 9; i++ {
// 		//建立暫存,用來驗證是否已經有存在過
// 		tmp := [10]int{}
// 		for j := 0; j < 9; j++ {
// 			//取橫的row的值,第一個array固定,後面的array不斷+1
// 			cellVal := board[i][j : j+1]
// 			//fmt.Println(string(cellVal))
// 			if string(cellVal) != "." {
// 				//把文字轉換成數字
// 				index, _ := strconv.Atoi(string(cellVal))
// 				//如果數字不在1~9之間,直接回傳false
// 				if index > 9 || index < 1 {
// 					return false
// 				}
// 				//如果數字已經存在暫存中,回傳false
// 				if tmp[index] == 1 {
// 					return false
// 				}
// 				//如果沒有存在過,存入暫存中,並給一個假設的值1
// 				tmp[index] = 1
// 			}
// 		}
// 	}
// 	// 判断列 column(直的)
// 	//先固定第二個array的值
// 	for i := 0; i < 9; i++ {
// 		tmp := [10]int{}
// 		//循環第一個array的值,就能取到每個column的第一個值
// 		for j := 0; j < 9; j++ {
// 			cellVal := board[j][i]
// 			if string(cellVal) != "." {
// 				index, _ := strconv.Atoi(string(cellVal))
// 				if index > 9 || index < 1 {
// 					return false
// 				}
// 				if tmp[index] == 1 {
// 					return false
// 				}
// 				tmp[index] = 1
// 			}
// 		}
// 	}
// 	// 判断 9宫格 3X3 cell
// 	//先定row要做幾次9宮格判斷,row要做3次
// 	for i := 0; i < 3; i++ {
// 		//再定column要做幾次9宮格判斷,column也要做3次
// 		for j := 0; j < 3; j++ {
// 			tmp := [10]int{}
// 			//九宮格循環,先做一個九宮格的row,i=0就是第一個,,i=1就是第二個
// 			//然後乘上3就能取到row
// 			//i*3+3,只比到九宮格第一個後面兩位
// 			for ii := i * 3; ii < i*3+3; ii++ {
// 				//九宮格循環,九宮格的column,j=0就是第一個,j=1就是第二個
// 				// j*3+3,只比到九宮格第一個後面兩位

// 				for jj := j * 3; jj < j*3+3; jj++ {
// 					cellVal := board[ii][jj]
// 					if string(cellVal) != "." {
// 						index, _ := strconv.Atoi(string(cellVal))
// 						if tmp[index] == 1 {
// 							return false
// 						}
// 						tmp[index] = 1
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return true
// }

// 解法二 添加缓存，时间复杂度 O(n^2)
func isValidSudoku(board [][]byte) bool {
	//建立緩存
	// 全部遍历一遍，行/列/box是否都只出现过一次
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	//建立兩層的array
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	// 遍历一次，添加缓存

	//先固定橫的row
	for r := 0; r < 9; r++ {
		//循環執行直的column
		for c := 0; c < 9; c++ {
			//如果欄位不是空值
			if board[r][c] != '.' {
				//把数字-1化成索引下标,c是字符串要减去字符串，-1会报错。
				num := board[r][c] - '0' - byte(1)
				//如果橫的row,直的column,九宮格其中一個有存過這個數字(true)
				//直接回傳false
				//九宮格循環,先做一個九宮格的row,i=0就是第一個,,i=1就是第二個
				//除以三以後乘上3就能取到row
				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				//沒有的話將這個數字存成true
				rowbuf[r][num] = true
				colbuf[c][num] = true
				//小方块编号和行列的关系为：r/3*3+c/3
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}
	return true
}

// func isValidSudoku(board [][]byte) bool {
// 	var row, col, sbox [9][9]int
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] != '.' {
// 				num := board[i][j] - '1'
// 				index_box := (i/3)*3 + j/3
// 				//橫的row
// 				if row[i][num] == 1 {
// 					return false
// 				} else {
// 					row[i][num] = 1
// 				}
// 				//直的column
// 				if col[j][num] == 1 {
// 					return false
// 				} else {
// 					col[j][num] = 1
// 				}
// 				//九宮格
// 				if sbox[index_box][num] == 1 {
// 					return false
// 				} else {
// 					sbox[index_box][num] = 1
// 				}
// 			}
// 		}
// 	}
// 	return true
// }
