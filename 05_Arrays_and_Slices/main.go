package main

import (
	"fmt"
)

func main() {
	// Arrays
	fmt.Println("Arrays:")

	// 1. Declare and initialize an array
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println("Array:", numbers)

	// 2. Initialize an array using an array literal
	names := [3]string{"Alice", "Bob", "Charlie"}
	fmt.Println("Array of names:", names)

	// 3. Access elements of an array
	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])

	// 4. Get the length of an array
	arrayLength := len(numbers)
	fmt.Println("Array length:", arrayLength)

	// 5. Iterating over an array using a for loop
	fmt.Println("Iterating over the numbers array:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Slices
	fmt.Println("\nSlices:")

	// 1. Create a slice from an array
	slice1 := numbers[1:4] // [2 3 4]
	fmt.Println("Slice 1:", slice1)

	// 2. Create a slice using a slice literal
	slice2 := []int{5, 6, 7, 8, 9}
	fmt.Println("Slice 2:", slice2)

	// 3. Modify a slice (which affects the underlying array)
	slice1[0] = 20
	fmt.Println("Modified slice 1:", slice1)
	fmt.Println("Original array:", numbers)

	// 4. Append elements to a slice
	slice2 = append(slice2, 10, 11)
	fmt.Println("Updated slice 2:", slice2)

	// 5. Create a slice using make
	slice3 := make([]int, 3, 5) // A slice with length 3 and capacity 5
	fmt.Println("Slice 3:", slice3)

	// 6. Copy slices
	copy(slice3, slice2)
	fmt.Println("Copied slice 3 from slice 2:", slice3)

	// 7. Slicing a slice to create a new slice
	slice4 := slice2[:3] // Take the first 3 elements of slice2
	fmt.Println("Slice 4:", slice4)

	// 8. Get the length and capacity of a slice
	fmt.Println("Length of slice 2:", len(slice2))
	fmt.Println("Capacity of slice 2:", cap(slice2))

	// 9. Nil slices
	var nilSlice []int
	fmt.Println("Nil slice:", nilSlice)
	fmt.Println("Is nilSlice nil?", nilSlice == nil)
}
