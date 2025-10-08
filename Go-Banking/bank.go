package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	// _ dipake buat mau handle value, tapi belom mau dipake
	// kalo pake variable, bakal ada error karena unused
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read file.")
	}

	balanceText := string(data)
	// 2 params
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func saveBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	
	// 0644 : file permission
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if(err != nil){
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")

		// kalo misal mau stop aplikasi
		// panic("Can't Continue, sorry")
	}

	fmt.Println("Welcome to Go Bank!")

	// bakal jalan terus
	for {
		fmt.Println("What do you want to do")
		fmt.Println("1. Check the Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Print("Your Choice : ")
		fmt.Scan(&choice)
	
		// wantsCheckBalance := choice == 1

		// if-else statement
	
		// if choice == 1 {
		// 	fmt.Println("Your balance is ", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
			
		// 	fmt.Print("your Deposit : ")
		// 	fmt.Scan(&depositAmount)
	
		// 	if depositAmount <= 0 {
		// 		fmt.Print("Invalid Amount, must be greater than 0.")
		// 		continue
		// 	}
	
		// 	accountBalance += depositAmount
			
		// 	fmt.Println("Current Account Balance: ", accountBalance)
		// } else if choice == 3 {
		// 	var withdrawAmount float64
	
		// 	fmt.Println("Current Account Balance: ", accountBalance)
		// 	fmt.Print("Your withdraw : ")
		// 	fmt.Scan(&withdrawAmount)
	
		// 	if withdrawAmount <= 0 {
		// 		fmt.Print("Invalidate Amount, must be greater than 0.")
		// 		continue
		// 	}
	
		// 	if withdrawAmount > accountBalance {
		// 		fmt.Println("insufficient Balance!")
		// 	} else {
		// 		accountBalance -= withdrawAmount
	
		// 		fmt.Println("Current Account Balance: ", accountBalance)
		// 	}
		// } else {
		// 	fmt.Println("Goodbye!")
		// 	break
		// }

		// switch statement

		switch choice {
		case 1: 
			fmt.Println("Your balance is ", accountBalance)	

		case 2:
			var depositAmount float64
			
			fmt.Print("your Deposit : ")
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Print("Invalid Amount, must be greater than 0.")
				continue
			}
	
			accountBalance += depositAmount
			
			fmt.Println("Current Account Balance: ", accountBalance)
			saveBalanceToFile(accountBalance)
		case 3:
			var withdrawAmount float64
	
			fmt.Println("Current Account Balance: ", accountBalance)
			fmt.Print("Your withdraw : ")
			fmt.Scan(&withdrawAmount)
	
			if withdrawAmount <= 0 {
				fmt.Print("Invalidate Amount, must be greater than 0.")
				continue
			}
	
			if withdrawAmount > accountBalance {
				fmt.Println("insufficient Balance!")
			} else {
				accountBalance -= withdrawAmount
	
				fmt.Println("Current Account Balance: ", accountBalance)
				saveBalanceToFile(accountBalance)
			}

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for chhosing our bank!")
			// break di swicth, cuma keluar dari switch
			// break
			return
		}
	}
}

func CheckTheBalance () {

}