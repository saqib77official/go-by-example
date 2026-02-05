package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Random Numbers ===")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Random integer
	fmt.Printf("Random int (0-99): %d\n", rand.Intn(100))

	// Random int in range
	fmt.Printf("Random int (10-20): %d\n", rand.Intn(11)+10)

	// Random float64
	fmt.Printf("Random float64: %f\n", rand.Float64())

	// Random float64 in range
	fmt.Printf("Random float64 (0-10): %f\n", rand.Float64()*10)

	// Random boolean
	fmt.Printf("Random boolean: %t\n", rand.Intn(2) == 1)

	// Random choice from slice
	fruits := []string{"Apple", "Banana", "Cherry", "Date"}
	choice := fruits[rand.Intn(len(fruits))]
	fmt.Printf("Random fruit: %s\n", choice)

	// Random string
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Printf("Random string: %s\n", string(b))

	// Random UUID-like
	uuid := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		rand.Uint32(), rand.Uint16(), rand.Uint16(),
		rand.Uint16(), rand.Uint64())
	fmt.Printf("Random UUID: %s\n", uuid)

	// Shuffle slice
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Printf("Shuffled: %v\n", numbers)

	// Random password
	password := make([]byte, 12)
	for i := range password {
		switch rand.Intn(3) {
		case 0:
			password[i] = byte(rand.Intn(26) + 'a')
		case 1:
			password[i] = byte(rand.Intn(26) + 'A')
		case 2:
			password[i] = byte(rand.Intn(10) + '0')
		}
	}
	fmt.Printf("Random password: %s\n", string(password))
}
