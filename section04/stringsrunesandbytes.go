package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Strings and UTF-8 Characters
	myString := "Hello, 世界!"
	fmt.Println("String:", myString)

	// Indexing string reveals UTF-8 byte values
	fmt.Printf("Index 2: %v\n", myString[2])

	// Iterating over string using range decodes multi-byte characters
	fmt.Println("Iterating over runes:")
	for index, char := range myString {
		fmt.Printf("Index: %d, Rune: %c, Unicode Point: %U\n", index, char, char)
	}

	// Using strings.Builder for efficient string building
	var strBuilder strings.Builder

	strBuilder.WriteString("This is ")
	strBuilder.WriteString("a concatenated ")
	strBuilder.WriteString("string.")

	resultString := strBuilder.String()
	fmt.Println("Result String (using strings.Builder):", resultString)

	// Using bytes.Buffer for efficient byte buffer building
	var byteBuffer bytes.Buffer

	byteBuffer.WriteString("This is ")
	byteBuffer.WriteString("a byte buffer ")
	byteBuffer.WriteString("example.")

	resultBytes := byteBuffer.Bytes()
	fmt.Println("Result Bytes (using bytes.Buffer):", resultBytes)

	// Casting strings to an array of runes
	runesArray := []rune(myString)
	fmt.Println("Array of Runes:", runesArray)

	// UTF-8 encoding uses a variable number of bytes for characters, making it efficient.
	// To work with characters directly, cast strings to an array of runes, which represents Unicode points.
	// String concatenation using the plus symbol creates new strings, but strings are immutable.
	// For efficient string building, use the strings package and the strings.Builder type.
	// strings.Builder appends values, and the final string is created only when the String() method is called.
}
