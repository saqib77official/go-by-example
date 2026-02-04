package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("=== Strings and Runes Examples ===")

	// 1. Basic string operations
	fmt.Println("\n1. Basic string operations:")
	str := "Hello, World!"
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Length: %d\n", len(str))
	fmt.Printf("First character: %c\n", str[0])
	fmt.Printf("Last character: %c\n", str[len(str)-1])

	// 2. String literals
	fmt.Println("\n2. String literals:")
	raw := `This is a raw string\nwith newlines and\ttabs`
	interpreted := "This is an interpreted string\nwith newlines and\ttabs"
	
	fmt.Printf("Raw string: %s\n", raw)
	fmt.Printf("Interpreted string: %s\n", interpreted)

	// 3. String concatenation
	fmt.Println("\n3. String concatenation:")
	first := "Hello"
	second := "World"
	combined := first + ", " + second + "!"
	fmt.Printf("Combined: %s\n", combined)
	
	// Using strings.Builder for efficiency
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World")
	builder.WriteString("!")
	fmt.Printf("Builder result: %s\n", builder.String())

	// 4. Runes and Unicode
	fmt.Println("\n4. Runes and Unicode:")
	unicodeStr := "Hello, 世界! 🌍"
	fmt.Printf("Unicode string: %s\n", unicodeStr)
	fmt.Printf("Byte length: %d\n", len(unicodeStr))
	fmt.Printf("Rune count: %d\n", len([]rune(unicodeStr)))
	
	// Accessing runes
	runes := []rune(unicodeStr)
	for i, r := range runes {
		fmt.Printf("Index %d: Rune '%c' (Unicode: %U)\n", i, r, r)
	}

	// 5. String iteration
	fmt.Println("\n5. String iteration:")
	text := "Go Programming"
	
	fmt.Println("Iterating by bytes:")
	for i := 0; i < len(text); i++ {
		fmt.Printf("Byte %d: %c\n", i, text[i])
	}
	
	fmt.Println("Iterating by runes:")
	for i, r := range text {
		fmt.Printf("Rune %d: %c\n", i, r)
	}

	// 6. String manipulation functions
	fmt.Println("\n6. String manipulation functions:")
	sample := "  Go Programming Language  "
	
	fmt.Printf("Original: '%s'\n", sample)
	fmt.Printf("Trim: '%s'\n", strings.TrimSpace(sample))
	fmt.Printf("Upper: '%s'\n", strings.ToUpper(sample))
	fmt.Printf("Lower: '%s'\n", strings.ToLower(sample))
	fmt.Printf("Title: '%s'\n", strings.Title(strings.TrimSpace(sample)))

	// 7. String searching and splitting
	fmt.Println("\n7. String searching and splitting:")
	sentence := "The quick brown fox jumps over the lazy dog"
	
	fmt.Printf("Original: %s\n", sentence)
	fmt.Printf("Contains 'fox': %t\n", strings.Contains(sentence, "fox"))
	fmt.Printf("Starts with 'The': %t\n", strings.HasPrefix(sentence, "The"))
	fmt.Printf("Ends with 'dog': %t\n", strings.HasSuffix(sentence, "dog"))
	fmt.Printf("Index of 'fox': %d\n", strings.Index(sentence, "fox"))
	
	words := strings.Split(sentence, " ")
	fmt.Printf("Split by space: %v\n", words)
	fmt.Printf("Join with '-': %s\n", strings.Join(words, "-"))

	// 8. String replacement
	fmt.Println("\n8. String replacement:")
	oldText := "Hello World, Hello Universe"
	
	fmt.Printf("Original: %s\n", oldText)
	fmt.Printf("Replace 'Hello' with 'Hi': %s\n", strings.Replace(oldText, "Hello", "Hi", -1))
	fmt.Printf("Replace first 'Hello': %s\n", strings.Replace(oldText, "Hello", "Hi", 1))
	fmt.Printf("ReplaceAll: %s\n", strings.ReplaceAll(oldText, "Hello", "Hi"))

	// 9. Working with substrings
	fmt.Println("\n9. Working with substrings:")
	text = "Hello, Go Programming!"
	
	fmt.Printf("Original: %s\n", text)
	fmt.Printf("First 5 chars: %s\n", text[:5])
	fmt.Printf("Last 5 chars: %s\n", text[len(text)-5:])
	fmt.Printf("Middle substring: %s\n", text[7:9])

	// 10. Rune operations
	fmt.Println("\n10. Rune operations:")
	runeStr := "Hello123!@#"
	
	var letters, digits, symbols int
	for _, r := range runeStr {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsSymbol(r) || unicode.IsPunct(r):
			symbols++
		}
	}
	
	fmt.Printf("String: %s\n", runeStr)
	fmt.Printf("Letters: %d, Digits: %d, Symbols: %d\n", letters, digits, symbols)

	// 11. String formatting
	fmt.Println("\n11. String formatting:")
	name := "Alice"
	age := 25
	score := 95.5
	
	formatted := fmt.Sprintf("Name: %s, Age: %d, Score: %.1f", name, age, score)
	fmt.Printf("Formatted: %s\n", formatted)
	
	// Using strings.Builder with formatting
	var fmtBuilder strings.Builder
	fmtBuilder.WriteString("User Info:\n")
	fmtBuilder.WriteString(fmt.Sprintf("  Name: %s\n", name))
	fmtBuilder.WriteString(fmt.Sprintf("  Age: %d\n", age))
	fmtBuilder.WriteString(fmt.Sprintf("  Score: %.1f\n", score))
	fmt.Printf("Builder with formatting:\n%s", fmtBuilder.String())

	// 12. String comparison
	fmt.Println("\n12. String comparison:")
	str1 := "Hello"
	str2 := "hello"
	str3 := "Hello"
	
	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' == '%s': %t\n", str1, str3, str1 == str3)
	fmt.Printf("'%s' < '%s': %t\n", str1, str2, str1 < str2)
	fmt.Printf("EqualFold('%s', '%s'): %t\n", str1, str2, strings.EqualFold(str1, str2))

	// 13. String to number conversion
	fmt.Println("\n13. String to number conversion:")
	numStr := "123"
	floatStr := "45.67"
	
	// Using fmt.Sscanf
	var num int
	var f float64
	fmt.Sscanf(numStr, "%d", &num)
	fmt.Sscanf(floatStr, "%f", &f)
	
	fmt.Printf("String '%s' to int: %d\n", numStr, num)
	fmt.Printf("String '%s' to float: %.2f\n", floatStr, f)

	// 14. Multiline strings
	fmt.Println("\n14. Multiline strings:")
	multiline := `This is a multiline string.
It can span multiple lines.
And preserve formatting.
    With indentation too!`
	
	fmt.Printf("Multiline string:\n%s\n", multiline)

	// 15. String and byte slice conversion
	fmt.Println("\n15. String and byte slice conversion:")
	original := "Hello, Go!"
	
	// String to byte slice
	bytes := []byte(original)
	fmt.Printf("String to bytes: %v\n", bytes)
	
	// Byte slice to string
	backToString := string(bytes)
	fmt.Printf("Bytes to string: %s\n", backToString)
	
	// Rune slice to string
	runes = []rune{'H', 'e', 'l', 'l', 'o'}
	runeString := string(runes)
	fmt.Printf("Runes to string: %s\n", runeString)

	// 16. String patterns and validation
	fmt.Println("\n16. String patterns and validation:")
	email := "user@example.com"
	password := "Secure123!"
	
	// Simple email validation
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	fmt.Printf("Email '%s' is valid: %t\n", email, isValidEmail)
	
	// Password validation
	var hasUpper, hasLower, hasDigit bool
	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		}
	}
	isStrongPassword := hasUpper && hasLower && hasDigit && len(password) >= 8
	fmt.Printf("Password is strong: %t\n", isStrongPassword)
}
