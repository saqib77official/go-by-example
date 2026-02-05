package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Timeouts Examples ===")

	// 1. Basic timeout with time.After
	fmt.Println("\n1. Basic timeout:")
	ch := make(chan string)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "result"
	}()
	
	select {
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Operation timed out")
	}

	// 2. Timeout with longer duration
	fmt.Println("\n2. Timeout with longer duration:")
	ch2 := make(chan int)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- 42
	}()
	
	select {
	case result := <-ch2:
		fmt.Printf("Received: %d\n", result)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation timed out")
	}

	// 3. Multiple timeouts
	fmt.Println("\n3. Multiple timeouts:")
	ch3 := make(chan string)
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch3 <- "data"
	}()
	
	select {
	case data := <-ch3:
		fmt.Printf("Received: %s\n", data)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Short timeout (100ms)")
	case <-time.After(300 * time.Millisecond):
		fmt.Println("Long timeout (300ms)")
	}

	// 4. Timeout with function returning error
	fmt.Println("\n4. Timeout with error handling:")
	
	operationWithTimeout := func() (string, error) {
		ch := make(chan string)
		
		go func() {
			time.Sleep(200 * time.Millisecond)
			ch <- "operation completed"
		}()
		
		select {
		case result := <-ch:
			return result, nil
		case <-time.After(100 * time.Millisecond):
			return "", fmt.Errorf("operation timed out")
		}
	}
	
	result, err := operationWithTimeout()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success: %s\n", result)
	}

	// 5. Timeout for network simulation
	fmt.Println("\n5. Network operation timeout:")
	
	simulateNetworkCall := func(timeout time.Duration) (string, error) {
		ch := make(chan string)
		
		go func() {
			// Simulate network delay
			time.Sleep(150 * time.Millisecond)
			ch <- "network response"
		}()
		
		select {
		case response := <-ch:
			return response, nil
		case <-time.After(timeout):
			return "", fmt.Errorf("network timeout after %v", timeout)
		}
	}
	
	// Try with short timeout
	response, err := simulateNetworkCall(100 * time.Millisecond)
	if err != nil {
		fmt.Printf("Short timeout error: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", response)
	}
	
	// Try with longer timeout
	response, err = simulateNetworkCall(200 * time.Millisecond)
	if err != nil {
		fmt.Printf("Long timeout error: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", response)
	}

	// 6. Timeout with context cancellation
	fmt.Println("\n6. Timeout with cancellation:")
	
	type Result struct {
		Data  string
		Error error
	}
	
	ch4 := make(chan Result)
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch4 <- Result{Data: "processed data", Error: nil}
	}()
	
	select {
	case result := <-ch4:
		if result.Error != nil {
			fmt.Printf("Operation failed: %v\n", result.Error)
		} else {
			fmt.Printf("Operation succeeded: %s\n", result.Data)
		}
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Operation cancelled due to timeout")
	}

	// 7. Timeout with retry logic
	fmt.Println("\n7. Timeout with retry:")
	
	operationWithRetry := func(maxRetries int, timeout time.Duration) (string, error) {
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("Attempt %d\n", attempt)
			
			ch := make(chan string)
			go func() {
				// Simulate work that might take different times
				delay := time.Duration(attempt*50) * time.Millisecond
				time.Sleep(delay)
				ch <- fmt.Sprintf("success on attempt %d", attempt)
			}()
			
			select {
			case result := <-ch:
				return result, nil
			case <-time.After(timeout):
				fmt.Printf("Attempt %d timed out\n", attempt)
				if attempt == maxRetries {
					return "", fmt.Errorf("operation failed after %d attempts", maxRetries)
				}
			}
		}
		return "", fmt.Errorf("should not reach here")
	}
	
	result, err = operationWithRetry(3, 100*time.Millisecond)
	if err != nil {
		fmt.Printf("Final error: %v\n", err)
	} else {
		fmt.Printf("Final success: %s\n", result)
	}

	// 8. Timeout for concurrent operations
	fmt.Println("\n8. Timeout for concurrent operations:")
	
	ch5 := make(chan int)
	ch6 := make(chan int)
	
	// Start two operations
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch5 <- 100
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch6 <- 200
	}()
	
	// Wait for first result with timeout
	select {
	case result1 := <-ch5:
		fmt.Printf("First operation completed: %d\n", result1)
	case result2 := <-ch6:
		fmt.Printf("Second operation completed: %d\n", result2)
	case <-time.After(150 * time.Millisecond):
		fmt.Println("Both operations timed out")
	}

	// 9. Timeout with cleanup
	fmt.Println("\n9. Timeout with cleanup:")
	
	operationWithCleanup := func(timeout time.Duration) error {
		ch := make(chan string)
		done := make(chan bool)
		
		// Worker goroutine
		go func() {
			defer func() {
				done <- true
			}()
			
			time.Sleep(200 * time.Millisecond)
			ch <- "work completed"
		}()
		
		select {
		case result := <-ch:
			fmt.Printf("Work completed: %s\n", result)
			return nil
		case <-time.After(timeout):
			fmt.Println("Timeout occurred, cleaning up...")
			// Wait for worker to finish cleanup
			<-done
			return fmt.Errorf("operation timed out after %v", timeout)
		}
	}
	
	err = operationWithCleanup(100 * time.Millisecond)
	fmt.Printf("Result: %v\n", err)

	// 10. Timeout with progress reporting
	fmt.Println("\n10. Timeout with progress:")
	
	operationWithProgress := func(timeout time.Duration) {
		ch := make(chan string)
		progress := make(chan int)
		
		// Worker with progress
		go func() {
			defer close(ch)
			
			for i := 1; i <= 5; i++ {
				time.Sleep(50 * time.Millisecond)
				progress <- i * 20 // Progress percentage
			}
			ch <- "completed"
		}()
		
		ticker := time.NewTicker(30 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case result := <-ch:
				fmt.Printf("Final result: %s\n", result)
				return
			case p := <-progress:
				fmt.Printf("Progress: %d%%\n", p)
			case <-ticker.C:
				fmt.Println("Working...")
			case <-time.After(timeout):
				fmt.Printf("Operation timed out after %v\n", timeout)
				return
			}
		}
	}
	
	operationWithProgress(200 * time.Millisecond)

	// 11. Timeout with deadline
	fmt.Println("\n11. Timeout with deadline:")
	
	operationWithDeadline := func(deadline time.Time) error {
		ch := make(chan string)
		
		go func() {
			time.Sleep(200 * time.Millisecond)
			ch <- "result"
		}()
		
		timeout := time.Until(deadline)
		select {
		case result := <-ch:
			fmt.Printf("Completed: %s\n", result)
			return nil
		case <-time.After(timeout):
			return fmt.Errorf("missed deadline at %v", deadline)
		}
	}
	
	deadline := time.Now().Add(100 * time.Millisecond)
	err = operationWithDeadline(deadline)
	fmt.Printf("Deadline result: %v\n", err)

	// 12. Timeout pattern: Circuit breaker
	fmt.Println("\n12. Circuit breaker pattern:")
	
	type CircuitBreaker struct {
		failures    int
		maxFailures int
		timeout     time.Duration
		lastFailure time.Time
	}
	
	(cb *CircuitBreaker) Call() (string, error) {
		// Check if circuit is open
		if cb.failures >= cb.maxFailures {
			if time.Since(cb.lastFailure) < cb.timeout {
				return "", fmt.Errorf("circuit breaker is open")
			}
			// Reset after timeout
			cb.failures = 0
		}
		
		ch := make(chan string)
		go func() {
			// Simulate operation
			if time.Now().Unix()%2 == 0 { // Random failure
				time.Sleep(150 * time.Millisecond)
				return
			}
			time.Sleep(50 * time.Millisecond)
			ch <- "success"
		}()
		
		select {
		case result := <-ch:
			cb.failures = 0
			return result, nil
		case <-time.After(100 * time.Millisecond):
			cb.failures++
			cb.lastFailure = time.Now()
			return "", fmt.Errorf("operation timed out, failures: %d", cb.failures)
		}
	}
	
	breaker := &CircuitBreaker{
		maxFailures: 2,
		timeout:     500 * time.Millisecond,
	}
	
	// Test circuit breaker
	for i := 1; i <= 5; i++ {
		result, err := breaker.Call()
		if err != nil {
			fmt.Printf("Call %d failed: %v\n", i, err)
		} else {
			fmt.Printf("Call %d succeeded: %s\n", i, result)
		}
		time.Sleep(100 * time.Millisecond)
	}

	// 13. Timeout with graceful shutdown
	fmt.Println("\n13. Graceful shutdown with timeout:")
	
	shutdown := make(chan struct{})
	workDone := make(chan bool)
	
	// Worker
	go func() {
		defer close(workDone)
		
		for {
			select {
			case <-shutdown:
				fmt.Println("Received shutdown signal, cleaning up...")
				time.Sleep(50 * time.Millisecond) // Simulate cleanup
				fmt.Println("Cleanup completed")
				return
			default:
				// Do work
				time.Sleep(30 * time.Millisecond)
				fmt.Println("Working...")
			}
		}
	}()
	
	// Let worker run for a bit
	time.Sleep(100 * time.Millisecond)
	
	// Initiate shutdown with timeout
	close(shutdown)
	
	select {
	case <-workDone:
		fmt.Println("Graceful shutdown completed")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Forceful shutdown (timeout)")
	}

	// 14. Timeout with resource management
	fmt.Println("\n14. Timeout with resource management:")
	
	acquireResource := func() (string, func(), error) {
		ch := make(chan string)
		release := make(chan func())
		
		// Resource manager
		go func() {
			resourceID := "resource-123"
			fmt.Printf("Acquiring resource %s...\n", resourceID)
			time.Sleep(50 * time.Millisecond)
			
			ch <- resourceID
			release <- func() {
				fmt.Printf("Releasing resource %s\n", resourceID)
			}
		}()
		
		select {
		case id := <-ch:
			releaseFunc := <-release
			return id, releaseFunc, nil
		case <-time.After(100 * time.Millisecond):
			return "", nil, fmt.Errorf("resource acquisition timeout")
		}
	}
	
	resourceID, release, err := acquireResource()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Using resource %s\n", resourceID)
		time.Sleep(30 * time.Millisecond)
		release()
	}

	fmt.Println("All timeout examples completed!")
}
