package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Defer Examples ===")

	// 1. Basic defer
	fmt.Println("\n1. Basic defer:")
	fmt.Println("Start")
	defer fmt.Println("End")
	fmt.Println("Middle")

	// 2. Multiple defers (LIFO order)
	fmt.Println("\n2. Multiple defers (LIFO):")
	fmt.Println("Start")
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("Middle")

	// 3. Defer with file operations
	fmt.Println("\n3. Defer with file operations:")
	createFile := func() {
		file, err := os.Create("test.txt")
		if err != nil {
			fmt.Printf("Error creating file: %v\n", err)
			return
		}
		defer file.Close()
		
		fmt.Println("File created")
		file.WriteString("Hello, defer!")
		fmt.Println("Data written")
	}
	
	createFile()

	// 4. Defer with function return
	fmt.Println("\n4. Defer with function return:")
	deferExample := func() string {
		defer fmt.Println("Deferred in function")
		return "Function result"
	}
	
	result := deferExample()
	fmt.Printf("Result: %s\n", result)

	// 5. Defer with panic recovery
	fmt.Println("\n5. Defer with panic recovery:")
	deferWithPanic := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		
		fmt.Println("About to panic")
		// panic("Test panic")
		fmt.Println("Panic commented out")
	}
	
	deferWithPanic()

	// 6. Defer with resource cleanup
	fmt.Println("\n6. Defer with resource cleanup:")
	resourceManager := func() {
		fmt.Println("Acquiring resource 1")
		defer fmt.Println("Releasing resource 1")
		
		fmt.Println("Acquiring resource 2")
		defer fmt.Println("Releasing resource 2")
		
		fmt.Println("Doing work with resources")
	}
	
	resourceManager()

	// 7. Defer with named return values
	fmt.Println("\n7. Defer with named return values:")
	deferWithNamedReturn := func() (result int) {
		defer func() {
			result *= 2
			fmt.Printf("Deferred doubled result to: %d\n", result)
		}()
		
		result = 5
		fmt.Printf("Initial result: %d\n", result)
		return
	}
	
	finalResult := deferWithNamedReturn()
	fmt.Printf("Final result: %d\n", finalResult)

	fmt.Println("All defer examples completed!")
}
