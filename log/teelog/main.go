// 	Copyright 2017, Google, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Command logpipe is a service that will let you pipe logs directly to Stackdriver Logging.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"

	flags "github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

func main() {

	var opts struct {
		LogName  string `short:"l" long:"logname" description:"The name of the log to write to" default:"default"`
		LogLevel string `short:"v" long:"loglevel" description:"The level of the log for debug application" default:"info"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("flags.Parse")
		os.Exit(2)
	}

	if opts.LogLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	}

	log.WithFields(log.Fields{}).Debug("log-service start")

	// Check if Standard In is coming from a pipe
	fi, err := os.Stdin.Stat()
	if err != nil {
		errorf("Could not stat standard input: %v", err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		errorf("Nothing is piped in so there is nothing to log!")
	}

	lines := make(chan string)
	go func() {
		defer close(lines)
		// Read from Stdin and log it to Stdout and Stackdriver
		s := bufio.NewScanner(io.TeeReader(os.Stdin, os.Stdout))
		for s.Scan() {
			log.WithFields(log.Fields{
				"Text": s.Text(),
			}).Debug("s.Text()")
			lines <- s.Text()
		}
		if err := s.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to scan input: %v\n", err)
		}
	}()

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)

loop:
	for {
		select {
		case line, ok := <-lines:
			if !ok {
				break loop
			}
			log.WithFields(log.Fields{
				"payload": line,
				"LogName": opts.LogName,
			}).Info("line")
		case s := <-signals:
			fmt.Fprintf(os.Stderr, "Terminating program after receiving signal: %v\n", s)
			break loop
		}
	}

}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(2)
}
