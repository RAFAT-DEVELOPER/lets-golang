package main

import (
	"fmt"
	"myprogram/helpers" // Import the "helpers" package
)

func main() {
	helpers.Greet("Alice") // Call a function from the "helpers" package
	result := helpers.Add(3, 4)
	fmt.Printf("The result of 3 + 4 is %d\n", result)
}
