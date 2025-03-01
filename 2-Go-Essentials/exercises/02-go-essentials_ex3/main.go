package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
// => Show error mesage and exit if invalid input is provided
// - No negative numbers
// - Not 0
// - 2) store results into file

func main() {
	revenue, errRevenue := getUserInput("Revenue: ")
	expenses, errExpenses := getUserInput("Expenses: ")
	taxRate, errTaxRate := getUserInput("Tax Rate: ")

	if errRevenue != nil || errExpenses != nil || errTaxRate != nil {
		panic("invalid value provided")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writeFile(ebt, profit, ratio)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func writeFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("ebt=%.2f, profit=%.2f, ratio=%.2f", ebt, profit, ratio)
	os.WriteFile("data.log", []byte(results), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return userInput, errors.New("invalid input, must be value bigger than zero")
	}
	return userInput, nil
}
