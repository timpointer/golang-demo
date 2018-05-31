package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"

	"cloud.google.com/go/logging"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	var opts struct {
		ProjectID string `short:"p" long:"project" description:"Google Cloud Platform Project ID" required:"true"`
		LogName   string `short:"l" long:"logname" description:"The name of the log to write to" default:"default"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(2)
	}

	// Check if Standard In is coming from a pipe
	fi, err := os.Stdin.Stat()
	if err != nil {
		errorf("Could not stat standard input: %v", err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		errorf("Nothing is piped in so there is nothing to log!")
	}

	// Creates a client.
	ctx := context.Background()

	lines := make(chan string)
	go func() {
		defer close(lines)
		// Read from Stdin and log it to Stdout and Stackdriver
		s := bufio.NewScanner(io.TeeReader(os.Stdin, os.Stdout))
		for s.Scan() {
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
			logger.Log(logging.Entry{Payload: line})
		case s := <-signals:
			fmt.Fprintf(os.Stderr, "Terminating program after receiving signal: %v\n", s)
			break loop
		}
	}

	// Closes the client and flushes the buffer to the Stackdriver Logging
	// service.
	if err := client.Close(); err != nil {
		errorf("Failed to close client: %v", err)
	}
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(2)
}
