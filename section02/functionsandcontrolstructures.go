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
	for _, value := range values {
		fmt.Println(value)
	}
}

// Function with named return values
func divideWithNamedResult(numerator, denominator int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("panic occurred")
		}
	}()

	if denominator == 0 {
		panic("division by zero")
	}

	result = numerator / denominator
	return result, nil
}

func main() {

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
	fmt.Println("Function with Named Return Values and Defer")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Function with Named Return Values and Defer
	numerator := 10
	denominator := 2
	result, err := divideWithNamedResult(numerator, denominator)
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
