package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Tickers Examples ===")

	// 1. Basic ticker
	fmt.Println("\n1. Basic ticker:")
	ticker1 := time.NewTicker(500 * time.Millisecond)
	defer ticker1.Stop()
	
	for i := 0; i < 3; i++ {
		<-ticker1.C
		fmt.Printf("Tick %d at %v\n", i+1, time.Now().Format("15:04:05.000"))
	}

	// 2. Ticker with select
	fmt.Println("\n2. Ticker with select:")
	ticker2 := time.NewTicker(300 * time.Millisecond)
	defer ticker2.Stop()
	
	ch := make(chan string)
	
	go func() {
		time.Sleep(800 * time.Millisecond)
		ch <- "message"
	}()
	
	for i := 0; i < 5; i++ {
		select {
		case <-ticker2.C:
			fmt.Printf("Tick %d\n", i+1)
		case msg := <-ch:
			fmt.Printf("Received: %s\n", msg)
			break
		}
	}

	// 3. Multiple tickers
	fmt.Println("\n3. Multiple tickers:")
	fastTicker := time.NewTicker(200 * time.Millisecond)
	slowTicker := time.NewTicker(500 * time.Millisecond)
	defer fastTicker.Stop()
	defer slowTicker.Stop()
	
	count := 0
	for count < 6 {
		select {
		case <-fastTicker.C:
			fmt.Printf("Fast tick at %v\n", time.Now().Format("15:04:05.000"))
			count++
		case <-slowTicker.C:
			fmt.Printf("Slow tick at %v\n", time.Now().Format("15:04:05.000"))
			count++
		}
	}

	// 4. Ticker with stop condition
	fmt.Println("\n4. Ticker with stop condition:")
	ticker3 := time.NewTicker(100 * time.Millisecond)
	defer ticker3.Stop()
	
	go func() {
		time.Sleep(550 * time.Millisecond)
		ticker3.Stop()
		fmt.Println("Ticker stopped from goroutine")
	}()
	
	count = 0
	for {
		select {
		case <-ticker3.C:
			count++
			fmt.Printf("Tick %d\n", count)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout reached")
			return
		}
		
		if count >= 10 {
			break
		}
	}

	// 5. Ticker for periodic tasks
	fmt.Println("\n5. Ticker for periodic tasks:")
	
	periodicTask := func() {
		ticker := time.NewTicker(400 * time.Millisecond)
		defer ticker.Stop()
		
		tasks := []string{"Task A", "Task B", "Task C", "Task D"}
		taskIndex := 0
		
		for i := 0; i < len(tasks)*2; i++ {
			<-ticker.C
			task := tasks[taskIndex%len(tasks)]
			fmt.Printf("Executing %s at %v\n", task, time.Now().Format("15:04:05.000"))
			taskIndex++
		}
	}
	
	periodicTask()

	// 6. Ticker with rate limiting
	fmt.Println("\n6. Ticker with rate limiting:")
	
	rateLimited := func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		
		requests := []int{1, 2, 3, 4, 5, 6, 7, 8}
		
		for _, req := range requests {
			<-ticker.C // Wait for ticker
			fmt.Printf("Processing request %d at %v\n", req, time.Now().Format("15:04:05.000"))
		}
	}
	
	rateLimited()

	// 7. Ticker with timeout
	fmt.Println("\n7. Ticker with timeout:")
	
	tickerWithTimeout := func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()
		
		timeout := time.After(2 * time.Second)
		count := 0
		
		for {
			select {
			case <-ticker.C:
				count++
				fmt.Printf("Tick %d\n", count)
			case <-timeout:
				fmt.Printf("Timeout after %d ticks\n", count)
				return
			}
		}
	}
	
	tickerWithTimeout()

	// 8. Ticker for monitoring
	fmt.Println("\n8. Ticker for monitoring:")
	
	monitor := func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		// Simulate system metrics
		getCPU := func() float64 {
			// Simulate CPU usage
			return float64(count%100) / 100.0
		}
		
		getMemory := func() int {
			// Simulate memory usage
			return 50 + count%50
		}
		
		count := 0
		for count < 5 {
			<-ticker.C
			cpu := getCPU()
			mem := getMemory()
			fmt.Printf("Monitor %d: CPU=%.2f%%, Memory=%dMB\n", count+1, cpu*100, mem)
			count++
		}
	}
	
	monitor()

	// 9. Ticker with dynamic interval
	fmt.Println("\n9. Ticker with dynamic interval:")
	
	dynamicTicker := func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		interval := 100 * time.Millisecond
		count := 0
		
		for count < 10 {
			<-ticker.C
			
			// Change interval every 3 ticks
			if count > 0 && count%3 == 0 {
				interval += 100 * time.Millisecond
				ticker.Stop()
				ticker = time.NewTicker(interval)
				fmt.Printf("New interval: %v\n", interval)
			}
			
			fmt.Printf("Tick %d (interval: %v)\n", count+1, interval)
			count++
		}
	}
	
	dynamicTicker()

	// 10. Ticker for heartbeat
	fmt.Println("\n10. Ticker for heartbeat:")
	
	heartbeat := func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		
		stop := make(chan struct{})
		
		// Stop after 5 heartbeats
		go func() {
			time.Sleep(5 * time.Second)
			close(stop)
		}()
		
		beatCount := 0
		for {
			select {
			case <-ticker.C:
				beatCount++
				fmt.Printf("❤️ Heartbeat %d at %v\n", beatCount, time.Now().Format("15:04:05"))
			case <-stop:
				fmt.Println("Heartbeat stopped")
				return
			}
		}
	}
	
	heartbeat()

	// 11. Ticker with data collection
	fmt.Println("\n11. Ticker with data collection:")
	
	dataCollector := func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()
		
		// Simulate sensor data
		getSensorData := func() (temp float64, humidity int) {
			return 20.0 + float64(count%10), 40 + count%20
		}
		
		var readings []struct {
			temp     float64
			humidity int
			time     time.Time
		}
		
		count := 0
		for count < 5 {
			<-ticker.C
			temp, humidity := getSensorData()
			reading := struct {
				temp     float64
				humidity int
				time     time.Time
			}{temp, humidity, time.Now()}
			
			readings = append(readings, reading)
			fmt.Printf("Reading %d: Temp=%.1f°C, Humidity=%d%%\n", 
				count+1, temp, humidity)
			count++
		}
		
		// Calculate averages
		if len(readings) > 0 {
			var totalTemp float64
			var totalHumidity int
			
			for _, r := range readings {
				totalTemp += r.temp
				totalHumidity += r.humidity
			}
			
			avgTemp := totalTemp / float64(len(readings))
			avgHumidity := totalHumidity / len(readings)
			
			fmt.Printf("Averages: Temp=%.1f°C, Humidity=%d%%\n", avgTemp, avgHumidity)
		}
	}
	
	dataCollector()

	// 12. Ticker with batch processing
	fmt.Println("\n12. Ticker with batch processing:")
	
	batchProcessor := func() {
		ticker := time.NewTicker(400 * time.Millisecond)
		defer ticker.Stop()
		
		// Input generator
		input := make(chan int)
		go func() {
			defer close(input)
			for i := 1; i <= 15; i++ {
				input <- i
				time.Sleep(50 * time.Millisecond)
			}
		}()
		
		batch := make([]int, 0, 5)
		
		for {
			select {
			case <-ticker.C:
				if len(batch) > 0 {
					fmt.Printf("Processing batch: %v\n", batch)
					batch = batch[:0] // Clear batch
				} else {
					fmt.Println("No items to process")
				}
			case item, ok := <-input:
				if !ok {
					// Process remaining items
					if len(batch) > 0 {
						fmt.Printf("Final batch: %v\n", batch)
					}
					return
				}
				batch = append(batch, item)
				fmt.Printf("Added to batch: %d\n", item)
			}
		}
	}
	
	batchProcessor()

	// 13. Ticker with statistics
	fmt.Println("\n13. Ticker with statistics:")
	
	tickerStats := func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		
		var tickTimes []time.Time
		count := 0
		
		for count < 10 {
			tickTime := <-ticker.C
			tickTimes = append(tickTimes, tickTime)
			fmt.Printf("Tick %d at %v\n", count+1, tickTime.Format("15:04:05.000"))
			count++
		}
		
		// Calculate intervals
		if len(tickTimes) > 1 {
			var totalInterval time.Duration
			minInterval := tickTimes[1].Sub(tickTimes[0])
			maxInterval := minInterval
			
			for i := 1; i < len(tickTimes); i++ {
				interval := tickTimes[i].Sub(tickTimes[i-1])
				totalInterval += interval
				
				if interval < minInterval {
					minInterval = interval
				}
				if interval > maxInterval {
					maxInterval = interval
				}
			}
			
			avgInterval := totalInterval / time.Duration(len(tickTimes)-1)
			
			fmt.Printf("Interval stats: Min=%v, Max=%v, Avg=%v\n", 
				minInterval, maxInterval, avgInterval)
		}
	}
	
	tickerStats()

	// 14. Ticker with graceful shutdown
	fmt.Println("\n14. Ticker with graceful shutdown:")
	
	gracefulShutdown := func() {
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()
		
		shutdown := make(chan struct{})
		done := make(chan struct{})
		
		// Initiate shutdown after 1 second
		go func() {
			time.Sleep(1 * time.Second)
			close(shutdown)
		}()
		
		// Worker
		go func() {
			defer close(done)
			for {
				select {
				case <-ticker.C:
					fmt.Printf("Working at %v\n", time.Now().Format("15:04:05.000"))
				case <-shutdown:
					fmt.Println("Received shutdown signal")
					return
				}
			}
		}()
		
		<-done
		fmt.Println("Worker shutdown gracefully")
	}
	
	gracefulShutdown()

	// 15. Ticker vs Timer comparison
	fmt.Println("\n15. Ticker vs Timer comparison:")
	
	compareTickerTimer := func() {
		fmt.Println("Using Timer:")
		timer := time.NewTimer(0)
		defer timer.Stop()
		
		for i := 0; i < 3; i++ {
			<-timer.C
			fmt.Printf("Timer tick %d\n", i+1)
			timer.Reset(300 * time.Millisecond)
		}
		
		fmt.Println("\nUsing Ticker:")
		ticker := time.NewTicker(300 * time.Millisecond)
		defer ticker.Stop()
		
		for i := 0; i < 3; i++ {
			<-ticker.C
			fmt.Printf("Ticker tick %d\n", i+1)
		}
	}
	
	compareTickerTimer()

	fmt.Println("All ticker examples completed!")
}
