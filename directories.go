package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Directories ===")

	// Create directory
	err := os.Mkdir("testdir", 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
	} else {
		fmt.Printf("Created directory: testdir\n")
	}

	// Create nested directories
	err = os.MkdirAll("nested/deep/path", 0755)
	if err != nil {
		fmt.Printf("Error creating nested directories: %v\n", err)
	} else {
		fmt.Printf("Created nested directories\n")
	}

	// Read directory contents
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	fmt.Println("\nCurrent directory contents:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("  DIR:  %s\n", entry.Name())
		} else {
			fmt.Printf("  FILE: %s\n", entry.Name())
		}
	}

	// Check if path exists and is directory
	if info, err := os.Stat("testdir"); err == nil {
		fmt.Printf("testdir exists: %t\n", info.IsDir())
	}

	// Create directory with files
	os.Mkdir("withfiles", 0755)
	os.WriteFile("withfiles/file1.txt", []byte("content1"), 0644)
	os.WriteFile("withfiles/file2.txt", []byte("content2"), 0644)

	// List files in directory
	if entries, err := os.ReadDir("withfiles"); err == nil {
		fmt.Println("\nFiles in 'withfiles' directory:")
		for _, entry := range entries {
			if !entry.IsDir() {
				fmt.Printf("  %s\n", entry.Name())
			}
		}
	}

	// Get current working directory
	wd, err := os.Getwd()
	if err == nil {
		fmt.Printf("\nCurrent working directory: %s\n", wd)
	}

	// Change directory
	err = os.Chdir("testdir")
	if err == nil {
		fmt.Printf("Changed to testdir\n")
		os.Chdir("..") // Go back
	}

	// Remove directory (must be empty)
	err = os.Remove("testdir")
	if err == nil {
		fmt.Printf("Removed empty directory: testdir\n")
	}

	// Remove directory and all contents
	err = os.RemoveAll("nested")
	if err == nil {
		fmt.Printf("Removed nested directory tree\n")
	}

	err = os.RemoveAll("withfiles")
	if err == nil {
		fmt.Printf("Removed withfiles directory\n")
	}

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "example")
	if err == nil {
		fmt.Printf("Created temp directory: %s\n", tempDir)
		os.RemoveAll(tempDir) // Clean up
	}
}
