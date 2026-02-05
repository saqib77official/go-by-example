package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=== Signals ===")

	// Create channel for signals
	sigChan := make(chan os.Signal, 1)

	// Notify for specific signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	fmt.Println("Signal handler started")
	fmt.Printf("Process PID: %d\n", os.Getpid())
	fmt.Println("Send signals to test:")
	fmt.Println("  Ctrl+C (SIGINT)")
	fmt.Println("  kill -TERM <pid>")
	fmt.Println("  kill -INT <pid>")
	fmt.Println("Press Ctrl+C to exit gracefully")

	// Handle signals
	for sig := range sigChan {
		switch sig {
		case syscall.SIGINT:
			fmt.Println("\nReceived SIGINT (Interrupt)")
			fmt.Println("Cleaning up...")
			time.Sleep(1 * time.Second)
			fmt.Println("Goodbye!")
			return
		case syscall.SIGTERM:
			fmt.Println("\nReceived SIGTERM (Terminate)")
			fmt.Println("Terminating gracefully...")
			return
		case os.Interrupt:
			fmt.Println("\nReceived Interrupt signal")
			return
		default:
			fmt.Printf("\nReceived signal: %v\n", sig)
		}
	}
}

// Signal handling with context
func signalWithContext() {
	fmt.Println("\n--- Signal with Context ---")
	
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	
	// Notify for SIGINT and SIGTERM
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Handle signals in goroutine
	go func() {
		sig := <-sigChan
		fmt.Printf("Received signal %v, cancelling context\n", sig)
		cancel()
	}()
	
	// Wait for context cancellation
	<-ctx.Done()
	fmt.Println("Context cancelled, shutting down")
}

// Multiple signal handlers
func multipleSignalHandlers() {
	fmt.Println("\n--- Multiple Signal Handlers ---")
	
	intChan := make(chan os.Signal, 1)
	termChan := make(chan os.Signal, 1)
	hupChan := make(chan os.Signal, 1)
	
	signal.Notify(intChan, syscall.SIGINT)
	signal.Notify(termChan, syscall.SIGTERM)
	signal.Notify(hupChan, syscall.SIGHUP)
	
	for {
		select {
		case <-intChan:
			fmt.Println("SIGINT received")
			return
		case <-termChan:
			fmt.Println("SIGTERM received")
			return
		case <-hupChan:
			fmt.Println("SIGHUP received - reloading configuration")
			// In real app, reload config here
		}
	}
}

// Signal masking
func signalMasking() {
	fmt.Println("\n--- Signal Masking ---")
	
	// Block SIGINT during critical section
	sigChan := make(chan os.Signal, 1)
	
	// Create a mask to block SIGINT
	mask := syscall.SIGINT
	oldMask := syscall.Sigset{}
	
	fmt.Println("Blocking SIGINT during critical operation...")
	// syscall.Sigprocmask(syscall.SIG_BLOCK, &mask, &oldMask)
	
	// Simulate critical work
	time.Sleep(2 * time.Second)
	
	// Restore original mask
	// syscall.Sigprocmask(syscall.SIG_SETMASK, &oldMask, nil)
	fmt.Println("SIGINT unblocked")
	
	signal.Notify(sigChan, syscall.SIGINT)
	<-sigChan
	fmt.Println("SIGINT received after unblock")
}
