package main

import (
	"fmt"

	".example/bank/fileops"
	"github.com/pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if(err != nil){
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------------")

		// kalo misal mau stop aplikasi
		// panic("Can't Continue, sorry")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7, ", randomdata.PhoneNumber())

	// bakal jalan terus
	for {
		presentOption()
	
		var choice int
		fmt.Print("Your Choice : ")

		// &schoice adalah pointer ke variable choice
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
			fileops.SaveFloatToFile(accountBalance, accountBalanceFile)
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
				fileops.SaveFloatToFile(accountBalance, accountBalanceFile)
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
