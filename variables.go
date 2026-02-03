package main

import "fmt"

func main() {
	// Variable declarations with var keyword
	var name string = "John"
	var age int = 25
	var height float64 = 5.9
	var isStudent bool = true

	// Type inference
	var city = "New York"
	var score = 95.5

	// Short declaration (only inside functions)
	country := "USA"
	grade := "A"

	// Multiple variable declaration
	var (
		x int = 10
		y int = 20
		z int = 30
	)

	// Zero values
	var defaultInt int
	var defaultString string
	var defaultBool bool

	fmt.Println("=== Variables Example ===")
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Student: %t\n", name, age, height, isStudent)
	fmt.Printf("City: %s, Score: %.1f\n", city, score)
	fmt.Printf("Country: %s, Grade: %s\n", country, grade)
	fmt.Printf("Coordinates: (%d, %d, %d)\n", x, y, z)
	fmt.Printf("Zero values: %d, '%s', %t\n", defaultInt, defaultString, defaultBool)

	// Variable reassignment
	age = 26
	fmt.Printf("Updated age: %d\n", age)
}
