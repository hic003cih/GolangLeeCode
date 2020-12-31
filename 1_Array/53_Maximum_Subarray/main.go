package main

func main() {

}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if num[i-1] > 0 {
			num[i] += num[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
		// if nums[i]+nums[i-1] > nums[i] {
		// 	nums[i] += nums[i-1]
		// }
		// if nums[i] > max {
		// 	max = nums[i]
		// }
	}
	return max

}
