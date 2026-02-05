package main

import (
	"fmt"
	"testing"
)

// Function to test
func Add(a, b int) int {
	return a + b
}

func Multiply(a, b int) int {
	return a * b
}

// Basic test
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	expected := 12
	if result != expected {
		t.Errorf("Multiply(3, 4) = %d; want %d", result, expected)
	}
}

// Table-driven tests
func TestAddTable(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d+%d", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
			}
		})
	}
}

// Benchmark
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, i+1)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiply(i, i+1)
	}
}

// Example function
func ExampleAdd() {
	result := Add(2, 3)
	fmt.Println(result)
	// Output: 5
}

func main() {
	fmt.Println("=== Testing and Benchmarking ===")
	fmt.Println("Run tests with: go test")
	fmt.Println("Run benchmarks with: go test -bench=.")
	fmt.Println("Run specific test: go test -run TestAdd")
	fmt.Println("Run with verbose: go test -v")

	// Demonstrate manual testing
	fmt.Println("\nManual testing:")
	fmt.Printf("Add(2, 3) = %d (expected: 5)\n", Add(2, 3))
	fmt.Printf("Multiply(3, 4) = %d (expected: 12)\n", Multiply(3, 4))

	// Test helper functions
	fmt.Println("\nTest helpers:")
	fmt.Println("t.Errorf() - Report test failure")
	fmt.Println("t.Fatalf() - Report fatal failure")
	fmt.Println("t.Log() - Log information")
	fmt.Println("t.Skip() - Skip test")
	fmt.Println("t.Run() - Run subtest")
	fmt.Println("t.Parallel() - Run test in parallel")

	// Benchmark helpers
	fmt.Println("\nBenchmark helpers:")
	fmt.Println("b.N - Number of iterations")
	fmt.Println("b.ResetTimer() - Reset timer")
	fmt.Println("b.StopTimer() - Stop timer")
	fmt.Println("b.StartTimer() - Start timer")
	fmt.Println("b.ReportAllocs() - Report memory allocations")
}
