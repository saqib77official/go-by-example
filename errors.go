package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("=== Errors Examples ===")

	// 1. Basic error creation and handling
	fmt.Println("\n1. Basic error handling:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
	
	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}

	// 2. Using errors.New
	fmt.Println("\n2. Using errors.New:")
	err = errors.New("something went wrong")
	fmt.Printf("Custom error: %v\n", err)
	fmt.Printf("Error type: %T\n", err)

	// 3. Using fmt.Errorf
	fmt.Println("\n3. Using fmt.Errorf:")
	name := "Alice"
	age := -5
	if age < 0 {
		err = fmt.Errorf("invalid age %d for user %s", age, name)
		fmt.Printf("Formatted error: %v\n", err)
	}

	// 4. Error wrapping (Go 1.13+)
	fmt.Println("\n4. Error wrapping:")
	baseErr := errors.New("database connection failed")
	wrappedErr := fmt.Errorf("failed to save user: %w", baseErr)
	
	fmt.Printf("Wrapped error: %v\n", wrappedErr)
	fmt.Printf("Unwrapped error: %v\n", errors.Unwrap(wrappedErr))
	fmt.Printf("Is database error: %t\n", errors.Is(wrappedErr, baseErr))

	// 5. Error checking with errors.Is
	fmt.Println("\n5. Error checking with errors.Is:")
	err = processUser("admin")
	if errors.Is(err, ErrUserNotFound) {
		fmt.Printf("User not found: %v\n", err)
	} else if errors.Is(err, ErrPermissionDenied) {
		fmt.Printf("Permission denied: %v\n", err)
	}

	// 6. Error type checking with errors.As
	fmt.Println("\n6. Error type checking with errors.As:")
	err = readFile("nonexistent.txt")
	
	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Printf("Path error: Op=%s, Path=%s, Err=%v\n", 
			pathError.Op, pathError.Path, pathError.Err)
	}

	// 7. Multiple error handling
	fmt.Println("\n7. Multiple error handling:")
	errs := validateUser("bob@example.com", 15)
	if len(errs) > 0 {
		fmt.Printf("Validation errors:\n")
		for i, validationErr := range errs {
			fmt.Printf("  %d: %v\n", i+1, validationErr)
		}
	}

	// 8. Error with additional context
	fmt.Println("\n8. Error with context:")
	err = processOrder("order123")
	if err != nil {
		fmt.Printf("Order processing failed: %v\n", err)
		
		// Get error chain
		fmt.Println("Error chain:")
		for err := err; err != nil; err = errors.Unwrap(err) {
			fmt.Printf("  - %v\n", err)
		}
	}

	// 9. Sentinel errors
	fmt.Println("\n9. Sentinel errors:")
	err = checkPermission("guest", "admin")
	if errors.Is(err, ErrPermissionDenied) {
		fmt.Printf("Access denied: %v\n", err)
	}

	// 10. Custom error types
	fmt.Println("\n10. Custom error types:")
	err = &ValidationError{
		Field:   "email",
		Value:   "invalid-email",
		Message: "invalid email format",
	}
	
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Printf("Validation error: %s\n", validationErr.Error())
		fmt.Printf("Field: %s, Value: %s\n", validationErr.Field, validationErr.Value)
	}

	// 11. Error handling in functions
	fmt.Println("\n11. Error handling patterns:")
	
	// Pattern 1: Return error immediately
	err = saveToDatabase("user123")
	if err != nil {
		fmt.Printf("Database save failed: %v\n", err)
		return
	}
	
	// Pattern 2: Collect multiple errors
	errs = []error{}
	err1 := validateEmail("invalid")
	err2 := validateAge(-5)
	
	if err1 != nil {
		errs = append(errs, err1)
	}
	if err2 != nil {
		errs = append(errs, err2)
	}
	
	if len(errs) > 0 {
		fmt.Printf("Multiple validation errors: %v\n", errs)
	}

	// 12. Panic and recover
	fmt.Println("\n12. Panic and recover:")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		
		// This will panic
		causePanic()
	}()

	// 13. Error handling with defer
	fmt.Println("\n13. Error handling with defer:")
	err = processFile("test.txt")
	if err != nil {
		fmt.Printf("File processing error: %v\n", err)
	}

	// 14. Error aggregation
	fmt.Println("\n14. Error aggregation:")
	errs = []error{
		errors.New("first error"),
		errors.New("second error"),
		errors.New("third error"),
	}
	
	combined := errors.Join(errs...)
	fmt.Printf("Combined error: %v\n", combined)
	
	// Check if combined error contains specific error
	if errors.Is(combined, errors.New("second error")) {
		fmt.Println("Combined error contains 'second error'")
	}

	// 15. Error handling best practices
	fmt.Println("\n15. Error handling best practices:")
	
	// Good: Provide context
	err = fmt.Errorf("failed to process payment for order %s: %w", "ORD123", baseErr)
	fmt.Printf("Good error with context: %v\n", err)
	
	// Bad: Generic error (for demonstration)
	// err = errors.New("something failed")
	// fmt.Printf("Bad generic error: %v\n", err)
}

// Helper functions and types

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrPermissionDenied = errors.New("permission denied")
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func processUser(username string) error {
	users := map[string]bool{
		"admin": true,
		"user":  true,
	}
	
	if !users[username] {
		return ErrUserNotFound
	}
	
	if username == "admin" {
		return nil
	}
	
	return ErrPermissionDenied
}

func readFile(filename string) error {
	_, err := os.Open(filename)
	return err
}

func validateUser(email string, age int) []error {
	var errs []error
	
	if email == "" {
		errs = append(errs, errors.New("email is required"))
	} else if !contains(email, "@") {
		errs = append(errs, errors.New("invalid email format"))
	}
	
	if age < 0 {
		errs = append(errs, errors.New("age cannot be negative"))
	} else if age < 18 {
		errs = append(errs, errors.New("user must be at least 18 years old"))
	}
	
	return errs
}

func processOrder(orderID string) error {
	// Simulate validation error
	if orderID == "" {
		return errors.New("order ID is required")
	}
	
	// Simulate database error
	dbErr := errors.New("database connection failed")
	return fmt.Errorf("failed to process order %s: %w", orderID, dbErr)
}

func checkPermission(role, resource string) error {
	permissions := map[string][]string{
		"admin": {"read", "write", "delete"},
		"user":  {"read", "write"},
		"guest": {"read"},
	}
	
	userPerms, exists := permissions[role]
	if !exists {
		return ErrPermissionDenied
	}
	
	for _, perm := range userPerms {
		if perm == resource {
			return nil
		}
	}
	
	return ErrPermissionDenied
}

type ValidationError struct {
	Field   string
	Value   string
	Message string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("validation error for field '%s': %s", ve.Field, ve.Message)
}

func saveToDatabase(userID string) error {
	// Simulate successful save
	fmt.Printf("User %s saved to database\n", userID)
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	if !contains(email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	if age > 120 {
		return errors.New("age seems unrealistic")
	}
	return nil
}

func causePanic() {
	panic("something terrible happened!")
}

func processFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file: %v\n", closeErr)
		}
	}()
	
	// Process file content
	fmt.Printf("File %s processed successfully\n", filename)
	return nil
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || 
		(len(s) > len(substr) && 
			(s[:len(substr)] == substr || 
			 s[len(s)-len(substr):] == substr ||
			 findSubstring(s, substr))))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
