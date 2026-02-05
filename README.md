# Go by Example

A comprehensive collection of Go programming examples covering fundamental concepts, advanced features, and modern Go patterns. Each file contains practical examples with detailed explanations to help you learn Go programming from beginner to advanced level.

## 📁 Files Overview

### 🌱 **Beginner Examples (13 files)**

### 📋 [variables.go](./variables.go)
**Variable Declarations and Usage**
- Variable declarations with `var` keyword
- Type inference with `:=`
- Zero values
- Multiple variable declarations
- Variable reassignment

**Key Concepts:**
```go
var name string = "John"
age := 25
var x, y, z int = 10, 20, 30
```

---

### 🔢 [constants.go](./constants.go)
**Constants and Enumerations**
- Constant declarations
- Grouped constants
- Enumerated constants using `iota`
- Typed constants

**Key Concepts:**
```go
const PI = 3.14159
const (
    RED = iota
    GREEN
    BLUE
)
```

---

### 🔄 [for.go](./for.go)
**Loop Constructs**
- Basic for loops
- For loops as while loops
- Range loops over slices, maps, strings
- Nested loops
- Loop control with `break` and `continue`

**Key Concepts:**
```go
for i := 0; i < 5; i++ { }
for condition { }
for index, value := range slice { }
```

---

### 🔀 [if-else.go](./if-else.go)
**Conditional Statements**
- Basic if statements
- If-else chains
- If with initialization
- Nested conditions
- Error handling patterns

**Key Concepts:**
```go
if age >= 18 { }
if x := getValue(); x > 0 { }
if err != nil { }
```

---

### 🎛️ [switch.go](./switch.go)
**Switch Statements**
- Basic switch cases
- Switch without expression
- Multiple case values
- Fallthrough behavior
- Type switches

**Key Concepts:**
```go
switch value {
case 1: // case 1
case 2, 3: // case 2 or 3
}
switch v := value.(type) { } // type switch
```

---

### 📦 [arrays.go](./arrays.go)
**Array Operations**
- Array declaration and initialization
- Array length and capacity
- Multidimensional arrays
- Array comparison
- Value vs reference behavior

**Key Concepts:**
```go
var arr [5]int
arr := [3]string{"a", "b", "c"}
matrix := [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
```

---

### 🔪 [slices.go](./slices.go)
**Slice Operations**
- Slice creation and initialization
- `append` and `copy` operations
- Slicing and reslicing
- Slice growth and capacity
- Reference type behavior

**Key Concepts:**
```go
slice := []int{1, 2, 3}
slice = append(slice, 4)
copy(dest, src)
slice = slice[1:4]
```

---

### 🗺️ [maps.go](./maps.go)
**Map Operations**
- Map creation and initialization
- Adding, accessing, and deleting elements
- Checking key existence
- Iterating over maps
- Nested maps and complex values

**Key Concepts:**
```go
m := make(map[string]int)
m["key"] = value
if val, exists := m["key"]; exists { }
delete(m, "key")
```

---

### ⚙️ [functions.go](./functions.go)
**Function Definitions**
- Basic functions
- Functions with parameters
- Named return values
- Higher-order functions
- Function expressions

**Key Concepts:**
```go
func add(a, b int) int { return a + b }
func calculate() (area, perimeter float64) { }
func filter(nums []int, pred func(int) bool) { }
```

---

### 🔄 [multiple-return-values.go](./multiple-return-values.go)
**Multiple Return Values**
- Functions returning multiple values
- Named return values
- Error handling patterns
- Ignoring return values
- Tuple-like behavior

**Key Concepts:**
```go
func divide(a, b float64) (float64, error) { }
result, err := divide(10, 2)
name, age, email := getPersonInfo(1)
```

---

### 📝 [variadic-functions.go](./variadic-functions.go)
**Variadic Functions**
- Functions with variable parameters
- Passing slices to variadic functions
- Mixed regular and variadic parameters
- Practical use cases

**Key Concepts:**
```go
func sum(numbers ...int) int { }
func greet(greeting string, names ...string) { }
result := sum(slice...)
```

---

### 🎯 [closures.go](./closures.go)
**Closures and Anonymous Functions**
- Function literals
- Captured variables
- Functions returning closures
- Practical closure patterns
- Common pitfalls

**Key Concepts:**
```go
adder := func(x int) int { return x + 1 }
makeMultiplier := func(factor int) func(int) int { }
counter := func() func() int { }
```

---

### 🌳 [recursion.go](./recursion.go)
**Recursive Functions**
- Basic recursion patterns
- Factorial and Fibonacci
- Tree traversal
- Memoization
- Complex recursive algorithms

**Key Concepts:**
```go
func factorial(n int) int { }
func fibonacci(n int) int { }
func binarySearch(arr []int, target, left, right int) int { }
```

---

## 🔧 Advanced Go Concepts

### 🚀 **Intermediate & Advanced Examples (15 files)**

### 🔄 [range-over-built-in-types.go](./range-over-built-in-types.go)
**Range Over Built-in Types**
- Range over slices, arrays, strings, maps, channels
- Different iteration patterns (index/value, values only, keys only)
- Range over complex data structures
- Range with early termination and modification

**Key Concepts:**
```go
for index, value := range slice { }
for key, value := range map { }
for index, rune := range string { }
for value := range channel { }
```

---

### 🎯 [pointers.go](./pointers.go)
**Pointer Operations**
- Pointer declaration and dereferencing
- Pointers with functions and structs
- Pointer to pointer concepts
- Memory management and comparison

**Key Concepts:**
```go
var p *int = &x
*p = 42
func modify(ptr *Type) { }
```

---

### 📝 [strings-and-runes.go](./strings-and-runes.go)
**String and Rune Manipulation**
- String operations and formatting
- Unicode and rune handling
- String conversion and validation
- Text processing patterns

**Key Concepts:**
```go
runes := []rune("Hello, 世界")
strings.Contains(s, substr)
strings.Split(s, sep)
```

---

### 🏗️ [structs.go](./structs.go)
**Struct Operations**
- Struct definition and initialization
- Nested structs and pointers
- Struct comparison and copying
- Struct tags and zero values

**Key Concepts:**
```go
type Person struct { Name string; Age int }
person := Person{Name: "Alice", Age: 25}
personPtr := &person
```

---

### ⚙️ [methods.go](./methods.go)
**Method Definitions**
- Value vs pointer receivers
- Method promotion and overriding
- Method expressions and values
- Interface implementation through methods

**Key Concepts:**
```go
func (r Rectangle) Area() float64 { }
func (c *Counter) Increment() { }
areaFunc := Rectangle.Area
```

---

### 🔌 [interfaces.go](./interfaces.go)
**Interface Implementation**
- Interface definition and implementation
- Empty interface and type assertions
- Interface composition
- Dynamic polymorphism

**Key Concepts:**
```go
type Shape interface { Area() float64 }
var s Shape = &Circle{}
if val, ok := s.(*Circle) { }
```

---

### 🏷️ [enums.go](./enums.go)
**Enumeration Patterns**
- Enum patterns with iota
- String enums and bitmask enums
- Enum validation and iteration
- State machine patterns

**Key Concepts:**
```go
const (
    RED = iota
    GREEN
    BLUE
)
type Status int
```

---

### 🔗 [struct-embedding.go](./struct-embedding.go)
**Struct Embedding**
- Basic and multiple embedding
- Method promotion and overriding
- Name conflict resolution
- Composition patterns

**Key Concepts:**
```go
type Employee struct {
    Person
    EmployeeID int
}
emp.Name // Promoted field
```

---

### 🔧 [generics.go](./generics.go)
**Generic Programming**
- Generic functions and structs
- Type constraints and interfaces
- Generic collections and algorithms
- Real-world generic patterns

**Key Concepts:**
```go
func Print[T any](value T) { }
type Container[T any] struct { data []T }
type Number interface { int | float64 }
```

---

### 🔄 [range-over-iterators.go](./range-over-iterators.go)
**Custom Iterators**
- Custom iterator functions
- Iterator composition and chaining
- Tree traversal with iterators
- Iterator with early termination

**Key Concepts:**
```go
func numbers() iter.Seq[int] { }
for num := range numbers() { }
type TreeNode struct { ... }
```

---

### ❌ [errors.go](./errors.go)
**Error Handling**
- Error creation and handling
- Error wrapping and unwrapping
- Error checking with `errors.Is` and `errors.As`
- Panic and recover patterns

**Key Concepts:**
```go
errors.New("error message")
fmt.Errorf("context: %w", err)
errors.Is(err, targetErr)
```

---

### 🚨 [custom-errors.go](./custom-errors.go)
**Custom Error Types**
- Custom error types with methods
- Error with context and metadata
- Business logic errors
- Error aggregation and retry patterns

**Key Concepts:**
```go
type AppError struct { Code int; Message string }
func (ae *AppError) Error() string { }
type ValidationError struct { Field string }
```

---

### 🚀 [goroutines.go](./goroutines.go)
**Concurrent Programming**
- Basic goroutine usage
- WaitGroup synchronization
- Worker pool patterns
- Atomic operations and mutexes

**Key Concepts:**
```go
go func() { }()
var wg sync.WaitGroup
wg.Add(1); defer wg.Done()
atomic.AddInt64(&counter, 1)
```

---

### 📡 [channels.go](./channels.go)
**Channel Communication**
- Channel operations and directions
- Select statements and timeouts
- Fan-in/fan-out patterns
- Channel pipelines

**Key Concepts:**
```go
ch := make(chan Type)
select { case <-ch: case <-time.After(): }
for val := range ch { }
```

---

### 📦 [channel-buffering.go](./channel-buffering.go)
**Buffered Channels**
- Buffered vs unbuffered channels
- Buffering for performance
- Semaphore patterns
- Rate limiting and batching

**Key Concepts:**
```go
ch := make(chan int, 10)
semaphore := make(chan struct{}, 3)
select { case ch <- val: default: }
```

---

## 🔧 Expert Go Concepts

### 🚀 **Expert Examples (20 files)**

### 🔄 [channel-synchronization.go](./channel-synchronization.go)
**Channel Synchronization**
- Basic synchronization patterns
- Pipeline coordination
- Producer-consumer patterns
- Resource access control

**Key Concepts:**
```go
done := make(chan bool)
<-done // Wait for completion
close(done) // Signal completion
```

---

### 🎯 [channel-directions.go](./channel-directions.go)
**Channel Directions**
- Send-only channels
- Receive-only channels
- Function parameters with directions
- Type safety with directional channels

**Key Concepts:**
```go
var sendOnly chan<- int
var receiveOnly <-chan int
func sendData(ch chan<- int) { }
```

---

### 🔀 [select.go](./select.go)
**Select Statements**
- Multiple channel operations
- Non-blocking selects
- Timeout patterns
- Random selection

**Key Concepts:**
```go
select {
case <-ch1:
case <-ch2:
default:
}
```

---

### ⏱️ [timeouts.go](./timeouts.go)
**Timeout Patterns**
- Basic timeouts with time.After
- Timeout with error handling
- Retry mechanisms
- Circuit breaker patterns

**Key Concepts:**
```go
select {
case <-ch:
case <-time.After(timeout):
}
```

---

### 🚫 [non-blocking-channel-operations.go](./non-blocking-channel-operations.go)
**Non-Blocking Channel Operations**
- Non-blocking sends/receives
- Backpressure handling
- Load shedding
- Health checks

**Key Concepts:**
```go
select {
case ch <- value:
default:
    // Channel blocked
}
```

---

### 🔒 [closing-channels.go](./closing-channels.go)
**Channel Closing**
- Safe channel closing
- Detecting closed channels
- Graceful shutdown
- Resource cleanup

**Key Concepts:**
```go
close(ch)
value, ok := <-ch
for value := range ch { }
```

---

### 📡 [range-over-channels.go](./range-over-channels.go)
**Range Over Channels**
- Basic channel iteration
- Early termination
- Multiplexing with select
- Statistics collection

**Key Concepts:**
```go
for value := range ch {
    // Process value
}
```

---

### ⏰ [timers.go](./timers.go)
**Timer Operations**
- Basic timer usage
- Timer reset and stop
- Timeout patterns
- Debouncing

**Key Concepts:**
```go
timer := time.NewTimer(duration)
<-timer.C
timer.Reset(duration)
```

---

### 🔄 [tickers.go](./tickers.go)
**Ticker Operations**
- Periodic operations
- Rate limiting
- Heartbeat patterns
- Monitoring

**Key Concepts:**
```go
ticker := time.NewTicker(interval)
<-ticker.C
ticker.Stop()
```

---

### 👥 [worker-pools.go](./worker-pools.go)
**Worker Pools**
- Basic worker pools
- Dynamic scaling
- Load balancing
- Circuit breakers

**Key Concepts:**
```go
jobs := make(chan Job, 100)
results := make(chan Result, 100)
for w := 1; w <= numWorkers; w++ {
    go worker(w, jobs, results)
}
```

---

### ⏳ [waitgroups.go](./waitgroups.go)
**WaitGroup Synchronization**
- Basic WaitGroup usage
- Nested WaitGroups
- Error handling
- Progress tracking

**Key Concepts:**
```go
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done() }()
wg.Wait()
```

---

### 🚦 [rate-limiting.go](./rate-limiting.go)
**Rate Limiting**
- Token bucket algorithm
- Sliding window
- Leaky bucket
- Adaptive control

**Key Concepts:**
```go
ticker := time.NewTicker(rate)
<-ticker.C // Rate limit
```

---

### 🔢 [atomic-counters.go](./atomic-counters.go)
**Atomic Operations**
- Atomic counters
- Compare and swap
- Memory barriers
- Lock-free programming

**Key Concepts:**
```go
atomic.AddInt64(&counter, 1)
atomic.CompareAndSwapInt64(&ptr, old, new)
atomic.LoadInt64(&value)
```

---

### 🔐 [mutexes.go](./mutexes.go)
**Mutex Synchronization**
- Basic mutex usage
- RWMutex for readers/writers
- Defer unlock
- Deadlock prevention

**Key Concepts:**
```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
var rwMu sync.RWMutex
rwMu.RLock()
```

---

### 🧠 [stateful-goroutines.go](./stateful-goroutines.go)
**Stateful Goroutines**
- Maintaining state
- Thread-safe operations
- State management patterns
- Concurrent state updates

**Key Concepts:**
```go
type Worker struct {
    mu    sync.Mutex
    state int
}
```

---

### 📊 [sorting.go](./sorting.go)
**Sorting Operations**
- Basic slice sorting
- Custom type sorting
- Reverse sorting
- Partial sorting

**Key Concepts:**
```go
sort.Ints(slice)
sort.Strings(slice)
sort.Slice(custom, func(i, j int) bool { })
```

---

### 🎯 [sorting-by-functions.go](./sorting-by-functions.go)
**Custom Sorting**
- Multi-criteria sorting
- Custom comparators
- Complex sorting logic
- Performance optimization

**Key Concepts:**
```go
sort.Slice(items, func(i, j int) bool {
    return items[i].field < items[j].field
})
```

---

### 😱 [panic.go](./panic.go)
**Panic Handling**
- Basic panic usage
- Panic sources
- Panic vs errors
- Safe operations

**Key Concepts:**
```go
panic("error message")
var slice []int
slice[0] = 1 // Panics
```

---

### ⏪ [defer.go](./defer.go)
**Defer Statements**
- Basic defer usage
- LIFO execution order
- Resource cleanup
- Function return values

**Key Concepts:**
```go
defer file.Close()
defer fmt.Println("cleanup")
```

---

### 🛡️ [recover.go](./recover.go)
**Panic Recovery**
- Basic recover patterns
- Error handling
- Cleanup on panic
- Goroutine recovery

**Key Concepts:**
```go
defer func() {
    if r := recover(); r != nil {
        // Handle panic
    }
}()
```

---

## � Getting Started

### Prerequisites
- Go installed (version 1.18 or later)
- Basic understanding of programming concepts

### Running Examples

Each file can be run independently:

```bash
# Run a specific example
go run variables.go

# Run with verbose output
go run -v functions.go

# Build and run
go build functions.go && ./functions
```

### Learning Path

#### 🌱 Beginner Level
1. **Start with basics:** `variables.go` → `constants.go` → `functions.go`
2. **Control flow:** `for.go` → `if-else.go` → `switch.go`
3. **Data structures:** `arrays.go` → `slices.go` → `maps.go`
4. **Function concepts:** `multiple-return-values.go` → `variadic-functions.go` → `closures.go` → `recursion.go`

#### 🚀 Intermediate Level
5. **Range operations:** `range-over-built-in-types.go`
6. **Memory management:** `pointers.go`
7. **Text processing:** `strings-and-runes.go`
8. **Data modeling:** `structs.go` → `methods.go` → `struct-embedding.go`
9. **Abstraction:** `interfaces.go` → `enums.go`

#### 🔥 Advanced Level
10. **Modern Go:** `generics.go` → `range-over-iterators.go`
11. **Error handling:** `errors.go` → `custom-errors.go`
12. **Concurrency:** `goroutines.go` → `channels.go` → `channel-buffering.go`

#### 📚 Quick Reference
- **Total Examples:** 28 files
- **Beginner:** 13 files
- **Intermediate:** 7 files  
- **Advanced:** 8 files
- **Estimated Learning Time:** 2-4 weeks

### Code Style

All examples follow Go conventions:
- Package naming: `package main`
- Clear function and variable names
- Comprehensive comments
- Error handling patterns
- Idiomatic Go code

## 📚 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go Tour](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go by Example (Official)](https://gobyexample.com/)

## 🤝 Contributing

Feel free to submit issues or enhancement requests!

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

---

**Happy Coding! 🎉**

Each example is designed to be self-contained and educational. Run them, modify them, and experiment to deepen your understanding of Go programming!
