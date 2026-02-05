package main

import "fmt"

func main() {
	fmt.Println("=== Channel Directions Examples ===")

	// 1. Bidirectional channel (default)
	fmt.Println("\n1. Bidirectional channel:")
	bidirectional := make(chan int)
	fmt.Printf("Bidirectional channel type: %T\n", bidirectional)
	
	// Can send and receive
	go func() {
		bidirectional <- 42
	}()
	value := <-bidirectional
	fmt.Printf("Received from bidirectional: %d\n", value)

	// 2. Send-only channel
	fmt.Println("\n2. Send-only channel:")
	var sendOnly chan<- int = bidirectional
	fmt.Printf("Send-only channel type: %T\n", sendOnly)
	
	// Can only send
	go func() {
		sendOnly <- 100
		// sendOnly <- 200 // This would work
		// value := <-sendOnly // This would cause compile error
	}()
	
	received := <-bidirectional // Receive from original bidirectional
	fmt.Printf("Received: %d\n", received)

	// 3. Receive-only channel
	fmt.Println("\n3. Receive-only channel:")
	var receiveOnly <-chan int = bidirectional
	fmt.Printf("Receive-only channel type: %T\n", receiveOnly)
	
	// Send to bidirectional, receive from receive-only
	go func() {
		bidirectional <- 200
	}()
	
	received = <-receiveOnly
	fmt.Printf("Received from receive-only: %d\n", received)
	
	// receiveOnly <- 300 // This would cause compile error

	// 4. Function with send-only parameter
	fmt.Println("\n4. Function with send-only parameter:")
	
	sendData := func(ch chan<- int, data int) {
		ch <- data
		fmt.Printf("Sent %d to send-only channel\n", data)
	}
	
	sendChannel := make(chan int)
	go sendData(sendChannel, 300)
	
	received = <-sendChannel
	fmt.Printf("Received: %d\n", received)

	// 5. Function with receive-only parameter
	fmt.Println("\n5. Function with receive-only parameter:")
	
	receiveData := func(ch <-chan int) {
		data := <-ch
		fmt.Printf("Received %d from receive-only channel\n", data)
	}
	
	receiveChannel := make(chan int)
	go func() {
		receiveChannel <- 400
	}()
	
	receiveData(receiveChannel)

	// 6. Function with both send-only and receive-only parameters
	fmt.Println("\n6. Function with both directions:")
	
	bridgeData := func(input <-chan int, output chan<- int) {
		data := <-input
		output <- data * 2
		fmt.Printf("Bridged %d to %d\n", data, data*2)
	}
	
	inputChan := make(chan int)
	outputChan := make(chan int)
	
	go func() {
		inputChan <- 50
	}()
	
	go bridgeData(inputChan, outputChan)
	
	result := <-outputChan
	fmt.Printf("Final result: %d\n", result)

	// 7. Channel directions in struct fields
	fmt.Println("\n7. Channel directions in struct:")
	
	type DataProcessor struct {
		input  <-chan int
		output chan<- int
	}
	
	newProcessor := func(input <-chan int, output chan<- int) *DataProcessor {
		return &DataProcessor{
			input:  input,
			output: output,
		}
	}
	
	procInput := make(chan int)
	procOutput := make(chan int)
	processor := newProcessor(procInput, procOutput)
	
	go func() {
		procInput <- 75
	}()
	
	go func() {
		data := <-processor.input
		processor.output <- data + 25
	}()
	
	finalResult := <-procOutput
	fmt.Printf("Processor result: %d\n", finalResult)

	// 8. Returning directional channels
	fmt.Println("\n8. Returning directional channels:")
	
	createChannels := func() (<-chan int, chan<- int) {
		ch := make(chan int)
		return ch, ch
	}
	
	readOnly, writeOnly := createChannels()
	
	go func() {
		writeOnly <- 500
	}()
	
	readValue := <-readOnly
	fmt.Printf("Read value: %d\n", readValue)
	
	// writeOnly <- 600 // This would work
	// readOnly <- 700   // This would cause compile error

	// 9. Channel conversion
	fmt.Println("\n9. Channel conversion:")
	
	// Start with bidirectional
	bidirectionalChan := make(chan string)
	
	// Convert to send-only
	var sendOnlyChan chan<- string = bidirectionalChan
	
	// Convert to receive-only
	var receiveOnlyChan <-chan string = bidirectionalChan
	
	go func() {
		sendOnlyChan <- "Hello from send-only"
	}()
	
	message := <-receiveOnlyChan
	fmt.Printf("Message: %s\n", message)

	// 10. Practical example: Producer-Consumer with directions
	fmt.Println("\n10. Producer-Consumer with directions:")
	
	producer := func(output chan<- int) {
		for i := 1; i <= 3; i++ {
			output <- i
			fmt.Printf("Produced: %d\n", i)
		}
		close(output)
	}
	
	consumer := func(input <-chan int) {
		for value := range input {
			fmt.Printf("Consumed: %d\n", value)
		}
	}
	
	prodConsChan := make(chan int)
	
	go producer(prodConsChan)
	consumer(prodConsChan)

	// 11. Pipeline with directional channels
	fmt.Println("\n11. Pipeline with directions:")
	
	stage1 := func(input <-chan int, output chan<- int) {
		for value := range input {
			output <- value * 2
		}
		close(output)
	}
	
	stage2 := func(input <-chan int, output chan<- int) {
		for value := range input {
			output <- value + 10
		}
		close(output)
	}
	
	// Create pipeline
	pipe1 := make(chan int)
	pipe2 := make(chan int)
	pipe3 := make(chan int)
	
	// Input stage
	go func() {
		defer close(pipe1)
		for i := 1; i <= 3; i++ {
			pipe1 <- i
		}
	}()
	
	// Processing stages
	go stage1(pipe1, pipe2)
	go stage2(pipe2, pipe3)
	
	// Collect results
	for result := range pipe3 {
		fmt.Printf("Pipeline result: %d\n", result)
	}

	// 12. Fan-out with directional channels
	fmt.Println("\n12. Fan-out with directional channels:")
	
	distributor := func(input <-chan int, outputs []chan<- int) {
		for value := range input {
			for _, output := range outputs {
				output <- value
			}
		}
		for _, output := range outputs {
			close(output)
		}
	}
	
	worker := func(id int, input <-chan int, results chan<- int) {
		for value := range input {
			result := value * id
			results <- result
			fmt.Printf("Worker %d: %d -> %d\n", id, value, result)
		}
	}
	
	// Setup fan-out
	input := make(chan int)
	outputs := make([]chan int, 3)
	for i := range outputs {
		outputs[i] = make(chan int)
	}
	
	// Start distributor
	go distributor(input, outputs)
	
	// Start workers
	results := make(chan int, 9)
	for i := 0; i < 3; i++ {
		go worker(i+1, outputs[i], results)
	}
	
	// Send input data
	go func() {
		defer close(input)
		for i := 1; i <= 3; i++ {
			input <- i
		}
	}()
	
	// Collect results
	for i := 0; i < 9; i++ {
		result := <-results
		fmt.Printf("Final result: %d\n", result)
	}

	// 13. Type safety with directions
	fmt.Println("\n13. Type safety demonstration:")
	
	// This function enforces that you can only send to the channel
	safeSender := func(ch chan<- int) {
		ch <- 999
		fmt.Println("Successfully sent to send-only channel")
		// value := <-ch // This would cause compile error
	}
	
	// This function enforces that you can only receive from the channel
	safeReceiver := func(ch <-chan int) {
		value := <-ch
		fmt.Printf("Successfully received %d from receive-only channel\n", value)
		// ch <- 888 // This would cause compile error
	}
	
	safeChan := make(chan int)
	
	go safeSender(safeChan)
	safeReceiver(safeChan)

	fmt.Println("All channel direction examples completed!")
}
