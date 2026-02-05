package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func main() {
	fmt.Println("=== XML ===")

	// Marshal (Go to XML)
	person := Person{
		Name:  "Alice",
		Age:   25,
		Email: "alice@example.com",
	}

	xmlData, err := xml.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("XML: %s\n", xmlData)

	// Marshal with indentation and header
	prettyXML, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Pretty XML:\n%s\n", prettyXML)

	// Unmarshal (XML to Go)
	xmlStr := `<person><name>Bob</name><age>30</age><email>bob@example.com</email></person>`
	var decoded Person
	err = xml.Unmarshal([]byte(xmlStr), &decoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded: %+v\n", decoded)

	// Working with nested structures
	type Company struct {
		XMLName xml.Name  `xml:"company"`
		Name    string    `xml:"name"`
		People []Person  `xml:"people>person"`
	}

	company := Company{
		Name: "Tech Corp",
		People: []Person{
			{Name: "Charlie", Age: 35},
			{Name: "Diana", Age: 28},
		},
	}

	companyXML, _ := xml.MarshalIndent(company, "", "  ")
	fmt.Printf("Company XML:\n%s\n", companyXML)
}
