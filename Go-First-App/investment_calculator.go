package main

import (
	"fmt"
	"math"
)

func main() {
	// explicit type
	// var investmentAmount float64 = 1000
	// var expectedReturnRate = 5.5
	// var years float64 = 10

	// alternative variables

	// var investmentAmount, years  = 1000, "10" 
	// var investmentAmount, years float64 = 1000, 10

	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0,  5.5

	// Kalo gk ada value, selalu pake var sama datatype-nya
	var investmentAmount float64
	var years float64
	expectedReturnRate :=  5.5

	const inflationRate = 2.5

	// var futureValue = float64(investmentAmount) * math.Pow(1 + (expectedReturnRate) / 100, float64(years))

	// need user input
	// jangan lupa pake &
	fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	fmt.Print("How long is the investment : ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + (expectedReturnRate) / 100, years)

	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}