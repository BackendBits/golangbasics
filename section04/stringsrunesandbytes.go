package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// SECTION 4: Strings, Runes and Bytes in Golang
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 4: Strings, Runes and Bytes in Golang")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// =====================================================================================================
	// Example 1: Counting Runes in a String
	// =====================================================================================================
	city := "Kraków"
	fmt.Println("Original String:", city)

	// utf8.RuneCountInString counts the number of runes (Unicode characters) in a string
	runeCount := utf8.RuneCountInString(city)
	fmt.Println("Number of Runes:", runeCount)

	// len() counts the number of bytes in a string, not runes
	byteCount := len(city)
	fmt.Println("Number of Bytes:", byteCount)

	fmt.Println("\n------------------------\n")

	// =====================================================================================================
	// Example 2: Working with Runes in a String
	// =====================================================================================================
	text := "Hello, 世界" // Hello, World in English and Sekai means world in Japanese

	// Extracting a specific rune
	firstRune, _ := utf8.DecodeRuneInString(text)
	fmt.Println("\nFirst Rune in the String:", firstRune)

	// Slicing a string by runes
	slicedString := text[:3] // Take the first 3 runes
	fmt.Println("\nSliced String by Runes:", slicedString)

	// =====================================================================================================
	// Example 3: Strings and UTF-8 Characters
	// =====================================================================================================
	myString := "Hello, 世界!" // Sekai means world
	fmt.Println("String:", myString)

	// Indexing string reveals UTF-8 byte values
	fmt.Printf("Index 2: %v\n", myString[2])

	// Iterating over string using range decodes multi-byte characters
	fmt.Println("Iterating over runes:")
	for index, char := range myString {
		// %U prints the Unicode format, %#U adds the "U+" prefix
		fmt.Printf("Index: %d, Rune: %c, Unicode Point: %U\n", index, char, char)
	}

	// =====================================================================================================
	// Example 4: Efficient String Building with strings.Builder
	// =====================================================================================================
	var strBuilder strings.Builder

	strBuilder.WriteString("This is ")
	strBuilder.WriteString("a concatenated ")
	strBuilder.WriteString("string.")

	resultString := strBuilder.String()
	fmt.Println("Result String (using strings.Builder):", resultString)

	// =====================================================================================================
	// Example 5: Efficient Byte Buffer Building with bytes.Buffer
	// =====================================================================================================
	var byteBuffer bytes.Buffer

	byteBuffer.WriteString("This is ")
	byteBuffer.WriteString("a byte buffer ")
	byteBuffer.WriteString("example.")

	resultBytes := byteBuffer.Bytes()
	fmt.Println("Result Bytes (using bytes.Buffer):", resultBytes)

	// =====================================================================================================
	// Example 6: Casting Strings to an Array of Runes
	// =====================================================================================================
	runesArray := []rune(myString)
	fmt.Println("Array of Runes:", runesArray)

	// =====================================================================================================
	// Example 7: Creating Runes
	// =====================================================================================================
	rune1 := 'A'
	fmt.Printf("\nRune 1: %c, Unicode Point: %U\n", rune1, rune1)

	rune2 := '世'
	fmt.Printf("Rune 2: %c, Unicode Point: %U\n", rune2, rune2)

	// =====================================================================================================
	// Additional Comments Section
	// =====================================================================================================
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("UTF-8 encoding uses a variable number of bytes for characters, making it efficient.")
	fmt.Println("To work with characters directly, cast strings to an array of runes, which represents Unicode points.")
	fmt.Println("String concatenation using the plus symbol creates new strings, but strings are immutable.")
	fmt.Println("For efficient string building, use the strings package and the strings.Builder type.")
	fmt.Println("strings.Builder appends values, and the final string is created only when the String() method is called.")
	fmt.Println("======================================================================================================")
	fmt.Println()
}
