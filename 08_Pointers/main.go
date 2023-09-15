package main

import (
	"fmt"
)

func main() {
	// 1. Declare a pointer
	var intPtr *int

	// 2. Initialize a pointer
	x := 42
	intPtr = &x

	// 3. Access pointer value and address
	fmt.Println("Pointer Value:", *intPtr)
	fmt.Println("Pointer Address:", intPtr)

	// 4. Update the value through a pointer
	*intPtr = 43
	fmt.Println("Updated Value:", x)

	// 5. Zero value of a pointer
	var nilPtr *int
	fmt.Println("Nil Pointer:", nilPtr)

	// 6. Compare two pointers
	y := 42
	yPtr := &y
	fmt.Println("Are intPtr and yPtr equal?", intPtr == yPtr)

	// 7. Passing a pointer to a function
	incrementValue(intPtr)
	fmt.Println("Incremented Value:", x)

	// 8. Returning a pointer from a function
	ptr := createPointer()
	fmt.Println("Returned Pointer:", *ptr)

	// 9. Pointer to an array
	array := [3]int{1, 2, 3}
	arrayPtr := &array
	fmt.Println("Array through Pointer:", (*arrayPtr)[0])

	// 10. Pointer to a struct
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "Alice", Age: 30}
	personPtr := &person
	fmt.Println("Person through Pointer:", (*personPtr).Name)

	// 11. Shortcut for accessing struct fields through a pointer
	fmt.Println("Person Name through Pointer (Shortcut):", personPtr.Name)
}

// 12. Passing a pointer to a function
func incrementValue(ptr *int) {
	*ptr++
}

// 13. Returning a pointer from a function
func createPointer() *int {
	value := 100
	return &value
}
