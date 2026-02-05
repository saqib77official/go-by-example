package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Range Over Channels Examples ===")

	// 1. Basic range over channel
	fmt.Println("\n1. Basic range over channel:")
	ch := make(chan int)
	
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("Sent: %d\n", i)
		}
	}()
	
	fmt.Println("Receiving with range:")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	fmt.Println("Range completed")

	// 2. Range over buffered channel
	fmt.Println("\n2. Range over buffered channel:")
	buffered := make(chan string, 3)
	
	// Pre-fill buffer
	buffered <- "first"
	buffered <- "second"
	buffered <- "third"
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		buffered <- "fourth"
		time.Sleep(100 * time.Millisecond)
		buffered <- "fifth"
		close(buffered)
	}()
	
	fmt.Println("Ranging over buffered channel:")
	for value := range buffered {
		fmt.Printf("Received: %s\n", value)
	}

	// 3. Range with early termination
	fmt.Println("\n3. Range with early termination:")
	ch2 := make(chan int)
	
	go func() {
		defer close(ch2)
		for i := 1; i <= 10; i++ {
			ch2 <- i
			fmt.Printf("Sent: %d\n", i)
		}
	}()
	
	fmt.Println("Ranging with early termination:")
	count := 0
	for value := range ch2 {
		fmt.Printf("Received: %d\n", value)
		count++
		if count >= 5 {
			fmt.Println("Early termination after 5 items")
			break
		}
	}

	// 4. Range over multiple channels (sequentially)
	fmt.Println("\n4. Range over multiple channels:")
	ch3 := make(chan int)
	ch4 := make(chan string)
	
	go func() {
		defer close(ch3)
		for i := 1; i <= 3; i++ {
			ch3 <- i
		}
	}()
	
	go func() {
		defer close(ch4)
		for i := 1; i <= 3; i++ {
			ch4 <- fmt.Sprintf("item-%d", i)
		}
	}()
	
	fmt.Println("Ranging over ch3:")
	for value := range ch3 {
		fmt.Printf("From ch3: %d\n", value)
	}
	
	fmt.Println("Ranging over ch4:")
	for value := range ch4 {
		fmt.Printf("From ch4: %s\n", value)
	}

	// 5. Range with select (multiplexing)
	fmt.Println("\n5. Range with select:")
	ch5 := make(chan int)
	ch6 := make(chan int)
	
	go func() {
		defer close(ch5)
		for i := 1; i <= 5; i++ {
			ch5 <- i * 10
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(ch6)
		for i := 1; i <= 5; i++ {
			ch6 <- i * 100
			time.Sleep(30 * time.Millisecond)
		}
	}()
	
	// Use select to multiplex
	ch5Open, ch6Open := true, true
	for ch5Open || ch6Open {
		select {
		case value, ok := <-ch5:
			if ok {
				fmt.Printf("From ch5: %d\n", value)
			} else {
				fmt.Println("ch5 closed")
				ch5Open = false
			}
		case value, ok := <-ch6:
			if ok {
				fmt.Printf("From ch6: %d\n", value)
			} else {
				fmt.Println("ch6 closed")
				ch6Open = false
			}
		}
	}

	// 6. Range with timeout
	fmt.Println("\n6. Range with timeout:")
	ch7 := make(chan int)
	
	go func() {
		defer close(ch7)
		for i := 1; i <= 10; i++ {
			ch7 <- i
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	fmt.Println("Ranging with timeout:")
	timeout := time.After(1 * time.Second)
	
Loop:
	for {
		select {
		case value, ok := <-ch7:
			if !ok {
				fmt.Println("Channel closed")
				break Loop
			}
			fmt.Printf("Received: %d\n", value)
		case <-timeout:
			fmt.Println("Timeout reached")
			break Loop
		}
	}

	// 7. Range with filtering
	fmt.Println("\n7. Range with filtering:")
	input := make(chan int)
	filtered := make(chan int)
	
	// Producer
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()
	
	// Filter
	go func() {
		defer close(filtered)
		for value := range input {
			if value%2 == 0 {
				filtered <- value
				fmt.Printf("Filtered: %d\n", value)
			}
		}
	}()
	
	fmt.Println("Filtered results:")
	for value := range filtered {
		fmt.Printf("Even number: %d\n", value)
	}

	// 8. Range with transformation
	fmt.Println("\n8. Range with transformation:")
	source := make(chan int)
	transformed := make(chan int)
	
	// Source
	go func() {
		defer close(source)
		for i := 1; i <= 5; i++ {
			source <- i
		}
	}()
	
	// Transformer
	go func() {
		defer close(transformed)
		for value := range source {
			result := value * value
			transformed <- result
			fmt.Printf("Transformed %d -> %d\n", value, result)
		}
	}()
	
	fmt.Println("Transformed results:")
	for value := range transformed {
		fmt.Printf("Square: %d\n", value)
	}

	// 9. Range with aggregation
	fmt.Println("\n9. Range with aggregation:")
	numbers := make(chan int)
	
	go func() {
		defer close(numbers)
		for i := 1; i <= 10; i++ {
			numbers <- i
		}
	}()
	
	sum := 0
	count := 0
	for value := range numbers {
		sum += value
		count++
		fmt.Printf("Accumulating: %d (sum: %d, count: %d)\n", value, sum, count)
	}
	
	fmt.Printf("Final: sum=%d, count=%d, average=%.2f\n", sum, count, float64(sum)/float64(count))

	// 10. Range with batch processing
	fmt.Println("\n10. Range with batch processing:")
	items := make(chan int)
	
	go func() {
		defer close(items)
		for i := 1; i <= 12; i++ {
			items <- i
		}
	}()
	
	batch := make([]int, 0, 3)
	for item := range items {
		batch = append(batch, item)
		
		if len(batch) == 3 {
			fmt.Printf("Processing batch: %v\n", batch)
			batch = batch[:0] // Reset batch
		}
	}
	
	// Process remaining items
	if len(batch) > 0 {
		fmt.Printf("Processing final batch: %v\n", batch)
	}

	// 11. Range with context cancellation
	fmt.Println("\n11. Range with context cancellation:")
	work := make(chan int)
	stop := make(chan struct{})
	
	go func() {
		defer close(work)
		for i := 1; i <= 20; i++ {
			select {
			case work <- i:
				fmt.Printf("Sent work: %d\n", i)
			case <-stop:
				fmt.Println("Producer stopped")
				return
			}
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Consumer with cancellation
	go func() {
		time.Sleep(300 * time.Millisecond)
		close(stop)
	}()
	
	fmt.Println("Consuming with cancellation:")
	for value := range work {
		fmt.Printf("Received work: %d\n", value)
	}

	// 12. Range with error handling
	fmt.Println("\n12. Range with error handling:")
	type Result struct {
		Value int
		Error error
	}
	
	results := make(chan Result)
	
	go func() {
		defer close(results)
		for i := 1; i <= 8; i++ {
			if i == 4 || i == 7 {
				results <- Result{Error: fmt.Errorf("error processing %d", i)}
				continue
			}
			results <- Result{Value: i * 10}
		}
	}()
	
	fmt.Println("Processing with error handling:")
	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		} else {
			fmt.Printf("Success: %d\n", result.Value)
		}
	}

	// 13. Range with rate limiting
	fmt.Println("\n13. Range with rate limiting:")
	data := make(chan int)
	limiter := time.NewTicker(100 * time.Millisecond)
	defer limiter.Stop()
	
	go func() {
		defer close(data)
		for i := 1; i <= 5; i++ {
			data <- i
		}
	}()
	
	fmt.Println("Consuming with rate limiting:")
	for value := range data {
		<-limiter.C // Wait for ticker
		fmt.Printf("Processed: %d at %v\n", value, time.Now().Format("15:04:05.000"))
	}

	// 14. Range with fan-in pattern
	fmt.Println("\n14. Range with fan-in:")
	inputs := []chan int{
		make(chan int),
		make(chan int),
		make(chan int),
	}
	
	// Start producers
	for i, input := range inputs {
		go func(id int, ch chan int) {
			defer close(ch)
			for j := 1; j <= 3; j++ {
				ch <- id*10 + j
			}
		}(i+1, input)
	}
	
	// Fan-in
	fanIn := make(chan int)
	
	for _, input := range inputs {
		go func(ch <-chan int) {
			for value := range ch {
				fanIn <- value
			}
		}(input)
	}
	
	// Close fan-in when all inputs done
	go func() {
		for _, input := range inputs {
			<-input // Wait for close (this is a trick to wait)
		}
		close(fanIn)
	}()
	
	fmt.Println("Fan-in results:")
	for value := range fanIn {
		fmt.Printf("Received: %d\n", value)
	}

	// 15. Range with statistics collection
	fmt.Println("\n15. Range with statistics:")
	metrics := make(chan int)
	
	type Stats struct {
		Count int
		Sum   int
		Min   int
		Max   int
	}
	
	go func() {
		defer close(metrics)
		numbers := []int{5, 2, 8, 1, 9, 3, 7, 4, 6}
		for _, num := range numbers {
			metrics <- num
		}
	}()
	
	stats := Stats{Min: 999999, Max: -999999}
	for value := range metrics {
		stats.Count++
		stats.Sum += value
		if value < stats.Min {
			stats.Min = value
		}
		if value > stats.Max {
			stats.Max = value
		}
		fmt.Printf("Processing: %d (min: %d, max: %d, sum: %d)\n", 
			value, stats.Min, stats.Max, stats.Sum)
	}
	
	fmt.Printf("Final stats: Count=%d, Sum=%d, Min=%d, Max=%d, Avg=%.2f\n",
		stats.Count, stats.Sum, stats.Min, stats.Max, 
		float64(stats.Sum)/float64(stats.Count))

	fmt.Println("All range over channels examples completed!")
}
