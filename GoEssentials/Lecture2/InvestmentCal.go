package main

import (
	"fmt"
	"math"
)

// func main() {
// 	var investmentAmount = 1000
// 	var expectedReturnRate = 5.5
// 	var years = 10

// 	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
// 	fmt.Println(futureValue)
// }

// func main() {
// 	var investmentAmount float64= 1000
// 	var expectedReturnRate float64= 5.5
// 	var years float64= 10

// 	var futureValue = investmentAmount* math.Pow(1+expectedReturnRate/100, years)
// 	fmt.Println(futureValue)
// }

// func main() {
// 	var investmentAmount,years float64= 1000,10
//     expectedReturnRate := 5.5

// 	futureValue := investmentAmount* math.Pow(1+expectedReturnRate/100, years)
// 	fmt.Println(futureValue)
// }



func main() {

	const inflationRate=2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5
	


	fmt.Print("InvestmentAmount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("ExpectedReturnRate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Years: ")
	fmt.Scan(&years)


	var futureValue = investmentAmount* math.Pow(1+expectedReturnRate/100, years)

	futureRealValue:=futureValue/math.Pow((1+inflationRate/100),years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}


