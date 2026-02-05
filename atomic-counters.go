package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	fmt.Println("=== Atomic Counters Examples ===")

	// 1. Basic atomic counter
	fmt.Println("\n1. Basic atomic counter:")
	var counter int64
	
	// Increment from multiple goroutines
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", atomic.LoadInt64(&counter))

	// 2. Compare and swap
	fmt.Println("\n2. Compare and swap:")
	var value int64 = 100
	
	// Try to swap from different goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			old := int64(id * 10)
			new := int64(id * 20)
			
			if atomic.CompareAndSwapInt64(&value, old, new) {
				fmt.Printf("Goroutine %d: Swapped %d -> %d\n", id, old, new)
			} else {
				fmt.Printf("Goroutine %d: Swap failed (current: %d)\n", id, atomic.LoadInt64(&value))
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final value: %d\n", atomic.LoadInt64(&value))

	// 3. Add and fetch
	fmt.Println("\n3. Add and fetch:")
	var counter2 int64
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			oldValue := atomic.AddInt64(&counter2, int64(id+1))
			fmt.Printf("Goroutine %d: Added %d, old value was %d\n", id, id+1, oldValue-id-1)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter: %d\n", atomic.LoadInt64(&counter2))

	// 4. Load and store
	fmt.Println("\n4. Load and store:")
	var config int64
	
	// Store configuration
	go func() {
		atomic.StoreInt64(&config, 42)
		fmt.Println("Configuration stored")
	}()
	
	time.Sleep(50 * time.Millisecond)
	
	// Load configuration
	loadedConfig := atomic.LoadInt64(&config)
	fmt.Printf("Loaded configuration: %d\n", loadedConfig)

	// 5. Atomic counter with mutex comparison
	fmt.Println("\n5. Atomic vs Mutex counter:")
	
	// Atomic version
	var atomicCounter int64
	start := time.Now()
	
	var atomicWg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		atomicWg.Add(1)
		go func() {
			defer atomicWg.Done()
			atomic.AddInt64(&atomicCounter, 1)
		}()
	}
	atomicWg.Wait()
	atomicDuration := time.Since(start)
	
	// Mutex version
	var mutexCounter int64
	var mu sync.Mutex
	start = time.Now()
	
	var mutexWg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		mutexWg.Add(1)
		go func() {
			defer mutexWg.Done()
			mu.Lock()
			mutexCounter++
			mu.Unlock()
		}()
	}
	mutexWg.Wait()
	mutexDuration := time.Since(start)
	
	fmt.Printf("Atomic counter: %d (took %v)\n", atomic.LoadInt64(&atomicCounter), atomicDuration)
	fmt.Printf("Mutex counter: %d (took %v)\n", mutexCounter, mutexDuration)

	// 6. Atomic boolean operations
	fmt.Println("\n6. Atomic boolean operations:")
	var flag int32 // Use int32 for atomic operations
	
	// Set flag
	go func() {
		time.Sleep(100 * time.Millisecond)
		atomic.StoreInt32(&flag, 1)
		fmt.Println("Flag set")
	}()
	
	// Wait for flag
	for atomic.LoadInt32(&flag) == 0 {
		fmt.Println("Waiting for flag...")
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("Flag detected!")

	// 7. Atomic pointer operations
	fmt.Println("\n7. Atomic pointer operations:")
	type Data struct {
		Value int
	}
	
	var dataPtr unsafe.Pointer
	data := &Data{Value: 42}
	atomic.StorePointer(&dataPtr, unsafe.Pointer(data))
	
	// Load and use
	loadedPtr := atomic.LoadPointer(&dataPtr)
	loadedData := (*Data)(loadedPtr)
	fmt.Printf("Loaded data value: %d\n", loadedData.Value)

	// 8. Atomic counter with overflow handling
	fmt.Println("\n8. Atomic counter with overflow:")
	var counter3 uint32
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			atomic.AddUint32(&counter3, uint32(id))
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final counter: %d\n", atomic.LoadUint32(&counter3))

	// 9. Atomic operations for statistics
	fmt.Println("\n9. Atomic statistics:")
	type Stats struct {
		count   int64
		sum     int64
		min     int64
		max     int64
	}
	
	var stats Stats
	atomic.StoreInt64(&stats.min, 999999)
	atomic.StoreInt64(&stats.max, -999999)
	
	numbers := []int64{10, 5, 15, 3, 8, 12, 7}
	
	for _, num := range numbers {
		wg.Add(1)
		go func(n int64) {
			defer wg.Done()
			
			// Update count
			atomic.AddInt64(&stats.count, 1)
			
			// Update sum
			atomic.AddInt64(&stats.sum, n)
			
			// Update min
			for {
				currentMin := atomic.LoadInt64(&stats.min)
				if n < currentMin {
					if atomic.CompareAndSwapInt64(&stats.min, currentMin, n) {
						break
					}
				} else {
					break
				}
			}
			
			// Update max
			for {
				currentMax := atomic.LoadInt64(&stats.max)
				if n > currentMax {
					if atomic.CompareAndSwapInt64(&stats.max, currentMax, n) {
						break
					}
				} else {
					break
				}
			}
		}(num)
	}
	
	wg.Wait()
	
	count := atomic.LoadInt64(&stats.count)
	sum := atomic.LoadInt64(&stats.sum)
	min := atomic.LoadInt64(&stats.min)
	max := atomic.LoadInt64(&stats.max)
	
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Min: %d\n", min)
	fmt.Printf("Max: %d\n", max)
	if count > 0 {
		fmt.Printf("Average: %.2f\n", float64(sum)/float64(count))
	}

	// 10. Atomic counter with reset
	fmt.Println("\n10. Atomic counter with reset:")
	var counter4 int64
	var resetFlag int32
	
	// Incrementer
	go func() {
		for i := 0; i < 100; i++ {
			atomic.AddInt64(&counter4, 1)
			time.Sleep(10 * time.Millisecond)
			
			// Check for reset
			if atomic.LoadInt32(&resetFlag) == 1 {
				fmt.Printf("Reset detected at iteration %d\n", i)
				break
			}
		}
	}()
	
	// Resetter
	go func() {
		time.Sleep(50 * time.Millisecond)
		atomic.StoreInt64(&counter4, 0)
		atomic.StoreInt32(&resetFlag, 1)
		fmt.Println("Counter reset")
	}()
	
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Final counter: %d\n", atomic.LoadInt64(&counter4))

	// 11. Atomic operations for rate limiting
	fmt.Println("\n11. Atomic rate limiting:")
	type RateLimiter struct {
		tokens    int64
		maxTokens int64
		lastRefill int64
	}
	
	var limiter RateLimiter
	atomic.StoreInt64(&limiter.maxTokens, 10)
	atomic.StoreInt64(&limiter.tokens, 10)
	atomic.StoreInt64(&limiter.lastRefill, time.Now().Unix())
	
	allowRequest := func() bool {
		now := time.Now().Unix()
		lastRefill := atomic.LoadInt64(&limiter.lastRefill)
		
		// Refill tokens (1 per second)
		if now-lastRefill >= 1 {
			atomic.StoreInt64(&limiter.tokens, atomic.LoadInt64(&limiter.maxTokens))
			atomic.StoreInt64(&limiter.lastRefill, now)
		}
		
		// Take token
		for {
			currentTokens := atomic.LoadInt64(&limiter.tokens)
			if currentTokens > 0 {
				if atomic.CompareAndSwapInt64(&limiter.tokens, currentTokens, currentTokens-1) {
					return true
				}
			} else {
				return false
			}
		}
	}
	
	// Test rate limiter
	for i := 1; i <= 15; i++ {
		if allowRequest() {
			fmt.Printf("Request %d: Allowed (tokens: %d)\n", i, atomic.LoadInt64(&limiter.tokens))
		} else {
			fmt.Printf("Request %d: Rate limited (tokens: %d)\n", i, atomic.LoadInt64(&limiter.tokens))
		}
		time.Sleep(200 * time.Millisecond)
	}

	// 12. Atomic operations for reference counting
	fmt.Println("\n12. Atomic reference counting:")
	type Resource struct {
		id     int
		name   string
		refCount int32
	}
	
	resource := &Resource{id: 1, name: "shared-resource"}
	atomic.StoreInt32(&resource.refCount, 1)
	
	// Use resource from multiple goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Acquire reference
			atomic.AddInt32(&resource.refCount, 1)
			fmt.Printf("Goroutine %d: Acquired resource (refs: %d)\n", id, atomic.LoadInt32(&resource.refCount))
			
			time.Sleep(50 * time.Millisecond)
			
			// Release reference
			atomic.AddInt32(&resource.refCount, -1)
			fmt.Printf("Goroutine %d: Released resource (refs: %d)\n", id, atomic.LoadInt32(&resource.refCount))
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("Final reference count: %d\n", atomic.LoadInt32(&resource.refCount))

	// 13. Atomic operations for circular buffer
	fmt.Println("\n13. Atomic circular buffer:")
	type CircularBuffer struct {
		buffer [8]int64
		head   int64
		tail   int64
		size   int64
	}
	
	var cb CircularBuffer
	
	// Producer
	go func() {
		for i := 0; i < 20; i++ {
			head := atomic.LoadInt64(&cb.head)
			tail := atomic.LoadInt64(&cb.tail)
			size := atomic.LoadInt64(&cb.size)
			
			if size < 8 {
				atomic.StoreInt64(&cb.buffer[head%8], int64(i))
				atomic.StoreInt64(&cb.head, (head+1)%8)
				atomic.AddInt64(&cb.size, 1)
				fmt.Printf("Produced %d at position %d\n", i, head%8)
			} else {
				fmt.Printf("Buffer full, cannot produce %d\n", i)
			}
			
			time.Sleep(50 * time.Millisecond)
		}
	}()
	
	// Consumer
	for i := 0; i < 15; i++ {
		tail := atomic.LoadInt64(&cb.tail)
		size := atomic.LoadInt64(&cb.size)
		
		if size > 0 {
			value := atomic.LoadInt64(&cb.buffer[tail%8])
			atomic.StoreInt64(&cb.tail, (tail+1)%8)
			atomic.AddInt64(&cb.size, -1)
			fmt.Printf("Consumed %d from position %d\n", value, tail%8)
		} else {
			fmt.Printf("Buffer empty\n")
		}
		
		time.Sleep(70 * time.Millisecond)
	}

	// 14. Atomic operations for bit flags
	fmt.Println("\n14. Atomic bit flags:")
	var flags int32
	
	setFlag := func(flag int32) {
		for {
			old := atomic.LoadInt32(&flags)
			new := old | flag
			if atomic.CompareAndSwapInt32(&flags, old, new) {
				fmt.Printf("Set flag %d (was: %d, now: %d)\n", flag, old, new)
				break
			}
		}
	}
	
	clearFlag := func(flag int32) {
		for {
			old := atomic.LoadInt32(&flags)
			new := old &^ flag
			if atomic.CompareAndSwapInt32(&flags, old, new) {
				fmt.Printf("Cleared flag %d (was: %d, now: %d)\n", flag, old, new)
				break
			}
		}
	}
	
	checkFlag := func(flag int32) bool {
		return atomic.LoadInt32(&flags)&flag != 0
	}
	
	// Set some flags
	setFlag(1) // 001
	setFlag(2) // 010
	setFlag(4) // 100
	
	fmt.Printf("Flag 1 set: %t\n", checkFlag(1))
	fmt.Printf("Flag 2 set: %t\n", checkFlag(2))
	fmt.Printf("Flag 4 set: %t\n", checkFlag(4))
	fmt.Printf("Flag 8 set: %t\n", checkFlag(8))
	
	clearFlag(2)
	fmt.Printf("Flag 2 cleared: %t\n", checkFlag(2))

	// 15. Atomic operations for high-frequency counting
	fmt.Println("\n15. High-frequency atomic counting:")
	var highFreqCounter int64
	
	// High-frequency incrementer
	go func() {
		for i := 0; i < 100000; i++ {
			atomic.AddInt64(&highFreqCounter, 1)
		}
	}()
	
	// Monitor counter
	for i := 0; i < 10; i++ {
		time.Sleep(10 * time.Millisecond)
		count := atomic.LoadInt64(&highFreqCounter)
		fmt.Printf("Count at %dms: %d\n", (i+1)*10, count)
	}
	
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Final count: %d\n", atomic.LoadInt64(&highFreqCounter))

	fmt.Println("All atomic counter examples completed!")
}
