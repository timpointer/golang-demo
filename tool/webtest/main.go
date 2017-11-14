package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for index := 0; index < 50; index++ {
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			multRequest(t)
		}(index)
	}
	wg.Wait()

}

func multRequest(t int) {
	for index := 0; index < 50; index++ {
		request()
		time.Sleep(time.Second)
		log.Println("t:", t, "-s:", index)
	}
}

func request() {
	res, err := http.Get("http://sr.metrowechat.com/mpg/getlist?openid=oJ9_hjuFCT_Z4zoQ2PfuCQA6_39k")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", len(robots))
}
