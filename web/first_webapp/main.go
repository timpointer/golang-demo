// Package main provides ...
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("we.html"))
	fmt.Fprintf(w, "hello world, %s!", r.URL.Path[1:])
}
