package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Rate Limiting Examples ===")

	// 1. Basic rate limiting with ticker
	fmt.Println("\n1. Basic rate limiting with ticker:")
	
	basicRateLimit := func() {
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()
		
		requests := []int{1, 2, 3, 4, 5, 6, 7, 8}
		
		for _, req := range requests {
			<-ticker.C // Wait for ticker
			fmt.Printf("Processing request %d at %v\n", req, time.Now().Format("15:04:05.000"))
		}
	}
	
	basicRateLimit()

	// 2. Token bucket rate limiter
	fmt.Println("\n2. Token bucket rate limiter:")
	
	tokenBucket := func() {
		type TokenBucket struct {
			tokens     int
			maxTokens  int
			refillRate int
			lastRefill time.Time
			mu         sync.Mutex
		}
		
		bucket := &TokenBucket{
			maxTokens:  5,
			tokens:     5,
			refillRate: 2, // tokens per second
			lastRefill: time.Now(),
		}
		
		takeToken := func() bool {
			bucket.mu.Lock()
			defer bucket.mu.Unlock()
			
			// Refill tokens
			now := time.Now()
			elapsed := now.Sub(bucket.lastRefill)
			tokensToAdd := int(elapsed.Seconds()) * bucket.refillRate
			bucket.tokens += tokensToAdd
			if bucket.tokens > bucket.maxTokens {
				bucket.tokens = bucket.maxTokens
			}
			bucket.lastRefill = now
			
			// Take token
			if bucket.tokens > 0 {
				bucket.tokens--
				return true
			}
			return false
		}
		
		// Process requests
		for i := 1; i <= 15; i++ {
			if takeToken() {
				fmt.Printf("Request %d: Allowed (tokens: %d)\n", i, bucket.tokens)
			} else {
				fmt.Printf("Request %d: Rate limited (tokens: %d)\n", i, bucket.tokens)
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	tokenBucket()

	// 3. Sliding window rate limiter
	fmt.Println("\n3. Sliding window rate limiter:")
	
	slidingWindow := func() {
		type Window struct {
			requests []time.Time
			maxRequests int
			duration   time.Duration
			mu        sync.Mutex
		}
		
		window := &Window{
			maxRequests: 3,
			duration:   time.Second,
		}
		
		allowRequest := func() bool {
			window.mu.Lock()
			defer window.mu.Unlock()
			
			now := time.Now()
			
			// Remove old requests
			validRequests := make([]time.Time, 0)
			for _, req := range window.requests {
				if now.Sub(req) <= window.duration {
					validRequests = append(validRequests, req)
				}
			}
			
			// Check if we can add new request
			if len(validRequests) < window.maxRequests {
				window.requests = append(validRequests, now)
				return true
			}
			
			window.requests = validRequests
			return false
		}
		
		// Process requests
		for i := 1; i <= 10; i++ {
			if allowRequest() {
				fmt.Printf("Request %d: Allowed\n", i)
			} else {
				fmt.Printf("Request %d: Rate limited\n", i)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
	
	slidingWindow()

	// 4. Fixed window counter rate limiter
	fmt.Println("\n4. Fixed window counter rate limiter:")
	
	fixedWindow := func() {
		type FixedWindow struct {
			count     int
			maxCount  int
			windowEnd time.Time
			duration  time.Duration
			mu        sync.Mutex
		}
		
		window := &FixedWindow{
			maxCount: 5,
			duration:  time.Second,
			windowEnd: time.Now().Add(time.Second),
		}
		
		allowRequest := func() bool {
			window.mu.Lock()
			defer window.mu.Unlock()
			
			now := time.Now()
			
			// Reset window if expired
			if now.After(window.windowEnd) {
				window.count = 0
				window.windowEnd = now.Add(window.duration)
			}
			
			// Check if we can add request
			if window.count < window.maxCount {
				window.count++
				return true
			}
			
			return false
		}
		
		// Process requests
		for i := 1; i <= 15; i++ {
			if allowRequest() {
				fmt.Printf("Request %d: Allowed (count: %d/%d)\n", i, window.count, window.maxCount)
			} else {
				fmt.Printf("Request %d: Rate limited (count: %d/%d)\n", i, window.count, window.maxCount)
			}
			time.Sleep(150 * time.Millisecond)
		}
	}
	
	fixedWindow()

	// 5. Leaky bucket rate limiter
	fmt.Println("\n5. Leaky bucket rate limiter:")
	
	leakyBucket := func() {
		type LeakyBucket struct {
			capacity   int
			content    int
			leakRate  time.Duration
			lastLeak   time.Time
			requests   chan struct{}
			mu         sync.Mutex
		}
		
		bucket := &LeakyBucket{
			capacity:  5,
			leakRate: 200 * time.Millisecond,
			requests:  make(chan struct{}, 100),
		}
		
		// Leaker goroutine
		go func() {
			ticker := time.NewTicker(bucket.leakRate)
			defer ticker.Stop()
			
			for range ticker.C {
				bucket.mu.Lock()
				if bucket.content > 0 {
					bucket.content--
				}
				bucket.mu.Unlock()
			}
		}()
		
		allowRequest := func() bool {
			bucket.mu.Lock()
			defer bucket.mu.Unlock()
			
			if bucket.content < bucket.capacity {
				bucket.content++
				return true
			}
			return false
		}
		
		// Process requests
		for i := 1; i <= 15; i++ {
			if allowRequest() {
				fmt.Printf("Request %d: Accepted (content: %d/%d)\n", i, bucket.content, bucket.capacity)
			} else {
				fmt.Printf("Request %d: Rejected (content: %d/%d)\n", i, bucket.content, bucket.capacity)
			}
			time.Sleep(50 * time.Millisecond)
		}
	}
	
	leakyBucket()

	// 6. Rate limiting with multiple tiers
	fmt.Println("\n6. Rate limiting with multiple tiers:")
	
	multiTier := func() {
		type Tier struct {
			name       string
			maxRequests int
			duration   time.Duration
			count      int
			windowEnd  time.Time
		}
		
		tiers := []Tier{
			{"Basic", 5, time.Second},
			{"Premium", 10, time.Second},
			{"Enterprise", 20, time.Second},
		}
		
		allowRequest := func(tierIndex int) bool {
			tier := &tiers[tierIndex]
			now := time.Now()
			
			// Reset window if expired
			if now.After(tier.windowEnd) {
				tier.count = 0
				tier.windowEnd = now.Add(tier.duration)
			}
			
			if tier.count < tier.maxRequests {
				tier.count++
				return true
			}
			return false
		}
		
		// Test different tiers
		for tierIndex, tier := range tiers {
			fmt.Printf("Testing %s tier:\n", tier.name)
			for i := 1; i <= 8; i++ {
				if allowRequest(tierIndex) {
					fmt.Printf("  Request %d: Allowed (%d/%d)\n", i, tier.count, tier.maxRequests)
				} else {
					fmt.Printf("  Request %d: Rate limited (%d/%d)\n", i, tier.count, tier.maxRequests)
				}
				time.Sleep(100 * time.Millisecond)
			}
			fmt.Println()
		}
	}
	
	multiTier()

	// 7. Rate limiting with burst capacity
	fmt.Println("\n7. Rate limiting with burst capacity:")
	
	burstRateLimit := func() {
		type BurstLimiter struct {
			tokens     int
			maxTokens  int
			burstSize  int
			refillRate int
			lastRefill time.Time
			mu         sync.Mutex
		}
		
		limiter := &BurstLimiter{
			maxTokens:  3,
			burstSize:  8,
			tokens:    8, // Start with full burst
			refillRate: 1, // 1 token per second
			lastRefill: time.Now(),
		}
		
		allowRequest := func() bool {
			limiter.mu.Lock()
			defer limiter.mu.Unlock()
			
			// Refill tokens
			now := time.Now()
			elapsed := now.Sub(limiter.lastRefill)
			tokensToAdd := int(elapsed.Seconds()) * limiter.refillRate
			limiter.tokens += tokensToAdd
			if limiter.tokens > limiter.maxTokens {
				limiter.tokens = limiter.maxTokens
			}
			limiter.lastRefill = now
			
			// Take token
			if limiter.tokens > 0 {
				limiter.tokens--
				return true
			}
			return false
		}
		
		// Process burst of requests
		for i := 1; i <= 15; i++ {
			if allowRequest() {
				fmt.Printf("Request %d: Allowed (tokens: %d)\n", i, limiter.tokens)
			} else {
				fmt.Printf("Request %d: Rate limited (tokens: %d)\n", i, limiter.tokens)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
	
	burstRateLimit()

	// 8. Rate limiting with priority
	fmt.Println("\n8. Rate limiting with priority:")
	
	priorityRateLimit := func() {
		type PriorityRequest struct {
			priority int
			id       int
		}
		
		requests := make(chan PriorityRequest, 20)
		limiter := make(chan struct{}, 3) // 3 concurrent requests
		
		// Request generator
		go func() {
			defer close(requests)
			priorities := []int{1, 3, 2, 3, 1, 2, 3, 1}
			for i, priority := range priorities {
				requests <- PriorityRequest{priority: priority, id: i + 1}
			}
		}()
		
		// Priority processor
		go func() {
			for req := range requests {
				limiter <- struct{}{}
				fmt.Printf("Processing request %d (priority %d)\n", req.id, req.priority)
				time.Sleep(300 * time.Millisecond)
				<-limiter
			}
		}()
		
		time.Sleep(3 * time.Second)
	}
	
	priorityRateLimit()

	// 9. Rate limiting with backpressure
	fmt.Println("\n9. Rate limiting with backpressure:")
	
	backpressureRateLimit := func() {
		requests := make(chan int, 20)
		processed := make(chan int, 20)
		
		// Rate limited processor
		go func() {
			ticker := time.NewTicker(200 * time.Millisecond)
			defer ticker.Stop()
			
			for req := range requests {
				<-ticker.C // Wait for rate limit
				processed <- req
				fmt.Printf("Processed request %d\n", req)
			}
		}()
		
		// Request generator
		go func() {
			defer close(requests)
			for i := 1; i <= 10; i++ {
				select {
				case requests <- i:
					fmt.Printf("Sent request %d\n", i)
				case <-time.After(50 * time.Millisecond):
					fmt.Printf("Request %d dropped (backpressure)\n", i)
				}
			}
		}()
		
		// Collect results
		for i := 0; i < 10; i++ {
			select {
			case result := <-processed:
				fmt.Printf("Received processed %d\n", result)
			case <-time.After(1 * time.Second):
				fmt.Println("Timeout waiting for results")
				return
			}
		}
	}
	
	backpressureRateLimit()

	// 10. Rate limiting with adaptive control
	fmt.Println("\n10. Rate limiting with adaptive control:")
	
	adaptiveRateLimit := func() {
		type AdaptiveLimiter struct {
			currentRate time.Duration
			minRate    time.Duration
			maxRate    time.Duration
			errors      int
			mu          sync.Mutex
		}
		
		limiter := &AdaptiveLimiter{
			currentRate: 200 * time.Millisecond,
			minRate:    100 * time.Millisecond,
			maxRate:    500 * time.Millisecond,
		}
		
		allowRequest := func() bool {
			limiter.mu.Lock()
			defer limiter.mu.Unlock()
			
			// Simulate error rate based processing
			time.Sleep(limiter.currentRate)
			
			// Simulate random errors
			if time.Now().Unix()%5 == 0 {
				limiter.errors++
				if limiter.errors > 2 {
					// Slow down due to errors
					limiter.currentRate *= 2
					if limiter.currentRate > limiter.maxRate {
						limiter.currentRate = limiter.maxRate
					}
				}
				return false
			}
			
			// Speed up if no errors
			if limiter.errors == 0 {
				limiter.currentRate = limiter.currentRate * 3 / 4
				if limiter.currentRate < limiter.minRate {
					limiter.currentRate = limiter.minRate
				}
			}
			
			return true
		}
		
		// Process requests
		for i := 1; i <= 20; i++ {
			if allowRequest() {
				fmt.Printf("Request %d: Success (rate: %v)\n", i, limiter.currentRate)
			} else {
				fmt.Printf("Request %d: Failed (rate: %v)\n", i, limiter.currentRate)
			}
		}
	}
	
	adaptiveRateLimit()

	// 11. Rate limiting with distributed coordination
	fmt.Println("\n11. Rate limiting with distributed coordination:")
	
	distributedRateLimit := func() {
		type Node struct {
			id       int
			tokens   int
			maxTokens int
			mu       sync.Mutex
		}
		
		nodes := []*Node{
			{id: 1, tokens: 5, maxTokens: 5},
			{id: 2, tokens: 5, maxTokens: 5},
			{id: 3, tokens: 5, maxTokens: 5},
		}
		
		// Token redistribution
		redistributeTokens := func() {
			totalTokens := 0
			for _, node := range nodes {
				totalTokens += node.tokens
			}
			
			avgTokens := totalTokens / len(nodes)
			for _, node := range nodes {
				if node.tokens < avgTokens-1 {
					node.tokens = avgTokens
				} else if node.tokens > avgTokens+1 {
					node.tokens = avgTokens
				}
			}
		}
		
		allowRequest := func(nodeID int) bool {
			node := nodes[nodeID-1]
			node.mu.Lock()
			defer node.mu.Unlock()
			
			if node.tokens > 0 {
				node.tokens--
				return true
			}
			return false
		}
		
		// Process requests on different nodes
		for i := 1; i <= 15; i++ {
			nodeID := (i-1)%3 + 1
			if allowRequest(nodeID) {
				fmt.Printf("Request %d: Allowed on node %d (tokens: %d)\n", i, nodeID, nodes[nodeID-1].tokens)
			} else {
				fmt.Printf("Request %d: Rate limited on node %d (tokens: %d)\n", i, nodeID, nodes[nodeID-1].tokens)
			}
			
			// Redistribute every 5 requests
			if i%5 == 0 {
				redistributeTokens()
				fmt.Println("Redistributed tokens")
			}
			
			time.Sleep(100 * time.Millisecond)
		}
	}
	
	distributedRateLimit()

	// 12. Rate limiting with graceful degradation
	fmt.Println("\n12. Rate limiting with graceful degradation:")
	
	gracefulDegradation := func() {
		type DegradationLimiter struct {
			currentLimit int
			maxLimit     int
			minLimit     int
			load         int
			mu           sync.Mutex
		}
		
		limiter := &DegradationLimiter{
			currentLimit: 10,
			maxLimit:     10,
			minLimit:     2,
		}
		
		adjustLimit := func() {
			limiter.mu.Lock()
			defer limiter.mu.Unlock()
			
			// Adjust based on load
			if limiter.load > 8 {
				limiter.currentLimit = limiter.maxLimit
			} else if limiter.load > 5 {
				limiter.currentLimit = limiter.maxLimit * 3 / 4
			} else {
				limiter.currentLimit = limiter.minLimit
			}
		}
		
		allowRequest := func() bool {
			limiter.mu.Lock()
			defer limiter.mu.Unlock()
			
			if limiter.currentLimit > 0 {
				limiter.currentLimit--
				return true
			}
			return false
		}
		
		// Simulate varying load
		loads := []int{3, 7, 12, 8, 4, 15, 6, 9}
		
		for i, load := range loads {
			limiter.load = load
			adjustLimit()
			
			fmt.Printf("Load: %d, Limit: %d\n", load, limiter.currentLimit)
			
			// Process requests
			for j := 1; j <= 5; j++ {
				if allowRequest() {
					fmt.Printf("  Request %d.%d: Allowed\n", i+1, j)
				} else {
					fmt.Printf("  Request %d.%d: Rate limited\n", i+1, j)
				}
			}
		}
	}
	
	gracefulDegradation()

	fmt.Println("All rate limiting examples completed!")
}
