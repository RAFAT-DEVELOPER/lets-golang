package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	// 1. Deferred Function Calls
	defer fmt.Println("1. Deferred function 1")
	defer fmt.Println("1. Deferred function 2")

	// 2. Resource Cleanup using Defer
	file := createFile("example.txt")
	defer closeFile(file)

	// 3. Logging and Debugging
	logError()
	debug()

	// 4. Transaction Handling
	doDatabaseTransaction()

	// 5. Mutex Unlocking
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	// 6. Resource Allocation and Deallocation
	allocateAndDeallocateMemory()

	// 7. Recover from Panics
	performOperation()

	// 8. Error Handling and Propagation
	_, err := simulateError()
	if err != nil {
		fmt.Println("8. Error Handling:", err)
	}

	// 9. Graceful Program Termination
	defer cleanupBeforeExit()

	// 10. Stack Unwinding
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("10. Recovered from panic:", r)
		}
	}()

	// 11. Assertions and Validations
	validateInput(42)

	// 12. Graceful Failure Handling
	handleFailure()

	// 13. Graceful Degradation
	gracefulDegradation()

	// 14. Middleware and Frameworks
	handleRequest()

	fmt.Println("Main function exits.")
}

func createFile(filename string) *os.File {
	fmt.Println("2. Creating file...")
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(file *os.File) {
	fmt.Println("2. Closing file...")
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func logError() {
	fmt.Println("3. Simulating an error...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("3. Recovered from panic:", r)
		}
	}()
	panic("3. This is an error.")
}

func debug() {
	fmt.Println("3. Debugging...")
	// Your debugging logic here
}

func doDatabaseTransaction() {
	fmt.Println("4. Starting database transaction...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("4. Rolling back transaction due to panic:", r)
			// Rollback transaction logic
		} else {
			fmt.Println("4. Committing transaction...")
			// Commit transaction logic
		}
	}()
	// Simulate database transaction
	panic("4. Something went wrong!")
}

func allocateAndDeallocateMemory() {
	fmt.Println("6. Allocating and deallocating memory...")
	data := make([]int, 10)
	defer func() {
		fmt.Println("6. Deallocating memory...")
		data = nil
	}()
	// Use allocated memory
	data[5] = 42
	fmt.Println("6. Data[5]:", data[5])
}

func performOperation() {
	fmt.Println("7. Performing an operation...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("7. Recovered from panic:", r)
		}
	}()
	// Simulate an error
	var data []int
	data[0] = 42 // This will cause a panic
}

func simulateError() (int, error) {
	fmt.Println("8. Simulating an error...")
	return 0, fmt.Errorf("8. This is an error.")
}

func cleanupBeforeExit() {
	fmt.Println("9. Cleaning up before exit...")
	// Cleanup logic before program exit
}

func validateInput(input int) {
	fmt.Println("11. Validating input...")
	if input < 0 {
		panic("11. Input should be a positive number.")
	}
	fmt.Println("11. Input is valid.")
}

func handleFailure() {
	fmt.Println("12. Handling graceful failure...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("12. Recovered from panic:", r)
			// Handle the failure gracefully
		}
	}()
	// Simulate a failure
	panic("12. A failure occurred.")
}

func gracefulDegradation() {
	fmt.Println("13. Simulating graceful degradation...")
	// Implement graceful degradation logic
}

func middlewareFunction(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("14. Middleware: Preprocessing")
		next(w, r)
		fmt.Println("14. Middleware: Postprocessing")
	}
}

func handleRequest() {
	http.HandleFunc("/", middlewareFunction(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("14. Handling the request")
		// Actual request handling logic
	}))
	http.ListenAndServe(":8080", nil)
}
