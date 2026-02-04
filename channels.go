package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Channels Examples ===")

	// 1. Basic channel operations
	fmt.Println("\n1. Basic channel operations:")
	ch := make(chan string)
	
	// Send data in goroutine
	go func() {
		ch <- "Hello from goroutine!"
	}()
	
	// Receive data
	message := <-ch
	fmt.Printf("Received: %s\n", message)

	// 2. Channel with direction
	fmt.Println("\n2. Directional channels:")
	// Send-only channel
	sendOnly := make(chan<- int)
	// Receive-only channel
	receiveOnly := make(<-chan int)
	
	fmt.Printf("Send-only channel type: %T\n", sendOnly)
	fmt.Printf("Receive-only channel type: %T\n", receiveOnly)

	// 3. Multiple goroutines communicating
	fmt.Println("\n3. Multiple goroutines communication:")
	ch2 := make(chan int)
	
	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			ch2 <- i
			fmt.Printf("Sent: %d\n", i)
		}
		close(ch2)
	}()
	
	// Consumer
	for value := range ch2 {
		fmt.Printf("Received: %d\n", value)
	}

	// 4. Channel as function parameter
	fmt.Println("\n4. Channel as function parameter:")
	ch3 := make(chan string)
	
	go sendData(ch3, "Hello", "World", "Go")
	
	for msg := range ch3 {
		fmt.Printf("Got: %s\n", msg)
	}

	// 5. Select statement with channels
	fmt.Println("\n5. Select statement:")
	ch4 := make(chan string)
	ch5 := make(chan int)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch4 <- "From channel 4"
	}()
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch5 <- 42
	}()
	
	select {
	case msg := <-ch4:
		fmt.Printf("Received from ch4: %s\n", msg)
	case num := <-ch5:
		fmt.Printf("Received from ch5: %d\n", num)
	default:
		fmt.Println("No data available")
	}

	// 6. Timeout with select
	fmt.Println("\n6. Timeout with select:")
	ch6 := make(chan string)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch6 <- "Delayed message"
	}()
	
	select {
	case msg := <-ch6:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout occurred")
	}

	// 7. Non-blocking channel operations
	fmt.Println("\n7. Non-blocking operations:")
	ch7 := make(chan int)
	
	select {
	case val := <-ch7:
		fmt.Printf("Received: %d\n", val)
	default:
		fmt.Println("No data to receive")
	}
	
	select {
	case ch7 <- 42:
		fmt.Println("Sent data")
	default:
		fmt.Println("Could not send data (channel blocking)")
	}

	// 8. Closing channels
	fmt.Println("\n8. Closing channels:")
	ch8 := make(chan int)
	
	go func() {
		for i := 1; i <= 3; i++ {
			ch8 <- i
		}
		close(ch8)
	}()
	
	for {
		value, ok := <-ch8
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Printf("Received: %d\n", value)
	}

	// 9. Range over channel
	fmt.Println("\n9. Range over channel:")
	ch9 := make(chan string)
	
	go func() {
		defer close(ch9)
		items := []string{"apple", "banana", "cherry"}
		for _, item := range items {
			ch9 <- item
		}
	}()
	
	for item := range ch9 {
		fmt.Printf("Item: %s\n", item)
	}

	// 10. Worker pool with channels
	fmt.Println("\n10. Worker pool:")
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go poolWorker(i, jobs, results, &wg)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}

	// 11. Fan-in pattern
	fmt.Println("\n11. Fan-in pattern:")
	input1 := make(chan int)
	input2 := make(chan int)
	output := make(chan int)
	
	// Producers
	go func() {
		defer close(input1)
		for i := 1; i <= 3; i++ {
			input1 <- i
		}
	}()
	
	go func() {
		defer close(input2)
		for i := 4; i <= 6; i++ {
			input2 <- i
		}
	}()
	
	// Fan-in
	go func() {
		defer close(output)
		for {
			select {
			case val, ok := <-input1:
				if !ok {
					input1 = nil
				} else {
					output <- val
				}
			case val, ok := <-input2:
				if !ok {
					input2 = nil
				} else {
					output <- val
				}
			}
			if input1 == nil && input2 == nil {
				break
			}
		}
	}()
	
	for val := range output {
		fmt.Printf("Fan-in result: %d\n", val)
	}

	// 12. Fan-out pattern
	fmt.Println("\n12. Fan-out pattern:")
	input := make(chan int)
	outputs := []chan int{make(chan int), make(chan int), make(chan int)}
	
	// Distribute to multiple workers
	for i, output := range outputs {
		go func(id int, out chan int) {
			for val := range input {
				result := val * id
				out <- result
			}
			close(out)
		}(i+1, output)
	}
	
	// Send input data
	go func() {
		defer close(input)
		for i := 1; i <= 3; i++ {
			input <- i
		}
	}()
	
	// Collect from all outputs
	for i, output := range outputs {
		fmt.Printf("Worker %d results: ", i+1)
		for result := range output {
			fmt.Printf("%d ", result)
		}
		fmt.Println()
	}

	// 13. Channel for signaling
	fmt.Println("\n13. Channel for signaling:")
	done := make(chan bool)
	
	go func() {
		fmt.Println("Worker started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Worker finished")
		done <- true
	}()
	
	fmt.Println("Waiting for worker...")
	<-done
	fmt.Println("Worker completed")

	// 14. Channel for cancellation
	fmt.Println("\n14. Channel for cancellation:")
	stop := make(chan bool)
	
	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				fmt.Println("Working...")
			case <-stop:
				fmt.Println("Stopping work...")
				return
			}
		}
	}()
	
	time.Sleep(200 * time.Millisecond)
	stop <- true
	time.Sleep(50 * time.Millisecond)

	// 15. Channel with struct for complex data
	fmt.Println("\n15. Channel with struct data:")
	type Message struct {
		ID      int
		Content string
		Time    time.Time
	}
	
	msgCh := make(chan Message, 3)
	
	go func() {
		defer close(msgCh)
		messages := []Message{
			{ID: 1, Content: "First message", Time: time.Now()},
			{ID: 2, Content: "Second message", Time: time.Now()},
			{ID: 3, Content: "Third message", Time: time.Now()},
		}
		
		for _, msg := range messages {
			msgCh <- msg
		}
	}()
	
	for msg := range msgCh {
		fmt.Printf("Message %d: %s at %v\n", msg.ID, msg.Content, msg.Time.Format("15:04:05"))
	}

	// 16. Channel pipeline
	fmt.Println("\n16. Channel pipeline:")
	// Stage 1: Generate numbers
	numbers := make(chan int)
	go func() {
		defer close(numbers)
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
	}()
	
	// Stage 2: Square numbers
	squares := make(chan int)
	go func() {
		defer close(squares)
		for num := range numbers {
			squares <- num * num
		}
	}()
	
	// Stage 3: Add 10
	results := make(chan int)
	go func() {
		defer close(results)
		for square := range squares {
			results <- square + 10
		}
	}()
	
	// Collect final results
	for result := range results {
		fmt.Printf("Pipeline result: %d\n", result)
	}

	// 17. Channel for rate limiting
	fmt.Println("\n17. Rate limiting with channel:")
	requests := make(chan int, 5)
	
	// Simulate incoming requests
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	
	// Rate limiter
	limiter := time.NewTicker(100 * time.Millisecond)
	defer limiter.Stop()
	
	for req := range requests {
		<-limiter.C // Wait for ticker
		fmt.Printf("Processing request %d at %v\n", req, time.Now().Format("15:04:05.000"))
	}

	// 18. Channel with timeout pattern
	fmt.Println("\n18. Timeout pattern:")
	processWithTimeout := func() (string, error) {
		ch := make(chan string)
		
		go func() {
			// Simulate work that might take time
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "Work completed"
		}()
		
		select {
		case result := <-ch:
			return result, nil
		case <-time.After(100 * time.Millisecond):
			return "", fmt.Errorf("operation timed out")
		}
	}
	
	for i := 1; i <= 3; i++ {
		result, err := processWithTimeout()
		if err != nil {
			fmt.Printf("Attempt %d: %v\n", i, err)
		} else {
			fmt.Printf("Attempt %d: %s\n", i, result)
		}
	}

	fmt.Println("All channel examples completed!")
}

// Helper functions

func sendData(ch chan<- string, messages ...string) {
	defer close(ch)
	for _, msg := range messages {
		ch <- msg
	}
}

func poolWorker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		result := job * 2
		results <- result
		fmt.Printf("Worker %d completed job %d\n", id, job)
	}
}
