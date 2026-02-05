package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== String Functions ===")

	// Length
	text := "Hello, World!"
	fmt.Printf("Length: %d\n", len(text))

	// Contains
	fmt.Printf("Contains 'World': %t\n", strings.Contains(text, "World"))

	// Index
	fmt.Printf("Index of 'World': %d\n", strings.Index(text, "World"))

	// Replace
	fmt.Printf("Replace: %s\n", strings.Replace(text, "World", "Go", 1))

	// ToUpper/ToLower
	fmt.Printf("Upper: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower: %s\n", strings.ToLower(text))

	// Trim
	whitespace := "   Hello   "
	fmt.Printf("Trimmed: '%s'\n", strings.TrimSpace(whitespace))

	// Split
	words := strings.Split(text, " ")
	fmt.Printf("Split: %v\n", words)

	// Join
	joined := strings.Join(words, "-")
	fmt.Printf("Joined: %s\n", joined)

	// Prefix/Suffix
	fmt.Printf("HasPrefix 'Hello': %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("HasSuffix '!': %t\n", strings.HasSuffix(text, "!"))
}
