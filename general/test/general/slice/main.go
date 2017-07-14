package main

import "log"

func main() {
	s := make([]int, 0, 9)
	s = append(s, 5)
	log.Println(s)
}
