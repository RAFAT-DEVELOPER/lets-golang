package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	// Middleware: Logging
	n := negroni.New()
	n.Use(negroni.NewLogger())

	// Middleware: Authentication (simulated)
	authMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate authentication logic here
			// In a real application, you would check for valid tokens or sessions.
			if r.Header.Get("Authorization") != "Bearer valid_token" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Middleware: Rate Limiting
	rateLimitMiddleware := func(next http.Handler) http.Handler {
		// In a real application, you would implement rate limiting logic here.
		// For simplicity, we're allowing one request per second for demonstration.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Second) // Simulate rate limiting
			next.ServeHTTP(w, r)
		})
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Apply middleware
	n.Use(negroni.HandlerFunc(authMiddleware))
	n.Use(negroni.HandlerFunc(rateLimitMiddleware))

	// Attach the router to the middleware
	n.UseHandler(r)

	// Start the server
	http.ListenAndServe(":8080", n)
}
