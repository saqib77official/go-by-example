package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func main() {
	fmt.Println("=== Mutexes Examples ===")

	// 1. Basic mutex
	fmt.Println("\n1. Basic mutex:")
	counter := Counter{}
	var wg sync.WaitGroup
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			counter.mu.Lock()
			counter.value++
			fmt.Printf("Goroutine %d: Counter = %d\n", id, counter.value)
			counter.mu.Unlock()
		}(i)
	}
	wg.Wait()

	// 2. Mutex with defer
	fmt.Println("\n2. Mutex with defer:")
	var mu sync.Mutex
	data := []int{}
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			data = append(data, id)
			fmt.Printf("Goroutine %d: Data = %v\n", id, data)
		}(i)
	}
	wg.Wait()

	// 3. RWMutex
	fmt.Println("\n3. RWMutex:")
	var rwMu sync.RWMutex
	sharedData := map[string]int{"a": 1, "b": 2}
	
	// Readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			rwMu.RLock()
			defer rwMu.RUnlock()
			fmt.Printf("Reader %d: %v\n", id, sharedData)
		}(i)
	}
	
	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		rwMu.Lock()
		defer rwMu.Unlock()
		sharedData["c"] = 3
		fmt.Printf("Writer: %v\n", sharedData)
	}()
	wg.Wait()

	fmt.Println("All mutex examples completed!")
}
