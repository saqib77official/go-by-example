package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Time Formatting & Parsing ===")

	// Current time
	now := time.Now()

	// Common format layouts
	fmt.Printf("RFC1123: %s\n", now.Format(time.RFC1123))
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Kitchen: %s\n", now.Format(time.Kitchen))

	// Custom formatting
	fmt.Printf("Custom: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("US format: %s\n", now.Format("01/02/2006 03:04 PM"))
	fmt.Printf("ISO: %s\n", now.Format("2006-01-02T15:04:05Z07:00"))

	// Parse time from string
	timeStr := "2023-12-25 15:30:00"
	parsed, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed time: %v\n", parsed)

	// Parse different formats
	formats := []string{
		"Jan 2, 2006",
		"2006/01/02",
		"02-01-2006 15:04",
	}

	timeStrings := []string{
		"Dec 25, 2023",
		"2023/12/25",
		"25-12-2023 15:30",
	}

	for i, format := range formats {
		if t, err := time.Parse(format, timeStrings[i]); err == nil {
			fmt.Printf("Parsed '%s' with format '%s': %v\n", timeStrings[i], format, t)
		}
	}

	// Time zone formatting
	fmt.Printf("Time zone: %s\n", now.Format("MST"))
	fmt.Printf("Time zone offset: %s\n", now.Format("-0700"))

	// Only date or time
	fmt.Printf("Date only: %s\n", now.Format("2006-01-02"))
	fmt.Printf("Time only: %s\n", now.Format("15:04:05"))

	// Weekday and month names
	fmt.Printf("Weekday: %s\n", now.Format("Monday"))
	fmt.Printf("Month: %s\n", now.Format("January"))
}
