package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Number Parsing ===")

	// Parse int
	intStr := "42"
	if num, err := strconv.Atoi(intStr); err == nil {
		fmt.Printf("Parsed int: %d\n", num)
	}

	// Parse int with base
	hexStr := "FF"
	if num, err := strconv.ParseInt(hexStr, 16, 64); err == nil {
		fmt.Printf("Parsed hex: %d\n", num)
	}

	// Parse float
	floatStr := "3.14159"
	if num, err := strconv.ParseFloat(floatStr, 64); err == nil {
		fmt.Printf("Parsed float: %f\n", num)
	}

	// Parse with error handling
	invalid := "not a number"
	if _, err := strconv.Atoi(invalid); err != nil {
		fmt.Printf("Error parsing '%s': %v\n", invalid, err)
	}

	// Format numbers
	num := 42
	fmt.Printf("String from int: %s\n", strconv.Itoa(num))

	// Format with base
	fmt.Printf("Binary: %s\n", strconv.FormatInt(int64(num), 2))
	fmt.Printf("Hex: %s\n", strconv.FormatInt(int64(num), 16))
	fmt.Printf("Octal: %s\n", strconv.FormatInt(int64(num), 8))

	// Parse bool
	trueStr := "true"
	falseStr := "false"
	
	if b, err := strconv.ParseBool(trueStr); err == nil {
		fmt.Printf("Parsed true: %t\n", b)
	}
	
	if b, err := strconv.ParseBool(falseStr); err == nil {
		fmt.Printf("Parsed false: %t\n", b)
	}

	// Quote and unquote
	quoted := strconv.Quote("Hello, World!")
	fmt.Printf("Quoted: %s\n", quoted)
	
	if unquoted, err := strconv.Unquote(quoted); err == nil {
		fmt.Printf("Unquoted: %s\n", unquoted)
	}

	// Parse with different bases
	bases := []int{2, 8, 10, 16}
	for _, base := range bases {
		if num, err := strconv.ParseInt("1010", base, 64); err == nil {
			fmt.Printf("1010 in base %d = %d\n", base, num)
		}
	}
}
