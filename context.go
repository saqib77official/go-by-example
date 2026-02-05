package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Context ===")

	// Basic context with timeout
	fmt.Println("--- Context with Timeout ---")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed (this won't print)")
	case <-ctx.Done():
		fmt.Printf("Operation cancelled: %v\n", ctx.Err())
	}

	// Context with cancellation
	fmt.Println("\n--- Context with Cancellation ---")
	ctx, cancel = context.WithCancel(context.Background())

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling context...")
		cancel()
	}()

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Printf("Operation cancelled: %v\n", ctx.Err())
	}

	// Context with value
	fmt.Println("\n--- Context with Value ---")
	ctx = context.WithValue(context.Background(), "userID", 12345)
	ctx = context.WithValue(ctx, "role", "admin")

	userID := ctx.Value("userID")
	role := ctx.Value("role")
	
	fmt.Printf("User ID: %v\n", userID)
	fmt.Printf("Role: %v\n", role)

	// Context with deadline
	fmt.Println("\n--- Context with Deadline ---")
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()

	fmt.Printf("Deadline: %v\n", deadline)
	if d, ok := ctx.Deadline(); ok {
		fmt.Printf("Time until deadline: %v\n", time.Until(d))
	}

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Operation completed before deadline")
	case <-ctx.Done():
		fmt.Printf("Operation cancelled: %v\n", ctx.Err())
	}

	// Context propagation
	fmt.Println("\n--- Context Propagation ---")
	parentCtx, parentCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer parentCancel()

	// Create child context
	childCtx, childCancel := context.WithCancel(parentCtx)
	defer childCancel()

	// Simulate work
	go func() {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Child operation completed")
		case <-childCtx.Done():
			fmt.Printf("Child cancelled: %v\n", childCtx.Err())
		}
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Parent operation completed")
	case <-parentCtx.Done():
		fmt.Printf("Parent cancelled: %v\n", parentCtx.Err())
	}

	// Context in HTTP requests (simulation)
	fmt.Println("\n--- Context in HTTP Simulation ---")
	requestCtx := context.WithValue(context.Background(), "requestID", "req-123")
	requestCtx = context.WithTimeout(requestCtx, 1*time.Second)

	// Simulate HTTP handler
	handleRequest := func(ctx context.Context) {
		requestID := ctx.Value("requestID")
		fmt.Printf("Handling request %s\n", requestID)

		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Request %s completed\n", requestID)
		case <-ctx.Done():
			fmt.Printf("Request %s cancelled: %v\n", requestID, ctx.Err())
		}
	}

	handleRequest(requestCtx)

	// Context error types
	fmt.Println("\n--- Context Error Types ---")
	testContextError := func(err error) {
		switch err {
		case context.Canceled:
			fmt.Printf("Error: Context was cancelled\n")
		case context.DeadlineExceeded:
			fmt.Printf("Error: Context deadline exceeded\n")
		default:
			fmt.Printf("Error: %v\n", err)
		}
	}

	// Test cancelled context
	ctx, cancel = context.WithCancel(context.Background())
	cancel()
	testContextError(ctx.Err())

	// Test deadline exceeded context
	ctx, _ = context.WithTimeout(context.Background(), 1*time.Nanosecond)
	time.Sleep(1 * time.Millisecond)
	testContextError(ctx.Err())
}
