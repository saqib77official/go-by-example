package main

import "fmt"

// Basic variadic function
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Variadic function with other parameters
func greet(greeting string, names ...string) {
	fmt.Printf("%s", greeting)
	for i, name := range names {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(name)
	}
	fmt.Println("!")
}

// Variadic function returning multiple values
func calculate(operation string, numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}

	switch operation {
	case "sum":
		total := 0.0
		for _, num := range numbers {
			total += num
		}
		return total, nil
	case "average":
		total := 0.0
		for _, num := range numbers {
			total += num
		}
		return total / float64(len(numbers)), nil
	case "max":
		max := numbers[0]
		for _, num := range numbers {
			if num > max {
				max = num
			}
		}
		return max, nil
	case "min":
		min := numbers[0]
		for _, num := range numbers {
			if num < min {
				min = num
			}
		}
		return min, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}

// Variadic function that modifies slice
func filter(predicate func(int) bool, numbers ...int) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

// Variadic function with string formatting
func formatMessage(template string, args ...interface{}) string {
	return fmt.Sprintf(template, args...)
}

// Variadic function that returns a slice
func collectStrings(items ...string) []string {
	return items
}

// Variadic function with error handling
func divideNumbers(divisor float64, numbers ...float64) ([]float64, error) {
	if divisor == 0 {
		return nil, fmt.Errorf("division by zero")
	}

	result := make([]float64, len(numbers))
	for i, num := range numbers {
		result[i] = num / divisor
	}
	return result, nil
}

// Variadic function that builds a map
func buildMap(pairs ...string) map[string]string {
	result := make(map[string]string)
	
	for i := 0; i < len(pairs); i += 2 {
		if i+1 < len(pairs) {
			result[pairs[i]] = pairs[i+1]
		}
	}
	return result
}

// Variadic function with default values
func printNumbers(prefix string, numbers ...int) {
	if len(numbers) == 0 {
		fmt.Printf("%s: No numbers provided\n", prefix)
		return
	}
	
	fmt.Printf("%s: ", prefix)
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}

// Variadic function that concatenates strings
func concatenate(sep string, parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += sep + parts[i]
	}
	return result
}

// Variadic function with validation
func validateAndProcess(min, max int, numbers ...int) ([]int, error) {
	var valid []int
	
	for _, num := range numbers {
		if num < min || num > max {
			return nil, fmt.Errorf("number %d is out of range [%d, %d]", num, min, max)
		}
		valid = append(valid, num)
	}
	
	return valid, nil
}

func main() {
	fmt.Println("=== Variadic Functions Examples ===")

	// 1. Basic variadic function
	fmt.Println("\n1. Basic variadic function:")
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5: %d\n", total)
	
	total = sum(10, 20, 30)
	fmt.Printf("Sum of 10,20,30: %d\n", total)
	
	total = sum() // No arguments
	fmt.Printf("Sum of no numbers: %d\n", total)

	// 2. Variadic function with other parameters
	fmt.Println("\n2. Variadic with other parameters:")
	greet("Hello", "Alice", "Bob", "Charlie")
	greet("Hi", "David")

	// 3. Variadic function returning multiple values
	fmt.Println("\n3. Variadic returning multiple values:")
	result, err := calculate("sum", 1.5, 2.5, 3.5)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Sum: %.1f\n", result)
	}
	
	result, err = calculate("average", 10, 20, 30, 40, 50)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Average: %.1f\n", result)
	}
	
	result, err = calculate("max", 5, 2, 8, 1, 9)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Max: %.1f\n", result)
	}

	// 4. Passing slice to variadic function
	fmt.Println("\n4. Passing slice to variadic function:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total = sum(numbers...)
	fmt.Printf("Sum of slice %v: %d\n", numbers, total)

	// 5. Variadic function with predicate
	fmt.Println("\n5. Variadic with predicate function:")
	even := filter(func(n int) bool { return n%2 == 0 }, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Even numbers: %v\n", even)
	
	positive := filter(func(n int) bool { return n > 0 }, -5, -3, 0, 2, 4, 6)
	fmt.Printf("Positive numbers: %v\n", positive)

	// 6. Variadic function with interface{}
	fmt.Println("\n6. Variadic with interface{}:")
	message := formatMessage("Hello %s, you are %d years old and scored %.1f%%", "Alice", 25, 95.5)
	fmt.Printf("Formatted message: %s\n", message)

	// 7. Variadic function returning slice
	fmt.Println("\n7. Variadic returning slice:")
	collection := collectStrings("apple", "banana", "orange", "grape")
	fmt.Printf("Collected strings: %v\n", collection)

	// 8. Variadic function with error handling
	fmt.Println("\n8. Variadic with error handling:")
	results, err := divideNumbers(2, 10, 20, 30, 40)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Division results: %v\n", results)
	}
	
	_, err = divideNumbers(0, 10, 20)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 9. Variadic function building map
	fmt.Println("\n9. Variadic building map:")
	config := buildMap("host", "localhost", "port", "8080", "debug", "true")
	fmt.Printf("Config map: %v\n", config)

	// 10. Variadic function with default behavior
	fmt.Println("\n10. Variadic with default behavior:")
	printNumbers("Numbers", 1, 2, 3)
	printNumbers("Empty")

	// 11. Variadic function for string concatenation
	fmt.Println("\n11. String concatenation:")
	joined := concatenate("-", "2023", "12", "25")
	fmt.Printf("Joined with '-': %s\n", joined)
	
	joined = concatenate(" ", "Hello", "beautiful", "world")
	fmt.Printf("Joined with ' ': %s\n", joined)

	// 12. Variadic function with validation
	fmt.Println("\n12. Variadic with validation:")
	valid, err := validateAndProcess(1, 100, 5, 10, 15, 20)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Valid numbers: %v\n", valid)
	}
	
	_, err = validateAndProcess(1, 10, 5, 15, 0, 8) // 0 is out of range
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 13. Mixing regular and variadic parameters
	fmt.Println("\n13. Mixed parameters:")
	process := func(op string, x int, nums ...int) int {
		total := x
		for _, n := range nums {
			if op == "add" {
				total += n
			} else if op == "multiply" {
				total *= n
			}
		}
		return total
	}
	
	result = process("add", 10, 5, 3, 2)
	fmt.Printf("Add 10 + 5 + 3 + 2 = %d\n", result)
	
	result = process("multiply", 2, 3, 4)
	fmt.Printf("Multiply 2 * 3 * 4 = %d\n", result)
}
