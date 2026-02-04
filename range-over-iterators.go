package main

import (
	"fmt"
	"iter"
	"strings"
)

// 1. Basic iterator function
func numbers() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= 5; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// 2. Iterator with two values (key-value)
func keyValuePairs() iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		pairs := map[string]int{
			"apple":  5,
			"banana": 3,
			"orange": 8,
			"grape":  12,
		}
		for key, value := range pairs {
			if !yield(key, value) {
				return
			}
		}
	}
}

// 3. Iterator that can be stopped early
func fibonacci(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

// 4. Iterator over custom data structure
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) InOrder() iter.Seq[int] {
	return func(yield func(int) bool) {
		var traverse func(node *TreeNode) bool
		traverse = func(node *TreeNode) bool {
			if node == nil {
				return true
			}
			// Traverse left subtree
			if !traverse(node.Left) {
				return false
			}
			// Visit current node
			if !yield(node.Value) {
				return false
			}
			// Traverse right subtree
			return traverse(node.Right)
		}
		traverse(t)
	}
}

func (t *TreeNode) PreOrder() iter.Seq[int] {
	return func(yield func(int) bool) {
		var traverse func(node *TreeNode) bool
		traverse = func(node *TreeNode) bool {
			if node == nil {
				return true
			}
			// Visit current node
			if !yield(node.Value) {
				return false
			}
			// Traverse left subtree
			if !traverse(node.Left) {
				return false
			}
			// Traverse right subtree
			return traverse(node.Right)
		}
		traverse(t)
	}
}

// 5. Iterator with filtering
func evenNumbers(max int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i <= max; i++ {
			if i%2 == 0 {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// 6. Iterator over string words
func words(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		wordStart := -1
		for i, r := range s {
			if r == ' ' || r == '\t' || r == '\n' {
				if wordStart != -1 {
					if !yield(s[wordStart:i]) {
						return
					}
					wordStart = -1
				}
			} else if wordStart == -1 {
				wordStart = i
			}
		}
		// Yield the last word if there is one
		if wordStart != -1 {
			yield(s[wordStart:])
		}
	}
}

// 7. Iterator that transforms data
func squares(numbers iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		numbers(func(n int) bool {
			return yield(n * n)
		})
	}
}

// 8. Iterator that combines two iterators
func merge(seq1, seq2 iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		// Create channels to collect values from both sequences
		ch1 := make(chan int, 10)
		ch2 := make(chan int, 10)
		
		// Start goroutines to feed channels
		go func() {
			defer close(ch1)
			seq1(func(n int) bool {
				ch1 <- n
				return true
			})
		}()
		
		go func() {
			defer close(ch2)
			seq2(func(n int) bool {
				ch2 <- n
				return true
			})
		}()
		
		// Yield values from both channels
		for ch1 != nil || ch2 != nil {
			select {
			case n, ok := <-ch1:
				if ok {
					if !yield(n) {
						return
					}
				} else {
					ch1 = nil
				}
			case n, ok := <-ch2:
				if ok {
					if !yield(n) {
						return
					}
				} else {
					ch2 = nil
				}
			}
		}
	}
}

// 9. Iterator with state
type Counter struct {
	current int
	max     int
	step    int
}

func (c *Counter) Iterator() iter.Seq[int] {
	return func(yield func(int) bool) {
		for c.current <= c.max {
			if !yield(c.current) {
				return
			}
			c.current += c.step
		}
	}
}

// 10. Iterator over file lines (simulated)
type File struct {
	lines []string
}

func (f *File) Lines() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, line := range f.lines {
			if !yield(line) {
				return
			}
		}
	}
}

// 11. Iterator with error handling
func safeDivide(dividend, divisor int) iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		for i := 1; i <= 10; i++ {
			if divisor == 0 {
				if !yield(0, fmt.Errorf("division by zero")) {
					return
				}
				continue
			}
			result := dividend * i / divisor
			if !yield(result, nil) {
				return
			}
		}
	}
}

// 12. Iterator that generates permutations
func permutations(arr []int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		var generate func([]int, int) bool
		generate = func(arr []int, n int) bool {
			if n == 1 {
				perm := make([]int, len(arr))
				copy(perm, arr)
				if !yield(perm) {
					return false
				}
			} else {
				for i := 0; i < n; i++ {
					generate(arr, n-1)
					if n%2 == 0 {
						arr[i], arr[n-1] = arr[n-1], arr[i]
					} else {
						arr[0], arr[n-1] = arr[n-1], arr[0]
					}
				}
			}
			return true
		}
		generate(arr, len(arr))
	}
}

// 13. Iterator with backpressure simulation
func slowProducer(max int, delay int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 1; i <= max; i++ {
			// Simulate work
			for j := 0; j < delay; j++ {
				// Busy wait to simulate delay
			}
			if !yield(i) {
				return
			}
		}
	}
}

func main() {
	fmt.Println("=== Range Over Iterators Examples ===")

	// 1. Basic iterator
	fmt.Println("\n1. Basic iterator:")
	fmt.Print("Numbers 1-5: ")
	for num := range numbers() {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 2. Key-value iterator
	fmt.Println("\n2. Key-value iterator:")
	fmt.Println("Fruit counts:")
	for fruit, count := range keyValuePairs() {
		fmt.Printf("  %s: %d\n", fruit, count)
	}

	// 3. Iterator with early stopping
	fmt.Println("\n3. Iterator with early stopping:")
	fmt.Print("First 5 Fibonacci numbers: ")
	count := 0
	for num := range fibonacci(10) {
		fmt.Printf("%d ", num)
		count++
		if count >= 5 {
			break
		}
	}
	fmt.Println()

	// 4. Tree traversal iterators
	fmt.Println("\n4. Tree traversal iterators:")
	root := &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 3,
			Left:  &TreeNode{Value: 2},
			Right: &TreeNode{Value: 4},
		},
		Right: &TreeNode{
			Value: 7,
			Left:  &TreeNode{Value: 6},
			Right: &TreeNode{Value: 8},
		},
	}
	
	fmt.Print("In-order traversal: ")
	for val := range root.InOrder() {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	
	fmt.Print("Pre-order traversal: ")
	for val := range root.PreOrder() {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// 5. Filtering iterator
	fmt.Println("\n5. Filtering iterator:")
	fmt.Print("Even numbers up to 10: ")
	for num := range evenNumbers(10) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 6. String words iterator
	fmt.Println("\n6. String words iterator:")
	text := "Hello world from Go iterators"
	fmt.Printf("Words in '%s': ", text)
	for word := range words(text) {
		fmt.Printf("[%s] ", word)
	}
	fmt.Println()

	// 7. Transforming iterator
	fmt.Println("\n7. Transforming iterator:")
	fmt.Print("Squares of 1-5: ")
	for square := range squares(numbers()) {
		fmt.Printf("%d ", square)
	}
	fmt.Println()

	// 8. Merging iterators
	fmt.Println("\n8. Merging iterators:")
	fmt.Print("Merged sequences: ")
	for num := range merge(numbers(), evenNumbers(8)) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 9. Stateful iterator
	fmt.Println("\n9. Stateful iterator:")
	counter := Counter{current: 1, max: 10, step: 2}
	fmt.Print("Counter (1 to 10, step 2): ")
	for num := range counter.Iterator() {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 10. File lines iterator
	fmt.Println("\n10. File lines iterator:")
	file := File{
		lines: []string{
			"First line",
			"Second line",
			"Third line",
			"Fourth line",
		},
	}
	fmt.Println("File lines:")
	for line := range file.Lines() {
		fmt.Printf("  %s\n", line)
	}

	// 11. Iterator with error handling
	fmt.Println("\n11. Iterator with error handling:")
	fmt.Println("Safe division results:")
	for result, err := range safeDivide(100, 10) {
		if err != nil {
			fmt.Printf("  Error: %v\n", err)
		} else {
			fmt.Printf("  Result: %d\n", result)
		}
	}

	// 12. Permutations iterator
	fmt.Println("\n12. Permutations iterator:")
	arr := []int{1, 2, 3}
	fmt.Printf("Permutations of %v (first 3):\n", arr)
	count = 0
	for perm := range permutations(arr) {
		fmt.Printf("  %v\n", perm)
		count++
		if count >= 3 {
			break
		}
	}

	// 13. Iterator with backpressure
	fmt.Println("\n13. Iterator with backpressure:")
	fmt.Print("Slow producer (first 3): ")
	count = 0
	for num := range slowProducer(10, 1000000) {
		fmt.Printf("%d ", num)
		count++
		if count >= 3 {
			break
		}
	}
	fmt.Println()

	// 14. Using iterators with built-in functions
	fmt.Println("\n14. Using iterators with built-in functions:")
	
	// Collect all values from iterator
	var allNumbers []int
	for num := range numbers() {
		allNumbers = append(allNumbers, num)
	}
	fmt.Printf("Collected numbers: %v\n", allNumbers)
	
	// Find first matching value
	var firstEven int
	found := false
	for num := range evenNumbers(20) {
		if num > 10 {
			firstEven = num
			found = true
			break
		}
	}
	if found {
		fmt.Printf("First even number > 10: %d\n", firstEven)
	}

	// 15. Chaining iterators
	fmt.Println("\n15. Chaining iterators:")
	
	// Chain: numbers -> squares -> filter for > 10
	fmt.Print("Numbers -> squares -> > 10: ")
	for num := range squares(numbers()) {
		if num > 10 {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()

	// 16. Iterator with custom logic
	fmt.Println("\n16. Iterator with custom logic:")
	
	// Iterator that yields prime numbers
	primes := func(max int) iter.Seq[int] {
		return func(yield func(int) bool) {
			for num := 2; num <= max; num++ {
				isPrime := true
				for i := 2; i*i <= num; i++ {
					if num%i == 0 {
						isPrime = false
						break
					}
				}
				if isPrime {
					if !yield(num) {
						return
					}
				}
			}
		}
	}
	
	fmt.Print("Prime numbers up to 20: ")
	for prime := range primes(20) {
		fmt.Printf("%d ", prime)
	}
	fmt.Println()
}
