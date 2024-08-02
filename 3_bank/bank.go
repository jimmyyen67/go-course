package main

import (
	"example.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("-------------------------------")
		fmt.Println("ERROR:", err)
		fmt.Println("-------------------------------")
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()
		var choice int
		fmt.Print("Your choice: ")
		_, errScan := fmt.Scan(&choice)
		if errScan != nil {
			return
		}

		switch choice {
		case 1:
			fmt.Println("Your balance is ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			_, errScan := fmt.Scan(&depositAmount)
			if errScan != nil {
				return
			}

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("How much do you want to withdraw? ")
			var withdrawAmount float64
			_, errScan := fmt.Scan(&withdrawAmount)
			if errScan != nil {
				return
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Your balance is not enough. PLease try again.")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Withdraw success! Your remain amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our 2_bank.")
			return
		}
	}
}
