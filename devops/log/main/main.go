package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("./logfile", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	logger := log.New(f, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")
	f.Close()
}
