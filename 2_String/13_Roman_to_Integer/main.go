package main

import (
	"fmt"
)

/* const I = 1
const V = 5
const X = 10
const L = 50
const C = 100
const D = 500
const M = 1000 */

func main() {

	fmt.Println(romanToInt("MCMXCIV"))
}

/* 1994
1000 +100 -1000 +10 -100 +1 -5

XIV 14

10 - 1 +5

XXVII 27

10 + 10 + 5 + 1+ 1 */

//把羅馬號碼對應到的值加到map表內
//然後字母的部分轉成byte
var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {

	//計算羅馬字串的長度
	//字串可以直接計算長度出來
	n := len(s)
	//迴圈比對每個字母
	for i := range s {
		fmt.Println(i)
		fmt.Println(s[i])
		//s[i]->會取出string每個字母的byte值
		//然後對應到map表內取出對應的數值
		value := symbolValues[s[i]]
		//如果index小於index總長度(表示不是最後一個字母)
		//然後小於index右邊一位羅馬字母的值
		//就要用減的
		//反之則用加的
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}

//自己寫的錯誤版
/* func romanToInt(s string) int {

	//把羅馬號碼對應到的值加到map表內
	romanMap := map[string]int{}
	romanMap["I"] = 1
	romanMap["V"] = 5
	romanMap["X"] = 10
	romanMap["L"] = 50
	romanMap["C"] = 100
	romanMap["D"] = 500
	romanMap["M"] = 1000
	romanMap["IV"] = 4
	romanMap["IX"] = 9
	romanMap["XL"] = 40
	romanMap["XC"] = 90
	romanMap["XL"] = 40
	romanMap["CD"] = 400
	romanMap["CM"] = 900
	//預設加總的值
	sum := 0
	for i, v := range s {
		fmt.Println(i)
		fmt.Println(v)
	}
	//先比對是不是特殊的減法
	x, ok := romanMap[s]
	if ok {
		sum = sum + x
	} else {

		//分解每個字母
		temp := strings.Split(s, "")
		//迴圈比對分解的每個字母
		//如果有在羅馬號碼map表
		//則加上對應的值
		for _, v := range temp {
			x, ok := romanMap[v]
			if ok {
				sum = sum + x
			}
		}
	}

	return sum

} */

//正確版
/* 通常情况下，罗马数字中小的数字在大的数字的右边。若输入的字符串满足该情况，那么可以将每个字符视作一个单独的值，累加每个字符对应的数值即可。

例如 \texttt{XXVII}XXVII 可视作 \texttt{X}+\texttt{X}+\texttt{V}+\texttt{I}+\texttt{I}=10+10+5+1+1=27X+X+V+I+I=10+10+5+1+1=27。

若存在小的数字在大的数字的左边的情况，根据规则需要减去小的数字。对于这种情况，我们也可以将每个字符视作一个单独的值，若一个数字右侧的数字比它大，则将该数字的符号取反。

例如 \texttt{XIV}XIV 可视作 \texttt{X}-\texttt{I}+\texttt{V}=10-1+5=14X−I+V=10−1+5=14。
*/
/* var symbolValues = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) (ans int) {
    n := len(s)
    for i := range s {
        value := symbolValues[s[i]]
        if i < n-1 && value < symbolValues[s[i+1]] {
            ans -= value
        } else {
            ans += value
        }
    }
    return
} */

//自己寫的
/* //把羅馬號碼對應到的值加到map表內
romanMap := map[string]int{}
romanMap["I"] = 1
romanMap["V"] = 5
romanMap["X"] = 10
romanMap["L"] = 50
romanMap["C"] = 100
romanMap["D"] = 500
romanMap["M"] = 1000
romanMap["IV"] = 4
romanMap["IX"] = 9
romanMap["XL"] = 40
romanMap["XC"] = 90
romanMap["XL"] = 40
romanMap["CD"] = 400
romanMap["CM"] = 900
//預設加總的值
sum := 0
for _, v := range romanMap {
	//分解每個字母
	temp = strings.Split(s, v)
}
//先比對是不是特殊的減法
x, ok := romanMap[s]
if ok {
	sum = sum + x
} else {

	//分解每個字母
	temp := strings.Split(s, "")
	//迴圈比對分解的每個字母
	//如果有在羅馬號碼map表
	//則加上對應的值
	for _, v := range temp {
		x, ok := romanMap[v]
		if ok {
			sum = sum + x
		}
	}
}

return sum */
