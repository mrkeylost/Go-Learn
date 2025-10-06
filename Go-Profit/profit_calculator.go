package main

import (
	"fmt"
)

func main() {
	var revenue, expenses float64
	taxRate := 11.0

	revenue = inputText("Revenue : ")
	expenses = inputText("Expenses : ")
	taxRate = inputText("Tax Rate : ")

	EBT, profit, ratio := calculate(revenue, expenses, taxRate)

	// EBT := revenue - expenses
	// profit := EBT - (EBT * (taxRate / 100))

	// fmt.Println("EBT : ", EBT)
	// fmt.Println("Profit : ", profit)

	// ratio := EBT / profit

	// fmt.Println("Ratio : ", ratio)

	fmt.Printf("EBT : %v\n", EBT)
	fmt.Printf("Profit : %v\n", profit)

	fmt.Printf("Ratio : %.2f\n", ratio)
}

func inputText(text string) float64 {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

func calculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt/profit 

	return ebt, profit, ratio
}