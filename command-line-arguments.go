package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Command Line Arguments ===")

	// Get all arguments
	args := os.Args
	fmt.Printf("Program name: %s\n", args[0])
	fmt.Printf("Number of arguments: %d\n", len(args)-1)

	// Print all arguments
	fmt.Println("\nAll arguments:")
	for i, arg := range args {
		fmt.Printf("  [%d]: %s\n", i, arg)
	}

	// Process arguments (skip program name)
	if len(args) > 1 {
		fmt.Println("\nProcessing arguments:")
		for i := 1; i < len(args); i++ {
			arg := args[i]
			fmt.Printf("  Arg %d: %s\n", i, arg)
			
			// Handle different argument types
			switch arg {
			case "-h", "--help":
				fmt.Println("    Help flag detected")
			case "-v", "--version":
				fmt.Println("    Version flag detected")
			case "-f", "--file":
				if i+1 < len(args) {
					fmt.Printf("    File: %s\n", args[i+1])
					i++ // Skip next argument as it's the value
				}
			default:
				fmt.Printf("    Unknown argument: %s\n", arg)
			}
		}
	} else {
		fmt.Println("\nNo arguments provided")
		fmt.Println("Usage: program [options] [files...]")
		fmt.Println("Options:")
		fmt.Println("  -h, --help     Show help")
		fmt.Println("  -v, --version  Show version")
		fmt.Println("  -f, --file     Specify file")
	}

	// Example usage simulation
	fmt.Println("\nExample usage patterns:")
	fmt.Println("  ./program -h")
	fmt.Println("  ./program -v")
	fmt.Println("  ./program -f input.txt")
	fmt.Println("  ./program file1.txt file2.txt")

	// Argument validation
	fmt.Println("\nArgument validation:")
	if len(args) < 2 {
		fmt.Println("Error: No arguments provided")
		os.Exit(1)
	}

	// Count specific argument types
	flagCount := 0
	fileCount := 0
	
	for i := 1; i < len(args); i++ {
		if args[i][0] == '-' {
			flagCount++
		} else {
			fileCount++
		}
	}
	
	fmt.Printf("Flags: %d, Files: %d\n", flagCount, fileCount)
}
