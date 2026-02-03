package main

import "fmt"

func main() {
	fmt.Println("=== Switch Examples ===")

	// Basic switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with initialization
	switch grade := 85; {
	case grade >= 90:
		fmt.Println("Excellent")
	case grade >= 80:
		fmt.Println("Good")
	case grade >= 70:
		fmt.Println("Average")
	case grade >= 60:
		fmt.Println("Below Average")
	default:
		fmt.Println("Poor")
	}

	// Switch with multiple cases
	month := "January"
	switch month {
	case "January", "February", "December":
		fmt.Println("Winter")
	case "March", "April", "May":
		fmt.Println("Spring")
	case "June", "July", "August":
		fmt.Println("Summer")
	case "September", "October", "November":
		fmt.Println("Fall")
	default:
		fmt.Println("Invalid month")
	}

	// Switch without expression (like if-else chain)
	age := 25
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 20:
		fmt.Println("Teenager")
	case age < 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior")
	}

	// Switch with fallthrough
	fmt.Println("\nFallthrough example:")
	switch num := 2; num {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// Switch on type
	var value interface{} = "Hello"
	switch v := value.(type) {
	case string:
		fmt.Printf("String: %s\n", v)
	case int:
		fmt.Printf("Integer: %d\n", v)
	case float64:
		fmt.Printf("Float: %.2f\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}

	// Switch with function call
	fmt.Println("\nSwitch with function:")
	switch getDayOfWeek() {
	case "Saturday", "Sunday":
		fmt.Println("Weekend - Relax!")
	case "Monday":
		fmt.Println("Monday - Work starts")
	case "Friday":
		fmt.Println("Friday - TGIF!")
	default:
		fmt.Println("Weekday - Keep working")
	}
}

func getDayOfWeek() string {
	// Simulate getting current day of week
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	return days[3] // Return Thursday for this example
}
