package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== HTTP Server ===")

	// Basic handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! Requested path: %s", r.URL.Path)
	})

	// JSON handler
	http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"users": [{"id": 1, "name": "Alice"}, {"id": 2, "name": "Bob"}]}`)
	})

	// Form handler
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			if err := r.ParseForm(); err == nil {
				name := r.FormValue("name")
				email := r.FormValue("email")
				fmt.Fprintf(w, "Received: Name=%s, Email=%s", name, email)
			} else {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
			}
		} else {
			fmt.Fprintf(w, `<html><body>
				<form method="post">
					Name: <input name="name"><br>
					Email: <input name="email"><br>
					<input type="submit">
				</form>
			</body></html>`)
		}
	})

	// Headers handler
	http.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Custom-Header", "CustomValue")
		
		fmt.Fprintf(w, `{"method": "%s", "headers": {`, r.Method)
		first := true
		for key, values := range r.Header {
			if !first {
				fmt.Fprintf(w, ",")
			}
			fmt.Fprintf(w, `"%s": "%s"`, key, values[0])
			first = false
		}
		fmt.Fprintf(w, `}}`)
	})

	// Middleware example
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			log.Printf("%s %s", r.Method, r.URL.Path)
			
			next.ServeHTTP(w, r)
			
			duration := time.Since(start)
			log.Printf("Request completed in %v", duration)
		})
	}

	// Apply middleware
	handler := loggingMiddleware(http.DefaultServeMux)

	// Server configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  /                    - Basic response")
	fmt.Println("  /api/users           - JSON response")
	fmt.Println("  /submit              - Form handling")
	fmt.Println("  /headers             - Request headers")
	fmt.Println("\nPress Ctrl+C to stop the server")

	// Start server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server error:", err)
	}
}
