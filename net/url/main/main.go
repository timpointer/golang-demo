package main

import (
	"fmt"
	"log"
	"net/url"
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
}
