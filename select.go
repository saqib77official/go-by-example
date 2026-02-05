package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select Examples ===")

	// 1. Basic select with multiple channels
	fmt.Println("\n1. Basic select:")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from channel 1"
	}()
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "from channel 2"
	}()
	
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received: %s\n", msg2)
	}

	// 2. Select with default case (non-blocking)
	fmt.Println("\n2. Select with default:")
	ch3 := make(chan int)
	
	select {
	case value := <-ch3:
		fmt.Printf("Received: %d\n", value)
	default:
		fmt.Println("No data available, default case executed")
	}
	
	// Send with default
	select {
	case ch3 <- 42:
		fmt.Println("Sent successfully")
	default:
		fmt.Println("Could not send, channel blocked")
	}

	// 3. Select with timeout
	fmt.Println("\n3. Select with timeout:")
	ch4 := make(chan string)
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch4 <- "delayed message"
	}()
	
	select {
	case msg := <-ch4:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timeout occurred")
	}

	// 4. Select with multiple cases
	fmt.Println("\n4. Select with multiple cases:")
	ch5 := make(chan int)
	ch6 := make(chan string)
	ch7 := make(chan bool)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch5 <- 100
	}()
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch6 <- "hello"
	}()
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch7 <- true
	}()
	
	for i := 0; i < 3; i++ {
		select {
		case num := <-ch5:
			fmt.Printf("Received number: %d\n", num)
		case str := <-ch6:
			fmt.Printf("Received string: %s\n", str)
		case flag := <-ch7:
			fmt.Printf("Received boolean: %t\n", flag)
		}
	}

	// 5. Select in loop (channel multiplexing)
	fmt.Println("\n5. Select in loop:")
	input1 := make(chan int)
	input2 := make(chan int)
	
	// Producers
	go func() {
		defer close(input1)
		for i := 1; i <= 3; i++ {
			input1 <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	go func() {
		defer close(input2)
		for i := 4; i <= 6; i++ {
			input2 <- i
			time.Sleep(30 * time.Millisecond)
		}
	}()
	
	// Consumer with select
	for {
		select {
		case val, ok := <-input1:
			if !ok {
				input1 = nil
			} else {
				fmt.Printf("From input1: %d\n", val)
			}
		case val, ok := <-input2:
			if !ok {
				input2 = nil
			} else {
				fmt.Printf("From input2: %d\n", val)
			}
		}
		
		if input1 == nil && input2 == nil {
			break
		}
	}

	// 6. Select for random selection
	fmt.Println("\n6. Select for random selection:")
	ch8 := make(chan string)
	ch9 := make(chan string)
	
	// Send to both channels
	go func() {
		ch8 <- "first"
	}()
	
	go func() {
		ch9 <- "second"
	}()
	
	// Random selection
	select {
	case msg1 := <-ch8:
		fmt.Printf("Selected: %s\n", msg1)
	case msg2 := <-ch9:
		fmt.Printf("Selected: %s\n", msg2)
	}

	// 7. Select with empty case
	fmt.Println("\n7. Select with empty case:")
	ch10 := make(chan int)
	
	// This select will block until ch10 receives data
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch10 <- 42
	}()
	
	select {
	case value := <-ch10:
		fmt.Printf("Received: %d\n", value)
	}

	// 8. Select with send operations
	fmt.Println("\n8. Select with send operations:")
	ch11 := make(chan int)
	ch12 := make(chan int)
	
	// Receiver
	go func() {
		for i := 0; i < 3; i++ {
			select {
			case val := <-ch11:
				fmt.Printf("Received from ch11: %d\n", val)
			case val := <-ch12:
				fmt.Printf("Received from ch12: %d\n", val)
			}
		}
	}()
	
	// Senders
	for i := 1; i <= 3; i++ {
		select {
		case ch11 <- i * 10:
			fmt.Printf("Sent %d to ch11\n", i*10)
		case ch12 <- i * 100:
			fmt.Printf("Sent %d to ch12\n", i*100)
		}
		time.Sleep(50 * time.Millisecond)
	}

	// 9. Select with nil channels
	fmt.Println("\n9. Select with nil channels:")
	var ch13 chan int // nil channel
	ch14 := make(chan int)
	
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch14 <- 200
	}()
	
	select {
	case val := <-ch13:
		fmt.Printf("From nil channel: %d\n", val) // This will never execute
	case val := <-ch14:
		fmt.Printf("From valid channel: %d\n", val)
	default:
		fmt.Println("Default case")
	}

	// 10. Select for load balancing
	fmt.Println("\n10. Select for load balancing:")
	workers := []chan int{make(chan int), make(chan int), make(chan int)}
	
	// Start workers
	for i, worker := range workers {
		go func(id int, w chan int) {
			for job := range w {
				fmt.Printf("Worker %d processing job %d\n", id, job)
				time.Sleep(50 * time.Millisecond)
			}
		}(i, worker)
	}
	
	// Distribute jobs using select
	for job := 1; job <= 6; job++ {
		select {
		case workers[0] <- job:
			fmt.Printf("Sent job %d to worker 0\n", job)
		case workers[1] <- job:
			fmt.Printf("Sent job %d to worker 1\n", job)
		case workers[2] <- job:
			fmt.Printf("Sent job %d to worker 2\n", job)
		}
	}
	
	// Close workers
	for _, worker := range workers {
		close(worker)
	}
	time.Sleep(400 * time.Millisecond)

	// 11. Select with multiple timeouts
	fmt.Println("\n11. Select with multiple timeouts:")
	ch15 := make(chan string)
	
	go func() {
		time.Sleep(150 * time.Millisecond)
		ch15 <- "ready"
	}()
	
	select {
	case msg := <-ch15:
		fmt.Printf("Received: %s\n", msg)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Short timeout")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Long timeout") // This won't execute if short timeout fires first
	}

	// 12. Select for channel closing detection
	fmt.Println("\n12. Select for channel closing:")
	ch16 := make(chan int)
	
	go func() {
		for i := 1; i <= 3; i++ {
			ch16 <- i
		}
		close(ch16)
	}()
	
	for {
		select {
		case value, ok := <-ch16:
			if ok {
				fmt.Printf("Received: %d\n", value)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}

	// 13. Select with ticker integration
	fmt.Println("\n13. Select with ticker:")
	ch17 := make(chan int)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	
	go func() {
		time.Sleep(250 * time.Millisecond)
		ch17 <- 42
	}()
	
	for i := 0; i < 5; i++ {
		select {
		case value := <-ch17:
			fmt.Printf("Received from channel: %d\n", value)
			return
		case <-ticker.C:
			fmt.Printf("Tick %d\n", i+1)
		}
	}

	// 14. Select pattern: Quit channel
	fmt.Println("\n14. Select with quit channel:")
	workChan := make(chan int)
	quitChan := make(chan bool)
	
	// Worker
	go func() {
		for {
			select {
			case work := <-workChan:
				fmt.Printf("Processing work: %d\n", work)
			case <-quitChan:
				fmt.Println("Received quit signal")
				return
			}
		}
	}()
	
	// Send some work
	for i := 1; i <= 3; i++ {
		workChan <- i
	}
	
	time.Sleep(100 * time.Millisecond)
	
	// Send quit signal
	quitChan <- true
	time.Sleep(50 * time.Millisecond)

	// 15. Select for heartbeat pattern
	fmt.Println("\n15. Select for heartbeat:")
	dataChan := make(chan int)
	heartbeatChan := make(chan time.Time)
	
	// Data producer
	go func() {
		for i := 1; i <= 3; i++ {
			dataChan <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(dataChan)
	}()
	
	// Heartbeat generator
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case heartbeatChan <- time.Now():
			case <-time.After(50 * time.Millisecond):
				return
			}
		}
	}()
	
	// Consumer with heartbeat
	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				fmt.Println("Data channel closed")
				goto done
			}
			fmt.Printf("Data: %d\n", data)
		case heartbeat := <-heartbeatChan:
			fmt.Printf("Heartbeat: %v\n", heartbeat.Format("15:04:05.000"))
		}
	}
	
done:
	fmt.Println("All select examples completed!")
}
