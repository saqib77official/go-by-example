package main

import "fmt"

func main() {
	fmt.Println("=== Range Over Built-in Types Examples ===")

	// 1. Range over slice
	fmt.Println("\n1. Range over slice:")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("Numbers: %v\n", numbers)
	
	fmt.Println("With index and value:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	
	fmt.Println("Only values:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
	
	fmt.Println("Only indices:")
	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
	}

	// 2. Range over array
	fmt.Println("\n2. Range over array:")
	colors := [3]string{"red", "green", "blue"}
	fmt.Printf("Colors: %v\n", colors)
	
	for i, color := range colors {
		fmt.Printf("Index %d: %s\n", i, color)
	}

	// 3. Range over string
	fmt.Println("\n3. Range over string:")
	message := "Hello, 世界"
	fmt.Printf("Message: %s\n", message)
	
	fmt.Println("With index and rune:")
	for index, runeValue := range message {
		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", index, runeValue, runeValue)
	}
	
	fmt.Println("Only runes:")
	for _, runeValue := range message {
		fmt.Printf("Rune: %c\n", runeValue)
	}

	// 4. Range over map
	fmt.Println("\n4. Range over map:")
	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
		"Diana":   28,
	}
	fmt.Printf("Ages map: %v\n", ages)
	
	fmt.Println("Key-value pairs:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
	
	fmt.Println("Only keys:")
	for name := range ages {
		fmt.Printf("Name: %s\n", name)
	}
	
	fmt.Println("Only values:")
	for _, age := range ages {
		fmt.Printf("Age: %d\n", age)
	}

	// 5. Range over map with different value types
	fmt.Println("\n5. Range over map with slice values:")
	studentsByGrade := map[string][]string{
		"A": {"Alice", "Adam"},
		"B": {"Bob", "Bella"},
		"C": {"Charlie", "Carol"},
	}
	
	for grade, students := range studentsByGrade {
		fmt.Printf("Grade %s: %v\n", grade, students)
	}

	// 6. Range over channels
	fmt.Println("\n6. Range over channel:")
	ch := make(chan int)
	
	// Start goroutine to send values
	go func() {
		defer close(ch)
		for i := 1; i <= 5; i++ {
			ch <- i * 10
		}
	}()
	
	fmt.Println("Receiving from channel:")
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}

	// 7. Range over byte slice
	fmt.Println("\n7. Range over byte slice:")
	data := []byte{72, 101, 108, 108, 111} // "Hello" in ASCII
	fmt.Printf("Byte slice: %v\n", data)
	
	for i, b := range data {
		fmt.Printf("Index %d: Byte %d, Char: %c\n", i, b, b)
	}

	// 8. Range over rune slice
	fmt.Println("\n8. Range over rune slice:")
	runes := []rune{'H', 'e', 'l', 'l', 'o', '世', '界'}
	fmt.Printf("Rune slice: %v\n", runes)
	
	for i, r := range runes {
		fmt.Printf("Index %d: Rune %c, Unicode: %U\n", i, r, r)
	}

	// 9. Range over empty collections
	fmt.Println("\n9. Range over empty collections:")
	var emptySlice []int
	var emptyMap map[string]int
	var emptyString string
	
	fmt.Println("Empty slice:")
	for i, v := range emptySlice {
		fmt.Printf("This won't print: %d, %d\n", i, v)
	}
	
	fmt.Println("Empty map:")
	for k, v := range emptyMap {
		fmt.Printf("This won't print: %s, %d\n", k, v)
	}
	
	fmt.Println("Empty string:")
	for i, r := range emptyString {
		fmt.Printf("This won't print: %d, %c\n", i, r)
	}

	// 10. Range with modification
	fmt.Println("\n10. Range with modification:")
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n", slice)
	
	// Modify values using range
	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Printf("After modification: %v\n", slice)

	// 11. Range over multidimensional structures
	fmt.Println("\n11. Range over multidimensional slice:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	for rowIdx, row := range matrix {
		fmt.Printf("Row %d: ", rowIdx)
		for colIdx, value := range row {
			fmt.Printf("[%d]=%d ", colIdx, value)
		}
		fmt.Println()
	}

	// 12. Range over interface{} slice
	fmt.Println("\n12. Range over interface{} slice:")
	mixed := []interface{}{1, "hello", 3.14, true, []int{1, 2, 3}}
	
	for i, value := range mixed {
		fmt.Printf("Index %d: Value %v (Type: %T)\n", i, value, value)
	}

	// 13. Range with early termination
	fmt.Println("\n13. Range with early termination:")
	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	fmt.Println("Find first even number:")
	for i, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("Found even number %d at index %d\n", num, i)
			break
		}
	}
	
	fmt.Println("Skip first 3 elements:")
	for i, num := range numbers {
		if i < 3 {
			continue
		}
		fmt.Printf("Index %d: %d\n", i, num)
		if i >= 6 {
			break
		}
	}
}
