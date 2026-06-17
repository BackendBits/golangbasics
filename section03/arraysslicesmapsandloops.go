package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 3: Arrays, Slices, Maps and Loops in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	fmt.Println("======================================================================================================")
	fmt.Println("Arrays in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	var arr1 [3]int           // Fixed-size array. Values default to zero.
	arr2 := [3]int{1, 2, 3}   // Explicit length.
	arr3 := [...]int{4, 5, 6} // Compiler infers the length.

	fmt.Println("Array1:", arr1)
	fmt.Println("Array2:", arr2)
	fmt.Println("Array3:", arr3)
	fmt.Println("Element at index 1 in Array2:", arr2[1])

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Slices in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	var slice1 []int
	slice2 := make([]int, 0, 5) // len=0, cap=5
	slice3 := make([]int, 3, 5) // len=3, cap=5, values default to zero

	slice1 = append(slice1, 4, 5, 6)

	fmt.Printf("Slice1: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("Slice2: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("Slice3: %v, len=%d, cap=%d\n", slice3, len(slice3), cap(slice3))

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Creating Slices from Arrays")
	fmt.Println("======================================================================================================")
	fmt.Println()

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4] // start is inclusive, end is exclusive.

	fmt.Println("Array:", arr)
	fmt.Println("Slice arr[1:4]:", slice)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Maps in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	ages := map[string]uint8{
		"John": 30,
		"Mark": 28,
	}

	fmt.Println("Map:", ages)
	for key, value := range ages {
		fmt.Println(key, value)
	}

	unknownAge, unknownExists := ages["Unknown"]
	fmt.Printf("Unknown age: %d, exists: %t\n", unknownAge, unknownExists)

	delete(ages, "Mark")
	fmt.Println("Map after deleting Mark:", ages)

	if markAge, markExists := ages["Mark"]; markExists {
		fmt.Printf("Age for Mark is %d\n", markAge)
	} else {
		fmt.Println("Age for Mark not found after delete.")
	}

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Loops in Go")
	fmt.Println("======================================================================================================")
	fmt.Println()

	i := 0
	for i < 5 { // Go does not have a separate while keyword.
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	for j := 0; ; j++ {
		if j >= 5 {
			break
		}
		fmt.Print(j, " ")
	}
	fmt.Println()

	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		fmt.Println("Index:", index, "Name:", name)
	}

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Performance Test: Append With and Without Preallocation")
	fmt.Println("======================================================================================================")
	fmt.Println()

	iterations := 1_000_000
	timeWithoutPreallocation := timeLoop(iterations, false)
	timeWithPreallocation := timeLoop(iterations, true)

	fmt.Printf("Time without preallocation: %.6f seconds\n", timeWithoutPreallocation)
	fmt.Printf("Time with preallocation: %.6f seconds\n", timeWithPreallocation)

	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println()
}

func timeLoop(iterations int, preallocate bool) float64 {
	start := time.Now()

	var values []int
	if preallocate {
		values = make([]int, 0, iterations)
	} else {
		values = make([]int, 0)
	}

	for i := 0; i < iterations; i++ {
		values = append(values, i)
	}

	return time.Since(start).Seconds()
}
