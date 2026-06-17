# Go Learning Roadmap

This repository currently covers Part 1: Go language fundamentals.

## Part 1: Core syntax and language basics

Covered in this repo:

- variables, constants, and data types
- functions and control flow
- arrays, slices, maps, and loops
- strings, runes, and bytes
- structs and interfaces
- pointers
- generics
- file read/write

Expected result: you should be comfortable reading and writing small Go programs without copying code blindly.

## Part 2: Standard library foundations

Recommended next topics:

- errors and wrapping with `fmt.Errorf`
- time and duration
- sorting
- regex
- JSON and CSV
- command-line flags
- environment variables
- logging

Expected result: you can build useful CLI tools using only the standard library.

## Part 3: Testing and code quality

Recommended next topics:

- unit tests with `testing`
- table-driven tests
- benchmarks
- examples
- `go test`, `go test -race`, and coverage
- package organization

Expected result: you can write code that is easy to verify and maintain.

## Part 4: Concurrency

Recommended next topics:

- goroutines
- channels
- buffered channels
- select
- wait groups
- mutexes
- context cancellation
- worker pools

Expected result: you understand Go's concurrency model beyond simple examples.

## Part 5: Backend development with Go

Recommended next topics:

- HTTP servers
- middleware
- REST APIs
- database access
- transactions
- configuration
- graceful shutdown
- Docker

Expected result: you can build a small production-style Go backend service.
