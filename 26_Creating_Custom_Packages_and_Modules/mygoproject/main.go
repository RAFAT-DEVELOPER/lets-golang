package main

import (
	"fmt"
	"mygoproject/apiendpoints"
	"mygoproject/codeorg"
	"mygoproject/customtypes"
	"mygoproject/reusable"
	"mygoproject/testingutil"
)

func main() {
	// Use Case 1: Code Organization
	fmt.Println("Use Case 1: Code Organization")
	codeorg.OrganizationDemo()

	// Use Case 2: Reusable Utilities
	fmt.Println("\nUse Case 2: Reusable Utilities")
	result := reusable.Add(10, 20)
	fmt.Printf("Addition Result: %d\n", result)

	// Use Case 3: API Endpoints
	fmt.Println("\nUse Case 3: API Endpoints")
	api := apiendpoints.NewAPI()
	api.Handle("/hello", apiendpoints.HelloHandler)
	api.Handle("/user", apiendpoints.UserHandler)
	go api.StartServer()

	// Use Case 4: Database Abstraction (Not demonstrated in this simple example)

	// Use Case 5: Middleware (Not demonstrated in this simple example)

	// Use Case 6: Custom Data Types
	fmt.Println("\nUse Case 6: Custom Data Types")
	user := customtypes.NewUser("John", "Doe", 30)
	fmt.Printf("User Details: %s\n", user)

	// Use Case 7: Testing
	fmt.Println("\nUse Case 7: Testing")
	testingutil.RunTests()

	// Use Case 8: Third-party Integrations (Not demonstrated in this simple example)

	// Block to keep the server running
	select {}
}
