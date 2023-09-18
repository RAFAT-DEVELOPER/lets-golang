package main

import (
	"errors"
	"fmt"
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

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	// Case 1: Defer for cleanup or finalization
	file, err := os.Open("nonexistent.txt")
	defer func() {
		if file != nil {
			file.Close()
		}
	}()
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	// Case 2: Error checking with type assertions
	var val interface{} = "Hello, Go"
	str, ok := val.(string)
	if !ok {
		fmt.Println("Value is not a string")
	} else {
		fmt.Println("Value:", str)
	}

	// Case 3: Panicking and recovering from errors
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("This is a panic!")

	// Case 4: Basic error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Case 5: Custom error handling
	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Custom Error:", err)
	}

	// Case 6: Using custom error types
	customErr := &MyError{Code: 100, Message: "Custom error message"}
	fmt.Println("Custom Error:", customErr)

	// Case 7: Error wrapping with additional context
	wrappedErr := fmt.Errorf("Wrapping error: %w", customErr)
	fmt.Println("Wrapped Error:", wrappedErr)

	// Case 8: Handling multiple errors with the `errors` package
	err1 := errors.New("Error 1")
	err2 := errors.New("Error 2")
	err3 := errors.New("Error 3")

	multiErr := errors.New("Multiple errors")
	multiErr = fmt.Errorf("%w\n\t%v\n\t%v\n\t%v", multiErr, err1, err2, err3)
	fmt.Println("Multiple Errors:", multiErr)
}
