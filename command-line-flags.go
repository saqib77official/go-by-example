package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Command Line Flags ===")

	// Define flags
	name := flag.String("name", "World", "Name to greet")
	age := flag.Int("age", 0, "Your age")
	verbose := flag.Bool("verbose", false, "Enable verbose output")
	count := flag.Int("count", 1, "Number of times to repeat")
	help := flag.Bool("help", false, "Show help")

	// Custom flag type
	var output string
	flag.StringVar(&output, "output", "", "Output file")

	// Parse flags
	flag.Parse()

	// Show help if requested
	if *help {
		fmt.Println("Usage: program [options]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Use flag values
	fmt.Printf("Hello, %s!\n", *name)
	
	if *age > 0 {
		fmt.Printf("You are %d years old\n", *age)
	}

	if *verbose {
		fmt.Println("Verbose mode enabled")
	}

	if *count > 1 {
		for i := 0; i < *count; i++ {
			fmt.Printf("Greeting %d: Hello, %s!\n", i+1, *name)
		}
	}

	if *output != "" {
		fmt.Printf("Output will be written to: %s\n", *output)
	}

	// Show remaining arguments
	args := flag.Args()
	if len(args) > 0 {
		fmt.Printf("Additional arguments: %v\n", args)
	}

	// Flag examples
	fmt.Println("\nFlag usage examples:")
	fmt.Println("  ./program -name=Alice -age=25")
	fmt.Println("  ./program -verbose -count=3")
	fmt.Println("  ./program -output=result.txt file1.txt file2.txt")
	fmt.Println("  ./program -help")

	// Demonstrate flag parsing
	fmt.Println("\nFlag parsing demonstration:")
	fmt.Printf("name flag: %s (default: World)\n", *name)
	fmt.Printf("age flag: %d (default: 0)\n", *age)
	fmt.Printf("verbose flag: %t (default: false)\n", *verbose)
	fmt.Printf("count flag: %d (default: 1)\n", *count)
	fmt.Printf("help flag: %t (default: false)\n", *help)

	// Check required flags
	if *name == "World" {
		fmt.Println("\nWarning: Using default name. Use -name to specify.")
	}

	// Validate flag values
	if *age < 0 {
		fmt.Println("Error: Age cannot be negative")
		os.Exit(1)
	}

	if *count < 1 || *count > 10 {
		fmt.Println("Error: Count must be between 1 and 10")
		os.Exit(1)
	}
}
