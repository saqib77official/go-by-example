package main

import "fmt"

// Constants declared at package level
const PI = 3.14159
const GREETING = "Hello, World!"
const MAX_USERS = 1000

// Grouped constants
const (
	STATUS_ACTIVE   = "active"
	STATUS_INACTIVE = "inactive"
	STATUS_PENDING  = "pending"
)

// Enumerated constants using iota
const (
	RED = iota
	ORANGE
	YELLOW
	GREEN
	BLUE
	INDIGO
	VIOLET
)

// Typed constants
const (
	SECONDS_IN_MINUTE int = 60
	MINUTES_IN_HOUR   int = 60
	HOURS_IN_DAY      int = 24
)

func main() {
	// Local constants
	const TAX_RATE = 0.08
	const MIN_AGE = 18

	fmt.Println("=== Constants Example ===")
	fmt.Printf("PI: %.5f\n", PI)
	fmt.Printf("Greeting: %s\n", GREETING)
	fmt.Printf("Max users: %d\n", MAX_USERS)
	fmt.Printf("Status values: %s, %s, %s\n", STATUS_ACTIVE, STATUS_INACTIVE, STATUS_PENDING)
	fmt.Printf("Color values: %d, %d, %d, %d, %d, %d, %d\n", RED, ORANGE, YELLOW, GREEN, BLUE, INDIGO, VIOLET)
	fmt.Printf("Time constants: %d seconds/minute, %d minutes/hour, %d hours/day\n", 
		SECONDS_IN_MINUTE, MINUTES_IN_HOUR, HOURS_IN_DAY)
	fmt.Printf("Tax rate: %.2f, Minimum age: %d\n", TAX_RATE, MIN_AGE)

	// Constants cannot be reassigned (this would cause a compile error)
	// PI = 3.14 // Error: cannot assign to PI
}
