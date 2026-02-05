package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Writing Files ===")

	// Write entire file
	content := "Hello, World!\nThis is a test file."
	err := os.WriteFile("output.txt", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote to output.txt\n")

	// Append to file
	appendContent := "\nThis line was appended."
	err = os.WriteFile("output.txt", []byte(appendContent), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Appended to output.txt\n")

	// Write with file handle
	file, err := os.Create("handle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString("Line 1\n")
	file.WriteString("Line 2\n")
	file.WriteString("Line 3\n")
	fmt.Printf("Wrote to handle.txt\n")

	// Write with buffered writer
	bufFile, err := os.Create("buffered.txt")
	if err != nil {
		panic(err)
	}
	defer bufFile.Close()

	writer := bufio.NewWriter(bufFile)
	writer.WriteString("Buffered line 1\n")
	writer.WriteString("Buffered line 2\n")
	writer.Flush() // Important: flush to write to file
	fmt.Printf("Wrote to buffered.txt\n")

	// Write bytes
	binFile, err := os.Create("binary.dat")
	if err != nil {
		panic(err)
	}
	defer binFile.Close()

	bytes := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello"
	binFile.Write(bytes)
	fmt.Printf("Wrote binary data to binary.dat\n")

	// Write multiple lines with newline
	linesFile, err := os.Create("lines.txt")
	if err != nil {
		panic(err)
	}
	defer linesFile.Close()

	lines := []string{"First line", "Second line", "Third line"}
	for _, line := range lines {
		linesFile.WriteString(line + "\n")
	}
	fmt.Printf("Wrote lines to lines.txt\n")

	// Write formatted content
	fmtFile, err := os.Create("formatted.txt")
	if err != nil {
		panic(err)
	}
	defer fmtFile.Close()

	fmt.Fprintf(fmtFile, "Name: %s\n", "Alice")
	fmt.Fprintf(fmtFile, "Age: %d\n", 25)
	fmt.Fprintf(fmtFile, "Score: %.2f\n", 95.5)
	fmt.Printf("Wrote formatted content to formatted.txt\n")

	// Clean up
	files := []string{"output.txt", "handle.txt", "buffered.txt", "binary.dat", "lines.txt", "formatted.txt"}
	for _, file := range files {
		os.Remove(file)
		fmt.Printf("Removed %s\n", file)
	}
}
