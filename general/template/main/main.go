package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {
	var buf bytes.Buffer
	p := struct {
		Title string
		Body  string
	}{
		"tim", "This is a body",
	}
	t, _ := template.ParseFiles("../template/report.html")
	t.Execute(&buf, p)
	fmt.Print(&buf)
}
