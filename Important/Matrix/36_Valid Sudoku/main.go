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
	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}

	x := isValidSudoku(matrix)
	fmt.Print(x)
}

func isValidSudoku(board [][]byte) bool {

}
