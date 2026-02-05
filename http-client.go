package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	fmt.Println("=== HTTP Client ===")

	// Simple GET request
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
		return
	}

	fmt.Printf("Response body (first 100 chars): %s...\n", string(body[:100]))

	// POST request
	fmt.Println("\n--- POST Request ---")
	postData := `{"name": "Alice", "age": 25}`
	
	resp, err = http.Post("https://httpbin.org/post", "application/json", 
		strings.NewReader(postData))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("POST Response: %s\n", string(body))

	// Custom client with timeout
	fmt.Println("\n--- Custom Client with Timeout ---")
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://httpbin.org/delay/2", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Timeout error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Request completed within timeout\n")

	// Adding headers
	fmt.Println("\n--- Request with Headers ---")
	req, _ = http.NewRequest("GET", "https://httpbin.org/headers", nil)
	req.Header.Set("User-Agent", "Go-HTTP-Client/1.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Custom-Header", "custom-value")

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	fmt.Printf("Headers response: %s\n", string(body))

	// Handling different status codes
	fmt.Println("\n--- Status Code Handling ---")
	urls := []string{
		"https://httpbin.org/status/200",
		"https://httpbin.org/status/404",
		"https://httpbin.org/status/500",
	}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 200:
			fmt.Printf("✓ %s: Success\n", url)
		case 404:
			fmt.Printf("✗ %s: Not Found\n", url)
		case 500:
			fmt.Printf("✗ %s: Server Error\n", url)
		default:
			fmt.Printf("? %s: Unknown status %d\n", url, resp.StatusCode)
		}
	}

	// Download file
	fmt.Println("\n--- Download File ---")
	resp, err = http.Get("https://httpbin.org/bytes/1024")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	downloaded, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error downloading: %v\n", err)
		return
	}

	fmt.Printf("Downloaded %d bytes\n", len(downloaded))

	fmt.Println("HTTP client examples completed!")
}
