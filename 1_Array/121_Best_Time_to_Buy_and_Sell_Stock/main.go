package main

import "fmt"

func main(){
	var prices =[]int{0,1}
	fmt.Println(maxProfit(prices))

}


func maxProfit(prices []int) int {

	//每個數比對
	//max :=0
	// for leftPrices := 0; leftPrices <= len(prices)-1; leftPrices++ {
	// 	for rightPrices := (leftPrices+1); rightPrices <= len(prices)-1; rightPrices++ {
	// 		if (prices[rightPrices]-prices[leftPrices]) > max{
	// 			max = prices[rightPrices]-prices[leftPrices]
	// 		}
	// 	}
	// }
	//return max

	if len(prices) == 0 {
        return 0
    }
    Min := prices[0]        // 记录当前股票最小价格
	Max := 0                // 记录当前最大获利
	//循環比對最大值的時候
	//同時找出最小直
	//因為只能往右找,所以不用每個都比對
    for i:=1; i<len(prices); i++ {
        if prices[i] - Min > Max {
            Max = prices[i] - Min
        }
        if prices[i] < Min {
            Min = prices[i]
        }
    }
    return Max
}

