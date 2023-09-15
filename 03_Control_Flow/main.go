package main

import (
	"fmt"
)

func main() {
	// If statement
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	// For loop
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While-like loop
	fmt.Println("Counting from 5 to 1:")
	j := 5
	for j >= 1 {
		fmt.Println(j)
		j--
	}

	// Switch statement
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("It's an apple")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("It's something else")
	}

	// Break and continue
	fmt.Println("Counting from 1 to 10 but skipping 5:")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue // Skip iteration 5
		}
		fmt.Println(i)
	}

	// Nested loops
	fmt.Println("Nested loop example:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("(%d, %d) ", i, j)
		}
		fmt.Println()
	}

	// Defer statement
	deferExample()

	// Panic and recover
	panicAndRecover()
}

func deferExample() {
	defer fmt.Println("Deferred statement 1")
	defer fmt.Println("Deferred statement 2")
	fmt.Println("Function body")
}

func panicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Before panic")
	panic("Panic!")
	fmt.Println("After panic (this won't be executed)")
}
