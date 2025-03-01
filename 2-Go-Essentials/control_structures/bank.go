package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	// data, err :=os.ReadFile(accountBalanceFile) // because this func returns two values
	//data, _ := os.ReadFile(accountBalanceFile) // if we don't want to work with the second value we use underscore, otherwise will get an error if we dont use the latest one
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64) //convert string to float
	if err != nil {
		return 1000, errors.New("failed to parse file")
	}

	return balance, nil
}

func main() {
	// var accountBalance float64 = 1000
	var accountBalance, err = getBalanceFromFile()
	fmt.Println("Welcom to GO Bank!")

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		panic("Can't continue sorry.")
	}
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		wantsCheckBalance := choice == 1 // for short conditions we don't need this var, its just an example of bool variable

		if wantsCheckBalance {
			fmt.Printf("The account balance is %.2f\n", accountBalance)
		} else if choice == 2 {
			var depositAmmount float64
			fmt.Print("Your deposit: ")
			fmt.Scan(&depositAmmount)

			if depositAmmount <= 0 {
				fmt.Println("Invalid ammount.")
				// return
				continue
			}
			// accountBalance = accountBalance + depositAmmount
			// short version
			accountBalance += depositAmmount
			fmt.Printf("Balance updated, new ammount: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			var withdrawAmmount float64
			fmt.Print("Your withdraw: ")
			fmt.Scan(&withdrawAmmount)

			if withdrawAmmount <= 0 || withdrawAmmount > accountBalance {
				fmt.Println("Invalid ammont")
				// return
				continue
			}
			accountBalance -= withdrawAmmount
			fmt.Printf("Balance updated, new ammount: %.2f\n", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}

	fmt.Println("Thanks for using our bank")

}
