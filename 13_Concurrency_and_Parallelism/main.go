package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Use Case 1: Basic Goroutine
	fmt.Println("Use Case 1: Basic Goroutine")

	go func() {
		fmt.Println("Hello from Goroutine 1")
	}()

	// Wait for a moment to allow Goroutine to run
	time.Sleep(time.Millisecond * 100)

	// Use Case 2: Concurrency with Goroutines
	fmt.Println("\nUse Case 2: Concurrency with Goroutines")

	const numTasks = 5

	var wg sync.WaitGroup
	wg.Add(numTasks)

	for i := 1; i <= numTasks; i++ {
		go func(taskID int) {
			defer wg.Done()
			fmt.Printf("Task %d is running\n", taskID)
		}(i)
	}

	wg.Wait()

	// Use Case 3: Parallelism with Goroutines
	fmt.Println("\nUse Case 3: Parallelism with Goroutines")

	const numWorkers = 3

	var wg2 sync.WaitGroup
	taskCh := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		wg2.Add(1)
		go func(workerID int) {
			defer wg2.Done()
			for task := range taskCh {
				fmt.Printf("Worker %d is processing task %d\n", workerID, task)
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			}
		}(i)
	}

	// Send tasks to be processed in parallel
	for i := 1; i <= numTasks; i++ {
		taskCh <- i
	}

	close(taskCh)
	wg2.Wait()

	// Use Case 4: Channel Synchronization
	fmt.Println("\nUse Case 4: Channel Synchronization")

	var wg3 sync.WaitGroup
	doneCh := make(chan struct{})

	wg3.Add(1)
	go func() {
		defer wg3.Done()
		fmt.Println("Worker is doing some work...")
		time.Sleep(1 * time.Second)
		fmt.Println("Worker is done.")
		doneCh <- struct{}{}
	}()

	// Wait for the worker to finish
	<-doneCh
	wg3.Wait()

	// Use Case 5: Parallel Execution with WaitGroup
	fmt.Println("\nUse Case 5: Parallel Execution with WaitGroup")

	var parallelWg sync.WaitGroup

	parallelWg.Add(2)

	go func() {
		defer parallelWg.Done()
		fmt.Println("Task A is running")
		time.Sleep(1 * time.Second)
	}()

	go func() {
		defer parallelWg.Done()
		fmt.Println("Task B is running")
		time.Sleep(1 * time.Second)
	}()

	parallelWg.Wait()

	// Use Case 6: Parallel Execution with Goroutines
	fmt.Println("\nUse Case 6: Parallel Execution with Goroutines")

	done := make(chan bool)
	for i := 1; i <= 3; i++ {
		go func(taskID int) {
			fmt.Printf("Task %d is running\n", taskID)
			time.Sleep(1 * time.Second)
			done <- true
		}(i)
	}

	for i := 1; i <= 3; i++ {
		<-done
	}

	fmt.Println("\nAll tasks are completed.")
}
