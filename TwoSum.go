/* Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

import "fmt"

func main() {
	n := twoSum([]int{2, 7, 11, 15}, 26)
	fmt.Println(n)
}

/* //暴力法
func twoSum(nums []int, target int) []int {
	//创建一个初始元素长度为2的数组切片，元素初始值为0：
	res := make([]int, 2)
	for key, value := range nums {
		fmt.Println("loop 1")
		res[0] = key
		fmt.Println("key" ,key)
		fmt.Println("res[0]" , res[0])
		v := target - value
		fmt.Println("v" , v)
		for i := key + 1; i < len(nums); i++ {
			fmt.Println("loop 2")
			fmt.Println("i" , i)
			if nums[i] == v {
				res[1] = i
				fmt.Println("res" , res)
				return res
			}
		}
	}

	return nil

} */

//map法
func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	res := make([]int, 2)

	for k, v := range nums {
		tmp[v] = k
	}

	for k, v := range nums {
		res[0] = k
		value := target - v

		//check the existence of the key in the map
		//如果value值存在tmp這個map中,則返回index值和ok值給ture
		//如果沒有index值則給0,ok值給false
		index, ok := tmp[value]
		//
		if ok && index != k {
			res[1] = index
			break
		}
	}

	return res
}
