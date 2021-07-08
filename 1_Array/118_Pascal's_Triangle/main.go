package main

import "fmt"

func main() {
	//var prices = []int{8, 8, 8, 7, 7}
	//target := 3
	fmt.Println(generate(4))

}
func generate(numRows int) [][]int {
	//初始數組裡面的數組行數
	ans := make([][]int, numRows)
	for i := range ans {
		//初始
		ans[i] = make([]int, i+1)
		//設定頭尾都是1
		ans[i][0] = 1
		ans[i][i] = 1
		//每行的最大數量就是該行的行數
		//頭尾都已經設定為1,只要取中間的數就可以
		//把左上跟右上的數字加起來存入
		for j := 1; j < i; j++ {
			//
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
		//fmt.Println(ans)
	}
	return ans
}
