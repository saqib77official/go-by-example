# Go by Example

A comprehensive collection of Go programming examples covering fundamental concepts and language features. Each file contains practical examples with detailed explanations to help you learn Go programming.

## 📁 Files Overview

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

## 🚀 Getting Started

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

1. **Start with basics:** `variables.go` → `constants.go` → `functions.go`
2. **Control flow:** `for.go` → `if-else.go` → `switch.go`
3. **Data structures:** `arrays.go` → `slices.go` → `maps.go`
4. **Advanced concepts:** `multiple-return-values.go` → `variadic-functions.go` → `closures.go` → `recursion.go`

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
