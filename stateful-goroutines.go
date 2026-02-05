package main

import (
	"fmt"
	"sync"
	"time"
)

type StatefulWorker struct {
	id     int
	state  int
	mu     sync.Mutex
}

func (w *StatefulWorker) Process(input int) int {
	w.mu.Lock()
	defer w.mu.Unlock()
	
	w.state += input
	fmt.Printf("Worker %d: state = %d (added %d)\n", w.id, w.state, input)
	return w.state
}

func (w *StatefulWorker) GetState() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.state
}

func main() {
	fmt.Println("=== Stateful Goroutines Examples ===")

	// 1. Basic stateful worker
	fmt.Println("\n1. Basic stateful worker:")
	worker := StatefulWorker{id: 1, state: 0}
	
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			worker.Process(val)
		}(i)
	}
	wg.Wait()
	
	fmt.Printf("Final state: %d\n", worker.GetState())

	// 2. Multiple stateful workers
	fmt.Println("\n2. Multiple stateful workers:")
	workers := []*StatefulWorker{
		{id: 1, state: 0},
		{id: 2, state: 0},
		{id: 3, state: 0},
	}
	
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			worker := workers[val%3]
			worker.Process(val)
		}(i)
	}
	wg.Wait()
	
	for _, w := range workers {
		fmt.Printf("Worker %d final state: %d\n", w.id, w.GetState())
	}

	fmt.Println("All stateful goroutine examples completed!")
}
