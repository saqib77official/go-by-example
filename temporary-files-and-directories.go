package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("=== Temporary Files and Directories ===")

	// Create temporary file
	tempFile, err := os.CreateTemp("", "example-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tempFile.Name()) // Clean up
	defer tempFile.Close()

	fmt.Printf("Created temp file: %s\n", tempFile.Name())

	// Write to temp file
	tempFile.WriteString("This is temporary content")
	tempFile.WriteString("\nLine 2")
	fmt.Printf("Wrote content to temp file\n")

	// Read from temp file
	tempFile.Seek(0, 0)
	content, err := io.ReadAll(tempFile)
	if err == nil {
		fmt.Printf("Read from temp file: %s\n", string(content))
	}

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "example-dir-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	fmt.Printf("Created temp directory: %s\n", tempDir)

	// Create file in temp directory
	tempFileInDir, err := os.CreateTemp(tempDir, "data-*.txt")
	if err != nil {
		panic(err)
	}
	defer tempFileInDir.Close()

	fmt.Printf("Created file in temp dir: %s\n", tempFileInDir.Name())

	// List contents of temp directory
	entries, err := os.ReadDir(tempDir)
	if err == nil {
		fmt.Printf("Contents of temp directory:\n")
		for _, entry := range entries {
			fmt.Printf("  %s\n", entry.Name())
		}
	}

	// Multiple temp files
	fmt.Println("\nCreating multiple temp files:")
	for i := 0; i < 3; i++ {
		file, err := os.CreateTemp("", "multi-*.tmp")
		if err != nil {
			continue
		}
		defer os.Remove(file.Name())
		defer file.Close()

		file.WriteString(fmt.Sprintf("File %d content", i+1))
		fmt.Printf("  %s\n", file.Name())
	}

	// Custom temp file name pattern
	customFile, err := os.CreateTemp("", "custom-*.log")
	if err != nil {
		panic(err)
	}
	defer os.Remove(customFile.Name())
	defer customFile.Close()

	fmt.Printf("\nCustom pattern temp file: %s\n", customFile.Name())

	// Temp file with specific permissions
	permFile, err := os.CreateTemp("", "perm-*.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(permFile.Name())
	defer permFile.Close()

	// Set specific permissions
	permFile.Chmod(0600) // Read/write for owner only
	fmt.Printf("Created temp file with restricted permissions: %s\n", permFile.Name())

	// Working with temp file paths
	tempPath := tempFile.Name()
	fmt.Printf("\nTemp file path operations:\n")
	fmt.Printf("  Full path: %s\n", tempPath)
	fmt.Printf("  Directory: %s\n", tempPath[:len(tempPath)-len(tempFile.Name())])
	fmt.Printf("  File name: %s\n", tempFile.Name())

	// Copy temp file to new location
	copyFile := func(src, dst string) error {
		source, err := os.Open(src)
		if err != nil {
			return err
		}
		defer source.Close()

		destination, err := os.Create(dst)
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		return err
	}

	finalFile := "final-output.txt"
	err = copyFile(tempFile.Name(), finalFile)
	if err == nil {
		fmt.Printf("Copied temp file to: %s\n", finalFile)
		defer os.Remove(finalFile)
	}

	fmt.Println("Temporary files and directories will be cleaned up automatically")
}
