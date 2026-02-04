package main

import "fmt"

func main() {
	fmt.Println("=== Struct Embedding Examples ===")

	// 1. Basic struct embedding
	fmt.Println("\n1. Basic struct embedding:")
	type Person struct {
		Name string
		Age  int
	}
	
	type Employee struct {
		Person     // Embedded struct
		EmployeeID int
		Department string
	}
	
	emp := Employee{
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		EmployeeID: 1001,
		Department: "Engineering",
	}
	
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Name (promoted field): %s\n", emp.Name) // Access Person.Name directly
	fmt.Printf("Age (promoted field): %d\n", emp.Age)    // Access Person.Age directly
	fmt.Printf("EmployeeID: %d\n", emp.EmployeeID)
	fmt.Printf("Department: %s\n", emp.Department)
	fmt.Printf("Person struct: %+v\n", emp.Person) // Access embedded struct directly

	// 2. Multiple embedded structs
	fmt.Println("\n2. Multiple embedded structs:")
	type Contact struct {
		Email string
		Phone string
	}
	
	type Address struct {
		Street  string
		City    string
		Country string
	}
	
	type FullProfile struct {
		Person  // Embedded
		Contact // Embedded
		Address // Embedded
		Website string
	}
	
	profile := FullProfile{
		Person: Person{
			Name: "Jane Smith",
			Age:  28,
		},
		Contact: Contact{
			Email: "jane@example.com",
			Phone: "+1-555-0123",
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "New York",
			Country: "USA",
		},
		Website: "https://janesmith.com",
	}
	
	fmt.Printf("Full Profile: %+v\n", profile)
	fmt.Printf("Name: %s\n", profile.Name)
	fmt.Printf("Email: %s\n", profile.Email)
	fmt.Printf("City: %s\n", profile.City)

	// 3. Method promotion from embedded structs
	fmt.Println("\n3. Method promotion:")
	
	// Add methods to Person
	func (p Person) Greet() string {
		return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
	}
	
	func (p Person) IsAdult() bool {
		return p.Age >= 18
	}
	
	// Add methods to Contact
	func (c Contact) GetContactInfo() string {
		return fmt.Sprintf("Email: %s, Phone: %s", c.Email, c.Phone)
	}
	
	// Employee now has access to Person and Contact methods
	emp2 := Employee{
		Person: Person{
			Name: "Alice Johnson",
			Age:  25,
		},
		EmployeeID: 2002,
		Department: "Marketing",
	}
	
	fmt.Printf("Greeting: %s\n", emp2.Greet())
	fmt.Printf("Is adult: %t\n", emp2.IsAdult())

	// 4. Method overriding with embedding
	fmt.Println("\n4. Method overriding:")
	type Manager struct {
		Employee
		TeamSize int
	}
	
	// Override Greet method
	func (m Manager) Greet() string {
		return fmt.Sprintf("Hello! I'm %s, a manager with %d team members", m.Name, m.TeamSize)
	}
	
	manager := Manager{
		Employee: Employee{
			Person: Person{
				Name: "Bob Wilson",
				Age:  35,
			},
			EmployeeID: 3003,
			Department: "Sales",
		},
		TeamSize: 8,
	}
	
	fmt.Printf("Manager greeting: %s\n", manager.Greet())
	fmt.Printf("Employee greeting (accessing embedded): %s\n", manager.Employee.Greet())

	// 5. Name conflicts with embedding
	fmt.Println("\n5. Name conflicts:")
	type A struct {
		Value string
	}
	
	type B struct {
		Value int
	}
	
	type C struct {
		A
		B
	}
	
	c := C{
		A: A{Value: "string value"},
		B: B{Value: 42},
	}
	
	// Ambiguous access - would cause compile error
	// fmt.Printf("Value: %v\n", c.Value)
	
	// Must specify which embedded field
	fmt.Printf("A.Value: %s\n", c.A.Value)
	fmt.Printf("B.Value: %d\n", c.B.Value)

	// 6. Pointer embedding
	fmt.Println("\n6. Pointer embedding:")
	type Engine struct {
		Type  string
		Power int
	}
	
	func (e *Engine) Start() {
		fmt.Printf("Starting %s engine with %d HP\n", e.Type, e.Power)
	}
	
	func (e *Engine) Stop() {
		fmt.Printf("Stopping %s engine\n", e.Type)
	}
	
	type Car struct {
		Make   string
		Model  string
		Engine *Engine // Pointer embedding
	}
	
	car := Car{
		Make:  "Toyota",
		Model: "Camry",
		Engine: &Engine{
			Type:  "V6",
			Power: 300,
		},
	}
	
	fmt.Printf("Car: %s %s\n", car.Make, car.Model)
	car.Start() // Method promotion works with pointers
	car.Stop()

	// 7. Anonymous struct embedding
	fmt.Println("\n7. Anonymous struct embedding:")
	type Product struct {
		Name  string
		Price float64
		struct {
			Category string
			Brand    string
		}
	}
	
	product := Product{
		Name:  "Laptop",
		Price: 999.99,
		struct {
			Category string
			Brand    string
		}{
			Category: "Electronics",
			Brand:    "TechCo",
		},
	}
	
	fmt.Printf("Product: %+v\n", product)
	fmt.Printf("Category: %s\n", product.Category)
	fmt.Printf("Brand: %s\n", product.Brand)

	// 8. Interface embedding
	fmt.Println("\n8. Interface embedding:")
	type Writer interface {
		Write(data string) error
	}
	
	type Reader interface {
		Read() (string, error)
	}
	
	type ReadWriter interface {
		Writer
		Reader
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

	// 9. Embedding with composition patterns
	fmt.Println("\n9. Embedding with composition patterns:")
	type Logger struct {
		logLevel string
	}
	
	func (l *Logger) Log(message string) {
		fmt.Printf("[%s] %s\n", l.logLevel, message)
	}
	
	func (l *Logger) SetLevel(level string) {
		l.logLevel = level
	}
	
	type Service struct {
		Name   string
		Logger // Embedded logger
	}
	
	func (s *Service) Start() {
		s.Log(fmt.Sprintf("Starting service: %s", s.Name))
		// Service logic here
		s.Log("Service started successfully")
	}
	
	service := Service{
		Name: "UserService",
		Logger: Logger{
			logLevel: "INFO",
		},
	}
	
	service.Start()
	service.SetLevel("DEBUG")
	service.Log("Debug message")

	// 10. Embedding for behavior extension
	fmt.Println("\n10. Embedding for behavior extension:")
	type Animal struct {
		Name string
	}
	
	func (a Animal) Speak() string {
		return fmt.Sprintf("%s makes a sound", a.Name)
	}
	
	type Flyable interface {
		Fly() string
	}
	
	type Bird struct {
		Animal
		Wingspan int
	}
	
	func (b Bird) Speak() string {
		return fmt.Sprintf("%s chirps", b.Name)
	}
	
	func (b Bird) Fly() string {
		return fmt.Sprintf("%s flies with %d cm wingspan", b.Name, b.Wingspan)
	}
	
	type Eagle struct {
		Bird
		HuntingSkill string
	}
	
	func (e Eagle) Speak() string {
		return fmt.Sprintf("%s screeches", e.Name)
	}
	
	eagle := Eagle{
		Bird: Bird{
			Animal:    Animal{Name: "Golden Eagle"},
			Wingspan: 200,
		},
		HuntingSkill: "Excellent",
	}
	
	fmt.Printf("Eagle speaks: %s\n", eagle.Speak())
	fmt.Printf("Eagle flies: %s\n", eagle.Fly())
	fmt.Printf("Animal speak: %s\n", eagle.Animal.Speak())

	// 11. Embedding with validation
	fmt.Println("\n11. Embedding with validation:")
	type Timestamp struct {
		CreatedAt string
		UpdatedAt string
	}
	
	func (t *Timestamp) Touch() {
		t.UpdatedAt = "2023-12-25T10:00:00Z" // In real app, use current time
	}
	
	type AuditableModel struct {
		ID        int
		Timestamp // Embedded
	}
	
	func (am *AuditableModel) Validate() error {
		if am.ID <= 0 {
			return fmt.Errorf("invalid ID: %d", am.ID)
		}
		if am.CreatedAt == "" {
			return fmt.Errorf("created_at is required")
		}
		return nil
	}
	
	model := AuditableModel{
		ID: 1,
		Timestamp: Timestamp{
			CreatedAt: "2023-12-25T09:00:00Z",
			UpdatedAt: "2023-12-25T09:00:00Z",
		},
	}
	
	if err := model.Validate(); err != nil {
		fmt.Printf("Validation error: %v\n", err)
	} else {
		fmt.Printf("Model is valid: %+v\n", model)
	}
	
	model.Touch()
	fmt.Printf("After touch: %+v\n", model.Timestamp)

	// 12. Embedding in slices and maps
	fmt.Println("\n12. Embedding in slices and maps:")
	type BaseItem struct {
		ID    int
		Name  string
		Price float64
	}
	
	type Book struct {
		BaseItem
		Author string
		Pages  int
	}
	
	type Electronics struct {
		BaseItem
		Brand     string
		Warranty  int
	}
	
	items := []interface{}{
		Book{
			BaseItem: BaseItem{ID: 1, Name: "Go Programming", Price: 29.99},
			Author:   "John Doe",
			Pages:    300,
		},
		Electronics{
			BaseItem: BaseItem{ID: 2, Name: "Laptop", Price: 999.99},
			Brand:    "TechCo",
			Warranty: 24,
		},
	}
	
	for i, item := range items {
		switch v := item.(type) {
		case Book:
			fmt.Printf("Book %d: %s by %s, $%.2f\n", i+1, v.Name, v.Author, v.Price)
		case Electronics:
			fmt.Printf("Electronics %d: %s by %s, $%.2f\n", i+1, v.Name, v.Brand, v.Price)
		}
	}
}
