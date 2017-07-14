package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:1520/gather/stream?storekey=37&custkey=912832")
	if err != nil {
		log.Println(err)
	}
	if res.StatusCode == http.StatusTooManyRequests {

		log.Fatal("err")
	}
	if res.StatusCode != 200 {
		log.Println(res.StatusCode, res.Status)
	}
	log.Println(res)
}
