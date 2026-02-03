package main

import "fmt"

// Function that returns a closure
func getAdder() func(int) int {
	return func(x int) int {
		return x + 1
	}
}

// Function that returns a closure with captured variable
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Function that returns a closure with multiple captured variables
func calculator() func(string, int) int {
	sum := 0
	return func(operation string, value int) int {
		switch operation {
		case "add":
			sum += value
		case "subtract":
			sum -= value
		case "multiply":
			sum *= value
		case "reset":
			sum = 0
		}
		return sum
	}
}

// Function returning closure that maintains state
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Closure that captures multiple values
func makeGreeter(greeting string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s!", greeting, name)
	}
}

// Closure with filter function
func makeFilter(predicate func(int) bool) func([]int) []int {
	return func(numbers []int) []int {
		var result []int
		for _, num := range numbers {
			if predicate(num) {
				result = append(result, num)
			}
		}
		return result
	}
}

// Closure that modifies external variable
func accumulator() func(int) int {
	total := 0
	return func(x int) int {
		total += x
		return total
	}
}

// Closure with deferred execution
func deferExample() func() {
	message := "Hello!"
	return func() {
		fmt.Println(message)
	}
}

// Closure that creates a generator
func fibonacciGenerator() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// Closure with multiple return values
func makeValidator(min, max int) func(int) (bool, string) {
	return func(value int) (bool, string) {
		if value < min {
			return false, fmt.Sprintf("value %d is below minimum %d", value, min)
		}
		if value > max {
			return false, fmt.Sprintf("value %d is above maximum %d", value, max)
		}
		return true, fmt.Sprintf("value %d is within range [%d, %d]", value, min, max)
	}
}

// Closure that maintains a cache
func memoize(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(x int) int {
		if val, exists := cache[x]; exists {
			return val
		}
		result := fn(x)
		cache[x] = result
		return result
	}
}

func main() {
	fmt.Println("=== Closures Examples ===")

	// 1. Basic closure
	fmt.Println("\n1. Basic closure:")
	adder := getAdder()
	result := adder(5)
	fmt.Printf("5 + 1 = %d\n", result)

	// 2. Closure with captured variable
	fmt.Println("\n2. Closure with captured variable:")
	times2 := makeMultiplier(2)
	times3 := makeMultiplier(3)
	times5 := makeMultiplier(5)
	
	fmt.Printf("10 * 2 = %d\n", times2(10))
	fmt.Printf("10 * 3 = %d\n", times3(10))
	fmt.Printf("10 * 5 = %d\n", times5(10))

	// 3. Closure maintaining state
	fmt.Println("\n3. Closure maintaining state:")
	calc := calculator()
	fmt.Printf("Initial: %d\n", calc("reset", 0))
	fmt.Printf("Add 10: %d\n", calc("add", 10))
	fmt.Printf("Add 5: %d\n", calc("add", 5))
	fmt.Printf("Subtract 3: %d\n", calc("subtract", 3))
	fmt.Printf("Multiply 2: %d\n", calc("multiply", 2))

	// 4. Counter closure
	fmt.Println("\n4. Counter closure:")
	count1 := counter()
	count2 := counter()
	
	fmt.Printf("Counter 1: %d\n", count1())
	fmt.Printf("Counter 1: %d\n", count1())
	fmt.Printf("Counter 2: %d\n", count2())
	fmt.Printf("Counter 1: %d\n", count1())
	fmt.Printf("Counter 2: %d\n", count2())

	// 5. Greeter closure
	fmt.Println("\n5. Greeter closure:")
	helloGreeter := makeGreeter("Hello")
	hiGreeter := makeGreeter("Hi")
	
	fmt.Printf("%s\n", helloGreeter("Alice"))
	fmt.Printf("%s\n", hiGreeter("Bob"))
	fmt.Printf("%s\n", helloGreeter("Charlie"))

	// 6. Filter closure
	fmt.Println("\n6. Filter closure:")
	evenFilter := makeFilter(func(n int) bool { return n%2 == 0 })
	positiveFilter := makeFilter(func(n int) bool { return n > 0 })
	
	numbers := []int{-2, -1, 0, 1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Even: %v\n", evenFilter(numbers))
	fmt.Printf("Positive: %v\n", positiveFilter(numbers))

	// 7. Accumulator closure
	fmt.Println("\n7. Accumulator closure:")
	acc := accumulator()
	fmt.Printf("Add 10: %d\n", acc(10))
	fmt.Printf("Add 5: %d\n", acc(5))
	fmt.Printf("Add 15: %d\n", acc(15))

	// 8. Closure with deferred execution
	fmt.Println("\n8. Deferred closure:")
	deferred := deferExample()
	fmt.Println("Creating deferred closure...")
	deferred() // Executes when called

	// 9. Fibonacci generator closure
	fmt.Println("\n9. Fibonacci generator:")
	fib := fibonacciGenerator()
	fmt.Print("First 10 Fibonacci numbers: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()

	// 10. Validator closure
	fmt.Println("\n10. Validator closure:")
	validateAge := makeValidator(0, 120)
	validateScore := makeValidator(0, 100)
	
	if valid, msg := validateAge(25); valid {
		fmt.Printf("✓ %s\n", msg)
	} else {
		fmt.Printf("✗ %s\n", msg)
	}
	
	if valid, msg := validateAge(150); valid {
		fmt.Printf("✓ %s\n", msg)
	} else {
		fmt.Printf("✗ %s\n", msg)
	}
	
	if valid, msg := validateScore(85); valid {
		fmt.Printf("✓ %s\n", msg)
	} else {
		fmt.Printf("✗ %s\n", msg)
	}

	// 11. Memoization closure
	fmt.Println("\n11. Memoization closure:")
	slowFunction := func(n int) int {
		fmt.Printf("Computing factorial(%d)...\n", n)
		if n <= 1 {
			return 1
		}
		return n * slowFunction(n-1)
	}
	
	memoizedFactorial := memoize(slowFunction)
	fmt.Printf("First call: %d\n", memoizedFactorial(5))
	fmt.Printf("Second call (cached): %d\n", memoizedFactorial(5))
	fmt.Printf("Third call (cached): %d\n", memoizedFactorial(5))

	// 12. Closure in loops (common pitfall and solution)
	fmt.Println("\n12. Closure in loops:")
	
	// Wrong way - captures the same variable
	var wrongFuncs []func() int
	for i := 0; i < 3; i++ {
		wrongFuncs = append(wrongFuncs, func() int {
			return i
		})
	}
	fmt.Print("Wrong way: ")
	for _, f := range wrongFuncs {
		fmt.Printf("%d ", f())
	}
	fmt.Println()
	
	// Correct way - capture loop variable
	var correctFuncs []func() int
	for i := 0; i < 3; i++ {
		i := i // Create new variable for each iteration
		correctFuncs = append(correctFuncs, func() int {
			return i
		})
	}
	fmt.Print("Correct way: ")
	for _, f := range correctFuncs {
		fmt.Printf("%d ", f())
	}
	fmt.Println()

	// 13. Closure as method-like function
	fmt.Println("\n13. Closure as method-like function:")
	type Person struct {
		Name string
		Age  int
	}
	
	makePersonGreeter := func(p Person) func(string) string {
		return func(title string) string {
			return fmt.Sprintf("%s %s, age %d", title, p.Name, p.Age)
		}
	}
	
	person := Person{Name: "Alice", Age: 25}
	greeter := makePersonGreeter(person)
	fmt.Printf("%s\n", greeter("Mr."))
	fmt.Printf("%s\n", greeter("Dr."))
}
