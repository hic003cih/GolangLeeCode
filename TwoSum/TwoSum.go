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

		//先將nums的index k,存入res[0]
		//下面做相減比對,如果有比對成功的,將另一個index存到res[1]
		res[0] = k
		value := target - v

		//check the existence of the key in the map
		//如果value值存在tmp這個map中,則返回index值和ok值給ture
		//如果沒有index值則給0,ok值給false
		index, ok := tmp[value]
		//如果傳的是true,還有index的值不等於一開始的
		if ok && index != k {
			res[1] = index
			break
		}
	}

	return res
}
([]int{2, 7, 11, 15}, 26)


map[int]int: map[]
[]int, 2: [0 0]
nums: [2 7 11 15]
k: 0
v: 2
nums: [2 7 11 15]
k: 1
v: 7
nums: [2 7 11 15]
k: 2
v: 11
nums: [2 7 11 15]
k: 3
v: 15


res[0]: 0
value: 24
target: 26
v: 2
tmp[value]: 0
index: 0
ok: false


res[0]: 1
value: 19
target: 26
v: 7
tmp[value]: 0
index: 0
ok: false

res[0]: 2
value: 15
target: 26
v: 11
tmp[value]: 3
index: 3
ok: true