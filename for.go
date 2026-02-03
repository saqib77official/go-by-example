package main

import "fmt"

func main() {
	fmt.Println("=== For Loop Examples ===")

	// Basic for loop (like C/Java for loop)
	fmt.Println("\n1. Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// For loop as while loop
	fmt.Println("\n2. For loop as while loop:")
	count := 0
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// Infinite loop with break
	fmt.Println("\n3. Infinite loop with break:")
	num := 0
	for {
		if num >= 3 {
			break
		}
		fmt.Printf("Number: %d\n", num)
		num++
	}

	// For loop with continue
	fmt.Println("\n4. For loop with continue:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("Odd number: %d\n", i)
	}

	// For range over slice
	fmt.Println("\n5. For range over slice:")
	fruits := []string{"apple", "banana", "orange"}
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// For range over map
	fmt.Println("\n6. For range over map:")
	ages := map[string]int{"Alice": 25, "Bob": 30, "Charlie": 35}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// For range over string
	fmt.Println("\n7. For range over string:")
	message := "Hello"
	for index, char := range message {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// Nested for loops
	fmt.Println("\n8. Nested for loops:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
