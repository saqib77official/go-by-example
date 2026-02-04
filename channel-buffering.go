package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Channel Buffering Examples ===")

	// 1. Unbuffered vs Buffered channel
	fmt.Println("\n1. Unbuffered vs Buffered channel:")
	
	// Unbuffered channel (synchronous)
	unbuffered := make(chan int)
	fmt.Printf("Unbuffered channel cap: %d\n", cap(unbuffered))
	
	// Buffered channel (asynchronous)
	buffered := make(chan int, 3)
	fmt.Printf("Buffered channel cap: %d\n", cap(buffered))

	// 2. Buffered channel basics
	fmt.Println("\n2. Buffered channel basics:")
	ch := make(chan string, 2)
	
	// Send to buffered channel (won't block)
	ch <- "First"
	fmt.Println("Sent: First")
	
	ch <- "Second"
	fmt.Println("Sent: Second")
	
	// Receive from buffered channel
	fmt.Printf("Received: %s\n", <-ch)
	fmt.Printf("Received: %s\n", <-ch)

	// 3. Buffering prevents blocking
	fmt.Println("\n3. Buffering prevents blocking:")
	bufferedCh := make(chan int, 3)
	
	// Send multiple values without receiver
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	fmt.Println("Sent 3 values to buffered channel")
	
	// Now receive them
	for i := 0; i < 3; i++ {
		val := <-bufferedCh
		fmt.Printf("Received: %d\n", val)
	}

	// 4. Buffer overflow (blocking)
	fmt.Println("\n4. Buffer overflow demonstration:")
	overflowCh := make(chan int, 2)
	
	// Fill the buffer
	overflowCh <- 1
	overflowCh <- 2
	fmt.Println("Buffer filled (2/2)")
	
	// This would block (commented out to avoid hanging)
	// overflowCh <- 3 // This would block
	
	// Receive one to make space
	fmt.Printf("Received: %d\n", <-overflowCh)
	
	// Now we can send again
	overflowCh <- 3
	fmt.Println("Sent third value after receiving one")
	
	fmt.Printf("Remaining: %d\n", <-overflowCh)
	fmt.Printf("Remaining: %d\n", <-overflowCh)

	// 5. Buffered channel with goroutines
	fmt.Println("\n5. Buffered channel with goroutines:")
	workCh := make(chan int, 5)
	
	// Producer (fast)
	go func() {
		defer close(workCh)
		for i := 1; i <= 5; i++ {
			workCh <- i
			fmt.Printf("Produced: %d\n", i)
		}
	}()
	
	// Consumer (slow)
	time.Sleep(100 * time.Millisecond) // Let producer fill buffer
	
	for work := range workCh {
		fmt.Printf("Consumed: %d\n", work)
		time.Sleep(50 * time.Millisecond)
	}

	// 6. Using len() and cap() with buffered channels
	fmt.Println("\n6. Channel length and capacity:")
	metricsCh := make(chan string, 3)
	
	fmt.Printf("Initial - Len: %d, Cap: %d\n", len(metricsCh), cap(metricsCh))
	
	metricsCh <- "A"
	fmt.Printf("After 1 send - Len: %d, Cap: %d\n", len(metricsCh), cap(metricsCh))
	
	metricsCh <- "B"
	fmt.Printf("After 2 sends - Len: %d, Cap: %d\n", len(metricsCh), cap(metricsCh))
	
	metricsCh <- "C"
	fmt.Printf("After 3 sends - Len: %d, Cap: %d\n", len(metricsCh), cap(metricsCh))
	
	<-metricsCh
	fmt.Printf("After 1 receive - Len: %d, Cap: %d\n", len(metricsCh), cap(metricsCh))

	// 7. Buffered channel as semaphore
	fmt.Println("\n7. Buffered channel as semaphore:")
	// Create semaphore with capacity 3 (max 3 concurrent operations)
	semaphore := make(chan struct{}, 3)
	
	var wg sync.WaitGroup
	
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Acquire semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Release semaphore
			
			fmt.Printf("Task %d started\n", id)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Task %d completed\n", id)
		}(i)
	}
	
	wg.Wait()

	// 8. Buffered channel for rate limiting
	fmt.Println("\n8. Rate limiting with buffered channel:")
	// Create a bucket with capacity 5
	rateLimiter := make(chan time.Time, 5)
	
	// Fill the bucket initially
	for i := 0; i < 5; i++ {
		rateLimiter <- time.Now()
	}
	
	// Refill at rate of 1 per second
	go func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		
		for range ticker.C {
			select {
			case rateLimiter <- time.Now():
				// Token added
			default:
				// Bucket full
			}
		}
	}()
	
	// Process requests
	for i := 1; i <= 8; i++ {
		<-rateLimiter // Wait for token
		fmt.Printf("Request %d processed at %v\n", i, time.Now().Format("15:04:05"))
	}

	// 9. Buffered channel for batching
	fmt.Println("\n9. Batching with buffered channel:")
	batchCh := make(chan int, 10)
	
	// Producer
	go func() {
		defer close(batchCh)
		for i := 1; i <= 15; i++ {
			batchCh <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Batch processor
	batchSize := 5
	for {
		batch := make([]int, 0, batchSize)
		
		// Collect batch
		for len(batch) < batchSize {
			item, ok := <-batchCh
			if !ok {
				break
			}
			batch = append(batch, item)
		}
		
		if len(batch) == 0 {
			break
		}
		
		fmt.Printf("Processing batch: %v\n", batch)
		time.Sleep(100 * time.Millisecond)
	}

	// 10. Buffered channel for fan-out
	fmt.Println("\n10. Fan-out with buffered channel:")
	input := make(chan int, 10)
	
	// Start multiple workers
	workerCount := 3
	var wg2 sync.WaitGroup
	
	for i := 1; i <= workerCount; i++ {
		wg2.Add(1)
		go func(workerID int) {
			defer wg2.Done()
			for item := range input {
				fmt.Printf("Worker %d processing: %d\n", workerID, item)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}
	
	// Send work
	go func() {
		defer close(input)
		for i := 1; i <= 9; i++ {
			input <- i
			fmt.Printf("Sent: %d\n", i)
		}
	}()
	
	wg2.Wait()

	// 11. Buffered channel with timeout
	fmt.Println("\n11. Buffered channel with timeout:")
	timeoutCh := make(chan string, 3)
	
	// Try to send with timeout
	go func() {
		time.Sleep(100 * time.Millisecond)
		timeoutCh <- "Delayed message"
	}()
	
	select {
	case msg := <-timeoutCh:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timeout occurred")
	}
	
	// Try again with longer timeout
	select {
	case msg := <-timeoutCh:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout occurred")
	}

	// 12. Buffered channel for work queue
	fmt.Println("\n12. Work queue with buffered channel:")
	workQueue := make(chan func(), 5)
	
	// Add work to queue
	for i := 1; i <= 5; i++ {
		taskID := i
		workQueue <- func() {
			fmt.Printf("Executing task %d\n", taskID)
			time.Sleep(50 * time.Millisecond)
		}
	}
	
	// Worker processes tasks
	var wg3 sync.WaitGroup
	wg3.Add(1)
	
	go func() {
		defer wg3.Done()
		for task := range workQueue {
			task()
		}
	}()
	
	time.Sleep(300 * time.Millisecond) // Let tasks process
	close(workQueue)
	wg3.Wait()

	// 13. Buffered channel performance consideration
	fmt.Println("\n13. Buffer size performance:")
	
	// Test different buffer sizes
	bufferSizes := []int{0, 1, 10, 100}
	
	for _, size := range bufferSizes {
		testCh := make(chan int, size)
		start := time.Now()
		
		var wg4 sync.WaitGroup
		
		// Producer
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			defer close(testCh)
			for i := 0; i < 1000; i++ {
				testCh <- i
			}
		}()
		
		// Consumer
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			for range testCh {
				// Process
			}
		}()
		
		wg4.Wait()
		duration := time.Since(start)
		fmt.Printf("Buffer size %d: %v\n", size, duration)
	}

	// 14. Buffered channel for resource pooling
	fmt.Println("\n14. Resource pool with buffered channel:")
	type Resource struct {
		ID int
	}
	
	// Create pool of 3 resources
	pool := make(chan *Resource, 3)
	
	// Initialize pool
	for i := 1; i <= 3; i++ {
		pool <- &Resource{ID: i}
	}
	
	var wg5 sync.WaitGroup
	
	// Use resources
	for i := 1; i <= 5; i++ {
		wg5.Add(1)
		go func(taskID int) {
			defer wg5.Done()
			
			// Acquire resource
			resource := <-pool
			fmt.Printf("Task %d acquired resource %d\n", taskID, resource.ID)
			
			// Use resource
			time.Sleep(100 * time.Millisecond)
			
			// Release resource
			pool <- resource
			fmt.Printf("Task %d released resource %d\n", taskID, resource.ID)
		}(i)
	}
	
	wg5.Wait()

	// 15. Buffered channel with select and default
	fmt.Println("\n15. Select with default on buffered channel:")
	selectCh := make(chan string, 2)
	
	// Pre-fill channel
	selectCh <- "Message 1"
	selectCh <- "Message 2"
	
	for i := 1; i <= 4; i++ {
		select {
		case msg := <-selectCh:
			fmt.Printf("Received: %s\n", msg)
		default:
			fmt.Printf("No message available (iteration %d)\n", i)
		}
		
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("All channel buffering examples completed!")
}
