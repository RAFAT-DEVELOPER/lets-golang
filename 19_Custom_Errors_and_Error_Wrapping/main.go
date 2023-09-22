package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// Custom error type
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Function that returns a custom error
func customErrorExample() error {
	return &MyError{Code: 100, Message: "Custom error occurred"}
}

// Function that wraps an error with additional context
func wrapErrorExample() error {
	originalErr := customErrorExample()
	return fmt.Errorf("Wrapped error: %w", originalErr)
}

func main() {
	// Case 1: Custom error example
	customErr := customErrorExample()
	fmt.Println("Custom Error:", customErr)

	// Case 2: Error wrapping example
	wrappedErr := wrapErrorExample()
	fmt.Println("Wrapped Error:", wrappedErr)

	// Case 3: Unwrap and inspect wrapped error
	unwrappedErr := errors.Unwrap(wrappedErr)
	if unwrappedErr != nil {
		fmt.Println("Unwrapped Error:", unwrappedErr)
	}

	// Case 4: Checking for specific error type
	if customErr, ok := unwrappedErr.(*MyError); ok {
		fmt.Printf("Custom Error Code: %d, Message: %s\n", customErr.Code, customErr.Message)
	}

	// Case 5: Writing wrapped error to a file
	file, err := os.Create("error_log.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, writeErr := io.WriteString(file, wrappedErr.Error())
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
	}

	fmt.Println("Error written to 'error_log.txt'")
}
