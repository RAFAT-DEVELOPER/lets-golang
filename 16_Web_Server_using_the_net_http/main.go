package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Middleware function to log incoming requests
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request URL: %s", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Struct for JSON response
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Case 1: Basic HTTP server with a simple response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Case 2: Handling GET requests with URL parameters
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	// Case 3: Handling POST requests with form data
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}
			message := r.FormValue("message")
			fmt.Fprintf(w, "Received message: %s", message)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	// Case 4: Serving static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Case 5: Custom 404 Page
	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Custom 404 Page - Page not found!")
	})

	// Case 6: Middleware for request logging
	mux := http.NewServeMux()
	mux.Handle("/", logRequest(http.DefaultServeMux))

	// Case 7: Handling JSON Responses
	http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		data := Response{Message: "This is a JSON response"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
