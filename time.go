package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Time ===")

	// Current time
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)

	// Time components
	fmt.Printf("Year: %d\n", now.Year())
	fmt.Printf("Month: %d\n", now.Month())
	fmt.Printf("Day: %d\n", now.Day())
	fmt.Printf("Hour: %d\n", now.Hour())
	fmt.Printf("Minute: %d\n", now.Minute())
	fmt.Printf("Second: %d\n", now.Second())

	// Weekday
	fmt.Printf("Weekday: %v\n", now.Weekday())

	// Creating specific time
	specific := time.Date(2023, 12, 25, 15, 30, 0, 0, time.UTC)
	fmt.Printf("Specific time: %v\n", specific)

	// Time arithmetic
	future := now.Add(24 * time.Hour)
	past := now.Add(-7 * 24 * time.Hour)
	fmt.Printf("Future (24h): %v\n", future)
	fmt.Printf("Past (7 days): %v\n", past)

	// Duration
	duration := future.Sub(now)
	fmt.Printf("Duration: %v\n", duration)

	// Time zones
	local := now.Local()
	utc := now.UTC()
	fmt.Printf("Local: %v\n", local)
	fmt.Printf("UTC: %v\n", utc)

	// Time comparison
	fmt.Printf("Future > Now: %t\n", future.After(now))
	fmt.Printf("Past < Now: %t\n", past.Before(now))
}
