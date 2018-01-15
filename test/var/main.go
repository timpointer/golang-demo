package main

import (
	"fmt"
)

func main() {
	var num int
	var err error
	if num, err = getNum(); err != nil {

		return
	}
	fmt.Println("num", num)
}

func getNum() (int, error) {
	return 23423, nil
}
