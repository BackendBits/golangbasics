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
	// Creating items
	item1 := Item{Name: "Laptop", Price: 1000.0}
	item2 := Item{Name: "Mouse", Price: 20.0}

	// Displaying initial inventory
	fmt.Println("Initial Inventory:")
	fmt.Println("Item 1:", item1)
	fmt.Println("Item 2:", item2)

	// Updating the price of an item using a pointer
	UpdatePrice(&item1, 1200.0)

	// Displaying updated inventory
	fmt.Println("\nUpdated Inventory:")
	fmt.Println("Item 1:", item1) // Price of item1 is updated

	// Creating a pointer to an item
	itemPointer := &item2

	// Updating the price of an item using the pointer
	UpdatePrice(itemPointer, 25.0)
	fmt.Println("Item 2 (after regular update):", item2) // Price of item2 is updated

	// Updating the price of an item using unsafe.Pointer
	UnsafeUpdatePrice(itemPointer, 30.0)
	fmt.Println("Item 2 (after unsafe update):", item2) // Price of item2 is updated using unsafe.Pointer
}
