package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Use Case 1: Creating and running Goroutines
	fmt.Println("Use Case 1: Creating and Running Goroutines")

	// Create a WaitGroup to wait for Goroutines to finish
	var wg sync.WaitGroup

	// Launch multiple Goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is running\n", id)
		}(i)
	}

	// Wait for all Goroutines to finish
	wg.Wait()

	// Use Case 2: Using Channels for Communication
	fmt.Println("\nUse Case 2: Using Channels for Communication")

	// Create a channel to pass integers between Goroutines
	ch := make(chan int)

	// Launch a Goroutine to send data to the channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel when done sending
	}()

	// Launch a Goroutine to receive data from the channel
	go func() {
		for num := range ch {
			fmt.Printf("Received: %d\n", num)
		}
	}()

	// Use Case 3: Buffered Channels
	fmt.Println("\nUse Case 3: Buffered Channels")

	// Create a buffered channel with a capacity of 3
	bufCh := make(chan string, 3)

	// Launch Goroutines to send data to the buffered channel
	go func() {
		bufCh <- "A"
		bufCh <- "B"
		bufCh <- "C"
		close(bufCh) // Close the channel when done sending
	}()

	// Launch a Goroutine to receive data from the buffered channel
	go func() {
		for item := range bufCh {
			fmt.Printf("Received: %s\n", item)
		}
	}()

	// Use Case 4: Select Statement for Channel Operations
	fmt.Println("\nUse Case 4: Select Statement for Channel Operations")

	// Create two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Launch Goroutines to send data to channels
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch1 <- "Message from Channel 1"
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		ch2 <- "Message from Channel 2"
	}()

	// Use a select statement to receive from the first available channel
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}

	// Use Case 5: Error Handling with Channels
	fmt.Println("\nUse Case 5: Error Handling with Channels")

	// Create a channel to simulate an error
	errCh := make(chan error)

	// Launch a Goroutine to send an error to the channel
	go func() {
		time.Sleep(time.Second)
		errCh <- fmt.Errorf("Simulated error")
	}()

	// Use a select statement to handle errors
	select {
	case err := <-errCh:
		fmt.Printf("Error received: %v\n", err)
	case <-time.After(2 * time.Second):
		fmt.Println("No error received within the timeout")
	}
}
