package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("=== Base64 Encoding ===")

	// Encode string
	text := "Hello, World!"
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Printf("Encoded: %s\n", encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err == nil {
		fmt.Printf("Decoded: %s\n", decoded)
	}

	// URL-safe encoding
	urlSafe := base64.URLEncoding.EncodeToString([]byte("Hello+World/"))
	fmt.Printf("URL-safe encoded: %s\n", urlSafe)

	// Raw encoding (no padding)
	rawEncoded := base64.RawStdEncoding.EncodeToString([]byte(text))
	fmt.Printf("Raw encoded: %s\n", rawEncoded)

	// Encode binary data
	data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f} // "Hello" in bytes
	binaryEncoded := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("Binary encoded: %s\n", binaryEncoded)

	// Decode with validation
	validEncoded := "SGVsbG8sIFdvcmxkIQ=="
	invalidEncoded := "Invalid!@#$"

	if decoded, err := base64.StdEncoding.DecodeString(validEncoded); err == nil {
		fmt.Printf("Valid decoded: %s\n", decoded)
	}

	if _, err := base64.StdEncoding.DecodeString(invalidEncoded); err != nil {
		fmt.Printf("Invalid decode error: %v\n", err)
	}

	// Encode large data
	largeText := "This is a longer string that will demonstrate base64 encoding with more content."
	largeEncoded := base64.StdEncoding.EncodeToString([]byte(largeText))
	fmt.Printf("Large encoded: %s\n", largeEncoded)

	// Different encodings
	encodings := []struct {
		name string
		enc  *base64.Encoding
	}{
		{"Standard", base64.StdEncoding},
		{"URL Safe", base64.URLEncoding},
		{"Raw Std", base64.RawStdEncoding},
		{"Raw URL", base64.RawURLEncoding},
	}

	for _, e := range encodings {
		encoded := e.enc.EncodeToString([]byte(text))
		fmt.Printf("%s: %s\n", e.name, encoded)
	}

	// Check if string is valid base64
	isValidBase64 := func(s string) bool {
		_, err := base64.StdEncoding.DecodeString(s)
		return err == nil
	}

	fmt.Printf("'SGVsbG8=' is valid base64: %t\n", isValidBase64("SGVsbG8="))
	fmt.Printf("'Invalid!' is valid base64: %t\n", isValidBase64("Invalid!"))
}
