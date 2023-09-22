package apiendpoints

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User endpoint - GET request")
}

// Include more handlers for your API endpoints.
