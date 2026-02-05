package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("=== SHA256 Hashes ===")

	// Hash string
	text := "Hello, World!"
	hash := sha256.Sum256([]byte(text))
	fmt.Printf("SHA256 of '%s': %x\n", text, hash)

	// Hash with hex encoding
	hashHex := hex.EncodeToString(hash[:])
	fmt.Printf("SHA256 hex: %s\n", hashHex)

	// Hash file content
	data := []byte("This is some data to hash")
	fileHash := sha256.Sum256(data)
	fmt.Printf("Data hash: %x\n", fileHash)

	// Compare hashes
	text2 := "Hello, World!"
	hash2 := sha256.Sum256([]byte(text2))
	
	fmt.Printf("Hashes equal: %t\n", hash == hash2)

	// Hash different text
	text3 := "Hello, world!"
	hash3 := sha256.Sum256([]byte(text3))
	fmt.Printf("Hashes equal (case sensitive): %t\n", hash == hash3)

	// Hash empty string
	emptyHash := sha256.Sum256([]byte(""))
	fmt.Printf("Empty string hash: %x\n", emptyHash)

	// Hash multiple times (same result)
	hash4 := sha256.Sum256([]byte("test"))
	hash5 := sha256.Sum256([]byte("test"))
	fmt.Printf("Same input, same hash: %t\n", hash4 == hash5)

	// Create hash incrementally
	h := sha256.New()
	h.Write([]byte("Hello, "))
	h.Write([]byte("World!"))
	incrementalHash := h.Sum(nil)
	fmt.Printf("Incremental hash: %x\n", incrementalHash)

	// Verify incremental hash equals direct hash
	fmt.Printf("Incremental equals direct: %t\n", incrementalHash == hash[:])

	// Hash large data
	largeData := make([]byte, 1000)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}
	largeHash := sha256.Sum256(largeData)
	fmt.Printf("Large data hash: %x\n", largeHash)

	// Hash length (always 32 bytes for SHA256)
	fmt.Printf("Hash length: %d bytes\n", len(hash))
}
