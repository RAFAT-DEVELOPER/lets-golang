# Go Reflection Example
This Go program demonstrates various use cases for reflection, a powerful feature that allows you to inspect and manipulate variables and structs at runtime in a dynamic way.



# Use Cases
### Case 1: Getting the Type of a Variable
Demonstrates how to get the type of a variable using reflect.TypeOf.
### Case 2: Getting the Type of a Struct
Demonstrates how to get the type of a struct using reflect.TypeOf.
### Case 3: Getting the Kind of a Type
Shows how to get the kind of a type (e.g., struct, int, string) using reflect.ValueOf(...).Kind().
### Case 4: Creating and Initializing a Variable Using Reflection
Illustrates how to create and initialize a variable dynamically using reflection with reflect.New and modify its value.
### Case 5: Getting and Setting Values of Struct Fields
Demonstrates how to get and set values of struct fields dynamically using reflection.
### Case 6: Calling a Method on a Struct Using Reflection
Shows how to call a method on a struct using reflection with reflect.ValueOf(...).MethodByName.