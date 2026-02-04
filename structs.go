package main

import "fmt"

func main() {
	fmt.Println("=== Structs Examples ===")

	// 1. Basic struct definition and usage
	fmt.Println("\n1. Basic struct definition:")
	type Person struct {
		Name string
		Age  int
		City string
	}
	
	// Creating struct instances
	var person1 Person
	person1.Name = "Alice"
	person1.Age = 25
	person1.City = "New York"
	
	person2 := Person{Name: "Bob", Age: 30, City: "Los Angeles"}
	person3 := Person{"Charlie", 35, "Chicago"} // Positional initialization
	
	fmt.Printf("Person1: %+v\n", person1)
	fmt.Printf("Person2: %+v\n", person2)
	fmt.Printf("Person3: %+v\n", person3)

	// 2. Struct with different field types
	fmt.Println("\n2. Struct with different field types:")
	type Employee struct {
		ID        int
		Name      string
		Salary    float64
		IsActive  bool
		Tags      []string
	}
	
	emp := Employee{
		ID:       1001,
		Name:     "John Doe",
		Salary:   75000.50,
		IsActive: true,
		Tags:     []string{"developer", "senior"},
	}
	
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Name: %s, Salary: %.2f\n", emp.Name, emp.Salary)

	// 3. Struct with nested structs
	fmt.Println("\n3. Nested structs:")
	type Address struct {
		Street  string
		City    string
		State   string
		ZipCode string
	}
	
	type PersonWithAddress struct {
		Name    string
		Age     int
		Address Address
	}
	
	person := PersonWithAddress{
		Name: "Jane Smith",
		Age:  28,
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			State:   "MA",
			ZipCode: "02101",
		},
	}
	
	fmt.Printf("Person: %+v\n", person)
	fmt.Printf("Address: %s, %s, %s %s\n", 
		person.Address.Street, person.Address.City, 
		person.Address.State, person.Address.ZipCode)

	// 4. Struct pointers
	fmt.Println("\n4. Struct pointers:")
	personPtr := &Person{Name: "Mike", Age: 40, City: "Seattle"}
	
	fmt.Printf("Pointer: %p\n", personPtr)
	fmt.Printf("Dereferenced: %+v\n", *personPtr)
	fmt.Printf("Field access: %s\n", personPtr.Name) // Automatic dereferencing
	
	// Modifying through pointer
	personPtr.Age = 41
	fmt.Printf("After modification: %+v\n", *personPtr)

	// 5. Struct comparison
	fmt.Println("\n5. Struct comparison:")
	p1 := Person{Name: "Alice", Age: 25, City: "NYC"}
	p2 := Person{Name: "Alice", Age: 25, City: "NYC"}
	p3 := Person{Name: "Alice", Age: 26, City: "NYC"}
	
	fmt.Printf("p1 == p2: %t\n", p1 == p2) // true
	fmt.Printf("p1 == p3: %t\n", p1 == p3) // false
	
	// Structs with slices cannot be compared directly
	type PersonWithTags struct {
		Name string
		Tags []string
	}
	
	pt1 := PersonWithTags{Name: "Alice", Tags: []string{"developer"}}
	pt2 := PersonWithTags{Name: "Alice", Tags: []string{"developer"}}
	// pt1 == pt2 would cause compile error

	// 6. Anonymous structs
	fmt.Println("\n6. Anonymous structs:")
	anonymous := struct {
		Name string
		Age  int
	}{
		Name: "Anonymous User",
		Age:  99,
	}
	
	fmt.Printf("Anonymous struct: %+v\n", anonymous)
	
	// Slice of anonymous structs
	people := []struct {
		Name string
		Age  int
	}{
		{"Person A", 25},
		{"Person B", 30},
		{"Person C", 35},
	}
	
	fmt.Printf("People slice: %+v\n", people)

	// 7. Struct with methods (will be covered more in methods.go)
	fmt.Println("\n7. Basic struct methods:")
	type Rectangle struct {
		Width  float64
		Height float64
	}
	
	// Method with value receiver
	func (r Rectangle) Area() float64 {
		return r.Width * r.Height
	}
	
	// Method with pointer receiver
	func (r *Rectangle) Scale(factor float64) {
		r.Width *= factor
		r.Height *= factor
	}
	
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	
	rect.Scale(2)
	fmt.Printf("After scaling: %+v\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())

	// 8. Struct tags
	fmt.Println("\n8. Struct tags:")
	type User struct {
		ID       int    `json:"id" db:"user_id"`
		Username string `json:"username" db:"username"`
		Email    string `json:"email" db:"email"`
		Active   bool   `json:"active" db:"is_active"`
	}
	
	user := User{
		ID:       1,
		Username: "johndoe",
		Email:    "john@example.com",
		Active:   true,
	}
	
	fmt.Printf("User: %+v\n", user)
	// In real applications, tags are used by encoding/json, database libraries, etc.

	// 9. Zero values of structs
	fmt.Println("\n9. Zero values:")
	var zeroPerson Person
	var zeroEmployee Employee
	
	fmt.Printf("Zero Person: %+v\n", zeroPerson)
	fmt.Printf("Zero Employee: %+v\n", zeroEmployee)
	
	// Checking if struct is zero value
	isZero := zeroPerson.Name == "" && zeroPerson.Age == 0
	fmt.Printf("Is zeroPerson zero value? %t\n", isZero)

	// 10. Struct copying
	fmt.Println("\n10. Struct copying:")
	original := Person{Name: "Original", Age: 25, City: "Original City"}
	copy := original
	
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copy: %+v\n", copy)
	
	// Modifying copy doesn't affect original
	copy.Name = "Modified"
	fmt.Printf("After modifying copy:\n")
	fmt.Printf("Original: %+v\n", original)
	fmt.Printf("Copy: %+v\n", copy)

	// 11. Struct as function parameters and return values
	fmt.Println("\n11. Structs in functions:")
	
	// Pass by value
	func printPerson(p Person) {
		fmt.Printf("Person: %+v\n", p)
	}
	
	// Pass by pointer
	func updateAge(p *Person, newAge int) {
		p.Age = newAge
	}
	
	// Return struct
	func createPerson(name string, age int) Person {
		return Person{Name: name, Age: age, City: "Unknown"}
	}
	
	newPerson := createPerson("New Person", 22)
	printPerson(newPerson)
	
	updateAge(&newPerson, 23)
	printPerson(newPerson)

	// 12. Struct composition basics
	fmt.Println("\n12. Struct composition:")
	type Engine struct {
		Type  string
		Power int
	}
	
	type Car struct {
		Make    string
		Model   string
		Year    int
		Engine  // Embedded struct
	}
	
	car := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2022,
		Engine: Engine{
			Type:  "V6",
			Power: 300,
		},
	}
	
	fmt.Printf("Car: %+v\n", car)
	fmt.Printf("Engine type: %s\n", car.Engine.Type)
	fmt.Printf("Engine type (promoted): %s\n", car.Type) // Promoted field

	// 13. Struct with interface fields
	fmt.Println("\n13. Struct with interface fields:")
	type Logger interface {
		Log(message string)
	}
	
	type ConsoleLogger struct{}
	
	func (cl ConsoleLogger) Log(message string) {
		fmt.Printf("LOG: %s\n", message)
	}
	
	type Service struct {
		Name   string
		Logger Logger
	}
	
	service := Service{
		Name:   "MyService",
		Logger: ConsoleLogger{},
	}
	
	service.Logger.Log("Service started")
	fmt.Printf("Service: %+v\n", service)
}
