package main

import (
	"errors"
	"fmt"
)

// Function with variadic parameters
func printValues(values ...string) {
	fmt.Println("Printed values:")
	if values == nil {
		fmt.Println("none")
	}

	// First is Index (In this case its given as '_') and Second is Value
	for _, value := range values {
		fmt.Println(value)
	}
}

// Function with named return values
func divide(numerator, denominator int) (result int, err error) {
	if denominator == 0 {
		return 0, errors.New("division by zero")
	}
	return numerator / denominator, nil
}

func main() {

	// SECTION 2: Functions and Control Structures in Golang
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 2: Functions and Control Structures in Golang")
	fmt.Println("======================================================================================================")
	fmt.Println()

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Calling Functions")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Calling Functions
	printValues()

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Function with Parameter")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Function with Parameter
	printValues("Hello, everyone!")

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Function with Variadic Parameters")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Function with Variadic Parameters
	printValues("One", "Two", "Three", "Four")

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Function with Error Handling")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Function with Error Handling
	numerator := 10
	denominator := 2
	result, err := divide(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of %d / %d is: %d\n", numerator, denominator, result)
	}

	// Function with Error Handling (Error Case)
	fmt.Println("Divide By Zero Case")
	denominator = 0
	result, err = divide(numerator, denominator)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result of %d / %d is: %d\n", numerator, denominator, result)
	}

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Conditional Statements")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Conditional Statements
	if result > 5 {
		fmt.Println("Result is greater than 5")
	} else if result == 0 {
		fmt.Println("Result is zero")
	} else {
		fmt.Println("Result is less than 5")
	}

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Switch Statements")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Switch Statements
	switch result {
	case 0:
		fmt.Println("Result is zero")
	case 1, 2:
		fmt.Println("Result is 1 or 2")
	default:
		fmt.Println("Result is not 0, 1, or 2")
	}

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Conditional Switch Statements")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Conditional Switch Statements
	switch {
	case result > 5:
		fmt.Println("Result is greater than 5")
	case result == 0:
		fmt.Println("Result is zero")
	default:
		fmt.Println("Result is less than 5")
	}

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println()
}
