package main

import (
	"fmt"
	"math"
	"strings"
)

// Function returning two values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Function returning multiple values of different types
func getPersonInfo(id int) (string, int, string, bool) {
	// Simulate database lookup
	people := map[int]struct {
		name  string
		age   int
		email string
		active bool
	}{
		1: {"Alice", 25, "alice@example.com", true},
		2: {"Bob", 30, "bob@example.com", false},
		3: {"Charlie", 35, "charlie@example.com", true},
	}

	if person, exists := people[id]; exists {
		return person.name, person.age, person.email, person.active
	}
	return "", 0, "", false
}

// Function with named return values
func calculateStats(numbers []float64) (count int, sum float64, average float64, min float64, max float64) {
	if len(numbers) == 0 {
		return 0, 0, 0, 0, 0
	}

	count = len(numbers)
	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		sum += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	average = sum / float64(count)
	return
}

// Function returning coordinates
func getCoordinates() (float64, float64) {
	return 40.7128, -74.0060 // New York coordinates
}

// Function returning validation results
func validateInput(username, password string) (bool, string) {
	if len(username) < 3 {
		return false, "Username must be at least 3 characters"
	}
	if len(password) < 6 {
		return false, "Password must be at least 6 characters"
	}
	return true, "Validation successful"
}

// Function returning both result and status
func findMinMax(numbers []int) (int, int, bool) {
	if len(numbers) == 0 {
		return 0, 0, false
	}

	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max, true
}

// Function returning multiple strings
func splitName(fullName string) (string, string) {
	parts := strings.Split(fullName, " ")
	if len(parts) >= 2 {
		return parts[0], strings.Join(parts[1:], " ")
	}
	return fullName, ""
}

// Function returning calculation and error
func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number")
	}
	return math.Sqrt(num), nil
}

// Function returning multiple values with different purposes
func analyzeText(text string) (int, int, int, []string) {
	words := strings.Fields(text)
	sentences := strings.Split(text, ".")
	paragraphs := strings.Split(text, "\n\n")
	
	// Count words starting with capital letter
	var capitalizedWords []string
	for _, word := range words {
		if len(word) > 0 && strings.Title(word[:1]) == word[:1] {
			capitalizedWords = append(capitalizedWords, word)
		}
	}
	
	return len(words), len(sentences), len(paragraphs), capitalizedWords
}

func main() {
	fmt.Println("=== Multiple Return Values Examples ===")

	// 1. Function returning result and error
	fmt.Println("\n1. Result and error:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", result)
	}

	// 2. Function returning multiple different types
	fmt.Println("\n2. Multiple different types:")
	name, age, email, active := getPersonInfo(1)
	fmt.Printf("Person: %s, Age: %d, Email: %s, Active: %t\n", name, age, email, active)

	// 3. Function with named return values
	fmt.Println("\n3. Named return values:")
	numbers := []float64{10, 20, 30, 40, 50}
	count, sum, avg, min, max := calculateStats(numbers)
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Count: %d, Sum: %.1f, Average: %.1f, Min: %.1f, Max: %.1f\n", count, sum, avg, min, max)

	// 4. Function returning coordinates
	fmt.Println("\n4. Coordinates:")
	lat, lon := getCoordinates()
	fmt.Printf("Location: (%.4f, %.4f)\n", lat, lon)

	// 5. Function returning validation result
	fmt.Println("\n5. Validation:")
	valid, message := validateInput("john", "password123")
	fmt.Printf("Validation: %t, Message: %s\n", valid, message)

	valid, message = validateInput("jo", "123")
	fmt.Printf("Validation: %t, Message: %s\n", valid, message)

	// 6. Function returning min/max and status
	fmt.Println("\n6. Min/Max with status:")
	numbersInt := []int{5, 2, 8, 1, 9, 3}
	minVal, maxVal, found := findMinMax(numbersInt)
	if found {
		fmt.Printf("Min: %d, Max: %d\n", minVal, maxVal)
	} else {
		fmt.Println("No numbers found")
	}

	// 7. Function returning split strings
	fmt.Println("\n7. String splitting:")
	firstName, lastName := splitName("John Doe Smith")
	fmt.Printf("First name: %s, Last name: %s\n", firstName, lastName)

	// 8. Function returning calculation and error
	fmt.Println("\n8. Square root with error:")
	sqrt, err := squareRoot(16)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Square root of 16: %.2f\n", sqrt)
	}

	sqrt, err = squareRoot(-4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Square root of -4: %.2f\n", sqrt)
	}

	// 9. Function returning multiple analysis results
	fmt.Println("\n9. Text analysis:")
	text := "Hello World. This is a test. Go programming is fun!"
	wordCount, sentenceCount, paragraphCount, capitalizedWords := analyzeText(text)
	fmt.Printf("Text: %s\n", text)
	fmt.Printf("Words: %d, Sentences: %d, Paragraphs: %d\n", wordCount, sentenceCount, paragraphCount)
	fmt.Printf("Capitalized words: %v\n", capitalizedWords)

	// 10. Ignoring return values
	fmt.Println("\n10. Ignoring return values:")
	_, _, _, capitalizedWords = analyzeText("Another Example With Capitalized Words")
	fmt.Printf("Only capitalized words: %v\n", capitalizedWords)

	// 11. Using multiple return values in function calls
	fmt.Println("\n11. Chaining function calls:")
	if valid, _ := validateInput("alice", "secure123"); valid {
		fmt.Println("User validation passed")
	}

	// 12. Multiple assignment with function calls
	fmt.Println("\n12. Multiple assignment:")
	a, b := 10, 20
	sum := func(x, y int) int { return x + y }(a, b)
	fmt.Printf("Sum of %d and %d is %d\n", a, b, sum)
}
