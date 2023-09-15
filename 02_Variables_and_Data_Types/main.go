package main

import (
	"fmt"
	"math"
)

func main() {
	// Integer types
	var integerVar int
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807
	integerVar = 42

	// Floating-point types
	var float32Var float32 = 3.14
	var float64Var float64 = 3.14159265359

	// Boolean type
	var boolVar bool = true

	// String type
	var stringVar string = "Hello, Go!"

	// Custom type
	type MyType int
	var customVar MyType = 100

	// Constants
	const pi = 3.14159

	// Arrays and Slices
	var intArray [3]int
	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3

	// Slices
	var intSlice []int
	intSlice = append(intSlice, 4, 5, 6)

	// Maps
	var myMap map[string]int
	myMap = make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	// Structs
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}

	// Pointers
	var intPointer *int
	intPointer = &integerVar

	// Functions
	add := func(x, y int) int {
		return x + y
	}

	// Math operations
	sqrt := math.Sqrt(16)
	fmt.Println("Square root of 16:", sqrt)

	// Type conversion
	var a int = 42
	var b float64 = float64(a)

	// Type assertion
	var i interface{} = 42
	j, ok := i.(int)

	// Nil values
	var nilIntVar int
	var nilSliceVar []int
	var nilMapVar map[string]int
	var nilPtrVar *int

	// Print variable values
	fmt.Println("Integer Variable:", integerVar)
	fmt.Println("Integer 8 Variable:", int8Var)
	fmt.Println("Integer 16 Variable:", int16Var)
	fmt.Println("Integer 32 Variable:", int32Var)
	fmt.Println("Integer 64 Variable:", int64Var)
	fmt.Println("Float 32 Variable:", float32Var)
	fmt.Println("Float 64 Variable:", float64Var)

	fmt.Println("Boolean Variable:", boolVar)
	fmt.Println("String Variable:", stringVar)
	fmt.Println("Custom Type Variable:", customVar)
	fmt.Println("Pi Constant:", pi)
	fmt.Println("Integer Array:", intArray)
	fmt.Println("Integer Slice:", intSlice)
	fmt.Println("Map:", myMap)
	fmt.Println("Struct:", person)
	fmt.Println("Integer Pointer:", *intPointer)
	fmt.Println("Add Function Result:", add(5, 3))
	fmt.Println("Square Root:", sqrt)
	fmt.Println("Type Conversion (int to float):", b)
	fmt.Println("Type Assertion (interface{} to int):", j, ok)
	fmt.Println("Nil Integer Variable:", nilIntVar)
	fmt.Println("Nil Slice Variable:", nilSliceVar)
	fmt.Println("Nil Map Variable:", nilMapVar)
	fmt.Println("Nil Pointer Variable:", nilPtrVar)
}
