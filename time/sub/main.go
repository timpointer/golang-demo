package main

import (
	"fmt"
	"time"
)

func main() {
	StartDate := time.Date(2017, 9, 1, 0, 0, 0, 0, time.UTC)
	EndDate := time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC)
	sub := EndDate.Sub(StartDate)
	fmt.Println("sub", sub)
	lastDate := StartDate.Add(-sub)
	fmt.Println("last", lastDate)
}
