package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Command Line Subcommands ===")

	// Check if any arguments provided
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	// Route to appropriate subcommand
	switch command {
	case "add":
		handleAdd(args)
	case "list":
		handleList(args)
	case "delete":
		handleDelete(args)
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func handleAdd(args []string) {
	fmt.Println("=== Add Command ===")
	
	addFlags := flag.NewFlagSet("add", flag.ExitOnError)
	name := addFlags.String("name", "", "Item name")
	value := addFlags.Int("value", 0, "Item value")
	
	addFlags.Parse(args)
	
	if *name == "" {
		fmt.Println("Error: --name is required")
		addFlags.PrintDefaults()
		return
	}
	
	fmt.Printf("Added item: %s with value: %d\n", *name, *value)
}

func handleList(args []string) {
	fmt.Println("=== List Command ===")
	
	listFlags := flag.NewFlagSet("list", flag.ExitOnError)
	all := listFlags.Bool("all", false, "List all items")
	count := listFlags.Int("count", 10, "Number of items to show")
	
	listFlags.Parse(args)
	
	if *all {
		fmt.Println("Listing all items:")
	} else {
		fmt.Printf("Listing %d items:\n", *count)
	}
	
	// Simulate listing items
	items := []string{"Item1", "Item2", "Item3", "Item4", "Item5"}
	limit := len(items)
	if !*all && *count < limit {
		limit = *count
	}
	
	for i := 0; i < limit; i++ {
		fmt.Printf("  %d. %s\n", i+1, items[i])
	}
}

func handleDelete(args []string) {
	fmt.Println("=== Delete Command ===")
	
	deleteFlags := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteFlags.Int("id", 0, "Item ID to delete")
	force := deleteFlags.Bool("force", false, "Force deletion")
	
	deleteFlags.Parse(args)
	
	if *id == 0 {
		fmt.Println("Error: --id is required")
		deleteFlags.PrintDefaults()
		return
	}
	
	if *force {
		fmt.Printf("Force deleted item with ID: %d\n", *id)
	} else {
		fmt.Printf("Deleted item with ID: %d\n", *id)
	}
}

func printUsage() {
	fmt.Println("Usage: program <command> [options]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  add     Add a new item")
	fmt.Println("  list    List items")
	fmt.Println("  delete  Delete an item")
	fmt.Println("  help    Show this help")
	fmt.Println("\nExamples:")
	fmt.Println("  program add --name=item1 --value=100")
	fmt.Println("  program list --all")
	fmt.Println("  program list --count=5")
	fmt.Println("  program delete --id=1 --force")
	fmt.Println("\nUse 'program help <command>' for command-specific help")
}
