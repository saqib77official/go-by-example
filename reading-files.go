package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("=== Reading Files ===")

	// Create a test file
	content := "Line 1\nLine 2\nLine 3\nLine 4\nLine 5"
	err := os.WriteFile("test.txt", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	defer os.Remove("test.txt")

	// Read entire file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read all: %s\n", string(data))

	// Read file line by line
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Lines:")
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	// Read with buffer
	file.Seek(0, 0) // Reset to beginning
	reader := bufio.NewReader(file)
	buffer := make([]byte, 10)
	
	fmt.Println("Reading with buffer:")
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
	}

	// Read specific number of bytes
	file.Seek(0, 0)
	specific := make([]byte, 5)
	n, err := file.Read(specific)
	if err == nil {
		fmt.Printf("First 5 bytes: %s\n", string(specific[:n]))
	}

	// Check file info
	info, err := os.Stat("test.txt")
	if err == nil {
		fmt.Printf("File size: %d bytes\n", info.Size())
		fmt.Printf("File mode: %v\n", info.Mode())
		fmt.Printf("Is directory: %t\n", info.IsDir())
	}

	// Read from stdin simulation
	fmt.Println("Reading from stdin (simulated):")
	stdin := os.Stdin
	fmt.Print("Enter something: ")
	input := make([]byte, 100)
	n, err = stdin.Read(input)
	if err == nil {
		fmt.Printf("You entered: %s\n", string(input[:n]))
	}
}
