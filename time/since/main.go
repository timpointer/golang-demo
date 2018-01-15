package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	time.Sleep(time.Millisecond)
	log.Println(time.Since(now).Nanoseconds())
}
