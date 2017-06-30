package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/test", handler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v \n", r.URL)
	log.Printf("%v \n", r.URL.RawQuery)
	query := r.URL.RawQuery
	qs := strings.Split(query, "_")
	log.Printf("%v \n", qs)
	log.Printf("%v \n", r.URL.EscapedPath)
	return
}
