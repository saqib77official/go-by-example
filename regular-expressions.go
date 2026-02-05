package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("=== Regular Expressions ===")

	// Basic match
	text := "The quick brown fox jumps over the lazy dog"
	
	// Find all words
	words := regexp.MustCompile(`\w+`)
	fmt.Printf("Words: %v\n", words.FindAllString(text, -1))

	// Find email pattern
	email := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	testEmail := "Contact us at admin@example.com or support@company.org"
	fmt.Printf("Emails: %v\n", email.FindAllString(testEmail, -1))

	// Replace
	replaced := regexp.MustCompile(`\bfox\b`).ReplaceAllString(text, "cat")
	fmt.Printf("Replaced: %s\n", replaced)

	// Split by pattern
	parts := regexp.MustCompile(`\s+`).Split(text, -1)
	fmt.Printf("Split: %v\n", parts)

	// Match and capture groups
	datePattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2023-12-25"
	matches := datePattern.FindStringSubmatch(date)
	if len(matches) > 0 {
		fmt.Printf("Date parts: Year=%s, Month=%s, Day=%s\n", matches[1], matches[2], matches[3])
	}

	// Validate format
	phonePattern := regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`)
	phones := []string{"123-456-7890", "123-456-789", "abc-def-ghij"}
	for _, phone := range phones {
		fmt.Printf("Phone %s valid: %t\n", phone, phonePattern.MatchString(phone))
	}
}
