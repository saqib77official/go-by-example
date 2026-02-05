package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Sorting Examples ===")

	// 1. Basic slice sorting
	fmt.Println("\n1. Basic slice sorting:")
	numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
	fmt.Printf("Original: %v\n", numbers)
	
	sort.Ints(numbers)
	fmt.Printf("Sorted: %v\n", numbers)

	// 2. String sorting
	fmt.Println("\n2. String sorting:")
	words := []string{"zebra", "apple", "orange", "banana", "grape"}
	fmt.Printf("Original: %v\n", words)
	
	sort.Strings(words)
	fmt.Printf("Sorted: %v\n", words)

	// 3. Check if sorted
	fmt.Println("\n3. Check if sorted:")
	fmt.Printf("Numbers sorted: %t\n", sort.IntsAreSorted(numbers))
	fmt.Printf("Words sorted: %t\n", sort.StringsAreSorted(words))

	// 4. Reverse sorting
	fmt.Println("\n4. Reverse sorting:")
	reverseNumbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", reverseNumbers)
	
	sort.Sort(sort.Reverse(sort.IntSlice(reverseNumbers)))
	fmt.Printf("Reverse sorted: %v\n", reverseNumbers)

	// 5. Partial sorting
	fmt.Println("\n5. Partial sorting:")
	partial := []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	fmt.Printf("Original: %v\n", partial)
	
	sort.Slice(partial[:5], func(i, j int) bool {
		return partial[i] < partial[j]
	})
	fmt.Printf("First 5 sorted: %v\n", partial)

	// 6. Sorting with custom type
	fmt.Println("\n6. Custom type sorting:")
	type Person struct {
		Name string
		Age  int
	}
	
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
		{"Diana", 22},
	}
	
	fmt.Printf("Original: %v\n", people)
	
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("Sorted by age: %v\n", people)

	fmt.Println("All sorting examples completed!")
}
