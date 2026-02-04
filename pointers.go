package main

import "fmt"

func main() {
	fmt.Println("=== Pointers Examples ===")

	// 1. Basic pointer declaration and usage
	fmt.Println("\n1. Basic pointer operations:")
	var x int = 42
	var p *int = &x // p is a pointer to x
	
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("Value of p (address of x): %p\n", p)
	fmt.Printf("Value pointed to by p: %d\n", *p)

	// 2. Pointer dereferencing
	fmt.Println("\n2. Pointer dereferencing:")
	*p = 100 // Modify x through pointer
	fmt.Printf("After *p = 100, x = %d\n", x)
	
	y := *p // Copy value from pointer
	fmt.Printf("y = *p, y = %d\n", y)

	// 3. Pointer to pointer
	fmt.Println("\n3. Pointer to pointer:")
	var pp **int = &p // pp is a pointer to p
	fmt.Printf("Value of pp (address of p): %p\n", pp)
	fmt.Printf("Value pointed to by pp: %p\n", *pp)
	fmt.Printf("Value pointed to by *pp: %d\n", **pp)

	// 4. Nil pointers
	fmt.Println("\n4. Nil pointers:")
	var nilPtr *int
	fmt.Printf("Nil pointer value: %p\n", nilPtr)
	fmt.Printf("Is nil pointer nil? %t\n", nilPtr == nil)
	
	// This would cause a panic: *nilPtr = 42
	// fmt.Printf("Dereferencing nil pointer: %d\n", *nilPtr)

	// 5. Pointers with functions
	fmt.Println("\n5. Pointers with functions:")
	a := 10
	b := 20
	
	fmt.Printf("Before swap: a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap: a = %d, b = %d\n", a, b)
	
	// Function returning pointer
	ptr := createPointer(99)
	fmt.Printf("Function returned pointer: %p, value: %d\n", ptr, *ptr)

	// 6. Pointers and structs
	fmt.Println("\n6. Pointers and structs:")
	type Person struct {
		Name string
		Age  int
	}
	
	person := Person{Name: "Alice", Age: 25}
	personPtr := &person
	
	fmt.Printf("Struct: %+v\n", person)
	fmt.Printf("Struct pointer: %p\n", personPtr)
	fmt.Printf("Access through pointer: %+v\n", *personPtr)
	
	// Field access through pointer (automatic dereferencing)
	personPtr.Name = "Bob"
	personPtr.Age = 30
	fmt.Printf("Modified through pointer: %+v\n", person)

	// 7. Pointers and arrays/slices
	fmt.Println("\n7. Pointers and arrays/slices:")
	arr := [5]int{10, 20, 30, 40, 50}
	arrPtr := &arr
	
	fmt.Printf("Array: %v\n", arr)
	fmt.Printf("Array pointer: %p\n", arrPtr)
	fmt.Printf("First element through pointer: %d\n", (*arrPtr)[0])
	
	// Pointer to first element
	firstPtr := &arr[0]
	fmt.Printf("Pointer to first element: %p, value: %d\n", firstPtr, *firstPtr)

	// 8. Pointers and maps
	fmt.Println("\n8. Pointers and maps:")
	m := map[string]*int{
		"a": new(int),
		"b": new(int),
	}
	
	*m["a"] = 100
	*m["b"] = 200
	
	fmt.Printf("Map values: ")
	for key, ptr := range m {
		fmt.Printf("%s: %d ", key, *ptr)
	}
	fmt.Println()

	// 9. Pointer arithmetic (not allowed in Go)
	fmt.Println("\n9. Pointer arithmetic:")
	// Go does not allow pointer arithmetic like in C/C++
	// The following would be illegal:
	// ptr++
	// ptr = ptr + 1
	
	// But you can use slice operations for similar functionality
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Slice: %v\n", slice)
	fmt.Printf("Slice[1:3]: %v\n", slice[1:3])

	// 10. Comparing pointers
	fmt.Println("\n10. Comparing pointers:")
	x, y := 10, 10
	ptrX, ptrY := &x, &y
	ptrZ := &x
	
	fmt.Printf("ptrX == ptrY: %t (different variables)\n", ptrX == ptrY)
	fmt.Printf("ptrX == ptrZ: %t (same variable)\n", ptrX == ptrZ)
	fmt.Printf("*ptrX == *ptrY: %t (same values)\n", *ptrX == *ptrY)

	// 11. Pointers in practice - modifying large data
	fmt.Println("\n11. Pointers for efficiency:")
	type LargeStruct struct {
		Data [1000]int
	}
	
	// Without pointer (copies entire struct)
	ls := LargeStruct{Data: [1000]int{1, 2, 3}}
	fmt.Printf("Without pointer - first elements: %d, %d, %d\n", ls.Data[0], ls.Data[1], ls.Data[2])
	
	// With pointer (only copies pointer)
	lsPtr := &LargeStruct{Data: [1000]int{4, 5, 6}}
	fmt.Printf("With pointer - first elements: %d, %d, %d\n", lsPtr.Data[0], lsPtr.Data[1], lsPtr.Data[2])

	// 12. Pointer receiver methods
	fmt.Println("\n12. Pointer receiver methods:")
	type Counter struct {
		count int
	}
	
	// Value receiver
	func (c Counter) getValue() int {
		return c.count
	}
	
	// Pointer receiver
	func (c *Counter) increment() {
		c.count++
	}
	
	counter := Counter{count: 0}
	fmt.Printf("Initial count: %d\n", counter.getValue())
	
	counter.increment()
	fmt.Printf("After increment: %d\n", counter.getValue())
	
	// Can call pointer receiver on value (Go automatically takes address)
	counter.increment()
	fmt.Printf("After another increment: %d\n", counter.getValue())

	// 13. Common pointer patterns
	fmt.Println("\n13. Common pointer patterns:")
	
	// Factory function returning pointer
	newPerson := func(name string, age int) *Person {
		return &Person{Name: name, Age: age}
	}
	
	p1 := newPerson("Charlie", 35)
	fmt.Printf("Created person: %+v\n", *p1)
	
	// Optional return value pattern
	findUser := func(id int) *Person {
		users := map[int]Person{
			1: {Name: "User1", Age: 25},
			2: {Name: "User2", Age: 30},
		}
		
		if user, exists := users[id]; exists {
			return &user
		}
		return nil // User not found
	}
	
	user := findUser(1)
	if user != nil {
		fmt.Printf("Found user: %+v\n", *user)
	} else {
		fmt.Println("User not found")
	}
	
	user = findUser(99)
	if user != nil {
		fmt.Printf("Found user: %+v\n", *user)
	} else {
		fmt.Println("User not found")
	}
}

// Helper functions
func swap(a, b *int) {
	*a, *b = *b, *a
}

func createPointer(value int) *int {
	ptr := new(int)
	*ptr = value
	return ptr
}
