package main

import (
	"errors"
	"fmt"
	"os"
)

// CustomError represents a custom error type.
type CustomError struct {
	Msg string
}

// Implement the error interface for CustomError.
func (ce CustomError) Error() string {
	return ce.Msg
}

// divide divides two numbers and returns the result. It handles division by zero error.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// openFile opens a file and returns an error if the file does not exist.
func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	// Use Case 1: Creating custom errors
	customErr := CustomError{Msg: "This is a custom error"}
	fmt.Println("Use Case 1: Creating Custom Errors")
	fmt.Println(customErr)

	// Use Case 2: Handling division by zero error
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("\nUse Case 2: Handling Division by Zero Error")
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	// Use Case 3: Handling file not found error
	filename := "nonexistent.txt"
	file, err := openFile(filename)
	if err != nil {
		fmt.Println("\nUse Case 3: Handling File Not Found Error")
		fmt.Printf("Error: %v\n", err)
	} else {
		defer file.Close()
		fmt.Printf("File %s opened successfully.\n", filename)
	}

	// Use Case 4: Error propagation
	result, err = divide(20, 5)
	if err != nil {
		fmt.Println("\nUse Case 4: Error Propagation")
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("\nResult: %d\n", result)
	}
}
