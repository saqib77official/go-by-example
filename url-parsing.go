package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("=== URL Parsing ===")

	// Parse URL
	rawURL := "https://user:pass@example.com:8080/path/to/resource?query=value&other=123#fragment"
	
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Scheme: %s\n", parsed.Scheme)
	fmt.Printf("User: %s\n", parsed.User.Username())
	if pass, ok := parsed.User.Password(); ok {
		fmt.Printf("Password: %s\n", pass)
	}
	fmt.Printf("Host: %s\n", parsed.Host)
	fmt.Printf("Path: %s\n", parsed.Path)
	fmt.Printf("RawQuery: %s\n", parsed.RawQuery)
	fmt.Printf("Fragment: %s\n", parsed.Fragment)

	// Parse query parameters
	query := parsed.Query()
	fmt.Printf("Query params: %v\n", query)
	fmt.Printf("Query 'query': %s\n", query.Get("query"))
	fmt.Printf("Query 'other': %s\n", query.Get("other"))

	// Build URL
	newURL := url.URL{
		Scheme:   "https",
		Host:     "api.example.com",
		Path:     "/v1/users",
		RawQuery: "page=1&limit=10",
	}
	fmt.Printf("Built URL: %s\n", newURL.String())

	// Add query parameters
	values := url.Values{}
	values.Add("search", "golang")
	values.Add("sort", "date")
	values.Add("sort", "relevance")
	
	baseURL := "https://example.com/search"
	fullURL := baseURL + "?" + values.Encode()
	fmt.Printf("URL with query: %s\n", fullURL)

	// Parse path segments
	path := "/a/b/c/d"
	segments := strings.Split(strings.Trim(path, "/"), "/")
	fmt.Printf("Path segments: %v\n", segments)

	// URL encoding
	encoded := url.QueryEscape("Hello, World!")
	fmt.Printf("Encoded: %s\n", encoded)
	
	decoded, err := url.QueryUnescape(encoded)
	if err == nil {
		fmt.Printf("Decoded: %s\n", decoded)
	}

	// Check if URL is absolute
	absURL := "https://example.com/path"
	relURL := "/relative/path"
	
	fmt.Printf("Is absolute '%s': %t\n", absURL, isAbsolute(absURL))
	fmt.Printf("Is absolute '%s': %t\n", relURL, isAbsolute(relURL))
}

func isAbsolute(u string) bool {
	parsed, err := url.Parse(u)
	return err == nil && parsed.IsAbs()
}
