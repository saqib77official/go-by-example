package main

import "fmt"

func main() {
	fmt.Println("=== Slices Examples ===")

	// Creating slices from arrays
	arr := [5]int{1, 2, 3, 4, 5}
	var slice []int = arr[1:4] // Elements from index 1 to 3 (4 is exclusive)
	fmt.Printf("Slice from array: %v\n", slice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	// Slice literal
	fruits := []string{"apple", "banana", "orange", "grape"}
	fmt.Printf("Fruits slice: %v\n", fruits)

	// Empty slice
	var emptySlice []int
	fmt.Printf("Empty slice: %v, Length: %d, Is nil: %t\n", emptySlice, len(emptySlice), emptySlice == nil)

	// Slice with make
	numbers := make([]int, 5, 10) // Length 5, Capacity 10
	fmt.Printf("Made slice: %v, Length: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

	// Adding elements to slice
	numbers = append(numbers, 1, 2, 3)
	fmt.Printf("After append: %v, Length: %d, Capacity: %d\n", numbers, len(numbers), cap(numbers))

	// Appending another slice
	moreNumbers := []int{4, 5, 6}
	numbers = append(numbers, moreNumbers...)
	fmt.Printf("After appending slice: %v\n", numbers)

	// Copying slices
	source := []int{10, 20, 30, 40, 50}
	destination := make([]int, 3)
	copied := copy(destination, source)
	fmt.Printf("Source: %v\n", source)
	fmt.Printf("Destination: %v\n", destination)
	fmt.Printf("Number of elements copied: %d\n", copied)

	// Slice operations
	fmt.Println("\nSlice operations:")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Slicing
	fmt.Printf("Original: %v\n", data)
	fmt.Printf("data[2:5]: %v\n", data[2:5])   // Elements 2,3,4
	fmt.Printf("data[:3]: %v\n", data[:3])     // Elements 0,1,2
	fmt.Printf("data[7:]: %v\n", data[7:])     // Elements 7,8,9
	fmt.Printf("data[:]: %v\n", data[:])       // All elements

	// Modifying slice elements
	slice := []int{1, 2, 3, 4, 5}
	slice[2] = 99
	fmt.Printf("Modified slice: %v\n", slice)

	// Reslicing
	slice = slice[1:4]
	fmt.Printf("Resliced: %v\n", slice)

	// Iterating over slices
	fmt.Println("\nIterating over slice:")
	for i, value := range slice {
		fmt.Printf("Index %d: %d\n", i, value)
	}

	// Slices are reference types
	fmt.Println("\nSlices are reference types:")
	original := []int{1, 2, 3}
	reference := original
	reference[0] = 999
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Reference: %v\n", reference)

	// Growing slice beyond capacity
	fmt.Println("\nGrowing slice beyond capacity:")
	small := make([]int, 3, 3) // Length 3, Capacity 3
	fmt.Printf("Before: %v, Len: %d, Cap: %d\n", small, len(small), cap(small))
	small = append(small, 4) // This will create a new underlying array
	fmt.Printf("After: %v, Len: %d, Cap: %d\n", small, len(small), cap(small))

	// Multi-dimensional slices
	fmt.Println("\nMulti-dimensional slices:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	for i, row := range matrix {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	// Filtering slices
	fmt.Println("\nFiltering slice:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var even []int
	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	fmt.Printf("Even numbers: %v\n", even)

	// Slice deletion (remove element at index)
	fmt.Println("\nDeleting from slice:")
	items := []string{"a", "b", "c", "d", "e"}
	index := 2
	items = append(items[:index], items[index+1:]...)
	fmt.Printf("After removing index %d: %v\n", index, items)
}
