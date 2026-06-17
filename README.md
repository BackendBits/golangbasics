# Go Basics

A beginner-friendly Go learning repository with small, runnable examples for the core language features.

This repo is designed for three outcomes:

1. learn Go syntax by running code,
2. revise Go fundamentals quickly before interviews, and
3. build confidence by doing small hands-on exercises after every topic.

## Prerequisites

- Go 1.21 or later
- Git
- Any editor such as VS Code, GoLand, or Vim

Check your Go installation:

```bash
go version
```

## How to run

Clone the repository:

```bash
git clone https://github.com/BackendBits/golangbasics.git
cd golangbasics
```

Run any section directly:

```bash
go run ./section01
go run ./section02
go run ./section03
go run ./section04
go run ./section05
go run ./section06
go run ./section07
go run ./section08
```

Run formatting and compile checks:

```bash
gofmt -w .
go test ./...
```

## Learning path

| Order | Topic | File | What you should understand |
|---:|---|---|---|
| 1 | Constants, variables, and basic data types | [`section01/constantandvariableanddatatypes.go`](section01/constantandvariableanddatatypes.go) | Static typing, strong typing, exported names, constants, integer/float/string/bool types |
| 2 | Functions and control structures | [`section02/functionsandcontrolstructures.go`](section02/functionsandcontrolstructures.go) | Functions, variadic parameters, named returns, errors, if/else, switch |
| 3 | Arrays, slices, maps, and loops | [`section03/arraysslicesmapsandloops.go`](section03/arraysslicesmapsandloops.go) | Fixed arrays, dynamic slices, map lookup/delete, range loops, preallocation |
| 4 | Strings, runes, and bytes | [`section04/stringsrunesandbytes.go`](section04/stringsrunesandbytes.go) | UTF-8, byte indexing, rune iteration, safe character slicing, string builders |
| 5 | Structs and interfaces | [`section05/structsandinterfaces.go`](section05/structsandinterfaces.go) | Struct fields, methods, implicit interface implementation, polymorphism |
| 6 | Pointers | [`section06/pointers.go`](section06/pointers.go) | Pointer receivers, mutation through pointers, `new`, and why `unsafe` should be avoided in normal code |
| 7 | Generics | [`section07/generics.go`](section07/generics.go) | Type parameters and reusable generic functions |
| 8 | File read/write | [`section08/filesreadwrite.go`](section08/filesreadwrite.go) | JSON encoding/decoding, file append/rewrite, CLI input, CRUD-style flow |

## Recommended study method

For each section:

1. Read the file once without running it.
2. Predict the output.
3. Run the file using `go run ./sectionXX`.
4. Change one example and rerun it.
5. Write a 5-line summary in your own words.
6. Solve the related task from [`PRACTICE.md`](PRACTICE.md).

## Interview revision checklist

Before saying you know Go basics, make sure you can answer these:

- Why is Go called statically typed and strongly typed?
- What is the difference between `var`, `:=`, and `const`?
- Why are exported names capitalized in Go?
- What is the difference between an array and a slice?
- What are `len` and `cap` for a slice?
- Why should large append loops often preallocate slice capacity?
- What happens when reading a missing key from a map?
- Why does `len("Kraków")` return bytes, not characters?
- What is a rune?
- How does Go implement interfaces implicitly?
- When should you pass a pointer to a function?
- Why is `unsafe.Pointer` rarely used in application code?
- What problem do generics solve?
- How do you return and handle errors in idiomatic Go?
- How do you read and write JSON data using the standard library?

## Repo quality checks

This repository includes a GitHub Actions workflow that runs:

```bash
gofmt -w .
go test ./...
```

Use this before pushing changes locally:

```bash
gofmt -w . && go test ./...
```

## Next improvements planned

- Add unit tests for reusable functions.
- Add section-level notes with diagrams.
- Add beginner exercises with expected outputs.
- Add intermediate examples for goroutines, channels, context, HTTP servers, and testing.
