package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Worker Pools Examples ===")

	// 1. Basic worker pool
	fmt.Println("\n1. Basic worker pool:")
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	
	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Printf("Sent job %d\n", j)
	}
	close(jobs)
	
	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Received result %d\n", result)
	}

	// 2. Worker pool with WaitGroup
	fmt.Println("\n2. Worker pool with WaitGroup:")
	var wg sync.WaitGroup
	jobs2 := make(chan int, 20)
	results2 := make(chan int, 20)
	
	// Start workers
	for w := 1; w <= 4; w++ {
		wg.Add(1)
		go workerWithWaitGroup(w, jobs2, results2, &wg)
	}
	
	// Send jobs
	for j := 1; j <= 8; j++ {
		jobs2 <- j
	}
	close(jobs2)
	
	// Start result collector
	go func() {
		wg.Wait()
		close(results2)
	}()
	
	// Collect results
	for result := range results2 {
		fmt.Printf("Result: %d\n", result)
	}

	// 3. Dynamic worker pool
	fmt.Println("\n3. Dynamic worker pool:")
	
	dynamicPool := func(numWorkers, numJobs int) {
		jobs := make(chan int, numJobs)
		results := make(chan int, numJobs)
		var wg sync.WaitGroup
		
		// Start workers
		for i := 1; i <= numWorkers; i++ {
			wg.Add(1)
			go workerWithWaitGroup(i, jobs, results, &wg)
		}
		
		// Send jobs
		go func() {
			defer close(jobs)
			for j := 1; j <= numJobs; j++ {
				jobs <- j
			}
		}()
		
		// Wait for completion
		go func() {
			wg.Wait()
			close(results)
		}()
		
		// Collect results
		var allResults []int
		for result := range results {
			allResults = append(allResults, result)
		}
		
		fmt.Printf("Processed %d jobs with %d workers\n", len(allResults), numWorkers)
		fmt.Printf("Results: %v\n", allResults)
	}
	
	dynamicPool(3, 6)

	// 4. Worker pool with timeout
	fmt.Println("\n4. Worker pool with timeout:")
	
	poolWithTimeout := func() {
		jobs := make(chan int, 10)
		results := make(chan string, 10)
		
		// Start workers with timeout
		for w := 1; w <= 3; w++ {
			go timeoutWorker(w, jobs, results)
		}
		
		// Send jobs
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs)
		
		// Collect results
		for i := 0; i < 5; i++ {
			result := <-results
			fmt.Printf("Result: %s\n", result)
		}
	}
	
	poolWithTimeout()

	// 5. Worker pool with load balancing
	fmt.Println("\n5. Worker pool with load balancing:")
	
	loadBalancedPool := func() {
		jobs := make(chan int, 20)
		results := make(chan string, 20)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 4; w++ {
			wg.Add(1)
			go loadBalancedWorker(w, jobs, results, &wg)
		}
		
		// Send jobs with different complexities
		go func() {
			defer close(jobs)
			jobComplexities := []int{1, 3, 1, 2, 3, 2, 1, 3}
			for _, complexity := range jobComplexities {
				jobs <- complexity
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		for result := range results {
			fmt.Printf("Completed: %s\n", result)
		}
	}
	
	loadBalancedPool()

	// 6. Worker pool with priority queue
	fmt.Println("\n6. Worker pool with priority queue:")
	
	priorityPool := func() {
		type Job struct {
			priority int
			id       int
		}
		
		jobs := make(chan Job, 20)
		results := make(chan string, 20)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 2; w++ {
			wg.Add(1)
			go priorityWorker(w, jobs, results, &wg)
		}
		
		// Send priority jobs
		go func() {
			defer close(jobs)
			priorityJobs := []Job{
				{priority: 1, id: 1},
				{priority: 3, id: 2},
				{priority: 2, id: 3},
				{priority: 3, id: 4},
				{priority: 1, id: 5},
			}
			
			for _, job := range priorityJobs {
				jobs <- job
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		for result := range results {
			fmt.Printf("Completed: %s\n", result)
		}
	}
	
	priorityPool()

	// 7. Worker pool with retry mechanism
	fmt.Println("\n7. Worker pool with retry mechanism:")
	
	retryPool := func() {
		jobs := make(chan int, 20)
		results := make(chan string, 20)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 2; w++ {
			wg.Add(1)
			go retryWorker(w, jobs, results, &wg)
		}
		
		// Send jobs
		go func() {
			defer close(jobs)
			for j := 1; j <= 6; j++ {
				jobs <- j
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		for result := range results {
			fmt.Printf("Result: %s\n", result)
		}
	}
	
	retryPool()

	// 8. Worker pool with graceful shutdown
	fmt.Println("\n8. Worker pool with graceful shutdown:")
	
	gracefulPool := func() {
		jobs := make(chan int, 20)
		results := make(chan string, 20)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 3; w++ {
			wg.Add(1)
			go gracefulWorker(w, jobs, results, &wg)
		}
		
		// Send jobs
		go func() {
			defer close(jobs)
			for j := 1; j <= 8; j++ {
				jobs <- j
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		for result := range results {
			fmt.Printf("Result: %s\n", result)
		}
	}
	
	gracefulPool()

	// 9. Worker pool with statistics
	fmt.Println("\n9. Worker pool with statistics:")
	
	poolWithStats := func() {
		type Stats struct {
			jobsProcessed int
			totalTime     time.Duration
		}
		
		jobs := make(chan int, 20)
		results := make(chan Stats, 20)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 3; w++ {
			wg.Add(1)
			go statsWorker(w, jobs, results, &wg)
		}
		
		// Send jobs
		go func() {
			defer close(jobs)
			for j := 1; j <= 9; j++ {
				jobs <- j
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		// Aggregate statistics
		var totalJobs int
		var totalTime time.Duration
		
		for stat := range results {
			totalJobs += stat.jobsProcessed
			totalTime += stat.totalTime
		}
		
		fmt.Printf("Total jobs processed: %d\n", totalJobs)
		fmt.Printf("Total time: %v\n", totalTime)
		if totalJobs > 0 {
			avgTime := totalTime / time.Duration(totalJobs)
			fmt.Printf("Average time per job: %v\n", avgTime)
		}
	}
	
	poolWithStats()

	// 10. Worker pool with batch processing
	fmt.Println("\n10. Worker pool with batch processing:")
	
	batchPool := func() {
		jobs := make(chan []int, 10)
		results := make(chan string, 10)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 2; w++ {
			wg.Add(1)
			go batchWorker(w, jobs, results, &wg)
		}
		
		// Send batch jobs
		go func() {
			defer close(jobs)
			batches := [][]int{
				{1, 2, 3},
				{4, 5},
				{6, 7, 8, 9},
			}
			
			for _, batch := range batches {
				jobs <- batch
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		for result := range results {
			fmt.Printf("Batch result: %s\n", result)
		}
	}
	
	batchPool()

	// 11. Worker pool with backpressure
	fmt.Println("\n11. Worker pool with backpressure:")
	
	backpressurePool := func() {
		jobs := make(chan int, 5) // Small buffer for backpressure
		results := make(chan string, 20)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 2; w++ {
			wg.Add(1)
			go backpressureWorker(w, jobs, results, &wg)
		}
		
		// Send jobs (may block due to backpressure)
		go func() {
			defer close(jobs)
			for j := 1; j <= 10; j++ {
				select {
				case jobs <- j:
					fmt.Printf("Sent job %d\n", j)
				case <-time.After(100 * time.Millisecond):
					fmt.Printf("Job %d dropped (backpressure)\n", j)
				}
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		for result := range results {
			fmt.Printf("Result: %s\n", result)
		}
	}
	
	backpressurePool()

	// 12. Worker pool with circuit breaker
	fmt.Println("\n12. Worker pool with circuit breaker:")
	
	circuitBreakerPool := func() {
		type CircuitBreaker struct {
			failures    int
			maxFailures int
			open        bool
		}
		
		cb := &CircuitBreaker{maxFailures: 3}
		
		jobs := make(chan int, 20)
		results := make(chan string, 20)
		var wg sync.WaitGroup
		
		// Start workers
		for w := 1; w <= 2; w++ {
			wg.Add(1)
			go circuitBreakerWorker(w, jobs, results, &wg, cb)
		}
		
		// Send jobs
		go func() {
			defer close(jobs)
			for j := 1; j <= 10; j++ {
				jobs <- j
			}
		}()
		
		// Wait and collect
		go func() {
			wg.Wait()
			close(results)
		}()
		
		for result := range results {
			fmt.Printf("Result: %s\n", result)
		}
	}
	
	circuitBreakerPool()

	fmt.Println("All worker pool examples completed!")
}

// Worker functions

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(100 * time.Millisecond) // Simulate work
		results <- j * 2
	}
}

func workerWithWaitGroup(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(50 * time.Millisecond)
		results <- j * 3
	}
}

func timeoutWorker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		
		// Simulate work with potential timeout
		time.Sleep(time.Duration(j*50) * time.Millisecond)
		
		if j > 3 {
			results <- fmt.Sprintf("Worker %d: Job %d timed out", id, j)
		} else {
			results <- fmt.Sprintf("Worker %d: Job %d completed", id, j)
		}
	}
}

func loadBalancedWorker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		start := time.Now()
		fmt.Printf("Worker %d processing job %d (complexity: %d)\n", id, j, j)
		time.Sleep(time.Duration(j*100) * time.Millisecond)
		elapsed := time.Since(start)
		results <- fmt.Sprintf("Worker %d: Job %d done in %v", id, j, elapsed)
	}
}

func priorityWorker(id int, jobs <-chan struct{priority int; id int}, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d (priority: %d)\n", id, job.id, job.priority)
		time.Sleep(100 * time.Millisecond)
		results <- fmt.Sprintf("Worker %d: Job %d (priority %d) completed", id, job.id, job.priority)
	}
}

func retryWorker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		var result string
		for attempt := 1; attempt <= 3; attempt++ {
			fmt.Printf("Worker %d attempting job %d (attempt %d)\n", id, j, attempt)
			time.Sleep(50 * time.Millisecond)
			
			// Simulate success on attempt 2 or 3
			if attempt >= 2 {
				result = fmt.Sprintf("Worker %d: Job %d succeeded on attempt %d", id, j, attempt)
				break
			}
		}
		results <- result
	}
}

func gracefulWorker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d gracefully processing job %d\n", id, j)
		time.Sleep(80 * time.Millisecond)
		results <- fmt.Sprintf("Worker %d: Job %d completed gracefully", id, j)
	}
}

func statsWorker(id int, jobs <-chan int, results chan<- struct{jobsProcessed int; totalTime time.Duration}, wg *sync.WaitGroup) {
	defer wg.Done()
	var stats struct{jobsProcessed int; totalTime time.Duration}
	
	for j := range jobs {
		start := time.Now()
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Duration(j*30) * time.Millisecond)
		elapsed := time.Since(start)
		
		stats.jobsProcessed++
		stats.totalTime += elapsed
	}
	
	results <- stats
}

func batchWorker(id int, jobs <-chan []int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for batch := range jobs {
		fmt.Printf("Worker %d processing batch %v\n", id, batch)
		time.Sleep(150 * time.Millisecond)
		sum := 0
		for _, num := range batch {
			sum += num
		}
		results <- fmt.Sprintf("Worker %d: Batch %v sum = %d", id, batch, sum)
	}
}

func backpressureWorker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(200 * time.Millisecond) // Slow processing
		results <- fmt.Sprintf("Worker %d: Job %d completed", id, j)
	}
}

func circuitBreakerWorker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup, cb *struct{failures int; maxFailures int; open bool}) {
	defer wg.Done()
	for j := range jobs {
		if cb.open {
			results <- fmt.Sprintf("Worker %d: Circuit breaker OPEN, job %d rejected", id, j)
			continue
		}
		
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(50 * time.Millisecond)
		
		// Simulate failure for jobs 4, 5, 6
		if j >= 4 && j <= 6 {
			cb.failures++
			fmt.Printf("Worker %d: Job %d failed (failures: %d)\n", id, j, cb.failures)
			if cb.failures >= cb.maxFailures {
				cb.open = true
				fmt.Println("Circuit breaker OPENED")
			}
			results <- fmt.Sprintf("Worker %d: Job %d failed", id, j)
		} else {
			results <- fmt.Sprintf("Worker %d: Job %d succeeded", id, j)
		}
	}
}
