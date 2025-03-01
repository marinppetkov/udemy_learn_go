package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := userInput("Revenue: ")
	expenses := userInput("Expenses: ")
	taxRate := userInput("Tax Rate: ")

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit
	ebt, profit, ratio := calculateValies(revenue, expenses, taxRate)
	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func userInput(text string) float64 {
	var varValue float64
	print(text)
	fmt.Scan(&varValue)
	return varValue
}

func calculateValies(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return ebt, profit, ratio
}
