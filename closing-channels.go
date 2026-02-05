package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Closing Channels Examples ===")

	// 1. Basic channel closing
	fmt.Println("\n1. Basic channel closing:")
	ch := make(chan int)
	
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("Sent: %d\n", i)
		}
		close(ch)
		fmt.Println("Channel closed")
	}()
	
	// Receive until closed
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Range completed")

	// 2. Checking if channel is closed
	fmt.Println("\n2. Checking if channel is closed:")
	ch2 := make(chan string)
	
	go func() {
		ch2 <- "hello"
		ch2 <- "world"
		close(ch2)
	}()
	
	// Receive with ok check
	for {
		value, ok := <-ch2
		if !ok {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Printf("Received: %s\n", value)
	}

	// 3. Closing multiple channels
	fmt.Println("\n3. Closing multiple channels:")
	ch3 := make(chan int)
	ch4 := make(chan string)
	
	go func() {
		for i := 1; i <= 3; i++ {
			ch3 <- i
			ch4 <- fmt.Sprintf("item-%d", i)
		}
		close(ch3)
		close(ch4)
	}()
	
	// Receive from both channels
	ch3Open, ch4Open := true, true
	for ch3Open || ch4Open {
		select {
		case num, ok := <-ch3:
			if ok {
				fmt.Printf("Number: %d\n", num)
			} else {
				fmt.Println("ch3 closed")
				ch3Open = false
			}
		case str, ok := <-ch4:
			if ok {
				fmt.Printf("String: %s\n", str)
			} else {
				fmt.Println("ch4 closed")
				ch4Open = false
			}
		}
	}

	// 4. Closing buffered channels
	fmt.Println("\n4. Closing buffered channels:")
	buffered := make(chan int, 3)
	
	// Fill buffer
	buffered <- 1
	buffered <- 2
	buffered <- 3
	
	fmt.Printf("Buffer length before close: %d\n", len(buffered))
	close(buffered)
	
	// Can still receive buffered items after close
	for len(buffered) > 0 {
		value := <-buffered
		fmt.Printf("Received from closed buffered: %d\n", value)
	}
	
	// Receive from closed empty channel
	value, ok := <-buffered
	fmt.Printf("Receive from closed empty: %v, %t\n", value, ok)

	// 5. Panic on closed channel
	fmt.Println("\n5. Panic on closed channel:")
	ch5 := make(chan int)
	
	go func() {
		ch5 <- 42
		close(ch5)
	}()
	
	// Safe receive
	value, ok := <-ch5
	fmt.Printf("Safe receive: %d, %t\n", value, ok)
	
	// This would panic (commented out)
	// ch5 <- 100 // Panic: send on closed channel

	// 6. Closing channels from receiver side
	fmt.Println("\n6. Closing from receiver side:")
	ch6 := make(chan int)
	
	go func() {
		// Producer
		for i := 1; i <= 5; i++ {
			select {
			case ch6 <- i:
				fmt.Printf("Sent: %d\n", i)
			case <-time.After(10 * time.Millisecond):
				fmt.Printf("Send timeout for %d\n", i)
			}
		}
	}()
	
	// Receiver that decides when to close
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(ch6)
		fmt.Println("Receiver closed channel")
	}()
	
	// Receive until closed
	for value := range ch6 {
		fmt.Printf("Received: %d\n", value)
	}

	// 7. Closing with select
	fmt.Println("\n7. Closing with select:")
	ch7 := make(chan string)
	stopChan := make(chan bool)
	
	go func() {
		for i := 1; i <= 5; i++ {
			select {
			case ch7 <- fmt.Sprintf("message-%d", i):
				fmt.Printf("Sent: message-%d\n", i)
			case <-stopChan:
				fmt.Println("Received stop signal")
				close(ch7)
				return
			}
			time.Sleep(20 * time.Millisecond)
		}
		close(ch7)
	}()
	
	// Stop after 3 messages
	time.Sleep(70 * time.Millisecond)
	stopChan <- true
	
	// Receive remaining messages
	for msg := range ch7 {
		fmt.Printf("Received: %s\n", msg)
	}

	// 8. Channel closing patterns
	fmt.Println("\n8. Channel closing patterns:")
	
	// Pattern 1: Producer closes
	producerCloses := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 1; i <= 3; i++ {
				ch <- i
			}
		}()
		return ch
	}
	
	// Pattern 2: Separate done channel
	producerWithDone := func() (<-chan int, <-chan struct{}) {
		ch := make(chan int)
		done := make(chan struct{})
		
		go func() {
			defer close(ch)
			for i := 1; i <= 3; i++ {
				ch <- i
			}
		}()
		
		return ch, done
	}
	
	// Test pattern 1
	fmt.Println("Pattern 1 - Producer closes:")
	ch8 := producerCloses()
	for value := range ch8 {
		fmt.Printf("Received: %d\n", value)
	}
	
	// Test pattern 2
	fmt.Println("Pattern 2 - Separate done channel:")
	ch9, done := producerWithDone()
	
	for {
		select {
		case value, ok := <-ch9:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Received: %d\n", value)
		case <-done:
			fmt.Println("Done signal received")
			return
		}
	}

	// 9. Closing channels in fan-out
	fmt.Println("\n9. Closing channels in fan-out:")
	input := make(chan int)
	outputs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}
	
	// Distributor
	go func() {
		defer close(input)
		for i := 1; i <= 6; i++ {
			input <- i
		}
	}()
	
	// Workers
	for i, output := range outputs {
		go func(id int, out chan int) {
			defer close(out)
			for value := range input {
				out <- value * id
				fmt.Printf("Worker %d: %d -> %d\n", id, value, value*id)
			}
		}(i+1, output)
	}
	
	// Collect from all outputs
	for i := 0; i < 6*3; i++ { // 6 inputs * 3 workers
		select {
		case value := <-outputs[0]:
			fmt.Printf("From worker 1: %d\n", value)
		case value := <-outputs[1]:
			fmt.Printf("From worker 2: %d\n", value)
		case value := <-outputs[2]:
			fmt.Printf("From worker 3: %d\n", value)
		}
	}

	// 10. Graceful shutdown with channel closing
	fmt.Println("\n10. Graceful shutdown:")
	workChan := make(chan int)
	shutdownChan := make(chan struct{})
	
	// Worker
	go func() {
		for {
			select {
			case work := <-workChan:
				fmt.Printf("Processing work: %d\n", work)
				time.Sleep(30 * time.Millisecond)
			case <-shutdownChan:
				fmt.Println("Received shutdown signal")
				return
			}
		}
	}()
	
	// Send work
	for i := 1; i <= 5; i++ {
		workChan <- i
	}
	
	// Initiate graceful shutdown
	close(shutdownChan)
	time.Sleep(50 * time.Millisecond)
	close(workChan)

	// 11. Channel closing with resource cleanup
	fmt.Println("\n11. Channel closing with cleanup:")
	resourceChan := make(chan string)
	
	// Resource manager
	go func() {
		defer fmt.Println("Resource cleaned up")
		
		for resource := range resourceChan {
			fmt.Printf("Using resource: %s\n", resource)
			time.Sleep(30 * time.Millisecond)
		}
	}()
	
	// Send resources
	resources := []string{"file1", "file2", "file3"}
	for _, resource := range resources {
		resourceChan <- resource
	}
	
	close(resourceChan) // Trigger cleanup
	time.Sleep(50 * time.Millisecond)

	// 12. Detecting closed channel without receiving
	fmt.Println("\n12. Detecting closed channel without receiving:")
	ch10 := make(chan int)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		close(ch10)
	}()
	
	// Check if channel is closed without consuming data
	for i := 0; i < 5; i++ {
		select {
		case <-ch10:
			fmt.Println("Channel has data or is closed")
		default:
			fmt.Printf("Channel appears open (check %d)\n", i+1)
		}
		time.Sleep(20 * time.Millisecond)
	}
	
	// Final check
	select {
	case <-ch10:
		fmt.Println("Channel confirmed closed")
	default:
		fmt.Println("Channel appears open")
	}

	// 13. Channel closing with error handling
	fmt.Println("\n13. Channel closing with error handling:")
	type Result struct {
		Value int
		Error error
	}
	
	resultChan := make(chan Result)
	
	go func() {
		defer close(resultChan)
		
		// Simulate some work with potential errors
		for i := 1; i <= 5; i++ {
			if i == 3 {
				resultChan <- Result{Error: fmt.Errorf("error processing %d", i)}
				continue
			}
			resultChan <- Result{Value: i * 10}
		}
	}()
	
	// Process results
	for result := range resultChan {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		} else {
			fmt.Printf("Success: %d\n", result.Value)
		}
	}

	// 14. Channel closing with timeout
	fmt.Println("\n14. Channel closing with timeout:")
	ch11 := make(chan int)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		close(ch11)
	}()
	
	// Wait for close with timeout
	timeout := time.After(100 * time.Millisecond)
	
Loop:
	for {
		select {
		case value, ok := <-ch11:
			if !ok {
				fmt.Println("Channel closed")
				break Loop
			}
			fmt.Printf("Received: %d\n", value)
		case <-timeout:
			fmt.Println("Timeout waiting for close")
			break Loop
		}
	}

	// 15. Channel closing statistics
	fmt.Println("\n15. Channel closing statistics:")
	statsChan := make(chan int)
	
	go func() {
		defer close(statsChan)
		sent := 0
		for i := 1; i <= 10; i++ {
			statsChan <- i
			sent++
			time.Sleep(10 * time.Millisecond)
		}
		fmt.Printf("Sent %d items before closing\n", sent)
	}()
	
	// Count received
	received := 0
	for value := range statsChan {
		received++
		fmt.Printf("Received: %d\n", value)
	}
	
	fmt.Printf("Total received: %d\n", received)

	fmt.Println("All channel closing examples completed!")
}
