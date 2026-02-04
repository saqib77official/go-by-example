package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== Goroutines Examples ===")

	// 1. Basic goroutine
	fmt.Println("\n1. Basic goroutine:")
	go sayHello("Goroutine 1")
	go sayHello("Goroutine 2")
	
	// Wait a bit to see the output
	time.Sleep(100 * time.Millisecond)

	// 2. Anonymous function goroutine
	fmt.Println("\n2. Anonymous function goroutine:")
	go func() {
		fmt.Println("Anonymous goroutine running")
	}()
	
	time.Sleep(50 * time.Millisecond)

	// 3. Goroutine with parameters
	fmt.Println("\n3. Goroutine with parameters:")
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Worker %d started\n", id)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Printf("Worker %d finished\n", id)
		}(i)
	}
	
	time.Sleep(200 * time.Millisecond)

	// 4. Using WaitGroup for synchronization
	fmt.Println("\n4. WaitGroup synchronization:")
	var wg sync.WaitGroup
	
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Task %d started\n", id)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Printf("Task %d completed\n", id)
		}(i)
	}
	
	fmt.Println("Waiting for all tasks to complete...")
	wg.Wait()
	fmt.Println("All tasks completed")

	// 5. Common pitfall: loop variable capture
	fmt.Println("\n5. Loop variable capture (correct way):")
	var wg2 sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			fmt.Printf("Correct capture: %d\n", id)
		}(i) // Pass i as parameter
	}
	
	wg2.Wait()

	// 6. Goroutine with channels
	fmt.Println("\n6. Goroutine with channels:")
	ch := make(chan string)
	
	go func() {
		ch <- "Message from goroutine"
	}()
	
	message := <-ch
	fmt.Printf("Received: %s\n", message)

	// 7. Multiple goroutines sending to one channel
	fmt.Println("\n7. Multiple producers:")
	ch2 := make(chan int)
	var wg3 sync.WaitGroup
	
	// Producer goroutines
	for i := 1; i <= 3; i++ {
		wg3.Add(1)
		go func(producerID int) {
			defer wg3.Done()
			for j := 1; j <= 3; j++ {
				value := producerID*10 + j
				ch2 <- value
				fmt.Printf("Producer %d sent: %d\n", producerID, value)
			}
		}(i)
	}
	
	// Consumer goroutine
	go func() {
		wg3.Wait()
		close(ch2)
	}()
	
	// Receive values
	for value := range ch2 {
		fmt.Printf("Consumer received: %d\n", value)
	}

	// 8. Worker pool pattern
	fmt.Println("\n8. Worker pool:")
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start workers
	var wg4 sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg4.Add(1)
		go worker(i, jobs, results, &wg4)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Printf("Sent job %d\n", j)
	}
	close(jobs)
	
	// Wait for workers to finish
	go func() {
		wg4.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}

	// 9. Atomic operations
	fmt.Println("\n9. Atomic operations:")
	var counter int64
	
	// Start multiple goroutines that increment counter
	var wg5 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg5.Add(1)
		go func() {
			defer wg5.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	
	wg5.Wait()
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))

	// 10. Mutex for protecting shared state
	fmt.Println("\n10. Mutex for shared state:")
	var mu sync.Mutex
	balance := 1000
	
	var wg6 sync.WaitGroup
	
	// Depositors
	for i := 1; i <= 5; i++ {
		wg6.Add(1)
		go func(amount int) {
			defer wg6.Done()
			mu.Lock()
			balance += amount
			fmt.Printf("Deposited: %d, New balance: %d\n", amount, balance)
			mu.Unlock()
		}(100)
	}
	
	// Withdrawers
	for i := 1; i <= 3; i++ {
		wg6.Add(1)
		go func(amount int) {
			defer wg6.Done()
			mu.Lock()
			if balance >= amount {
				balance -= amount
				fmt.Printf("Withdrew: %d, New balance: %d\n", amount, balance)
			} else {
				fmt.Printf("Insufficient funds for withdrawal: %d\n", amount)
			}
			mu.Unlock()
		}(200)
	}
	
	wg6.Wait()
	fmt.Printf("Final balance: %d\n", balance)

	// 11. Select statement with channels
	fmt.Println("\n11. Select statement:")
	ch3 := make(chan string)
	ch4 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch3 <- "From channel 3"
	}()
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch4 <- "From channel 4"
	}()
	
	select {
	case msg1 := <-ch3:
		fmt.Printf("Received from ch3: %s\n", msg1)
	case msg2 := <-ch4:
		fmt.Printf("Received from ch4: %s\n", msg2)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Timeout occurred")
	}

	// 12. Fan-out/Fan-in pattern
	fmt.Println("\n12. Fan-out/Fan-in pattern:")
	input := make(chan int)
	
	// Fan-out: distribute work to multiple workers
	output1 := make(chan int)
	output2 := make(chan int)
	
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()
	
	go squareWorker(input, output1)
	go squareWorker(input, output2)
	
	// Fan-in: collect results
	go func() {
		var wg7 sync.WaitGroup
		wg7.Add(2)
		
		go func() {
			defer wg7.Done()
			for val := range output1 {
				fmt.Printf("Worker 1 result: %d\n", val)
			}
		}()
		
		go func() {
			defer wg7.Done()
			for val := range output2 {
				fmt.Printf("Worker 2 result: %d\n", val)
			}
		}()
		
		wg7.Wait()
	}()
	
	time.Sleep(100 * time.Millisecond)

	// 13. Goroutine leak prevention
	fmt.Println("\n13. Goroutine leak prevention:")
	processWithTimeout := func() error {
		ch := make(chan string)
		
		go func() {
			time.Sleep(200 * time.Millisecond) // Simulate work
			ch <- "result"
		}()
		
		select {
		case result := <-ch:
			fmt.Printf("Work completed: %s\n", result)
			return nil
		case <-time.After(100 * time.Millisecond):
			fmt.Println("Work timed out")
			return fmt.Errorf("operation timed out")
		}
	}
	
	err := processWithTimeout()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// 14. Using once for initialization
	fmt.Println("\n14. Using sync.Once:")
	var once sync.Once
	var config map[string]string
	
	loadConfig := func() {
		fmt.Println("Loading configuration...")
		config = map[string]string{
			"host": "localhost",
			"port": "8080",
		}
		time.Sleep(50 * time.Millisecond) // Simulate loading time
	}
	
	var wg8 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg8.Add(1)
		go func(id int) {
			defer wg8.Done()
			once.Do(loadConfig) // Only executed once
			fmt.Printf("Goroutine %d: config loaded\n", id)
		}(i)
	}
	
	wg8.Wait()
	fmt.Printf("Config: %v\n", config)

	// 15. Goroutine for background tasks
	fmt.Println("\n15. Background task with goroutine:")
	stop := make(chan bool)
	
	// Background task
	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				fmt.Println("Background task running...")
			case <-stop:
				fmt.Println("Background task stopped")
				return
			}
		}
	}()
	
	// Let it run for a while
	time.Sleep(1600 * time.Millisecond)
	
	// Stop the background task
	stop <- true
	time.Sleep(100 * time.Millisecond)

	fmt.Println("All examples completed!")
}

// Helper functions

func sayHello(message string) {
	fmt.Printf("Hello from %s!\n", message)
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		result := job * 2
		results <- result
		fmt.Printf("Worker %d completed job %d with result %d\n", id, job, result)
	}
}

func squareWorker(input <-chan int, output chan<- int) {
	for num := range input {
		result := num * num
		output <- result
	}
}
