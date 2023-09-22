package main

import (
	"fmt"
	"reflect"
)

// Struct with fields
type Person struct {
	Name  string
	Age   int
	Title string
}

func main() {
	// Case 1: Getting the type of a variable
	var x int
	var y string
	fmt.Println("Type of x:", reflect.TypeOf(x))
	fmt.Println("Type of y:", reflect.TypeOf(y))

	// Case 2: Getting the type of a struct
	p := Person{Name: "Alice", Age: 30, Title: "Engineer"}
	pType := reflect.TypeOf(p)
	fmt.Println("\nType of Person struct:", pType)

	// Case 3: Getting the kind of a type (e.g., struct, int, string)
	fmt.Println("Kind of x:", reflect.ValueOf(x).Kind())
	fmt.Println("Kind of y:", reflect.ValueOf(y).Kind())
	fmt.Println("Kind of Person struct:", pType.Kind())

	// Case 4: Creating and initializing a variable using reflection
	newInt := reflect.New(reflect.TypeOf(x)).Elem()
	newInt.SetInt(42)
	fmt.Println("\nNew integer created with reflection:", newInt.Interface())

	// Case 5: Getting and setting values of struct fields
	fieldName := "Name"
	field := reflect.ValueOf(&p).Elem().FieldByName(fieldName)
	if field.IsValid() {
		fmt.Printf("Value of %s: %v\n", fieldName, field.Interface())
	} else {
		fmt.Printf("Field %s not found\n", fieldName)
	}

	// Case 6: Calling a method on a struct using reflection
	methodName := "PrintInfo"
	method := reflect.ValueOf(&p).MethodByName(methodName)
	if method.IsValid() {
		method.Call(nil)
	} else {
		fmt.Printf("Method %s not found\n", methodName)
	}
}

// PrintInfo method on the Person struct
func (p *Person) PrintInfo() {
	fmt.Printf("Name: %s\nAge: %d\nTitle: %s\n", p.Name, p.Age, p.Title)
}
