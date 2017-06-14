package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func main() {
	tmpl := template.Must(template.ParseFiles("./template/template.html"))
	buf := bytes.NewBuffer([]byte{})
	err := tmpl.Execute(buf, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
}
