package main

import (
	"fmt"
	"time"
)

// Assume functions sendMessage and sendEmail are implemented
func sendMessage(prefix, website string) {
	fmt.Println(prefix, website)
}

func sendEmail(prefix, website string) {
	fmt.Println(prefix, website)
}

func main() {
	// Creating a buffered channel for integers with a capacity of 5
	bufferedCh := make(chan int, 5)

	// Sending values to the buffered channel
	sendToChannel(bufferedCh)

	// Receiving values from the buffered channel
	receiveFromChannel(bufferedCh)

	// Closing the buffered channel
	close(bufferedCh)

	// Using select statement to handle multiple channels
	chickenChannel := make(chan string)
	tofuChannel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		chickenChannel <- "www.chicken.com"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		tofuChannel <- "www.tofu.com"
	}()

	// Use a goroutine for the select statement
	go func() {
		select {
		case chickenWebsite := <-chickenChannel:
			sendMessage("Chicken on sale at", chickenWebsite)
		case tofuWebsite := <-tofuChannel:
			sendEmail("Tofu on sale at", tofuWebsite)
		default:
			fmt.Println("No website information received.")
		}
	}()

	// Simulate all channels including the default case
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Simulating all channels:")
		select {
		case chickenWebsite := <-chickenChannel:
			sendMessage("Chicken on sale at", chickenWebsite)
		case tofuWebsite := <-tofuChannel:
			sendEmail("Tofu on sale at", tofuWebsite)
		default:
			fmt.Println("No website information received.")
		}
	}()

	// Sleep to allow goroutines to finish
	time.Sleep(5 * time.Second)
}

func sendToChannel(ch chan<- int) {
	// Sending values to the channel
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func receiveFromChannel(ch <-chan int) {
	// Receiving values from the channel
	for i := 0; i < 5; i++ {
		value := <-ch
		fmt.Println("Received value:", value)
	}
}
