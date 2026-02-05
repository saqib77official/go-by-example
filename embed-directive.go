package main

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed hello.txt
var helloFile string

//go:embed files/*.txt
var textFiles embed.FS

func main() {
	fmt.Println("=== Embed Directive ===")

	// Embed single file as string
	fmt.Printf("Embedded file content: %s\n", helloFile)

	// Read embedded file system
	fmt.Println("\nReading from embedded FS:")
	
	// List embedded files
	entries, err := textFiles.ReadDir("files")
	if err != nil {
		panic(err)
	}

	fmt.Println("Embedded files:")
	for _, entry := range entries {
		fmt.Printf("  %s\n", entry.Name())
	}

	// Read specific embedded file
	content, err := textFiles.ReadFile("files/test.txt")
	if err == nil {
		fmt.Printf("Content of test.txt: %s\n", string(content))
	}

	// Read all embedded files
	fs.WalkDir(textFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			content, err := textFiles.ReadFile(path)
			if err == nil {
				fmt.Printf("\nFile: %s\n", path)
				fmt.Printf("Content: %s\n", string(content))
			}
		}
		return nil
	})

	// Check if file exists in embedded FS
	if _, err := textFiles.Stat("files/test.txt"); err == nil {
		fmt.Println("files/test.txt exists in embedded FS")
	}

	// Open embedded file as file
	file, err := textFiles.Open("files/test.txt")
	if err == nil {
		defer file.Close()
		
		buffer := make([]byte, 100)
		n, err := file.Read(buffer)
		if err == nil {
			fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
		}
	}

	// Embed multiple file types
	fmt.Println("\nEmbed demonstration complete")
	fmt.Println("Note: This example requires actual embedded files to work properly")
	fmt.Println("Create hello.txt and files/ directory with test.txt to see full functionality")
}
