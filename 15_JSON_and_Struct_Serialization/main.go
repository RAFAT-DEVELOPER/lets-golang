package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Alice", Age: 30}

	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("JSON:", string(data))
}
