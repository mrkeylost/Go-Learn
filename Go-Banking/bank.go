package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do")
	fmt.Println("1. Check the Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice : ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is ", accountBalance)
	} else if choice == 2 {
		var depositAmount float64
		
		fmt.Print("your Deposit : ")
		fmt.Scan(&depositAmount)

		accountBalance += depositAmount
		
		fmt.Println("Current Account Balance: ", accountBalance)
	}
}

func CheckTheBalance () {

}