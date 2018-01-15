package main

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	now := time.Now()
	resp, err := http.Get("http://localhost:8080/bar")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Get")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Get")
	}

	log.WithFields(log.Fields{
		"body":  string(body),
		"since": time.Since(now).Seconds(),
	}).Info("Get")
}
