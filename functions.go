package main

import "fmt"

// Basic function without parameters
func greet() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func greetPerson(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Function with multiple parameters
func add(a, b int) int {
	return a + b
}

// Function with multiple parameters of different types
func printDetails(name string, age int, city string) {
	fmt.Printf("Name: %s, Age: %d, City: %s\n", name, age, city)
}

// Function with named return values
func calculateRectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Naked return - returns named variables
}

// Function that performs operations
func isEven(num int) bool {
	return num%2 == 0
}

// Function with string manipulation
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Function that modifies slice
func doubleNumbers(numbers []int) []int {
	for i := range numbers {
		numbers[i] *= 2
	}
	return numbers
}

// Function with default-like behavior using zero values
func createUser(name string, age int, email string) string {
	if age == 0 {
		age = 18 // Default age
	}
	if email == "" {
		email = "unknown@example.com" // Default email
	}
	return fmt.Sprintf("User: %s, Age: %d, Email: %s", name, age, email)
}

// Recursive function (basic example)
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Function that calls other functions
func performCalculations(x, y int) {
	sum := add(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, sum)
	
	isEvenResult := isEven(sum)
	fmt.Printf("Is %d even? %t\n", sum, isEvenResult)
	
	if x > 0 && y > 0 {
		fact := factorial(sum)
		fmt.Printf("Factorial of %d is %d\n", sum, fact)
	}
}

// Function with early return
func getGrade(score int) string {
	if score < 0 || score > 100 {
		return "Invalid score"
	}
	
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func main() {
	fmt.Println("=== Functions Examples ===")

	// Calling basic functions
	fmt.Println("\n1. Basic functions:")
	greet()
	greetPerson("Alice")

	// Function with parameters and return value
	fmt.Println("\n2. Function with parameters and return:")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	// Function with multiple parameters
	fmt.Println("\n3. Function with multiple parameters:")
	printDetails("Bob", 25, "New York")

	// Function with named return values
	fmt.Println("\n4. Function with named return values:")
	area, perimeter := calculateRectangle(5.0, 3.0)
	fmt.Printf("Rectangle 5x3: Area=%.1f, Perimeter=%.1f\n", area, perimeter)

	// Boolean function
	fmt.Println("\n5. Boolean function:")
	fmt.Printf("Is 4 even? %t\n", isEven(4))
	fmt.Printf("Is 7 even? %t\n", isEven(7))

	// String manipulation function
	fmt.Println("\n6. String manipulation:")
	original := "Hello"
	reversed := reverseString(original)
	fmt.Printf("Original: %s, Reversed: %s\n", original, reversed)

	// Function that modifies slice
	fmt.Println("\n7. Function modifying slice:")
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", numbers)
	doubled := doubleNumbers(numbers)
	fmt.Printf("Doubled: %v\n", doubled)

	// Function with default-like behavior
	fmt.Println("\n8. Function with default-like behavior:")
	user1 := createUser("John", 25, "john@example.com")
	user2 := createUser("Jane", 0, "") // Using defaults
	fmt.Printf("%s\n", user1)
	fmt.Printf("%s\n", user2)

	// Recursive function
	fmt.Println("\n9. Recursive function:")
	fmt.Printf("Factorial of 5: %d\n", factorial(5))
	fmt.Printf("Factorial of 6: %d\n", factorial(6))

	// Function that calls other functions
	fmt.Println("\n10. Function composition:")
	performCalculations(4, 3)

	// Function with early return
	fmt.Println("\n11. Function with early return:")
	fmt.Printf("Score 95: Grade %s\n", getGrade(95))
	fmt.Printf("Score 75: Grade %s\n", getGrade(75))
	fmt.Printf("Score -5: Grade %s\n", getGrade(-5))

	// Function expressions (anonymous functions)
	fmt.Println("\n12. Function expressions:")
	add := func(a, b int) int {
		return a + b
	}
	fmt.Printf("Anonymous function result: %d\n", add(10, 20))

	// Higher-order function (function as parameter)
	fmt.Println("\n13. Higher-order function:")
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := filter(numbers, isEven)
	fmt.Printf("Even numbers: %v\n", evenNumbers)
}

// Higher-order function that takes a function as parameter
func filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}
