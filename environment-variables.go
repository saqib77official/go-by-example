package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Environment Variables ===")

	// Get specific environment variable
	home := os.Getenv("HOME")
	fmt.Printf("HOME: %s\n", home)

	path := os.Getenv("PATH")
	fmt.Printf("PATH: %s\n", path)

	// Get all environment variables
	fmt.Println("\nAll environment variables:")
	for _, env := range os.Environ() {
		fmt.Printf("  %s\n", env)
	}

	// Set environment variable
	os.Setenv("MY_APP_VAR", "Hello from Go")
	fmt.Printf("\nSet MY_APP_VAR: %s\n", os.Getenv("MY_APP_VAR"))

	// Check if variable exists
	if value, exists := os.LookupEnv("NON_EXISTENT"); exists {
		fmt.Printf("NON_EXISTENT: %s\n", value)
	} else {
		fmt.Println("NON_EXISTENT: not set")
	}

	// Clear environment variable
	os.Unsetenv("MY_APP_VAR")
	if value, exists := os.LookupEnv("MY_APP_VAR"); exists {
		fmt.Printf("MY_APP_VAR after unset: %s\n", value)
	} else {
		fmt.Println("MY_APP_VAR after unset: not set")
	}

	// Working with PATH
	fmt.Println("\nPATH analysis:")
	pathEntries := strings.Split(path, string(os.PathListSeparator))
	fmt.Printf("PATH has %d entries\n", len(pathEntries))
	for i, entry := range pathEntries {
		if i < 5 { // Show first 5 entries
			fmt.Printf("  %d: %s\n", i, entry)
		}
	}
	if len(pathEntries) > 5 {
		fmt.Printf("  ... and %d more\n", len(pathEntries)-5)
	}

	// Environment variable expansion
	fmt.Println("\nEnvironment variable patterns:")
	fmt.Printf("Current working directory: %s\n", os.Getenv("PWD"))
	fmt.Printf("User: %s\n", os.Getenv("USER"))
	fmt.Printf("Shell: %s\n", os.Getenv("SHELL"))

	// Common environment variables
	commonVars := []string{
		"HOME", "USER", "PATH", "SHELL", "TERM", 
		"LANG", "PWD", "GOPATH", "GOROOT",
	}

	fmt.Println("\nCommon environment variables:")
	for _, varName := range commonVars {
		if value := os.Getenv(varName); value != "" {
			fmt.Printf("  %s: %s\n", varName, value)
		} else {
			fmt.Printf("  %s: (not set)\n", varName)
		}
	}

	// Environment variable for configuration
	fmt.Println("\nUsing environment variables for configuration:")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	
	fmt.Printf("Database: %s:%s\n", dbHost, dbPort)

	// Environment variable validation
	fmt.Println("\nEnvironment variable validation:")
	requiredVars := []string{"HOME", "PATH"}
	allSet := true
	
	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			fmt.Printf("Error: Required variable %s is not set\n", varName)
			allSet = false
		}
	}
	
	if allSet {
		fmt.Println("All required environment variables are set")
	}
}
