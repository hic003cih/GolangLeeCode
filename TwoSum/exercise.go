package main

var nums = []int{4, 8, 9, 20}

var target = 28

func main() {
	TwoSum(nums, target)
}

func TwoSum(nums []int, target int) []int {
	tmp := make(map[int]int)
	res := make([]int, 2)

	for v, k := range nums {
		tmp[v] = k
	}

	for v, k := range nums {
		res[0] = v
		value := target - v

		index, ok := tmp[value]

		if ok,value
	}
}
