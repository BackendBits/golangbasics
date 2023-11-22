package commonpackage

// Constants for the project
const (
	Pi    = 3.14
	Point = 10 // Constant should start with Captial letter or else it will not work
)

// Variables for the project
var (
	MarkAge  = 25 // MarkAge is a public variable
	JohnAge  = 25 // JohnAge is a public variable
	kevinAge = 30 // kevinAge is a package-level variable
)

// Package level function
func printLocalVariablePackageFunction() int {
	var NoelAge = 21 // Local varibale within function
	return NoelAge
}

// public level function
func PrintLocalVariablePublicFunction() int {
	var NoelAge = 21
	return NoelAge
}
