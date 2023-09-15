package main

import (
	"fmt"
)

// 1. Define a struct
type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
	Street  string
	City    string
	ZipCode string
}

func main() {
	// 2. Create an instance of a struct
	person1 := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			ZipCode: "12345",
		},
	}

	// 3. Access struct fields
	fmt.Println("Name:", person1.Name)
	fmt.Println("Age:", person1.Age)
	fmt.Println("Address:", person1.Address)

	// 4. Update struct fields
	person1.Age = 31
	fmt.Println("Updated Age:", person1.Age)

	// 5. Anonymous structs
	employee := struct {
		Name   string
		Salary float64
	}{
		Name:   "Bob",
		Salary: 50000.0,
	}
	fmt.Println("Employee:", employee)

	// 6. Zero value of a struct
	var person2 Person // All fields will have their zero values
	fmt.Println("Zero Value Person:", person2)

	// 7. Compare two structs for equality
	areEqual := person1 == person2
	fmt.Println("Are person1 and person2 equal?", areEqual)

	// 8. Nested structs
	fmt.Println("Street:", person1.Address.Street)
	fmt.Println("City:", person1.Address.City)

	// 9. Struct as a function parameter
	printPerson(person1)

	// 10. Pointer to a struct
	person3 := &Person{Name: "Charlie", Age: 25}
	fmt.Println("Pointer to Person:", person3.Name)

	// 11. Modify struct through a pointer
	person3.Age = 26
	fmt.Println("Modified Age via Pointer:", person3.Age)

	// 12. Struct methods
	person1.PrintName()

	// 13. Struct composition
	type Employee struct {
		Person
		EmployeeID int
	}
	employee1 := Employee{
		Person:     Person{Name: "David", Age: 35},
		EmployeeID: 12345,
	}
	fmt.Println("Employee Name:", employee1.Name)
	fmt.Println("Employee Age:", employee1.Age)
	fmt.Println("Employee ID:", employee1.EmployeeID)
}

// 13. Struct as a function parameter
func printPerson(p Person) {
	fmt.Println("Printing Person:")
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

// 14. Struct methods
func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
