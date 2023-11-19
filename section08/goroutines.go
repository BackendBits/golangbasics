package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to simulate a time-consuming task
func timeConsumingTask(taskID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d started\n", taskID)
	time.Sleep(2 * time.Second) // Simulate a time-consuming task
	fmt.Printf("Task %d completed\n", taskID)
}

func main() {
	// Example 1: Basic Goroutine
	go func() {
		fmt.Println("Hello from Goroutine!")
		// Tricky Aspect: This Goroutine may not have time to execute before the program exits.
	}()

	// Example 2: Multiple Goroutines with WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go timeConsumingTask(i, &wg)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Tricky Aspect: Without WaitGroup, the main Goroutine may not wait for others to finish, and the program might exit prematurely.

	fmt.Println("Main function completed")
}
