package log

import (
	"log"
	"os"
)

func logToFile(name string, message string) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	logger := log.New(f, "logger: ", log.Lshortfile)
	logger.Print(message)
	return nil
}
