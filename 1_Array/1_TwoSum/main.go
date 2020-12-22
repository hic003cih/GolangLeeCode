package main
)
func main() {
	nums := []int{3, 2, 4}
	target := 6
	x := twoSum(nums, target)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		y := i
		for y = (y + 1); y < len(nums); y++ {
			if target == nums[i]+nums[y] {
				return []int{i, y}
			}
		}
	}
	return nil
}
