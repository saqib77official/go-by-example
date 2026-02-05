package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Non-Blocking Channel Operations Examples ===")

	// 1. Non-blocking receive
	fmt.Println("\n1. Non-blocking receive:")
	ch := make(chan int)
	
	select {
	case value := <-ch:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("No data available (non-blocking receive)")
	}
	
	// Send data and try non-blocking receive
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- 42
	}()
	
	time.Sleep(100 * time.Millisecond) // Wait for send
	
	select {
	case value := <-ch:
		fmt.Printf("Now received: %d\n", value)
	default:
		fmt.Println("Still no data")
	}

	// 2. Non-blocking send
	fmt.Println("\n2. Non-blocking send:")
	ch2 := make(chan string) // Unbuffered
	
	select {
	case ch2 <- "hello":
		fmt.Println("Sent successfully")
	default:
		fmt.Println("Could not send (channel blocking)")
	}
	
	// Non-blocking send to buffered channel
	ch3 := make(chan string, 1)
	
	select {
	case ch3 <- "hello":
		fmt.Println("Sent to buffered channel")
	default:
		fmt.Println("Could not send")
	}
	
	// Try to send to full buffered channel
	ch4 := make(chan int, 1)
	ch4 <- 100 // Fill buffer
	
	select {
	case ch4 <- 200:
		fmt.Println("Sent to full buffered channel")
	default:
		fmt.Println("Buffered channel full, cannot send")
	}

	// 3. Non-blocking operations in loop
	fmt.Println("\n3. Non-blocking operations in loop:")
	ch5 := make(chan int, 2)
	
	// Fill buffer partially
	ch5 <- 1
	ch5 <- 2
	
	// Try to send more
	for i := 3; i <= 5; i++ {
		select {
		case ch5 <- i:
			fmt.Printf("Sent: %d\n", i)
		default:
			fmt.Printf("Could not send %d (channel full)\n", i)
		}
	}
	
	// Drain channel
	for len(ch5) > 0 {
		value := <-ch5
		fmt.Printf("Drained: %d\n", value)
	}

	// 4. Non-blocking receive from multiple channels
	fmt.Println("\n4. Non-blocking receive from multiple channels:")
	ch6 := make(chan string)
	ch7 := make(chan int)
	
	// Send to one channel
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch6 <- "message"
	}()
	
	// Try non-blocking receive from both
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch6:
			fmt.Printf("Received string: %s\n", msg)
		case num := <-ch7:
			fmt.Printf("Received int: %d\n", num)
		default:
			fmt.Printf("No data available (attempt %d)\n", i+1)
		}
		time.Sleep(30 * time.Millisecond)
	}

	// 5. Non-blocking with timeout simulation
	fmt.Println("\n5. Non-blocking with timeout simulation:")
	ch8 := make(chan int)
	
	// Simulate timeout using non-blocking select
	deadline := time.Now().Add(100 * time.Millisecond)
	
	for {
		select {
		case value := <-ch8:
			fmt.Printf("Received: %d\n", value)
			break
		default:
			if time.Now().After(deadline) {
				fmt.Println("Timeout reached")
				break
			}
			fmt.Println("Waiting...")
			time.Sleep(20 * time.Millisecond)
		}
	}

	// 6. Non-blocking producer-consumer
	fmt.Println("\n6. Non-blocking producer-consumer:")
	prodChan := make(chan int, 3)
	consChan := make(chan int, 3)
	
	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			select {
			case prodChan <- i:
				fmt.Printf("Produced: %d\n", i)
			default:
				fmt.Printf("Production buffer full, skipping: %d\n", i)
			}
			time.Sleep(50 * time.Millisecond)
		}
		close(prodChan)
	}()
	
	// Consumer
	go func() {
		for {
			select {
			case item, ok := <-prodChan:
				if !ok {
					close(consChan)
					return
				}
				select {
				case consChan <- item:
					fmt.Printf("Consumed: %d\n", item)
				default:
					fmt.Printf("Consumer buffer full, dropping: %d\n", item)
				}
			}
		}
	}()
	
	// Collect results
	for item := range consChan {
		fmt.Printf("Final result: %d\n", item)
		time.Sleep(30 * time.Millisecond)
	}

	// 7. Non-blocking with multiple sends
	fmt.Println("\n7. Non-blocking with multiple sends:")
	outputs := []chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}
	
	data := []int{10, 20, 30, 40, 50}
	
	for _, item := range data {
		sent := false
		for i, ch := range outputs {
			select {
			case ch <- item:
				fmt.Printf("Sent %d to output %d\n", item, i)
				sent = true
				break
			default:
				continue
			}
		}
		if !sent {
			fmt.Printf("Could not send %d to any output\n", item)
		}
		time.Sleep(30 * time.Millisecond)
	}

	// 8. Non-blocking fan-in
	fmt.Println("\n8. Non-blocking fan-in:")
	inputs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}
	
	// Start producers
	for i, input := range inputs {
		go func(id int, ch chan int) {
			for j := 1; j <= 3; j++ {
				ch <- id*10 + j
				time.Sleep(50 * time.Millisecond)
			}
			close(ch)
		}(i, input)
	}
	
	// Non-blocking fan-in
	for {
		received := false
		for i, input := range inputs {
			select {
			case value, ok := <-input:
				if ok {
					fmt.Printf("Received %d from input %d\n", value, i)
					received = true
				}
			default:
				continue
			}
		}
		if !received {
			// Check if all inputs are closed
			allClosed := true
			for _, input := range inputs {
				select {
				case <-input:
					allClosed = false
				default:
					if len(input) > 0 {
						allClosed = false
					}
				}
			}
			if allClosed {
				break
			}
		}
		time.Sleep(20 * time.Millisecond)
	}

	// 9. Non-blocking with buffer management
	fmt.Println("\n9. Non-blocking with buffer management:")
	bufferedChan := make(chan int, 3)
	
	// Fill buffer
	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	
	fmt.Printf("Buffer length: %d, capacity: %d\n", len(bufferedChan), cap(bufferedChan))
	
	// Try to add more
	for i := 4; i <= 6; i++ {
		select {
		case bufferedChan <- i:
			fmt.Printf("Added %d to buffer\n", i)
		default:
			fmt.Printf("Buffer full, removing oldest item\n")
			<-bufferedChan // Remove oldest
			bufferedChan <- i // Add new item
			fmt.Printf("Added %d after removal\n", i)
		}
		fmt.Printf("Buffer length: %d\n", len(bufferedChan))
	}

	// 10. Non-blocking with priority
	fmt.Println("\n10. Non-blocking with priority:")
	highPriority := make(chan int, 2)
	normalPriority := make(chan int, 2)
	lowPriority := make(chan int, 2)
	
	// Send to different priority channels
	go func() {
		highPriority <- 1
		time.Sleep(20 * time.Millisecond)
		normalPriority <- 2
		time.Sleep(20 * time.Millisecond)
		lowPriority <- 3
		time.Sleep(20 * time.Millisecond)
		highPriority <- 4
	}()
	
	// Process with priority
	for i := 0; i < 4; i++ {
		select {
		case item := <-highPriority:
			fmt.Printf("High priority: %d\n", item)
		case item := <-normalPriority:
			fmt.Printf("Normal priority: %d\n", item)
		case item := <-lowPriority:
			fmt.Printf("Low priority: %d\n", item)
		default:
			fmt.Println("No items available")
		}
		time.Sleep(30 * time.Millisecond)
	}

	// 11. Non-blocking with backpressure handling
	fmt.Println("\n11. Non-blocking with backpressure:")
	workChan := make(chan int, 2)
	
	// Producer with backpressure handling
	go func() {
		for i := 1; i <= 8; i++ {
			select {
			case workChan <- i:
				fmt.Printf("Accepted work: %d\n", i)
			default:
				fmt.Printf("Rejected work: %d (backpressure)\n", i)
			}
			time.Sleep(30 * time.Millisecond)
		}
		close(workChan)
	}()
	
	// Consumer
	for work := range workChan {
		fmt.Printf("Processing work: %d\n", work)
		time.Sleep(50 * time.Millisecond)
	}

	// 12. Non-blocking with load shedding
	fmt.Println("\n12. Non-blocking with load shedding:")
	requestChan := make(chan int, 3)
	processedChan := make(chan int, 3)
	
	// Request generator
	go func() {
		for i := 1; i <= 10; i++ {
			select {
			case requestChan <- i:
				fmt.Printf("Request %d accepted\n", i)
			default:
				fmt.Printf("Request %d dropped (load shedding)\n", i)
			}
			time.Sleep(20 * time.Millisecond)
		}
		close(requestChan)
	}()
	
	// Processor
	go func() {
		for req := range requestChan {
			select {
			case processedChan <- req:
				fmt.Printf("Request %d processed\n", req)
			default:
				fmt.Printf("Request %d dropped (processing overflow)\n", req)
			}
			time.Sleep(40 * time.Millisecond)
		}
		close(processedChan)
	}()
	
	// Collect processed requests
	for processed := range processedChan {
		fmt.Printf("Final processed: %d\n", processed)
	}

	// 13. Non-blocking with health checks
	fmt.Println("\n13. Non-blocking with health checks:")
	dataChan := make(chan int)
	healthChan := make(chan string)
	
	// Data producer
	go func() {
		for i := 1; i <= 5; i++ {
			dataChan <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(dataChan)
	}()
	
	// Health checker
	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				select {
				case healthChan <- "healthy":
					// Health report sent
				default:
					// Health channel full, skip this report
				}
			}
		}
	}()
	
	// Consumer with health monitoring
	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				goto done
			}
			fmt.Printf("Data: %d\n", data)
		case health := <-healthChan:
			fmt.Printf("Health: %s\n", health)
		}
	}
	
done:
	fmt.Println("Data channel closed")

	// 14. Non-blocking with batch processing
	fmt.Println("\n14. Non-blocking batch processing:")
	batchChan := make(chan int, 10)
	
	// Producer
	go func() {
		for i := 1; i <= 15; i++ {
			batchChan <- i
			time.Sleep(20 * time.Millisecond)
		}
		close(batchChan)
	}()
	
	// Batch processor
	for {
		batch := make([]int, 0, 5)
		
		// Collect batch non-blocking
		for len(batch) < 5 {
			select {
			case item, ok := <-batchChan:
				if !ok {
					// Channel closed, process remaining batch
					if len(batch) > 0 {
						fmt.Printf("Final batch: %v\n", batch)
					}
					goto batchDone
				}
				batch = append(batch, item)
			default:
				// No more items right now
				break
			}
		}
		
		if len(batch) > 0 {
			fmt.Printf("Processing batch: %v\n", batch)
		} else {
			// Check if channel is closed
			select {
			case _, ok := <-batchChan:
				if !ok {
					goto batchDone
				}
			default:
				// Wait for more items
				time.Sleep(30 * time.Millisecond)
			}
		}
	}
	
batchDone:
	fmt.Println("All non-blocking examples completed!")
}
