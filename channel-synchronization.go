package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Channel Synchronization Examples ===")

	// 1. Basic synchronization with channels
	fmt.Println("\n1. Basic synchronization:")
	done := make(chan bool)
	
	go func() {
		fmt.Println("Worker started")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Worker finished")
		done <- true
	}()
	
	fmt.Println("Main waiting for worker...")
	<-done
	fmt.Println("Main received completion signal")

	// 2. Synchronizing multiple goroutines
	fmt.Println("\n2. Multiple goroutine synchronization:")
	numWorkers := 3
	doneChan := make(chan bool, numWorkers)
	
	for i := 1; i <= numWorkers; i++ {
		go func(id int) {
			fmt.Printf("Worker %d started\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d finished\n", id)
			doneChan <- true
		}(i)
	}
	
	// Wait for all workers
	for i := 0; i < numWorkers; i++ {
		<-doneChan
	}
	fmt.Println("All workers completed")

	// 3. Using channels for coordination
	fmt.Println("\n3. Coordination pattern:")
	startChan := make(chan struct{})
	workerDone := make(chan struct{})
	
	go func() {
		<-startChan // Wait for start signal
		fmt.Println("Worker received start signal")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Worker completed work")
		workerDone <- struct{}{}
	}()
	
	fmt.Println("Main sending start signal")
	close(startChan) // Broadcast start signal
	
	<-workerDone
	fmt.Println("Main received completion signal")

	// 4. Pipeline synchronization
	fmt.Println("\n4. Pipeline synchronization:")
	stage1 := make(chan int, 3)
	stage2 := make(chan int, 3)
	stage3 := make(chan int, 3)
	
	// Stage 1: Generate numbers
	go func() {
		defer close(stage1)
		for i := 1; i <= 5; i++ {
			stage1 <- i
			fmt.Printf("Stage 1 generated: %d\n", i)
		}
	}()
	
	// Stage 2: Process numbers
	go func() {
		defer close(stage2)
		for num := range stage1 {
			result := num * 2
			stage2 <- result
			fmt.Printf("Stage 2 processed: %d -> %d\n", num, result)
		}
	}()
	
	// Stage 3: Final processing
	go func() {
		defer close(stage3)
		for num := range stage2 {
			result := num + 10
			stage3 <- result
			fmt.Printf("Stage 3 processed: %d -> %d\n", num, result)
		}
	}()
	
	// Collect results
	for result := range stage3 {
		fmt.Printf("Final result: %d\n", result)
	}

	// 5. Fan-in synchronization
	fmt.Println("\n5. Fan-in synchronization:")
	input := make(chan int)
	output := make(chan int)
	
	// Multiple workers processing same input
	worker := func(id int, input <-chan int, output chan<- int) {
		for num := range input {
			result := num * id
			output <- result
			fmt.Printf("Worker %d processed: %d -> %d\n", id, num, result)
		}
	}
	
	// Start 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, input, output)
	}
	
	// Send input data
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()
	
	// Collect results
	for i := 0; i < 15; i++ { // 5 inputs * 3 workers
		result := <-output
		fmt.Printf("Received: %d\n", result)
	}

	// 6. Synchronization with struct
	fmt.Println("\n6. Struct-based synchronization:")
	type Worker struct {
		id     int
		input  chan int
		output chan int
		done   chan bool
	}
	
	newWorker := func(id int) *Worker {
		w := &Worker{
			id:     id,
			input:  make(chan int),
			output: make(chan int),
			done:   make(chan bool),
		}
		
		go func() {
			for num := range w.input {
				result := num * w.id
				w.output <- result
				fmt.Printf("Worker %d: %d -> %d\n", w.id, num, result)
			}
			w.done <- true
		}()
		
		return w
	}
	
	workers := []*Worker{newWorker(2), newWorker(3)}
	
	// Send work to all workers
	go func() {
		for _, w := range workers {
			go func(worker *Worker) {
				defer close(worker.input)
				for i := 1; i <= 3; i++ {
					worker.input <- i
				}
			}(w)
		}
	}()
	
	// Wait for all workers to complete
	for _, w := range workers {
		<-w.done
		fmt.Printf("Worker %d completed\n", w.id)
	}

	// 7. Synchronization with timeout
	fmt.Println("\n7. Synchronization with timeout:")
	syncChan := make(chan bool)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		syncChan <- true
	}()
	
	select {
	case <-syncChan:
		fmt.Println("Operation completed successfully")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Operation timed out")
	}

	// 8. Barrier synchronization pattern
	fmt.Println("\n8. Barrier pattern:")
	barrier := make(chan struct{})
	workerCount := 3
	
	for i := 1; i <= workerCount; i++ {
		go func(id int) {
			fmt.Printf("Worker %d phase 1\n", id)
			time.Sleep(time.Duration(id*50) * time.Millisecond)
			
			// Wait at barrier
			<-barrier
			fmt.Printf("Worker %d phase 2\n", id)
		}(i)
	}
	
	// Open barrier after all workers reach it
	time.Sleep(200 * time.Millisecond)
	close(barrier)
	time.Sleep(100 * time.Millisecond)

	// 9. Synchronization with WaitGroup and channels
	fmt.Println("\n9. WaitGroup + channels:")
	var wg sync.WaitGroup
	results := make(chan int, 10)
	
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id*50) * time.Millisecond)
			result := id * 10
			results <- result
			fmt.Printf("Worker %d produced: %d\n", id, result)
		}(i)
	}
	
	// Close results channel when all workers done
	go func() {
		wg.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Printf("Collected: %d\n", result)
	}

	// 10. Producer-consumer synchronization
	fmt.Println("\n10. Producer-consumer:")
	producerChan := make(chan int, 5)
	consumerDone := make(chan bool)
	
	// Producer
	go func() {
		defer close(producerChan)
		for i := 1; i <= 5; i++ {
			producerChan <- i
			fmt.Printf("Produced: %d\n", i)
		}
		fmt.Println("Producer finished")
	}()
	
	// Consumer
	go func() {
		for item := range producerChan {
			fmt.Printf("Consumed: %d\n", item)
			time.Sleep(50 * time.Millisecond)
		}
		consumerDone <- true
	}()
	
	<-consumerDone
	fmt.Println("Consumer finished")

	// 11. Synchronization for resource access
	fmt.Println("\n11. Resource access synchronization:")
	resourceChan := make(chan struct{}, 1) // Semaphore
	
	accessResource := func(id int) {
		resourceChan <- struct{}{} // Acquire
		defer func() { <-resourceChan }() // Release
		
		fmt.Printf("Goroutine %d accessing resource\n", id)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("Goroutine %d finished accessing resource\n", id)
	}
	
	// Multiple goroutines accessing shared resource
	for i := 1; i <= 3; i++ {
		go accessResource(i)
	}
	
	time.Sleep(500 * time.Millisecond)

	// 12. Synchronization with context signaling
	fmt.Println("\n12. Context signaling:")
	stopChan := make(chan struct{})
	dataChan := make(chan int)
	
	// Data producer
	go func() {
		i := 1
		for {
			select {
			case <-stopChan:
				fmt.Println("Producer stopping...")
				return
			case dataChan <- i:
				fmt.Printf("Produced: %d\n", i)
				i++
				time.Sleep(50 * time.Millisecond)
			}
		}
	}()
	
	// Consumer with stop condition
	go func() {
		count := 0
		for item := range dataChan {
			fmt.Printf("Consumed: %d\n", item)
			count++
			if count >= 5 {
				close(stopChan) // Signal producer to stop
				break
			}
		}
		close(dataChan)
	}()
	
	time.Sleep(400 * time.Millisecond)

	fmt.Println("All synchronization examples completed!")
}
