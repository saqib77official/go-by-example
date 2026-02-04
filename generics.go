package main

import "fmt"

// 1. Basic generic function
func Print[T any](value T) {
	fmt.Printf("Value: %v (Type: %T)\n", value, value)
}

// 2. Generic function with multiple type parameters
func Pair[T, U any](first T, second U) {
	fmt.Printf("First: %v (%T), Second: %v (%T)\n", first, first, second, second)
}

// 3. Generic function with type constraint
type Number interface {
	int | int64 | float32 | float64
}

func Add[T Number](a, b T) T {
	return a + b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 4. Generic struct
type Container[T any] struct {
	data []T
}

func (c *Container[T]) Add(item T) {
	c.data = append(c.data, item)
}

func (c *Container[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index >= len(c.data) {
		return zero, fmt.Errorf("index out of bounds")
	}
	return c.data[index], nil
}

func (c *Container[T]) Size() int {
	return len(c.data)
}

// 5. Generic struct with type constraint
type Stack[T comparable] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, fmt.Errorf("stack is empty")
	}
	
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack[T]) Contains(item T) bool {
	for _, i := range s.items {
		if i == item {
			return true
		}
	}
	return false
}

// 6. Generic interface
type Comparable[T any] interface {
	Compare(other T) int
}

// 7. Generic function with interface constraint
type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

// 8. Generic struct with multiple type parameters
type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

func (kvp KeyValuePair[K, V]) String() string {
	return fmt.Sprintf("%v: %v", kvp.Key, kvp.Value)
}

// 9. Generic map type
type Dictionary[K comparable, V any] struct {
	data map[K]V
}

func NewDictionary[K comparable, V any]() *Dictionary[K, V] {
	return &Dictionary[K, V]{
		data: make(map[K]V),
	}
}

func (d *Dictionary[K, V]) Set(key K, value V) {
	d.data[key] = value
}

func (d *Dictionary[K, V]) Get(key K) (V, bool) {
	value, exists := d.data[key]
	return value, exists
}

func (d *Dictionary[K, V]) Delete(key K) {
	delete(d.data, key)
}

func (d *Dictionary[K, V]) Keys() []K {
	keys := make([]K, 0, len(d.data))
	for k := range d.data {
		keys = append(keys, k)
	}
	return keys
}

// 10. Generic function with pointer type constraint
type Pointer[T any] interface {
	*T
}

func UpdateValue[T any](ptr *T, newValue T) {
	*ptr = newValue
}

// 11. Generic slice operations
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T, U any](slice []T, transform func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = transform(item)
	}
	return result
}

func Reduce[T any](slice []T, initial T, reducer func(T, T) T) T {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// 12. Generic linked list
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Size int
}

func (ll *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

func (ll *LinkedList[T]) ToSlice() []T {
	var result []T
	current := ll.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// 13. Generic constraint with underlying types
type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func Sum[T Integer](values []T) T {
	var total T
	for _, v := range values {
		total += v
	}
	return total
}

// Custom types for demonstration
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

func (p Person) Compare(other Person) int {
	if p.Age < other.Age {
		return -1
	} else if p.Age > other.Age {
		return 1
	}
	return 0
}

type Product struct {
	Name  string
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf("%s: $%.2f", p.Name, p.Price)
}

func main() {
	fmt.Println("=== Generics Examples ===")

	// 1. Basic generic function
	fmt.Println("\n1. Basic generic function:")
	Print(42)
	Print("Hello, Generics!")
	Print(3.14)
	Print(Person{Name: "Alice", Age: 30})

	// 2. Generic function with multiple type parameters
	fmt.Println("\n2. Multiple type parameters:")
	Pair("Name", "Alice")
	Pair(100, 200)
	Pair("Age", 25)

	// 3. Generic function with type constraint
	fmt.Println("\n3. Type constraint functions:")
	fmt.Printf("Add(10, 20) = %d\n", Add(10, 20))
	fmt.Printf("Add(3.14, 2.86) = %.2f\n", Add(3.14, 2.86))
	fmt.Printf("Max(15, 25) = %d\n", Max(15, 25))
	fmt.Printf("Max(3.5, 2.8) = %.1f\n", Max(3.5, 2.8))

	// 4. Generic struct
	fmt.Println("\n4. Generic struct:")
	intContainer := Container[int]{}
	intContainer.Add(10)
	intContainer.Add(20)
	intContainer.Add(30)
	
	fmt.Printf("Int container size: %d\n", intContainer.Size())
	if value, err := intContainer.Get(1); err == nil {
		fmt.Printf("Value at index 1: %d\n", value)
	}
	
	stringContainer := Container[string]{}
	stringContainer.Add("Hello")
	stringContainer.Add("World")
	
	fmt.Printf("String container size: %d\n", stringContainer.Size())
	if value, err := stringContainer.Get(0); err == nil {
		fmt.Printf("Value at index 0: %s\n", value)
	}

	// 5. Generic struct with comparable constraint
	fmt.Println("\n5. Stack with comparable constraint:")
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	
	fmt.Printf("Stack contains 20: %t\n", intStack.Contains(20))
	
	if item, err := intStack.Pop(); err == nil {
		fmt.Printf("Popped: %d\n", item)
	}
	
	if item, err := intStack.Pop(); err == nil {
		fmt.Printf("Popped: %d\n", item)
	}

	// 6. Generic function with Stringer constraint
	fmt.Println("\n6. Stringer constraint:")
	PrintString(Person{Name: "Bob", Age: 25})
	PrintString(Product{Name: "Laptop", Price: 999.99})

	// 7. Generic struct with multiple type parameters
	fmt.Println("\n7. KeyValuePair:")
	kvp1 := KeyValuePair[string, int]{Key: "age", Value: 25}
	kvp2 := KeyValuePair[int, string]{Key: 1, Value: "first"}
	
	fmt.Printf("KVP1: %s\n", kvp1.String())
	fmt.Printf("KVP2: %s\n", kvp2.String())

	// 8. Generic dictionary
	fmt.Println("\n8. Generic Dictionary:")
	stringDict := NewDictionary[string, int]()
	stringDict.Set("apple", 5)
	stringDict.Set("banana", 3)
	stringDict.Set("orange", 8)
	
	if value, exists := stringDict.Get("apple"); exists {
		fmt.Printf("Apple count: %d\n", value)
	}
	
	fmt.Printf("All keys: %v\n", stringDict.Keys())
	
	intDict := NewDictionary[int, string]()
	intDict.Set(1, "first")
	intDict.Set(2, "second")
	
	if value, exists := intDict.Get(2); exists {
		fmt.Printf("Value for key 2: %s\n", value)
	}

	// 9. Generic slice operations
	fmt.Println("\n9. Generic slice operations:")
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Filter even numbers
	even := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", even)
	
	// Map to squares
	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Printf("Squares: %v\n", squares)
	
	// Reduce to sum
	sum := Reduce(numbers, 0, func(acc, n int) int { return acc + n })
	fmt.Printf("Sum: %d\n", sum)

	// 10. Generic linked list
	fmt.Println("\n10. Generic linked list:")
	stringList := LinkedList[string]{}
	stringList.Append("Hello")
	stringList.Append("World")
	stringList.Append("Generics")
	
	fmt.Printf("String list: %v\n", stringList.ToSlice())
	
	intList := LinkedList[int]{}
	intList.Append(10)
	intList.Append(20)
	intList.Append(30)
	
	fmt.Printf("Int list: %v\n", intList.ToSlice())

	// 11. Generic constraint with underlying types
	fmt.Println("\n11. Integer constraint:")
	var int8Values []int8 = {1, 2, 3, 4, 5}
	var uintValues []uint = {10, 20, 30}
	
	fmt.Printf("Sum of int8 values: %d\n", Sum(int8Values))
	fmt.Printf("Sum of uint values: %d\n", Sum(uintValues))

	// 12. Generic function with pointer
	fmt.Println("\n12. Generic pointer function:")
	x := 42
	y := "Hello"
	
	fmt.Printf("Before: x = %d, y = %s\n", x, y)
	UpdateValue(&x, 100)
	UpdateValue(&y, "World")
	fmt.Printf("After: x = %d, y = %s\n", x, y)

	// 13. Complex generic operations
	fmt.Println("\n13. Complex operations:")
	people := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Diana", Age: 28},
	}
	
	// Filter people older than 28
	olderPeople := Filter(people, func(p Person) bool { return p.Age > 28 })
	fmt.Printf("People older than 28: %v\n", olderPeople)
	
	// Map to names
	names := Map(people, func(p Person) string { return p.Name })
	fmt.Printf("All names: %v\n", names)
	
	// Find oldest person
	oldest := Reduce(people, people[0], func(older, current Person) Person {
		if current.Age > older.Age {
			return current
		}
		return older
	})
	fmt.Printf("Oldest person: %s\n", oldest.String())
}
