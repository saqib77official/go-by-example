package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Methods Examples ===")

	// 1. Basic method with value receiver
	fmt.Println("\n1. Value receiver methods:")
	type Rectangle struct {
		Width  float64
		Height float64
	}
	
	// Value receiver method
	func (r Rectangle) Area() float64 {
		return r.Width * r.Height
	}
	
	// Another value receiver method
	func (r Rectangle) Perimeter() float64 {
		return 2 * (r.Width + r.Height)
	}
	
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())

	// 2. Method with pointer receiver
	fmt.Println("\n2. Pointer receiver methods:")
	
	// Pointer receiver method
	func (r *Rectangle) Scale(factor float64) {
		r.Width *= factor
		r.Height *= factor
	}
	
	// Another pointer receiver method
	func (r *Rectangle) SetDimensions(width, height float64) {
		r.Width = width
		r.Height = height
	}
	
	fmt.Printf("Before scaling: %+v\n", rect)
	rect.Scale(2)
	fmt.Printf("After scaling by 2: %+v\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())
	
	rect.SetDimensions(15, 8)
	fmt.Printf("After setting dimensions: %+v\n", rect)

	// 3. Value vs Pointer receiver behavior
	fmt.Println("\n3. Value vs Pointer receiver behavior:")
	type Counter struct {
		count int
	}
	
	// Value receiver - doesn't modify original
	func (c Counter) GetValue() int {
		return c.count
	}
	
	// Pointer receiver - modifies original
	func (c *Counter) Increment() {
		c.count++
	}
	
	// Value receiver - creates copy
	func (c Counter) IncrementAndReturn() int {
		c.count++
		return c.count
	}
	
	counter := Counter{count: 0}
	fmt.Printf("Initial count: %d\n", counter.GetValue())
	
	counter.Increment()
	fmt.Printf("After Increment(): %d\n", counter.GetValue())
	
	result := counter.IncrementAndReturn()
	fmt.Printf("IncrementAndReturn() result: %d\n", result)
	fmt.Printf("Count after IncrementAndReturn(): %d\n", counter.GetValue()) // Still 1

	// 4. Method expressions and values
	fmt.Println("\n4. Method expressions and values:")
	
	// Method expression
	areaFunc := Rectangle.Area
	perimeterFunc := (*Rectangle).Perimeter
	
	fmt.Printf("Area function result: %.2f\n", areaFunc(rect))
	fmt.Printf("Perimeter function result: %.2f\n", perimeterFunc(&rect))
	
	// Method value
	areaMethod := rect.Area
	fmt.Printf("Area method result: %.2f\n", areaMethod())

	// 5. Methods on non-struct types
	fmt.Println("\n5. Methods on non-struct types:")
	type MyInt int
	
	// Method on basic type
	func (m MyInt) Double() MyInt {
		return m * 2
	}
	
	func (m MyInt) String() string {
		return fmt.Sprintf("MyInt(%d)", m)
	}
	
	var num MyInt = 42
	fmt.Printf("Original: %s\n", num.String())
	fmt.Printf("Double: %s\n", num.Double().String())

	// 6. Methods with interface types
	fmt.Println("\n6. Methods with interface types:")
	type Shape interface {
		Area() float64
		Perimeter() float64
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
	
	circle := Circle{Radius: 5}
	shapes := []Shape{rect, circle}
	
	fmt.Printf("Rectangle area: %.2f\n", shapes[0].Area())
	fmt.Printf("Circle area: %.2f\n", shapes[1].Area())

	// 7. Method chaining
	fmt.Println("\n7. Method chaining:")
	type StringBuilder struct {
		data string
	}
	
	func (sb *StringBuilder) Append(s string) *StringBuilder {
		sb.data += s
		return sb
	}
	
	func (sb *StringBuilder) ToUpper() *StringBuilder {
		sb.data = fmt.Sprintf("%s", sb.data)
		return sb
	}
	
	func (sb *StringBuilder) String() string {
		return sb.data
	}
	
	result := new(StringBuilder).
		Append("Hello").
		Append(" ").
		Append("World").
		String()
	
	fmt.Printf("Chained result: %s\n", result)

	// 8. Methods with variadic parameters
	fmt.Println("\n8. Methods with variadic parameters:")
	type Calculator struct {
		result float64
	}
	
	func (c *Calculator) Add(numbers ...float64) *Calculator {
		for _, num := range numbers {
			c.result += num
		}
		return c
	}
	
	func (c *Calculator) Multiply(numbers ...float64) *Calculator {
		for _, num := range numbers {
			c.result *= num
		}
		return c
	}
	
	func (c *Calculator) Reset() *Calculator {
		c.result = 0
		return c
	}
	
	func (c *Calculator) Result() float64 {
		return c.result
	}
	
	calc := new(Calculator)
	result := calc.Reset().Add(10, 20, 30).Multiply(2).Result()
	fmt.Printf("Calculation result: %.2f\n", result)

	// 9. Methods returning multiple values
	fmt.Println("\n9. Methods returning multiple values:")
	type Point struct {
		X, Y float64
	}
	
	func (p Point) Distance(other Point) (float64, error) {
		if p.X == 0 && p.Y == 0 {
			return 0, fmt.Errorf("invalid point")
		}
		dx := p.X - other.X
		dy := p.Y - other.Y
		return math.Sqrt(dx*dx + dy*dy), nil
	}
	
	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 3, Y: 4}
	
	if distance, err := p1.Distance(p2); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Distance: %.2f\n", distance)
	}

	// 10. Methods with embedded types
	fmt.Println("\n10. Methods with embedded types:")
	type Animal struct {
		Name string
	}
	
	func (a Animal) Speak() string {
		return fmt.Sprintf("%s makes a sound", a.Name)
	}
	
	type Dog struct {
		Animal
		Breed string
	}
	
	// Override Speak method
	func (d Dog) Speak() string {
		return fmt.Sprintf("%s barks", d.Name)
	}
	
	// New method specific to Dog
	func (d Dog) WagTail() string {
		return fmt.Sprintf("%s wags tail happily", d.Name)
	}
	
	animal := Animal{Name: "Generic Animal"}
	dog := Dog{
		Animal: Animal{Name: "Buddy"},
		Breed:  "Golden Retriever",
	}
	
	fmt.Printf("Animal: %s\n", animal.Speak())
	fmt.Printf("Dog: %s\n", dog.Speak())
	fmt.Printf("Dog action: %s\n", dog.WagTail())
	fmt.Printf("Dog using Animal method: %s\n", dog.Animal.Speak())

	// 11. Method sets
	fmt.Println("\n11. Method sets:")
	type Writer interface {
		Write(data string)
	}
	
	type FileWriter struct {
		filename string
	}
	
	func (fw FileWriter) Write(data string) {
		fmt.Printf("Writing '%s' to file %s\n", data, fw.filename)
	}
	
	func (fw *FileWriter) Close() {
		fmt.Printf("Closing file %s\n", fw.filename)
	}
	
	// FileWriter value has Write method
	// *FileWriter pointer has both Write and Close methods
	
	fw := FileWriter{filename: "test.txt"}
	var writer Writer = fw // OK: fw implements Writer
	
	writer.Write("Hello, World!")
	
	// fw.Close() // Error: fw doesn't have Close method
	// (&fw).Close() // OK: pointer has Close method

	// 12. Methods with receiver as interface
	fmt.Println("\n12. Methods with receiver as interface:")
	type Processor interface {
		Process(data string) string
	}
	
	type UppercaseProcessor struct{}
	
	func (up UppercaseProcessor) Process(data string) string {
		return fmt.Sprintf("UPPERCASE: %s", data)
	}
	
	type LowercaseProcessor struct{}
	
	func (lp LowercaseProcessor) Process(data string) string {
		return fmt.Sprintf("lowercase: %s", data)
	}
	
	type Manager struct {
		processors []Processor
	}
	
	func (m *Manager) AddProcessor(p Processor) {
		m.processors = append(m.processors, p)
	}
	
	func (m *Manager) ProcessAll(data string) []string {
		var results []string
		for _, processor := range m.processors {
			results = append(results, processor.Process(data))
		}
		return results
	}
	
	manager := Manager{}
	manager.AddProcessor(UppercaseProcessor{})
	manager.AddProcessor(LowercaseProcessor{})
	
	results := manager.ProcessAll("Hello World")
	fmt.Printf("Processing results: %v\n", results)

	// 13. Method visibility and naming conventions
	fmt.Println("\n13. Method visibility:")
	type PublicStruct struct {
		publicField  string
		privateField string
	}
	
	// Public method (starts with uppercase)
	func (ps PublicStruct) PublicMethod() string {
		return "This is a public method"
	}
	
	// Private method (starts with lowercase)
	func (ps PublicStruct) privateMethod() string {
		return "This is a private method"
	}
	
	ps := PublicStruct{publicField: "public", privateField: "private"}
	fmt.Printf("Public method: %s\n", ps.PublicMethod())
	// ps.privateMethod() // Error: cannot call private method from outside package
}
