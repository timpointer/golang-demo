package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Wrote %d bytes.\n", bytesWritten)
}
