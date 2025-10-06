package main

import (
	"fmt"
	"math"
)

// Global Const
const inflationRate = 2.5

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

	// var futureValue = float64(investmentAmount) * math.Pow(1 + (expectedReturnRate) / 100, float64(years))

	// need user input
	// jangan lupa pake &

	printText("Investment Amount : ")
	// fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	printText("How long is the investment : ")
	// fmt.Print("How long is the investment : ")
	fmt.Scan(&years)

	printText("Expected Return Rate : ")
	//fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	// futureValue := investmentAmount * math.Pow(1 + (expectedReturnRate) / 100, years)

	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	// Ambil hasil multiple return function disini
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years) 

	// // formatted string
	// formattedFV := fmt.
	// 
	// Sprintf("Future Value : %.2f\n", futureValue)
	// formattedFRV := fmt.Sprintf("Future Value : %.2f\n", futureRealValue)

	// fmt.Println("Future Value : ",futureValue)
	// fmt.Println("Future Value After Inflation : ", futureRealValue)

	fmt.Printf("Future Value : %.2f\n", futureValue)
	fmt.Printf("Future Value : %.2f\n", futureRealValue)

	// fmt.Print(formattedFV, formattedFRV)

	// fmt.Printf(`Future Value : %.2f
	// Future Value : %.2f`, futureValue, futureRealValue)
}

// functions

func printText(text string){
	fmt.Print(text)
} 

// return multiple values from function
// func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {

// 	fv := investmentAmount * math.Pow(1 + (expectedReturnRate) / 100, years)

// 	rfv := fv / math.Pow(1 + inflationRate/100, years)

// 	return  fv, rfv
// }

// alternate syntax 
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {

	fv = investmentAmount * math.Pow(1 + (expectedReturnRate) / 100, years)

	rfv = fv / math.Pow(1 + inflationRate/100, years)

	// prefer gini aj biar keliatan yang direturn
	return  fv, rfv

	// otomatis return 2 data yang dibracket 2
	// return
}