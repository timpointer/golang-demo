package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	log.WithFields(log.Fields{
		"type": "log",
	}).Info("info")

	log.WithFields(log.Fields{
		"type": "log",
	}).Error("error")
}
