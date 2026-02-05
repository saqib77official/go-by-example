package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Line Filters ===")

	// Create test data
	lines := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
		"grape",
	}

	// Write test file
	file, err := os.Create("fruits.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, line := range lines {
		file.WriteString(line + "\n")
	}

	// Read and filter lines
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	fmt.Println("Lines starting with 'a' or 'b':")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "a") || strings.HasPrefix(line, "b") {
			fmt.Printf("  %s\n", line)
		}
	}

	// Filter by length
	file.Seek(0, 0)
	fmt.Println("\nLines longer than 5 characters:")
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 5 {
			fmt.Printf("  %s (%d chars)\n", line, len(line))
		}
	}

	// Filter by content
	file.Seek(0, 0)
	fmt.Println("\nLines containing 'e':")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "e") {
			fmt.Printf("  %s\n", line)
		}
	}

	// Transform lines
	file.Seek(0, 0)
	fmt.Println("\nUppercase lines:")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("  %s\n", strings.ToUpper(line))
	}

	// Count lines
	file.Seek(0, 0)
	count := 0
	for scanner.Scan() {
		count++
	}
	fmt.Printf("\nTotal lines: %d\n", count)

	// Filter with custom function
	filterLines := func(lines []string, predicate func(string) bool) []string {
		var result []string
		for _, line := range lines {
			if predicate(line) {
				result = append(result, line)
			}
		}
		return result
	}

	// Use custom filter
	longLines := filterLines(lines, func(s string) bool {
		return len(s) >= 5
	})
	fmt.Printf("Lines with 5+ chars: %v\n", longLines)

	// Clean up
	os.Remove("fruits.txt")
}
