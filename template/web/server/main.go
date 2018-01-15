package main

import (
	"fmt"
	"html"
	"log/syslog"
	"net/http"

	"github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

var log *logrus.Logger

func main() {
	app := &application{}
	http.Handle("/foo", app)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log = logrus.New()
	hook, err := lSyslog.NewSyslogHook("", "", syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
	}

	go testlogger()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
