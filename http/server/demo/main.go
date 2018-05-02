package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helloworld"))
		log.WithFields(log.Fields{}).Info("helloworld bar")
	})

	s := &http.Server{Addr: "localhost:8080"}
	log.WithFields(log.Fields{}).Info("start serve")
	log.Fatal(s.ListenAndServe())
}
