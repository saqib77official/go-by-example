package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Panic Examples ===")

	// 1. Basic panic
	fmt.Println("\n1. Basic panic:")
	fmt.Println("About to panic...")
	// panic("Something went wrong!")
	fmt.Println("This line won't execute if panic is uncommented")

	// 2. Panic from invalid operation
	fmt.Println("\n2. Panic from invalid operation:")
	var slice []int
	// slice[0] = 1 // This would panic
	fmt.Println("Slice access commented out to avoid panic")

	// 3. Panic from nil pointer
	fmt.Println("\n3. Panic from nil pointer:")
	var ptr *int
	// *ptr = 42 // This would panic
	fmt.Println("Nil pointer access commented out to avoid panic")

	// 4. Panic from type assertion
	fmt.Println("\n4. Panic from type assertion:")
	var i interface{} = "hello"
	// num := i.(int) // This would panic
	fmt.Println("Type assertion commented out to avoid panic")

	// 5. Safe type assertion
	fmt.Println("\n5. Safe type assertion:")
	if num, ok := i.(int); ok {
		fmt.Printf("Number: %d\n", num)
	} else {
		fmt.Printf("Not a number, it's a %T\n", i)
	}

	// 6. Function that might panic
	fmt.Println("\n6. Function that might panic:")
	mightPanic := func(shouldPanic bool) {
		if shouldPanic {
			panic("Intentional panic!")
		}
		fmt.Println("Function completed successfully")
	}
	
	mightPanic(false)
	// mightPanic(true) // This would panic

	// 7. Panic with formatted message
	fmt.Println("\n7. Panic with formatted message:")
	// panic(fmt.Sprintf("Error code: %d", 404))
	fmt.Println("Formatted panic commented out to avoid panic")

	fmt.Println("All panic examples completed!")
}
