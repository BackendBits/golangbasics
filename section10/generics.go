package main

import (
	"fmt"
	"sort"
)

// Declare a generic type parameter `T`
type Container[T any] struct {
	data []T
}

// Define a generic function `Swap` that swaps two elements in a container
func Swap[T any](c *Container[T], i, j int) {
	c.data[i], c.data[j] = c.data[j], c.data[i]
}

// Define a generic function `Sort` that sorts a container using a comparison function
func Sort[T any](c *Container[T], compare func(a, b T) bool) *Container[T] {
	for i := 0; i < len(c.data)-1; i++ {
		for j := i + 1; j < len(c.data); j++ {
			if compare(c.data[j], c.data[i]) {
				Swap(c, i, j)
			}
		}
	}
	return c
}

func main() {
	// Create a container of strings
	cStrings := Container[string]{}
	cStrings.data = []string{"Alice", "Bob", "Charlie", "David", "Emily"}

	// Sort the container of strings in alphabetical order
	sort.Strings(cStrings.data)

	fmt.Println("Sorted strings:", cStrings.data)
}
