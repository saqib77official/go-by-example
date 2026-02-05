package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Timers Examples ===")

	// 1. Basic timer
	fmt.Println("\n1. Basic timer:")
	timer1 := time.NewTimer(2 * time.Second)
	defer timer1.Stop()
	
	<-timer1.C
	fmt.Println("Timer 1 fired!")

	// 2. Timer with select
	fmt.Println("\n2. Timer with select:")
	timer2 := time.NewTimer(1 * time.Second)
	defer timer2.Stop()
	
	ch := make(chan string)
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch <- "message"
	}()
	
	select {
	case <-ch:
		fmt.Println("Received message first")
	case <-timer2.C:
		fmt.Println("Timer fired first")
	}

	// 3. Timer reset
	fmt.Println("\n3. Timer reset:")
	timer3 := time.NewTimer(2 * time.Second)
	
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Resetting timer...")
		if !timer3.Stop() {
			fmt.Println("Timer already fired")
		}
		timer3.Reset(1 * time.Second)
	}()
	
	<-timer3.C
	fmt.Println("Timer fired after reset")

	// 4. Multiple timers
	fmt.Println("\n4. Multiple timers:")
	timerA := time.NewTimer(1 * time.Second)
	timerB := time.NewTimer(2 * time.Second)
	defer timerA.Stop()
	defer timerB.Stop()
	
	for i := 0; i < 2; i++ {
		select {
		case <-timerA.C:
			fmt.Println("Timer A fired (1s)")
		case <-timerB.C:
			fmt.Println("Timer B fired (2s)")
		}
	}

	// 5. Timer with duration calculation
	fmt.Println("\n5. Timer with duration calculation:")
	start := time.Now()
	timer4 := time.NewTimer(1500 * time.Millisecond)
	defer timer4.Stop()
	
	<-timer4.C
	elapsed := time.Since(start)
	fmt.Printf("Timer fired after: %v\n", elapsed)

	// 6. Timer in loop
	fmt.Println("\n6. Timer in loop:")
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	count := 0
	for count < 3 {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("Tick %d at %v\n", count, time.Now().Format("15:04:05.000"))
		}
	}

	// 7. Timer with timeout pattern
	fmt.Println("\n7. Timer with timeout pattern:")
	
	operation := func(timeout time.Duration) (string, error) {
		ch := make(chan string)
		
		go func() {
			// Simulate work
			time.Sleep(800 * time.Millisecond)
			ch <- "operation completed"
		}()
		
		select {
		case result := <-ch:
			return result, nil
		case <-time.After(timeout):
			return "", fmt.Errorf("operation timed out after %v", timeout)
		}
	}
	
	// Test with short timeout
	result, err := operation(500 * time.Millisecond)
	if err != nil {
		fmt.Printf("Short timeout: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}
	
	// Test with longer timeout
	result, err = operation(1 * time.Second)
	if err != nil {
		fmt.Printf("Long timeout: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}

	// 8. Timer for periodic tasks
	fmt.Println("\n8. Timer for periodic tasks:")
	
	periodicTask := func(interval time.Duration, iterations int) {
		timer := time.NewTimer(0) // Start immediately
		defer timer.Stop()
		
		for i := 0; i < iterations; i++ {
			<-timer.C
			fmt.Printf("Task iteration %d at %v\n", i+1, time.Now().Format("15:04:05.000"))
			timer.Reset(interval)
		}
	}
	
	go periodicTask(300*time.Millisecond, 3)
	time.Sleep(1100 * time.Millisecond)

	// 9. Timer with cancellation
	fmt.Println("\n9. Timer with cancellation:")
	
	cancellableOperation := func() {
		timer := time.NewTimer(2 * time.Second)
		defer timer.Stop()
		
		cancel := make(chan struct{})
		
		// Cancel after 1 second
		go func() {
			time.Sleep(1 * time.Second)
			close(cancel)
		}()
		
		select {
		case <-timer.C:
			fmt.Println("Operation completed normally")
		case <-cancel:
			fmt.Println("Operation cancelled")
			if !timer.Stop() {
				<-timer.C // Drain if already fired
			}
		}
	}
	
	cancellableOperation()

	// 10. Timer for debouncing
	fmt.Println("\n10. Timer for debouncing:")
	
	debouncer := func() {
		input := make(chan string)
		timer := time.NewTimer(0)
		defer timer.Stop()
		
		// Input generator
		go func() {
			defer close(input)
			inputs := []string{"a", "b", "c", "d", "e"}
			for _, item := range inputs {
				input <- item
				time.Sleep(100 * time.Millisecond)
			}
		}()
		
		for {
			select {
			case item, ok := <-input:
				if !ok {
					return
				}
				fmt.Printf("Received: %s (resetting timer)\n", item)
				if !timer.Stop() {
					<-timer.C // Drain
				}
				timer.Reset(300 * time.Millisecond)
			case <-timer.C:
				fmt.Printf("Debounced output at %v\n", time.Now().Format("15:04:05.000"))
			}
		}
	}
	
	debouncer()

	// 11. Timer for heartbeat
	fmt.Println("\n11. Timer for heartbeat:")
	
	heartbeat := func() {
		timer := time.NewTimer(0)
		defer timer.Stop()
		
		stop := make(chan struct{})
		
		// Stop after 2 seconds
		go func() {
			time.Sleep(2 * time.Second)
			close(stop)
		}()
		
		count := 0
		for {
			select {
			case <-timer.C:
				count++
				fmt.Printf("Heartbeat %d at %v\n", count, time.Now().Format("15:04:05.000"))
				timer.Reset(500 * time.Millisecond)
			case <-stop:
				fmt.Println("Heartbeat stopped")
				return
			}
		}
	}
	
	heartbeat()

	// 12. Timer with multiple durations
	fmt.Println("\n12. Timer with multiple durations:")
	
	multiTimer := func() {
		shortTimer := time.NewTimer(500 * time.Millisecond)
		mediumTimer := time.NewTimer(1 * time.Second)
		longTimer := time.NewTimer(2 * time.Second)
		
		defer shortTimer.Stop()
		defer mediumTimer.Stop()
		defer longTimer.Stop()
		
		for {
			select {
			case <-shortTimer.C:
				fmt.Printf("Short timer fired at %v\n", time.Now().Format("15:04:05.000"))
				shortTimer.Reset(500 * time.Millisecond)
			case <-mediumTimer.C:
				fmt.Printf("Medium timer fired at %v\n", time.Now().Format("15:04:05.000"))
				mediumTimer.Reset(1 * time.Second)
			case <-longTimer.C:
				fmt.Printf("Long timer fired at %v\n", time.Now().Format("15:04:05.000"))
				longTimer.Reset(2 * time.Second)
				return // Stop after long timer
			}
		}
	}
	
	multiTimer()

	// 13. Timer for retry mechanism
	fmt.Println("\n13. Timer for retry mechanism:")
	
	retryOperation := func(maxRetries int) error {
		timer := time.NewTimer(0)
		defer timer.Stop()
		
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("Attempt %d\n", attempt)
			
			// Simulate operation
			go func() {
				time.Sleep(300 * time.Millisecond)
				if attempt == 3 { // Success on 3rd attempt
					timer.Stop() // Stop the timer
				}
			}()
			
			select {
			case <-timer.C:
				fmt.Printf("Attempt %d timed out\n", attempt)
				if attempt == maxRetries {
					return fmt.Errorf("failed after %d attempts", maxRetries)
				}
				timer.Reset(500 * time.Millisecond)
			case <-time.After(100 * time.Millisecond):
				// Timer was stopped (success)
				fmt.Printf("Attempt %d succeeded\n", attempt)
				return nil
			}
		}
		return nil
	}
	
	err := retryOperation(5)
	if err != nil {
		fmt.Printf("Retry failed: %v\n", err)
	} else {
		fmt.Println("Retry succeeded")
	}

	// 14. Timer with statistics
	fmt.Println("\n14. Timer with statistics:")
	
	timerStats := func() {
		var durations []time.Duration
		
		for i := 0; i < 3; i++ {
			start := time.Now()
			timer := time.NewTimer(time.Duration(100+i*50) * time.Millisecond)
			
			<-timer.C
			elapsed := time.Since(start)
			durations = append(durations, elapsed)
			
			fmt.Printf("Timer %d: expected %v, actual %v\n", 
				i+1, time.Duration(100+i*50)*time.Millisecond, elapsed)
		}
		
		// Calculate statistics
		var total time.Duration
		min := durations[0]
		max := durations[0]
		
		for _, d := range durations {
			total += d
			if d < min {
				min = d
			}
			if d > max {
				max = d
			}
		}
		
		avg := total / time.Duration(len(durations))
		
		fmt.Printf("Statistics: Min=%v, Max=%v, Avg=%v\n", min, max, avg)
	}
	
	timerStats()

	// 15. Timer with resource cleanup
	fmt.Println("\n15. Timer with resource cleanup:")
	
	resourceWithTimer := func() {
		timer := time.NewTimer(1 * time.Second)
		defer timer.Stop()
		
		// Simulate resource
		resource := "shared-resource"
		fmt.Printf("Acquired resource: %s\n", resource)
		
		// Work with timeout
		select {
		case <-timer.C:
			fmt.Printf("Work completed, releasing resource: %s\n", resource)
		case <-time.After(500 * time.Millisecond):
			fmt.Printf("Work timed out, force cleanup: %s\n", resource)
		}
	}
	
	resourceWithTimer()

	fmt.Println("All timer examples completed!")
}
