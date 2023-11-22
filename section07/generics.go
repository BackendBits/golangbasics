package main

import (
	"fmt"
)

// Swap takes a slice of any type and swaps elements at indices i and j
func Swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	// SECTION 7: Generics in Golang
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 7: Generics in Golang")
	fmt.Println("======================================================================================================")
	fmt.Println()
	// Create a slice of integers
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", intSlice)

	// Swap elements at indices 1 and 3
	Swap(intSlice, 1, 3)
	fmt.Println("Slice after swapping:", intSlice)

	// Create a slice of strings
	strSlice := []string{"apple", "banana", "cherry"}
	fmt.Println("\nOriginal string slice:", strSlice)

	// Swap elements at indices 0 and 2
	Swap(strSlice, 0, 2)
	fmt.Println("String slice after swapping:", strSlice)
}
