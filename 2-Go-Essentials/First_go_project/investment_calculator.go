package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// const inflationRate = 2.5 // moving to global scope
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	var years float64 = 10

	// fmt.Print("Investment Ammount: ")
	outputText("Investment Ammount:")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// moved to calculateFutureValue function
	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// output information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future real value: ", futureRealValue)
	// fmt.Printf("Future Value: %v\nFuture real value: %v\n", futureValue, futureRealValue)
	// fmt.Printf("Future Value: %.2f\nFuture real value: %.1f\n", futureValue, futureRealValue)

	formattedFV := fmt.Sprintf("Future Values: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value: %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// alternative syntax
// func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }
