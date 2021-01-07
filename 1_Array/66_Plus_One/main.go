package main

import "fmt"

func main() {
	digits := []int{1, 3, 5}
	plusOne(digits)
}

func plusOne(digits []int) []int {

	lastNum := digits[len(digits-1)]
	fmt.Println(lastNum)
}
