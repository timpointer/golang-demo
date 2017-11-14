package main

import (
	"fmt"
	"time"
)

func main() {
	date1 := time.Date(2016, 1, 1, 5, 0, 0, 0, time.Local)
	date2 := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC).Add(-8 * time.Hour)
	fmt.Println(date1, date2)
	if date1.Before(date2) {
		fmt.Println("before")
	} else {
		fmt.Println("after")
	}
}
