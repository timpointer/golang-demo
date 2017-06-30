package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	values, err := url.ParseQuery("name=time&age=3")
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range values {
		for _, v := range value {
			fmt.Println(key, v)
		}
	}
	fmt.Println(values.Encode())

	s := ""
	sl := strings.Split(s, "_")
	fmt.Println("sl len ", len(sl))
	fmt.Println(sl[0])
}
