package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== String Formatting ===")

	// Basic formatting
	name := "Alice"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Printf with different formats
	pi := 3.14159
	fmt.Printf("Pi: %.2f\n", pi)
	fmt.Printf("Pi: %.4f\n", pi)

	// Sprintf (returns string)
	formatted := fmt.Sprintf("User: %s (%d years old)", name, age)
	fmt.Printf("Sprintf result: %s\n", formatted)

	// Width and alignment
	fmt.Printf("|%-10s|%10s|\n", "Left", "Right")
	fmt.Printf("|%-10d|%10d|\n", 42, 42)

	// Verbs
	fmt.Printf("Binary: %b\n", 42)
	fmt.Printf("Hex: %x\n", 42)
	fmt.Printf("Octal: %o\n", 42)

	// String builder
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	fmt.Printf("Builder: %s\n", builder.String())
}
