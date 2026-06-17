package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 4: Strings, Runes and Bytes in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	city := "Kraków"
	fmt.Println("Original string:", city)
	fmt.Println("Rune count:", utf8.RuneCountInString(city))
	fmt.Println("Byte count:", len(city))

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Bytes vs Runes")
	fmt.Println("======================================================================================================")
	fmt.Println()

	text := "Hello, 世界"
	firstRune, size := utf8.DecodeRuneInString(text)
	fmt.Printf("First rune: %c, bytes used: %d\n", firstRune, size)

	// String indexes are byte indexes. text[:3] is safe here only because the
	// first three characters are ASCII and each ASCII character is one byte.
	fmt.Println("First 3 bytes:", text[:3])

	// To slice by characters, convert to []rune first.
	runes := []rune(text)
	fmt.Println("First 3 runes:", string(runes[:3]))
	fmt.Println("First 8 runes:", string(runes[:8]))

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Iterating over Unicode Text")
	fmt.Println("======================================================================================================")
	fmt.Println()

	myString := "Hello, 世界!"
	fmt.Println("String:", myString)
	fmt.Printf("Byte at index 2: %v\n", myString[2])

	for index, char := range myString {
		fmt.Printf("Byte index: %d, Rune: %c, Unicode Point: %U\n", index, char, char)
	}

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Efficient String and Byte Building")
	fmt.Println("======================================================================================================")
	fmt.Println()

	var strBuilder strings.Builder
	strBuilder.WriteString("This is ")
	strBuilder.WriteString("a concatenated ")
	strBuilder.WriteString("string.")
	fmt.Println("Result string using strings.Builder:", strBuilder.String())

	var byteBuffer bytes.Buffer
	byteBuffer.WriteString("This is ")
	byteBuffer.WriteString("a byte buffer ")
	byteBuffer.WriteString("example.")
	fmt.Println("Result bytes using bytes.Buffer:", byteBuffer.Bytes())

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Creating Runes")
	fmt.Println("======================================================================================================")
	fmt.Println()

	rune1 := 'A'
	rune2 := '世'
	fmt.Printf("Rune 1: %c, Unicode Point: %U\n", rune1, rune1)
	fmt.Printf("Rune 2: %c, Unicode Point: %U\n", rune2, rune2)

	fmt.Println()
	fmt.Println("Key point: len(string) counts bytes. range over a string decodes runes.")
	fmt.Println("Convert to []rune when you need character-based slicing.")
	fmt.Println("======================================================================================================")
	fmt.Println()
}
