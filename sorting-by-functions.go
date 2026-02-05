package main

import (
	"fmt"
	"sort"
	"time"
)

type Product struct {
	Name     string
	Price    float64
	Category string
	Rating   float64
}

func main() {
	fmt.Println("=== Sorting by Functions Examples ===")

	// 1. Sort by price
	fmt.Println("\n1. Sort by price:")
	products := []Product{
		{"Laptop", 999.99, "Electronics", 4.5},
		{"Phone", 699.99, "Electronics", 4.7},
		{"Book", 19.99, "Books", 4.2},
		{"Headphones", 149.99, "Electronics", 4.3},
		{"Coffee", 9.99, "Food", 4.8},
	}
	
	fmt.Printf("Original: %v\n", products)
	
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
	fmt.Printf("Sorted by price: %v\n", products)

	// 2. Sort by name
	fmt.Println("\n2. Sort by name:")
	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
	fmt.Printf("Sorted by name: %v\n", products)

	// 3. Sort by rating (descending)
	fmt.Println("\n3. Sort by rating (descending):")
	sort.Slice(products, func(i, j int) bool {
		return products[i].Rating > products[j].Rating
	})
	fmt.Printf("Sorted by rating: %v\n", products)

	// 4. Multi-criteria sorting
	fmt.Println("\n4. Multi-criteria sorting (category, then price):")
	sort.Slice(products, func(i, j int) bool {
		if products[i].Category != products[j].Category {
			return products[i].Category < products[j].Category
		}
		return products[i].Price < products[j].Price
	})
	fmt.Printf("Multi-criteria sorted: %v\n", products)

	// 5. Sort with custom comparator function
	fmt.Println("\n5. Sort with custom comparator:")
	compareProducts := func(a, b Product) int {
		if a.Category != b.Category {
			if a.Category < b.Category {
				return -1
			}
			return 1
		}
		if a.Price < b.Price {
			return -1
		} else if a.Price > b.Price {
			return 1
		}
		return 0
	}
	
	sort.Slice(products, func(i, j int) bool {
		return compareProducts(products[i], products[j]) < 0
	})
	fmt.Printf("Custom comparator sorted: %v\n", products)

	fmt.Println("All sorting by functions examples completed!")
}
