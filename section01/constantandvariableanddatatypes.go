package main

import (
	"fmt"

	"github.com/BackendBits/golangbasics/commonpackage"
)

func main() {
	// SECTION 1: Constants Variables and Basic Data Types

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Golang is a Statically Typed Language")
	fmt.Println("======================================================================================================")
	fmt.Println()

	/** Golang is a Statically Typed Language **/

	// Explicit declaration with type
	var explicitVariable string = "This is explicitly typed."

	// Inferred type during initialization
	inferredVariable := "This type is inferred."
	var inferredVariable2 = "This type is inferred."

	fmt.Printf("Explicit Variable: %s\n", explicitVariable)
	fmt.Printf("Inferred Variable: %s\n", inferredVariable)
	fmt.Printf("Inferred Variable2: %s\n", inferredVariable2)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Golang is a Strongly Typed Language")
	fmt.Println("======================================================================================================")
	fmt.Println()

	/** Golang is a Strongly Typed Language **/

	// Uncommenting the following line will cause a compilation error
	// result := maleAge + " years old"

	fmt.Printf("Male Age: %d\n", commonpackage.MaleAge)
	fmt.Printf("Female Age: %d\n", commonpackage.FemaleAge)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Constants Variables and Basic Data Types")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Constants
	fmt.Printf("Value of pi: %f\n", commonpackage.Pi)

	// Uncommenting the line below will result in a compilation error
	// commonpackage.Pi = 3.14159 // Error: cannot assign to pi (declared const)

	// Variables
	fmt.Printf("\nInitial Age: %d\n", commonpackage.MaleAge)
	fmt.Printf("Initial Age: %d\n", commonpackage.MaleAge)

	// Reassigning a new value to the variable
	commonpackage.MaleAge = 30
	fmt.Printf("Updated age: %d\n", commonpackage.MaleAge)

	// Letters can be Unicode letters
	var π = 22 / 7.0
	fmt.Printf("π: %.2f\n", π)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Constants Variables and Basic Data Types")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Basic Data Types

	var intValue int = 42
	var int8Value int8 = 8
	var int16Value int16 = 16
	var int32Value int32 = 32
	var int64Value int64 = 64

	var uintValue uint = 42
	var uint8Value uint8 = 8
	var uint16Value uint16 = 16
	var uint32Value uint32 = 32
	var uint64Value uint64 = 64

	var float32Value float32 = 3.14
	var float64Value float64 = 3.14

	var complex64Value complex64 = complex(1, 2)
	var complex128Value complex128 = complex(3, 4)

	var stringValue string = "Hello, Go!"

	var boolValue bool = true

	// Output

	fmt.Println("\nData Types:")
	fmt.Printf("Integer Values: %d, %d, %d, %d, %d\n", intValue, int8Value, int16Value, int32Value, int64Value)
	fmt.Printf("Unsigned Integer Values: %d, %d, %d, %d, %d\n", uintValue, uint8Value, uint16Value, uint32Value, uint64Value)
	fmt.Printf("Floating-Point Values: %f, %f\n", float32Value, float64Value)
	fmt.Printf("Complex Values: %v, %v\n", complex64Value, complex128Value)
	fmt.Printf("String Value: %s\n", stringValue)
	fmt.Printf("Boolean Value: %t\n", boolValue)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println()

}
