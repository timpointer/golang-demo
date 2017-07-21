package main

import (
	"bytes"
	"io"
	"log"
)

const debug = false

func main() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
	if debug {
		// ... use buf
	}
}

func f(out io.Writer) {
	if out != nil {
		log.Println("done")
		out.Write([]byte("done!\n"))
	}
}
