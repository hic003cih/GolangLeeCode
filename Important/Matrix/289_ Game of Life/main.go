//主要是结构化，把计算周围死活细胞抽离成函数，然后依次写判断条件即可
//注释掉的判断条件是因为符合细胞状态不变这一特点，因此没必要加

// func gameOfLife(board [][]int, target int) {
// 	//1 原来活现在活
// 	//0 原来死现在死
// 	//2 原来死现在活
// 	//-1 原来活现在死
// 	for i, rowArr := range board {
// 		for j, _ := range rowArr {
// 			temp := isAlive(i-1, j-1, board) + isAlive(i-1, j, board) + isAlive(i-1, j+1, board) + isAlive(i, j-1, board) + isAlive(i, j+1, board) + isAlive(i+1, j-1, board) + isAlive(i+1, j, board) + isAlive(i+1, j+1, board)
// 			if temp < 2 && board[i][j] == 1 {
// 				board[i][j] = -1
// 				// } else if ((temp == 2 || temp == 3) && board[i][j] == 1) {
// 				// board[i][j] = 1
// 			} else if temp > 3 && board[i][j] == 1 {
// 				board[i][j] = -1
// 			} else if temp == 3 && board[i][j] == 0 {
// 				board[i][j] = 2
// 			}
// 		}
// 	}
// 	for i, rowArr := range board {
// 		for j, _ := range rowArr {
// 			if board[i][j] == 1 || board[i][j] == 2 {
// 				board[i][j] = 1
// 			} else {
// 				board[i][j] = 0
// 			}
// 		}
// 	}
// }

// //判斷是否還存活
// func isAlive(i int, j int, board [][]int) int {
// 	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
// 		return 0
// 	}
// 	if board[i][j] == 1 || board[i][j] == -1 {
// 		return 1
// 	}
// 	return 0
// }

// 遍历周围八个点

// 注意：不是遍历周边最新状态，而是最初的状态
//規定要取最初的狀態來作全部的判斷,全部判斷完以後才統一一起更新
func gameOfLife(board [][]int) {
	//一个最简单的解决方法就是复制一份原始数组，复制的那一份永远不修改，只作为更新规则的引用
	//建立暫存board的array row的長度
	temp := make([][]int, len(board))
	//建立內部的array board col的長度
	for i := 0; i < len(board); i++ {
		temp[i] = make([]int, len(board[i]))
	}
	//先迴圈row
	for i := 0; i < len(board); i++ {
		//再迴圈col
		for j := 0; j < len(board[i]); j++ {
			//用來儲存周圍數值的總合
			nums := 0
			//設定大於等於0和小於board長度,防止溢出長度
			// 上面
			if i-1 >= 0 {
				nums += board[i-1][j]
			}
			// 左面
			if j-1 >= 0 {
				nums += board[i][j-1]
			}
			// 下面
			if i+1 < len(board) {
				nums += board[i+1][j]
			}
			// 右面
			if j+1 < len(board[i]) {
				nums += board[i][j+1]
			}
			// 左上
			if i-1 >= 0 && j-1 >= 0 {
				nums += board[i-1][j-1]
			}
			// 右上
			if i-1 >= 0 && j+1 < len(board[i]) {
				nums += board[i-1][j+1]
			}
			// 左下
			if i+1 < len(board) && j-1 >= 0 {
				nums += board[i+1][j-1]
			}
			// 右下
			if j+1 < len(board[i]) && i+1 < len(board) {
				nums += board[i+1][j+1]
			}
			//一个最简单的解决方法就是复制一份原始数组，复制的那一份永远不修改，只作为更新规则的引用
			temp[i][j] = board[i][j]
			//變更副本
			switch {
			//如果周圍活著的少於2,則死亡
			case nums < 2:
				temp[i][j] = 0
			//如果周圍活著的等於3或0,則活著
			case nums == 3 && temp[i][j] == 0:
				temp[i][j] = 1
			//如果周圍活著的超過3,則死亡
			case nums > 3:
				temp[i][j] = 0
			}
		}
	}
	//作完全部的判斷,將副本複製到原始array中覆蓋
	copy(board, temp)
}
