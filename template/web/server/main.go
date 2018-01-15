package main

import (
	"fmt"
	"html"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := &application{}
	http.Handle("/foo", app)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type application struct{}

func (a *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
