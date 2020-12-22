package main

func main() {

}

func reverse(x int) int {

	sum := 0

	/* //取末位,用除以10剩下個位數餘數的方法,例如123除以10,餘3
	lastDigit := x % 10

	//捨去最後的個位數,用除以10的方法,例如123除以10,得到12
	leftDigits := x / 10 */

	/* for _,x = range {
		lastDigit * 10
	} */

	for {
		leftDigits := x / 10 //12
		lastDigit := x % 10  //3
		x = leftDigits       //12

		sum = sum*10 + lastDigit

		if 0 == leftDigits {
			break
		}
	}

	//sum := lastDigit*10 + leftDigits

	//sum = sum*10 + lastDigit

	return sum
}
