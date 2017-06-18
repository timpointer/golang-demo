package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("start")
	time.Sleep(5 * time.Second)
	fmt.Println("end")
	log.Fatal("error")
}
