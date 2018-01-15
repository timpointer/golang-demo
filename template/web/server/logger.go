package main

import (
	"time"

	"github.com/sirupsen/logrus"
)

func testlogger() {
	for {
		log.WithFields(logrus.Fields{
			"channel": "web",
			"err":     "1111",
		}).Debug("message")
		log.WithFields(logrus.Fields{
			"channel": "web",
			"err":     "2222",
		}).Info("message")
		log.WithFields(logrus.Fields{
			"channel": "web",
			"err":     "3333",
		}).Warn("message")
		log.WithFields(logrus.Fields{
			"channel": "web",
			"err":     "4444",
		}).Error("message")
		time.Sleep(time.Second * 1)
	}
}
