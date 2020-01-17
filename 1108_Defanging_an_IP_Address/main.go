/* Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".

Example 1:

Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"

*/

package main

import (
	"fmt"
)

func main() {
	//defangIPaddr("1.1.1.1")
	fmt.Println(defangIPaddr("1.15~.5.1"))
}

/* func defangIPaddr(address string) []string {
	var splits = strings.Split(address, ".")
	return splits
} */

/* func defangIPaddr(address string) string {
	for i := 0; i < len(address); i++ {
		if address[i] == 46 {
			address = address[:i] + "[.]" + address[i+1:]
			i += 3
		}
	}
	return address
} */

/* func defangIPaddr(address string) string {

	a := "."
	b := "[.]"
	rep := ""
	for _, s := range address {
		fmt.Println("s:", s)
		if string(s) == a {
			rep += b
			fmt.Println("rep:", rep)
			fmt.Println("string(s):", string(s))
			continue
		}
		rep += string(s)
	}
	return rep
}
*/

func defangIPaddr(address string) string {
	res := ""

	//s的实际类型是 rune 类型
	//而rune用来表示Unicode的code point
	for _, s := range address {
		fmt.Println("s:", s)

		//rune 类型可以直接等於字符,也可以等於code point
		//逗號的code point是46,可以改成 if s == 46
		//這邊要用''(單引號)來框住逗號
		if s == '.' {
			fmt.Println("enter .")
			res += "[.]"
			fmt.Println("res:", res)
		} else {
			fmt.Println("enter string")
			res += string(s)
			fmt.Println("string(s):", string(s))
		}
	}
	return res
}
