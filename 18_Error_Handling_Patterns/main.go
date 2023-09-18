package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	// Simulate an error
	return errors.New("something went wrong")
}

func main() {
	if err := doSomething(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("No errors")
}
