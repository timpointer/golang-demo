package main

import (
	"log"
)

func main() {
	l := []int{}
	l = append(l, 3)
	l = append(l, 2)
	l = append(l, 1)
	l = append(l, 3)
	l = append(l, 2)
	log.Println("l", l)
	var ls []string = []string(nil)
	_ = ls

	var lint int = int(4)
	_ = lint
}
