package main

import (
	"os"
	"text/template"
)

func main() {
	fmt.Println("=== Text Templates ===")

	// Basic template
	tmpl := `Hello {{.Name}}! You are {{.Age}} years old.`

	// Parse template
	t, err := template.New("greeting").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Data for template
	data := struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  30,
	}

	// Execute template
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	fmt.Println()

	// Template with condition
	condTmpl := `{{if .Admin}}Welcome Admin!{{else}}Welcome User!{{end}}`
	t2, _ := template.New("welcome").Parse(condTmpl)

	adminData := struct{ Admin bool }{Admin: true}
	t2.Execute(os.Stdout, adminData)
	fmt.Println()

	userData := struct{ Admin bool }{Admin: false}
	t2.Execute(os.Stdout, userData)
	fmt.Println()

	// Template with loop
	loopTmpl := `Items: {{range .}}- {{.}} {{end}}`
	t3, _ := template.New("items").Parse(loopTmpl)
	items := []string{"Apple", "Banana", "Cherry"}
	t3.Execute(os.Stdout, items)
	fmt.Println()
}
