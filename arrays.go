package main

import "fmt"

func main() {
	fmt.Println("=== Arrays Examples ===")

	// Array declaration with specified size and initialization
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers array: %v\n", numbers)
	fmt.Printf("Length: %d\n", len(numbers))

	// Array with type inference
	fruits := [3]string{"apple", "banana", "orange"}
	fmt.Printf("Fruits array: %v\n", fruits)

	// Array with partial initialization (remaining elements get zero values)
	var partial [5]int = [5]int{10, 20} // Other elements will be 0
	fmt.Printf("Partial array: %v\n", partial)

	// Array with ellipsis (compiler determines size)
	grades := [...]int{90, 85, 78, 92, 88}
	fmt.Printf("Grades array: %v\n", grades)
	fmt.Printf("Length (determined by compiler): %d\n", len(grades))

	// Accessing array elements
	fmt.Printf("First fruit: %s\n", fruits[0])
	fmt.Printf("Last fruit: %s\n", fruits[len(fruits)-1])

	// Modifying array elements
	numbers[0] = 100
	fmt.Printf("Modified numbers array: %v\n", numbers)

	// Iterating over arrays
	fmt.Println("\nIterating with index:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	fmt.Println("\nIterating with range:")
	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}

	// Multidimensional arrays
	fmt.Println("\nMultidimensional arrays:")
	var matrix [3][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9

	fmt.Printf("Matrix:\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Multidimensional array initialization
	matrix2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Printf("Matrix2: %v\n", matrix2)

	// Arrays are value types
	fmt.Println("\nArrays are value types:")
	original := [3]int{1, 2, 3}
	copy := original
	copy[0] = 999
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Copy: %v\n", copy)

	// Array comparison (arrays of same type can be compared)
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	arr3 := [3]int{1, 2, 4}

	fmt.Printf("arr1 == arr2: %t\n", arr1 == arr2)
	fmt.Printf("arr1 == arr3: %t\n", arr1 == arr3)

	// Zero value of an array
	var zeroArray [5]int
	fmt.Printf("Zero array: %v\n", zeroArray)
}
