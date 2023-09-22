package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// 1. Cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go performTask(ctx)

	// Simulate cancellation after 2 seconds
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling the task...")
	cancel()

	// 2. Timeout
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelTimeout()

	go performTaskWithTimeout(ctxTimeout)

	// 3. Deadlines
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancelDeadline()

	go performTaskWithDeadline(ctxDeadline)

	// 4. Propagation
	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	go parentGoroutine(parentCtx)

	// 5. Context Values
	valueCtx := context.WithValue(context.Background(), "user", "john.doe")
	go printUser(valueCtx)

	// 6. Context Hierarchies
	rootCtx := context.Background()
	childCtx1, _ := context.WithTimeout(rootCtx, 2*time.Second)
	childCtx2, _ := context.WithTimeout(rootCtx, 3*time.Second)

	go processWithHierarchy(rootCtx, "Root Context")
	go processWithHierarchy(childCtx1, "Child Context 1")
	go processWithHierarchy(childCtx2, "Child Context 2")

	// 7. Selective Cancellation
	parentContext, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()

	go selectiveCancellation(parentContext)

	// 8. Graceful Shutdown
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt)

	// Handle graceful shutdown
	go func() {
		<-shutdownSignal
		fmt.Println("\nReceived interrupt signal. Shutting down gracefully...")
		cancelParent()
	}()

	// 9. Resource Cleanup
	resourceCtx, cancelResource := context.WithCancel(context.Background())
	defer cancelResource()

	go resourceCleanup(resourceCtx)

	// Wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(9)
	wg.Wait()
	fmt.Println("All goroutines have finished.")
}

func performTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task has been cancelled.")
			return
		default:
			fmt.Println("Performing a task...")
			time.Sleep(1 * time.Second)
		}
	}
}

func performTaskWithTimeout(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task with timeout has been cancelled.")
	case <-time.After(4 * time.Second):
		fmt.Println("Task with timeout has completed.")
	}
}

func performTaskWithDeadline(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Task with deadline has been cancelled.")
	case <-time.After(6 * time.Second):
		fmt.Println("Task with deadline has completed.")
	}
}

func parentGoroutine(ctx context.Context) {
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go childGoroutine(childCtx)
}

func childGoroutine(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Child goroutine received cancellation signal.")
	}
}

func printUser(ctx context.Context) {
	user := ctx.Value("user")
	if user != nil {
		fmt.Println("User:", user)
	}
}

func processWithHierarchy(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		fmt.Printf("%s has been cancelled.\n", name)
	}
}

func selectiveCancellation(parentContext context.Context) {
	childContext1, _ := context.WithCancel(parentContext)
	childContext2, _ := context.WithCancel(parentContext)

	go func() {
		select {
		case <-childContext1.Done():
			fmt.Println("Child 1 has been cancelled.")
		}
	}()

	go func() {
		select {
		case <-childContext2.Done():
			fmt.Println("Child 2 has been cancelled.")
		}
	}()

	// Simulate cancelling only one child
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling Child 1...")
	childContext1.Done()
}

func resourceCleanup(ctx context.Context) {
	// Simulate resource cleanup (closing a file, database connection, etc.)
	<-ctx.Done()
	fmt.Println("Resource cleanup completed.")
}
