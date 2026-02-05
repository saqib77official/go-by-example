package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== WaitGroups Examples ===")

	// 1. Basic WaitGroup
	fmt.Println("\n1. Basic WaitGroup:")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d started\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}
	
	fmt.Println("Waiting for goroutines...")
	wg.Wait()
	fmt.Println("All goroutines completed")

	// 2. WaitGroup with return values
	fmt.Println("\n2. WaitGroup with return values:")
	var wg2 sync.WaitGroup
	results := make(chan int, 3)
	
	for i := 1; i <= 3; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			time.Sleep(50 * time.Millisecond)
			results <- id * 10
		}(i)
	}
	
	// Wait for completion
	go func() {
		wg2.Wait()
		close(results)
	}()
	
	// Collect results
	for result := range results {
		fmt.Printf("Received result: %d\n", result)
	}

	// 3. Nested WaitGroups
	fmt.Println("\n3. Nested WaitGroups:")
	var outerWG sync.WaitGroup
	
	for i := 1; i <= 2; i++ {
		outerWG.Add(1)
		go func(outerID int) {
			defer outerWG.Done()
			
			var innerWG sync.WaitGroup
			fmt.Printf("Outer goroutine %d started\n", outerID)
			
			for j := 1; j <= 2; j++ {
				innerWG.Add(1)
				go func(innerID int) {
					defer innerWG.Done()
					fmt.Printf("  Inner goroutine %d.%d started\n", outerID, innerID)
					time.Sleep(50 * time.Millisecond)
					fmt.Printf("  Inner goroutine %d.%d finished\n", outerID, innerID)
				}(j)
			}
			
			innerWG.Wait()
			fmt.Printf("Outer goroutine %d finished\n", outerID)
		}(i)
	}
	
	outerWG.Wait()
	fmt.Println("All nested goroutines completed")

	// 4. WaitGroup with error handling
	fmt.Println("\n4. WaitGroup with error handling:")
	var wg3 sync.WaitGroup
	errors := make(chan error, 3)
	
	for i := 1; i <= 3; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			time.Sleep(50 * time.Millisecond)
			
			// Simulate error for goroutine 2
			if id == 2 {
				errors <- fmt.Errorf("error in goroutine %d", id)
			} else {
				errors <- nil
			}
		}(i)
	}
	
	// Wait and collect errors
	go func() {
		wg3.Wait()
		close(errors)
	}()
	
	for err := range errors {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println("Success")
		}
	}

	// 5. WaitGroup with timeout
	fmt.Println("\n5. WaitGroup with timeout:")
	var wg4 sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("Goroutine %d completed\n", id)
		}(i)
	}
	
	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg4.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		fmt.Println("All goroutines completed normally")
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout reached")
	}

	// 6. WaitGroup with counter
	fmt.Println("\n6. WaitGroup with counter:")
	var wg5 sync.WaitGroup
	counter := 0
	var mu sync.Mutex
	
	for i := 1; i <= 5; i++ {
		wg5.Add(1)
		go func(id int) {
			defer wg5.Done()
			time.Sleep(50 * time.Millisecond)
			
			mu.Lock()
			counter++
			fmt.Printf("Goroutine %d completed (total: %d)\n", id, counter)
			mu.Unlock()
		}(i)
	}
	
	wg5.Wait()
	fmt.Printf("Final counter: %d\n", counter)

	// 7. WaitGroup with dynamic addition
	fmt.Println("\n7. WaitGroup with dynamic addition:")
	var wg6 sync.WaitGroup
	work := make(chan int, 10)
	
	// Worker
	worker := func() {
		for job := range work {
			wg6.Add(1)
			go func(j int) {
				defer wg6.Done()
				time.Sleep(50 * time.Millisecond)
				fmt.Printf("Processed job %d\n", j)
			}(job)
		}
	}
	
	// Start 3 workers
	for i := 0; i < 3; i++ {
		go worker()
	}
	
	// Send work
	for j := 1; j <= 9; j++ {
		work <- j
	}
	close(work)
	
	wg6.Wait()
	fmt.Println("All jobs processed")

	// 8. WaitGroup with pipeline
	fmt.Println("\n8. WaitGroup with pipeline:")
	var wg7 sync.WaitGroup
	
	stage1 := make(chan int, 10)
	stage2 := make(chan int, 10)
	stage3 := make(chan int, 10)
	
	// Stage 1
	wg7.Add(1)
	go func() {
		defer wg7.Done()
		defer close(stage1)
		for i := 1; i <= 5; i++ {
			stage1 <- i
			fmt.Printf("Stage 1 produced: %d\n", i)
		}
	}()
	
	// Stage 2
	wg7.Add(1)
	go func() {
		defer wg7.Done()
		defer close(stage2)
		for num := range stage1 {
			result := num * 2
			stage2 <- result
			fmt.Printf("Stage 2 processed: %d -> %d\n", num, result)
		}
	}()
	
	// Stage 3
	wg7.Add(1)
	go func() {
		defer wg7.Done()
		defer close(stage3)
		for num := range stage2 {
			result := num + 10
			stage3 <- result
			fmt.Printf("Stage 3 processed: %d -> %d\n", num, result)
		}
	}()
	
	// Wait for all stages and collect results
	go func() {
		wg7.Wait()
	}()
	
	for result := range stage3 {
		fmt.Printf("Final result: %d\n", result)
	}

	// 9. WaitGroup with resource pool
	fmt.Println("\n9. WaitGroup with resource pool:")
	var wg8 sync.WaitGroup
	resources := make(chan string, 3)
	
	// Initialize resources
	go func() {
		resources <- "resource-1"
		resources <- "resource-2"
		resources <- "resource-3"
	}()
	
	// Use resources
	for i := 1; i <= 5; i++ {
		wg8.Add(1)
		go func(id int) {
			defer wg8.Done()
			
			// Acquire resource
			resource := <-resources
			fmt.Printf("Goroutine %d acquired %s\n", id, resource)
			
			time.Sleep(50 * time.Millisecond)
			
			// Release resource
			resources <- resource
			fmt.Printf("Goroutine %d released %s\n", id, resource)
		}(i)
	}
	
	wg8.Wait()
	fmt.Println("All resource operations completed")

	// 10. WaitGroup with fan-out/fan-in
	fmt.Println("\n10. WaitGroup with fan-out/fan-in:")
	var wg9 sync.WaitGroup
	
	input := make(chan int, 20)
	outputs := []chan int{
		make(chan int, 10),
		make(chan int, 10),
		make(chan int, 10),
	}
	
	// Distributor
	wg9.Add(1)
	go func() {
		defer wg9.Done()
		defer close(input)
		for i := 1; i <= 9; i++ {
			input <- i
		}
	}()
	
	// Workers (fan-out)
	for i, output := range outputs {
		wg9.Add(1)
		go func(id int, out chan int) {
			defer wg9.Done()
			defer close(out)
			for num := range input {
				result := num * (id + 1)
				out <- result
				fmt.Printf("Worker %d: %d -> %d\n", id+1, num, result)
			}
		}(i, output)
	}
	
	// Collector (fan-in)
	wg9.Add(1)
	go func() {
		defer wg9.Done()
		var wgCollect sync.WaitGroup
		
		for i, output := range outputs {
			wgCollect.Add(1)
			go func(id int, ch <-chan int) {
				defer wgCollect.Done()
				for result := range ch {
					fmt.Printf("Collected from worker %d: %d\n", id+1, result)
				}
			}(i, output)
		}
		
		wgCollect.Wait()
	}()
	
	wg9.Wait()
	fmt.Println("Fan-out/fan-in completed")

	// 11. WaitGroup with batch processing
	fmt.Println("\n11. WaitGroup with batch processing:")
	var wg10 sync.WaitGroup
	
	batches := [][]int{
		{1, 2, 3},
		{4, 5},
		{6, 7, 8, 9},
	}
	
	for i, batch := range batches {
		wg10.Add(1)
		go func(batchID int, b []int) {
			defer wg10.Done()
			fmt.Printf("Processing batch %d: %v\n", batchID+1, b)
			
			sum := 0
			for _, num := range b {
				sum += num
			}
			
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Batch %d sum: %d\n", batchID+1, sum)
		}(i, batch)
	}
	
	wg10.Wait()
	fmt.Println("All batches processed")

	// 12. WaitGroup with progress tracking
	fmt.Println("\n12. WaitGroup with progress tracking:")
	var wg11 sync.WaitGroup
	totalTasks := 10
	completed := make(chan int, totalTasks)
	
	for i := 1; i <= totalTasks; i++ {
		wg11.Add(1)
		go func(id int) {
			defer wg11.Done()
			time.Sleep(time.Duration(id*50) * time.Millisecond)
			completed <- id
		}(i)
	}
	
	// Progress tracker
	go func() {
		count := 0
		for id := range completed {
			count++
			percentage := float64(count) / float64(totalTasks) * 100
			fmt.Printf("Task %d completed (%.1f%%)\n", id, percentage)
		}
	}()
	
	wg11.Wait()
	fmt.Println("All tasks completed")

	// 13. WaitGroup with cancellation
	fmt.Println("\n13. WaitGroup with cancellation:")
	var wg12 sync.WaitGroup
	stop := make(chan struct{})
	
	for i := 1; i <= 5; i++ {
		wg12.Add(1)
		go func(id int) {
			defer wg12.Done()
			for j := 1; j <= 10; j++ {
				select {
				case <-stop:
					fmt.Printf("Goroutine %d stopped\n", id)
					return
				default:
					fmt.Printf("Goroutine %d working on step %d\n", id, j)
					time.Sleep(50 * time.Millisecond)
				}
			}
		}(i)
	}
	
	// Stop after 300ms
	time.Sleep(300 * time.Millisecond)
	close(stop)
	
	wg12.Wait()
	fmt.Println("All goroutines stopped")

	// 14. WaitGroup with retry mechanism
	fmt.Println("\n14. WaitGroup with retry mechanism:")
	var wg13 sync.WaitGroup
	results := make(chan string, 15)
	
	for i := 1; i <= 5; i++ {
		wg13.Add(1)
		go func(id int) {
			defer wg13.Done()
			
			var err error
			for attempt := 1; attempt <= 3; attempt++ {
				// Simulate work
				time.Sleep(30 * time.Millisecond)
				
				// Simulate success on attempt 2 or 3
				if attempt >= 2 {
					err = nil
					break
				}
				err = fmt.Errorf("attempt %d failed", attempt)
			}
			
			if err != nil {
				results <- fmt.Sprintf("Task %d failed: %v", id, err)
			} else {
				results <- fmt.Sprintf("Task %d succeeded", id)
			}
		}(i)
	}
	
	// Wait and collect
	go func() {
		wg13.Wait()
		close(results)
	}()
	
	for result := range results {
		fmt.Printf("Result: %s\n", result)
	}

	// 15. WaitGroup with resource cleanup
	fmt.Println("\n15. WaitGroup with resource cleanup:")
	var wg14 sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		wg14.Add(1)
		go func(id int) {
			defer wg14.Done()
			
			// Acquire resource
			resource := fmt.Sprintf("resource-%d", id)
			fmt.Printf("Goroutine %d acquired %s\n", id, resource)
			
			// Simulate work
			time.Sleep(100 * time.Millisecond)
			
			// Cleanup
			fmt.Printf("Goroutine %d cleaning up %s\n", id, resource)
		}(i)
	}
	
	wg14.Wait()
	fmt.Println("All resources cleaned up")

	fmt.Println("All WaitGroup examples completed!")
}
