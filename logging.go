package main

import (
	"log"
	"os"
)

func main() {
	fmt.Println("=== Logging ===")

	// Basic logging
	log.Println("This is a basic log message")
	log.Printf("Formatted log: %s is %d years old", "Alice", 25)

	// Log to file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error creating log file:", err)
	}
	defer file.Close()

	// Create new logger that writes to file
	fileLogger := log.New(file, "APP: ", log.LstdFlags)
	fileLogger.Println("This goes to file")
	fileLogger.Printf("User %s logged in", "Bob")

	// Log with different flags
	fmt.Println("\nDifferent log flags:")
	
	// Standard logger with flags
	log.SetFlags(log.LstdFlags)
	log.Println("With standard flags")

	// With microseconds
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("With microseconds")

	// With short file name
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("With short file name")

	// With long file name
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("With long file name")

	// Custom logger
	customLogger := log.New(os.Stdout, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)
	customLogger.Println("Custom logger message")

	// Log levels simulation
	fmt.Println("\nSimulated log levels:")
	
	info := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	warning := log.New(os.Stdout, "WARN: ", log.LstdFlags)
	error := log.New(os.Stdout, "ERROR: ", log.LstdFlags)
	
	info.Println("Application started")
	warning.Println("Low disk space")
	error.Println("Database connection failed")

	// Fatal logging (this would exit the program)
	fmt.Println("\nFatal logging (commented out to avoid exit):")
	// log.Fatal("This would terminate the program")

	// Panic logging (this would panic)
	fmt.Println("Panic logging (commented out to avoid panic):")
	// log.Panic("This would cause a panic")

	// Structured logging simulation
	fmt.Println("\nStructured logging simulation:")
	logStruct := func(level, message string, data map[string]interface{}) {
		log.Printf("[%s] %s %+v", level, message, data)
	}

	logStruct("INFO", "User action", map[string]interface{}{
		"user_id": 123,
		"action":  "login",
		"ip":      "192.168.1.1",
	})

	// Performance logging
	startTime := log.New(os.Stdout, "PERF: ", log.LstdFlags)
	startTime.Println("Starting operation")
	
	// Simulate work
	for i := 0; i < 3; i++ {
		log.Printf("Processing item %d", i+1)
	}
	
	startTime.Println("Operation completed")

	// Clean up
	os.Remove("app.log")
	fmt.Println("Log file cleaned up")
}
