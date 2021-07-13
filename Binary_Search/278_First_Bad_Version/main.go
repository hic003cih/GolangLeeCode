/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

package main

import "fmt"

func main() {
	n := 8
	bad := 4

	fmt.Println(firstBadVersion(n))
}
func firstBadVersion(int n) {
	left := 0
	right := n

	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left

}

//做不出來版
/*func firstBadVersion(n int) int {

		if n==1 {
			return 1
		}
		mid := n /2
		if isBadVersion(mid){
			mid--

		}

		for i := mid; i <= n; i++ {
			if isBadVersion(i){
				return i
			}
		}

		for i := 1; i < mid; i++ {
			if isBadVersion(i){
				return i
			}
		}

	return n
}*/

//正確版
/*func firstBadVersion(n int) int {
	left, right := 1, n
	// 循环直至右邊小於左邊時
	//表示已經找完了,
	for left<=right{
		//找到中間值
		mid:=(left+right)/2
		//如果中間值是錯誤的版本
		//代表右邊指針要繼續往左邊找最初的錯誤版本
		//所以要-1
		if isBadVersion(mid){
			right = mid-1 // 答案在区间 [left, mid] 中
		}else{
		//如果中間值不是錯誤的版本
		//代表左邊指針要繼續往右邊找最初的錯誤版本
		//所以要+1
			left = mid+1 // 答案在区间 [mid+1, right] 中
		}
	}
	// 此时有 left == right，区间缩为一个点，即为答案
	return left
}
*/
