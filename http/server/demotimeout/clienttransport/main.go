package main

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	now := time.Now()
	tr := &http.Transport{
		ResponseHeaderTimeout: 2 * time.Second,
		MaxIdleConns:          10,
		IdleConnTimeout:       30 * time.Second,
		DisableCompression:    true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://localhost:8080/bar")
	if err != nil {
		log.WithFields(log.Fields{
			"err":   err,
			"since": time.Since(now).Seconds(),
		}).Fatal("Get")
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
