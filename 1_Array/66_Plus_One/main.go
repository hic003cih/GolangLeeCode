package main

import "fmt"

func main() {
<<<<<<< HEAD
	digits := []int{1, 3, 5}
	plusOne(digits)
=======
	digits := []int{9}
	fmt.Print(plusOne(digits))
>>>>>>> 8217556e21b16d0fb854f639eba014f7d9e88d92
}

func plusOne(digits []int) []int {

<<<<<<< HEAD
	lastNum := digits[len(digits-1)]
	fmt.Println(lastNum)
=======
	digitLen := len(digits)
	//lastPlusOneDigit := digits[digitLen-1] +1
	if digitLen == 1 {

		if digits[0] == 0 {
			//fmt.Println("hello")
			digits[0] = 1
		} else {
			digits[0] += 1
		}
	} else {
		for i := (digitLen - 1); i > 0; i-- {
			if i == (digitLen - 1) {
				digits[i] += 1
			}
			if digits[i] == 10 {
				digits[i-1] += 1
				digits[i] = 0
			}
		}
	}

	if digits[0] == 10 {
		digits[0] = 1
		digits = append(digits[:digitLen], 0)
	}
	//digits =append(digits[:0],plusOneDigit)
	//return digits
	//digits =append(digits[:digitLen-1],plusOneDigit)
	return digits
>>>>>>> 8217556e21b16d0fb854f639eba014f7d9e88d92
}
