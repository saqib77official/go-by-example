package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Admin   bool   `json:"admin,omitempty"`
}

func main() {
	fmt.Println("=== JSON ===")

	// Marshal (Go to JSON)
	person := Person{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
		Admin: false,
	}

	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %s\n", jsonData)

	// Marshal with indentation
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pretty JSON:\n%s\n", prettyJSON)

	// Unmarshal (JSON to Go)
	jsonStr := `{"name":"Bob","age":30,"email":"bob@example.com","admin":true}`
	var decoded Person
	err = json.Unmarshal([]byte(jsonStr), &decoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded: %+v\n", decoded)

	// Working with arrays
	people := []Person{
		{Name: "Charlie", Age: 35, Email: "charlie@example.com"},
		{Name: "Diana", Age: 28, Email: "diana@example.com"},
	}
	
	peopleJSON, _ := json.Marshal(people)
	fmt.Printf("People JSON: %s\n", peopleJSON)

	// Dynamic JSON with interface
	var data map[string]interface{}
	json.Unmarshal([]byte(`{"name":"Eve","age":40,"skills":["Go","Python"]}`), &data)
	fmt.Printf("Dynamic: %+v\n", data)
}
