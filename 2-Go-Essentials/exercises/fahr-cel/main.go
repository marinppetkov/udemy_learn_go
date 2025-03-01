package main

import (
	"errors"
	"fmt"
)

func main() {
	var choice int
	for {
		fmt.Println("Make a choice")
		fmt.Println("1. Convert from Fahrenheit to Celsius")
		fmt.Println("2. Convert from Celsius to Fahrenheit")
		fmt.Println("3. Exit choose 3")
		fmt.Print("Your choice:")
		fmt.Scan(&choice)
		if choice == 1 {
			// fmt.Println(fahrToCel())
			fahrToCel, err := fahrToCel()
			if err != nil {
				panic("Wrong input for fahrenheit")
			}
			fmt.Println(fahrToCel)
		} else if choice == 2 {
			celToFahr, err := celToFahr()
			if err != nil {
				panic("Wrong input for celsius")
			}
			fmt.Println(celToFahr)
		} else if choice == 3 {
			return
		} else {
			fmt.Println("Wrong choice.\nPlease choose between 1 or 2.")
		}
	}
}

func fahrToCel() (result string, errs error) {
	var fahrenheit float64
	fmt.Print("Enter fahrenheit: ")
	_, err := fmt.Scan(&fahrenheit)
	if err != nil {
		fmt.Println(err)
		// panic("Invalid value for fahrenheit")
		return "something", errors.New("invalid valie for fahrenheit")
	}
	result = fmt.Sprintf("The %.2f in fahrenheit is %.2f in celsius", fahrenheit, (fahrenheit-32)*5/9)
	return result, nil
}

func celToFahr() (string, error) {
	var celsius float64
	fmt.Print("Enter celsius: ")
	// fmt.Scan(&celsius)
	_, err := fmt.Scan(&celsius)
	if err != nil {
		fmt.Println(err)
		return "Invalid value for celsius", errors.New("invalid value for celsius")
	}
	result := fmt.Sprintf("The %.2f in celsius is %.2f in fahrenheit", celsius, celsius*9/5+32)
	return result, nil
}
