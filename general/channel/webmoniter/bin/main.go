package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	time.Sleep(5 * time.Second)
	fmt.Println("**")
	time.Sleep(5 * time.Second)
	fmt.Println("****")
	time.Sleep(5 * time.Second)
	fmt.Println("*****")
	time.Sleep(5 * time.Second)
	fmt.Println("*******")
	time.Sleep(5 * time.Second)
	fmt.Println("************")
	fmt.Println("end")
}
