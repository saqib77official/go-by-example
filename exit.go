package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Exit ===")

	// Different exit codes
	fmt.Println("--- Exit Codes ---")
	fmt.Println("0: Success")
	fmt.Println("1: General error")
	fmt.Println("2: Misuse of shell builtins")
	fmt.Println("126: Command invoked cannot execute")
	fmt.Println("127: Command not found")
	fmt.Println("128: Invalid exit argument")
	fmt.Println("130: Script terminated by Control-C")
	fmt.Println("255*: Exit status out of range")

	// Exit with success code
	fmt.Println("\n--- Exit with Success ---")
	fmt.Println("This will exit with code 0 (success)")
	// os.Exit(0) // Commented out to allow program to continue

	// Exit with error code
	fmt.Println("\n--- Exit with Error ---")
	fmt.Println("This would exit with code 1 (error)")
	// os.Exit(1) // Commented out to allow program to continue

	// Conditional exit
	fmt.Println("\n--- Conditional Exit ---")
	shouldExit := false
	if shouldExit {
		fmt.Println("Exiting due to condition")
		// os.Exit(2)
	} else {
		fmt.Println("Continuing execution")
	}

	// Exit with cleanup
	fmt.Println("\n--- Exit with Cleanup ---")
	fmt.Println("Performing cleanup before exit...")
	
	// In a real application, you might:
	// 1. Close database connections
	// 2. Save state
	// 3. Release resources
	// 4. Log shutdown
	
	fmt.Println("Cleanup completed")
	// os.Exit(0)

	// Exit from function
	fmt.Println("\n--- Exit from Function ---")
	exitFromFunction := func(code int) {
		fmt.Printf("Exiting with code %d from function\n", code)
		// os.Exit(code)
		fmt.Println("This line won't be reached after os.Exit")
	}
	
	exitFromFunction(3)
	fmt.Println("This also won't be reached")

	// Panic vs Exit
	fmt.Println("\n--- Panic vs Exit ---")
	fmt.Println("Panic:")
	fmt.Println("  - Shows stack trace")
	fmt.Println("  - Can be recovered")
	fmt.Println("  - Indicates unexpected error")
	fmt.Println()
	fmt.Println("Exit:")
	fmt.Println("  - No stack trace")
	fmt.Println("  - Cannot be recovered")
	fmt.Println("  - Indicates controlled termination")

	// Exit status checking (shell simulation)
	fmt.Println("\n--- Exit Status Checking ---")
	fmt.Println("In shell, you can check exit status:")
	fmt.Println("  echo $?  # Shows last exit status")
	fmt.Println("  command1 && command2  # Run command2 only if command1 succeeds")
	fmt.Println("  command1 || command2  # Run command2 only if command1 fails")

	// Common exit patterns
	fmt.Println("\n--- Common Exit Patterns ---")
	
	patterns := []struct {
		situation string
		code      int
		reason    string
	}{
		{"Normal completion", 0, "Program completed successfully"},
		{"Invalid arguments", 1, "User provided invalid arguments"},
		{"File not found", 2, "Required file doesn't exist"},
		{"Permission denied", 3, "Insufficient permissions"},
		{"Network error", 4, "Network connectivity issue"},
		{"Configuration error", 5, "Invalid configuration"},
		{"Out of memory", 6, "Insufficient memory"},
		{"Timeout", 7, "Operation timed out"},
	}

	for _, p := range patterns {
		fmt.Printf("  %s: %d (%s)\n", p.situation, p.code, p.reason)
	}

	// Graceful shutdown
	fmt.Println("\n--- Graceful Shutdown ---")
	fmt.Println("Best practices for graceful shutdown:")
	fmt.Println("1. Handle signals (SIGINT, SIGTERM)")
	fmt.Println("2. Close resources properly")
	fmt.Println("3. Save state if needed")
	fmt.Println("4. Use appropriate exit codes")
	fmt.Println("5. Log shutdown process")

	fmt.Println("\nProgram completed normally")
	fmt.Println("Use os.Exit(code) to terminate with specific status")
}
