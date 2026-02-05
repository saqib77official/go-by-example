package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Epoch Time ===")

	// Current epoch time (seconds since Jan 1, 1970)
	now := time.Now()
	epochSeconds := now.Unix()
	epochMillis := now.UnixMilli()
	epochNanos := now.UnixNano()

	fmt.Printf("Current epoch (seconds): %d\n", epochSeconds)
	fmt.Printf("Current epoch (milliseconds): %d\n", epochMillis)
	fmt.Printf("Current epoch (nanoseconds): %d\n", epochNanos)

	// Convert epoch to time
	epochTime := time.Unix(epochSeconds, 0)
	fmt.Printf("Epoch to time: %v\n", epochTime)

	// Convert milliseconds to time
	millisTime := time.UnixMilli(epochMillis)
	fmt.Printf("Millis to time: %v\n", millisTime)

	// Common epoch values
	unixStart := time.Unix(0, 0)
	fmt.Printf("Unix epoch start: %v\n", unixStart)

	// Y2K epoch
	y2k := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	y2kEpoch := y2k.Unix()
	fmt.Printf("Y2K epoch: %d\n", y2kEpoch)

	// Calculate age from birth epoch
	birthEpoch := int64(946684800) // Jan 1, 2000
	age := (epochSeconds - birthEpoch) / (365 * 24 * 3600)
	fmt.Printf("Age from 2000: %d years\n", age)

	// Duration in different units
	hours := float64(epochSeconds) / 3600
	days := float64(epochSeconds) / (24 * 3600)
	years := float64(epochSeconds) / (365.25 * 24 * 3600)

	fmt.Printf("Hours since epoch: %.0f\n", hours)
	fmt.Printf("Days since epoch: %.0f\n", days)
	fmt.Printf("Years since epoch: %.1f\n", years)
}
