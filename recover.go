package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Recover Examples ===")

	// 1. Basic recover
	fmt.Println("\n1. Basic recover:")
	basicRecover := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered: %v\n", r)
			}
		}()
		
		panic("Basic panic example")
	}
	
	basicRecover()

	// 2. Recover in different function
	fmt.Println("\n2. Recover in different function:")
	panickingFunction := func() {
		panic("Panic in nested function")
	}
	
	recoveringFunction := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from nested: %v\n", r)
			}
		}()
		
		panickingFunction()
	}
	
	recoveringFunction()

	// 3. Recover with error handling
	fmt.Println("\n3. Recover with error handling:")
	safeOperation := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic recovered: %v", r)
			}
		}()
		
		// Simulate operation that might panic
		var slice []int
		slice[0] = 1 // This will panic
		
		return nil
	}
	
	if err := safeOperation(); err != nil {
		fmt.Printf("Operation failed: %v\n", err)
	}

	// 4. Recover with cleanup
	fmt.Println("\n4. Recover with cleanup:")
	operationWithCleanup := func() {
		fmt.Println("Starting operation")
		
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic occurred: %v\n", r)
				fmt.Println("Performing cleanup...")
			}
			fmt.Println("Cleanup completed")
		}()
		
		fmt.Println("Performing work...")
		panic("Something went wrong")
	}
	
	operationWithCleanup()

	// 5. Recover with multiple panics
	fmt.Println("\n5. Recover with multiple panics:")
	multiplePanics := func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("First recover: %v\n", r)
			}
		}()
		
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Second recover: %v\n", r)
			}
		}()
		
		panic("Multiple panic test")
	}
	
	multiplePanics()

	// 6. Recover with goroutines
	fmt.Println("\n6. Recover with goroutines:")
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Goroutine recovered: %v\n", r)
			}
		}()
		
		panic("Goroutine panic")
	}()
	
	// Give goroutine time to execute
	fmt.Println("Main function continuing...")
	fmt.Println("Goroutine panic handled separately")

	fmt.Println("All recover examples completed!")
}
