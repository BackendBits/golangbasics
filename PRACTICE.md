# Go Basics Practice

Use these exercises after reading each section. The goal is not only to run the examples, but to change them and explain the result.

## Section 01: Constants, variables, and data types

1. Create variables using explicit type, inferred type, and short declaration.
2. Create one exported variable and one unexported variable in a package. Explain which one can be accessed from another package.
3. Print the type of five values using `fmt.Printf("%T", value)`.
4. Try assigning an `int` to a `string`. Explain why it fails.
5. Create constants for tax rate, app name, and max retry count.

## Section 02: Functions and control structures

1. Write a function that returns both quotient and remainder.
2. Write a function that returns an error when the input is invalid.
3. Use `if`, `else if`, and `else` to classify a score.
4. Use `switch` to map status codes to messages.
5. Write a variadic function that sums integers.

## Section 03: Arrays, slices, maps, and loops

1. Create an array of five numbers and print each index and value.
2. Convert an array range into a slice.
3. Append values to a slice and print `len` and `cap` after each append.
4. Create a map of employee names to ages.
5. Delete a map key and verify it using the `value, ok := map[key]` pattern.
6. Compare append performance with and without preallocation.

## Section 04: Strings, runes, and bytes

1. Print `len()` and `utf8.RuneCountInString()` for an English word and a Unicode word.
2. Iterate over a string using index access and then using `range`. Explain the difference.
3. Safely slice the first three characters from a Unicode string.
4. Build a long string using `strings.Builder`.
5. Convert a string to `[]byte` and `[]rune`. Explain when each is useful.

## Section 05: Structs and interfaces

1. Create a `User` struct with name, email, and age.
2. Add a method called `IsAdult()` to the `User` struct.
3. Create an interface with one method and implement it using two structs.
4. Pass different structs into a function that accepts the interface.
5. Explain why Go does not need an `implements` keyword.

## Section 06: Pointers

1. Write a function that updates a struct field using a pointer.
2. Write the same function without a pointer and explain why the original value does not change.
3. Print the address and value of a variable.
4. Create a value using `new` and update it.
5. Explain why `unsafe.Pointer` should be avoided unless there is a specific low-level reason.

## Section 07: Generics

1. Write a generic `Swap` function for a slice.
2. Write a generic `Contains` function for comparable values.
3. Write a generic `PrintAll` function for any type.
4. Explain the difference between `any` and `comparable`.
5. Refactor duplicate int/string logic using generics.

## Section 08: File read/write

1. Add a todo with a title containing spaces.
2. View all todos from the JSON lines file.
3. Edit an existing todo.
4. Delete a todo and then add a new one. Confirm that the ID is not duplicated.
5. Change the file path and verify that `filepath.Join` works across operating systems.

## Mini project after Part 1

Build a small CLI contact book with these commands:

- add contact
- list contacts
- search contact by name
- update contact email
- delete contact

Store contacts in a JSON lines file. Use structs, slices, maps, errors, file handling, and functions.
