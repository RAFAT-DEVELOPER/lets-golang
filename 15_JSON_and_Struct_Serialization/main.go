package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct for a sample object
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Case 1: Creating a Struct and Serializing to JSON
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	// Serialize the struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal("Error encoding JSON:", err)
	}
	fmt.Println("Serialized JSON data:")
	fmt.Println(string(jsonData))

	// Case 2: Deserializing JSON into a Struct
	jsonStr := `{"name":"Alice","age":25,"address":"456 Elm St"}`

	// Create a new Person struct
	var anotherPerson Person

	// Deserialize JSON into the struct
	err = json.Unmarshal([]byte(jsonStr), &anotherPerson)
	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	fmt.Println("\nDeserialized Struct:")
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", anotherPerson.Name, anotherPerson.Age, anotherPerson.Address)

	// Case 3: Handling JSON Errors
	invalidJSON := `{"name": "Bob", "age": "thirty"}`

	// Attempt to decode invalid JSON into a struct
	var invalidPerson Person
	err = json.Unmarshal([]byte(invalidJSON), &invalidPerson)
	if err != nil {
		fmt.Println("\nError decoding JSON:", err)
	}

	// Case 4: Working with JSON Arrays
	jsonArray := `[{"name":"Eve","age":22,"address":"789 Oak St"},{"name":"Mallory","age":35,"address":"101 Pine St"}]`

	// Create a slice of Person structs
	var people []Person

	// Deserialize JSON array into the slice
	err = json.Unmarshal([]byte(jsonArray), &people)
	if err != nil {
		log.Fatal("Error decoding JSON array:", err)
	}
	fmt.Println("\nDeserialized Slice of Structs:")
	for i, p := range people {
		fmt.Printf("Person %d:\nName: %s\nAge: %d\nAddress: %s\n", i+1, p.Name, p.Age, p.Address)
	}
}
