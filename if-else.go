package main

import "fmt"

func main() {
	fmt.Println("=== If/Else Examples ===")

	// Basic if statement
	age := 18
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	}

	// If-else statement
	temperature := 25
	if temperature > 30 {
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's not too hot")
	}

	// If-else if-else chain
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}

	// If with initialization statement
	if num := 10; num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

	// Nested if statements
	hours := 14
	isWeekend := false
	if hours >= 9 && hours <= 17 {
		if !isWeekend {
			fmt.Println("Business hours - Open")
		} else {
			fmt.Println("Weekend - Closed")
		}
	} else {
		fmt.Println("After hours - Closed")
	}

	// If with multiple conditions
	username := "admin"
	password := "secret123"
	if username == "admin" && password == "secret123" {
		fmt.Println("Login successful")
	} else {
		fmt.Println("Login failed")
	}

	// If with logical OR
	day := "Saturday"
	if day == "Saturday" || day == "Sunday" {
		fmt.Println("It's weekend!")
	}

	// Checking if map key exists
	grades := map[string]int{"Alice": 95, "Bob": 87}
	if grade, exists := grades["Alice"]; exists {
		fmt.Printf("Alice's grade: %d\n", grade)
	} else {
		fmt.Println("Alice's grade not found")
	}

	// If with error handling
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// If with error handling (error case)
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}
