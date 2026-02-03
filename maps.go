package main

import "fmt"

func main() {
	fmt.Println("=== Maps Examples ===")

	// Creating maps with make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	fmt.Printf("Ages map: %v\n", ages)

	// Map literal
	grades := map[string]int{
		"Math":     90,
		"Science":  85,
		"History":  78,
		"English":  92,
	}
	fmt.Printf("Grades map: %v\n", grades)

	// Empty map
	emptyMap := make(map[string]string)
	fmt.Printf("Empty map: %v, Length: %d\n", emptyMap, len(emptyMap))

	// Nil map
	var nilMap map[string]int
	fmt.Printf("Nil map: %v, Is nil: %t\n", nilMap, nilMap == nil)

	// Accessing map values
	fmt.Println("\nAccessing map values:")
	fmt.Printf("Alice's age: %d\n", ages["Alice"])
	fmt.Printf("Math grade: %d\n", grades["Math"])

	// Checking if key exists
	fmt.Println("\nChecking key existence:")
	if age, exists := ages["Alice"]; exists {
		fmt.Printf("Alice exists and is %d years old\n", age)
	} else {
		fmt.Println("Alice does not exist in the map")
	}

	if age, exists := ages["David"]; exists {
		fmt.Printf("David exists and is %d years old\n", age)
	} else {
		fmt.Println("David does not exist in the map")
	}

	// Modifying map values
	fmt.Println("\nModifying map values:")
	ages["Alice"] = 26
	fmt.Printf("Updated Alice's age: %d\n", ages["Alice"])

	// Adding new key-value pairs
	ages["David"] = 40
	fmt.Printf("After adding David: %v\n", ages)

	// Deleting from maps
	fmt.Println("\nDeleting from maps:")
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n", ages)

	// Iterating over maps
	fmt.Println("\nIterating over map:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Iterating over maps (keys only)
	fmt.Println("\nIterating over keys:")
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}

	// Map length
	fmt.Printf("\nNumber of people in ages map: %d\n", len(ages))

	// Maps with different value types
	fmt.Println("\nMaps with different value types:")

	// Map with slice values
	studentsByGrade := map[string][]string{
		"A": {"Alice", "Bob"},
		"B": {"Charlie", "David"},
		"C": {"Eve", "Frank"},
	}
	fmt.Printf("Students by grade: %v\n", studentsByGrade)

	// Map with struct values
	type Person struct {
		Name string
		Age  int
		City string
	}

	people := map[int]Person{
		1: {Name: "Alice", Age: 25, City: "New York"},
		2: {Name: "Bob", Age: 30, City: "Los Angeles"},
		3: {Name: "Charlie", Age: 35, City: "Chicago"},
	}
	fmt.Printf("People map: %v\n", people)

	// Nested maps
	fmt.Println("\nNested maps:")
	departments := map[string]map[string]int{
		"Engineering": {
			"Frontend": 5,
			"Backend":  8,
			"DevOps":   3,
		},
		"Marketing": {
			"Digital": 4,
			"Content": 2,
		},
	}
	
	for dept, teams := range departments {
		fmt.Printf("%s Department:\n", dept)
		for team, count := range teams {
			fmt.Printf("  %s: %d employees\n", team, count)
		}
	}

	// Map operations
	fmt.Println("\nMap operations:")
	
	// Counting occurrences
	words := []string{"apple", "banana", "apple", "orange", "banana", "apple"}
	wordCount := make(map[string]int)
	
	for _, word := range words {
		wordCount[word]++
	}
	
	fmt.Printf("Word count: %v\n", wordCount)

	// Finding max value in map
	fmt.Println("\nFinding maximum value:")
	maxGrade := ""
	maxScore := -1
	
	for subject, score := range grades {
		if score > maxScore {
			maxScore = score
			maxGrade = subject
		}
	}
	fmt.Printf("Highest grade: %s with %d\n", maxGrade, maxScore)

	// Clearing a map
	fmt.Println("\nClearing a map:")
	fmt.Printf("Before clear - Length: %d\n", len(ages))
	for key := range ages {
		delete(ages, key)
	}
	fmt.Printf("After clear - Length: %d\n", len(ages))
}
