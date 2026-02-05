package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("=== File Paths ===")

	// Current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Working directory: %s\n", wd)

	// Path components
	path := "/home/user/documents/file.txt"
	dir := filepath.Dir(path)
	file := filepath.Base(path)
	ext := filepath.Ext(path)
	
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Directory: %s\n", dir)
	fmt.Printf("File name: %s\n", file)
	fmt.Printf("Extension: %s\n", ext)

	// Join paths
	joined := filepath.Join("home", "user", "documents", "file.txt")
	fmt.Printf("Joined path: %s\n", joined)

	// Clean path
	messy := "home//user/../user/./documents/file.txt"
	clean := filepath.Clean(messy)
	fmt.Printf("Messy: %s\n", messy)
	fmt.Printf("Clean: %s\n", clean)

	// Absolute path
	abs, err := filepath.Abs("relative/path.txt")
	if err == nil {
		fmt.Printf("Absolute path: %s\n", abs)
	}

	// Split path
	dir2, file2 := filepath.Split("/path/to/file.txt")
	fmt.Printf("Split - Dir: %s, File: %s\n", dir2, file2)

	// Path separator
	fmt.Printf("Path separator: %s\n", string(filepath.Separator))

	// Match patterns
	files, err := filepath.Glob("*.go")
	if err == nil {
		fmt.Printf("Go files: %v\n", files)
	}

	// Walk directory
	fmt.Println("\nWalking current directory:")
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Printf("  %s\n", path)
		}
		return nil
	})

	// File extension operations
	filename := "archive.tar.gz"
	ext = filepath.Ext(filename)
	base := filename[:len(filename)-len(ext)]
	fmt.Printf("File: %s, Ext: %s, Base: %s\n", filename, ext, base)

	// OS-specific paths
	if runtime.GOOS == "windows" {
		fmt.Printf("Windows path: %s\n", filepath.Join("C:", "Program Files", "App"))
	} else {
		fmt.Printf("Unix path: %s\n", filepath.Join("/", "usr", "local", "bin"))
	}

	// Relative path
	rel, err := filepath.Rel(wd, filepath.Join(wd, "subdir", "file.txt"))
	if err == nil {
		fmt.Printf("Relative path: %s\n", rel)
	}

	// Check if path is absolute
	fmt.Printf("Is absolute '/home/user': %t\n", filepath.IsAbs("/home/user"))
	fmt.Printf("Is absolute 'relative/path': %t\n", filepath.IsAbs("relative/path"))
}
