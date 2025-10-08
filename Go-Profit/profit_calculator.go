package main

import (
	"fmt"
	"os"
	"errors"
)

// Goals
// validate user input
// shows error message and exit if invalid input is provided
// no negative numbers
// no 0
// Store calculated result into file

func saveToFile (EBT, profit, ratio float64) {
	result := fmt.Sprintf("EBT : %.2f\nProfit : %.2f\nRatio : %.2f", EBT, profit, ratio)

	os.WriteFile("result.txt", []byte(result), 0644)
}

func main() {
	var revenue, expenses float64
	taxRate := 11.0
	var err error

	revenue, err = inputText("Revenue : ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		return
	}

	expenses, err = inputText("Expenses : ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		return
	}

	taxRate, err = inputText("Tax Rate : ")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")
		return
	}

	EBT, profit, ratio := calculate(revenue, expenses, taxRate)
	saveToFile(EBT, profit, ratio)

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

func inputText(text string) (float64, error) {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	if(userInput < 0){
		return 0, errors.New("Invalid Item")	
	}

	return userInput, nil
}

func calculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt/profit 

	return ebt, profit, ratio
}