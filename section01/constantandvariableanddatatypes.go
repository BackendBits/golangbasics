package main

import (
	"fmt"

	"github.com/BackendBits/golangbasics/commonpackage"
)

var entryAge int = 18

func main() {
	// SECTION 1: Constants, Variables and Basic Data Types in Golang
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 1: Constants, Variables and Basic Data Types in Golang")
	fmt.Println("======================================================================================================")
	fmt.Println()

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
	// result := commonpackage.MarkAge + " years old"

	fmt.Printf("Mark Age is public: %d\n", commonpackage.MarkAge)
	fmt.Printf("John Age is public: %d\n", commonpackage.JohnAge)

	// Issue is public or global variable should start with capital letter in the variable name
	// fmt.Printf("Kevin Age: %d\n", commonpackage.kevinAge)

	fmt.Printf("Entry Age is Global: %d\n", entryAge)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Constants and Variables")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Constants
	fmt.Printf("Value of pi: %f\n", commonpackage.Pi)

	// Uncommenting the line below will result in a compilation error
	// commonpackage.Pi = 3.14159 // Error: cannot assign to pi (declared const)

	// Variables
	fmt.Printf("\nInitial Age: %d\n", commonpackage.MarkAge)
	fmt.Printf("Initial Age: %d\n", commonpackage.MarkAge)

	// Reassigning a new value to the variable
	commonpackage.MarkAge = 30
	fmt.Printf("Updated age: %d\n", commonpackage.MarkAge)

	// Letters can be Unicode letters
	var π = 22 / 7.0
	fmt.Printf("π: %.2f\n", π)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Basic Data Types")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Basic Data Types

	var intValue int = 42
	var int8Value int8 = 8
	var int16Value int16 = -30000
	var int32Value int32 = -2000000000
	var int64Value int64 = -9000000000000000000

	var uintValue uint = 42
	var uint8Value uint8 = 255
	var uint16Value uint16 = 65535
	var uint32Value uint32 = 4294967295
	var uint64Value uint64 = 18446744073709551615

	var float32Value float32 = 3.14
	var float64Value float64 = 3.141592653589793

	var complex64Value complex64 = complex(1.5, 2.5)
	var complex128Value complex128 = complex(3.5, 4.5)

	var stringValue string = "Hello, Go!"

	var boolValue bool = true

	// Output

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println()

	fmt.Println("\nData Types with Different Values:")

	fmt.Printf("Integer Values: %d, %d, %d, %d, %d\n", intValue, int8Value, int16Value, int32Value, int64Value)
	fmt.Printf("Unsigned Integer Values: %d, %d, %d, %d, %d\n", uintValue, uint8Value, uint16Value, uint32Value, uint64Value)
	fmt.Printf("Floating-Point Values: %f, %f\n", float32Value, float64Value)
	fmt.Printf("Complex Values: %v, %v\n", complex64Value, complex128Value)
	fmt.Printf("String Value: %s\n", stringValue)
	fmt.Printf("Boolean Value: %t\n", boolValue)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println()

	fmt.Println("\nData Types:")

	fmt.Println("int:", "Signed integer type with platform-dependent size, typically 32 or 64 bits, holding values in the range of -2^(n-1) to 2^(n-1)-1.")
	fmt.Println("int8:", "Signed 8-bit integer type, ranging from -128 to 127.")
	fmt.Println("int16:", "Signed 16-bit integer type, ranging from -32768 to 32767.")
	fmt.Println("int32:", "Signed 32-bit integer type, ranging from -2^31 to 2^31-1.")
	fmt.Println("int64:", "Signed 64-bit integer type, ranging from -2^63 to 2^63-1.")
	fmt.Println("uint:", "Unsigned integer type with platform-dependent size, typically 32 or 64 bits, holding values in the range of 0 to 2^n-1.")
	fmt.Println("uint8:", "Unsigned 8-bit integer type, ranging from 0 to 255.")
	fmt.Println("uint16:", "Unsigned 16-bit integer type, ranging from 0 to 65535.")
	fmt.Println("uint32:", "Unsigned 32-bit integer type, ranging from 0 to 2^32-1.")
	fmt.Println("uint64:", "Unsigned 64-bit integer type, ranging from 0 to 2^64-1.")
	fmt.Println("float32:", "32-bit floating-point type conforming to IEEE-754 standard, capable of representing a wide range of decimal numbers with moderate precision.")
	fmt.Println("float64:", "64-bit floating-point type conforming to IEEE-754 standard, providing higher precision than float32.")
	fmt.Println("complex64:", "Complex number type with float32 real and imaginary parts.")
	fmt.Println("complex128:", "Complex number type with float64 real and imaginary parts.")
	fmt.Println("string:", "Variable-length string type representing a sequence of characters.")
	fmt.Println("bool:", "Boolean type representing true or false values.")

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println()

}
