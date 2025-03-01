# Execution of the program
`go run app.go` or `go run .`

## Essentials
### First_go_program folder
1. Functions

* Like python the command is a function but its preceded with the package name
* Unlike python where the print function is built-in, Go requires explicit imports to keep core language small and efficient
* But we still have [builtin packages](https://pkg.go.dev/builtin). </br> 
```GO
package main

import "fmt"

func main() {
	fmt.Print("Hello world")
}
```
In this code the program start running package main and imports package fmt from go [Standard library](https://pkg.go.dev/std).

2. Packages

Every GO code needs package instruction like `package main` to organize the code. </br>
`main` is a special package name, its the **Entry point for Execution**. </br>
A GO program **must** have `main` package to **compile** into an executable. </br>

3. GO Modules 

[Documentation](https://go.dev/blog/using-go-modules)
> A module is a collection of Go packages stored in a file tree with a go.mod file at its root.

One module may consists of multiple packages. A module can be viewed as GO project </br>

* `go mod init example.com/first-app` a go.mod file is added
* `go build` we get first-app executable file and run it `./first-app` </br>
If we don't have the main package the executable will not be build


4. The `main` function

GO will execute that function when the program starts, so its important to have it. </br>
We can have only one main function accross two files that belongs to the same package </br>

### First_go_project folder

5. Variables

We use the var key word and the naming convetion is camel case notation (not mandatory) </br>
```GO
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10
```

6. Value types

Go is staticly typed language, so we may need to define type or help go to [infer](https://go.dev/tour/basics/14) the type.
```GO
	var investmentAmount = 1000.0
	var expectedReturnRate = 5.5
	var years = 10
	var futureValue = investmentAmount * (1 + expectedReturnRate/100)
```
or convert the integer to float with **note** the **built-in** function float64
```GO
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
```

7. Output

`fmt.Println(futureValue)`

8. Type Conversions and Explicit Type Assignment

As GO infers the value types we may need to assing a type explicitly to overweite the infered value
```GO
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10
	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
```

9. Alternative Variable declaration Styles

* If we don't have explicit var type we can ommit the `var` keyword and use an assignment operator `:=`.
Its recommanded to use this operator when the type should be infered by GO
```GO
expectedReturnRate := 5.5
```

* We can declare multiple variables in the same line
```GO
var investmentAmount, years float64 = 1000, 10
```
GO can still infer the values. In the below example the first is int and the second is string
```GO
var investmentAmount, years = 1000, "10"
```
Note that you cannot add multiple type assingments in one line

* We can do the same with assignment operator
```GO
investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
```

10. Constants

[Documentation](https://go.dev/tour/basics/15)
The variables can be reassigned unlike the constants, which cannot be changed after.

11. Fetch user input

We use the `Scan` function </br>
The pointer is important 
```GO
// & -> pointer
fmt.Scan(&investmentAmount)
```

12. Improve user input experience

```GO
	fmt.Print("Investment Ammount: ")
	fmt.Scan(&investmentAmount)
```
What about python?
```Python
investmentAmount = input("Investment Ammount: ")
```

13. Formatting strings

* We can pass multiple values
```GO
fmt.Println("Future Value: ", futureValue)
```
* We can print text with certain format with `Printf` function
%v -> placeholder of outputing value of variable or constant </br>
More about it [here](https://pkg.go.dev/fmt@go1.24.0#hdr-Printing) </br>
```GO
// Printf does not add line break so we may need to add it
fmt.Printf("Future Value: %v\n", futureValue)
```
or more placeholder with the values followed in the order we want to show
```GO
fmt.Printf("Future Value: %v\nFuture real value: %v\n", futureValue, futureRealValue)
```

14. Formtatting Floats in Strings

Same doc [here](https://pkg.go.dev/fmt@go1.24.0#hdr-Printing) and looks like C
```
%f     default width, default precision
%9f    width 9, default precision
%.2f   default width, precision 2
%9.2f  width 9, precision 2
%9.f   width 9, precision 0
```
```GO
fmt.Printf("Future Value: %.2f\nFuture real value: %.1f\n", futureValue, futureRealValue)
```

15. Creating Formatted strings

Sprintf formats according to a format specifier and returns the resulting string. </br>

```GO
	formattedFV := fmt.Sprintf("Future Values: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future real value: %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
```

16. Building multiple strings

We use the backtick character
```GO
fmt.Printf(`Future Value: %.2f\n
Future real value: %.1f\n`, futureValue, futureRealValue)
```
With backticks the line breaks no longer work!

17. Functions

* In python the function is defined with `def name():` in GO its `func name(){}`
* When adding parameter to a function we must define value type
```GO
func outputText(text string) {
	fmt.Println("some text")
}
// multiple values
func outputText2(text string, text2 string) {
	fmt.Println("some text")
}
// multiple values with the same type
func outputText3(text, text2 string) {
	fmt.Println("some text")
}
```

18. Function return values and scope

We can use global scope by defining variables and const outside the main or any other function. </br>
If the variables are scoped globally, they must be defined with the `var` keyword. </br>

In GO we can return multiple values sperated by comma

```GO
func calculator(){
	return value1, value2
}
```
Its important to define the return type
```GO
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
```
or in case its one value
```GO
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) float64 {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return rfv
}
```
In the main function we call our func and pass the values
```GO
futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
```

19. Alternative return value syntax

```GO
func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) { //that will instruct GO to create these values
	// since they are already created, we just need to assign the values and don't need to init the fv and rfv values with :=
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv // we can still use it, but we can just use return and go will return the declared variables
	return
}
```

### control_structures folder
20. Cotrol structures
- if, if else, else statement </br>
`else` should be on the same line as the closing curly bracket! 
```GO
	wantsCheckBalance := choice == 1 // for short conditions we don't need this var, its just an example of bool variable
	if wantsCheckBalance {
		fmt.Printf("The account balance is %.2f\n", accountBalance)
	} else if choice == 2 {
		var depositAmmount float64
		fmt.Print("Your deposit: ")
		fmt.Scan(&depositAmmount)
		accountBalance += depositAmmount
		fmt.Printf("Balance updated, new ammount: %.2f\n", accountBalance)
	} else if choice == 3 {
		var withdrawAmmount float64
		fmt.Print("Your withdraw: ")
		fmt.Scan(&withdrawAmmount)
		accountBalance -= withdrawAmmount
		fmt.Printf("Balance updated, new ammount: %.2f\n", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}
```
-  Nested "if" Statements & Using "return" To Stop Function Execution </br>
With the `return` keyword we exit the function and no other code is executed:
```GO
	if wantsCheckBalance {
		fmt.Printf("The account balance is %.2f\n", accountBalance)
	} else if choice == 2 {
		var depositAmmount float64
		fmt.Print("Your deposit: ")
		fmt.Scan(&depositAmmount)

		if depositAmmount <= 0 {
			fmt.Println("Invalid ammount.")
			return
		}
		accountBalance += depositAmmount
		fmt.Printf("Balance updated, new ammount: %.2f\n", accountBalance)
	}
```

- Loops </br>

In GO we can create only `for` [loops](https://go.dev/tour/flowcontrol/1). </br>
>the init statement: executed before the first iteration </br>
>the condition expression: evaluated before every iteration </br>
>the post statement: executed at the end of every iteration </br>
Similar to C and PowerShell
```GO
	for i := 0; i < 10; i++ {
		sum += i
	}
```

- Infinitive loops </br>
We just define the loop without any components 
```GO
for {
	println("Infitive loop")
	// unless we use return
	//return
	// or break
	//break
}
```
Keywords `break` and `continue`

- Conditional For loops
```GO
for someCondition {
  // do something ...
}
```
>someCondition is an expression that yields a boolean value or </br> a variable that contains a boolean value (i.e., true or false).

21. Switch
* Example [1](https://go.dev/tour/flowcontrol/9)
* Example [2](https://www.geeksforgeeks.org/switch-statement-in-go/)
```GO
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

```
Python uses match as from 3.10
```python
match term:
    case pattern-1:
         action-1
    case pattern-2:
         action-2
    case pattern-3:
         action-3
    case _:
        action-default
```
In `C` use break to exit the block and skip other and default cases, same in PowerShell </br>
But `the break statement that is needed at the end of each case in those languages is provided automatically in Go`
```C
#include <stdio.h>

int main() {
  
    // Switch variable
    int var = 1;

    // Switch statement
    switch (var) {
    case 1:
        printf("Case 1 is Matched.");
        break;

    case 2:
        printf("Case 2 is Matched.");
        break;

    case 3:
        printf("Case 3 is Matched.");
        break;

    default:
        printf("Default case is Matched.");
        break;
    }

    return 0;
}
```

Note that in `GO` if we add a `break` statement inside switch, this will not break a loop, so a `return` is better option to exit the function, unlike [Python](https://stackoverflow.com/questions/72273235/how-to-break-the-match-case-but-not-the-while-loop)

22. Writing to files
```GO
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}
```
[[]byte](https://medium.com/@tyler_brewer2/bits-bytes-and-byte-slices-in-go-8a99012dcc8f)
> Byte slices are a list of bytes that represent UTF-8 encodings of Unicode code points. </br>
`%d` - prints UTF-8 decimal value of each byte </br>
`%s` - converts the byte slice to a string </br>
```GO 
s := "Wow look at me"
bs := []byte(s)
fmt.Printf("%s", bs) // Output: Wow look at me
fmt.Printf("%d", bs) // Output: [87 111 119 32 108 111 111 107 32 97 116 32 109 101]
```

23. Reading files
`func os.ReadFile(name string) ([]byte, error)`
```GO
func getBalanceFromFile() float64 {
	// data, err :=os.ReadFile(accountBalanceFile) // because this func returns two values
	data, _ := os.ReadFile(accountBalanceFile) // if we don't want to work with the second value we use underscore, otherwise will get an error if we dont use the latest one
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64) //convert string to float
	return balance
}
```

24. Handling errors

In GO we don't use try catch statement </br>
In GO functions are writtent that the error don't crash the application </br>
We have special type in GO called `error` </br>
We have error package, with which we can create new error and return it with the `error` type </br>
This is how we create functions with error handling </br>
```Go
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Faield to read file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64) //convert string to float
	if err != nil {
		return 1000, errors.New("Failed to parse file")
	}

	return balance, nil
}
```

25. Panic
We can exit an error with return(not necessary can just leave it that way) or use panic which contains which part of the code has failed
```GO
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		panic("Can't continue sorry.")
	}
```
```shell
Welcom to GO Bank!
ERROR
Faield to read file
---------
panic: Can't continue sorry.

goroutine 1 [running]:
main.main()
	/Users/marinpetkov/Lab/hashicorpTrainings/code/GO/udemy_learn_go/2-Go-Essentials/control_structures/bank.go:42 +0x68c
exit status 2
```