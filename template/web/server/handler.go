package main

import "net/http"

type application struct{}

func (a *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
