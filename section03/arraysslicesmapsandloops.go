package main

import (
	"fmt"
	"time"
)

func main() {
	// SECTION 2: Arrays, Slices, Maps and Loopsin Golang
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 2: Arrays, Slices, Maps and Loopsin Golang")
	fmt.Println("======================================================================================================")
	fmt.Println()

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Arrays in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Different types of array initialization
	var arr1 [3]int           // All elements initialized to the zero value of int
	arr2 := [3]int{1, 2, 3}   // Initialization with values
	arr3 := [...]int{4, 5, 6} // Ellipsis (...) for inferring the length

	// Accessing elements in the array
	fmt.Println("Array1:", arr1)
	fmt.Println("Array2:", arr2)
	fmt.Println("Array3:", arr3)
	fmt.Println("Element at index 1 in Array2:", arr2[1])

	// Slices

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Slices in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Make keyword for creating slices
	var slice1 []int
	slice2 := make([]int, 0, 5) // Capacity is specified with make

	// Appending multiple values to the slice
	slice1 = append(slice1, 4, 5, 6)

	// Using make to create a slice
	slice3 := make([]int, 3, 5)

	// Capacity of the slices
	fmt.Println("Slice1:", slice1)
	fmt.Println("Slice2:", slice2)
	fmt.Println("Slice3:", slice3)
	fmt.Println("Capacity of Slice1:", cap(slice1))
	fmt.Println("Capacity of Slice2:", cap(slice2))
	fmt.Println("Capacity of Slice3:", cap(slice3))

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Creating Slices from Array in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Declare and initialize an array with five integers
	arr := [5]int{1, 2, 3, 4, 5}

	// Create a slice from the array that includes elements from index 1 to 3  and 4th (exclusive)
	// Syntax: arr[start:end]
	// - start: Index from which the slice begins (inclusive)
	// - end: Index where the slice ends (exclusive)
	slice := arr[1:4]

	// Display the original array
	fmt.Println("Array (arr):", arr)

	// Display the created slice
	fmt.Println("Slice (arr[1:4]):", slice)

	// Maps

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Maps in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Difference between array and slice
	// Arrays have fixed size, while slices are dynamic and can grow or shrink.

	// Difference between array and slice
	ages := make(map[string]uint8)
	ages["John"] = 30
	ages["Mark"] = 28

	// Iterating over the map
	fmt.Println("Map:", ages)
	for key, value := range ages {
		fmt.Println(key, value)
	}
	fmt.Println()
	// Accessing values in the map
	_, unknownExists := ages["Unknown"]

	// Deleting a value from the map
	delete(ages, "Mark")

	// Iterating over the map
	fmt.Println("Map:", ages)
	for key, value := range ages {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Error handling with maps
	if !unknownExists {
		fmt.Println("Age for Unknown not found.")
	}

	// Error handling and printing age for Mark
	markAge, markExists := ages["Mark"]
	if markExists {
		fmt.Printf("Age for Mark is %d\n", markAge)
	}

	// Loops

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Loops in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// While loop
	i := 0
	for i < 5 {
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// For loop with break
	for j := 0; ; j++ {
		if j >= 5 {
			break
		}
		fmt.Print(j, " ")
	}
	fmt.Println()

	// For loop with range
	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		fmt.Println("Index:", index, "Name:", name)
	}

	// Performance Test

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Performance Test with and without Preallocation on slice")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Why: To measure the time taken for append operations in a loop
	// How: Using time.Since to calculate elapsed time

	testSlice := make([]int, 0)
	timeWithoutPreallocation := timeLoop(testSlice, 1000000, false)
	timeWithPreallocation := timeLoop(testSlice, 1000000, true)

	fmt.Printf("Time without preallocation: %v seconds\n", timeWithoutPreallocation)
	fmt.Printf("Time with preallocation: %v seconds\n", timeWithPreallocation)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println()
}

// timeLoop measures the time taken to perform append operations in a loop.
// It takes a slice, number of iterations, and a boolean indicating whether to preallocate the slice.
// It returns the elapsed time in seconds.
func timeLoop(slice []int, iterations int, preallocate bool) float64 {
	start := time.Now()

	for i := 0; i < iterations; i++ {
		if preallocate {
			// Preallocate memory by doubling the capacity when nearing the limit
			if i >= cap(slice) {
				newSlice := make([]int, len(slice), 2*cap(slice))
				copy(newSlice, slice)
				slice = newSlice
			}
			slice = append(slice, i)
		} else {
			slice = append(slice, i)
		}
	}

	elapsed := time.Since(start).Seconds()
	return elapsed
}
