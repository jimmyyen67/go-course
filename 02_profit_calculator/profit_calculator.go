package main

import (
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//    => Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

func main() {
	revenue, err := getUserInput("Revenue")
	if err != nil {
		fmt.Println("ERROR:", err)
		panic("Can't continue, sorry.")
	}
	expenses, err := getUserInput("Expenses")
	if err != nil {
		fmt.Println("ERROR:", err)
		panic("Can't continue, sorry.")
	}
	taxRate, err := getUserInput("Tax Rate")
	if err != nil {
		fmt.Println("ERROR:", err)
		panic("Can't continue, sorry.")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.2f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0664)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	var error error = nil
	fmt.Print(infoText, ": ")
	fmt.Scan(&userInput)

	if userInput < 0 {
		error = fmt.Errorf("user's input should greater than 0, %.2f given", userInput)
	} else if userInput == 0 {
		error = fmt.Errorf("user's input should not be 0")
	}

	return userInput, error
}
