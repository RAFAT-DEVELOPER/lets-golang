package main

import (
	"fmt"
	"math"
)

// Shape is an interface for common shape operations.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle represents a rectangular shape.
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle represents a circular shape.
type Circle struct {
	Radius float64
}

// Implement methods for the Rectangle type.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Implement methods for the Circle type.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Greetable is an interface for objects that can be greeted.
type Greetable interface {
	Greet() string
}

// Person represents a person who can be greeted.
type Person struct {
	Name string
}

// Implement the Greet method for the Person type.
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s.", p.Name)
}

// Animal represents an animal that can be greeted.
type Animal struct {
	Species string
}

// Implement the Greet method for the Animal type.
func (a Animal) Greet() string {
	return fmt.Sprintf("Hello, I am a %s.", a.Species)
}

func main() {
	// Use case 1: Methods for shapes
	rectangle := Rectangle{Width: 5, Height: 10}
	circle := Circle{Radius: 3}

	fmt.Println("Use Case 1: Methods for Shapes")
	fmt.Println("Rectangle:")
	fmt.Printf("Area: %.2f\n", rectangle.Area())
	fmt.Printf("Perimeter: %.2f\n", rectangle.Perimeter())

	fmt.Println("\nCircle:")
	fmt.Printf("Area: %.2f\n", circle.Area())
	fmt.Printf("Perimeter: %.2f\n", circle.Perimeter())

	// Use case 2: Interfaces for greeting
	person := Person{Name: "Alice"}
	animal := Animal{Species: "Lion"}

	fmt.Println("\nUse Case 2: Interfaces for Greeting")
	greetables := []Greetable{person, animal}
	for _, greetable := range greetables {
		fmt.Println(greetable.Greet())
	}
}
