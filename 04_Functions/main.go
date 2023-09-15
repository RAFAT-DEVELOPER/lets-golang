package main

import (
	"fmt"
)

// Function with parameters and a return value
func add(x, y int) int {
	return x + y
}

// Function with multiple return values
func divideAndRemainder(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// Variadic function (accepts a variable number of arguments)
func sum(numbers ...int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}
	return result
}

// Function with a named return value (naming the return variable)
func multiply(x, y int) (product int) {
	product = x * y
	return // product is automatically returned
}

// Function as a variable (function closure)
func createIncrementer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Anonymous function
var square = func(x int) int {
	return x * x
}

// Function with defer statement
func deferExample() {
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")
	fmt.Println("Function body")
}

// Function with a pointer receiver (method)
type Circle struct {
	Radius float64
}

func (c *Circle) calculateArea() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	// Function with parameters and return value
	result := add(5, 3)
	fmt.Println("Result of add:", result)

	// Function with multiple return values
	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// Variadic function
	sumResult := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sumResult)

	// Named return value
	product := multiply(6, 7)
	fmt.Println("Product:", product)

	// Function as a variable
	increment := createIncrementer()
	fmt.Println("Increment:", increment())
	fmt.Println("Increment:", increment())

	// Anonymous function
	squareResult := square(4)
	fmt.Println("Square:", squareResult)

	// Defer statement
	deferExample()

	// Method (function with a receiver)
	circle := Circle{Radius: 2.5}
	area := circle.calculateArea()
	fmt.Println("Circle Area:", area)
}
