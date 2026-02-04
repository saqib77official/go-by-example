package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Interfaces Examples ===")

	// 1. Basic interface definition and implementation
	fmt.Println("\n1. Basic interface:")
	type Shape interface {
		Area() float64
		Perimeter() float64
	}
	
	type Rectangle struct {
		Width, Height float64
	}
	
	func (r Rectangle) Area() float64 {
		return r.Width * r.Height
	}
	
	func (r Rectangle) Perimeter() float64 {
		return 2 * (r.Width + r.Height)
	}
	
	type Circle struct {
		Radius float64
	}
	
	func (c Circle) Area() float64 {
		return math.Pi * c.Radius * c.Radius
	}
	
	func (c Circle) Perimeter() float64 {
		return 2 * math.Pi * c.Radius
	}
	
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	
	shapes := []Shape{rect, circle}
	for i, shape := range shapes {
		fmt.Printf("Shape %d - Area: %.2f, Perimeter: %.2f\n", 
			i+1, shape.Area(), shape.Perimeter())
	}

	// 2. Empty interface
	fmt.Println("\n2. Empty interface:")
	var data interface{} = 42
	fmt.Printf("Value: %v, Type: %T\n", data, data)
	
	data = "Hello, World!"
	fmt.Printf("Value: %v, Type: %T\n", data, data)
	
	data = []int{1, 2, 3}
	fmt.Printf("Value: %v, Type: %T\n", data, data)
	
	// Working with empty interface
	mixed := []interface{}{42, "hello", 3.14, true, []string{"a", "b"}}
	for i, item := range mixed {
		fmt.Printf("Item %d: %v (Type: %T)\n", i, item, item)
	}

	// 3. Type assertions
	fmt.Println("\n3. Type assertions:")
	var x interface{} = "Hello, Go!"
	
	// Safe type assertion
	if str, ok := x.(string); ok {
		fmt.Printf("String value: %s\n", str)
	} else {
		fmt.Println("Not a string")
	}
	
	// Type assertion that would panic
	// num := x.(int) // This would panic
	
	// Type switch
	x = 42
	switch v := x.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// 4. Interface composition
	fmt.Println("\n4. Interface composition:")
	type Writer interface {
		Write(data string) error
	}
	
	type Reader interface {
		Read() (string, error)
	}
	
	type ReadWriter interface {
		Reader
		Writer
	}
	
	type File struct {
		name    string
		content string
	}
	
	func (f *File) Write(data string) error {
		f.content += data
		return nil
	}
	
	func (f *File) Read() (string, error) {
		return f.content, nil
	}
	
	file := &File{name: "test.txt"}
	var rw ReadWriter = file
	
	rw.Write("Hello, ")
	rw.Write("World!")
	
	if content, err := rw.Read(); err == nil {
		fmt.Printf("File content: %s\n", content)
	}

	// 5. Interface with methods returning interfaces
	fmt.Println("\n5. Methods returning interfaces:")
	type Animal interface {
		Speak() string
	}
	
	type Dog struct {
		Name string
	}
	
	func (d Dog) Speak() string {
		return fmt.Sprintf("%s says Woof!", d.Name)
	}
	
	type Cat struct {
		Name string
	}
	
	func (c Cat) Speak() string {
		return fmt.Sprintf("%s says Meow!", c.Name)
	}
	
	func CreateAnimal(animalType, name string) Animal {
		switch animalType {
		case "dog":
			return Dog{Name: name}
		case "cat":
			return Cat{Name: name}
		default:
			return nil
		}
	}
	
	animals := []Animal{
		CreateAnimal("dog", "Buddy"),
		CreateAnimal("cat", "Whiskers"),
	}
	
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}

	// 6. Interface as function parameters
	fmt.Println("\n6. Interface as function parameters:")
	type Logger interface {
		Log(message string)
	}
	
	type ConsoleLogger struct{}
	
	func (cl ConsoleLogger) Log(message string) {
		fmt.Printf("CONSOLE: %s\n", message)
	}
	
	type FileLogger struct {
		filename string
	}
	
	func (fl FileLogger) Log(message string) {
		fmt.Printf("FILE[%s]: %s\n", fl.filename, message)
	}
	
	func ProcessData(data string, logger Logger) {
		logger.Log("Starting data processing")
		// Simulate processing
		logger.Log("Data processed successfully")
	}
	
	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{filename: "app.log"}
	
	ProcessData("sample data", consoleLogger)
	ProcessData("sample data", fileLogger)

	// 7. Interface nil values
	fmt.Println("\n7. Interface nil values:")
	var nilInterface Shape
	fmt.Printf("Nil interface value: %v\n", nilInterface)
	fmt.Printf("Is nil interface nil? %t\n", nilInterface == nil)
	
	// Non-nil interface with nil concrete value
	var nilRect *Rectangle
	var shapeInterface Shape = nilRect
	fmt.Printf("Interface with nil concrete: %v\n", shapeInterface)
	fmt.Printf("Is interface nil? %t\n", shapeInterface == nil)
	fmt.Printf("Is concrete value nil? %t\n", nilRect == nil)

	// 8. Interface comparison
	fmt.Println("\n8. Interface comparison:")
	rect1 := Rectangle{Width: 10, Height: 5}
	rect2 := Rectangle{Width: 10, Height: 5}
	circle1 := Circle{Radius: 5}
	
	var shape1, shape2, shape3 Shape
	shape1 = rect1
	shape2 = rect2
	shape3 = circle1
	
	fmt.Printf("shape1 == shape2: %t\n", shape1 == shape2) // true (same concrete values)
	fmt.Printf("shape1 == shape3: %t\n", shape1 == shape3) // false (different types)
	
	// Interfaces with uncomparable types cannot be compared
	// type SliceStruct struct { data []int }
	// var iface1, iface2 interface{} = SliceStruct{[]int{1}}, SliceStruct{[]int{1}}
	// iface1 == iface2 // This would panic

	// 9. Interface embedding
	fmt.Println("\n9. Interface embedding:")
	type Closer interface {
		Close() error
	}
	
	type ReadCloser interface {
		Reader
		Closer
	}
	
	type Buffer struct {
		data []byte
		pos  int
	}
	
	func (b *Buffer) Read() (string, error) {
		if b.pos >= len(b.data) {
			return "", fmt.Errorf("EOF")
		}
		result := string(b.data[b.pos])
		b.pos++
		return result, nil
	}
	
	func (b *Buffer) Write(data string) error {
		b.data = append(b.data, []byte(data)...)
		return nil
	}
	
	func (b *Buffer) Close() error {
		b.pos = 0
		b.data = nil
		return nil
	}
	
	buffer := &Buffer{}
	var rc ReadCloser = buffer
	
	rc.Write("Hello")
	rc.Write(" World!")
	
	for {
		if char, err := rc.Read(); err != nil {
			break
		} else {
			fmt.Printf("Read: %s\n", char)
		}
	}
	
	rc.Close()

	// 10. Type constraints in interfaces
	fmt.Println("\n10. Type constraints:")
	type Comparable interface {
		~int | ~float64 | ~string
	}
	
	func Max[T Comparable](a, b T) T {
		if a > b {
			return a
		}
		return b
	}
	
	fmt.Printf("Max of 10 and 20: %d\n", Max(10, 20))
	fmt.Printf("Max of 3.14 and 2.71: %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max of 'apple' and 'banana': %s\n", Max("apple", "banana"))

	// 11. Interface with pointer receivers
	fmt.Println("\n11. Interface with pointer receivers:")
	type Counter interface {
		Increment()
		GetValue() int
	}
	
	type SimpleCounter struct {
		value int
	}
	
	func (sc *SimpleCounter) Increment() {
		sc.value++
	}
	
	func (sc SimpleCounter) GetValue() int {
		return sc.value
	}
	
	counter := &SimpleCounter{value: 0}
	var c Counter = counter
	
	c.Increment()
	c.Increment()
	fmt.Printf("Counter value: %d\n", c.GetValue())

	// 12. Dynamic interface implementation
	fmt.Println("\n12. Dynamic interface implementation:")
	type Validator interface {
		Validate(value interface{}) bool
	}
	
	type StringValidator struct {
		minLength int
	}
	
	func (sv StringValidator) Validate(value interface{}) bool {
		if str, ok := value.(string); ok {
			return len(str) >= sv.minLength
		}
		return false
	}
	
	type NumberValidator struct {
		min, max float64
	}
	
	func (nv NumberValidator) Validate(value interface{}) bool {
		switch v := value.(type) {
		case int:
			return float64(v) >= nv.min && float64(v) <= nv.max
		case float64:
			return v >= nv.min && v <= nv.max
		default:
			return false
		}
	}
	
	validators := []Validator{
		StringValidator{minLength: 5},
		NumberValidator{min: 0, max: 100},
	}
	
	testValues := []interface{}{"hello", "hi", 50, 150, 3.14}
	
	for _, value := range testValues {
		fmt.Printf("Validating %v (%T):\n", value, value)
		for i, validator := range validators {
			isValid := validator.Validate(value)
			fmt.Printf("  Validator %d: %t\n", i+1, isValid)
		}
	}

	// 13. Interface as abstraction layer
	fmt.Println("\n13. Interface as abstraction layer:")
	type Database interface {
		Save(key string, value interface{}) error
		Get(key string) (interface{}, error)
	}
	
	type MemoryDatabase struct {
		data map[string]interface{}
	}
	
	func NewMemoryDatabase() *MemoryDatabase {
		return &MemoryDatabase{
			data: make(map[string]interface{}),
		}
	}
	
	func (md *MemoryDatabase) Save(key string, value interface{}) error {
		md.data[key] = value
		return nil
	}
	
	func (md *MemoryDatabase) Get(key string) (interface{}, error) {
		if value, exists := md.data[key]; exists {
			return value, nil
		}
		return nil, fmt.Errorf("key not found: %s", key)
	}
	
	type Service struct {
		db Database
	}
	
	func NewService(db Database) *Service {
		return &Service{db: db}
	}
	
	func (s *Service) StoreUserData(userID string, userData interface{}) error {
		return s.db.Save(userID, userData)
	}
	
	func (s *Service) GetUserData(userID string) (interface{}, error) {
		return s.db.Get(userID)
	}
	
	// Using the service with memory database
	service := NewService(NewMemoryDatabase())
	
	userData := map[string]interface{}{
		"name":  "John Doe",
		"email": "john@example.com",
		"age":   30,
	}
	
	service.StoreUserData("user123", userData)
	
	if retrieved, err := service.GetUserData("user123"); err == nil {
		fmt.Printf("Retrieved user data: %+v\n", retrieved)
	}
}
