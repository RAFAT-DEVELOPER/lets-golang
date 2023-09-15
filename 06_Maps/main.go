package main

import (
	"fmt"
)

func main() {
	// Maps
	fmt.Println("Maps:")

	// 1. Declare and initialize a map
	var emptyMap map[string]int
	fmt.Println("Empty Map:", emptyMap)

	// 2. Initialize a map using a map literal
	scores := map[string]int{
		"Alice": 85,
		"Bob":   92,
		"Eve":   78,
	}
	fmt.Println("Scores Map:", scores)

	// 3. Access elements of a map
	aliceScore := scores["Alice"]
	fmt.Println("Alice's Score:", aliceScore)

	// 4. Add an element to a map
	scores["Charlie"] = 89
	fmt.Println("Updated Scores Map:", scores)

	// 5. Delete an element from a map
	delete(scores, "Eve")
	fmt.Println("Scores Map after deleting Eve:", scores)

	// 6. Check if a key exists in a map
	_, exists := scores["Eve"]
	fmt.Println("Eve's Score Exists:", exists)

	// 7. Iterate over a map using a for loop
	fmt.Println("Iterating over the Scores Map:")
	for name, score := range scores {
		fmt.Println(name, ":", score)
	}

	// 8. Get the length of a map
	mapLength := len(scores)
	fmt.Println("Number of Entries in Scores Map:", mapLength)

	// 9. Nil maps
	var nilMap map[string]int
	fmt.Println("Nil Map:", nilMap)
	fmt.Println("Is nilMap nil?", nilMap == nil)

	// 10. Maps with slices as values
	fruits := map[string][]string{
		"fruits":     {"apple", "banana", "cherry"},
		"vegetables": {"carrot", "broccoli"},
	}
	fmt.Println("Fruits and Vegetables Map:", fruits)

	// 11. Nested maps
	nestedMap := map[string]map[string]int{
		"Math": {
			"Alice": 95,
			"Bob":   88,
		},
		"History": {
			"Eve":     75,
			"Charlie": 92,
		},
	}
	fmt.Println("Nested Scores Map:", nestedMap)
}
