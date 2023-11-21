package main

import (
	"fmt"
	"unsafe"
)

// Item represents an inventory item
type Item struct {
	Name  string
	Price float64
}

// UpdatePrice updates the price of an item using pointers
func UpdatePrice(item *Item, newPrice float64) {
	item.Price = newPrice
}

// In Go, the unsafe package provides a way to work with pointers and perform operations that are not type-safe or violate the language's safety guarantees.
// UnsafeUpdatePrice updates the price of an item using unsafe.Pointer
func UnsafeUpdatePrice(item *Item, newPrice float64) {
	// Convert the Item pointer to an unsafe.Pointer
	itemPointer := unsafe.Pointer(item)

	// Calculate the offset of the 'Price' field within the Item struct using unsafe.Offsetof
	// Add this offset to the itemPointer to get the pointer to the 'Price' field
	pricePointer := (*float64)(unsafe.Pointer(uintptr(itemPointer) + unsafe.Offsetof(item.Price)))

	// Dereference the pricePointer and update the price
	*pricePointer = newPrice
}

func main() {

	// SECTION 6: Pointers in Golang
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 6: Pointers in Golang")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// SECTION 6.1: Initial Inventory
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Initial Inventory")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Creating items
	item1 := Item{Name: "Laptop", Price: 1000.0}
	item2 := Item{Name: "Mouse", Price: 20.0}

	// Displaying initial inventory
	fmt.Println("Item 1:", item1)
	fmt.Println("Item 2:", item2)

	// SECTION 6.2: UpdatePrice with Pointer
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("UpdatePrice with Pointer")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Updating the price of an item using a pointer
	UpdatePrice(&item1, 1200.0)

	// Displaying updated inventory
	fmt.Println("Item 1 (after update):", item1) // Price of item1 is updated

	// SECTION 6.3: UpdatePrice with Item Pointer
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("UpdatePrice with Item Pointer")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Creating a pointer to an item
	itemPointer := &item2

	// Updating the price of an item using the pointer
	UpdatePrice(itemPointer, 25.0)
	fmt.Println("Item 2 (after regular update):", item2) // Price of item2 is updated

	// SECTION 6.4: UnsafeUpdatePrice
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("UnsafeUpdatePrice")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Updating the price of an item using unsafe.Pointer
	UnsafeUpdatePrice(itemPointer, 30.0)
	fmt.Println("Item 2 (after unsafe update):", item2) // Price of item2 is updated using unsafe.Pointer

	// SECTION 6.5: New Keyword
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("New Keyword")
	fmt.Println("======================================================================================================")
	fmt.Println()

	num := 42
	fmt.Printf("Address: %p, Value: %v\n", &num, num)

	var num1Pointer *int
	num1Pointer = new(int) // Allocate memory for an int and assign its address to num1Pointer

	// Setting a value through the pointer
	*num1Pointer = 42
	fmt.Printf("Pointer Address: %p, Pointer Value: %p, Pointed Value: %v\n", &num1Pointer, num1Pointer, *num1Pointer)
	fmt.Printf("Pointer Address: %p, Pointer Value: %p, Pointed Value: %v\n", &num1Pointer, num1Pointer, *num1Pointer)

}
