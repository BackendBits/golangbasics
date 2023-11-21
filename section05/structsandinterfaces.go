package main

import "fmt"

// GasEngine struct definition
type GasEngine struct {
	MilesPerGallon uint8
	Gallons        uint8
	OwnerInfo      struct {
		Name string
	}
}

// Method for GasEngine struct
func (g GasEngine) MilesLeft() uint8 {
	return g.MilesPerGallon * g.Gallons
}

// Engine interface definition
type Engine interface {
	MilesLeft() uint8
}

// ElectricEngine struct implementing the Engine interface
type ElectricEngine struct {
	MilesPerKWH   uint8
	KilowattHours uint8
	OwnerInfo     struct {
		Name string
	}
}

// MilesLeft method for ElectricEngine
func (e ElectricEngine) MilesLeft() uint8 {
	return e.MilesPerKWH * e.KilowattHours
}

// canMakeIt function using the Engine interface
func canMakeIt(e Engine, miles uint8) bool {
	return e.MilesLeft() >= miles
}

func main() {

	// SECTION 5: Structs and Interfaces in Golang
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("SECTION 5: Structs and Interfaces in Golang")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// SECTION 5.1: GasEngine Example
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("GasEngine Example")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Creating a GasEngine instance
	gasCar := GasEngine{
		MilesPerGallon: 25,
		Gallons:        12,
		OwnerInfo: struct {
			Name string
		}{
			Name: "John Doe",
		},
	}

	// Accessing struct fields
	fmt.Println("Gas Car Miles per Gallon:", gasCar.MilesPerGallon)
	fmt.Println("Gas Car Kilowatt Hours:", gasCar.Gallons)
	fmt.Println("Gas Car Owner Name:", gasCar.OwnerInfo.Name)

	// Using a struct method
	fmt.Println("Gas Car Miles Left:", gasCar.MilesLeft())

	// SECTION 5.2: ElectricEngine Example
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("ElectricEngine Example")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Creating an ElectricEngine instance
	electricCar := ElectricEngine{
		MilesPerKWH:   4,
		KilowattHours: 20,
		OwnerInfo: struct {
			Name string
		}{
			Name: "Mark Doe",
		},
	}

	// Accessing ElectricEngine struct fields
	fmt.Println("Electric Car Miles per KWH:", electricCar.MilesPerKWH)
	fmt.Println("Electric Car Kilowatt Hours:", electricCar.KilowattHours)
	fmt.Println("Electric Car Owner Name:", electricCar.OwnerInfo.Name)

	// Using an interface method
	fmt.Println("Electric Car Miles Left:", electricCar.MilesLeft())

	// SECTION 5.3: Interface Check Example
	fmt.Println()
	fmt.Println("======================================================================================================")
	fmt.Println("Interface Check Example")
	fmt.Println("======================================================================================================")
	fmt.Println()

	// Check if each car can make a certain distance
	fmt.Println("Gas car can make it:", canMakeIt(gasCar, 200))
	fmt.Println("Electric car can make it:", canMakeIt(electricCar, 50))
}
