package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		log.WithFields(log.Fields{}).Info("start")
		time.Sleep(time.Second * 10)

		log.WithFields(log.Fields{
			"since": time.Since(now).Seconds(),
		}).Info("end")
	})

	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
